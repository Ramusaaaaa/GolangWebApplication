package config

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	admin "ramusaa/admin/controllers"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	//ADMİN
	r.GET("/admin", admin.Dasboard{}.Index)
	r.GET("/admin/yeni-ekle", admin.Dasboard{}.NewItem)
	r.POST("/admin/add", admin.Dasboard{}.Add)
	r.GET("/admin/delete/:id", admin.Dasboard{}.Delete)
	r.GET("/admin/edit/:id", admin.Dasboard{}.Edit)
	r.POST("/admin/update/:id", admin.Dasboard{}.Update)
	// SERVE FİLES
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))
	return r

}
