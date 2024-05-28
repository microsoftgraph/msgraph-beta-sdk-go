package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveRecentGetResponseable instead.
type FilestorageDeletedcontainersItemDriveRecentResponse struct {
    FilestorageDeletedcontainersItemDriveRecentGetResponse
}
// NewFilestorageDeletedcontainersItemDriveRecentResponse instantiates a new FilestorageDeletedcontainersItemDriveRecentResponse and sets the default values.
func NewFilestorageDeletedcontainersItemDriveRecentResponse()(*FilestorageDeletedcontainersItemDriveRecentResponse) {
    m := &FilestorageDeletedcontainersItemDriveRecentResponse{
        FilestorageDeletedcontainersItemDriveRecentGetResponse: *NewFilestorageDeletedcontainersItemDriveRecentGetResponse(),
    }
    return m
}
// CreateFilestorageDeletedcontainersItemDriveRecentResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageDeletedcontainersItemDriveRecentResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageDeletedcontainersItemDriveRecentResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveRecentGetResponseable instead.
type FilestorageDeletedcontainersItemDriveRecentResponseable interface {
    FilestorageDeletedcontainersItemDriveRecentGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
