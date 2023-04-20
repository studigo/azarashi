/*
PATHパラメータありのURLから関数呼び出しするハンドラなど.
(net/httpのラッパー)
*/
package http

import (
	"net/http"
	"regexp"
)

type httpMethod string

const (
	POST   httpMethod = "POST"
	GET    httpMethod = "GET"
	PUT    httpMethod = "PUT"
	PATCH  httpMethod = "PATCH"
	DELETE httpMethod = "DELETE"
)

// httpメソッドとコールバック関数を纏めるための構造体.
type handle struct {
	method      string
	regexpCache *regexp.Regexp
	function    func(w http.ResponseWriter, r *http.Request)
}

// PATHのルーティングを行うための辞書.
var (
	routingMap []*handle                                    = []*handle{}
	e404f      func(w http.ResponseWriter, r *http.Request) = nil
	e405f      func(w http.ResponseWriter, r *http.Request) = nil
)

/*
PATHに対する処理を設定する(?でパスパラメータを指定可能).
ex) /tasks/?/children ← ? の部分は任意の文字列
*/
func HandleFunc(method httpMethod, path string, f func(w http.ResponseWriter, r *http.Request)) {
	routingMap = append(routingMap,
		&handle{
			method:      string(method),
			regexpCache: regexp.MustCompile(path),
			function:    f,
		},
	)
}

// 404の時に呼び出す処理を登録する.
func Error404(f func(w http.ResponseWriter, r *http.Request)) {
	e404f = f
}

// 405の時に呼び出す処理を登録する.
func Error405(f func(w http.ResponseWriter, r *http.Request)) {
	e405f = f
}

// 指定ポートを待ち受ける.
func ListenAndServe(port string) error {
	http.HandleFunc("/", routing)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		return err
	}

	return nil
}

// HTTPリクエストのルーティングを行う.
func routing(w http.ResponseWriter, r *http.Request) {

	var is405 bool

	for _, v := range routingMap {

		if !v.regexpCache.MatchString(r.URL.Path) {
			continue
		}

		if v.method != r.Method {
			is405 = true
			continue
		}

		v.function(w, r)
		return
	}

	// 405 Method Not Allowed.
	if is405 && e405f != nil {
		e405f(w, r)
		return
	}

	// 404 Not Found.
	if e404f != nil {
		e404f(w, r)
		return
	}
}
