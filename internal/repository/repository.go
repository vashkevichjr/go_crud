package repository

import "context"

type Repository interface {
	SaveNumber(ctx context.Context, number int) error
	GetSortedNums(ctx context.Context) ([]int, error)
}
