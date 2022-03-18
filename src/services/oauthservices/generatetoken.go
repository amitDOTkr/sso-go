package oauthservices

import (
	"io/ioutil"
	"time"

	"github.com/amitdotkr/sso-go/src/entities"
	"github.com/amitdotkr/sso-go/src/global"

	// "github.com/amitdotkr/sso-go/src/services/userservices"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Generatetoken(c *fiber.Ctx) error {
	// var query []primitive.M

	userId, err := global.ValidatingUser(c)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": entities.Error{Type: "Authentication Error", Detail: err.Error()},
		})
	}

	// log.Printf("userId: %v", userId)

	service := c.Query("service")

	// if err := userservices.CreateTokenPairGo(c, userId, "user"); err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": entities.Error{
	// 			Type:   "Token Generation Error",
	// 			Detail: err.Error()},
	// 	})
	// }

	token, err := CreateOauthToken(c, userId, service)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": entities.Error{
				Type:   "Token Creation Error",
				Detail: err.Error(),
			},
		})
	}

	// log.Printf("userId: %v", userId)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"oauth_token": token,
	})
}

func CreateOauthToken(c *fiber.Ctx, userid string, role string) (string, error) {

	// prvKey, err := ioutil.ReadFile(global.PRVKEY_LOC)
	prvKey, err := ioutil.ReadFile(global.PRVKEY_LOC)
	if err != nil {
		return "", err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(prvKey)
	if err != nil {
		return "", err
	}

	exp := time.Now().Add(time.Second * 15)

	atClaims := jwt.MapClaims{}
	atClaims["uid"] = userid
	atClaims["service"] = role
	atClaims["type"] = "oauth"
	atClaims["exp"] = exp.Unix()

	OauthToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims).SignedString(key)
	if err != nil {
		return "", err
	}

	// log.Printf("Secure Val: %v", ac.Secure)

	return OauthToken, nil
}
