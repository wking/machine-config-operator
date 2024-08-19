// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	configv1 "github.com/openshift/api/config/v1"
	operatorv1 "github.com/openshift/api/operator/v1"
)

// ConsoleCustomizationApplyConfiguration represents an declarative configuration of the ConsoleCustomization type for use
// with apply.
type ConsoleCustomizationApplyConfiguration struct {
	Capabilities         []CapabilityApplyConfiguration                          `json:"capabilities,omitempty"`
	Brand                *operatorv1.Brand                                       `json:"brand,omitempty"`
	DocumentationBaseURL *string                                                 `json:"documentationBaseURL,omitempty"`
	CustomProductName    *string                                                 `json:"customProductName,omitempty"`
	CustomLogoFile       *configv1.ConfigMapFileReference                        `json:"customLogoFile,omitempty"`
	DeveloperCatalog     *DeveloperConsoleCatalogCustomizationApplyConfiguration `json:"developerCatalog,omitempty"`
	ProjectAccess        *ProjectAccessApplyConfiguration                        `json:"projectAccess,omitempty"`
	QuickStarts          *QuickStartsApplyConfiguration                          `json:"quickStarts,omitempty"`
	AddPage              *AddPageApplyConfiguration                              `json:"addPage,omitempty"`
	Perspectives         []PerspectiveApplyConfiguration                         `json:"perspectives,omitempty"`
}

// ConsoleCustomizationApplyConfiguration constructs an declarative configuration of the ConsoleCustomization type for use with
// apply.
func ConsoleCustomization() *ConsoleCustomizationApplyConfiguration {
	return &ConsoleCustomizationApplyConfiguration{}
}

// WithCapabilities adds the given value to the Capabilities field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Capabilities field.
func (b *ConsoleCustomizationApplyConfiguration) WithCapabilities(values ...*CapabilityApplyConfiguration) *ConsoleCustomizationApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithCapabilities")
		}
		b.Capabilities = append(b.Capabilities, *values[i])
	}
	return b
}

// WithBrand sets the Brand field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Brand field is set to the value of the last call.
func (b *ConsoleCustomizationApplyConfiguration) WithBrand(value operatorv1.Brand) *ConsoleCustomizationApplyConfiguration {
	b.Brand = &value
	return b
}

// WithDocumentationBaseURL sets the DocumentationBaseURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DocumentationBaseURL field is set to the value of the last call.
func (b *ConsoleCustomizationApplyConfiguration) WithDocumentationBaseURL(value string) *ConsoleCustomizationApplyConfiguration {
	b.DocumentationBaseURL = &value
	return b
}

// WithCustomProductName sets the CustomProductName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CustomProductName field is set to the value of the last call.
func (b *ConsoleCustomizationApplyConfiguration) WithCustomProductName(value string) *ConsoleCustomizationApplyConfiguration {
	b.CustomProductName = &value
	return b
}

// WithCustomLogoFile sets the CustomLogoFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CustomLogoFile field is set to the value of the last call.
func (b *ConsoleCustomizationApplyConfiguration) WithCustomLogoFile(value configv1.ConfigMapFileReference) *ConsoleCustomizationApplyConfiguration {
	b.CustomLogoFile = &value
	return b
}

// WithDeveloperCatalog sets the DeveloperCatalog field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeveloperCatalog field is set to the value of the last call.
func (b *ConsoleCustomizationApplyConfiguration) WithDeveloperCatalog(value *DeveloperConsoleCatalogCustomizationApplyConfiguration) *ConsoleCustomizationApplyConfiguration {
	b.DeveloperCatalog = value
	return b
}

// WithProjectAccess sets the ProjectAccess field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProjectAccess field is set to the value of the last call.
func (b *ConsoleCustomizationApplyConfiguration) WithProjectAccess(value *ProjectAccessApplyConfiguration) *ConsoleCustomizationApplyConfiguration {
	b.ProjectAccess = value
	return b
}

// WithQuickStarts sets the QuickStarts field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the QuickStarts field is set to the value of the last call.
func (b *ConsoleCustomizationApplyConfiguration) WithQuickStarts(value *QuickStartsApplyConfiguration) *ConsoleCustomizationApplyConfiguration {
	b.QuickStarts = value
	return b
}

// WithAddPage sets the AddPage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AddPage field is set to the value of the last call.
func (b *ConsoleCustomizationApplyConfiguration) WithAddPage(value *AddPageApplyConfiguration) *ConsoleCustomizationApplyConfiguration {
	b.AddPage = value
	return b
}

// WithPerspectives adds the given value to the Perspectives field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Perspectives field.
func (b *ConsoleCustomizationApplyConfiguration) WithPerspectives(values ...*PerspectiveApplyConfiguration) *ConsoleCustomizationApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPerspectives")
		}
		b.Perspectives = append(b.Perspectives, *values[i])
	}
	return b
}
