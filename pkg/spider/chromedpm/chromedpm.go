package chromedpm

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

// create chrome instance
func NewAweChromeDp(timeOutSecond time.Duration, isDebug bool, flags ...chromedp.ExecAllocatorOption) (*context.Context, context.CancelFunc) {
	// 1.chrome conf
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", !isDebug),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	options = append(options, flags...)

	// 2.NewExecAllocator
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	//defer cancel()
	ctx1, cancel := chromedp.NewContext(allocCtx)
	//defer cancel()

	// 3. create a timeout
	ctx, cancel := context.WithTimeout(ctx1, timeOutSecond*time.Second)
	//defer cancel()

	return &ctx, cancel
}
