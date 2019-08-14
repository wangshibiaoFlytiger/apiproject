package test

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/wumansgy/goEncrypt"
	"testing"
)

/**
测试rsa加密
*/
func TestRsa(t *testing.T) {
	//调用GetKey()就可以在本地生成一个私钥文件private.pem, 一个公钥文件public.pem
	if goEncrypt.GetRsaKey() != nil {
		fmt.Println("生成公钥文件和私钥文件失败")
	}
	const privateKey = `-----BEGIN  WUMAN RSA PRIVATE KEY -----
MIIEpgIBAAKCAQEA5xKzsoh5vmLjQG8TI6E6p4mCmKlh8Yz/wVfmYydiCEv3wDAQ
qFE3Ivl4Z+z+jiFEqLkjsvNmzPxJOAZTKq3k5OfyYFO28jRmdqM+UnjRcDzh4kku
T4zuZ5OgmOu6CBCt1BRWTqhg/QbUJngC7V9pL4//Z4IEfyoKr9fHgeo0NzfzpAZR
240F081iO9/Gak5vL9syc3dYMfvU4AgHGp2romdZIXrPvBamQWCvnN3/tMxyMS5s
+MqjcWeTJOTUNyJmOwOzDp1pSdZiV12CyNuIgl3A5I05d0F0AsE/PkyHS7AquhLt
CpzvupinNm4vcfvxhCQZi2xAkTPlSHmc69GCbQIDAQABAoIBAQCqbp7f5c114VYg
ZkzFkNvESqZqlzRIh2j0YDLrm1axK+cWd90xaIW0ZDCs5p5ykpMjpzpveRuvwbz3
d6LL7erP80xQRC7Bwmh4xLro9Xf0wtMQpk1kG1yURMzFPci+vh6YOE/YQFUZC3pa
zVXLYv4gOKN24u0SJh4vyIW8itPsjQ9GF1oAjy17T9Acj3XLXp+isoYo1C8QM0t3
R+DUFu0KcRFTSD7PHkgFC2PPGGFMcKSX/B94oJ2yB4iinoktuGAcuTSk74UTCXFX
hjuKt0Btj9e/jXuYJWzZruDNELazTtpDN7qfNhII14TYzYKbepBN+ff/aKhbrG+z
TcKYPJvBAoGBAOr9U0q3AXbt9ZDDhl9B7EPIlX6URB03NAH5Qykq3mKhlV6a5Qb8
egCreHLsjbIKtcEIJb2ugXQGflmDk59Gc+H5OH5MGNukPpKMtA6KLNnsP2XpXw+x
xNr940vZG/9BtInPICAbx53QYGInlhsjavwbmJXdnDSGepKEFBBaQAQZAoGBAPu7
u10GGOYRbC9AG4hhdkSbBjr3ebCB3rx8TQcOHbJt9eBchrYOjIUHcmKS85FPyp0+
bRcGaU8uCtyvDz2KIayzzMmkMRkm9kJfJNtVIiPjrtXEjOnyDYNnwY+5Xg6yLUwQ
ccd7z2nhRZijbYllqU+bkMoTxMHSSPytFhvpSRt1AoGBAOTIapVtg1GE7/ZXRrXr
etmrqlCojYBcRJg/VHH6OLWSV9jZuW2J1kZcq1JImNPktXwQMJ5yDbsmr1D7V3hU
oXI6sBPWhsUhLYKE1rDpOi2ZY0gXY+Pl0aDcWrV31Vg5YkphjAd2xGwTeiNWI0Cc
xQZa373ZGsHQw/lt+hLJq1XRAoGBAPNn5dYf1Cl5xTq26Sho3NufblnL+w1xtfZY
1n4w5wQlkO59aQBWdwcWWUL5RqxGE+sRdpgh6efW5EfYuYrOqpr8S9LZgWQJqWop
51BxNA6x09b9MyfulaRuydl36nuZS3VWHT4++CIv4YSi2YJWG21PRaQvkzu6Yjp3
/n3Eh50VAoGBAL8D8rK2yyrsy0uGegWeWDdS27Yv8L/l45kctSmkIpF+Tlu81jwN
e23agpN5GzXuE7EgEUlU+q+0DEfWAeYL6ank8CrHyee3EYgQOcHkSQ2pElpEnJwt
EA43beE/2B0+jRMJKAFrPymrQKsv88j1Vf8XyOPRt4iX+WZgPOZORRt4
-----END  WUMAN RSA PRIVATE KEY -----
`
	const publicKey = `-----BEGIN  WUMAN  RSA PUBLIC KEY -----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA5xKzsoh5vmLjQG8TI6E6
p4mCmKlh8Yz/wVfmYydiCEv3wDAQqFE3Ivl4Z+z+jiFEqLkjsvNmzPxJOAZTKq3k
5OfyYFO28jRmdqM+UnjRcDzh4kkuT4zuZ5OgmOu6CBCt1BRWTqhg/QbUJngC7V9p
L4//Z4IEfyoKr9fHgeo0NzfzpAZR240F081iO9/Gak5vL9syc3dYMfvU4AgHGp2r
omdZIXrPvBamQWCvnN3/tMxyMS5s+MqjcWeTJOTUNyJmOwOzDp1pSdZiV12CyNuI
gl3A5I05d0F0AsE/PkyHS7AquhLtCpzvupinNm4vcfvxhCQZi2xAkTPlSHmc69GC
bQIDAQAB
-----END  WUMAN  RSA PUBLIC KEY -----
`

	privateKeyByte := []byte(privateKey)
	publicKeyByte := []byte(publicKey)

	plaintext := []byte("床前明月光，疑是地上霜,举头望明月，低头学编程")
	// 直接传入明文和公钥加密得到密文
	crypttext, err := goEncrypt.RsaEncrypt(plaintext, publicKeyByte)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("密文", hex.EncodeToString(crypttext))
	// 解密操作，直接传入密文和私钥解密操作，得到明文
	plaintextDec, errDec := goEncrypt.RsaDecrypt(crypttext, privateKeyByte)
	if errDec != nil {
		fmt.Println(errDec)
		return
	}
	fmt.Println("明文：", string(plaintextDec))
}

/**
AES的CBC模式
*/
func TestAesCbc(t *testing.T) {
	plaintext := []byte("床前明月光，疑是地上霜，举头望明月，学习go语言")
	fmt.Println("明文为：", string(plaintext))

	// 传入明文和自己定义的密钥，密钥为16字节 可以自己传入初始化向量,如果不传就使用默认的初始化向量,16字节
	cryptText, err := goEncrypt.AesCbcEncrypt(plaintext, []byte("wumansgygoaescry"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("AES的CBC模式加密后的密文为:", string(cryptText))
	fmt.Println("AES的CBC模式加密后的密文为:", base64.StdEncoding.EncodeToString(cryptText))

	// 传入密文和自己定义的密钥，需要和加密的密钥一样，不一样会报错 可以自己传入初始化向量,如果不传就使用默认的初始化向量,16字节
	newplaintext, err := goEncrypt.AesCbcDecrypt(cryptText, []byte("wumansgygoaescry"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("AES的CBC模式解密完：", string(newplaintext))
}

/**
AES的CTR模式
*/
func TestAesCrt(t *testing.T) {
	plaintext := []byte("床前明月光，疑是地上霜，举头望明月，学习go语言")
	fmt.Println("明文为：", string(plaintext))

	//传入明文和自己定义的密钥，密钥为16字节 可以自己传入初始化向量,如果不传就使用默认的初始化向量,16字节
	cryptText, err := goEncrypt.AesCtrEncrypt(plaintext, []byte("wumansgygoaesctr"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("AES的CTR模式加密后的密文为:", base64.StdEncoding.EncodeToString(cryptText))

	//传入密文和自己定义的密钥，需要和加密的密钥一样，不一样会报错 可以自己传入初始化向量,如果不传就使用默认的初始化向量,16字节
	newplaintext, err := goEncrypt.AesCtrDecrypt(cryptText, []byte("wumansgygoaesctr"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("AES的CTR模式解密完：", string(newplaintext))
}
