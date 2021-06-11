package main

import (
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)



chromedp.ListenTarget(ctx, func(ev interface{}) {
	if _, ok := ev.(*page.EventJavascriptDialogOpening); ok {
		t := page.HandleJavaScriptDialog(false)
		go func() {
			if err := chromedp.Run(ctx, t); err != nil {
				// handle error
			}

			// ok
		}()
	}
})
