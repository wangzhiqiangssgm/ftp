package main

import (
	"flag"
	"fmt"
	"ftp/ftp/ftp"
	"os"
	"path/filepath"
)

const ip = "172.16.8.141" //ip地址
const port = "21"         //端口

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fmt.Println(root)

	ftp, err := ftp.Connect(ip + ":" + port) //连接
	if err != nil {
		fmt.Println(err.Error()) //打印输入错误日志
	}
	err = ftp.Login("admin", "1") //登录
	if err != nil {
		fmt.Println(err.Error()) //打印输入错误日志
	}

	dir, _ := ftp.CurrentDir() //获取当前目录
	fmt.Println(dir)
	ftp.MakeDir("/Photo")   //创建目录
	ftp.ChangeDir("/Photo") //切换目录
	dir, _ = ftp.CurrentDir()

	fpath := root + "files" //文件路径
	err = filepath.Walk(fpath, func(ipath string, f os.FileInfo, err error) error {
		//遍历
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println(ipath)
		file, er := os.Open(ipath)
		if er != nil {
			fmt.Println(er.Error())
		}

		defer file.Close()
		err = ftp.Stor(f.Name(), file)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("成功上传文件:", f.Name())
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}

	ftp.Logout()
	ftp.Quit()

}
