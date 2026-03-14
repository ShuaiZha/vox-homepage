package other

import (
	"vox-homepage/global"
	"vox-homepage/lib/cache"
	"time"
)

func InitVerifyCodeCachePool() cache.Cacher[string] {
	return global.NewCache[string](10*time.Minute, 10*time.Minute, "VerifyCodeCachePool")

}
