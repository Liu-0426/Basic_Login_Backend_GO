package controller

import (
    "github.com/gin-gonic/gin"
    "easyBackend/model"
	
	"time"
    "net/http"
	"golang.org/x/crypto/bcrypt"
	// "encoding/json"
	"github.com/golang-jwt/jwt/v5"
)
var jwtSecret = []byte("5k4g4au4ul4") // 替換為更安全的密鑰

// GenerateToken 生成 JWT
func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token 過期時間 (24 小時)
		"iat":    time.Now().Unix(),                     // Token 發行時間
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func Login(c *gin.Context) {
    var loginData struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    user, err := model.GetUserByUsername(loginData.Username)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

    c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}


func RegisterHandler(c *gin.Context) {
	var newUser model.User

	// 從請求體中解析 JSON
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 檢查使用者名稱和密碼是否提供
	if newUser.Username == "" || newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password are required"})
		return
	}

	// 加密密碼
	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	newUser.Password = hashedPassword

	// 儲存用戶資料到資料庫
	err = model.SaveUserToDB(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user to database"})
		return
	}

	// 返回成功訊息
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}