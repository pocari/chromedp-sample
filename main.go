package main

import (
	"context"
	"fmt"

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

type result struct {
	Key1 string `json:"key1"`
}

func main() {
	ctx, cancel := newChromedpContext(context.Background(), false)
	defer cancel()

	var dummyRes []byte
	var res result
	err := chromedp.Run(
		ctx,
		chromedp.Navigate("https://google.co.jp"),
		chromedp.Evaluate(`function hoge() { return { "key1": "value1" }; };`, &dummyRes),
		chromedp.Evaluate(`hoge();`, &res),
	)

	fmt.Printf("res ... %+v\n", res)
	if err != nil {
		panic(err)
	}
}
