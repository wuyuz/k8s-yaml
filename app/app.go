package app

import (
	"go.uber.org/zap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
)

var (
	Log *zap.Logger
	App *fiber.App

	TemplateEngine *django.Engine
)

