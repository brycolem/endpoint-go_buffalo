package actions

import (
	"go_buffalo/services"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
)

func ApplicationIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	service := services.NewApplicationService(tx)

	applications, err := service.GetAllApplications()
	if err != nil {
		return c.Error(http.StatusInternalServerError, err)
	}

	return c.Render(http.StatusOK, r.JSON(applications))
}

func ApplicationShow(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	service := services.NewApplicationService(tx)

	application, err := service.GetApplicationByID(c.Param("id"))
	if err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return c.Render(http.StatusOK, r.JSON(application))
}
