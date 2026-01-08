package app

import (
	boards "github.com/helpyourselfes/mono-chan/internal/app/board/repo"
	posts "github.com/helpyourselfes/mono-chan/internal/app/post/repo"
	threads "github.com/helpyourselfes/mono-chan/internal/app/thread/repo"
)

type Repos struct {
	Boards  boards.BoardRepo
	Threads threads.ThreadRepo
	Posts   posts.PostRepo
}
