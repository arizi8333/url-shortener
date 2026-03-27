package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type client struct {
	lastRequest time.Time
	count       int
}

var clients = make(map[string]*client)
var mu sync.Mutex

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		cl, exists := clients[ip]

		if !exists {
			clients[ip] = &client{time.Now(), 1}
			mu.Unlock()
			c.Next()
			return
		}

		if time.Since(cl.lastRequest) > time.Second {
			cl.count = 1
			cl.lastRequest = time.Now()
			mu.Unlock()
			c.Next()
			return
		}

		if cl.count >= 5 {
			mu.Unlock()
			c.JSON(429, gin.H{"error": "rate limit exceeded"})
			c.Abort()
			return
		}

		cl.count++
		mu.Unlock()
		c.Next()
	}
}
