// Package v1beta1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1beta1

import (
	externalRef1 "github.com/meshery/schemas/models/core"
	"github.com/meshery/schemas/models/v1alpha2"
	"github.com/meshery/schemas/models/v1alpha3"
)

// Defines values for MesheryPatternCatalogDataClass.
const (
	Community MesheryPatternCatalogDataClass = "community"
	Official  MesheryPatternCatalogDataClass = "official"
	Verified  MesheryPatternCatalogDataClass = "verified"
)

// Defines values for MesheryPatternCatalogDataCompatibility.
const (
	Kubernetes MesheryPatternCatalogDataCompatibility = "kubernetes"
)

// Defines values for MesheryPatternCatalogDataType.
const (
	Deployment        MesheryPatternCatalogDataType = "Deployment"
	Observability     MesheryPatternCatalogDataType = "Observability"
	Resiliency        MesheryPatternCatalogDataType = "Resiliency"
	Scaling           MesheryPatternCatalogDataType = "Scaling"
	Security          MesheryPatternCatalogDataType = "Security"
	TrafficManagement MesheryPatternCatalogDataType = "Traffic-management"
	Troubleshooting   MesheryPatternCatalogDataType = "Troubleshooting"
	Workloads         MesheryPatternCatalogDataType = "Workloads"
)

// Defines values for MesheryPatternPatternFileComponentsModelStatus.
const (
	MesheryPatternPatternFileComponentsModelStatusDuplicate MesheryPatternPatternFileComponentsModelStatus = "duplicate"
	MesheryPatternPatternFileComponentsModelStatusEnabled   MesheryPatternPatternFileComponentsModelStatus = "enabled"
	MesheryPatternPatternFileComponentsModelStatusIgnored   MesheryPatternPatternFileComponentsModelStatus = "ignored"
)

// Defines values for MesheryPatternPatternFileRelationshipsKind.
const (
	Edge         MesheryPatternPatternFileRelationshipsKind = "Edge"
	Hierarchical MesheryPatternPatternFileRelationshipsKind = "Hierarchical"
	Sibling      MesheryPatternPatternFileRelationshipsKind = "Sibling"
)

// Defines values for MesheryPatternPatternFileRelationshipsModelStatus.
const (
	MesheryPatternPatternFileRelationshipsModelStatusDuplicate MesheryPatternPatternFileRelationshipsModelStatus = "duplicate"
	MesheryPatternPatternFileRelationshipsModelStatusEnabled   MesheryPatternPatternFileRelationshipsModelStatus = "enabled"
	MesheryPatternPatternFileRelationshipsModelStatusIgnored   MesheryPatternPatternFileRelationshipsModelStatus = "ignored"
)

// Defines values for MesheryPatternPatternFileRelationshipsSelectorsAllowFromModelStatus.
const (
	MesheryPatternPatternFileRelationshipsSelectorsAllowFromModelStatusDuplicate MesheryPatternPatternFileRelationshipsSelectorsAllowFromModelStatus = "duplicate"
	MesheryPatternPatternFileRelationshipsSelectorsAllowFromModelStatusEnabled   MesheryPatternPatternFileRelationshipsSelectorsAllowFromModelStatus = "enabled"
	MesheryPatternPatternFileRelationshipsSelectorsAllowFromModelStatusIgnored   MesheryPatternPatternFileRelationshipsSelectorsAllowFromModelStatus = "ignored"
)

// Defines values for MesheryPatternPatternFileRelationshipsSelectorsAllowFromPatchPatchStrategy.
const (
	MesheryPatternPatternFileRelationshipsSelectorsAllowFromPatchPatchStrategyReplace MesheryPatternPatternFileRelationshipsSelectorsAllowFromPatchPatchStrategy = "replace"
)

// Defines values for MesheryPatternPatternFileRelationshipsSelectorsAllowToModelStatus.
const (
	MesheryPatternPatternFileRelationshipsSelectorsAllowToModelStatusDuplicate MesheryPatternPatternFileRelationshipsSelectorsAllowToModelStatus = "duplicate"
	MesheryPatternPatternFileRelationshipsSelectorsAllowToModelStatusEnabled   MesheryPatternPatternFileRelationshipsSelectorsAllowToModelStatus = "enabled"
	MesheryPatternPatternFileRelationshipsSelectorsAllowToModelStatusIgnored   MesheryPatternPatternFileRelationshipsSelectorsAllowToModelStatus = "ignored"
)

// Defines values for MesheryPatternPatternFileRelationshipsSelectorsAllowToPatchPatchStrategy.
const (
	MesheryPatternPatternFileRelationshipsSelectorsAllowToPatchPatchStrategyReplace MesheryPatternPatternFileRelationshipsSelectorsAllowToPatchPatchStrategy = "replace"
)

// Defines values for MesheryPatternPatternFileRelationshipsSelectorsDenyFromModelStatus.
const (
	MesheryPatternPatternFileRelationshipsSelectorsDenyFromModelStatusDuplicate MesheryPatternPatternFileRelationshipsSelectorsDenyFromModelStatus = "duplicate"
	MesheryPatternPatternFileRelationshipsSelectorsDenyFromModelStatusEnabled   MesheryPatternPatternFileRelationshipsSelectorsDenyFromModelStatus = "enabled"
	MesheryPatternPatternFileRelationshipsSelectorsDenyFromModelStatusIgnored   MesheryPatternPatternFileRelationshipsSelectorsDenyFromModelStatus = "ignored"
)

// Defines values for MesheryPatternPatternFileRelationshipsSelectorsDenyFromPatchPatchStrategy.
const (
	MesheryPatternPatternFileRelationshipsSelectorsDenyFromPatchPatchStrategyReplace MesheryPatternPatternFileRelationshipsSelectorsDenyFromPatchPatchStrategy = "replace"
)

// Defines values for MesheryPatternPatternFileRelationshipsSelectorsDenyToModelStatus.
const (
	MesheryPatternPatternFileRelationshipsSelectorsDenyToModelStatusDuplicate MesheryPatternPatternFileRelationshipsSelectorsDenyToModelStatus = "duplicate"
	MesheryPatternPatternFileRelationshipsSelectorsDenyToModelStatusEnabled   MesheryPatternPatternFileRelationshipsSelectorsDenyToModelStatus = "enabled"
	MesheryPatternPatternFileRelationshipsSelectorsDenyToModelStatusIgnored   MesheryPatternPatternFileRelationshipsSelectorsDenyToModelStatus = "ignored"
)

// Defines values for MesheryPatternPatternFileRelationshipsSelectorsDenyToPatchPatchStrategy.
const (
	MesheryPatternPatternFileRelationshipsSelectorsDenyToPatchPatchStrategyReplace MesheryPatternPatternFileRelationshipsSelectorsDenyToPatchPatchStrategy = "replace"
)

// DeletePatternModel defines model for deletePatternModel.
type DeletePatternModel struct {
	ID   externalRef1.Id   `db:"id" json:"id"`
	Name externalRef1.Text `json:"name,omitempty"`
}

type PatternFile struct {
	// Components List of component declarations
	Components []Component `json:"components" yaml:"components"`

	// Name Name of the design; a descriptive, but concise title for the design document.
	Name string `json:"name" yaml:"name"`

	// Preferences Design-level preferences
	Preferences *struct {
		// Layers List of available layers
		Layers []string `json:"layers"`
	} `json:"preferences,omitempty"`

	// Relationships List of relationships between components
	Relationships []v1alpha3.Relationship `json:"relationships" yaml:"relationships"`

	// SchemaVersion Specifies the version of the schema to which the design conforms.
	SchemaVersion string `json:"schemaVersion" yaml:"schemaVersion"`

	// Version Revision of the design as expressed by an auto-incremented, SemVer-compliant version number. May be manually set by a user or third-party system, but will always be required to be of version number higher than the previously defined version number.
	Version string `json:"version" yaml:"version"`
}

// MesheryPattern defines model for mesheryPattern.
type MesheryPattern struct {
	CatalogData *v1alpha2.CatalogData  `json:"catalog_data,omitempty"`
	CreatedAt   externalRef1.Time      `json:"created_at,omitempty"`
	UserID      externalRef1.Id        `db:"user_id" json:"user_id"`
	Location    externalRef1.MapObject `json:"location,omitempty"`
	Name        externalRef1.Text      `json:"name,omitempty"`

	// PatternFile Designs are your primary tool for collaborative authorship of your infrastructure, workflow, and processes.
	PatternFile *PatternFile      `json:"pattern_file,omitempty"`
	UpdatedAt   externalRef1.Time `json:"updated_at,omitempty"`
	ID          externalRef1.Id   `db:"id" json:"id"`
	Visibility  externalRef1.Text `json:"visibility,omitempty"`
}

// MesheryPatternCatalogDataClass Published content is classifed by its support level. Content classes help you understand the origin and expected support level for each piece of content. It is important to note that the level of support may vary within each class, and you should exercise discretion when using community-contributed content. Content produced and fully supported by Meshery maintainers. This represents the highest level of support and is considered the most reliable. Content produced by partners and verified by Meshery maintainers. While not directly maintained by Meshery, it has undergone a verification process to ensure quality and compatibility. Content produced and supported by the respective project or organization responsible for the specific technology. This class offers a level of support from the project maintainers themselves. Content produced and shared by Meshery users. This includes a wide range of content, such as performance profiles, test results, filters, patterns, and applications. Community content may have varying levels of support and reliability.
type MesheryPatternCatalogDataClass string

// MesheryPatternCatalogDataCompatibility defines model for MesheryPattern.CatalogData.Compatibility.
type MesheryPatternCatalogDataCompatibility string

// MesheryPatternCatalogDataType Categorization of the type of design or operational flow depicted in this design.
type MesheryPatternCatalogDataType string

// MesheryPatternPatternFileComponentsFormat Format specifies the format used in the `component.schema` field. JSON is the default.
type MesheryPatternPatternFileComponentsFormat string

// MesheryPatternPatternFileComponentsMetadataShape Shape of the component used for UI representation.
type MesheryPatternPatternFileComponentsMetadataShape string

// MesheryPattern_PatternFile_Components_Metadata Metadata contains additional information associated with the component.
type MesheryPattern_PatternFile_Components_Metadata struct {
	// Capabilities Meshery manages components in accordance with their specific capabilities. This field explicitly identifies those capabilities largely by what actions a given component supports; e.g. metric-scrape, sub-interface, and so on. This field is extensible. ComponentDefinitions may define a broad array of capabilities, which are in-turn dynamically interpretted by Meshery for full lifecycle management.
	Capabilities *map[string]interface{} `json:"capabilities" yaml:"capabilities"`

	// Genealogy Genealogy represents the various representational states of the component.
	Genealogy *string `json:"genealogy" yaml:"genealogy"`

	// IsAnnotation Identifies whether the component is semantically meaningful or not; identifies whether the component should be treated as deployable entity or is for purposes of logical representation.
	IsAnnotation *bool `json:"isAnnotation" yaml:"isAnnotation"`

	// PrimaryColor Primary color of the component used for UI representation.
	PrimaryColor string `json:"primaryColor" yaml:"primaryColor"`

	// Published 'published' controls whether the component should be registered in Meshery Registry. When the same 'published' property in Models, is set to 'false', the Model property takes precedence with all Entities in the Model not being registered.
	Published *bool `json:"published" yaml:"published"`

	// SecondaryColor Secondary color of the component used for UI representation.
	SecondaryColor *string `json:"secondaryColor" yaml:"secondaryColor"`

	// Shape Shape of the component used for UI representation.
	Shape MesheryPatternPatternFileComponentsMetadataShape `json:"shape" yaml:"shape"`

	// SvgColor Colored SVG of the component used for UI representation on light background.
	SvgColor string `json:"svgColor" yaml:"svgColor"`

	// SvgComplete Complete SVG of the component used for UI representation, often inclusive of background.
	SvgComplete *string `json:"svgComplete" yaml:"svgComplete"`

	// SvgWhite White SVG of the component used for UI representation on dark background.
	SvgWhite             string                 `json:"svgWhite" yaml:"svgWhite"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// MesheryPatternPatternFileComponentsModelStatus Status of model, including:
// - duplicate: this component is a duplicate of another. The component that is to be the canonical reference and that is duplicated by other components should not be assigned the 'duplicate' status.
// - maintenance: model is unavailable for a period of time.
// - enabled: model is available for use for all users of this Meshery Server.
// - ignored: model is unavailable for use for all users of this Meshery Server.
type MesheryPatternPatternFileComponentsModelStatus string

// MesheryPatternPatternFileRelationshipsKind Kind of the Relationship. Learn more about relationships - https://docs.meshery.io/concepts/logical/relationships.
type MesheryPatternPatternFileRelationshipsKind string

// MesheryPatternPatternFileRelationshipsModelStatus Status of model, including:
// - duplicate: this component is a duplicate of another. The component that is to be the canonical reference and that is duplicated by other components should not be assigned the 'duplicate' status.
// - maintenance: model is unavailable for a period of time.
// - enabled: model is available for use for all users of this Meshery Server.
// - ignored: model is unavailable for use for all users of this Meshery Server.
type MesheryPatternPatternFileRelationshipsModelStatus string

// MesheryPatternPatternFileRelationshipsSelectorsAllowFromModelStatus Status of model, including:
// - duplicate: this component is a duplicate of another. The component that is to be the canonical reference and that is duplicated by other components should not be assigned the 'duplicate' status.
// - maintenance: model is unavailable for a period of time.
// - enabled: model is available for use for all users of this Meshery Server.
// - ignored: model is unavailable for use for all users of this Meshery Server.
type MesheryPatternPatternFileRelationshipsSelectorsAllowFromModelStatus string

// MesheryPatternPatternFileRelationshipsSelectorsAllowFromPatchPatchStrategy defines model for MesheryPattern.PatternFile.Relationships.Selectors.Allow.From.Patch.PatchStrategy.
type MesheryPatternPatternFileRelationshipsSelectorsAllowFromPatchPatchStrategy string

// MesheryPatternPatternFileRelationshipsSelectorsAllowToModelStatus Status of model, including:
// - duplicate: this component is a duplicate of another. The component that is to be the canonical reference and that is duplicated by other components should not be assigned the 'duplicate' status.
// - maintenance: model is unavailable for a period of time.
// - enabled: model is available for use for all users of this Meshery Server.
// - ignored: model is unavailable for use for all users of this Meshery Server.
type MesheryPatternPatternFileRelationshipsSelectorsAllowToModelStatus string

// MesheryPatternPatternFileRelationshipsSelectorsAllowToPatchPatchStrategy defines model for MesheryPattern.PatternFile.Relationships.Selectors.Allow.To.Patch.PatchStrategy.
type MesheryPatternPatternFileRelationshipsSelectorsAllowToPatchPatchStrategy string

// MesheryPatternPatternFileRelationshipsSelectorsDenyFromModelStatus Status of model, including:
// - duplicate: this component is a duplicate of another. The component that is to be the canonical reference and that is duplicated by other components should not be assigned the 'duplicate' status.
// - maintenance: model is unavailable for a period of time.
// - enabled: model is available for use for all users of this Meshery Server.
// - ignored: model is unavailable for use for all users of this Meshery Server.
type MesheryPatternPatternFileRelationshipsSelectorsDenyFromModelStatus string

// MesheryPatternPatternFileRelationshipsSelectorsDenyFromPatchPatchStrategy defines model for MesheryPattern.PatternFile.Relationships.Selectors.Deny.From.Patch.PatchStrategy.
type MesheryPatternPatternFileRelationshipsSelectorsDenyFromPatchPatchStrategy string

// MesheryPatternPatternFileRelationshipsSelectorsDenyToModelStatus Status of model, including:
// - duplicate: this component is a duplicate of another. The component that is to be the canonical reference and that is duplicated by other components should not be assigned the 'duplicate' status.
// - maintenance: model is unavailable for a period of time.
// - enabled: model is available for use for all users of this Meshery Server.
// - ignored: model is unavailable for use for all users of this Meshery Server.
type MesheryPatternPatternFileRelationshipsSelectorsDenyToModelStatus string

// MesheryPatternPatternFileRelationshipsSelectorsDenyToPatchPatchStrategy defines model for MesheryPattern.PatternFile.Relationships.Selectors.Deny.To.Patch.PatchStrategy.
type MesheryPatternPatternFileRelationshipsSelectorsDenyToPatchPatchStrategy string

// MesheryPatternPage defines model for mesheryPatternPage.
type MesheryPatternPage struct {
	Page       int               `json:"page,omitempty"`
	PageSize   int               `json:"page_size,omitempty"`
	Patterns   *[]MesheryPattern `json:"patterns,omitempty"`
	ResultType string            `json:"resultType,omitempty"`
	TotalCount int               `json:"total_count,omitempty"`
}

// MesheryPatternResource defines model for mesheryPatternResource.
type MesheryPatternResource struct {
	CreatedAt externalRef1.Time `json:"created_at,omitempty"`
	Deleted   *bool             `json:"deleted,omitempty"`
	ID        externalRef1.Id   `db:"id" json:"id"`
	Name      externalRef1.Text `json:"name,omitempty"`
	Namepace  externalRef1.Text `json:"namepace,omitempty"`
	OamType   externalRef1.Text `json:"oam_type,omitempty"`
	Type      externalRef1.Text `json:"type,omitempty"`
	UpdatedAt externalRef1.Time `json:"updated_at,omitempty"`
	UserID    externalRef1.Id   `db:"user_id" json:"user_id"`
}

// MesheryPatternResourcePage defines model for mesheryPatternResourcePage.
type MesheryPatternResourcePage struct {
	Page       int                       `json:"page,omitempty"`
	PageSize   int                       `json:"page_size,omitempty"`
	Resources  *[]MesheryPatternResource `json:"resources,omitempty"`
	ResultType string                    `json:"resultType,omitempty"`
	TotalCount int                       `json:"total_count,omitempty"`
}

// DesignShare defines model for designShare.
type DesignShare struct {
	ContentType string              `json:"content_type"`
	Emails      externalRef1.Emails `json:"emails"`
	ID          externalRef1.Id     `db:"id" json:"id"`
	Share       bool                `json:"share"`
}

// MesheryPatternDeleteRequestBody defines model for mesheryPatternDeleteRequestBody.
type MesheryPatternDeleteRequestBody struct {
	Patterns *[]DeletePatternModel `json:"patterns,omitempty"`
}

// MesheryPatternRequestBody defines model for mesheryPatternRequestBody.
type MesheryPatternRequestBody struct {
	Path        externalRef1.Text     `json:"path,omitempty"`
	PatternData *MesheryPattern       `json:"pattern_data,omitempty"`
	Save        *bool                 `json:"save,omitempty"`
	Url         externalRef1.Endpoint `json:"url,omitempty"`
}

// Getter for additional properties for MesheryPattern_PatternFile_Components_Metadata. Returns the specified
// element and whether it was found
func (a MesheryPattern_PatternFile_Components_Metadata) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}
