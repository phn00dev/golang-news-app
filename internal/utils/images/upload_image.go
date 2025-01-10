package images

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadImage(ctx *fiber.Ctx, inputName, publicPath, imagePath string) (*string, error) {

	file, err := ctx.FormFile(inputName)
	if err != nil {
		return nil, err
	}
	// check file extention
	fileExt := GetLastXChars(file.Filename, 3)
	if fileExt == "peg" {
		fileExt = "j" + fileExt
	}
	// file Ext validate -->(png , jpg, webm,peg)
	isSupported := CheckExtention(fileExt)
	if !isSupported {
		return nil, errors.New("file formady nädogry")
	}
	fullPath := fmt.Sprintf("%s/%s", publicPath, imagePath)
	err = os.MkdirAll(fullPath, 0755)
	if err != nil {
		return nil, err
	}
	uniqueId := uuid.New()
	filename := fmt.Sprintf("%s.%s", strings.Replace(uniqueId.String(), "-", "", -1), fileExt)
	filePath := fmt.Sprintf("%s/%s", fullPath, filename)
	err = ctx.SaveFile(file, filePath)
	if err != nil {
		return nil, err
	}
	dbPath := fmt.Sprintf("public/%s/%s", imagePath, filename)
	return &dbPath, nil
}
