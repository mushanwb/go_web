package article_modle

import (
	"go_web/pkg/model"
	"go_web/pkg/types"
)

// Article 文章模型
type Article struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Get(idstr string) (Article, error) {
	var article Article

	id := types.StringToInt64(idstr)

	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}
