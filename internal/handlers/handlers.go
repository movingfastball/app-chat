package handlers

import (
	"app-chat/domain"
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

// wsコネクションの基本設定
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var (
	// ペイロードチャネルを作成
	wsChan = make(chan domain.WsPayload)

	// コネクションマップを作成
	// keyはコネクション情報, valueにはユーザー名を入れる
	clients = make(map[domain.WebScoketConnection]string)
)

// WebSocketsのエンドポイント
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	// HTTPサーバーコネクションをWebSocketsプロトコルにアップグレード
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("OK Client Connecting")

	var response domain.WsJsonResponse
	response.Message = `<li>Connected to server</li>`

	// コネクション情報を格納
	conn := domain.WebScoketConnection{Conn: ws}
	// ブラウザが読み込まれた時に一度だけ呼び出されるのでユーザ名なし
	clients[conn] = ""

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

	go ListenForWs(&conn) // goroutineで呼び出し
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}
}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func ListenForWs(conn *domain.WebScoketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload domain.WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func broadcastToAll(response domain.WsJsonResponse) {
	// clientsには全ユーザーのコネクション情報が格納されている
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("websockets err")
			_ = client.Close()
			// clients Mapからclientの情報を消す
			delete(clients, client)
		}
	}
}

func ListenToWsChannel() {
	var response domain.WsJsonResponse

	for {
		// メッセージが入るまで、ここでブロック
		e := <-wsChan

		// アクションによって処理を分ける必要があるが便宜的に以下を宣言
		response.Action = "Got here"
		response.Message = fmt.Sprintf("Some message, and action was %s", e.Action)

		broadcastToAll(response)
	}
}
