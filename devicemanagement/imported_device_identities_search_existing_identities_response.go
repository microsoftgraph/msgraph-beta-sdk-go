package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImportedDeviceIdentitiesSearchExistingIdentitiesResponse 
// Deprecated: This class is obsolete. Use searchExistingIdentitiesPostResponse instead.
type ImportedDeviceIdentitiesSearchExistingIdentitiesResponse struct {
    ImportedDeviceIdentitiesSearchExistingIdentitiesPostResponse
}
// NewImportedDeviceIdentitiesSearchExistingIdentitiesResponse instantiates a new ImportedDeviceIdentitiesSearchExistingIdentitiesResponse and sets the default values.
func NewImportedDeviceIdentitiesSearchExistingIdentitiesResponse()(*ImportedDeviceIdentitiesSearchExistingIdentitiesResponse) {
    m := &ImportedDeviceIdentitiesSearchExistingIdentitiesResponse{
        ImportedDeviceIdentitiesSearchExistingIdentitiesPostResponse: *NewImportedDeviceIdentitiesSearchExistingIdentitiesPostResponse(),
    }
    return m
}
// CreateImportedDeviceIdentitiesSearchExistingIdentitiesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedDeviceIdentitiesSearchExistingIdentitiesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImportedDeviceIdentitiesSearchExistingIdentitiesResponse(), nil
}
// ImportedDeviceIdentitiesSearchExistingIdentitiesResponseable 
// Deprecated: This class is obsolete. Use searchExistingIdentitiesPostResponse instead.
type ImportedDeviceIdentitiesSearchExistingIdentitiesResponseable interface {
    ImportedDeviceIdentitiesSearchExistingIdentitiesPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
