package spec

import (
	"fmt"
)

// AttributeFormat represent the allowed format for an attribute.
type AttributeFormat string

// Various values for an AttributeFormat.
const (
	AttributeFormatFree  AttributeFormat = "free"
	AttributeFormatEmail AttributeFormat = "email"
	AttributeFormatPhone AttributeFormat = "phone"
	AttributeFormatIPv4  AttributeFormat = "ipv4"
	AttributeFormatIPv6  AttributeFormat = "ipv6"
	AttributeFormatCIDR  AttributeFormat = "cidr"
	AttributeFormatDate  AttributeFormat = "date"
)

// AttributeUniqueScope represent the unique scope for an attribute.
type AttributeUniqueScope string

// Various values for an AttributeUniqueScope.
const (
	AttributeUniqueScopeLocal  AttributeUniqueScope = "local"
	AttributeUniqueScopeGlobal AttributeUniqueScope = "global"
)

// AttributeType represents the various type for an attribute.
type AttributeType string

// Various values for AttributeType.
const (
	AttributeTypeString AttributeType = "string"
	AttributeTypeInt    AttributeType = "integer"
	AttributeTypeFloat  AttributeType = "float"
	AttributeTypeBool   AttributeType = "boolean"
	AttributeTypeEnum   AttributeType = "enum"
	AttributeTypeList   AttributeType = "list"
	AttributeTypeObject AttributeType = "object"
	AttributeTypeTime   AttributeType = "time"
	AttributeTypeExt    AttributeType = "external"
)

// AttributeNameConverterFunc is the type of a attribute name conveter.
type AttributeNameConverterFunc func(name string) string

// AttributeTypeConverterFunc is the type of a attribute type conveter.
type AttributeTypeConverterFunc func(typ AttributeType, subtype string) (converted string, provider string)

// An Attribute represents a monolithe specification attribute.
type Attribute struct {

	// NOTE: Order of attributes matters!
	// The YAML will be dumped respecting this order.

	Name         string        `yaml:"name,omitempty"           json:"name,omitempty"`
	Description  string        `yaml:"description,omitempty"    json:"description,omitempty"`
	Type         AttributeType `yaml:"type,omitempty"           json:"type,omitempty"`
	Exposed      bool          `yaml:"exposed,omitempty"        json:"exposed,omitempty"`
	SubType      string        `yaml:"subtype,omitempty"        json:"subtype,omitempty"`
	Stored       bool          `yaml:"stored,omitempty"         json:"stored,omitempty"`
	Required     bool          `yaml:"required,omitempty"       json:"required,omitempty"`
	ReadOnly     bool          `yaml:"read_only,omitempty"      json:"read_only,omitempty"`
	CreationOnly bool          `yaml:"creation_only,omitempty"  json:"creation_only,omitempty"`

	AllowedChars   string          `yaml:"allowed_chars,omitempty"      json:"allowed_chars,omitempty"`
	AllowedChoices []string        `yaml:"allowed_choices,omitempty"    json:"allowed_choices,omitempty"`
	Autogenerated  bool            `yaml:"autogenerated,omitempty"      json:"autogenerated,omitempty"`
	DefaultOrder   bool            `yaml:"default_order,omitempty"      json:"default_order,omitempty"`
	DefaultValue   interface{}     `yaml:"default_value,omitempty"      json:"default_value,omitempty"`
	Deprecated     bool            `yaml:"deprecated,omitempty"         json:"deprecated,omitempty"`
	ExampleValue   interface{}     `yaml:"example_value,omitempty"      json:"example_value,omitempty"`
	Filterable     bool            `yaml:"filterable,omitempty"         json:"filterable,omitempty"`
	ForeignKey     bool            `yaml:"foreign_key,omitempty"        json:"foreign_key,omitempty"`
	Format         AttributeFormat `yaml:"format,omitempty"             json:"format,omitempty"`
	Getter         bool            `yaml:"getter,omitempty"             json:"getter,omitempty"`
	Setter         bool            `yaml:"setter,omitempty"             json:"setter,omitempty"`
	Identifier     bool            `yaml:"identifier,omitempty"         json:"identifier,omitempty"`
	Index          bool            `yaml:"index,omitempty"              json:"index,omitempty"`
	MaxLength      uint16          `yaml:"max_length,omitempty"         json:"max_length,omitempty"`
	MinLength      uint16          `yaml:"min_length,omitempty"         json:"min_length,omitempty"`
	MaxValue       float64         `yaml:"max_value,omitempty"          json:"max_value,omitempty"`
	MinValue       float64         `yaml:"min_value,omitempty"          json:"min_value,omitempty"`
	Orderable      bool            `yaml:"orderable,omitempty"          json:"orderable,omitempty"`
	PrimaryKey     bool            `yaml:"primary_key,omitempty"        json:"primary_key,omitempty"`
	Secret         bool            `yaml:"secret,omitempty"             json:"secret,omitempty"`
	Transient      bool            `yaml:"transient,omitempty"          json:"transient,omitempty"`

	ConvertedName string `yaml:"-" json:"-"`
	ConvertedType string `yaml:"-" json:"-"`
	TypeProvider  string `yaml:"-" json:"-"`
	Initializer   string `yaml:"-" json:"-"`

	linkedSpecification *Specification
}

// Validate validates the attribute definition.
func (a *Attribute) Validate() []error {

	var errs []error

	if a.Required && a.DefaultValue == nil && a.ExampleValue == nil {
		errs = append(errs, fmt.Errorf("%s.spec: '%s' is required but has no defaultValue or example value", a.linkedSpecification.Model.RestName, a.Name))
	}

	if a.Description != "" && a.Description[len(a.Description)-1] != '.' && a.linkedSpecification != nil && a.linkedSpecification.Model != nil {
		errs = append(errs, fmt.Errorf("%s.spec: description of attribute '%s' must end with a period", a.linkedSpecification.Model.RestName, a.Name))
	}

	return errs
}
