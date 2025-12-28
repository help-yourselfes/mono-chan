package thread_repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/helpyourselfes/mono-chan/internal/app/thread/model"
	"github.com/helpyourselfes/mono-chan/internal/pkg/customErrors"
)

func (r *t) GetByID(ctx context.Context, inputID int64) (*model.Thread, error) {
	query := `
		SELECT key, board_key, caption, text, media_json, reply_count, password, created_at, updated_at
		FROM threads
		WHERE key = ?
	`
	var (
		id        int64
		boardKey  string
		caption   string
		text      string
		mediaJSON string
		replyCnt  int64
		password  sql.NullString
		createdAt time.Time
		updatedAt sql.NullTime
	)

	row := r.db.QueryRowContext(ctx, query, inputID)
	if err := row.Scan(
		&id, &boardKey, &caption, &text, &mediaJSON, &replyCnt, &password, &createdAt, &updatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, customErrors.ErrNotFound
		}
		return nil, err
	}

	var mediaLinks = make([]string, 0)
	if err := json.Unmarshal([]byte(mediaJSON), &mediaLinks); err != nil {
		return nil, err
	}

	thread := model.Thread{
		Id:         id,
		BoardKey:   boardKey,
		Caption:    caption,
		Text:       text,
		MediaLinks: mediaLinks,
		ReplyCount: replyCnt,
		Password:   "",
		CreatedAt:  &createdAt,
		UpdatedAt:  nil,
	}

	if password.Valid {
		thread.Password = password.String
	}
	if updatedAt.Valid {
		t := updatedAt.Time
		thread.UpdatedAt = &t
	}

	return &thread, nil
}
