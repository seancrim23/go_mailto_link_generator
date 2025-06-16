package services

import (
	"context"
	"errors"
	"fmt"
	"mailto_link_generator/models"
	"math/rand"
	"time"

	"cloud.google.com/go/firestore"
)

type FirestoreMailtoGeneratorService struct {
	Database *firestore.Client
}

func (f *FirestoreMailtoGeneratorService) CreateMailtoLink(ctx context.Context, to string, subject string, body string) (string, error) {
	shortUrl := generateShortUrl(8)

	builtUrl := &models.Mail{
		Id:      shortUrl,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	_, err := f.Database.Collection("URL").Doc(shortUrl).Set(ctx, builtUrl)
	if err != nil {
		fmt.Println("error generating short url")
		fmt.Println(err)
		return "", errors.New("error generating short url")
	}

	return shortUrl, nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// basically from the internet
// TODO update if theres any custom url or more fancy url generation i find
func generateShortUrl(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte

	for i := 0; i < length; i++ {
		index := seededRand.Intn(len(charset))
		result = append(result, charset[index])
	}

	return string(result)
}

func (f *FirestoreMailtoGeneratorService) GetMailtoLink(ctx context.Context, id string) (string, error) {
	dsnap, err := f.Database.Collection("URL").Doc(id).Get(ctx)
	if err != nil {
		fmt.Println("error getting long url")
		fmt.Println(err)
		return "", errors.New("error getting long url")
	}

	var m models.Mail
	err = dsnap.DataTo(&m)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("error getting long url")
	}

	mailToLink := generateMailToLink(m.To, m.Subject, m.Body)
	return mailToLink, nil
}

func generateMailToLink(to string, subject string, body string) string {
	return fmt.Sprintf("mailto:%s?subject=%s&body=%s", to, subject, body)
}
