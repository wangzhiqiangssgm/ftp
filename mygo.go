package main

import "fmt"
import "os"
import "./ftp"

func main() {
	ftp, err := ftp.Connect("172.16.8.136:21")
	if err != nil {
		fmt.Println(err)
	}
	err = ftp.Login("admin", "1")
	if err != nil {
		fmt.Println(err)
	}

	dir, err := ftp.CurrentDir()
	fmt.Println(dir)
	ftp.MakeDir("/Photo")
	ftp.ChangeDir("/Photo")
	dir, _ = ftp.CurrentDir()
	fmt.Println(dir)
	file, err := os.Open("files/1.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	err = ftp.Stor("b.txt", file)
	if err != nil {
		fmt.Println(err)
	}
	ftp.Logout()
	ftp.Quit()
	fmt.Println("success upload file:", "1.txt")
}
