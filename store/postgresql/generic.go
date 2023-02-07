package postgresql

import (
	"fmt"
	"github.com/adrien3d/base-api/store"
	"github.com/adrien3d/base-api/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create a generic model
func (db *PSQL) Create(c *store.Context, model store.Model) error {
	utils.EnsurePointer(model)
	store.EnsureGenericModel(model)

	if creator, ok := model.(store.BeforeCreator); ok {
		if err := creator.BeforeCreate(); err != nil {
			return errors.Wrap(err, "error in BeforeCreate")
		}
	}
	if creator, ok := model.(store.BeforeCreatorWithContext); ok {
		if err := creator.BeforeCreate(c); err != nil {
			return errors.Wrap(err, "error in BeforeCreatorWithContext")
		}
	}

	res := db.database.Create(&model)
	if res.Error != nil {
		logrus.WithError(res.Error).Errorln("cannot insert model")
		return errors.Wrap(res.Error, "cannot insert model")
	}

	return nil
}

// Find return a generic model
func (db *PSQL) Find(c *store.Context, filters bson.M, model store.Model, opts ...store.FindOption) error {
	utils.EnsurePointer(model)
	store.EnsureGenericModel(model)

	/*var sortQuery, sortValues string
	// apply sort
	if len(optValues.SortedFields) > 0 {
		sortBson := bson.D{}
		for i, sortedField := range optValues.SortedFields {
			sortedField.Field
		}
	}*/

	var filtersQuery string
	var filtersValue0 /*, filtersValue1*/ string
	var i int
	for key, value := range filters {
		//filtersValues = append(filtersValues, fmt.Sprint(value))
		if i == 0 {
			filtersQuery += key + " = ?"
			filtersValue0 = fmt.Sprint(value)
		} else {
			filtersQuery += " AND " + key + " = ?"
			//filtersValue1 = fmt.Sprint(value)
		}
		i++
	}
	db.database.Where(filtersQuery, filtersValue0).First(model) // find product with code D42)

	return nil
}

// FindAll return several generic models
func (db *PSQL) FindAll(c *store.Context, filters bson.M, results interface{}, opts ...store.FindOption) error {
	/*var sortQuery, sortValues string
	// apply sort
	if len(optValues.SortedFields) > 0 {
		sortBson := bson.D{}
		for i, sortedField := range optValues.SortedFields {
			sortedField.Field
		}
	}*/

	var filtersQuery string
	var filtersValues []string
	var i int
	for key, value := range filters {
		filtersValues = append(filtersValues, fmt.Sprint(value))
		if i == 0 {
			filtersQuery += key + " = ?"
		} else {
			filtersQuery += " AND " + key + " = ?"
		}
		i++
	}
	db.database.Where(filtersQuery, filtersValues).Find(results) // find product with code D42)

	return nil
}

// Update a generic model
func (db *PSQL) Update(c *store.Context, filter bson.M, model store.Model, opts ...store.UpdateOption) error {
	utils.EnsurePointer(model)
	store.EnsureGenericModel(model)

	db.database.Model(&model).Updates(model)

	return nil
}

// Delete a generic model
func (db *PSQL) Delete(c *store.Context, id string, model store.Model) error {
	utils.EnsurePointer(model)
	store.EnsureGenericModel(model)

	db.database.Delete(&model)

	return nil
}

// DeleteAll a generic model
func (db *PSQL) DeleteAll(c *store.Context, filter bson.M, model store.Model) (int64, error) {
	utils.EnsurePointer(model)
	store.EnsureGenericModel(model)

	db.database.Delete(&model)

	return 0, nil //res.DeletedCount, nil
}

func (db *PSQL) GetCollection(c *store.Context, model store.Model) *mongo.Collection {
	return nil
}
