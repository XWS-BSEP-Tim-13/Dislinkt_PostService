package domain

import "context"

type UploadImageStore interface {
	UploadObject(ctx context.Context, image []byte) (string, error)
	GetObject(ctx context.Context, filename string) []byte
	Start(ctx context.Context)
}
