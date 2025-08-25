package main

import (
	"log"

	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/database"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/models"
)

func main() {
	db, err := database.ConnectMysql()

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.AddressModel{}, &models.BankAccountModel{}, &models.CartModel{}, &models.CategoryModel{}, &models.OrderModel{}, &models.OrderItemModel{}, &models.PaymentModel{}, &models.ProductModel{}, &models.UserModel{})

	if err != nil {
		panic(err)
	}

	log.Println("Database migrated successfully")
}
