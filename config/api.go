package config

import "os"

var WAAPIQRLogin string = "https://api.wa.my.id/api/whatsauth/request"

var WAAPIMessage string = "https://api.wa.my.id/api/v2/send/message/text"

var WAAPIDocMessage string = "https://api.wa.my.id/api/send/message/document"

var WAAPIImageMessage string = "https://api.wa.my.id/api/send/message/document"

var WAAPITextMessage string = "https://api.wa.my.id/api/v2/send/message/text"

var WebHookBOTAPI string = "https://api.wa.my.id/api/signup"

var WAAPIGetToken string = "https://api.wa.my.id/api/signup"

var WAAPIGetDevice string = "https://api.wa.my.id/api/device/"

var PublicKeyWhatsAuth = "0d6171e848ee9efe0eca37a10813d12ecc9930d6f9b11d7ea594cac48648f022"

var WAAPIToken string = ""

var PRIVATEKEY string = "e4cb06d20bcce42bf4ac16c9b056bfaf1c6a5168c24692b38eb46d551777dc4147db091df55d64499fdf2ca85504ac4d320c4c645c9bef75efac0494314cae94"

var APIGETPDLMS string = "https://pamongdesa.kemendagri.go.id/webservice/public/user/get-by-phone?number="

var APITOKENPD string = os.Getenv("PDTOKEN")

var PUBLICKEY string = "47db091df55d64499fdf2ca85504ac4d320c4c645c9bef75efac0494314cae94"