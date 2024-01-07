package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"time"
)

/* 这个中间件的工作方式如下：
对每个请求，先检查是否已经缓存了响应。
如果找到缓存，直接返回缓存的内容，并中止后续处理。
如果没有找到缓存，设置一个自定义的 ResponseWriter 来截获响应数据。
请求处理完成后，将响应数据保存到缓存中。

注意：
这个缓存中间件示例使用了内存缓存，适用于简单场景。对于更复杂的应用，您可能需要考虑使用分布式缓存解决方案，如 Redis。
缓存策略（例如何时缓存、缓存多长时间、何时清除缓存）需要根据您的具体应用场景来设计。
*/

// 创建一个全局缓存实例
var c = cache.New(5*time.Minute, 10*time.Minute)

// CacheMiddleware 创建一个缓存中间件
func CacheMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.Request.URL.String()
		if cachedData, found := c.Get(key); found {
			// 如果缓存中存在数据，则直接返回缓存数据
			ctx.JSON(200, cachedData)
			ctx.Abort()
			return
		}

		// 保存原始的写方法
		writer := ctx.Writer
		newWriter := &cacheWriter{writer, nil}
		ctx.Writer = newWriter

		ctx.Next()

		// 保存响应数据到缓存中
		c.Set(key, newWriter.data, cache.DefaultExpiration)
	}
}

// cacheWriter 用于截获写入响应的数据
type cacheWriter struct {
	gin.ResponseWriter
	data interface{}
}

// Write 将数据写入响应，同时保存到cacheWriter的data属性中
func (w *cacheWriter) Write(b []byte) (int, error) {
	w.data = b
	return w.ResponseWriter.Write(b)
}
