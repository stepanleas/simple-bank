package db

import (
	"context"
	"fmt"
)

func (s *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %s, rb err: %s", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
