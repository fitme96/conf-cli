/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

// flag变量
var (
	provincelist string
	ipblacklist  string
	passwd       string
	passwdurl    string
	site         string
	middlewares  []string
	// service      string
	servers      string
	requestchall string
	nodedist     string
	gatewaydist  string
	urlwhite     string
	ipwhite      string
	white        string
)

// 构造结构体传入template
type middl struct {
	Site         []string
	Middlewares  []string
	Provincelist []string
	Requestchall []string
	Ipblacklist  []string
	// Service      string
	Servers   []string
	Passwd    string
	Passwdurl []string
	White     string
	Urlwhite  []string
	Ipwhite   []string
}

func slic(val string) []string {
	slicva := strings.Split(val, ",")
	return slicva
}

func makedata() *middl {

	// 生成中间件列表
	if requestchall != "" {
		middlewares = append(middlewares, "RequestChallenge")
	}
	if passwd != "" && passwdurl != "" {
		middlewares = append(middlewares, "PasswordRequire")
	}
	if provincelist != "" || ipblacklist != "" {
		middlewares = append(middlewares, "IPBlacklist")
	}
	if urlwhite != "" || ipwhite != "" {
		white = "IPWhiteList"
	}
	data := middl{
		Provincelist: slic(provincelist),
		Ipblacklist:  slic(ipblacklist),
		Passwdurl:    slic(passwdurl),
		Passwd:       passwd,
		Requestchall: slic(requestchall),
		Middlewares:  middlewares,
		Site:         slic(site),
		Servers:      slic(servers),
		Urlwhite:     slic(urlwhite),
		Ipwhite:      slic(ipwhite),
		White:        white,
		// Service:      service,
	}

	return &data
}

// 生成配置文件
func Exectemp() {
	var file []string
	// 构造中间件与模板名的对应关系
	mids := map[string]string{
		"IPBlacklist":      "blacklist",
		"RequestChallenge": "requestchallenge",
		"PasswordRequire":  "passwordreq",
	}
	file = append(file, "site")
	data := makedata()
	for _, k := range data.Middlewares {
		file = append(file, mids[k])
	}
	file = append(file, "services")
	if fileinfo, _ := os.Stat(nodedist); fileinfo != nil {
		os.Remove(nodedist)
	}
	if fileinfo, _ := os.Stat(gatewaydist); fileinfo != nil {
		os.Remove(gatewaydist)
	}
	// 生成node动态配置文件
	t := template.Must(template.ParseGlob("./nodetpl/*.tpl"))
	f, err := os.OpenFile(nodedist, os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		panic(err)
	}
	for _, i := range file {
		t.ExecuteTemplate(f, i, data)
	}

	// 生成网关动态配置文件
	t1 := template.Must(template.ParseGlob("./gatewaytpl/*.tpl"))
	f1, err := os.OpenFile(gatewaydist, os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		panic(err)
	}
	if white != "" {
		t1.ExecuteTemplate(f1, "gateway", data)
		t1.ExecuteTemplate(f1, "ipwhitelist", data)
		t1.ExecuteTemplate(f1, "services", data)
	} else {
		t1.ExecuteTemplate(f1, "gatewaydefault", "")
	}

}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use: "create",
	// Args:  cobra.MinimumNArgs(1),
	Short: "生成动态配置文件",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		Exectemp()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&provincelist, "provincelist", "P", "", "陕西")
	createCmd.Flags().StringVarP(&ipblacklist, "ipblacklist", "i", "", "192.168.1.1")
	createCmd.Flags().StringVarP(&passwd, "passwd", "p", "", "sec1024. 应该与passwdurl一同开启,开启一项此中间件不生效 ")
	createCmd.Flags().StringVarP(&passwdurl, "passwdurl", "u", "", "/test. 应该与passwd一同开启,开启一项此中间件不生效")
	createCmd.Flags().StringVarP(&site, "site", "s", "", "whoami,domain.local.cn(必须的)")
	createCmd.MarkFlagRequired("site")
	// nodeCmd.Flags().StringVarP(&service, "service", "v", "", "whoami")
	createCmd.Flags().StringVarP(&servers, "servers", "S", "", "http://192.168.1.1:82(必须的)")
	createCmd.MarkFlagRequired("servers")
	createCmd.Flags().StringVarP(&requestchall, "requestchall", "r", "", "2,2,2m,false")
	createCmd.PersistentFlags().StringVarP(&nodedist, "nodedist", "d", "/tmp/dynamic.yml", "/home/traefik/conf/dynamic.yml")
	createCmd.PersistentFlags().StringVarP(&gatewaydist, "gatewaydist", "D", "/tmp/dynamic_gatway.yml", "/home/traefik/conf/dynamic_gatway.yml")
	createCmd.PersistentFlags().StringVarP(&urlwhite, "urlwhite", "W", "", "/test/,/test2/")
	createCmd.PersistentFlags().StringVarP(&ipwhite, "ipwhite", "w", "", "10.10.1.3")

}
