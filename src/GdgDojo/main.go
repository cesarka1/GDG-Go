package main

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// configurando o logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// conectando db
	dsn := "host=localhost user=postgres password=default dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Falha em connectar com o banco de dados", zap.Error(err))
	}
	if err := db.AutoMigrate(&Estoque{}); err != nil  {
		logger.Fatal("Falha em migrar o banco de dados", zap.Error(err))
	} 

	//createProduct(db, logger, 50,"sabato",150.0)
	findID(db,logger, 5)
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

// função para buscar o id de um produto
func findID(db *gorm.DB, logger *zap.Logger, id uint) {
	var produto Estoque

	res := db.First(&produto, id)
	if(res.Error != nil) {
		logger.Fatal("Falha ao procurar produto", zap.Uint("ID", id), zap.Error(res.Error))
		return
	}
	logger.Info("Produto encontrado", zap.Uint("ID", produto.ID), zap.String("Product",produto.Product))
	//fmt.Printf("Produto ID: %d\n", produto.ID)
}

type Estoque struct {
	gorm.Model
	Quantidade int
	Product    string
	Price      float64
}

