package util

import (
	"bytes"
	"github.com/JPratama7/util/sync"
	"github.com/goccy/go-json"
	"net/http"
)

var internalPool = sync.NewPool(func() *bytes.Buffer {
	return bytes.NewBuffer(make([]byte, 0, 1024))
})

type Response[T any] struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Status  string `json:"status"`
	Data    T      `json:"data"`
}

type ErrorResponse = Response[*int]

func (r Response[T]) SetCode(code int) Response[T] {
	r.Code = code
	return r
}

func (r Response[T]) SetSuccess(success bool) Response[T] {
	r.Success = success
	return r
}

func (r Response[T]) SetStatus(status string) Response[T] {
	r.Status = status
	return r
}

func (r Response[T]) SetData(data T) Response[T] {
	r.Data = data
	return r
}

func (r Response[T]) Write(w http.ResponseWriter) (err error) {
	buf := internalPool.Get()
	defer internalPool.Put(buf)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)

	enc := json.NewEncoder(buf)

	enc.SetEscapeHTML(true)
	if err = enc.Encode(r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(buf.Bytes())
	return
}
