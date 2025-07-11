package rsa

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRsa(t *testing.T) {
	//RsaGenKey(2048)
	//var data = []byte("hello world")
	//encrypt := RSAEncrypt(data, []byte("public.pem"))
	//fmt.Println("加密后的数据:", string(encrypt))
	//decrypt := RSADecrypt(encrypt, []byte("private.pem"))
	//fmt.Println("解密后的数据:", string(decrypt))

	//privateKey, publicKey, err := B4RsaGenKey(1024)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("privateKey:", privateKey)
	//fmt.Println("publicKey:", publicKey)

	base64PrivKey := "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDWEFJQkFBS0JnUURHOTdCSHZIN1hNcm1WK2Y3VTFoQ3k4azVnOUszb05ENFg0QzEwUEtCOWhhemsyV01BCnBZMjNnVkFVeWxEQTlhWjBQK041N0RrbWE4SzNiREdRNXdWdUdYckxYOUN4MWtkMVdZWndIQ1UzUkVzWXVWWkcKUWtkaDBOV3RLTTZ1and0TXN0QjVBKzJHUzFVL2lsK3lvdE53ZUl0RWZLemJRalpVa3NGbHRDREJXd0lEQVFBQgpBb0dBRmVmc1ZTR1lOM3BDdTVQd0xoKzVhL0pJbGQwcWpuKzkvR1AyclM4Rm5ueEt1bEdDMkczaHlmdm5STHEzCnBGOEhhaElrcWVaTHJvSzd4L1VXQ2ZXZGY3MkIwNjBSQTFKKy9KVWJGbDdBayt2Yk9oTndybGJxaDRaTUZBeFEKbGVFQitIbUNBVjMrc3o0TjVKYnZGaU9GVGkvZHZnZWp3TEZQaVo2bi9KZ2VEcFVDUVFEY3FDdWozYjFtdHZIVApMNFVMNjRtQ3kzMzVqTUFoNi8xakVqZ1hPZ3dSeWt6SE5sa3o0MlQ1TVIzUFFvK3I3M0dEWW5KMlEvZTdmbi9PCm02QzY3YjFQQWtFQTV0WXFWbDJPVTQ5N2Fod2N3ZlhJZmVNN2Y5ZzRXa04veWhPZ3c5dlZwWjNuRE5qRVpvb3oKbWZJNXBtZk15QjBBNjdCajJVQ2dPRkdWSlNsY2x3cHdOUUpCQU53OGJmSk1hN0tOZnFpT2tYam9Tb1BsbjRMbwpYUXgrZ3BYVHBYQnBXNHFXSkRQaHB2OEhRODBFblFBMUt6a3M4RnQyYzFCZlhuQWhQbWQxSmVQdFRqa0NRSDZ6CnB1Rm8wS3BFNHpURzRYSUUrbHMrMG5YRnRJaTI2L2w2OUk4TXB5UVVtN0dOZVovZTA1djRQSFIrSFFUcUJvdFgKY2piWkpoaERqMTJxYWt2dFN0RUNRRzZKcEpUak80azgzbUt5NnhsSG1Ua2Z3SEFyWnFYeDhJRE0zMXJmYkZ2NQpHRWZxcUhIdFlTa2FMSkFVYnA1cjhHQ0NZMGk1YmxORjZjU2lTU0QvOGI0PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="
	//	publicKey: LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FERzk3Qkh2SDdYTXJtVitmN1UxaEN5OGs1Zwo5SzNvTkQ0WDRDMTBQS0I5aGF6azJXTUFwWTIzZ1ZBVXlsREE5YVowUCtONTdEa21hOEszYkRHUTV3VnVHWHJMClg5Q3gxa2QxV1lad0hDVTNSRXNZdVZaR1FrZGgwTld0S002dWp3dE1zdEI1QSsyR1MxVS9pbCt5b3ROd2VJdEUKZkt6YlFqWlVrc0ZsdENEQld3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=

	//
	//base64PubKey := "LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlHZk1BMEdDU3FHU0liM0RRRUJBUVVBQTRHTkFEQ0JpUUtCZ1FEUGxMazJORzVDeHEreldvbUl5MzM4aFZXTQp2Z0V2RjhPUmVkQ1N0cjZ1bE5YWUFWOG54Q1YyaW52c1dyczRNQXdVcFc3eitqckY4Mmkvek9lSlRXTDg5TGVuCi83dlphanZhRGc4N2JRMzI2dlJUS1JFeXVSZ0VWUlNCOUpoYS9jbnJ6Z2x4K2hVWTFlcTd3eEkyNHlsN0hpTTgKOTJ5ZnJEdWJZaWUrcnZkbHF3SURBUUFCCi0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo="
	//base64PrivKey := "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlDWEFJQkFBS0JnUURQbExrMk5HNUN4cSt6V29tSXkzMzhoVldNdmdFdkY4T1JlZENTdHI2dWxOWFlBVjhuCnhDVjJpbnZzV3JzNE1Bd1VwVzd6K2pyRjgyaS96T2VKVFdMODlMZW4vN3ZaYWp2YURnODdiUTMyNnZSVEtSRXkKdVJnRVZSU0I5SmhhL2NucnpnbHgraFVZMWVxN3d4STI0eWw3SGlNODkyeWZyRHViWWllK3J2ZGxxd0lEQVFBQgpBb0dBQjdvZzFKTktOYndqWURiSjJtMVF4RWdVRFFncG9iRUFEZmZYQkh6UlIzR001TnlMbzRlK3FVTWdicXFqCnJjVEVHY2s0Y0VIOTVUdzB4NVMyRWsybGxsQ29FWTR6RDU0T3lpU0dMa1NlenR0V1FteW1leTFuU2ZOdVFlK04KZEZ4SU5JellkNGNDdFY0WDZVcXlHZXZ4QVNyWmpUMFRmYTFIVEovV1pER2xRcEVDUVFEdjg0a0VsbjVYb2x5cQpTaHFmUXVRV3h1RzJJUUJraEJVZ1ZDd0JwZEcxSjlkNGFveUpzc2Z4UUY2c2JrM3hmcGVHQ09nbVNmQjNydllDCnptZEZ5K1BqQWtFQTNYYnd4QkNGeU1sWml5SGFrOGtQUnBkTjJ1RlgvQlVlZFBjdDFwYXVWc1VwS3AvN1B1T2wKNGZBWFNYYVdhaHpEZjdaTmR5aUtzbE5rRHNHUjduZHhtUUpBUkI1aC9vQnV1NWduTUZ5N05BMDhUVThHa3JySgpycjlrYy9vUlNDSjM1eks1VThFRHhxK1BYV2pGdDdQVXdzTUtadzJ1UWZocG9NQjFySlJGOHlXUHV3SkFOTngrCkRaNlFBR1FuTG5HUE5iVkpJeVZjWFFGcXVlM2tqakN0elVtOEpWUDhSb05YT2lTbVpLNmxNRkRSQ0ZVRFNRbWYKcGZDVVlvcExHTWFWYWFFekNRSkJBS3pIeEVWNjBTaVpRY1ExM3NxZE5RTGhVMXB0SVBQd29SMzdMdTJRM3FXLwpPRC9Lcm1USStmQjFpc3M2a0wweHZuV0ZGemM0bGc0RnVOaHd0bHRPM1VzPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="
	//// 加密前数据
	//data := []byte("Hello RSA")
	//// 假设 base64PubKey 是你存储在配置或环境变量中的 PEM 格式公钥的 base64 编码字符串
	//encrypted := B4RSAEncrypt(data, base64PubKey)
	//fmt.Println("加密后的数据:", encrypted)
	// 解密
	data, err := base64.StdEncoding.DecodeString("nCA4+Y+CooVtaFWrDTvibHX81G2HWmyCT+s+CLplERG/swhlSEDIXeEm/oyJjGCrlijzC5zEomG/erUymnxju5t4EZhl0C/62SrEjx9+Iw9Xvxv0gn2Ga3Bz1p6WEfQ1zrlJhi2PLu5vD7XkXGQLBDeUUGVBxWBz6NJk1ntrDcmecQ3hK3vT7e1Bsof1OezmpmdeI5tjwWXbmyvL5YPYuG3emhJxzztWO0BHdd8DKHc=")
	if err != nil {
		panic(err)
	}
	decrypted := B4RSADecrypt(data, base64PrivKey)
	fmt.Println("解密后的数据", string(decrypted))

}
