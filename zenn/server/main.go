package server

import (
	"context"
	"fmt"
)

type MyServer struct {
	router map[string]handlers.MyHandleFunc
}

func (srv *MyServer) Request(ctx context.Context, path string, token string) {
	// リクエストオブジェクト作成
	var req handlers.MyRequest
	req.SetPath(path)

	// (key:authToken <=> value:token)をcontextに入れる
	ctx = auth.SetAuthToken(ctx, token)

	// ルーティング操作
	if handler, ok := srv.router[req.GetPath()]; ok {
		handler(ctx, req)
	} else {
		handlers.NotFoundHandler(ctx, req)
	}
}

func (srv *MyServer) ListenAndServe() {
	for {
		var path, token string
		fmt.Scan(&path)
		fmt.Scan(&token)

		ctx := session.SetSessionID(context.Background())
		go srv.Request(ctx, path, token)
	}
}
