package dblayer

import (
	"errors"
	"gomusic/models"
)

var ErrINVALIDPASSWORD = errors.New("Invalid password")

type DBlayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProduct(uint) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SingInUser(username, password string) (models.Customer, error)
	SignOutUserByID(int) error
	GetCustomerOrderByID(int) ([]models.Order, error)
}
