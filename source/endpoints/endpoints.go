package endpoints

import (
	"net/http"
	"strconv"

	"github.com/studigo/azarashi/logic"
	"github.com/studigo/azarashi/task"
)

// パスパラメータの位置.
const (
	TASK_ID_POS = 2
)

// タスクを新規作成.
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var req task.Request
	getBody(r, &req)

	inner, err := logic.Create(&req)
	if err != nil {
		E500(w, r)
	}

	response(w, inner.ToResponse(), http.StatusCreated)
}

// 子タスクを追加する.
func AddChild(w http.ResponseWriter, r *http.Request) {
	var req task.Request
	getBody(r, &req)

	id, err := strconv.ParseInt(getPathParam(r.URL.Path, TASK_ID_POS), 10, 64)
	if err != nil {
		E500(w, r)
		return
	}

	parent, err := logic.Fetch(id)
	if err != nil {
		E404(w, r)
		return
	}

	_, c, err := logic.AddChild(parent, &req)
	if err != nil {
		E500(w, r)
		return
	}

	response(w, c.ToResponse(), http.StatusCreated)
}

// タスクを取得する.
func GetTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(getPathParam(r.URL.Path, TASK_ID_POS), 10, 64)
	if err != nil {
		E500(w, r)
		return
	}

	inner, err := logic.Fetch(id)
	if err != nil {
		E404(w, r)
		return
	}

	response(w, inner.ToResponse(), http.StatusOK)
}
