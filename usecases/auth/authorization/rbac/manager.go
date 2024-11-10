//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package rbac

import (
	"fmt"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/entities/models"
	"github.com/weaviate/weaviate/usecases/auth/authorization"
	"github.com/weaviate/weaviate/usecases/auth/authorization/errors"
)

type manager struct {
	*casbin.SyncedCachedEnforcer
	logger logrus.FieldLogger
}

func New(casbin *casbin.SyncedCachedEnforcer, logger logrus.FieldLogger) *manager {
	return &manager{casbin, logger}
}

// Authorize will give full access (to any resource!) if the user is part of
// the admin list or no access at all if they are not
func (m *manager) Authorize(principal *models.Principal, verb string, resources ...string) error {
	if m == nil {
		return fmt.Errorf("rbac enforcer expected but not set up")
	}
	if principal == nil {
		return fmt.Errorf("user is unauthenticated")
	}

	// TODO batch enforce
	for _, resource := range resources {
		m.logger.WithFields(logrus.Fields{
			"user":     principal.Username,
			"resource": resource,
			"action":   verb,
		}).Debug("checking for role")

		allow, err := m.Enforce(principal.Username, resource, verb)
		if err != nil {
			m.logger.WithFields(logrus.Fields{
				"user":     principal.Username,
				"resource": resource,
				"action":   verb,
			}).WithError(err).Error("failed to enforce policy")
			return err
		}

		// TODO audit-log ?
		if allow {
			return nil
		}
	}

	return errors.NewForbidden(principal, verb, resources...)
}

func (m *manager) CreateRoles(roles ...*models.Role) error {
	// TODO: block overriding existing roles
	for idx := range roles {
		for _, permission := range roles[idx].Permissions {
			// TODO verify slice position to avoid panics
			domain := strings.Split(*permission.Action, "_")[1]
			verb := strings.ToUpper(string(string(*permission.Action)[0]))
			if verb == "M" {
				verb = authorization.CRUD
			}

			resource := ""
			switch domain {
			case rolesD:
				resource = authorization.Roles()[0]
			case cluster:
				resource = authorization.Cluster()
			case collections:
				resource = authorization.Collections(*permission.Collection)[0]
			case tenants:
				resource = authorization.Shards(*permission.Collection, *permission.Tenant)[0]
			case objects:
				resource = authorization.Objects(*permission.Collection, *permission.Tenant, strfmt.UUID(*permission.Object))
			}

			// TODO prefix roles names
			// roleName := fmt.Sprintf("%s%s", rolePrefix, *roles[idx].Name)
			if _, err := m.AddNamedPolicy("p", *roles[idx].Name, resource, verb, string(domain)); err != nil {
				return err
			}
		}
	}
	if err := m.SavePolicy(); err != nil {
		return err
	}
	if err := m.InvalidateCache(); err != nil {
		return err
	}
	return nil
}

func (m *manager) GetRoles(names ...string) ([]*models.Role, error) {
	roles := []*models.Role{}
	rolesMap := make(map[string][]*models.Permission)
	if len(names) == 0 {
		// get all roles
		polices, err := m.GetNamedPolicy("p")
		if err != nil {
			return nil, err
		}

		for _, policy := range polices {
			rolesMap[policy[0]] = append(rolesMap[policy[0]], permission(policy))
		}
	} else {
		for _, name := range names {
			polices, err := m.GetFilteredNamedPolicy("p", 0, name) //fmt.Sprintf("'%s' == p.sub", name)
			if err != nil {
				return nil, err
			}

			for _, policy := range polices {
				rolesMap[name] = append(rolesMap[name], permission(policy))
			}
		}
	}

	for roleName, perms := range rolesMap {
		roles = append(roles, &models.Role{
			Name:        &roleName,
			Permissions: perms,
		})
	}
	return roles, nil
}

func (m *manager) DeleteRoles(roles ...string) error {
	// TODO: block deleting built in roles
	for _, role := range roles {
		ok, err := m.RemoveFilteredNamedPolicy("p", 0, role)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("failed to remove policy %v", role)
		}
	}
	if err := m.SavePolicy(); err != nil {
		return err
	}

	return m.InvalidateCache()
}

func (m *manager) AddRolesForUser(user string, roles []string) error {
	userName := fmt.Sprintf("%s%s", userPrefix, user)
	for _, role := range roles {
		roleName := fmt.Sprintf("%s%s", rolePrefix, role)
		if _, err := m.AddRoleForUser(userName, roleName); err != nil {
			return err
		}
	}

	return m.SavePolicy()
}

func (m *manager) GetRolesForUser(user string) ([]*models.Role, error) {
	// roles, err := m.GetRoleManager().GetRoles(user)
	// return m.GetRolesForUser(user)
	return nil, nil
}

func (m *manager) GetUsersForRole(role string) ([]string, error) {
	// return m.GetUsersForRole(role)
	return nil, nil
}

func (m *manager) RevokeRolesForUser(user string, roles ...string) error {
	for _, role := range roles {
		if _, err := m.DeleteRoleForUser(user, role); err != nil {
			return err
		}
	}
	if err := m.SavePolicy(); err != nil {
		return err
	}
	return m.InvalidateCache()
}
