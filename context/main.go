package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type stringKey string

const (
	userKey stringKey = "ID"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	ctx = context.WithValue(ctx, userKey, 1)
	result, err := reqHTTPS(ctx, "https://httpbin.org/delay/3")
	if err != nil {
		fmt.Println("Request error: ", err)
		return
	}

	fmt.Println("Result:", result)
}

func reqHTTPS(ctx context.Context, url string) (string, error) {
	err := userId(ctx)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Проверка входящих данных
func userId(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		if _, ok := ctx.Value(userKey).(int); !ok {
			return fmt.Errorf("wrong user-id format")
		}
	}
	return nil
}
