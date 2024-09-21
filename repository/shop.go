package repository

import (
	"context"

	"github.com/rnikrozoft/himymonsters-master-backend/model"
	"github.com/uptrace/bun"
)

type RepositoryIF interface {
	AddItems(m *model.StorageWrite) error
}

type newReposiroty struct {
	ctx context.Context
	db  *bun.DB
}

func NewRepository(ctx context.Context, db *bun.DB) RepositoryIF {
	return &newReposiroty{
		ctx: ctx,
		db:  db,
	}
}

func (r *newReposiroty) AddItems(m *model.StorageWrite) error {
	_, err := r.db.NewInsert().
		Model(m).
		On("CONFLICT (collection,key,user_id) DO UPDATE").
		Set("value = EXCLUDED.value").
		Exec(r.ctx)
	if err != nil {
		return err
	}
	return nil
}
