package domain

type UploadImageStore interface {
	UploadObject(image []byte) (string, error)
	GetObject(filename string) []byte
	Start()
}
