package repository

import (
	"database/sql"
	"errors"
	"vehicle-registration-manager/internal/adapters/repository/configs"
	"vehicle-registration-manager/internal/adapters/repository/entities"
	"vehicle-registration-manager/internal/adapters/repository/mappers"
	"vehicle-registration-manager/internal/core/domains"
	"vehicle-registration-manager/pkg/tracer"
)

type VehicleRepository struct {
	db *sql.DB
}

func NewVehicleRepository(config configs.DatabaseConfigs) *VehicleRepository {
	return &VehicleRepository{db: config.GetDB()}
}

const (
	insert   = "INSERT INTO vehicles (id, brand, model, year, color, price) VALUES ($1, $2, $3, $4, $5, $6)"
	update   = "UPDATE vehicles SET brand=$1, model=$2, year=$3, color=$4, price=$5 WHERE id=$6"
	findAll  = "SELECT id, brand, model, year, color, price FROM vehicles"
	findByID = "SELECT id, brand, model, year, color, price FROM vehicles WHERE id=$1"
)

const (
	msgFailedPrepare = "failed to prepare statement"
	msgFailedClose   = "failed to close statement"
)

func (r *VehicleRepository) Save(tcr *tracer.Tracer, vehicle domains.Vehicle) error {

	stmt, err := r.db.Prepare(insert)
	if err != nil {
		tcr.Logger.Error(msgFailedPrepare, err)
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			tcr.Logger.Error(msgFailedClose, err)
		}
	}(stmt)

	_, err = stmt.Exec(vehicle.ID, vehicle.Brand, vehicle.Model, vehicle.Year, vehicle.Color, vehicle.Price)
	if err != nil {
		tcr.Logger.Error("failed to execute statement", err)
		return err
	}
	return nil
}

func (r *VehicleRepository) Update(tcr *tracer.Tracer, vehicle domains.Vehicle) error {
	stmt, err := r.db.Prepare(update)
	if err != nil {
		tcr.Logger.Error(msgFailedPrepare, err)
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			tcr.Logger.Error(msgFailedClose, err)
		}
	}(stmt)

	_, err = stmt.Exec(vehicle.Brand, vehicle.Model, vehicle.Year, vehicle.Color, vehicle.Price, vehicle.ID)
	if err != nil {
		tcr.Logger.Error("failed to execute statement", err)
		return err
	}
	return nil
}

func (r *VehicleRepository) FindAll(tcr *tracer.Tracer) ([]domains.Vehicle, error) {
	rows, err := r.db.Query(findAll)
	if err != nil {
		tcr.Logger.Error("failed to query", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			tcr.Logger.Error("failed to close rows", err)
		}
	}(rows)

	var vehicles []domains.Vehicle
	for rows.Next() {
		var vehicle entities.Vehicle
		if err := rows.Scan(&vehicle.ID, &vehicle.Brand, &vehicle.Model, &vehicle.Year, &vehicle.Color, &vehicle.Price); err != nil {
			tcr.Logger.Error("failed to scan", err)
			return nil, errors.New("failed to scan")
		}
		vehicles = append(vehicles, mappers.EntityToDomain(vehicle))
	}
	return vehicles, nil
}

func (r *VehicleRepository) FindByID(tcr *tracer.Tracer, id string) (domains.Vehicle, error) {
	stmt, err := r.db.Prepare(findByID)
	if err != nil {
		tcr.Logger.Error(msgFailedPrepare, err)
		return domains.Vehicle{}, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			tcr.Logger.Error(msgFailedClose, err)
		}
	}(stmt)

	var vehicle entities.Vehicle
	err = stmt.QueryRow(id).Scan(&vehicle.ID, &vehicle.Brand, &vehicle.Model, &vehicle.Year, &vehicle.Color, &vehicle.Price)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		tcr.Logger.Error("failed to query row", err)
		return domains.Vehicle{}, err
	}

	return mappers.EntityToDomain(vehicle), nil
}
