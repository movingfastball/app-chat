// connect.go
package domain

// WebSocketsからの返却用データの構造体
type WsJsonResponse struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}
