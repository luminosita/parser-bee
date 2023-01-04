package documents

import "github.com/luminosita/parser-bee/internal/domain/entities"

type GetAllDocumentsRepositorerRequest = struct {
}
type GetAllDocumentsRepositorerResponse = struct {
	Documents []*entities.Document
}

type GetAllDocumentsRepositorer interface {
	GetAllDocuments(req *GetAllDocumentsRepositorerRequest) (*GetAllDocumentsRepositorerResponse, error)
}
