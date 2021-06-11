package main

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"

)

func verification()  {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

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

}


