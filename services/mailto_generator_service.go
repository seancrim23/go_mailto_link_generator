package services

import "context"

type MailtoGeneratorService interface {
	CreateMailtoLink(context.Context, string, string, string) (string, error)
	GetMailtoLink(context.Context, string) (string, error)
}
