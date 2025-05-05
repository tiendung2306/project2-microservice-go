package middleware

import (
	"net/http"
	"project2-microservice-go/internal/auth-service/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type JWTAuthMiddleware struct {
	jwtService service.IJWTService
}

func NewJWTAuthMiddleware(jwtService service.IJWTService) *JWTAuthMiddleware {
	return &JWTAuthMiddleware{
		jwtService: jwtService,
	}
}

// AuthRequired là middleware yêu cầu authentication
func (m *JWTAuthMiddleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Kiểm tra nếu header có định dạng "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer <token>"})
			c.Abort()
			return
		}

		// Lấy token từ header
		tokenString := parts[1]

		// Xác thực token
		claims, err := m.jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Lưu thông tin user vào context để sử dụng trong các handler tiếp theo
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Set("claims", claims)

		c.Next()
	}
}
