package cache

import "github.com/OrlovEvgeny/go-mcache"

var M *mcache.CacheDriver

func InitCache() {
	M = mcache.New()
}
