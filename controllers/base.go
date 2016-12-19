/**
 * controllerの基底クラス
 *
 * SetResponseにて画面への戻り値を設定
 */
package controllers

import (
    "encoding/json"
    "github.com/astaxie/beego"
    "github.com/mtgterao/derby-mas-memo/libraries"
)

// 構造体定義
type BaseController struct {
    beego.Controller
    response map[string]interface{}
    basicResponse map[string]interface{}
}

// 前処理
func (this *BaseController) Prepare() {
    this.basicResponse = map[string]interface{}{
        "title": beego.AppConfig.String("apptitle"),
    }
}

// 各controllerごとの画面へのパラメータを設定する処理
func (this *BaseController) SetResponse(res map[string]interface{}) {
    this.response = res
}

// 描画処理
// 概ね以下の処理
// ・responseとbasicResponseをマージして画面パラメータにする
// ・beego.ControllerのRenderを呼ぶ
func (this *BaseController) Render() error {
    res := libraries.MergeMap(this.response, this.basicResponse)

    resBytes, err := json.Marshal(res)

    if err == nil {
        this.Data["Data"] = string(resBytes)
    }

    return this.Controller.Render()
}
