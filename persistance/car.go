package persistance

import (
	"github.com/ivan-kostko/car-service/models"
	log "github.com/ivan-kostko/car-service/pkg/logger"
	"gorm.io/gorm"
)

type Car struct {
	db  *gorm.DB
	log log.FieldLogger
}

func (cs *Car) WithDB(db *gorm.DB) *Car {

	if err := db.AutoMigrate(&models.CarEntity{}); err != nil {
		panic("Failed to migrate CarEntity model")
	}

	cs.db = db
	return cs
}

func (cs *Car) WithLogger(l log.FieldLogger) *Car {
	cs.log = l
	return cs
}

func (cs *Car) Upsert(m *models.CarEntity) error {
	return cs.db.Create(m).Error
}

func (cs *Car) GetById(id int) (*models.CarEntity, error) {
	c := models.CarEntity{}

	l := cs.log.WithFields(map[string]interface{}{"method": "GetById", "id": id})
	l.Debugf("Calling Db to get first car", id)
	db := cs.db.First(&c, id)
	cs.log.Debugf("Db returned model %+v", c)

	if db.Error != nil {
		cs.log.Errorf("DataBase returned error: %e", db.Error)
	}
	return &(c), db.Error
}
