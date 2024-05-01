package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type KendaraanController struct {
	//Dependent services
}

var kendaraanController := *KendaraanController

func NewKendaraanController() *KendaraanController {
	return &KendaraanController{
		//Inject services
	}
}

func (r *KendaraanController) Index(ctx http.Context) http.Response {
	return ctx.Response().View().Make("kendaraan/lihat.tmpl", map[string]any{
		"hallo": "dimas",
	})
}

func (r *KendaraanController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *KendaraanController) Store(ctx http.Context) http.Response {
	return nil
}

func (r *KendaraanController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *KendaraanController) Destroy(ctx http.Context) http.Response {
	return nil
}
