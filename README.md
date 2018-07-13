# BeleineReCaptcha
ReCaptcha component for BeleineFramework  
Usage:
```
//Create new ReCaptcha struct
recaptcha := NewReCaptcha("privateKey", "HTML_Key")

//Attach ReCaptcha to page
page.Attach(&recaptcha)

//Verify response key
response, _ := recaptcha.VerifyKey("responseKey")

//Verify response
if response {
    //Succes
}
```
