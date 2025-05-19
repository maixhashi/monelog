package mocks

import (
	"monelog/model"

	"github.com/stretchr/testify/mock"
)

// MockCSVHistoryValidator はCSV履歴バリデータのモック実装です
type MockCSVHistoryValidator struct {
	mock.Mock
}

func (m *MockCSVHistoryValidator) ValidateCSVHistorySaveRequest(request model.CSVHistorySaveRequest) error {
	args := m.Called(request)
	return args.Error(0)
}