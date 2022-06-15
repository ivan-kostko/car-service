package car_test

import (
	"testing"

	"github.com/ivan-kostko/car-service/models"
	log "github.com/ivan-kostko/car-service/pkg/logger"
	"github.com/ivan-kostko/car-service/pkg/sequence"
	"github.com/ivan-kostko/car-service/services/car"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCarByIdGetter struct {
	mock.Mock
}

func (m *mockCarByIdGetter) GetById(id int) (*models.CarEntity, error) {
	args := m.Called(id)
	return args[0].(*models.CarEntity), args.Error(1)
}

func Test_GetCarById(t *testing.T) {

	type mockCase struct {
		ExpectedCallId int
		ReturnModel    *models.CarEntity
		ReturnError    error
	}

	type expectedOutput struct {
		Model *models.CarEntity
		Error error
	}

	testCases := []struct {
		Alias          string
		Id             int
		Mock           *mockCase
		ExpectedOutput expectedOutput
	}{
		{
			Alias: "Id = 10",
			Id:    10,
			// Should NOT fail as invalid Id
			Mock: &mockCase{
				ExpectedCallId: 10,
				ReturnModel:    &models.CarEntity{},
				ReturnError:    nil,
			},
			ExpectedOutput: expectedOutput{
				Model: &models.CarEntity{},
				Error: nil,
			},
		},
	}

	for _, testCase := range testCases {
		testFn := func(t *testing.T) {

			var mockedByIdGetter = (&mockCarByIdGetter{})
			if testCase.Mock != nil {
				mockedByIdGetter.On("GetById", testCase.Mock.ExpectedCallId).
					Return(testCase.Mock.ReturnModel, testCase.Mock.ReturnError)
			}

			service := car.NewService().
				WithByIdGetter(mockedByIdGetter).
				WithLogger(log.DefaultLogger).
				WithNexter(sequence.NewIntSequence(0))

			actualModel, actualError := service.GetCarById(testCase.Id)

			if testCase.Mock != nil {
				mockedByIdGetter.AssertExpectations(t)
			} else {
				mockedByIdGetter.AssertNotCalled(t, "GetById")
			}

			assert.Equal(t, testCase.ExpectedOutput.Model, actualModel)
			assert.Equal(t, testCase.ExpectedOutput.Error, actualError)
		}

		t.Run(testCase.Alias, testFn)
	}
}
