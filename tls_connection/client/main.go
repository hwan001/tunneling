package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	certPool := x509.NewCertPool()
	serverCert, err := ioutil.ReadFile("server.crt") // 'server.crt'는 서버 인증서 파일입니다.
	if err != nil {
		log.Fatalf("Failed to read server.crt: %s", err)
	}

	// 인증서를 CertPool에 추가
	if !certPool.AppendCertsFromPEM(serverCert) {
		log.Fatalf("Failed to append server cert")
	}

	conf := &tls.Config{
		InsecureSkipVerify: true,
		RootCAs:            certPool,
	}

	conn, err := tls.Dial("tcp", "localhost:443", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	println(string(buf[:n]))
}
