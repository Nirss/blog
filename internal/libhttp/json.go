package libhttp

import (
	"encoding/json"
	"log"
	"net/http"
)

const ContentTypeJson = "application/json"

type HTTPError struct {
	Error string
}

//JsonDecode Преобразовывает данные из json в пришедший формат
func JsonDecode(r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Printf("decode request data error: %v", err)
		return err
	}

	return nil
}

//JsonEncode Возвращает преобразованный в json ответ
func JsonEncode(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Add("Content-Type", ContentTypeJson)
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("encode data error: %v", err)
	}
}

func SendResponse(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func SendResponseWithValue(w http.ResponseWriter, status int, value interface{}) {
	w.WriteHeader(status)
}

func SendError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	JsonEncode(w, status, HTTPError{
		Error: err.Error(),
	})
}
