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
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Falha em connectar com o banco de dados", zap.Error(err))
	}
	if err := db.AutoMigrate(&Estoque{}); err != nil  {
		logger.Fatal("Falha em migrar o banco de dados", zap.Error(err))
	} 

	//createProduct(db, logger, 50,"caneta",10.0)
	//findID(db,logger, 1)
	updateById(db, logger, 2, "acerola", 10, 3.00)
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

func updateById(db *gorm.DB, logger *zap.Logger, id uint, product string, quantidade int, price float64) {
  var produtoAntigo Estoque
  // Buscar o produto pelo ID
  res := db.First(&produtoAntigo, id)
  if res.Error != nil {
    logger.Fatal("Falha em encontrar o produto", zap.Uint("ID", id), zap.Error(res.Error))
    return
  }
  // Atualizar o produto
  err := db.Model(&produtoAntigo).Updates(map[string]interface{}{"Product": product, "Quantidade": quantidade, "Price": price}).Error
	if(err != nil) {
		logger.Error("Falha em atualizar o produto", zap.Uint("ID", id), zap.Error(err))
	} else {
		logger.Info("Produto atualizado com sucesso", zap.Uint("ID", id),zap.String("Product", product), zap.Int("Quantidade", quantidade),zap.Float64("Price", price))
	}
}
type Estoque struct {
	gorm.Model
	Quantidade int
	Product    string
	Price      float64
}

