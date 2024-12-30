package repository_test

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
	"vehicle-registration-manager/internal/adapters/repository"
	"vehicle-registration-manager/internal/adapters/repository/configs"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/pkg/tracer"

	"github.com/stretchr/testify/suite"
)

type VehicleRepositoryTestSuite struct {
	suite.Suite
	db       *sql.DB
	repo     *repository.VehicleRepository
	mockDB   *configs.MockDatabaseConfig
	tracer   *tracer.Tracer
	sqlmockX sqlmock.Sqlmock
}

func (suite *VehicleRepositoryTestSuite) SetupSubTest() {
	suite.mockDB = new(configs.MockDatabaseConfig)
	suite.db, suite.sqlmockX, _ = sqlmock.New()
	suite.mockDB.On("GetDB").Return(suite.db)
	suite.repo = repository.NewVehicleRepository(suite.mockDB)
	suite.tracer = tracer.NewFakeTracer()
}

func (suite *VehicleRepositoryTestSuite) TearDownSubTest() {
	suite.mockDB.AssertExpectations(suite.T())
}

func TestVehicleRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(VehicleRepositoryTestSuite))
}

var columns = []string{"id", "brand", "model", "year", "color", "price"}

func (suite *VehicleRepositoryTestSuite) TestFindByID() {
	vehicleID := "1"
	findByID := `SELECT id, brand, model, year, color, price FROM vehicles WHERE id=\$1`
	suite.Run("Should return success", func() {
		row := suite.sqlmockX.NewRows(columns).
			AddRow(vehicleID, "Toyota", "Corolla", 2020, "Blue", 20000)
		suite.sqlmockX.ExpectPrepare(findByID).
			ExpectQuery().
			WithArgs(vehicleID).
			WillReturnRows(row).
			RowsWillBeClosed()

		vehicle, err := suite.repo.FindByID(suite.tracer, vehicleID)
		suite.NoError(err)
		suite.Equal(vehicleID, vehicle.ID)
	})

	suite.Run("Should return success - not found vehicle", func() {
		suite.sqlmockX.ExpectPrepare(findByID).
			ExpectQuery().
			WithArgs(vehicleID).
			WillReturnError(sql.ErrNoRows).
			RowsWillBeClosed()

		vehicle, err := suite.repo.FindByID(suite.tracer, vehicleID)
		suite.NoError(err)
		suite.Empty(vehicle)
	})

	suite.Run("Should return error- when prepare", func() {
		expected := errors.New("error in prepare")
		suite.sqlmockX.ExpectPrepare(findByID).WillReturnError(expected)

		vehicle, err := suite.repo.FindByID(suite.tracer, vehicleID)
		suite.Error(err)
		suite.Empty(vehicle)
		suite.Equal(expected, err)
	})

	suite.Run("Should return error- when scan", func() {
		expected := errors.New("error in Scan")
		suite.sqlmockX.ExpectPrepare(findByID).
			ExpectQuery().
			WithArgs(vehicleID).
			WillReturnError(expected)

		vehicle, err := suite.repo.FindByID(suite.tracer, vehicleID)
		suite.Error(err)
		suite.Empty(vehicle)
		suite.Equal(expected, err)
	})
}

func (suite *VehicleRepositoryTestSuite) TestFindAll() {
	findAll := `SELECT id, brand, model, year, color, price FROM vehicles`
	suite.Run("Should return success", func() {
		row := suite.sqlmockX.NewRows(columns).
			AddRow("1", "Toyota", "Corolla", 2020, "Blue", 20000).
			AddRow("2", "Toyota", "Corolla", 2020, "Green", 20000)
		suite.sqlmockX.ExpectQuery(findAll).WillReturnRows(row)

		vehicles, err := suite.repo.FindAll(suite.tracer)
		suite.NoError(err)
		suite.Len(vehicles, 2)
	})

	suite.Run("Should return error- when query", func() {
		expected := errors.New("error in query")
		suite.sqlmockX.ExpectQuery(findAll).WillReturnError(expected)

		vehicles, err := suite.repo.FindAll(suite.tracer)
		suite.Error(err)
		suite.Nil(vehicles)
		suite.Equal(expected, err)
	})

	suite.Run("Should return error - when scan to rows", func() {
		expected := errors.New("failed to scan")

		row := suite.sqlmockX.NewRows(columns).
			AddRow(1, "Toyota", "Corolla", 2020.66, "Blue", "44")

		suite.sqlmockX.ExpectQuery(findAll).WillReturnRows(row)

		vehicles, err := suite.repo.FindAll(suite.tracer)
		suite.Error(err)
		suite.Nil(vehicles)
		suite.Equal(expected, err)
	})

}

func (suite *VehicleRepositoryTestSuite) TestSave() {
	const insert = `INSERT INTO vehicles \(id, brand, model, year, color, price\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\)`

	domain := domains.Vehicle{
		ID:    "1",
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2020,
		Color: "Blue",
		Price: 20000.65,
	}
	suite.Run("Should return success", func() {
		suite.sqlmockX.ExpectPrepare(insert).
			ExpectExec().
			WithArgs("1", "Toyota", "Corolla", 2020, "Blue", 20000.65).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := suite.repo.Save(suite.tracer, domain)
		suite.NoError(err)
	})

	suite.Run("Should return error - when prepare", func() {
		expected := errors.New("error in prepare")
		suite.sqlmockX.ExpectPrepare(insert).WillReturnError(expected)

		err := suite.repo.Save(suite.tracer, domain)
		suite.Error(err)
		suite.Equal(expected, err)
	})

	suite.Run("Should return error - when exec", func() {
		expected := errors.New("error in exec")
		suite.sqlmockX.ExpectPrepare(insert).
			ExpectExec().
			WithArgs("1", "Toyota", "Corolla", 2020, "Blue", 20000.65).
			WillReturnError(expected)

		err := suite.repo.Save(suite.tracer, domain)
		suite.Error(err)
		suite.Equal(expected, err)
	})
}

func (suite *VehicleRepositoryTestSuite) TestUpdate() {
	const update = `UPDATE vehicles SET brand=\$1, model=\$2, year=\$3, color=\$4, price=\$5 WHERE id=\$6`

	domain := domains.Vehicle{
		ID:    "1",
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2020,
		Color: "Blue",
		Price: 20000.65,
	}
	suite.Run("Should return success", func() {
		suite.sqlmockX.ExpectPrepare(update).
			ExpectExec().
			WithArgs("Toyota", "Corolla", 2020, "Blue", 20000.65, "1").
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := suite.repo.Update(suite.tracer, domain)
		suite.NoError(err)
	})

	suite.Run("Should return error - when prepare", func() {
		expected := errors.New("error in prepare")
		suite.sqlmockX.ExpectPrepare(update).WillReturnError(expected)

		err := suite.repo.Update(suite.tracer, domain)
		suite.Error(err)
		suite.Equal(expected, err)
	})

	suite.Run("Should return error - when exec", func() {
		expected := errors.New("error in exec")
		suite.sqlmockX.ExpectPrepare(update).
			ExpectExec().
			WithArgs("Toyota", "Corolla", 2020, "Blue", 20000.65, "1").
			WillReturnError(expected)

		err := suite.repo.Update(suite.tracer, domain)
		suite.Error(err)
		suite.Equal(expected, err)
	})
}
