package services

import (
	"context"

	"cloud.google.com/go/firestore"
)

type FirestoreMailtoGeneratorService struct {
	Database *firestore.Client
}

func (f *FirestoreMailtoGeneratorService) CreateMailtoLink(ctx context.Context, to string, subject string, body string) (string, error) {
	return "finish this function", nil
}

func (f *FirestoreMailtoGeneratorService) GetMailtoLink(ctx context.Context, id string) (string, error) {
	return "finish this function", nil
}
