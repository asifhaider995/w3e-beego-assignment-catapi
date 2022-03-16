package routers

import (
	"catapi-bee/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/find-by-breed", &controllers.CatBreedController{})
	beego.Router("/filtered-by-breed", &controllers.CatBreedController{})
	beego.Router("/upload", &controllers.UploadController{})
	beego.Router("/post-upload", &controllers.UploadController{})
	beego.Router("/find-by-category", &controllers.CatCategoryController{})
	beego.Router("/filtered-by-category", &controllers.CatCategoryController{})
}
