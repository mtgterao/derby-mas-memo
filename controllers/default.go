package controllers

type MainController struct {
    BaseController
}

type response struct {
    BaseResponse
    Title string `json:"title"`
}

func (c *MainController) Get() {
    res := response{Title: "dery-mas-memo"}

    c.SetResponse(res)

    c.TplName = "index.tpl"
}
