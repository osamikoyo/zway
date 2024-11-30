package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"zway/internal/data"
	"zway/internal/data/models"
)

func pingHandler(w http.ResponseWriter, r *http.Request) error {
	if _, err := w.Write([]byte("PONG")); err != nil {
		return err
	}

	return nil
}

func getTextHandler(w http.ResponseWriter, r *http.Request) error {
	uniq := r.URL.Query().Get("hash")

	text, err := data.GetText(uniq)
	if err != nil {
		return err
	}

	body, err := json.Marshal(&text)
	if err != nil {
		return err
	}

	if _, err := w.Write(body); err != nil {
		return err
	}

	return nil
}

func AddText(w http.ResponseWriter, r *http.Request) error {
	var text models.Text
	text.Text = r.FormValue("content")
	text.Title = r.FormValue("title")
	text.Date = time.Now().GoString()

	return data.AddText(text)
}
