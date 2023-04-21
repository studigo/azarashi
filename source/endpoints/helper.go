package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Bodyを構造体に変換する.
func getBody(request *http.Request, v any) error {
	body := request.Body
	defer body.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, body)

	return json.Unmarshal(buf.Bytes(), v)
}

// レスポンスに結果を書き込む.
func response(w http.ResponseWriter, value any, statusCode int) {
	res, err := json.MarshalIndent(value, "", "	")
	if err != nil {
		E500(w, &http.Request{})
		return
	}
	w.WriteHeader(statusCode)
	fmt.Fprint(w, string(res))
}

// URLのn番目のPATHパラメータを取得する.
func getPathParam(path string, n int) string {
	return strings.Split(path, "/")[n]
}
