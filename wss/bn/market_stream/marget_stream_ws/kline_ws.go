package marketstreamws

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type KlineLineResponse struct {
	Ee string `json:"e"`
	E  int64  `json:"E"`
	S  string `json:"s"`
	K  struct {
		Tt int64  `json:"t"`
		T  int64  `json:"T"`
		S  string `json:"s"`
		I  string `json:"i"`
		F  int    `json:"f"`
		Ll int    `json:"L"`
		O  string `json:"o"`
		C  string `json:"c"`
		H  string `json:"h"`
		L  string `json:"l"`
		Vv string `json:"v"`
		N  int    `json:"n"`
		X  bool   `json:"x"`
		Qq string `json:"q"`
		V  string `json:"V"`
		Q  string `json:"Q"`
		B  string `json:"B"`
	} `json:"k"`
}

var WaitClientPongInterval = 10 * time.Second
var PingToClientInterval = 15 * time.Second
var (
	upgrader = websocket.Upgrader{}
)

type KlineServerWs struct { // method write, read, ping handler
	Ws                 *websocket.Conn
	WriteToTradeClient chan []byte
	ServerDone         chan bool
}

func (k *KlineServerWs) WriteTo(bn_client_done chan<- bool, bn_client_read <-chan []byte) {
	for {
		select {
		case msg, ok := <-bn_client_read:
			if !ok {
				k.Ws.WriteMessage(websocket.CloseMessage, []byte{})
				k.ServerDone <- true
				bn_client_done <- true
				return
			}
			err := k.Ws.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
					k.ServerDone <- true
					bn_client_done <- true
					close(k.WriteToTradeClient)
					return
				}
			}
		default:
		}
	}

}

type klineClientWs struct { // method write, read, pong handler
	Dial             *websocket.Conn
	ReadFromBnServer chan []byte
	ClientDone       chan bool
}

func (k *klineClientWs) ReadTo(trade_server_done chan<- bool) {
	for {
		_, msg, err := k.Dial.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				k.ClientDone <- true
				trade_server_done <- true
				close(k.ReadFromBnServer)
				return
			}
			k.ClientDone <- true
			trade_server_done <- true
			close(k.ReadFromBnServer)
			return
		}
		k.ReadFromBnServer <- msg
	}
}

func (m *marketStreamWs) KlineWs(
	c context.Context,
	w http.ResponseWriter,
	r *http.Request,
	h http.Header,
	symbol string,
	interval string,
) error {
	con, err := upgrader.Upgrade(w, r, h)
	if err != nil {
		return err
	}
	defer con.Close()

	trade_server := KlineServerWs{
		Ws:                 con,
		WriteToTradeClient: make(chan []byte),
		ServerDone:         make(chan bool, 1),
	}

	dial, _, err := websocket.DefaultDialer.Dial(
		fmt.Sprintf("wss://fstream.binance.com/ws/%v@kline_%v", symbol, interval),
		nil,
	)
	if err != nil {
		return err
	}
	defer dial.Close()
	dial.SetPingHandler(nil)

	bn_client := klineClientWs{
		Dial:             dial,
		ReadFromBnServer: make(chan []byte),
		ClientDone:       make(chan bool),
	}

	go trade_server.WriteTo(bn_client.ClientDone, bn_client.ReadFromBnServer)
	go bn_client.ReadTo(trade_server.ServerDone)

	<-trade_server.ServerDone
	<-bn_client.ClientDone

	return nil
}
