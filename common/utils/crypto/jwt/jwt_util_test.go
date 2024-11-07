package jwt

import (
	"fmt"
	JSON "github.com/litecodex/go-web-framework/common/utils/json"
	"testing"
)

func TestJwtService_HS256_SIGN(t *testing.T) {
	var signMethod string = "HS256"
	var signKey string = "abc"
	var verifyKey string = "abc"
	jwtService := NewJwtService(signMethod, signKey, verifyKey)

	jwtBody := map[string]interface{}{
		"app_key":   "",
		"device_id": "02:00:00:00:00:00",
		"iat":       1727154933,
		"im_id":     77520983,
		"iss":       "1",
		"platform":  "iphone",
		"session":   1727154933278,
		"user_type": 0,
	}
	token := jwtService.MustCreateToken(jwtBody, int64(10), HOURS)

	fmt.Println("Generated Token:", token)

	// 验证并解析token
	payload, err := jwtService.VerifyAndParseClaims(token)
	if err != nil {
		fmt.Println("Error verifying token:", err)
		return
	}

	fmt.Println("Token Payload:", JSON.MustStringify(payload))
}

func TestJwtService_RSA256_SIGN(t *testing.T) {
	var signMethod string = "RS256"
	var signKey string = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDvCDjm5SdSA8Hb+4T+Gm2jqlKi+JeLsnNL2KnonX4JsH9tFjJ0X7pedUnl13DkABdI2jnVt4wpQRaC4qDlBXN+tg/DTreiNPL/fqmUWZOk2zu77Z93h+/LtKtLrg1Eh9qGnuxyno0wxduAOsUsVezookSwbtOGUDhdLiHkFL6gnqdKeNjPqILw9e3jBdKjJ/rW5ZKWiKFAv187RG4AdC0cS7C5PygcbbSyD/KVLBXX/zMXgDQx2U61tKrlS+8NbJWg4p+YRe3xXyHtksUAUAFba9Wt5/mW1jRpNOq/+wzmeIJrkksHTDRt4J3wW2JaWOTsAXSPVbbZiq1IPLbE/nwjAgMBAAECggEAeddZGejo2BduM7HLorLZ/DkPkl7g8KZvutOgGCBfZJUA/xv3b/Zzyz5CAtSEiNO7CrmiDVxYJ5cz4Feg59yVeJtZAZcYZ6hRzQZFbocSiU/u7OY9CPLTuqRHRHZd8PbG3yQXJn3HPns8XeqXIvhRoGtGVCDJ1YcClAy13crtOHVqZKqkZSdvQFA3deORXXNfY7wDfvvc4So7jGi+2atbWvb//B8J5GuRUkWNMFkBXjrJuWBlwzU9/PrM7tczvPixIdN5cO8XWqeY4/oqYB9DhFODyU1aE2M6kIEy+yuqMlVIRHCp7FDALZ8beddYFZYj5FMQlXa5fy2PKyZCbi/HUQKBgQD3/+xIgRH9a9Bmb0lJYsR/BKPLIkislkViqnRVLooAlcwDMIr4/JT/9o8O+LNgpNfp7FAkDkvmXkbERGNxa3gJW3jntAxohmIL3YgIZdXqatkFezO+PJM1nxi41PbS5eizNXzsDsvtD5iiWa4x1tAxmZ/REZaEH901zttZvM8qJwKBgQD2vj3WlLlHdvVQOsZlmBTpxAnXzMf3BLITkOMjZGOs3iYWMNA3X875v9GiOpBUCP0jV2dt5QcLqPIEScpyOnOW169t5EBS7yCOLi5mJNkeAikW8HgYV0suFuIA/iOeyjPMn4nhg9xPqDjcwnFmq7vhRHE0F/AJX/pzEeo3vSzHpQKBgGwds0nMkyYzCXCO1ZlbqKRjRnD5aktrW6Zu/zZfiqREqeM+F2gC3YZVW/q/65uXYdXGQw3k+avdr+ZClkPNAVC7AxOoR7yN0VKw6mwW0VJX8HLWSjGGQPsgd+ukVFKPDoqKKALVVIvtv7IPfMSXjL4C5kyD6WWCarLZkoElsf8DAoGAbXwMxGJJtEQ8pdTuo7XP0cqC85aSRDF5MuVfZBzvfY01KTOPsIJ6vKc4xdtmn2M9r6jg5Ap0DeBxQyXbBsSY9Z3O4dweDq68q1oijIBdNsuOn/cj0ukpGtJchkQ+Wf8u7OT9sWtpHo9ua8Z7uysIuvQ7pvnYMNC9uMGCRClU7WECgYEA14/V4NM7pcbWVhVLzQEXIfPbpSxlEgjzAb91Of91yMyMBJL93hikjaGmeIn1Xvl64ij148c4w7GEvZtvpCODLTa2t3Exc6WPG2u0nQ+ZSU+AOwh//aiNIitcfLmtWaoavF89FY3TW3WqEXQTnw4zx0OA71U+H0ijbdpvx+0OK9k="
	var verifyKey string = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7wg45uUnUgPB2/uE/hpto6pSoviXi7JzS9ip6J1+CbB/bRYydF+6XnVJ5ddw5AAXSNo51beMKUEWguKg5QVzfrYPw063ojTy/36plFmTpNs7u+2fd4fvy7SrS64NRIfahp7scp6NMMXbgDrFLFXs6KJEsG7ThlA4XS4h5BS+oJ6nSnjYz6iC8PXt4wXSoyf61uWSloihQL9fO0RuAHQtHEuwuT8oHG20sg/ylSwV1/8zF4A0MdlOtbSq5UvvDWyVoOKfmEXt8V8h7ZLFAFABW2vVref5ltY0aTTqv/sM5niCa5JLB0w0beCd8FtiWljk7AF0j1W22YqtSDy2xP58IwIDAQAB"
	jwtService := NewJwtService(signMethod, signKey, verifyKey)

	jwtBody := map[string]interface{}{
		"username": "haha",
	}
	token := jwtService.MustCreateToken(jwtBody, 10, "seconds")

	fmt.Println("Generated Token:", token)

	// 验证并解析token
	payload, err := jwtService.VerifyAndParseClaims(token)
	if err != nil {
		fmt.Println("Error verifying token:", err)
		return
	}

	fmt.Println("Token Payload:", JSON.MustStringify(payload))
}
