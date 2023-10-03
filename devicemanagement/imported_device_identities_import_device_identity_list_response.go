package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImportedDeviceIdentitiesImportDeviceIdentityListResponse 
// Deprecated: This class is obsolete. Use importDeviceIdentityListPostResponse instead.
type ImportedDeviceIdentitiesImportDeviceIdentityListResponse struct {
    ImportedDeviceIdentitiesImportDeviceIdentityListPostResponse
}
// NewImportedDeviceIdentitiesImportDeviceIdentityListResponse instantiates a new ImportedDeviceIdentitiesImportDeviceIdentityListResponse and sets the default values.
func NewImportedDeviceIdentitiesImportDeviceIdentityListResponse()(*ImportedDeviceIdentitiesImportDeviceIdentityListResponse) {
    m := &ImportedDeviceIdentitiesImportDeviceIdentityListResponse{
        ImportedDeviceIdentitiesImportDeviceIdentityListPostResponse: *NewImportedDeviceIdentitiesImportDeviceIdentityListPostResponse(),
    }
    return m
}
// CreateImportedDeviceIdentitiesImportDeviceIdentityListResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedDeviceIdentitiesImportDeviceIdentityListResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImportedDeviceIdentitiesImportDeviceIdentityListResponse(), nil
}
// ImportedDeviceIdentitiesImportDeviceIdentityListResponseable 
// Deprecated: This class is obsolete. Use importDeviceIdentityListPostResponse instead.
type ImportedDeviceIdentitiesImportDeviceIdentityListResponseable interface {
    ImportedDeviceIdentitiesImportDeviceIdentityListPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
