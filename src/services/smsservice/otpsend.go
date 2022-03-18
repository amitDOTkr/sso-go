package smsservice

import (
	"context"
	"log"

	"github.com/amitdotkr/sso-go/src/entities"
	"github.com/amitdotkr/sso-go/src/global"
	"github.com/amitdotkr/sso-go/src/pb"

	// "github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func OtpSend(c *fiber.Ctx) error {

	var otpmobile entities.OtpMobile

	// validate := validator.New()

	if err := c.BodyParser(&otpmobile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": entities.Error{Type: "JSON Error", Detail: err.Error()},
		})
	}

	// if err := validate.Struct(otpmobile); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": entities.Error{Type: "Validation Error", Detail: err.Error()},
	// 	})
	// }

	cc, err := grpc.Dial(
		global.OTP_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Printf("error: %v", err)
	}

	defer cc.Close()

	gc := pb.NewOtpServiceClient(cc)

	mobile_number := &pb.OtpRequest{
		MobileNumber: otpmobile.Mobile,
	}

	res, err1 := gc.OtpSend(context.Background(), mobile_number)
	if err1 != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": entities.Error{
				Type:   "Internal Server Error",
				Detail: err1.Error(),
			},
		})
	}

	var responsebody entities.ResponseBody

	responsebody.Reason = res.Reason
	responsebody.Status = res.Status
	responsebody.StatusCode = res.StatusCode

	return c.Status(fiber.StatusOK).JSON(&responsebody)
}
