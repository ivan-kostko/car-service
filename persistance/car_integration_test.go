package persistance_test

import (
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ivan-kostko/car-service/models"
	"github.com/ivan-kostko/car-service/persistance"
	log "github.com/ivan-kostko/car-service/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestCreateAndGet(t *testing.T) {
	testCases := []struct {
		Alias                   string
		ModelIn                 *models.CarEntity
		ExpectedErrorOnCreation error
		ExpectedId              int
		ExpectedModelOut        *models.CarEntity
		ExpectedErrorOnGetting  error
	}{
		{
			Alias:                   "Nil model creation",
			ModelIn:                 nil,
			ExpectedErrorOnCreation: errors.New("invalid value, should be pointer to struct or slice"),
			ExpectedModelOut:        &models.CarEntity{},
			ExpectedErrorOnGetting:  errors.New("record not found"),
		},
		{
			Alias:                   "Empty model creation",
			ModelIn:                 &models.CarEntity{},
			ExpectedErrorOnCreation: nil,
			ExpectedId:              1,
			ExpectedModelOut: &models.CarEntity{
				Entity: models.Entity{ID: 1},
				Car:    models.Car{},
			},
		},
		{
			Alias: "Empty model with ID creation",
			ModelIn: &models.CarEntity{
				Entity: models.Entity{ID: 10},
				Car:    models.Car{},
			},
			ExpectedErrorOnCreation: nil,
			ExpectedId:              10,
			ExpectedModelOut: &models.CarEntity{
				Entity: models.Entity{ID: 10},
				Car:    models.Car{},
			},
		},
		{
			Alias: "Empty model with same ID creation",
			ModelIn: &models.CarEntity{
				Entity: models.Entity{ID: 10},
				Car:    models.Car{},
			},
			ExpectedErrorOnCreation: nil,
			ExpectedId:              10,
			ExpectedModelOut: &models.CarEntity{
				Entity: models.Entity{ID: 10},
				Car:    models.Car{},
			},
		},
		{
			Alias:                   "New empty model creation",
			ModelIn:                 &models.CarEntity{},
			ExpectedErrorOnCreation: nil,
			ExpectedId:              1,
			ExpectedModelOut: &models.CarEntity{
				Entity: models.Entity{ID: 1},
				Car:    models.Car{},
			},
			ExpectedErrorOnGetting: nil,
		},
		{
			Alias: "Not empty model creation",
			ModelIn: &models.CarEntity{
				Entity: models.Entity{ID: 11},
				Car: models.Car{
					Model:           "The Model",
					Engine:          "The Engine",
					Infotainment:    "The Infotainment",
					Interrior:       "The Interrior",
					CurrentLocation: "55:82",
				},
			},
			ExpectedErrorOnCreation: nil,
			ExpectedId:              11,
			ExpectedModelOut: &models.CarEntity{
				Entity: models.Entity{ID: 11},
				Car: models.Car{
					Model:           "The Model",
					Engine:          "The Engine",
					Infotainment:    "The Infotainment",
					Interrior:       "The Interrior",
					CurrentLocation: "55:82",
				},
			},
			ExpectedErrorOnGetting: nil,
		},
	}

	for _, testCase := range testCases {
		testFn := func(t *testing.T) {

			fileName := strings.ReplaceAll(testCase.Alias, " ", "_") + time.Now().GoString() + ".db"
			defer func() {
				os.Remove(fileName)
			}()
			service := (&persistance.Car{}).
				WithDB(persistance.DefaultSqliteDb(fileName)).
				WithLogger(log.DefaultLogger)

			actualError := service.Upsert(testCase.ModelIn)

			t.Logf("ModelIn after Upsert : %+v", testCase.ModelIn)

			if testCase.ModelIn != nil {
				assert.Equal(t, testCase.ExpectedId, testCase.ModelIn.ID)
			}

			assert.Equal(t, testCase.ExpectedErrorOnCreation, actualError)

			actualModelOut, actualErrorOnGetting := service.GetById(testCase.ExpectedId)

			assert.Equal(t, testCase.ExpectedModelOut, actualModelOut)
			assert.Equal(t, testCase.ExpectedErrorOnGetting, actualErrorOnGetting)
		}
		t.Run(testCase.Alias, testFn)

	}
}
