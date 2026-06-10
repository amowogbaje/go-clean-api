package repository

import (
	"database/sql"
	"go_clean_api/internal/domain"
)

type postgresMediaRepo struct {
	DB *sql.DB
}

func NewPostgresMediaRepository(db *sql.DB) domain.MediaRepository {
	return &postgresMediaRepo{DB: db}
}

func (r *postgresMediaRepo) Create(m *domain.Media) error {
	query := `INSERT INTO medias (url, type) VALUES ($1, $2) RETURNING id`
	err := r.DB.QueryRow(query, m.URL, m.Type).Scan(&m.ID)
	return err
}

func (r *postgresMediaRepo) FetchAll() ([]domain.Media, error) {
	query := `SELECT id, url, type FROM medias`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medias []domain.Media
	for rows.Next() {
		var m domain.Media
		if err := rows.Scan(&m.ID, &m.URL, &m.Type); err != nil {
			return nil, err
		}
		medias = append(medias, m)
	}
	return medias, nil
}