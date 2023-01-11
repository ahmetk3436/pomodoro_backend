package post

import (
	"math"
	"strconv"
	"time"

	"github.com/ahmetk3436/pkg/model"
	"gorm.io/gorm"
)

// Repository defines a struct for working with posts in the database
type Repository struct {
	db *gorm.DB
}

// NewRepository returns a new instance of Repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// FindByID retrieves an examdate by its ID
func (r *Repository) FindByID(id uint) (*model.ExamDate, error) {
	var examDate model.ExamDate
	result := r.db.Where("id = ?", id).First(&examDate)
	if result.Error != nil {
		return nil, result.Error
	}
	duration := examDate.ExamDate.Sub(time.Now())
	examDate.Duration = duration

	date := (strconv.Itoa(int(duration.Hours()/(24*30))) + "-" + strconv.Itoa(int(duration.Hours()/24)) + "-" + strconv.Itoa(int(int(duration.Hours())%24)) + "-" + strconv.Itoa(int(math.Mod(duration.Minutes(), 60.000000))) + "-" + strconv.Itoa(int(int(duration.Seconds())%60)))
	examDate.DurationTime = date
	return &examDate, nil
}

// Save ...
func (p *Repository) Save(examDate *model.ExamDate) (*model.ExamDate, error) {
	err := p.db.Create(&examDate).Error
	return examDate, err
}

// Migrate creates the tables for the corresponding model
func (r *Repository) Migrate() error {
	return r.db.AutoMigrate(&model.ExamDate{})
}
