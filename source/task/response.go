package task

// レスポンスを表す構造体.
type Response struct {
	ID          int64       `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	IsClosed    bool        `json:"is_closed"`
	Children    []*Response `json:"children"`
}
