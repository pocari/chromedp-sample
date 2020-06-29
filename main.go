package main

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

func newChromedpContext(ctx context.Context, headless bool) (context.Context, context.CancelFunc) {
	var opts []chromedp.ExecAllocatorOption
	for _, opt := range chromedp.DefaultExecAllocatorOptions {
		opts = append(opts, opt)
	}
	if !headless {
		opts = append(opts,
			chromedp.Flag("headless", false),
			chromedp.Flag("hide-scrollbars", false),
			chromedp.Flag("mute-audio", false),
		)
	}

	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, opts...)
	ctx, cancel := chromedp.NewContext(allocCtx)

	return ctx, func() {
		cancel()
		allocCancel()
	}
}

func main() {
	ctx, cancel := newChromedpContext(context.Background(), false)
	defer cancel()

	var res string
	err := chromedp.Run(
		ctx,
		chromedp.Navigate("https://golang.org/pkg/time/"),
		chromedp.Text("#pkg-overview", &res, chromedp.NodeVisible, chromedp.ByID),
	)

	if err != nil {
		panic(err)
	}

	log.Println(strings.TrimSpace(res))

}
