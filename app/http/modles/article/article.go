package article

import (
	"go_web/pkg/model"
	"go_web/pkg/types"
)

// Article 文章模型
type Article struct {
	ID    int64
	Title string
	Body  string
}

func Get(idstr string) (Article, error) {
	var article Article

	id := types.StringToInt64(idstr)

	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}

	return article, nil
}
