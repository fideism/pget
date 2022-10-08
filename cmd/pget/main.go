package main

import (
	"context"
	"fmt"
	"github.com/fideism/pget"
	"os"
)

func main() {

	//p := `E:\_ansible-2.9.15-1.el8.noarch.rpm.1/ansible-2.9.15-1.el8.noarch.rpm.1.0`
	//
	//panic(filepath.ToSlash(p))
	cli := pget.New()
	//err := cli.Save(context.Background(), `https://mirrors.aliyun.com/centos/8/storage/x86_64/ceph-nautilus/Packages/a/ansible-2.9.15-1.el8.noarch.rpm?spm=a2c6h.25603864.0.0.73896472ktubZE`, `E:\`, `ansible-2.9.15-1.el8.noarch.rpm`)
	err := cli.Save(context.Background(), `https://mirrors.aliyun.com/centos/8/storage/x86_64/ceph-nautilus/Packages/a/ansible-2.9.15-1.el8.noarch.rpm?spm=a2c6h.25603864.0.0.73896472ktubZE`, `/home/hyc/backup`, `ansible-2.9.15-1.el8.noarch.rpm`)

	if err != nil {
		if cli.Trace {
			fmt.Fprintf(os.Stderr, "Error:\n%+v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "Error:\n  %v\n", err)
		}
		os.Exit(1)
	}
}
