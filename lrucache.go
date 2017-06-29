package lrucache

type lruCacheImpl struct {
	maxSize  int
	contents map[string]interface{}
	access   *history
}

func (c *lruCacheImpl) Get(key string) interface{} {
	val, ok := c.contents[key]
	if !ok {
		return nil
	}
	c.access.add(key)
	return val
}

func (c *lruCacheImpl) Set(key string, val interface{}) {
	if len(c.contents) >= c.maxSize && c.access.tail.key != key {
		delete(c.contents, c.access.tail.key)
		c.access.remove(key)
	}
	c.contents[key] = val
	c.access.add(key)
}

func LRUCache(size int) Cache {
	return &lruCacheImpl{
		maxSize:  size,
		contents: make(map[string]interface{}, size),
	}
}
