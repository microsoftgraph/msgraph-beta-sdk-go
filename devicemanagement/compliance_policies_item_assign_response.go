package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CompliancePoliciesItemAssignResponse 
// Deprecated: This class is obsolete. Use assignPostResponse instead.
type CompliancePoliciesItemAssignResponse struct {
    CompliancePoliciesItemAssignPostResponse
}
// NewCompliancePoliciesItemAssignResponse instantiates a new CompliancePoliciesItemAssignResponse and sets the default values.
func NewCompliancePoliciesItemAssignResponse()(*CompliancePoliciesItemAssignResponse) {
    m := &CompliancePoliciesItemAssignResponse{
        CompliancePoliciesItemAssignPostResponse: *NewCompliancePoliciesItemAssignPostResponse(),
    }
    return m
}
// CreateCompliancePoliciesItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCompliancePoliciesItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompliancePoliciesItemAssignResponse(), nil
}
// CompliancePoliciesItemAssignResponseable 
// Deprecated: This class is obsolete. Use assignPostResponse instead.
type CompliancePoliciesItemAssignResponseable interface {
    CompliancePoliciesItemAssignPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
