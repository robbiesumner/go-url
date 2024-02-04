package routes

import "github.com/labstack/echo/v4"

type Url struct {
	Url string `json:"url"`
}

func RegisterUrl(c echo.Context) error {
	var url Url
	if err := c.Bind(&url); err != nil {
		return c.JSON(400, err)
	}

	if url.Url == "" {
		return c.JSON(400, "url is required")
	}

    // TODO: do something with the url
	return c.JSON(200, url)
}
