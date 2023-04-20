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
	if getBody(r, &req) != nil {
		E400(w, r)
		return
	}

	inner, err := logic.Create(&req)
	if err != nil {
		E500(w, r)
		return
	}

	response(w, inner.ToResponse(), http.StatusCreated)
}

// 子タスクを追加する.
func AddChild(w http.ResponseWriter, r *http.Request) {
	var req task.Request
	if getBody(r, &req) != nil {
		E400(w, r)
		return
	}

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

// タスクを削除する.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
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

	if logic.Delete(inner) != nil {
		E500(w, r)
		return
	}

	response(w, "", http.StatusNoContent)
}

// タスクを完了状態にする.
func PutClose(w http.ResponseWriter, r *http.Request) {
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

	if logic.Close(inner) != nil {
		E500(w, r)
		return
	}

	response(w, inner.ToResponse(), http.StatusOK)
}

// タスクを未完了状態にする.
func PutOpen(w http.ResponseWriter, r *http.Request) {
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

	if logic.Open(inner) != nil {
		E500(w, r)
		return
	}

	response(w, inner.ToResponse(), http.StatusOK)
}

// タスクのタイトルを変更する.
func PutTitle(w http.ResponseWriter, r *http.Request) {
	var newTitle struct {
		Title string `json:"title"`
	}
	if getBody(r, &newTitle) != nil {
		E400(w, r)
		return
	}

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

	if logic.ChangeTitle(inner, newTitle.Title) != nil {
		E500(w, r)
		return
	}

	response(w, inner.ToResponse(), http.StatusOK)
}

// タスクの説明文を変更する.
func PutDescription(w http.ResponseWriter, r *http.Request) {
	var newDescription struct {
		Description string `json:"description"`
	}
	if getBody(r, &newDescription) != nil {
		E400(w, r)
		return
	}

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

	if logic.ChangeDescription(inner, newDescription.Description) != nil {
		E500(w, r)
		return
	}

	response(w, inner.ToResponse(), http.StatusOK)
}

// タスクの親を変更する.
func PutParent(w http.ResponseWriter, r *http.Request) {
	var newParent struct {
		ID int64 `json:"parent_id"`
	}
	if getBody(r, &newParent) != nil {
		E400(w, r)
		return
	}

	id, err := strconv.ParseInt(getPathParam(r.URL.Path, TASK_ID_POS), 10, 64)
	if err != nil {
		E500(w, r)
		return
	}

	p, err := logic.Fetch(newParent.ID)
	if err != nil {
		E404(w, r)
		return
	}

	inner, err := logic.Fetch(id)
	if err != nil {
		E404(w, r)
		return
	}

	if logic.ChangeParent(inner, p) != nil {
		E500(w, r)
		return
	}

	response(w, inner.ToResponse(), http.StatusOK)
}
