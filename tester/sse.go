package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type Content struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Event struct {
	Contents []Content `json:"contents"`
	Time     string    `json:"time"`
}

func generateContents() []Content {
	var contents []Content
	for i := 1; i <= 10; i++ {
		contents = append(contents, Content{
			ID:    i,
			Title: fmt.Sprintf("Content Title %d", rand.Intn(100)),
		})
	}
	return contents
}

func main() {
	app := fiber.New()

	app.Get("/events", func(c *fiber.Ctx) error {
		ctx := c.Context()

		ctx.SetContentType("text/event-stream")
		ctx.Response.Header.Set("Cache-Control", "no-cache")
		ctx.Response.Header.Set("Connection", "keep-alive")
		ctx.Response.Header.Set("Transfer-Encoding", "chunked")
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Cache-Control")
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			fmt.Println("WRITER")
			for {
				event := Event{
					Contents: generateContents(),
					Time:     time.Now().Format(time.RFC3339),
				}
				data, _ := json.Marshal(event)
				msg := fmt.Sprintf("data: %s\n\n", data)
				fmt.Fprintf(w, "%s", msg)
				w.Flush()
				time.Sleep(1 * time.Second)
			}
		}))
		return nil
	})

	fmt.Println("Server started at :8080")
	app.Listen(":8080")
}
