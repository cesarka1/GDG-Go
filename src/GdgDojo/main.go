package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"go.uber.org/zap"
)

func main() {
	dsn := "host=localhost user=postgres password=default dbname=postgres port=5432 sslmode=disable"
	logger, _ := zap.NewProduction()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Erro em se conectar")
	}
	db.AutoMigrate(&Estoque{})
	createProduct(db, logger, 50,"sabato",150.0)


}

func createProduct (db *gorm.DB, logger *zap.Logger, quantidade int, product string, price float64) {
	estoque := Estoque {Quantidade: quantidade, Product: product, Price: price}
	result := db.Create(&estoque)
	if result.Error != nil {
		logger.Error("Erro na criação do produto")
	} else {
		logger.Info("Produto criado com sucesso")
	}
}

func searchProduct (db *gorm.DB, logger *zap.Logger, id uint) {
	
}

type Estoque struct {
	gorm.Model
	Quantidade int
	Product    string
	Price      float64
}

