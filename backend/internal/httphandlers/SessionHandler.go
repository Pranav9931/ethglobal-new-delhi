package httphandlers

import (
	"ethglobal-nd-backend/internal/application/helpers"
	store "ethglobal-nd-backend/internal/application/repositories"
	"ethglobal-nd-backend/internal/httpdtos"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// /session (create OR join)
func SessionHandler(ctx *fiber.Ctx) error {

	token := ctx.Locals("token")
	var userId string
	if token != nil {
		tokenJwt := ctx.Locals("token").(*jwt.Token)
		claims := tokenJwt.Claims.(jwt.MapClaims)

		userId = claims["sub"].(string)
	}

	privyRepo := helpers.PrivyConfig{
		BaseUri:  "https://auth.privy.io/api/v1/users/",
		UserName: "cmg17gsej004qla0dar9k36t8",
		Password: "",
		AppId:    "cmg17gsej004qla0dar9k36t8",
	}

	user, err := privyRepo.GetPrivyUser(userId)
	if err != nil {
		fmt.Printf("Error fetching user: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch user"})
	}

	address := user.LinkedAccounts[1].Address

	session, action, ok := store.CreateOrJoinSession(address)
	if !ok {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"error": action})
	}

	return ctx.JSON(fiber.Map{
		"action":  action,
		"session": session,
	})
}

func SessionApprovalHandler(ctx *fiber.Ctx) error {

	token := ctx.Locals("token")
	var userId string
	if token != nil {
		tokenJwt := ctx.Locals("token").(*jwt.Token)
		claims := tokenJwt.Claims.(jwt.MapClaims)

		userId = claims["sub"].(string)
	}

	privyRepo := helpers.PrivyConfig{
		BaseUri:  "https://auth.privy.io/api/v1/users/",
		UserName: "cmg17gsej004qla0dar9k36t8",
		Password: "",
		AppId:    "cmg17gsej004qla0dar9k36t8",
	}

	user, err := privyRepo.GetPrivyUser(userId)
	if err != nil {
		fmt.Printf("Error fetching user: %v\n", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to fetch user"})
	}

	address := user.LinkedAccounts[1].Address

	session, err := store.ApproveTransaction(address, token.(*jwt.Token).Raw)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Check if session is fully approved
	if session.SenderApproval && session.ReceiverApproval {
		session.Status = "completed"
		fmt.Printf("TRANSACTION INITIATING..\n\n")
		sessionDTO := httpdtos.TransactionDTO{
			TransactionId: session.SessionId,
			From:          session.SenderAddress,
			SessionJwt:    session.SenderJwt,
			To:            session.ReceiverAddress,
			Value:         "100000000000000000", // 0.1 PYUSD in wei
			Currency:      "PYUSD",
		}

		err := store.SendPYUSDTransaction(user.LinkedAccounts[1].ID, sessionDTO, privyRepo.AppId, privyRepo.Password, "http://localhost:5509")
		if err != nil {
			fmt.Printf("Error sending transaction: %v\n", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to send transaction"})
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(session)
}
