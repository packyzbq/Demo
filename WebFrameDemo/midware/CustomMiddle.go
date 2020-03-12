package midware

import (
	"WebFrameDemo/frame"
	"log"
)

func A() frame.HandlerFunc {
	return func(c *frame.Context) {
		log.Println("V2 Before handle request")
		//c.Next()
		log.Println("V2 After handle request")

	}
}
