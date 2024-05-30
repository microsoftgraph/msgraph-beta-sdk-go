package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use CompliancepoliciesItemAssignPostResponseable instead.
type CompliancepoliciesItemAssignResponse struct {
    CompliancepoliciesItemAssignPostResponse
}
// NewCompliancepoliciesItemAssignResponse instantiates a new CompliancepoliciesItemAssignResponse and sets the default values.
func NewCompliancepoliciesItemAssignResponse()(*CompliancepoliciesItemAssignResponse) {
    m := &CompliancepoliciesItemAssignResponse{
        CompliancepoliciesItemAssignPostResponse: *NewCompliancepoliciesItemAssignPostResponse(),
    }
    return m
}
// CreateCompliancepoliciesItemAssignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompliancepoliciesItemAssignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCompliancepoliciesItemAssignResponse(), nil
}
// Deprecated: This class is obsolete. Use CompliancepoliciesItemAssignPostResponseable instead.
type CompliancepoliciesItemAssignResponseable interface {
    CompliancepoliciesItemAssignPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
