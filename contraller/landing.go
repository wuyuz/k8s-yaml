package contraller

import (
	"github.com/gofiber/fiber/v2"
	"k8s_yaml/lib"
)

func Landing(c *fiber.Ctx) error {
	dirname := c.Query("filename")

	rootFile, File := lib.GetFile(dirname)
	return c.Render("index", fiber.Map{
		"rootFile": rootFile,
		"File":     File,
		"pwd":dirname,
	})
}

func Loading(c *fiber.Ctx) error {
	dirname:= c.Query("filename")
	c.Set("Content-Type", "application/pdf")
	return c.Download("./data/"+dirname);
}