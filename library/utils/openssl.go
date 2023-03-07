package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/forgoer/openssl"
)

type RsaOpenssl struct {
	PublicKey  []byte
	PrivateKey []byte
}

func NewRsaOpenssl(publicKey, privateKey []byte) *RsaOpenssl {
	return &RsaOpenssl{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}
}

// 加密
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem // 生成公钥
func (r *RsaOpenssl) RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(r.PublicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
//openssl genrsa -out rsa_private_key.pem 1024 // 生成私钥
func (r *RsaOpenssl) RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(r.PrivateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// RSA签名 - 私钥
func (r *RsaOpenssl) SignatureRSA(plainText []byte) []byte {
	//1. 使用pem对数据解码, 得到了pem.Block结构体变量
	block, _ := pem.Decode(r.PrivateKey)
	//2. x509将数据解析成私钥结构体 -> 得到了私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//3. 创建一个哈希对象 -> md5/sha1 -> sha512
	// sha512.Sum512()
	myhash := sha512.New()
	//4. 给哈希对象添加数据
	myhash.Write(plainText)
	//5. 计算哈希值
	hashText := myhash.Sum(nil)
	//6. 使用rsa中的函数对散列值签名
	sigText, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, hashText)
	if err != nil {
		panic(err)
	}
	return sigText
}

// RSA签名验证
func (r *RsaOpenssl) VerifyRSA(plainText, sigText []byte) bool {
	//1. 使用pem解码 -> 得到pem.Block结构体变量
	block, _ := pem.Decode(r.PublicKey)
	//2. 使用x509对pem.Block中的Bytes变量中的数据进行解析 ->  得到一接口
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//3. 进行类型断言 -> 得到了公钥结构体
	publicKey := pubInterface.(*rsa.PublicKey)
	//4. 对原始消息进行哈希运算(和签名使用的哈希算法一致) -> 散列值
	hashText := sha512.Sum512(plainText)
	//5. 签名认证 - rsa中的函数
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, hashText[:], sigText)
	if err == nil {
		return true
	}
	return false
}

type AesOpenssl struct {
	Key []byte
	IV  []byte
}

func NewAesOpenssl(key, iv []byte) *AesOpenssl {
	return &AesOpenssl{
		Key: key,
		IV:  iv,
	}
}

func (a *AesOpenssl) AesCBCEncrypt(data []byte) (string, error) {
	dst, err := openssl.AesCBCEncrypt(data, a.Key, a.IV, openssl.PKCS7_PADDING)
	if err != nil {
		return "", errors.New("AesCBCEncrypt error")
	}

	return base64.StdEncoding.EncodeToString(dst), nil
}

func (a AesOpenssl) AesCBCDecrypt(data []byte) ([]byte, error) {

	dst, err := openssl.AesCBCDecrypt(data, a.Key, a.IV, openssl.PKCS7_PADDING)
	if err != nil {
		return nil, errors.New("AesCBCDecrypt error")
	}
	return dst, nil
}
