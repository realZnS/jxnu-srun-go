package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"
)

func login() {
	ip := get_ip()
	token := get_token(ip)
	info := get_info(ip)

	x := get_xencode(info, token)
	if verbose {
		fmt.Printf("xencode:\n%s\n", x)
	}
	info = "{SRBX1}" + get_base64(x)
	if verbose {
		fmt.Printf("base64:\n%s\n", info)
	}
	hmd5 := get_md5(config.Password, token)
	if verbose {
		fmt.Printf("md5:\n%s\n", hmd5)
	}

	chkstr := token + config.Username +
		token + hmd5 +
		token + config.Ac_id +
		token + ip +
		token + config.N +
		token + config.Type +
		token + info

	chksum := get_sha1(chkstr)
	if verbose {
		fmt.Printf("sha1:\n%s\n", chksum)
	}

	Url, _ := url.Parse("http://172.16.8.6/cgi-bin/srun_portal")
	params := url.Values{
		"callback":     {"jQuery112408106679898811244_" + strconv.FormatInt(time.Now().Unix()/1000000, 10)},
		"action":       {"login"},
		"username":     {config.Username},
		"password":     {"{MD5}" + hmd5},
		"ac_id":        {config.Ac_id},
		"ip":           {ip},
		"chksum":       {chksum},
		"info":         {info},
		"n":            {config.N},
		"type":         {config.Type},
		"os":           {"Windows+11"},
		"name":         {"Windows"},
		"double_stack": {"0"},
		"_":            {strconv.FormatInt(time.Now().Unix()/1000000, 10)},
	}
	Url.RawQuery = params.Encode()
	// if verbose {
	// 	fmt.Println("Url:", Url.String())
	// }

	var client = &http.Client{
		Timeout: time.Second * 5,
	}
	rqst, _ := http.NewRequest("GET", Url.String(), nil)
	rqst.Header = config.Header
	rsps, err := client.Do(rqst)
	if err != nil {
		fmt.Println("Request failed:", err)
		os.Exit(1)
	}
	defer rsps.Body.Close()

	body, err := ioutil.ReadAll(rsps.Body)
	if err != nil {
		fmt.Println("Read body failed:", err)
		os.Exit(1)
	}
	fmt.Println(string(body))
}

func get_ip() string {
	var client = &http.Client{
		Timeout: time.Second * 5,
	}
	rqst, _ := http.NewRequest("GET", "http://172.16.8.6", nil)
	rqst.Header = config.Header
	rsps, err := client.Do(rqst)
	if err != nil {
		fmt.Println("Request failed:\n", err)
		os.Exit(1)
	}
	defer rsps.Body.Close()

	body, err := ioutil.ReadAll(rsps.Body)
	if err != nil {
		fmt.Println("Read body failed:", err)
		os.Exit(1)
	}

	re := regexp.MustCompile("id=\"user_ip\" value=\"(.*?)\"")
	ip := re.FindStringSubmatch(string(body))[1]

	fmt.Println("ip:", ip)
	return ip
}

func get_token(ip string) string {
	Url, _ := url.Parse("http://172.16.8.6/cgi-bin/get_challenge")
	params := url.Values{
		"callback": {"jQuery1124049533407110317169_" + strconv.FormatInt(time.Now().Unix()/1000000, 10)},
		"username": {config.Username},
		"ip":       {ip},
		"_":        {strconv.FormatInt(time.Now().Unix()/1000000, 10)},
	}
	/*
		params.Add("callback", "jQuery112404953340710317169_"+strconv.FormatInt(time.Now().UnixNano()/1000000, 10))
		params.Add("username", username)
		params.Add("ip", ip)
		params.Add("_", strconv.FormatInt(time.Now().UnixNano()/1000000, 10))
	*/
	Url.RawQuery = params.Encode()

	var client = &http.Client{
		Timeout: time.Second * 5,
	}
	rqst, _ := http.NewRequest("GET", Url.String(), nil)
	rqst.Header = config.Header
	rsps, err := client.Do(rqst)
	if err != nil {
		fmt.Println("Request failed:", err)
		os.Exit(1)
	}
	defer rsps.Body.Close()

	body, err := ioutil.ReadAll(rsps.Body)
	if err != nil {
		fmt.Println("Read body failed:", err)
		os.Exit(1)
	}
	re := regexp.MustCompile("\"challenge\":\"(.*?)\"")
	token := re.FindStringSubmatch(string(body))[1]

	fmt.Println("token:", token)
	return token
}

func get_info(ip string) string {
	info := "{\"username\":\"" + config.Username +
		"\",\"password\":\"" + config.Password +
		"\",\"ip\":\"" + ip +
		"\",\"acid\":\"" + config.Ac_id +
		"\",\"enc_ver\":\"" + config.Enc +
		"\"}"
	if verbose {
		fmt.Println("info:", info)
	}
	return info
}
