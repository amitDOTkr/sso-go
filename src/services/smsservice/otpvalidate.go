package smsservice

import (
	"context"
	"log"

	"github.com/amitdotkr/sso-go/src/entities"
	"github.com/amitdotkr/sso-go/src/global"
	"github.com/amitdotkr/sso-go/src/pb"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func OtpValidate(c *fiber.Ctx) error {

	var otpmobile entities.OtpMobile

	if err := c.BodyParser(&otpmobile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": entities.Error{Type: "JSON Error", Detail: err.Error()},
		})
	}

	// GRPC Dial up Connection to Internal Service
	cc, err := grpc.Dial(
		global.OTP_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("error: %v", err)
	}

	defer cc.Close()

	gc := pb.NewOtpServiceClient(cc)
	// GRPC connection code ends

	res, err := gc.OtpValidate(context.Background(), &pb.OtpValidateRequest{
		MobileNumber: otpmobile.Mobile,
		Otp:          otpmobile.Otp,
	})
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": entities.Error{
				Type:   "Internal Server Error",
				Detail: err.Error(),
			},
		})
	}

	if res.Validated {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"validated": res.Validated,
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"validated": false,
		})
	}

}
