// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"time"

	"github.com/seal-io/seal/pkg/dao/types"
)

// ProjectQueryInput is the input for the Project query.
type ProjectQueryInput struct {
	// ID holds the value of the "id" field.
	ID types.ID `uri:"id,omitempty" json:"id,omitempty"`
}

// Model converts the ProjectQueryInput to Project.
func (in ProjectQueryInput) Model() *Project {
	return &Project{
		ID: in.ID,
	}
}

// ProjectCreateInput is the input for the Project creation.
type ProjectCreateInput struct {
	// Name of the resource.
	Name string `json:"name"`
	// Description of the resource.
	Description string `json:"description,omitempty"`
	// Labels of the resource.
	Labels map[string]string `json:"labels,omitempty"`
}

// Model converts the ProjectCreateInput to Project.
func (in ProjectCreateInput) Model() *Project {
	var entity = &Project{
		Name:        in.Name,
		Description: in.Description,
		Labels:      in.Labels,
	}
	return entity
}

// ProjectUpdateInput is the input for the Project modification.
type ProjectUpdateInput struct {
	// ID holds the value of the "id" field.
	ID types.ID `uri:"id" json:"-"`
	// Name of the resource.
	Name string `json:"name,omitempty"`
	// Description of the resource.
	Description string `json:"description,omitempty"`
	// Labels of the resource.
	Labels map[string]string `json:"labels,omitempty"`
}

// Model converts the ProjectUpdateInput to Project.
func (in ProjectUpdateInput) Model() *Project {
	var entity = &Project{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		Labels:      in.Labels,
	}
	return entity
}

// ProjectOutput is the output for the Project.
type ProjectOutput struct {
	// ID holds the value of the "id" field.
	ID types.ID `json:"id,omitempty"`
	// Name of the resource.
	Name string `json:"name,omitempty"`
	// Description of the resource.
	Description string `json:"description,omitempty"`
	// Labels of the resource.
	Labels map[string]string `json:"labels,omitempty"`
	// Describe creation time.
	CreateTime *time.Time `json:"createTime,omitempty"`
	// Describe modification time.
	UpdateTime *time.Time `json:"updateTime,omitempty"`
	// Applications that belong to the project.
	Applications []*ApplicationOutput `json:"applications,omitempty"`
}

// ExposeProject converts the Project to ProjectOutput.
func ExposeProject(in *Project) *ProjectOutput {
	if in == nil {
		return nil
	}
	var entity = &ProjectOutput{
		ID:           in.ID,
		Name:         in.Name,
		Description:  in.Description,
		Labels:       in.Labels,
		CreateTime:   in.CreateTime,
		UpdateTime:   in.UpdateTime,
		Applications: ExposeApplications(in.Edges.Applications),
	}
	return entity
}

// ExposeProjects converts the Project slice to ProjectOutput pointer slice.
func ExposeProjects(in []*Project) []*ProjectOutput {
	var out = make([]*ProjectOutput, 0, len(in))
	for i := 0; i < len(in); i++ {
		var o = ExposeProject(in[i])
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
