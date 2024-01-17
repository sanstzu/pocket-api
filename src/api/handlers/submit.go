package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sanstzu/pocket-api/src/api/models"
	judgeRpc "github.com/sanstzu/pocket-api/src/services/judge-rpc"
	"github.com/sanstzu/pocket-api/src/utils"
)

func Submit(c echo.Context) error {
	type requestForm struct {
		file *multipart.FileHeader `form:"file"`
	}
	var err error

	file, err := c.FormFile("file")

	if err != nil {
		fmt.Println(err)
		return err
	}

	fileContent, err := utils.ReadMultipartFile(file)

	if err != nil {
		fmt.Println(err)
		return err
	}

	var testCases []string = []string{
		"1 2",
		"3 4",
		"5 6",
	}

	resp, err := judgeRpc.Judge(&models.JudgeRequest{
		Code:     fileContent,
		Language: "cpp",
		Stdin:    testCases,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"verdict": resp.Verdict,
		"stdout":  resp.Stdout,
		"stderr":  resp.Stderr,
	})

}
