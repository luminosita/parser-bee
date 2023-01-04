package documents

import (
	"github.com/luminosita/honeycomb/pkg/interfaces"
	"github.com/luminosita/parser-bee/internal/domain/entities"
)

type GetAllDocumenterRequest = struct {
}
type GetAllDocumenterResponse = struct {
	Documents []*entities.Document
}

type GetAllDocumenter interface {
	interfaces.UseCaser[*GetAllDocumenterRequest, *GetAllDocumenterResponse]
}
