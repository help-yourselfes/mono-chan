package thread_repo

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	m "github.com/helpyourselfes/mono-chan/internal/app/thread/model"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
)

func (r *t) Update(ctx context.Context, thread *m.Thread) error {
	if thread == nil {
		return errors.New("thread is nil")
	}

	mediaJSON := "null"
	if len(thread.MediaLinks) > 0 {
		b, err := json.Marshal(thread.MediaLinks)
		if err != nil {
			return err
		}
		mediaJSON = string(b)
	}

	now := time.Now().UTC()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	query := `
		UPDATE threads
		SET caption = ?, text = ?, media_json = ?, reply_count = ?, password = ?, updated_at = ?
		WHERE key = ?
	`

	var passwordVal interface{}
	if thread.Password == "" {
		passwordVal = nil
	} else {
		passwordVal = thread.Password
	}

	res, err := tx.ExecContext(ctx, query,
		thread.Caption,
		thread.Text,
		mediaJSON,
		thread.ReplyCount,
		passwordVal,
		now,
		thread.Id,
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	if affected == 0 {
		_ = tx.Rollback()
		return customErrors.ErrNotFound
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	thread.UpdatedAt = &now
	return nil
}
