package routes

import (
	"github.com/gofiber/fiber/v2"
	"mizuserver/src/pkg/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(fiberApp *fiber.App) {
	routeGroup := fiberApp.Group("/api")

	routeGroup.Get("/entries", controllers.GetEntries)        // get entries (base/thin entries)
	routeGroup.Get("/entries/:entryId", controllers.GetEntry) // get single (full) entry

	routeGroup.Get("/resetDB", controllers.DeleteAllEntries) // get single (full) entry

}