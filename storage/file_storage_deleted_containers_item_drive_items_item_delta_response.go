package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageDeletedContainersItemDriveItemsItemDeltaGetResponseable instead.
type FileStorageDeletedContainersItemDriveItemsItemDeltaResponse struct {
    FileStorageDeletedContainersItemDriveItemsItemDeltaGetResponse
}
// NewFileStorageDeletedContainersItemDriveItemsItemDeltaResponse instantiates a new FileStorageDeletedContainersItemDriveItemsItemDeltaResponse and sets the default values.
func NewFileStorageDeletedContainersItemDriveItemsItemDeltaResponse()(*FileStorageDeletedContainersItemDriveItemsItemDeltaResponse) {
    m := &FileStorageDeletedContainersItemDriveItemsItemDeltaResponse{
        FileStorageDeletedContainersItemDriveItemsItemDeltaGetResponse: *NewFileStorageDeletedContainersItemDriveItemsItemDeltaGetResponse(),
    }
    return m
}
// CreateFileStorageDeletedContainersItemDriveItemsItemDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageDeletedContainersItemDriveItemsItemDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageDeletedContainersItemDriveItemsItemDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageDeletedContainersItemDriveItemsItemDeltaGetResponseable instead.
type FileStorageDeletedContainersItemDriveItemsItemDeltaResponseable interface {
    FileStorageDeletedContainersItemDriveItemsItemDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
