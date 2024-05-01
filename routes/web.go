package routes

import (
	"goravel/app/http/controllers"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
)

func Web() {
	kendaraanController := controllers.NewKendaraanController()

	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})

	facades.Route().Resource("/welcome", kendaraanController)
	// facades.Route().Get("/welcome", func(ctx http.Context) http.Response {
	// 	return ctx.Response().View().First([]string{"kendaraan/lihat.tmpl", "lihat.tmpl"}, map[string]any{
	// 		"hallo": "dimas",
	// 	})
	// })
}
