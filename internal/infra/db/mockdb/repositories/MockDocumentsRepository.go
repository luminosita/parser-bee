package repositories

import (
	"context"
	"github.com/luminosita/parser-bee/internal/domain/entities"
	"github.com/luminosita/parser-bee/internal/interfaces/respositories/documents"
)

type MockDocumentsRepository struct {
	ctx context.Context
}

func MakeMockDocumentsRepository(ctx context.Context) *MockDocumentsRepository {
	return &MockDocumentsRepository{
		ctx: ctx,
	}
}

func (r *MockDocumentsRepository) GetAllDocuments(
	*documents.GetAllDocumentsRepositorerRequest) (*documents.GetAllDocumentsRepositorerResponse, error) {

	return &documents.GetAllDocumentsRepositorerResponse{
		Documents: append(make([]*entities.Document, 0), &entities.Document{
			Name:       "MockName",
			DocumentId: "MockId",
		}),
	}, nil
}

func (r *MockDocumentsRepository) GetDocument(
	docData *documents.GetDocumentRepositorerRequest) (*documents.GetDocumentRepositorerResponse, error) {
	return &documents.GetDocumentRepositorerResponse{
		Document: &entities.Document{
			Name:       "MockName",
			DocumentId: docData.DocumentID,
		},
	}, nil
}

func (r *MockDocumentsRepository) CreateDocument(
	*documents.CreateDocumentRepositorerRequest) (*documents.CreateDocumentRepositorerResponse, error) {
	return &documents.CreateDocumentRepositorerResponse{
		DocumentId: "MockId",
	}, nil
}
