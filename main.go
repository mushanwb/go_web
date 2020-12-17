package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"
)

// 使用 精准匹配 的 gorilla/mux
var router = mux.NewRouter()

// 数据库
var db *sql.DB

func initDB() {
	var err error
	config := mysql.Config{
		User:                 "homestead",
		Passwd:               "secret",
		Addr:                 "192.168.10.10:3306",
		Net:                  "tcp",
		DBName:               "go_web",
		AllowNativePasswords: true,
	}

	//准备数据库连接池
	db, err = sql.Open("mysql", config.FormatDSN())
	checkError(err)

	// 设置最大连接数
	db.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	db.SetConnMaxLifetime(5 * time.Minute)

	// 尝试连接，失败会报错
	err = db.Ping()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 创建 articles 数据表
func createTables() {
	createArticlesSQL := `CREATE TABLE IF NOT EXISTS articles(
    id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
    body longtext COLLATE utf8mb4_unicode_ci
); `

	_, err := db.Exec(createArticlesSQL)
	checkError(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(ReturnJson("首页访问", nil))
}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	// 获取 url 上的 id 参数
	id := getRouteVariable("id", r)

	// 读取对应文章的数据
	article, err := getArticleByID(id)

	// 如果出现错误
	if err != nil {
		if err == sql.ErrNoRows {
			// 数据没找到
			w.WriteHeader(http.StatusNotFound)
			w.Write(ReturnJson("文章不存在", nil))
		} else {
			// 数据库错误
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(ReturnJson("查询文章失败", nil))
		}
	} else {
		// 4. 读取成功，显示文章
		checkError(err)
		w.Write(ReturnJson("请求成功", article))
	}
}

// Article  对应一条文章数据
type Article struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// 返回的 json 数据格式
type JsonResult struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ReturnJson(message string, data interface{}) []byte {
	jsonData, _ := json.Marshal(JsonResult{message, data})
	return jsonData
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from articles")
	checkError(err)
	defer rows.Close()

	var articles []Article

	// 循环读取结果
	for rows.Next() {
		var article Article
		// 扫码每一行结果并赋值到一个 article 对象中
		err := rows.Scan(&article.ID, &article.Title, &article.Body)
		checkError(err)
		articles = append(articles, article)
	}

	// 检测遍历时是否发生错误
	err = rows.Err()
	checkError(err)
	w.Write(ReturnJson("文章列表查询成功", articles))
}

func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	// 使用这种方法，可以将接收的 application/json 数据转化为 map
	//param, _ := ioutil.ReadAll(r.Body)
	//m := make(map[string]interface{})
	//json.Unmarshal(param, &m)

	// 使用下面的方法，将只能接收 from-data 或者 application/x-www-form-urlencoded 格式数据
	// 接收不到 application/json 数据
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	errors := validateArticleFormData(title, body)

	// 检查是否有错误
	if len(errors) == 0 {
		lastInsertID, err := saveArticleToDB(title, body)
		if lastInsertID > 0 {
			w.Write(ReturnJson("插入文章成功", Article{lastInsertID, title, body}))
		} else {
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(ReturnJson("插入文章失败", nil))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ReturnJson("请求参数错误", errors))
	}
}

func saveArticleToDB(title string, body string) (int64, error) {
	// 变量初始化
	var (
		id   int64
		err  error
		rs   sql.Result
		stmt *sql.Stmt
	)

	// 1,获取一个 prepare 声明语句
	stmt, err = db.Prepare("INSERT INTO articles (title, body) VALUES(?,?)")
	// 例行检查错误
	if err != nil {
		return 0, err
	}

	// 2. 在此函数运行结束后关闭此语句，防止占用 SQL 连接
	defer stmt.Close()

	// 3. 执行请求，传参进入绑定的内容
	rs, err = stmt.Exec(title, body)
	if err != nil {
		return 0, err
	}

	// 4. 插入成功的话，会返回自增 ID
	if id, err = rs.LastInsertId(); id > 0 {
		return id, nil
	}

	return 0, err
}

func getRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

func getArticleByID(id string) (Article, error) {
	article := Article{}
	query := "SELECT * FROM articles WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)

	return article, err
}

func articlesUpdateHandler(w http.ResponseWriter, r *http.Request) {
	id := getRouteVariable("id", r)

	article, err := getArticleByID(id)

	// 如果出现错误
	if err != nil {
		if err == sql.ErrNoRows {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			w.Write(ReturnJson("文章不存在", nil))
		} else {
			// 3.2 数据库错误
			checkError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(ReturnJson("文章查询失败", nil))
		}
	} else {
		// 4.1 表单验证
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")

		errors := validateArticleFormData(title, body)

		if len(errors) == 0 {
			// 4.2 表单验证通过，更新数据

			query := "UPDATE articles SET title = ?, body = ? WHERE id = ?"
			rs, err := db.Exec(query, title, body, id)

			if err != nil {
				checkError(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(ReturnJson("文章更改失败", nil))
			}

			// √ 更新成功，跳转到文章详情页
			if n, _ := rs.RowsAffected(); n > 0 {
				w.Write(ReturnJson("文章更改成功", Article{article.ID, title, body}))
			} else {
				w.Write(ReturnJson("文章没有任何改动", article))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(ReturnJson("参数错误", errors))
		}
	}
}

func validateArticleFormData(title string, body string) map[string]string {
	errors := make(map[string]string)
	// 验证标题
	if title == "" {
		errors["title"] = "标题不能为空"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "标题长度需介于 3-40"
	}

	// 验证内容
	if body == "" {
		errors["body"] = "内容不能为空"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "内容长度需大于或等于 10 个字节"
	}

	return errors
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}

// 中间件,给每个请求头设置返回头数据格式
func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置返回头的数据格式
		w.Header().Set("Content-Type", "application/json")
		// 继续处理请求
		next.ServeHTTP(w, r)
	})
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	//初始化数据库
	initDB()
	// 初始化数据表
	createTables()

	// 后面的 Name 属性是给路由命名,和 laravel 路由的 name 属性差不多
	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")

	// 取 文章id 可以使用路由正则匹配
	router.HandleFunc("/articles/{id:[0-9]+}", articlesShowHandler).Methods("GET").Name("articles.show")

	// 同名的路由,根据请求的方式不同，选择进入不同的函数
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/{id:[0-9]+}", articlesUpdateHandler).Methods("PUT").Name("articles.update")

	// 自定义 404 页面
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// 中间件：强制内容类型为 JSON
	router.Use(forceHTMLMiddleware)

	// 通过命名路由获取 URL 示例
	//homeURL, _ := router.Get("home").URL()
	//fmt.Println("homeURL: ", homeURL)
	//articleURL, _ := router.Get("articles.show").URL("id", "23")
	//fmt.Println("articleURL: ", articleURL)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
