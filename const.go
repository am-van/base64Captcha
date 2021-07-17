package base64Captcha

const idLen = 20

// idChars are characters allowed in captcha id.
var idChars = []byte(TxtNumbers + TxtAlphabet)

const (
	imageStringDpi = 72.0
	//TxtNumbers chacters for numbers.
	TxtNumbers = "012346789"
	//TxtAlphabet characters for alphabet.
	TxtAlphabet = "ABCDEFGHJKMNOQRSTUVXYZabcdefghjkmnoqrstuvxyz"
	//MimeTypeImage output base64 mine-type.
	MimeTypeImage = "image/png"
)
