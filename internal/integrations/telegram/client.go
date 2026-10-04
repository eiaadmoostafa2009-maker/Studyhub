package telegram

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type Client interface {
    createMultipartBody(file io.Reader, filename string, caption string,) (io.Reader, string, error)
    SendVideo(ctx context.Context, file io.Reader, filename string, caption string) error
}
type client struct {
	Token  string
	ChatID string
}

func NewClient() Client {
	return &client{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		ChatID: os.Getenv("TELEGRAM_ADMIN_CHAT_ID"),
	}
}

func(c *client) createMultipartBody(file io.Reader, filename string, caption string,) (io.Reader, string, error) {

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)

	if err := writer.WriteField("chat_id", c.ChatID); err != nil {
		return nil, "", err
	}

	if caption != "" {
		if err := writer.WriteField("caption", caption); err != nil {
			return nil, "", err
		}
	}

	part, err := writer.CreateFormFile(
		"video",
		filename,
	)
	if err != nil {
		return nil, "", err
	}

	if _, err := io.Copy(part, file); err != nil {
		return nil, "", err
	}

	if err := writer.Close(); err != nil {
		return nil, "", err
	}

	return &body, writer.FormDataContentType(), nil
}

func (c *client) SendVideo(ctx context.Context, file io.Reader, filename string, caption string) error {

	apiURL := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendVideo",
		c.Token,
	)

	body, contentType, err := c.createMultipartBody(
		file,
		filename,
		caption,
	)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		apiURL,
		body,
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)

		return fmt.Errorf(
			"telegram returned %s: %s",
			resp.Status,
			string(data),
		)
	}

	return nil
}