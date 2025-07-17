package usecase

import (
	"errors"
	"testing"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do repositório
type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Save(order *entity.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) GetAll() ([]entity.Order, error) {
	args := m.Called()
	return args.Get(0).([]entity.Order), args.Error(1)
}

func TestListOrderUseCaseExecuteSuccess(t *testing.T) {
	// Arrange
	mockRepo := new(MockOrderRepository)
	useCase := NewListOrderUseCase(mockRepo)

	// Mock data
	mockOrders := []entity.Order{
		{ID: "1", Price: 100.0, Tax: 10.0, FinalPrice: 110.0},
		{ID: "2", Price: 200.0, Tax: 20.0, FinalPrice: 220.0},
		{ID: "3", Price: 300.0, Tax: 30.0, FinalPrice: 330.0},
	}

	mockRepo.On("GetAll").Return(mockOrders, nil)

	// Act
	input := ListOrderInputDTO{} // O input não é usado na implementação atual
	result, err := useCase.Execute(input)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 3)

	// Verificar primeiro pedido
	assert.Equal(t, "1", result[0].ID)
	assert.Equal(t, 100.0, result[0].Price)
	assert.Equal(t, 10.0, result[0].Tax)
	assert.Equal(t, 110.0, result[0].FinalPrice)

	// Verificar segundo pedido
	assert.Equal(t, "2", result[1].ID)
	assert.Equal(t, 200.0, result[1].Price)
	assert.Equal(t, 20.0, result[1].Tax)
	assert.Equal(t, 220.0, result[1].FinalPrice)

	// Verificar terceiro pedido
	assert.Equal(t, "3", result[2].ID)
	assert.Equal(t, 300.0, result[2].Price)
	assert.Equal(t, 30.0, result[2].Tax)
	assert.Equal(t, 330.0, result[2].FinalPrice)

	mockRepo.AssertExpectations(t)
}

func TestListOrderUseCaseExecuteEmptyList(t *testing.T) {
	// Arrange
	mockRepo := new(MockOrderRepository)
	useCase := NewListOrderUseCase(mockRepo)

	// Mock data - lista vazia
	mockOrders := []entity.Order{}
	mockRepo.On("GetAll").Return(mockOrders, nil)

	// Act
	input := ListOrderInputDTO{}
	result, err := useCase.Execute(input)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
	assert.Equal(t, []ListOrderOutputDTO{}, result)

	mockRepo.AssertExpectations(t)
}

func TestListOrderUseCaseExecuteRepositoryError(t *testing.T) {
	// Arrange
	mockRepo := new(MockOrderRepository)
	useCase := NewListOrderUseCase(mockRepo)

	// Mock error
	expectedError := errors.New("database connection error")
	mockRepo.On("GetAll").Return([]entity.Order{}, expectedError)

	// Act
	input := ListOrderInputDTO{}
	result, err := useCase.Execute(input)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, expectedError, err)

	mockRepo.AssertExpectations(t)
}

func TestListOrderUseCaseExecuteFinalPriceCalculation(t *testing.T) {
	// Arrange
	mockRepo := new(MockOrderRepository)
	useCase := NewListOrderUseCase(mockRepo)

	// Mock data com diferentes valores para testar o cálculo
	mockOrders := []entity.Order{
		{ID: "1", Price: 50.0, Tax: 5.0, FinalPrice: 55.0},
		{ID: "2", Price: 75.25, Tax: 7.75, FinalPrice: 83.0},
		{ID: "3", Price: 1000.0, Tax: 100.0, FinalPrice: 1100.0},
	}

	mockRepo.On("GetAll").Return(mockOrders, nil)

	// Act
	input := ListOrderInputDTO{}
	result, err := useCase.Execute(input)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 3)

	// Verificar cálculo do preço final (Price + Tax)
	assert.Equal(t, 55.0, result[0].FinalPrice)
	assert.Equal(t, 83.0, result[1].FinalPrice)
	assert.Equal(t, 1100.0, result[2].FinalPrice)

	mockRepo.AssertExpectations(t)
}

func TestNewListOrderUseCase(t *testing.T) {
	// Arrange
	mockRepo := new(MockOrderRepository)

	// Act
	useCase := NewListOrderUseCase(mockRepo)

	// Assert
	assert.NotNil(t, useCase)
	assert.Equal(t, mockRepo, useCase.OrderRepository)
}
