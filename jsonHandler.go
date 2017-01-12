package main

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, r *http.Request, data []byte, v interface{}) error {

	if err := json.Unmarshal(data, &v); err != nil {

		w.Header().Set("Content-Type:", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			return err
		}
		return err
	}
	return nil
}

func HttpResponse(w http.ResponseWriter, r *http.Request, httpStatus int) {
	w.Header().Set("Content-Type:", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)
}

func WrapJSON(v interface{}, status bool) ([]byte, error) {
	wrapped := map[string]interface{}{
		"success": status,
		"data":    v,
	}
	converted, err := json.Marshal(wrapped)
	if err != nil {
		panic(err)
	}
	return converted, nil
}
