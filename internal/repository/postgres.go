package repository

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type PostgresRepo struct {
	db *pgx.Conn
}

func NewPostgresRepo(db *pgx.Conn) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (repo *PostgresRepo) SaveNumber(ctx context.Context, number int) error {
	_, err := repo.db.Exec(ctx, "INSERT INTO nums (num) VALUES ($1)", number)
	if err != nil {
		return err
	}
	return nil
}

func (repo *PostgresRepo) GetSortedNums(ctx context.Context) ([]int, error) {
	var nums []int
	rows, err := repo.db.Query(ctx, "SELECT num FROM nums ORDER BY num ASC")
	if err != nil {
		return nil, err
	}

	err = pgxscan.ScanAll(&nums, rows)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return nums, nil
}
