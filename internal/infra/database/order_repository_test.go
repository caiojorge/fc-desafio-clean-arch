package database

import (
	"database/sql"
	"testing"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/stretchr/testify/suite"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	suite.Db.Exec("DELETE FROM orders")
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrderWhenSaveThenShouldSaveOrder() {
	order, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("Select id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestGivenMultipleOrdersWhenGetAllThenShouldReturnAllOrders() {
	// Arrange - criar múltiplos pedidos
	order1, err := entity.NewOrder("123", 10.0, 2.0)
	suite.NoError(err)
	suite.NoError(order1.CalculateFinalPrice())

	order2, err := entity.NewOrder("456", 20.0, 4.0)
	suite.NoError(err)
	suite.NoError(order2.CalculateFinalPrice())

	order3, err := entity.NewOrder("789", 30.0, 6.0)
	suite.NoError(err)
	suite.NoError(order3.CalculateFinalPrice())

	// Salvar os pedidos
	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order1)
	suite.NoError(err)
	err = repo.Save(order2)
	suite.NoError(err)
	err = repo.Save(order3)
	suite.NoError(err)

	// Act - buscar todos os pedidos
	orders, err := repo.GetAll()
	suite.NoError(err)

	// Assert - verificar se todos os pedidos foram retornados
	suite.Len(orders, 3)

	// Verificar se os pedidos estão corretos (pode não estar na mesma ordem)
	orderMap := make(map[string]entity.Order)
	for _, order := range orders {
		orderMap[order.ID] = order
	}

	// Verificar order1
	foundOrder1, exists := orderMap["123"]
	suite.True(exists)
	suite.Equal(order1.ID, foundOrder1.ID)
	suite.Equal(order1.Price, foundOrder1.Price)
	suite.Equal(order1.Tax, foundOrder1.Tax)
	suite.Equal(order1.FinalPrice, foundOrder1.FinalPrice)

	// Verificar order2
	foundOrder2, exists := orderMap["456"]
	suite.True(exists)
	suite.Equal(order2.ID, foundOrder2.ID)
	suite.Equal(order2.Price, foundOrder2.Price)
	suite.Equal(order2.Tax, foundOrder2.Tax)
	suite.Equal(order2.FinalPrice, foundOrder2.FinalPrice)

	// Verificar order3
	foundOrder3, exists := orderMap["789"]
	suite.True(exists)
	suite.Equal(order3.ID, foundOrder3.ID)
	suite.Equal(order3.Price, foundOrder3.Price)
	suite.Equal(order3.Tax, foundOrder3.Tax)
	suite.Equal(order3.FinalPrice, foundOrder3.FinalPrice)
}

func (suite *OrderRepositoryTestSuite) TestGivenEmptyDatabaseWhenGetAllThenShouldReturnEmptySlice() {
	// Arrange - limpar a base de dados
	suite.Db.Exec("DELETE FROM orders")

	// Act - buscar todos os pedidos
	repo := NewOrderRepository(suite.Db)
	orders, err := repo.GetAll()

	// Assert - verificar se retorna lista vazia
	suite.NoError(err)
	suite.NotNil(orders)
	suite.Len(orders, 0)
}
