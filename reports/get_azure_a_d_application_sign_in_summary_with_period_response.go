package reports

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use GetAzureADApplicationSignInSummaryWithPeriodGetResponseable instead.
type GetAzureADApplicationSignInSummaryWithPeriodResponse struct {
    GetAzureADApplicationSignInSummaryWithPeriodGetResponse
}
// NewGetAzureADApplicationSignInSummaryWithPeriodResponse instantiates a new GetAzureADApplicationSignInSummaryWithPeriodResponse and sets the default values.
func NewGetAzureADApplicationSignInSummaryWithPeriodResponse()(*GetAzureADApplicationSignInSummaryWithPeriodResponse) {
    m := &GetAzureADApplicationSignInSummaryWithPeriodResponse{
        GetAzureADApplicationSignInSummaryWithPeriodGetResponse: *NewGetAzureADApplicationSignInSummaryWithPeriodGetResponse(),
    }
    return m
}
// CreateGetAzureADApplicationSignInSummaryWithPeriodResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGetAzureADApplicationSignInSummaryWithPeriodResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetAzureADApplicationSignInSummaryWithPeriodResponse(), nil
}
// Deprecated: This class is obsolete. Use GetAzureADApplicationSignInSummaryWithPeriodGetResponseable instead.
type GetAzureADApplicationSignInSummaryWithPeriodResponseable interface {
    GetAzureADApplicationSignInSummaryWithPeriodGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
