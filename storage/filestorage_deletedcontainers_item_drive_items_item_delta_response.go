package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemDeltaGetResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemDeltaResponse struct {
    FilestorageDeletedcontainersItemDriveItemsItemDeltaGetResponse
}
// NewFilestorageDeletedcontainersItemDriveItemsItemDeltaResponse instantiates a new FilestorageDeletedcontainersItemDriveItemsItemDeltaResponse and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemDeltaResponse()(*FilestorageDeletedcontainersItemDriveItemsItemDeltaResponse) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemDeltaResponse{
        FilestorageDeletedcontainersItemDriveItemsItemDeltaGetResponse: *NewFilestorageDeletedcontainersItemDriveItemsItemDeltaGetResponse(),
    }
    return m
}
// CreateFilestorageDeletedcontainersItemDriveItemsItemDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageDeletedcontainersItemDriveItemsItemDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemDeltaGetResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemDeltaResponseable interface {
    FilestorageDeletedcontainersItemDriveItemsItemDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
