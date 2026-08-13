package rule

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/target/goalert/permission"
)

func TestRuleAdminPermissions(t *testing.T) {
	userCtx := permission.UserContext(context.Background(), "user-id", permission.RoleUser)
	adminCtx := permission.UserContext(context.Background(), "admin-id", permission.RoleAdmin)

	// Verify that regular user cannot pass Admin permission check for schedule rules
	err := permission.LimitCheckAny(userCtx, permission.Admin)
	assert.Error(t, err, "regular user should not have admin permissions for schedule rules")

	// Verify that admin can pass Admin permission check
	err = permission.LimitCheckAny(adminCtx, permission.Admin)
	assert.NoError(t, err, "admin should have admin permissions for schedule rules")
}
