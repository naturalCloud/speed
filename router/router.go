package router

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"speed/app/http/controllers"
	"speed/app/lib/log"
	"speed/app/middleware"
	"time"
)

var wsport string = "8086"

var OnlineMap =  make(map[string]*websocket.Conn)

func Router(c *gin.Engine) {

	c.Use(middleware.Trace)
	c.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello word!")

	})

	{
		c.GET("/index",controllers.HelloC.Index)
	}
	{
		c.GET("/first", func(context *gin.Context) {
			context.HTML(http.StatusOK,"Firstroom.html",gin.H{"wsport":wsport})

		})

		c.GET("/diff", func(context *gin.Context) {
			context.HTML(http.StatusOK,"Differentroom.html",gin.H{"wsport":wsport})

		})

		c.GET("/some", func(context *gin.Context) {
			context.HTML(http.StatusOK,"Sameroom.html",gin.H{"wsport":wsport})

		})
	}
	//webscoket 服务
	{
		c.GET("/ws", func(context *gin.Context) {
			//up
			var upgrader = websocket.Upgrader{
				CheckOrigin: func(r *http.Request) bool {
					return  true

				},
			}

			conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
			if err != nil {
				log.WithCtx(context).Error("握手错误!!!")
			}
			//read

			const  (
				// Time allowed to write a message to the peer.
				writeWait = 10 * time.Second

				// Time allowed to read the next pong message from the peer.
				pongWait = 60 * time.Second

				// Send pings to peer with this period. Must be less than pongWait.
				pingPeriod = (pongWait * 9) / 10

				// Maximum message size allowed from peer.
				maxMessageSize = 512
			)



			var (
				newline = []byte{'\n'}
				space   = []byte{' '}
			)

			OnlineMap[uuid.New().String()] = conn
			go func() {
				conn.SetReadLimit(maxMessageSize)
				conn.SetReadDeadline(time.Now().Add(pongWait))
				//conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
				for {
					_, message, err := conn.ReadMessage()

					if err != nil {
						if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
							fmt.Printf("error: %v", err)
						}
						break
					}
					message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
					fmt.Printf("收到消息--%s--,%s\n",string(message),conn.RemoteAddr())
					 conn.WriteMessage(websocket.TextMessage, []byte("我是websocket"))
					fmt.Println("返回消息---","我是websocket")

				}
			}()


		})
	}

}
