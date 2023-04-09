package endpoints

import (
	"fmt"
	"net/http"
)

// エラー情報を表す構造体.
type errorResponse struct {
	Status  int    `json:"error"`
	Message string `json:"message"`
}

// エラーを文字列化する.
func (p *errorResponse) ToString() string {
	return fmt.Sprintf("{\"error\":\"%d\",\"message\":\"%s\"}", p.Status, p.Message)
}

// エラーコードとエラーメッセージを設定する.
func (p *errorResponse) Set(errorCode int, message string) *errorResponse {
	p.Status = errorCode
	p.Message = message
	return p
}

// リクエストが間違っている場合.
func E400(w http.ResponseWriter, r *http.Request) {
	e := new(errorResponse).Set(400, "bad request")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, e.ToString())
}

// 存在しないpathにアクセスされた場合.
func E404(w http.ResponseWriter, r *http.Request) {
	e := new(errorResponse).Set(404, "resource not found")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, e.ToString())
}

// 未定義のHTTPメソッドでアクセスされた場合.
func E405(w http.ResponseWriter, r *http.Request) {
	e := new(errorResponse).Set(405, "method not allowed")
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprint(w, e.ToString())
}

// DB接続エラーなど処理中にエラーが発生した場合.
func E500(w http.ResponseWriter, r *http.Request) {
	e := new(errorResponse).Set(500, "internal server error")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, e.ToString())
}
