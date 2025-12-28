package thread_repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	m "github.com/helpyourselfes/mono-chan/internal/app/thread/model"
)

func (r *t) List(ctx context.Context, boardKey string) ([]*m.Thread, error) {
	query := `
		SELECT key, board_key, caption, text, media_json, reply_count, password, created_at, updated_at
		FROM threads
		WHERE board_key = ?
		ORDER BY reply_count
	`

	rows, err := r.db.QueryContext(ctx, query, boardKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads []*m.Thread
	for rows.Next() {
		var (
			id        int64
			bk        string
			caption   string
			text      string
			mediaJSON string
			replyCnt  int64
			password  sql.NullString
			createdAt time.Time
			updatedAt sql.NullTime
		)

		if err := rows.Scan(
			&id, &bk, &caption, &text, &mediaJSON, &replyCnt, &password, &createdAt, &updatedAt,
		); err != nil {
			return nil, err
		}

		var mediaLinks []string
		if mediaJSON != "" {
			if err := json.Unmarshal([]byte(mediaJSON), &mediaLinks); err != nil {
				return nil, err
			}
		}

		th := &m.Thread{
			BoardKey:   bk,
			Id:         id,
			Caption:    caption,
			Text:       text,
			MediaLinks: mediaLinks,
			ReplyCount: replyCnt,
			Password:   "",
			CreatedAt:  &createdAt,
			UpdatedAt:  nil,
		}
		if password.Valid {
			th.Password = password.String
		}
		if updatedAt.Valid {
			t := updatedAt.Time
			th.UpdatedAt = &t
		}

		threads = append(threads, th)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return threads, nil
}
