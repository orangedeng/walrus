// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/walrus/pkg/dao/model/resourcedefinition"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/utils/json"
)

// ResourceDefinition is the model entity for the ResourceDefinition schema.
type ResourceDefinition struct {
	config `json:"-"`
	// ID of the ent.
	ID object.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations holds the value of the "annotations" field.
	Annotations map[string]string `json:"annotations,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Type of the resources generated from the resource definition.
	Type string `json:"type,omitempty,cli-table-column"`
	// Generated schema of the resource definition.
	Schema types.Schema `json:"schema,omitempty"`
	// UI schema of the resource definition.
	UiSchema *types.UISchema `json:"uiSchema,omitempty"`
	// Indicate whether the resource definition is builtin, decided when creating.
	Builtin bool `json:"builtin,omitempty,cli-table-column"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResourceDefinitionQuery when eager-loading is set.
	Edges        ResourceDefinitionEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// ResourceDefinitionEdges holds the relations/edges for other nodes in the graph.
type ResourceDefinitionEdges struct {
	// MatchingRules holds the value of the matching_rules edge.
	MatchingRules []*ResourceDefinitionMatchingRule `json:"matching_rules,omitempty"`
	// Resources that use the definition.
	Resources []*Resource `json:"resources,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MatchingRulesOrErr returns the MatchingRules value or an error if the edge
// was not loaded in eager-loading.
func (e ResourceDefinitionEdges) MatchingRulesOrErr() ([]*ResourceDefinitionMatchingRule, error) {
	if e.loadedTypes[0] {
		return e.MatchingRules, nil
	}
	return nil, &NotLoadedError{edge: "matching_rules"}
}

// ResourcesOrErr returns the Resources value or an error if the edge
// was not loaded in eager-loading.
func (e ResourceDefinitionEdges) ResourcesOrErr() ([]*Resource, error) {
	if e.loadedTypes[1] {
		return e.Resources, nil
	}
	return nil, &NotLoadedError{edge: "resources"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ResourceDefinition) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case resourcedefinition.FieldLabels, resourcedefinition.FieldAnnotations, resourcedefinition.FieldSchema, resourcedefinition.FieldUiSchema:
			values[i] = new([]byte)
		case resourcedefinition.FieldID:
			values[i] = new(object.ID)
		case resourcedefinition.FieldBuiltin:
			values[i] = new(sql.NullBool)
		case resourcedefinition.FieldName, resourcedefinition.FieldDescription, resourcedefinition.FieldType:
			values[i] = new(sql.NullString)
		case resourcedefinition.FieldCreateTime, resourcedefinition.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ResourceDefinition fields.
func (rd *ResourceDefinition) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resourcedefinition.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				rd.ID = *value
			}
		case resourcedefinition.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rd.Name = value.String
			}
		case resourcedefinition.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				rd.Description = value.String
			}
		case resourcedefinition.FieldLabels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rd.Labels); err != nil {
					return fmt.Errorf("unmarshal field labels: %w", err)
				}
			}
		case resourcedefinition.FieldAnnotations:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field annotations", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rd.Annotations); err != nil {
					return fmt.Errorf("unmarshal field annotations: %w", err)
				}
			}
		case resourcedefinition.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				rd.CreateTime = new(time.Time)
				*rd.CreateTime = value.Time
			}
		case resourcedefinition.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				rd.UpdateTime = new(time.Time)
				*rd.UpdateTime = value.Time
			}
		case resourcedefinition.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				rd.Type = value.String
			}
		case resourcedefinition.FieldSchema:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field schema", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rd.Schema); err != nil {
					return fmt.Errorf("unmarshal field schema: %w", err)
				}
			}
		case resourcedefinition.FieldUiSchema:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field uiSchema", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &rd.UiSchema); err != nil {
					return fmt.Errorf("unmarshal field uiSchema: %w", err)
				}
			}
		case resourcedefinition.FieldBuiltin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field builtin", values[i])
			} else if value.Valid {
				rd.Builtin = value.Bool
			}
		default:
			rd.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ResourceDefinition.
// This includes values selected through modifiers, order, etc.
func (rd *ResourceDefinition) Value(name string) (ent.Value, error) {
	return rd.selectValues.Get(name)
}

// QueryMatchingRules queries the "matching_rules" edge of the ResourceDefinition entity.
func (rd *ResourceDefinition) QueryMatchingRules() *ResourceDefinitionMatchingRuleQuery {
	return NewResourceDefinitionClient(rd.config).QueryMatchingRules(rd)
}

// QueryResources queries the "resources" edge of the ResourceDefinition entity.
func (rd *ResourceDefinition) QueryResources() *ResourceQuery {
	return NewResourceDefinitionClient(rd.config).QueryResources(rd)
}

// Update returns a builder for updating this ResourceDefinition.
// Note that you need to call ResourceDefinition.Unwrap() before calling this method if this ResourceDefinition
// was returned from a transaction, and the transaction was committed or rolled back.
func (rd *ResourceDefinition) Update() *ResourceDefinitionUpdateOne {
	return NewResourceDefinitionClient(rd.config).UpdateOne(rd)
}

// Unwrap unwraps the ResourceDefinition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rd *ResourceDefinition) Unwrap() *ResourceDefinition {
	_tx, ok := rd.config.driver.(*txDriver)
	if !ok {
		panic("model: ResourceDefinition is not a transactional entity")
	}
	rd.config.driver = _tx.drv
	return rd
}

// String implements the fmt.Stringer.
func (rd *ResourceDefinition) String() string {
	var builder strings.Builder
	builder.WriteString("ResourceDefinition(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rd.ID))
	builder.WriteString("name=")
	builder.WriteString(rd.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(rd.Description)
	builder.WriteString(", ")
	builder.WriteString("labels=")
	builder.WriteString(fmt.Sprintf("%v", rd.Labels))
	builder.WriteString(", ")
	builder.WriteString("annotations=")
	builder.WriteString(fmt.Sprintf("%v", rd.Annotations))
	builder.WriteString(", ")
	if v := rd.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := rd.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(rd.Type)
	builder.WriteString(", ")
	builder.WriteString("schema=")
	builder.WriteString(fmt.Sprintf("%v", rd.Schema))
	builder.WriteString(", ")
	builder.WriteString("uiSchema=")
	builder.WriteString(fmt.Sprintf("%v", rd.UiSchema))
	builder.WriteString(", ")
	builder.WriteString("builtin=")
	builder.WriteString(fmt.Sprintf("%v", rd.Builtin))
	builder.WriteByte(')')
	return builder.String()
}

// ResourceDefinitions is a parsable slice of ResourceDefinition.
type ResourceDefinitions []*ResourceDefinition
