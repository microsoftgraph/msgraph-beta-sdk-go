package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemInvitePostResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemInviteResponse struct {
    FilestorageDeletedcontainersItemDriveItemsItemInvitePostResponse
}
// NewFilestorageDeletedcontainersItemDriveItemsItemInviteResponse instantiates a new FilestorageDeletedcontainersItemDriveItemsItemInviteResponse and sets the default values.
func NewFilestorageDeletedcontainersItemDriveItemsItemInviteResponse()(*FilestorageDeletedcontainersItemDriveItemsItemInviteResponse) {
    m := &FilestorageDeletedcontainersItemDriveItemsItemInviteResponse{
        FilestorageDeletedcontainersItemDriveItemsItemInvitePostResponse: *NewFilestorageDeletedcontainersItemDriveItemsItemInvitePostResponse(),
    }
    return m
}
// CreateFilestorageDeletedcontainersItemDriveItemsItemInviteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilestorageDeletedcontainersItemDriveItemsItemInviteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilestorageDeletedcontainersItemDriveItemsItemInviteResponse(), nil
}
// Deprecated: This class is obsolete. Use FilestorageDeletedcontainersItemDriveItemsItemInvitePostResponseable instead.
type FilestorageDeletedcontainersItemDriveItemsItemInviteResponseable interface {
    FilestorageDeletedcontainersItemDriveItemsItemInvitePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
