package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveListItemsDeltaGetResponseable instead.
type FilestorageDeletedcontainersItemDriveListItemsDeltaResponse struct {
    FilestorageDeletedcontainersItemDriveListItemsDeltaGetResponse
}
// NewFilestorageDeletedcontainersItemDriveListItemsDeltaResponse instantiates a new FilestorageDeletedcontainersItemDriveListItemsDeltaResponse and sets the default values.
func NewFilestorageDeletedcontainersItemDriveListItemsDeltaResponse()(*FilestorageDeletedcontainersItemDriveListItemsDeltaResponse) {
    m := &FilestorageDeletedcontainersItemDriveListItemsDeltaResponse{
        FilestorageDeletedcontainersItemDriveListItemsDeltaGetResponse: *NewFilestorageDeletedcontainersItemDriveListItemsDeltaGetResponse(),
    }
    return m
}
// CreateFilestorageDeletedcontainersItemDriveListItemsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageDeletedcontainersItemDriveListItemsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageDeletedcontainersItemDriveListItemsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveListItemsDeltaGetResponseable instead.
type FilestorageDeletedcontainersItemDriveListItemsDeltaResponseable interface {
    FilestorageDeletedcontainersItemDriveListItemsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
