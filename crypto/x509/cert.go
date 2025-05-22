package x509

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"
)

func getCaCertificate(name string) (*x509.Certificate, *rsa.PrivateKey, []byte, error) {
	//根证书信息
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2020),
		Subject: pkix.Name{
			Organization: []string{"Chainup Waas CA Cert, " + name},
			Country:      []string{"China"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(100, 0, 0),
		SubjectKeyId:          []byte{1, 2, 3, 4, 5},
		BasicConstraintsValid: true,
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	//根证书自签名
	privCa, _ := rsa.GenerateKey(rand.Reader, 2048)
	//pub 用于通信时验证双方证书 ，私钥用于签发证书
	_, pub, err := createCertificate(name+" RootCa", ca, privCa, ca, nil)
	return ca, privCa, pub, err

}

func getServerCertificate(name string, ca *x509.Certificate, privCa *rsa.PrivateKey, ip string) ([]byte, []byte, error) {
	server := &x509.Certificate{
		SerialNumber: big.NewInt(2020),
		Subject: pkix.Name{
			Organization: []string{"Chainup Waas Server pem, " + name},
			Country:      []string{"China"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(100, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	//绑定域名与IP
	/*hosts := []string{"localhost", "127.0.0.1"}
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			server.IPAddresses = append(server.IPAddresses, ip)
		} else {
			server.DNSNames = append(server.DNSNames, h)
		}
	}*/

	//server.IPAddresses = []net.IP{net.ParseIP(ip)}
	server.DNSNames = []string{"*"}

	privSer, _ := rsa.GenerateKey(rand.Reader, 2048)
	return createCertificate(name+" server", server, privSer, ca, privCa)
}

func getClientCertificate(name string, ca *x509.Certificate, privCa *rsa.PrivateKey) ([]byte, []byte, error) {
	client := &x509.Certificate{
		SerialNumber: big.NewInt(2020),
		Subject: pkix.Name{
			Organization: []string{"Chainup Waas client pem, " + name},
			Country:      []string{"China"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 7},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}
	privCli, _ := rsa.GenerateKey(rand.Reader, 2048)
	return createCertificate(name+" client", client, privCli, ca, privCa)
}

/**
 * name string 证书名
 * cert 要签名的 证书信息
 * key 要签名的证书私钥
 * caCert 根证书信息  如果cert 与caCert 相同 则证书自签名
 * caKey 根证书私钥（需要根证书自签）
 * return signed priv, signed pub, error
 */
func createCertificate(name string, cert *x509.Certificate, key *rsa.PrivateKey, caCert *x509.Certificate, caKey *rsa.PrivateKey) ([]byte, []byte, error) {
	fmt.Println("=======begin generate ", name, " =========")
	defer func() {
		fmt.Println("=======end generate ", name, " =========")
	}()

	priv := key
	pub := &priv.PublicKey

	//根证书私钥，没有传根证书私钥则用自身私钥签名（自签）
	privPm := priv
	if caKey != nil {
		privPm = caKey
	}

	signedCa, err := x509.CreateCertificate(rand.Reader, cert, caCert, pub, privPm)
	if err != nil {
		fmt.Println("create signed ca failed", err)
		return nil, nil, err
	}

	//公钥
	//ca_f := name + ".pem"
	//fmt.Println("write to pem", ca_f)
	var certificate = &pem.Block{Type: "CERTIFICATE",
		Headers: map[string]string{},
		Bytes:   signedCa}

	signedCaB64 := pem.EncodeToMemory(certificate)
	//ioutil.WriteFile(ca_f, signedCaB64, 0777)

	//私钥
	//priv_f := name + ".key"
	/*privDecode, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil{
		fmt.Println("gen pkcs8 priv error", err)
		return nil, nil, err
	}*/

	privDecode := x509.MarshalPKCS1PrivateKey(priv)
	//fmt.Println("write to key", priv_f)
	//ioutil.WriteFile(priv_f, priv_b, 0777)

	var privateKey = &pem.Block{Type: "RSA PRIVATE KEY",
		Headers: map[string]string{},
		Bytes:   privDecode}
	privDecodeB64 := pem.EncodeToMemory(privateKey)
	//ioutil.WriteFile(priv_f, privDecodeB64, 0777)

	return privDecodeB64, signedCaB64, nil
}

// 生成新的CA签发 新的服务端与客户端证书
func GenerateServerAndClient(appName string, listerIp string) (caPub []byte, serPriv []byte, serPub []byte, cliPriv []byte, cliPub []byte, err error) {

	var caPriv *rsa.PrivateKey
	var ca *x509.Certificate

	ca, caPriv, caPub, err = getCaCertificate(appName)
	if err != nil {
		fmt.Println("生成ca根证书失败：", appName, err)
		return nil, nil, nil, nil, nil, err
	}

	serPriv, serPub, err = getServerCertificate(appName, ca, caPriv, listerIp)
	if err != nil {
		fmt.Println("生成服务端证书失败：", appName, err)
		return nil, nil, nil, nil, nil, err
	}

	cliPriv, cliPub, err = getClientCertificate(appName, ca, caPriv)
	if err != nil {
		fmt.Println("生成客户端证书失败：", appName, err)
		return nil, nil, nil, nil, nil, err
	}

	//签发证书扣ca私钥无用了可以丢弃， caPub放到服务端与客户端用于验证双方证书
	return caPub, serPriv, serPub, cliPriv, cliPub, nil
}
