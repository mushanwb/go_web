package article_controller

import (
	"database/sql"
	"go_web/app/http/entity"
	"go_web/app/http/modles/article_modle"
	"go_web/pkg/logger"
	"go_web/pkg/route"
	"net/http"
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
