package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageDeletedContainersItemDriveItemsItemInvitePostResponseable instead.
type FileStorageDeletedContainersItemDriveItemsItemInviteResponse struct {
    FileStorageDeletedContainersItemDriveItemsItemInvitePostResponse
}
// NewFileStorageDeletedContainersItemDriveItemsItemInviteResponse instantiates a new FileStorageDeletedContainersItemDriveItemsItemInviteResponse and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemInviteResponse()(*FileStorageDeletedContainersItemDriveItemsItemInviteResponse) {
    m := &FileStorageDeletedContainersItemDriveItemsItemInviteResponse{
        FileStorageDeletedContainersItemDriveItemsItemInvitePostResponse: *NewFileStorageDeletedContainersItemDriveItemsItemInvitePostResponse(),
    }
    return m
}
// CreateFileStorageDeletedContainersItemDriveItemsItemInviteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageDeletedContainersItemDriveItemsItemInviteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageDeletedContainersItemDriveItemsItemInviteResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageDeletedContainersItemDriveItemsItemInvitePostResponseable instead.
type FileStorageDeletedContainersItemDriveItemsItemInviteResponseable interface {
    FileStorageDeletedContainersItemDriveItemsItemInvitePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
