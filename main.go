package main

import (
	"fmt"

	"github.com/maliaga-pantoja/sftp-connect-test/src"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func ConnectToSFTP(user string, passwd string, host string, port string) {
	fmt.Println("///////////////////////////////////////////////")
	fmt.Println("starting process")

	addr := fmt.Sprintf("%s:%s", host, port)
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd),
		},
	}
	fmt.Println("trying to loggin with user " + user)
	// testing connection
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic("Fail to dial: " + err.Error())
	}
	fmt.Print("Dial ok")
	// testing sftp
	client, err := sftp.NewClient(conn)
	if err != nil {
		panic("Fail create client: " + err.Error())
	}
	fmt.Println("sftp connect")
	defer client.Close()
	fmt.Println("Reading root directory /home/" + user)
	dirs, err := client.ReadDir("/home/" + user)
	if err != nil {
		panic("Error in command gcwd " + err.Error())
	}
	for _, v := range dirs {
		fmt.Print("found: " + v.Name())
	}
	fmt.Println("process end")
	fmt.Println("///////////////////////////////////////////////")
}

func main() {
	data := src.ReadJSONUsers("./data.json")
	for _, v := range data.Users {
		ConnectToSFTP(v.User, v.Passwd, v.Host, v.Port)
	}
}
