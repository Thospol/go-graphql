package repository

import (
	"github.com/thospol/go-graphql/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Repository repository interface
type Repository interface {
	Create(database *gorm.DB, i interface{}) error
	Update(database *gorm.DB, i interface{}) error
	Delete(database *gorm.DB, i interface{}) error
	FindByID(database *gorm.DB, id uint, i interface{}) error
	BulkInsert(database *gorm.DB, sliceValue interface{}) error
	BulkUpdate(database *gorm.DB, sliceValue interface{}) error
}

type repository struct{}

// NewRepository new repository
func NewRepository() Repository {
	return &repository{}
}

// Create create record database
func (r *repository) Create(database *gorm.DB, i interface{}) error {
	if m, ok := i.(model.ModelInterface); ok {
		m.Stamp()
	}

	if err := database.Omit(clause.Associations).Create(i).Error; err != nil {
		return err
	}

	return nil
}

// Update update record database
func (r *repository) Update(database *gorm.DB, i interface{}) error {
	if m, ok := i.(model.ModelInterface); ok {
		m.UpdateStamp()
	}

	if err := database.Omit(clause.Associations).Save(i).Error; err != nil {
		return err
	}

	return nil
}

// Delete delete record database
func (r *repository) Delete(database *gorm.DB, i interface{}) error {
	if m, ok := i.(model.ModelInterface); ok {
		m.DeleteStamp()
	}

	if err := database.Omit(clause.Associations).Delete(i).Error; err != nil {
		return err
	}

	return nil
}

// FindByID find by id record database
func (r *repository) FindByID(database *gorm.DB, id uint, i interface{}) error {
	if err := database.First(i, id).Error; err != nil {
		return err
	}

	return nil
}

// BulkInsert bulk insert into database
func (r *repository) BulkInsert(database *gorm.DB, sliceValue interface{}) error {
	if result := database.Create(sliceValue); result.Error != nil {
		return result.Error
	}

	return nil
}

// BulkUpdate bulk update into database
func (r *repository) BulkUpdate(database *gorm.DB, sliceValue interface{}) error {
	if result := database.Save(sliceValue); result.Error != nil {
		return result.Error
	}

	return nil
}
