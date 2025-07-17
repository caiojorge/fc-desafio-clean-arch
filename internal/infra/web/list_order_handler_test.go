package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const testEndpoint = "/orders"

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

func TestWebListOrderHandlerGetAllSuccess(t *testing.T) {
	// Arrange
	mockRepo := new(MockOrderRepository)
	handler := NewWebListOrderHandler(mockRepo)

	// Mock data
	mockOrders := []entity.Order{
		{ID: "1", Price: 100.0, Tax: 10.0, FinalPrice: 110.0},
		{ID: "2", Price: 200.0, Tax: 20.0, FinalPrice: 220.0},
	}

	mockRepo.On("GetAll").Return(mockOrders, nil)

	// Create request and response recorder
	req := httptest.NewRequest(http.MethodGet, testEndpoint, nil)
	rr := httptest.NewRecorder()

	// Act
	handler.GetAll(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verificar se o JSON retornado está correto
	var result []usecase.ListOrderOutputDTO
	err := json.Unmarshal(rr.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Len(t, result, 2)

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

	mockRepo.AssertExpectations(t)
}

func TestWebListOrderHandlerGetAllEmptyList(t *testing.T) {
	// Arrange
	mockRepo := new(MockOrderRepository)
	handler := NewWebListOrderHandler(mockRepo)

	// Mock data - lista vazia
	mockOrders := []entity.Order{}
	mockRepo.On("GetAll").Return(mockOrders, nil)

	// Create request and response recorder
	req := httptest.NewRequest(http.MethodGet, testEndpoint, nil)
	rr := httptest.NewRecorder()

	// Act
	handler.GetAll(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verificar se o JSON retornado está correto
	var result []usecase.ListOrderOutputDTO
	err := json.Unmarshal(rr.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Len(t, result, 0)

	mockRepo.AssertExpectations(t)
}

func TestWebListOrderHandlerGetAllRepositoryError(t *testing.T) {
	// Arrange
	mockRepo := new(MockOrderRepository)
	handler := NewWebListOrderHandler(mockRepo)

	// Mock error
	expectedError := errors.New("database connection error")
	mockRepo.On("GetAll").Return([]entity.Order{}, expectedError)

	// Create request and response recorder
	req := httptest.NewRequest(http.MethodGet, testEndpoint, nil)
	rr := httptest.NewRecorder()

	// Act
	handler.GetAll(rr, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "database connection error")

	mockRepo.AssertExpectations(t)
}

func TestWebListOrderHandlerGetAllHTTPMethod(t *testing.T) {
	// Arrange
	mockRepo := new(MockOrderRepository)
	handler := NewWebListOrderHandler(mockRepo)

	mockOrders := []entity.Order{
		{ID: "1", Price: 100.0, Tax: 10.0, FinalPrice: 110.0},
	}

	mockRepo.On("GetAll").Return(mockOrders, nil)

	// Test with different HTTP methods
	methods := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}

	for _, method := range methods {
		t.Run("Method_"+method, func(t *testing.T) {
			req := httptest.NewRequest(method, testEndpoint, nil)
			rr := httptest.NewRecorder()

			// Act
			handler.GetAll(rr, req)

			// Assert - o handler deve funcionar independente do método HTTP
			// (isso é responsabilidade do roteador, não do handler)
			assert.Equal(t, http.StatusOK, rr.Code)
		})
	}

	mockRepo.AssertExpectations(t)
}

func TestWebListOrderHandlerGetAllJSONResponse(t *testing.T) {
	// Arrange
	mockRepo := new(MockOrderRepository)
	handler := NewWebListOrderHandler(mockRepo)

	mockOrders := []entity.Order{
		{ID: "1", Price: 100.0, Tax: 10.0, FinalPrice: 110.0},
	}

	mockRepo.On("GetAll").Return(mockOrders, nil)

	// Create request and response recorder
	req := httptest.NewRequest(http.MethodGet, testEndpoint, nil)
	rr := httptest.NewRecorder()

	// Act
	handler.GetAll(rr, req)

	// Assert - verificar se o Content-Type está correto
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var result []usecase.ListOrderOutputDTO
	err := json.Unmarshal(rr.Body.Bytes(), &result)
	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestNewWebListOrderHandler(t *testing.T) {
	// Arrange
	mockRepo := new(MockOrderRepository)

	// Act
	handler := NewWebListOrderHandler(mockRepo)

	// Assert
	assert.NotNil(t, handler)
	assert.Equal(t, mockRepo, handler.OrderRepository)
}
