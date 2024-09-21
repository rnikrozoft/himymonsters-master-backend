package model

import "github.com/uptrace/bun"

type StorageWrite struct {
	bun.BaseModel `bun:"table:storage"`

	Collection      string `bun:"collection"`
	Key             string `bun:"key"`
	UserID          string `bun:"user_id"`
	Value           string `bun:"value"`
	Version         string `bun:"version"`
	PermissionRead  int    `bun:"read"`
	PermissionWrite int    `bun:"write"`
}
