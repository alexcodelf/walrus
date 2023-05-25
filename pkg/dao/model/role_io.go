// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "ent". DO NOT EDIT.

package model

import (
	"time"

	"github.com/seal-io/seal/pkg/dao/types"
)

// RoleQueryInput is the input for the Role query.
type RoleQueryInput struct {
	// It is also the name of the role.
	ID string `uri:"id,omitempty" json:"id,omitempty"`
}

// Model converts the RoleQueryInput to Role.
func (in RoleQueryInput) Model() *Role {
	return &Role{
		ID: in.ID,
	}
}

// RoleCreateInput is the input for the Role creation.
type RoleCreateInput struct {
	// The kind of the role.
	Kind string `json:"kind,omitempty"`
	// The detail of the role.
	Description string `json:"description,omitempty"`
	// The policy list of the role.
	Policies types.RolePolicies `json:"policies,omitempty"`
	// Indicate whether the role is session level, decide when creating.
	Session bool `json:"session,omitempty"`
	// Indicate whether the role is builtin, decide when creating.
	Builtin bool `json:"builtin,omitempty"`
}

// Model converts the RoleCreateInput to Role.
func (in RoleCreateInput) Model() *Role {
	var entity = &Role{
		Kind:        in.Kind,
		Description: in.Description,
		Policies:    in.Policies,
		Session:     in.Session,
		Builtin:     in.Builtin,
	}
	return entity
}

// RoleUpdateInput is the input for the Role modification.
type RoleUpdateInput struct {
	// It is also the name of the role.
	ID string `uri:"id" json:"-"`
	// The detail of the role.
	Description string `json:"description,omitempty"`
	// The policy list of the role.
	Policies types.RolePolicies `json:"policies,omitempty"`
}

// Model converts the RoleUpdateInput to Role.
func (in RoleUpdateInput) Model() *Role {
	var entity = &Role{
		ID:          in.ID,
		Description: in.Description,
		Policies:    in.Policies,
	}
	return entity
}

// RoleOutput is the output for the Role.
type RoleOutput struct {
	// It is also the name of the role.
	ID string `json:"id,omitempty"`
	// Describe creation time.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Describe modification time.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// The kind of the role.
	Kind string `json:"kind,omitempty"`
	// The detail of the role.
	Description string `json:"description,omitempty"`
	// The policy list of the role.
	Policies types.RolePolicies `json:"policies,omitempty"`
	// Indicate whether the role is session level, decide when creating.
	Session bool `json:"session,omitempty"`
	// Indicate whether the role is builtin, decide when creating.
	Builtin bool `json:"builtin,omitempty"`
	// Subjects holds the value of the subjects edge.
	Subjects []*SubjectRoleRelationshipOutput `json:"subjects,omitempty"`
}

// ExposeRole converts the Role to RoleOutput.
func ExposeRole(in *Role) *RoleOutput {
	if in == nil {
		return nil
	}
	var entity = &RoleOutput{
		ID:          in.ID,
		CreateTime:  in.CreateTime,
		UpdateTime:  in.UpdateTime,
		Kind:        in.Kind,
		Description: in.Description,
		Policies:    in.Policies,
		Session:     in.Session,
		Builtin:     in.Builtin,
		Subjects:    ExposeSubjectRoleRelationships(in.Edges.Subjects),
	}
	return entity
}

// ExposeRoles converts the Role slice to RoleOutput pointer slice.
func ExposeRoles(in []*Role) []*RoleOutput {
	var out = make([]*RoleOutput, 0, len(in))
	for i := 0; i < len(in); i++ {
		var o = ExposeRole(in[i])
		if o == nil {
			continue
		}
		out = append(out, o)
	}
	if len(out) == 0 {
		return nil
	}
	return out
}
