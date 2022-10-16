package main

const GlobalLimit = 100

const MaxCacheSize int = 10 * GlobalLimit

const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}
func cacheSet(key string, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

func getBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)

}

func setBook(isbn string, name string) {
	cacheSet(CacheKeyBook+isbn, name)
}
