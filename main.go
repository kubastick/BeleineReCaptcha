package recaptcha

import (
	json2 "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type reCaptchaResponse struct {
	success bool
}
type ReCaptcha struct {
	id         string
	privateKey string
	publicKey  string
}

func NewReCaptcha(privateKey string, publicKey string) ReCaptcha {
	return ReCaptcha{privateKey: privateKey, publicKey: publicKey}
}

func (r *ReCaptcha) VerifyKey(response string) (bool, error) {
	form := url.Values{}
	form.Add("secret", r.privateKey)
	form.Add("response", response)
	res, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", form)
	json, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}
	var result reCaptchaResponse
	json2.Unmarshal(json, result)
	fmt.Println(string(json))
	return result.success, nil
}

func (r *ReCaptcha) Render() string {
	result := "<script src='https://www.google.com/recaptcha/api.js'></script>"
	result += fmt.Sprintf(`<div class="g-recaptcha" data-sitekey="%s"></div>`, r.publicKey)
	return result
}

func (r *ReCaptcha) RenderJS() string {
	return ""
}

//Returns JS code as string, that returns response
func (r *ReCaptcha) GetResponseJS() string {
	return "grecaptcha.getResponse()"
}

func (r *ReCaptcha) GetID() string {
	return r.id
}
