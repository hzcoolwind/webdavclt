package main

import (
	"flag"
	"fmt"
	"github.com/studio-b12/gowebdav"
	"os"
)

func main() {

	host := flag.String("h", "", "webdav服务器地址")
	user := flag.String("u", "", "用户名")
	password := flag.String("p", "", "密码")
	cmd := flag.String("c", "", "执行命令")
	filepath := flag.String("d", "", "文件目录")
	fname := flag.String("f", "", "文件名")
	flag.Parse()
	fmt.Printf("path:[%s], filename:[%s]\n\r", *filepath, *fname)

	c := gowebdav.NewClient(*host, *user, *password)

	switch *cmd {
	case "DELETE":
		if len(*filepath) == 0 {
			fmt.Printf("DELETE command must have a webdav path: [%v]\n\r", *filepath)
			os.Exit(1)
		}
		if len(*fname) == 0 {
			c.RemoveAll(*filepath)
			fmt.Printf("delete path [%v]\n\r", *filepath)
			os.Exit(0)
		}

		if *fname == "all" || *fname == "*" || *fname == "*.*" {
			fmt.Printf("delete all [%s]\n\r", *fname)
			files, _ := c.ReadDir(*filepath)
			for _, file := range files {
				c.Remove(*filepath + "/" + file.Name())
				fmt.Printf("delete: " + file.Name() + "!\n\r")
			}
			os.Exit(0)
		} else {
			fmt.Printf("delete [%v]\n\r", *fname)
			c.Remove(*filepath + "/" + *fname)
		}
		break

	default:
		fmt.Println("Usage:  webdavclt -h host -u user -p password  -c cmd -d path [-f file]")
		fmt.Println("        cmd: DELETE         delete path/file/files(*/*.*)")
		os.Exit(0)
	}

}
