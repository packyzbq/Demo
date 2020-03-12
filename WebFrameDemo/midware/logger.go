package midware

import (
	"WebFrameDemo/frame"
	"log"
	"time"
)

func Logger() frame.HandlerFunc {
	return func(c *frame.Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
