package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/studigo/azarashi/endpoints"
	"github.com/studigo/azarashi/http"
	"github.com/studigo/azarashi/logic"
	"github.com/studigo/marinesnow/v2"
)

func initialize() error {

	// 環境変数読み込み.
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println(fmt.Errorf("info: %s", err))
	}

	// ID採番機の設定.
	var n int64 = 1024
	val, err := rand.Int(rand.Reader, big.NewInt(n))
	if err != nil {
		return fmt.Errorf("initialize failed: %s", err)
	}
	marinesnow.SetWorkerID(val.Int64())
	marinesnow.SetTimestampOffset(time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC).UnixMilli())

	// logicの初期化.
	if err := logic.Initialize(); err != nil {
		return fmt.Errorf("initialize failed: %s", err)
	}

	return nil
}

func main() {

	// 初期化.
	if err := initialize(); err != nil {
		fmt.Println(err)
		return
	}

	http.Error404(endpoints.E404)
	http.Error405(endpoints.E405)
	http.HandleFunc(http.POST, "^/tasks$", endpoints.CreateTask)
	http.HandleFunc(http.POST, "^/task/[0-9]*/children$", endpoints.AddChild)
	http.HandleFunc(http.GET, "^/tasks/[0-9]*$", endpoints.GetTask)
	http.HandleFunc(http.DELETE, "^/tasks/[0-9]*$", endpoints.DeleteTask)
	http.HandleFunc(http.PUT, "^/tasks/[0-9]*/close$", endpoints.PutClose)
	http.HandleFunc(http.PUT, "^/tasks/[0-9]*/open$", endpoints.PutOpen)
	http.HandleFunc(http.PUT, "^/tasks/[0-9]*/title$", endpoints.PutTitle)
	http.HandleFunc(http.PUT, "^/tasks/[0-9]*/description$", endpoints.PutDescription)
	http.HandleFunc(http.PUT, "^/tasks/[0-9]*/parent$", endpoints.PutParent)

	// Listenポートを環境変数から設定(環境変数がなければ8080にする).
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// サーバーを開始.
	if err := http.ListenAndServe(port); err != nil {
		fmt.Println(fmt.Errorf("app start failed: %s", err))
		return
	}
}
