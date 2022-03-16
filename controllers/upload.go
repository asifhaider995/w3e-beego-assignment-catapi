package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"log"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
	c.Data["Website"] = "thecatapi.com"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Message"] = "Please Upload a file"
	c.TplName = "upload.tpl"
}

func (c *UploadController) Post() {
	c.Data["Website"] = "thecatapi.com"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Message"] = "error"
	upFile, header, err := c.GetFile("file")
	if upFile != nil {
		fileName := header.Filename
		log.Println("Filename: ", fileName)
		log.Println("Up File: ", upFile)
		//timeout := time.Duration(5 * time.Second)
		//client := http.Client{
		//	Timeout: timeout,
		//}
		//req, err := http.NewRequest("POST", "https://api.thecatapi.com/v1/images/upload", bufio.NewReader(upFile))
		//req.Header.Set("Content-Type", "application/json")
		//req.Header.Set("Access-Control-Allow-Origin", "*")
		//req.Header.Set("x-api-key", "dcd005e2-35a8-45ec-843d-fcdfac3a4869")
		//
		//if err != nil {
		//	log.Fatalln(err)
		//}
		//resp, err := client.Do(req)
		//
		////resp, err := http.Post("https://api.thecatapi.com/v1/images/upload", "image/jpeg", upFile)
		//if err != nil {
		//	log.Fatalln(err)
		//}
		//defer resp.Body.Close()
		//
		//body, er := ioutil.ReadAll(resp.Body)
		//if er != nil {
		//	log.Fatalln(er)
		//}
		//log.Println(string(body))
		c.Data["Message"] = "File Uploaded"
	}
	if err != nil {
		log.Fatalln(err)
	}

	c.TplName = "upload.tpl"
}
