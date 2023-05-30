package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	// var name string

	// fmt.Print("Enter your name: ")
	// var g Greeter = &CliGreeterImpl{}
	// fmt.Scanln(&name)
	// g.Greeting(context.Background(), name)

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "form.html")
	})

	var h Handler = Handler{
		greeter: &WebGreeterImpl{},
	}

	http.HandleFunc("/action", h.GreetingHandler)

	http.ListenAndServe(":8080", nil)
}

type Handler struct {
	greeter Greeter
}

type writerKeyType string

const writerKey writerKeyType = "writer"

func (h *Handler) GreetingHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	ctx := context.WithValue(r.Context(), writerKey, w)
	h.greeter.Greeting(ctx, name)
}

type WebGreeterImpl struct {
}

func (g *WebGreeterImpl) Greeting(ctx context.Context, name string) {
	w := ctx.Value(writerKey).(http.ResponseWriter)
	w.Write([]byte("Hello, " + name + "!"))
}

type CliGreeterImpl struct{}

func (g *CliGreeterImpl) Greeting(ctx context.Context, name string) {
	fmt.Println("Hello, " + name + "!")
}

type Greeter interface {
	Greeting(ctx context.Context, name string)
}

func Greeting(name string) string {
	return "Hello, " + name + "!"
}
