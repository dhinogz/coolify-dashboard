package app

import (
	"github.com/dhinogz/coolify-dashboard/lib"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
)

func ProfileGet(c echo.Context) error {
	var record *models.Record = c.Get(apis.ContextAuthRecordKey).(*models.Record)
	return lib.Render(c, 200, Profile(record))
}
