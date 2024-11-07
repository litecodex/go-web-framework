package rsa

import (
	"fmt"
	"testing"
)

const rsaPublicKey string = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAttGE3BNuboYS/taUqOK1Zj2nxu+v3+H+Pv0r7rF+M6IIT3UhXxisaB6qwzo1mxIY73i+UNkltH5fUm3G7mCpIb9gfazMJChbR+6aPJM/cbf48aOFTjQG64ejWVS0T/+487CtXgutEFLD6xetqv5UpraNi1qENSVIhIRDOxIY0OSPTDJnldu2kwy/FUAnmC1c4O4SRaU8cGm4zBhmGdHG5M44gq+3nWF7wNnXoEtqppnA9fVHg2sM2fkO+sFRktsS7UYVr64VI1aAzToOF71GkKnQplh3dIY03UBTkpNbxG3/Tld6P2+EVfM3NrRDMrgW0QDEmrAKyIIV0p2MTJ0xywIDAQAB"
const rsaPrivateKey string = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC20YTcE25uhhL+1pSo4rVmPafG76/f4f4+/SvusX4zoghPdSFfGKxoHqrDOjWbEhjveL5Q2SW0fl9SbcbuYKkhv2B9rMwkKFtH7po8kz9xt/jxo4VONAbrh6NZVLRP/7jzsK1eC60QUsPrF62q/lSmto2LWoQ1JUiEhEM7EhjQ5I9MMmeV27aTDL8VQCeYLVzg7hJFpTxwabjMGGYZ0cbkzjiCr7edYXvA2degS2qmmcD19UeDawzZ+Q76wVGS2xLtRhWvrhUjVoDNOg4XvUaQqdCmWHd0hjTdQFOSk1vEbf9OV3o/b4RV8zc2tEMyuBbRAMSasArIghXSnYxMnTHLAgMBAAECggEAaasAmBPDKK7mG9X7ZwJixw2sBBhWF5mQUugSlIyS3VUyaHrTJxwjyqqvGNh0U4VKVF/94M0iNgk1H3fEG9RS7ean5vwRonSRDiqji4+whBJKGaDiVClONqTXjbKf5f1w8amVC17EUUMFasTs5IDMfO/XMEzJTc6W88Fe+q2jGlmHAqjsNF3GEaMCaU1O146nME78AHmDdSi+oFowbuZaU7sGrCgiipI9LcFpBQPWbIz1vqIiXkhgSUzWzUQ3OFupMmSASork3jesd4JhWgrsAUkYQZzUkK0MB1G1fgjE5/hauQtDIV4IUjXtzYJCVOYyYkUAs7S1UOWdJpd8y5sqAQKBgQDiVDKmVay/AszjuUGOhLkx36wkgmj2hATM3gQDHJy6aKMJtFuF2+2IeAklbK6HWtV+/Cioc7XYwqFekZYK+sQx1zELxYdbjjObZoBglYb7bdpVvJx9k4Z5LHK1tnBlsV0pL5Csrcb7JwvdrMOrm5rq5CXVkg/3hcZ9WhkXSg03KwKBgQDOyRIXO67C6F5uOeGQVqyRFmjTpxWJIsF0wyVpeDbg8sCMq01x6cDjHm/7M10fFvnSMrsmw0t2CiuEVU5/Ae4adn8Y19sJibKjVrrOrMlrOPFeZvg3o3YoqDjZmp73sQJv0aOGK9hfVbwLEQyLihSiU1A8V6+/UCzgPcVd5zSf4QKBgAfUN8hPMGGPHD1IXD8s0icqgI7mv/C/Eldv2p2s8LL2CaW9Smkv+WB+Hnrf7o2aE8aHvHRPRFwSJ3jY+mK41+6NbhHlLFB7c8eNXSV6Jqgt3Z6XnqYtYzpv0iv8+phZ8UoKbiu6+yYW7K8nWcFm6Y30hGaF3e2HAB237yRCGIDvAoGBAKXeQYWjWQ58x+pQwW8/JxMGT7Weq7ahy132ftb4F8Ue03bCnc+G+jL0Ikz0KXkbu+5wxRBVzPz1MWzn5JwaFzzg2hg6ZTdkXYeJtTS9Ap/gQDRCEk7G7qu0LE3YKjVypgq9tRaVquzl331dApwXeW+vtoeAqreh/y0sP1mQcPThAoGANfeTRUKw0N36vR/UwqC1in3WEKXkjjJ7eKnn0Iaz/e+xLrnCBEUo/14Vnw++zp15S0U8EnQmzZqbRn3fydzpA59ucI1KNdGo6WtjjyeDIMhAKjK5jkJXJqBM0+LHF6F++iQv0ZEgGKGObaZ7uSqFlOcrr8/dB0k5wCqaGx0wTUg="

func TestEncryptWithPublicKey(t *testing.T) {
	data3 := "abc"
	encryptData, err := RSAEncrypt(data3, rsaPublicKey)
	if err != nil {
		fmt.Println("加密错误信息：", err)
		return
	}
	fmt.Println("加密数据：", encryptData)

	//encryptData2 := "f4BAPmM9VhL5Plxq9w18QG2rD4A0hkD+nMgnWV1Uiksulv6Xdenn8qR8+x2bngD/2YqbGMqgntxbqA77TcJuqCI9qXp91Stf/8x0wbmG8+JkVmrT+RzDWF7oDgwbfwSxjqnQ9C5PtSRP/oCMcGDVn+coTaUH0j60kR/HkUBjQjnGi1E9fUVhzyZV+4JrIhYEpB40dXLpDNiLVNADu6c9kB6toWRete5Bl4fSRLrwMTowaVghvvLG8D4m7aTXTty3MuXAsZU3PC+YLmIZ7nGkHcWISoPmjbOS//CrXX8e7qJQ8hxH+3e0QdnB5UunQPbuo8s58bMTBLAaxxESqwwkjg=="
	data, err := RSADecryptLongText(encryptData, rsaPrivateKey)
	if err != nil {
		fmt.Println("解密错误信息：", err)
		return
	}
	fmt.Println("解密数据: ", data)
}