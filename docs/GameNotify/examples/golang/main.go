package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

func main() {
	data := "{\"Amount\":0.99,\"ExtraData\":\"{\\\"zoneId\\\":1000,\\\"gameGoodId\\\":\\\"4001\\\",\\\"money\\\":9900,\\\"roleName\\\":\\\"\\\",\\\"notifyUri\\\":\\\"http://localhost:8083/lt_charge\\\",\\\"roleId\\\":9192,\\\"extInfo\\\":2,\\\"goodGearId\\\":1,\\\"sdkProductId\\\":\\\"com.yostaren.revivedwitch.diamonds6\\\",\\\"productName\\\":\\\"6钻石\\\",\\\"orderId\\\":\\\"1229260561000\\\",\\\"ratio\\\":100}\",\"OrderID\":\"140088917161212164754\",\"ProductID\":\"com.yostaren.revivedwitch.diamonds6\",\"Type\":\"delivery\",\"UID\":\"1376172378933899192204\"}"
	sign := "QPsWdAq0Ywzh4CfvdUoCkZKsUCYfLIKuEhWZvymPm+ebU3QokcxWWN9JVzBuO92pob5qdqKGSkbvustnrNx5h39BNvXRaHRD5CuMlXNKG42vuzp+Dj7rVrIzhQPw8u8r4wvF1kRZ6FGbOWrqz9SsObvjQPKBCmtl7wNsuREEUfE="
	publicKey := "-----BEGIN PUBLIC KEY-----公钥内容xxx-----END PUBLIC KEY-----\n"
	ok := RsaVerifySha2(data, sign, publicKey)
	print(ok)
}

// RsaVerifySha2 RSA Sha256 公钥验签
// data: 原始数据
// sign: 收据密文
// publicKey: RSA公钥 包含 -----BEGIN PUBLIC KEY-----
func RsaVerifySha2(data string, sign string, publicKey string) bool {
	if publicKey == "" {
		return false
	}
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return false
	}
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false
	}
	pk := publicInterface.(*rsa.PublicKey)
	decodeSign, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false
	}
	sh2 := sha256.New()
	sh2.Write([]byte(data))
	hashed := sh2.Sum(nil)
	result := rsa.VerifyPKCS1v15(pk, crypto.SHA256, hashed, decodeSign)
	return result == nil
}
