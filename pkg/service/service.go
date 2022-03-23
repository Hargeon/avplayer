package service

import "context"

type CloudStorage interface {
	Download(ctx context.Context, id, path string) (string, error)
}
