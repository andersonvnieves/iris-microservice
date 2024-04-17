package repositories

import (
	"main/models"
)

type IrisRepository interface {
	ListAll() ([]models.Iris, error)
	Add(iris models.Iris) (models.Iris, error)
	Delete(iris models.Iris) (models.Iris, error)
	GetById(id int) (models.Iris, error)
	Update(iris models.Iris) (models.Iris, error)
}
