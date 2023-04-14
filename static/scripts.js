// scripts.js
 
let scoket = null;
 
// 最初のHTMLの読み込みと解析が完了したとき、スタイルシート、
// 画像、サブフレームの読み込みが完了するのを待たずに発生。
document.addEventListener("DOMContentLoaded", function(){
 
  // WebScoektオブジェクトの作成
  // https://developer.mozilla.org/ja/docs/Web/API/WebSockets_API/Writing_WebSocket_client_applications
  socket = new WebSocket("ws://127.0.0.1:8080/ws")
 
  //　接続を確立
  // https://developer.mozilla.org/ja/docs/Web/API/WebSocket/onopen
  socket.onopen = () => {
    console.log("Successfully connected")
  }

    // 接続がCLOSEDに変わった時に呼ばれる
  // https://developer.mozilla.org/ja/docs/Web/API/WebSocket/onclose
  socket.onclose = () => {
    console.log("connection closed")
  }
 
  // エラーが発生した時に呼び出される
  // https://developer.mozilla.org/ja/docs/Web/API/WebSocket/onerror
  socket.onerror = error => {
    console.log("there was an error")
  }
 
  // サーバーからメッセージが届いたときに呼び出される
  // https://developer.mozilla.org/ja/docs/Web/API/WebSocket/onmessage
  socket.onmessage = msg => {
    let j = JSON.parse(msg.data)
    console.log(j)
  }

  socket = new WebSocket("ws://127.0.0.1:8080/ws")
})