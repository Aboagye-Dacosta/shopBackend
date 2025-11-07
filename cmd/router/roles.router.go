package router

import (
	"github.com/Aboagye-Dacosta/shopBackend/cmd/controller"
	"github.com/Aboagye-Dacosta/shopBackend/cmd/middleware"
	"github.com/Aboagye-Dacosta/shopBackend/internal/constants"
	"github.com/Aboagye-Dacosta/shopBackend/internal/utils"
)

func (r *Router) initializeRolesRoutes(c *controller.Controller) {
	rolesRouter := r.router.PathPrefix("/roles").Subrouter()

	rolesRouter.Use(middleware.AuthMiddleWare)
	rolesRouter.HandleFunc("", utils.HandlePermissions(constants.ManageRoles, c.HttpCreateRole)).Methods("POST")
	rolesRouter.HandleFunc("", utils.HandlePermissions(constants.ManageRoles, c.HttpGetAllRoles)).Methods("GET")
	rolesRouter.HandleFunc("/{id}", utils.HandlePermissions(constants.ManageRoles, c.HttpGetRole)).Methods("GET")

	rolesRouter.HandleFunc("/{id}", utils.HandlePermissions(constants.ManageRoles, c.HttpUpdateRole)).Methods("PUT")
	rolesRouter.HandleFunc("/{id}", utils.HandlePermissions(constants.ManageRoles, c.HttpDeleteRole)).Methods("DELETE")

}
