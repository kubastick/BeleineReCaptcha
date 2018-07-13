# BeleineReCaptcha
ReCaptcha component for BeleineFramework
Usage:
```
//Create new ReCaptcha struct
recaptcha := NewReCaptcha("6LdLB2QUAAAAAK1TEGRMjZ3LeUR70Dkme335UH7-", "6LdLB2QUAAAAAMUJE9U_MFT3MXKT7KHZ0pKup7jo")
//Verify response key
response, _ := recaptcha.VerifyKey("responseKey")
//Verify response
if response {
    //Succes
}
```
