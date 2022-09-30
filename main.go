package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

var doubleEncode bool

func main() {
	author := cli.Author{
		Name: "无在无不在",
	}
	app := &cli.App{
		Name:      "urlencode",
		Usage:     "urlencode",
		UsageText: "echo 'http://target.com?id=1' | urlencode ",
		Version:   "v0.1",
		Authors:   []*cli.Author{&author},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "double encode",
				Aliases:     []string{"d"},
				Usage:       "double encode",
				Value:       false,
				Destination: &doubleEncode,
			},
		},
		Action: run,
	}
	//启动app
	if err := app.Run(os.Args); err != nil {
		logrus.Error(err)
	}
}

func run(c *cli.Context) (err error) {
	//从标准输入中读
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//获取输入的内容
		line := scanner.Text()
		//URL 编码
		result := url.QueryEscape(line)
		if doubleEncode {
			result = url.QueryEscape(result)
		}
		//输出到标准输出
		fmt.Println(result)
	}
	return nil
}
