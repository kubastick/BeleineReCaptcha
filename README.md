# BeleineReCaptcha
ReCaptcha component for BeleineFramework  
Usage:
```
//Create new ReCaptcha struct
recaptcha := NewReCaptcha("privateKey", "HTML_Key")
//Verify response key
response, _ := recaptcha.VerifyKey("responseKey")
//Verify response
if response {
    //Succes
}
```
