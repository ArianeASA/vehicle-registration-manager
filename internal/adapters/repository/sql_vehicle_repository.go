package repository

import (
	"database/sql"
	"vehicle-registration-manager/internal/adapters/repository/entities"
	"vehicle-registration-manager/internal/adapters/repository/mappers"
	"vehicle-registration-manager/internal/core/domain"
)

type SQLVehicleRepository struct {
	db *sql.DB
}

func NewSQLVehicleRepository(dataSourceName string) (*SQLVehicleRepository, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &SQLVehicleRepository{db: db}, nil
}

func (r *SQLVehicleRepository) Save(vehicle domain.Vehicle) error {
	_, err := r.db.Exec("INSERT INTO vehicles (id, brand, model, year, color, price) VALUES (?, ?, ?, ?, ?, ?)",
		vehicle.ID, vehicle.Brand, vehicle.Model, vehicle.Year, vehicle.Color, vehicle.Price)
	return err
}

func (r *SQLVehicleRepository) Update(vehicle domain.Vehicle) error {
	_, err := r.db.Exec("UPDATE vehicles SET brand=?, model=?, year=?, color=?, price=? WHERE id=?",
		vehicle.Brand, vehicle.Model, vehicle.Year, vehicle.Color, vehicle.Price, vehicle.ID)
	return err
}

func (r *SQLVehicleRepository) FindAll() ([]domain.Vehicle, error) {
	rows, err := r.db.Query("SELECT id, brand, model, year, color, price FROM vehicles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []domain.Vehicle
	for rows.Next() {
		var vehicle entities.Vehicle
		if err := rows.Scan(&vehicle.ID, &vehicle.Brand, &vehicle.Model, &vehicle.Year, &vehicle.Color, &vehicle.Price); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, mappers.EntityToDomain(vehicle))
	}
	return vehicles, nil
}

func (r *SQLVehicleRepository) FindByID(id string) (domain.Vehicle, error) {
	row := r.db.QueryRow("SELECT id, brand, model, year, color, price FROM vehicles WHERE id=?", id)
	var vehicle entities.Vehicle
	if err := row.Scan(&vehicle.ID, &vehicle.Brand, &vehicle.Model, &vehicle.Year, &vehicle.Color, &vehicle.Price); err != nil {
		return domain.Vehicle{}, err
	}

	return mappers.EntityToDomain(vehicle), nil
}
