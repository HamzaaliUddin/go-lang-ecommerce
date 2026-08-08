package middleware

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

func Security() gin.HandlerFunc {
	return secure.New(secure.Config{
		SSLRedirect: false,

		FrameDeny: true,

		ContentTypeNosniff: true,

		ContentSecurityPolicy: "default-src 'self'",

		ReferrerPolicy: "strict-origin-when-cross-origin",
	})
}