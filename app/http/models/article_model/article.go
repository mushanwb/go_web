package article_model

import (
	"go_web/app/http/models"
	"go_web/pkg/logger"
	"go_web/pkg/model"
	"go_web/pkg/types"
)

// Article 文章模型
type Article struct {
	models.BaseModel

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

func GetAll() ([]Article, error) {
	var article []Article

	if err := model.DB.Find(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (article *Article) Create() (err error) {
	result := model.DB.Create(&article)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Updates(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}

func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)

	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}
