package ext

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
)

func PerformHttpRequest(r http.Handler, method, path string, header http.Header, reqBody io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, reqBody)
	if header != nil {
		req.Header = header
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func DecodeJSON(r io.Reader) (map[string]interface{}, error) {
	var m map[string]interface{}
	err := DecodeJsonInto(r, &m)
	if err != nil {
		return nil, err
	}

	if len(m) <= 0 {
		return nil, errors.New("ext.DecodeJSON: no content")
	}

	return m, nil
}

func DecodeJsonInto(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(obj)
	if err != nil {
		return err
	}

	return nil
}

func DecodeJSONSlice(r io.Reader) ([]map[string]interface{}, error) {
	var m []map[string]interface{}
	err := DecodeJsonInto(r, &m)
	if err != nil {
		return nil, err
	}

	if len(m) <= 0 {
		return nil, errors.New("ext.DecodeJSONSlice: no content")
	}

	return m, nil
}

func ReadMockFile(filepath string) string {
	file, err := os.Open(filepath) // For read access.
	if err != nil {
		return ""
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	sdata := string(data)
	return sdata
}

func Pre(handlers map[string]http.HandlerFunc) {

	return
}
