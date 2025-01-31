// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spec

import "io"

// SpecificationSet represents an entire set of specification.
type SpecificationSet interface {

	// Specification returns the Specification with the given name.
	Specification(name string) Specification

	// SpecificationGroup returns the Specifications in the given group name.
	SpecificationGroup(groupName string) []Specification

	// Specifications returns all Specifications.
	Specifications() (specs []Specification)

	// Len returns the number of specifications in the set.
	Len() int

	// Relationships is better
	Relationships() map[string]*Relationship

	// RelationshipsByRestName returns the relationships indexed by rest name.
	RelationshipsByRestName() map[string]*Relationship

	// RelationshipsByResourceName returns the relationships indexed by resource name.
	RelationshipsByResourceName() map[string]*Relationship

	// Configuration returns the specification set Config.
	Configuration() *Config

	// TypeMapping returns the specification set TypeMapping.
	TypeMapping() TypeMapping

	// ValidationMapping returns the specification set ValidationMapping.
	ValidationMapping() ValidationMapping

	// APIInfo returns the specification set APIInfo.
	APIInfo() *APIInfo

	// Groups returns the list of group names.
	Groups() []string
}

// A Specification is the interface representing a Regolithe Specification.
type Specification interface {

	// Read reads and load the given reader containing a specification.
	// If validates is true, validations will be done.
	Read(reader io.Reader, validate bool) error

	// Writes write the current state of the Specification in the given
	// writer.
	Write(writer io.Writer) error

	// Validate validates the specification content.
	Validate() []error

	// ApplyBaseSpecifications applyes the given abstract specification to
	// the specification.
	ApplyBaseSpecifications(specs ...Specification) error

	// Model returns the Specification model.
	Model() *Model

	// Attribute returns the attribute with the given name in the given version.
	Attribute(name string, version string) *Attribute

	// Attributes returns all attributes for the given version.
	Attributes(version string) []*Attribute

	// ExposedAttributes returns only the exposed attributes in the given version.
	ExposedAttributes(version string) []*Attribute

	// OrderingAttributes returns the list of attributes used for ordering.
	OrderingAttributes(version string) []*Attribute

	// AttributeVersions returns all the versions of attributes.
	AttributeVersions() []string

	// LatestAttributesVersion returns the latest version of the attributes.
	LatestAttributesVersion() string

	// Relations returns the Specification relations.
	Relations() []*Relation

	// Relation returns the relation to the given restName.
	Relation(restName string) *Relation

	// Identitier returns the Attribute used as an identifier.
	Identifier() *Attribute
}
