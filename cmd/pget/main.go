package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Code-Hex/pget"
)

func main() {
	cli := pget.New()
	err := cli.Save(context.Background(), `https://mirrors.aliyun.com/centos/8.5.2111/BaseOS/x86_64/os/isolinux/initrd.img?spm=a2c6h.25603864.0.0.50b3d2f3EWIGky`, `/home/hyc/backup`, `initrd.img`)

	if err != nil {
		if cli.Trace {
			fmt.Fprintf(os.Stderr, "Error:\n%+v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "Error:\n  %v\n", err)
		}
		os.Exit(1)
	}
}
