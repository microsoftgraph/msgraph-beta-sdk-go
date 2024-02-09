package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type AssignmentFiltersItemGetSupportedPropertiesResponse struct {
    AssignmentFiltersItemGetSupportedPropertiesGetResponse
}
// NewAssignmentFiltersItemGetSupportedPropertiesResponse instantiates a new AssignmentFiltersItemGetSupportedPropertiesResponse and sets the default values.
func NewAssignmentFiltersItemGetSupportedPropertiesResponse()(*AssignmentFiltersItemGetSupportedPropertiesResponse) {
    m := &AssignmentFiltersItemGetSupportedPropertiesResponse{
        AssignmentFiltersItemGetSupportedPropertiesGetResponse: *NewAssignmentFiltersItemGetSupportedPropertiesGetResponse(),
    }
    return m
}
// CreateAssignmentFiltersItemGetSupportedPropertiesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAssignmentFiltersItemGetSupportedPropertiesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFiltersItemGetSupportedPropertiesResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type AssignmentFiltersItemGetSupportedPropertiesResponseable interface {
    AssignmentFiltersItemGetSupportedPropertiesGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
