package main

import (
	"flag"
	"os"
	"fmt"
	"bufio"
	"strings"
	"io"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func copyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer srcFile.Close()

	dstFile, err := os.Create(dst)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile);
}

// 拷贝文件的行为
func copyFileAction(src, dst string, showProgress, force bool) {
	if !force {
		if fileExists(dst) {
			fmt.Printf("%s exists override ? y/n\n", dst)
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()

			if strings.TrimSpace(string(data)) != "y" {
				return
			}
		}
	}

	copyFile(src, dst)

	if showProgress {
		fmt.Printf("'%s' -> '%s'\n", src, dst)
	}
}
func main() {
	var showProgress, force bool

	flag.BoolVar(&force, "f", false, "force copy when existing")
	flag.BoolVar(&showProgress, "v", false, "explain what is being done")

	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		return
	}

	copyFileAction(flag.Arg(0), flag.Arg(1), showProgress, force)
}
