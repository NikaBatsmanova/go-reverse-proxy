package utils

import (
	"github.com/go-resty/resty/v2"
	"go-reverse-proxy/project/cache"
	"log"
	"os"
	"time"
)

func CheckCache(key string) *resty.Response {
	if page, ok := cache.M.Get(key); ok {
		return page.(*resty.Response)
	} else {
		if cache.M.Len() == 4 {
			cache.M.Remove(os.Getenv("number"))
		}
		page, err := GetPage(key)
		if err != nil {
			log.Fatal(err)
		}
		err = cache.M.Set(key, page, time.Minute*20)
		if err != nil {
			log.Fatal(err)
		}
		if cache.M.Len() == 4 {
			err = os.Setenv("number", key)
			if err != nil {
				return nil
			}
		}
		return page
	}
}
func GetPage(url string) (*resty.Response, error) {
	client := resty.New()
	res, err := client.R().EnableTrace().Get(url)
	return res, err
}
