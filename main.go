package main

import (
	"net/http"
	admin_models "ramusaa/admin/models"
	"ramusaa/config"
)

func main() {
	admin_models.Post{}.Migrate()
	//ModelView Controller
	http.ListenAndServe(":8080", config.Routes())

}
