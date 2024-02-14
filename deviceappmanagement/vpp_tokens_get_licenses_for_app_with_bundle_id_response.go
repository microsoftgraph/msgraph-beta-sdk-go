package deviceappmanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type VppTokensGetLicensesForAppWithBundleIdResponse struct {
    VppTokensGetLicensesForAppWithBundleIdGetResponse
}
// NewVppTokensGetLicensesForAppWithBundleIdResponse instantiates a new VppTokensGetLicensesForAppWithBundleIdResponse and sets the default values.
func NewVppTokensGetLicensesForAppWithBundleIdResponse()(*VppTokensGetLicensesForAppWithBundleIdResponse) {
    m := &VppTokensGetLicensesForAppWithBundleIdResponse{
        VppTokensGetLicensesForAppWithBundleIdGetResponse: *NewVppTokensGetLicensesForAppWithBundleIdGetResponse(),
    }
    return m
}
// CreateVppTokensGetLicensesForAppWithBundleIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVppTokensGetLicensesForAppWithBundleIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVppTokensGetLicensesForAppWithBundleIdResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type VppTokensGetLicensesForAppWithBundleIdResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VppTokensGetLicensesForAppWithBundleIdGetResponseable
}
