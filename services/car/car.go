package car

import (
	"errors"
	"fmt"

	"github.com/ivan-kostko/car-service/models"
	log "github.com/ivan-kostko/car-service/pkg/logger"
)

// Errors
var (
	ErrInvalidId = errors.New("car id should be in range 1-2^64")
)

type CarByIdGetter interface {
	GetById(int) (*models.CarEntity, error)
}

type CarUpserter interface {
	Upsert(m *models.CarEntity) error
}

type IntNexter interface {
	Next() int
}

type CarService struct {
	log        log.FieldLogger
	byIdGetter CarByIdGetter
	upserter   CarUpserter
	nexter     IntNexter
}

func NewService() *CarService {
	return &CarService{log: log.DefaultLogger}
}

// Sets car byIdGetter
func (s *CarService) WithByIdGetter(in CarByIdGetter) *CarService {
	s.byIdGetter = in
	return s
}

// Sets car upserter
func (s *CarService) WithUpserted(in CarUpserter) *CarService {
	s.upserter = in
	return s
}

// Sets car Persistance
func (s *CarService) WithLogger(in log.FieldLogger) *CarService {
	s.log = in
	return s
}

// Sets car Nexter
func (s *CarService) WithNexter(in IntNexter) *CarService {
	s.nexter = in
	return s
}

// Returns CarEntity by id
func (cs *CarService) GetCarById(id int) (*models.CarEntity, error) {

	log := cs.log.WithFields(map[string]interface{}{"method": "GetCarById", "id": id})

	log.Debug("Starting")

	return cs.byIdGetter.GetById(id)
}

// Creates a new CarEntity entry in persistance.
// It overrides CarEntuty.Entity.ID with next sequence Id
// It is not concurrent safe and panics on nil model
func (cs *CarService) CreateCar(m *models.CarEntity) error {

	log := cs.log.WithFields(map[string]interface{}{"method": "CreateCar", "model": m})

	log.Debug("Assigning next ID")
	m.Entity.ID = int(cs.nexter.Next())
	log.Debugf("Assigned Id %i", m.Entity.ID)

	log.Debugf("Upserting model #+v", m)
	err := cs.upserter.Upsert(m)
	if err != nil {
		err := fmt.Errorf("database error: %e", err)
		log.Error(err)
	}

	return err
}
