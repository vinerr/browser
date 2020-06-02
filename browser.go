package ibrowser

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/sirupsen/logrus"

 )

type Browser struct {
	allocCtx  context.Context
	cancelCtx context.CancelFunc
}

type ScreenShotValue struct {
	ScreenShot     bool
	ScreenShotName []string
	Buf            [][]byte
}

func NewCaptureScreenShot(ScreenShot bool) *ScreenShotValue {
	return &ScreenShotValue{
		ScreenShot: ScreenShot,
	}
}

func NewBrowser(ctx context.Context,downloadPath string) *Browser {
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoSandbox,
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("window-size", "2560,1600"),
		chromedp.UserAgent(UserAgentForChrome),
		chromedp.Env("TZ="+TimeZone),
		chromedp.UserDataDir(downloadPath+"chrome"),
	)

	// Create contexts
	allocCtx, cancelCtx := chromedp.NewExecAllocator(ctx, opts...)
	browser := &Browser{
		allocCtx:  allocCtx,
		cancelCtx: cancelCtx,
	}
	return browser
}

// NewContext 在浏览器中tab可以使用context产生
func (brow *Browser) NewContext() (context.Context, context.CancelFunc) {
	return chromedp.NewContext(brow.allocCtx, chromedp.WithLogf(logrus.Printf))
}

// Close Browser对象 context取消它。
func (brow *Browser) Close() {
	brow.cancelCtx()
}

// 抓图
func CaptureScreen(screenShotName, loginFlag string, sleep int64, screenShot *ScreenShotValue) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		if screenShot.ScreenShot {
			var buf []byte
			err := chromedp.Sleep(time.Millisecond * time.Duration(sleep)).Do(ctx)
			if err != nil {
				logrus.Errorln(loginFlag, screenShotName, err)
				return err
			}
			err = chromedp.CaptureScreenshot(&buf).Do(ctx)
			if err != nil {
				logrus.Errorln(loginFlag, screenShotName, err)
				return err
			}
			screenShot.ScreenShotName = append(screenShot.ScreenShotName, loginFlag+screenShotName)
			screenShot.Buf = append(screenShot.Buf, buf)
		}
		return nil
	})
}

// 保存图片
func SaveScreenShot(path string, screenShot *ScreenShotValue) {
	if screenShot.ScreenShot {
		for i, v := range screenShot.ScreenShotName {
			if err := ioutil.WriteFile(fmt.Sprintf("%s/%s.png", path, v), screenShot.Buf[i], 0644); err != nil {
				logrus.Errorln(err)
			}
		}
		screenShot.ScreenShotName = screenShot.ScreenShotName[:0]
		screenShot.Buf = screenShot.Buf[:0]
	}
}

// elementScreenshot takes a screenshot of a specific element.
func ElementScreenShot(sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID),
	}
}

// FullScreenShot takes a screenshot of the entire browser viewport.
//
// Liberally copied from puppeteer's source.
//
// Note: this will override the viewport emulation settings.
func FullScreenShot(quality int64, screenShotName, loginFlag string, sleep int64, screenShot *ScreenShotValue) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			if screenShot.ScreenShot {
				logrus.Debugln(screenShotName)
				err := chromedp.Sleep(time.Millisecond * time.Duration(sleep)).Do(ctx)
				if err != nil {
					logrus.Errorln(loginFlag, screenShotName, err)
					return err
				}
				// get layout metrics
				_, _, contentSize, err := page.GetLayoutMetrics().Do(ctx)
				if err != nil {
					logrus.Errorln(err)
					return err
				}
				width, height := int64(math.Ceil(contentSize.Width)), int64(math.Ceil(contentSize.Height))

				// force viewport emulation
				err = emulation.SetDeviceMetricsOverride(width, height, 1, false).
					WithScreenOrientation(&emulation.ScreenOrientation{
						Type:  emulation.OrientationTypePortraitPrimary,
						Angle: 0,
					}).
					Do(ctx)
				if err != nil {
					logrus.Errorln(err)
					return err
				}

				var buf []byte
				// capture screenShot
				buf, err = page.CaptureScreenshot().
					WithQuality(quality).
					WithClip(&page.Viewport{
						X:      contentSize.X,
						Y:      contentSize.Y,
						Width:  contentSize.Width,
						Height: contentSize.Height,
						Scale:  1,
					}).Do(ctx)
				if err != nil {
					logrus.Errorln(err)
					return err
				}
				screenShot.ScreenShotName = append(screenShot.ScreenShotName, loginFlag+screenShotName)
				screenShot.Buf = append(screenShot.Buf, buf)
			}
			return nil
		}),
	}
}
