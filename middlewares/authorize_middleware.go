package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/casbin/casbin/v3"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RBAC struct {
	Enforcer *casbin.Enforcer
}

func NewRBAC(db *gorm.DB) *RBAC {
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		log.Fatal(err)
	}

	root, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	modelPath := filepath.Join(
		root,
		"rbac_model.conf",
	)

	enforcer, err := casbin.NewEnforcer(
		modelPath,
		adapter,
	)

	if err != nil {
		log.Fatal(err)
	}

	return &RBAC{
		Enforcer: enforcer,
	}
}

func (r *RBAC) GetEnforcer() *casbin.Enforcer {
	return r.Enforcer
}

func (r *RBAC) AuthorizeMiddleware(resource, action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetString("user_id")
		email := c.GetString("email")

		ok, err := r.Enforcer.Enforce(
			userId,
			resource,
			action,
		)

		fmt.Println(
			"user: ",
			email,
			userId,
			resource,
			action,
		)

		if err != nil || !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "forbidden",
			})
			return
		}

		c.Next()
	}
}
