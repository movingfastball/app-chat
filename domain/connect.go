// connect.go
package domain

import "github.com/gorilla/websocket"

// WebSocketsからの返却用データの構造体
type WsJsonResponse struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

// WebSocketsコネクション情報を格納
type WebScoketConnection struct {
	*websocket.Conn
}

// WebSockets送信データを格納
type WsPayload struct {
	Action   string              `json:"action"`
	Message  string              `json:"message"`
	Username string              `json:"username"`
	Conn     WebScoketConnection `json:"-"`
}
