package repositories

import (
	"context"
	"github.com/luminosita/parser-bee/internal/infra/db/mongodb"
	"github.com/luminosita/parser-bee/internal/interfaces/respositories/documents"
)

type GetDocumentRepository struct {
	ctx context.Context
}

func NewGetDocumentRepository(ctx context.Context) *GetDocumentRepository {
	return &GetDocumentRepository{
		ctx: ctx,
	}
}

func (r *GetDocumentRepository) GetDocument(
	docData *documents.GetDocumentRepositorerRequest) (*documents.GetDocumentRepositorerResponse, error) {
	col := mongodb.GetDbCollection(r.ctx)
	_ = col.FindOne(r.ctx, docData)

	return nil, nil
}
