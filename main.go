package main

import (
	"fmt"
	"strings"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	// File name arg (required to run Script)
	Domain string `short:"d" long:"domain" description:"domain name" required:"true"`
}

const (
	googleUrl = "https://www.google.com/search?q="
)

func main() {
	// Flage the args to start use it
	_, err := flags.Parse(&opts)
	// Check if the parse occurred any error
	if err != nil {
		fmt.Println("try to use -h to show the options and usage :))")
		return
	}
	// naming the options
	domain := opts.Domain
	dorks(domain)
}

func dorks(domain string) {
	domain = strings.ToLower(domain)
	d2 := strings.Replace(domain, ".", "-", 1)
	fmt.Printf("%ssite:s3.amazonaws.com %s\n", googleUrl, domain)
	fmt.Printf("%ssite:s3.amazonaws.com %s \n", googleUrl, d2)
	fmt.Printf("%ssite:*.%s inurl:url= \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:path= \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:file= \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:doc= \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:u= \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:uri= \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:redirect= \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:img_url= \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:wp-admin \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:wp-content \n", googleUrl, d2)
	fmt.Printf("%ssite:*.%s intitle:\"index of\" inurl:ftp \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s filetype:env \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s intext:\"Index of /\" +passwd \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s intext:\"Index of /\" +password.txt \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s intitle:\"index of\" passwd \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:?doc_url= \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:?document_url= \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:?imgurl= \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:?image_url= \n", googleUrl, domain)
	fmt.Printf("%ssite:*.%s inurl:?image_uri= \n", googleUrl, domain)
}
