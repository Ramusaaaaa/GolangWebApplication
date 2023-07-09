package controllers

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"ramusaa/admin/helpers"
	"ramusaa/admin/models"
	"strconv"
	"text/template"
)

type Dasboard struct{}

func (dashboard Dasboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println(helpers.GetAlert(w, r))
	view, err := template.ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll()
	data["Alert"] = helpers.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}
func (dashboard Dasboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ExecuteTemplate(w, "index", nil)
}
func (dashboard Dasboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")

	//Upload
	r.ParseMultipartForm(18 << 20)
	file, header, err := r.FormFile("blog-picture")
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return

	}
	_, err = io.Copy(f, file)
	//Upload End
	if err != nil {
		fmt.Println(err)
		return
	}
	models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  categoryID,
		Content:     content,
		Picture_url: "uploads/" + header.Filename,
	}.Add()
	helpers.SetAlert(w, r, "Kayıt başarıyla Eklendi.Yönlendiriliyorsunuz...")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}
func (dashboard Dasboard) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	post := models.Post{}.Get(params.ByName("id"))
	post.Delete()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
func (dashboard Dasboard) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Post"] = models.Post{}.Get(params.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}
func (dashboard Dasboard) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	post := models.Post{}.Get(params.ByName("id"))
	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category"))
	content := r.FormValue("blog-content")
	is_selected := r.FormValue("is_selected")
	var picture_url string

	if is_selected == "0" {
		r.ParseMultipartForm(18 << 20)
		file, header, err := r.FormFile("blog-picture")
		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return

		}
		_, err = io.Copy(f, file)
		//Upload End
		picture_url = "uploads/" + header.Filename
		os.Remove(post.Picture_url)
	} else {
		picture_url = post.Picture_url
	}
	post.Updates(models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  categoryID,
		Content:     content,
		Picture_url: picture_url,
	})
	http.Redirect(w, r, "/admin/edit/"+params.ByName("id"), http.StatusSeeOther)
}
