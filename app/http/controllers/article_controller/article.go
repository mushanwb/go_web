package article_controller

import (
	"database/sql"
	"go_web/app/http/entity"
	"go_web/app/http/modles/article_modle"
	"go_web/pkg/logger"
	"go_web/pkg/route"
	"go_web/pkg/types"
	"net/http"
	"unicode/utf8"
)

type ArticleController struct {
}

func (*ArticleController) ArticlesShowHandler(w http.ResponseWriter, r *http.Request) {
	// 获取 url 上的 id 参数
	id := route.GetRouteVariable("id", r)

	// 读取对应文章的数据
	article, err := article_modle.Get(id)

	// 如果出现错误
	if err != nil {
		if err == sql.ErrNoRows {
			// 数据没找到
			w.WriteHeader(http.StatusNotFound)
			w.Write(entity.ReturnJson("文章不存在", nil))
		} else {
			// 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(entity.ReturnJson("查询文章失败", nil))
		}
	} else {
		// 4. 读取成功，显示文章
		logger.LogError(err)
		w.Write(entity.ReturnJson("请求成功", article))
	}
}

func (*ArticleController) ArticlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := article_modle.GetAll()

	if err != nil {
		// 数据库错误
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(entity.ReturnJson("查询文章列表失败", nil))
	} else {
		w.Write(entity.ReturnJson("文章列表查询成功", articles))
	}

}

func (*ArticleController) ArticlesStoreHandler(w http.ResponseWriter, r *http.Request) {
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

		article := article_modle.Article{
			Title: title,
			Body:  body,
		}

		err := article.Create()

		if err == nil {
			w.Write(entity.ReturnJson("插入文章成功", article))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(entity.ReturnJson("插入文章失败", nil))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(entity.ReturnJson("请求参数错误", errors))
	}
}

func (*ArticleController) ArticlesUpdateHandler(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)

	_, err := article_modle.Get(id)

	// 如果出现错误
	if err != nil {
		if err == sql.ErrNoRows {
			// 3.1 数据未找到
			w.WriteHeader(http.StatusNotFound)
			w.Write(entity.ReturnJson("文章不存在", nil))
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(entity.ReturnJson("文章查询失败", nil))
		}
	} else {
		// 4.1 表单验证
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")

		errors := validateArticleFormData(title, body)

		if len(errors) == 0 {
			// 4.2 表单验证通过，更新数据
			article := article_modle.Article{
				ID:    types.StringToInt64(id),
				Title: title,
				Body:  body,
			}
			i, err := article.Update()
			if err != nil {
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(entity.ReturnJson("文章更改失败", nil))
			}

			// √ 更新成功，跳转到文章详情页
			if i > 0 {
				w.Write(entity.ReturnJson("文章更改成功", article))
			} else {
				w.Write(entity.ReturnJson("文章没有任何改动", article))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(entity.ReturnJson("参数错误", errors))
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
