package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
)


func verification(url string)  {

	//opts := append(chromedp.DefaultExecAllocatorOptions[:],
	//	chromedp.ProxyServer("http://127.0.0.1:8080"),
	//)
	//allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	//defer cancel()
	//ctx, cancel := chromedp.NewContext(allocCtx)
	//defer cancel()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()


	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if _, ok := ev.(*page.EventJavascriptDialogOpening); ok {
			fmt.Println("Found XSS [+]		" + url)
			t := page.HandleJavaScriptDialog(false)
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


