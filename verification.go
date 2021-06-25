package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"strings"
)


func verifyChromedp(url string) {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if _, ok := ev.(*page.EventJavascriptDialogOpening); ok {
			//fmt.Println("Found XSS [+]	" + url)
			xssFound = append(xssFound, url)

			t := page.HandleJavaScriptDialog(true)
			go func() {
				if err := chromedp.Run(ctx, t); err != nil {
					log.Fatal(err)
				}
			}()
		} else {
			//fmt.Println("No XSS")
		}
	})

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
	)
	if err != nil {
		log.Fatal(err)
	}

}

func VerifyReflection(body, payload string) bool {
	if strings.Contains(body, payload) {
		return true
	}
	return false
}

func VerifyDOM(s string) bool { //(body io.ReadCloser) bool {

	body := ioutil.NopCloser(strings.NewReader(s)) // r type is io.ReadCloser
	defer body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(body)
	check := false
	if err != nil {
		fmt.Println(err)
		return false
	}
	// Find the review items
	doc.Find(".xss").Each(func(i int, s *goquery.Selection) {
		check = true
	})
	if !check {
		doc.Find("xss").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			check = true
		})
	}
	return check
}
