package models

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// 創建store，用memory存。這個方法僅限於單機部署

//	若多機部署，要使用redis
var store = base64Captcha.DefaultMemStore

func MakeCaptcha() (string, string, string, error) {
	var driver base64Captcha.Driver
	//	可以參考官網的參數，顯示驗證碼的外觀
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driver = driverString.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, answer, err := c.Generate()
	return id, b64s, answer, err
}

// 驗證答案是否正確
func VerifyCaptcha(id, VerifyValue string) bool {
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}
