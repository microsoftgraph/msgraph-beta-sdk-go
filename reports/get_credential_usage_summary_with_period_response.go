package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetCredentialUsageSummaryWithPeriodResponse 
// Deprecated: This class is obsolete. Use getCredentialUsageSummaryWithPeriodGetResponse instead.
type GetCredentialUsageSummaryWithPeriodResponse struct {
    GetCredentialUsageSummaryWithPeriodGetResponse
}
// NewGetCredentialUsageSummaryWithPeriodResponse instantiates a new GetCredentialUsageSummaryWithPeriodResponse and sets the default values.
func NewGetCredentialUsageSummaryWithPeriodResponse()(*GetCredentialUsageSummaryWithPeriodResponse) {
    m := &GetCredentialUsageSummaryWithPeriodResponse{
        GetCredentialUsageSummaryWithPeriodGetResponse: *NewGetCredentialUsageSummaryWithPeriodGetResponse(),
    }
    return m
}
// CreateGetCredentialUsageSummaryWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetCredentialUsageSummaryWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetCredentialUsageSummaryWithPeriodResponse(), nil
}
// GetCredentialUsageSummaryWithPeriodResponseable 
// Deprecated: This class is obsolete. Use getCredentialUsageSummaryWithPeriodGetResponse instead.
type GetCredentialUsageSummaryWithPeriodResponseable interface {
    GetCredentialUsageSummaryWithPeriodGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
