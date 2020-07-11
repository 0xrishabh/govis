package main

import (
	"bufio"
	"context"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	nurl "net/url"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/0xrishabh/govis/pkg/util"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type recon struct {
	Url        string
	Title      string
	ImgPath    string
	StatusCode int64
	JsUrlsList []string
}

var reconUnit recon
var db []recon

func createDirectoryIfNotExists(directory string) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.Mkdir(directory, 0744)
	}
}
func saveImage(url string, image []byte, directory string) string {
	fileName := strings.TrimRight(base64.StdEncoding.EncodeToString([]byte(url)), "=")
	filePath := path.Join(directory, "images", fileName)
	ioutil.WriteFile(filePath, image, 0644)
	return path.Join("images", fileName)
}
func main() {
	threads := flag.Int("t", 10, "No of threads to use")
	directory := flag.String("dir", "./output/", "Directory to save data inside")
	flag.Parse()

	createDirectoryIfNotExists(*directory)
	createDirectoryIfNotExists(path.Join(*directory, "images"))

	var wg sync.WaitGroup
	var urlsChan = make(chan string)

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", "1"),
	)
	parentCtx, pcancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer pcancel()

	// Taking urls from STDIN and sending them to url channel
	f := bufio.NewScanner(os.Stdin)
	go func() {
		defer close(urlsChan)
		for f.Scan() {
			urlsChan <- f.Text()
		}
	}()

	// Creatring `n` go routines to start Screenshotting
	for i := 0; i < *threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for u := range urlsChan {
				data, image, err := screenshot(u, parentCtx, reconUnit)
				if err == nil {
					data.ImgPath = saveImage(u, image, *directory)
					db = append(db, data)
				}

			}
		}()
	}
	wg.Wait()
	jx, _ := json.Marshal(db)
	reconData := []byte("var Db=" + string(jx))
	ioutil.WriteFile(path.Join(*directory, "data.js"), reconData, 0644)
	ioutil.WriteFile(path.Join(*directory, "index.html"), util.Template(), 0644)
}

func screenshot(urlString string, parentCtx context.Context, data recon) (recon, []byte, error) {
	var url, title string
	var statusCode int64
	var jsUrlsList []string
	var buf []byte
	var err error

	code := `(function(){var scripts = document.getElementsByTagName('script');var list = [];for(i=0;i<scripts.length;i++){if(scripts[i].src){list.push(scripts[i].src)};}return list;}());`
	parsedUrl, err := nurl.Parse(urlString)
	if err != nil {
		fmt.Printf("[info] %s url is malformed\n", urlString)
		return data, buf, err
	}

	ctx, cancel := chromedp.NewContext(parentCtx)
	defer cancel()
	resp, err := chromedp.RunResponse(ctx, chromedp.Tasks{
		emulation.SetDeviceMetricsOverride(int64(1280), int64(1024), 1.0, false),
		chromedp.Navigate(parsedUrl.String()), // url to load
		chromedp.Title(&title),                // save the title
		chromedp.Evaluate(code, &jsUrlsList),  // collect all the url of loaded Javscript files
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, err = page.CaptureScreenshot().WithQuality(90).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}),
	})
	if err != nil {
		fmt.Printf("[info] %s screenshot failed\n", parsedUrl)
		return data, buf, err
	}
	url = parsedUrl.String()
	statusCode = resp.Status

	data.Url = url
	data.Title = title
	data.JsUrlsList = jsUrlsList
	data.StatusCode = statusCode
	return data, buf, nil

}
