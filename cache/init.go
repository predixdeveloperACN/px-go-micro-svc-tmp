package cache

import (
	"time"
	"os"
	"strconv"

	"github.com/allegro/bigcache"
)

var cache *bigcache.BigCache
var cacheExpiry int
var debug bool

func init() {

	var err error
	cacheExpiryStr := os.Getenv("CACHE_EXPIRY")
	cacheExpiry, err = strconv.Atoi(cacheExpiryStr)
	if err != nil {
		// if local
		cacheExpiry = 24
		return
	}

	// 24 hours cache expiry
	var eviction time.Duration
	eviction = time.Duration(cacheExpiry) * time.Minute
	config := bigcache.Config {
		// number of shards (must be a power of 2)
		Shards: 1024,
		// time after which entry can be evicted
		LifeWindow: eviction,
		// rps * lifeWindow, used only in initial memory allocation
		MaxEntriesInWindow: 1000 * 10 * 60,
		// max entry size in bytes, used only in initial memory allocation
		MaxEntrySize: 500,
		// prints information about additional memory allocation
		Verbose: true,
		// cache will not allocate more memory than this limit, value in MB
		// if value is reached then the oldest entries can be overridden for the new ones
		// 0 value means no size limit
		HardMaxCacheSize: 8192,
		// callback fired when the oldest entry is removed because of its
		// expiration time or no space left for the new entry. Default value is nil which
		// means no callback and it prevents from unwrapping the oldest entry.
		OnRemove: nil,
	}
	cache, err = bigcache.NewBigCache(config)
	if err != nil {
		panic(err)
		return
	}

	// get debug config
	debug, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		// if local
		debug = true
	}
}

func CacheExpiry() int {
	return cacheExpiry
}
