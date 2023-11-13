package tokens
import
(
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)
type Claims struct {
	 Username string `json:"username"`
	 jwt.StandardClaims
}
func GenerateToken(username string) (string, error) {
	// Create a new Claims instance with custom fields
	claims:= &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), //Token is valid for 1hr
		},
	}

	// Creating Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signing In
	secretKey := []byte("algomox")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func VerifyToken(tokenString string) (*Claims, error) {
	// Parse the token with the secret key
	secretKey := []byte("algomox")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Checking if the token is valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid token")
}
