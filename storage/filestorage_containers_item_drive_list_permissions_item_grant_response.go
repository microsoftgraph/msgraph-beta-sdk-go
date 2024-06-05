package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveListPermissionsItemGrantPostResponseable instead.
type FilestorageContainersItemDriveListPermissionsItemGrantResponse struct {
    FilestorageContainersItemDriveListPermissionsItemGrantPostResponse
}
// NewFilestorageContainersItemDriveListPermissionsItemGrantResponse instantiates a new FilestorageContainersItemDriveListPermissionsItemGrantResponse and sets the default values.
func NewFilestorageContainersItemDriveListPermissionsItemGrantResponse()(*FilestorageContainersItemDriveListPermissionsItemGrantResponse) {
    m := &FilestorageContainersItemDriveListPermissionsItemGrantResponse{
        FilestorageContainersItemDriveListPermissionsItemGrantPostResponse: *NewFilestorageContainersItemDriveListPermissionsItemGrantPostResponse(),
    }
    return m
}
// CreateFilestorageContainersItemDriveListPermissionsItemGrantResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageContainersItemDriveListPermissionsItemGrantResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageContainersItemDriveListPermissionsItemGrantResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageContainersItemDriveListPermissionsItemGrantPostResponseable instead.
type FilestorageContainersItemDriveListPermissionsItemGrantResponseable interface {
    FilestorageContainersItemDriveListPermissionsItemGrantPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
