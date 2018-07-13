package recaptcha

import "testing"

func TestReCaptcha(t *testing.T) {
	recaptcha := NewReCaptcha("6LdLB2QUAAAAAK1TEGRMjZ3LeUR70Dkme335UH7-", "6LdLB2QUAAAAAMUJE9U_MFT3MXKT7KHZ0pKup7jo")
	response, err := recaptcha.VerifyKey("wrongKey")
	if err != nil {
		t.Fail()
	}
	if response {
		t.Fail()
	}
}
