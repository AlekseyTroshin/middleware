package render

import (
	"bytes"
	"encoding/json"
	resp "hexlet-auth/internal/lib/api/response"
	"io"
	"net/http"
)

func DecoderJSON(body io.ReadCloser, r interface{}) error {
	return json.NewDecoder(body).Decode(r)
}

func JSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	err := json.NewEncoder(w).Encode(resp.Error("failed to decode request"))
	if err != nil {
		return
	}
}

func RequestBody(r *http.Request) ([]byte, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}
