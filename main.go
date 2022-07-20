package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type hello struct {
	app.Compo
	name string
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("输入名字"),
		app.Input().ID("name").Value(h.name).OnChange(h.ValueTo(&h.name)),
		app.Button().Text("确定").OnClick(func(ctx app.Context, e app.Event) {
			app.Log("click")
			app.Window().Call("alert", h.name)
		}),
	)
}

func main() {
	app.Route("/", &hello{})
	app.RunWhenOnBrowser()

	var t = flag.String("type", "online", "版本 online/offline")
	flag.Parse()
	if *t == "online" {
		http.Handle("/", &app.Handler{
			Name:        "go_h5app_demo",
			Description: "first demo",
		})
		log.Println("http://localhost:8080")
		if e := http.ListenAndServe(":8080", nil); e != nil {
			log.Panicln(e)
		}
	} else if *t == "offline" {
		e := app.GenerateStaticWebsite("dist", &app.Handler{
			Name:        "go_h5app_demo",
			Description: "first demo",
		})
		if e != nil {
			log.Panicln(e)
		} else {
			log.Println("生成到 ./dist")
		}
	} else {
		panic("type: online/offline")
	}
}
