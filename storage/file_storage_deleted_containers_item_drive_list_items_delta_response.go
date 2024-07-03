package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageDeletedContainersItemDriveListItemsDeltaGetResponseable instead.
type FileStorageDeletedContainersItemDriveListItemsDeltaResponse struct {
    FileStorageDeletedContainersItemDriveListItemsDeltaGetResponse
}
// NewFileStorageDeletedContainersItemDriveListItemsDeltaResponse instantiates a new FileStorageDeletedContainersItemDriveListItemsDeltaResponse and sets the default values.
func NewFileStorageDeletedContainersItemDriveListItemsDeltaResponse()(*FileStorageDeletedContainersItemDriveListItemsDeltaResponse) {
    m := &FileStorageDeletedContainersItemDriveListItemsDeltaResponse{
        FileStorageDeletedContainersItemDriveListItemsDeltaGetResponse: *NewFileStorageDeletedContainersItemDriveListItemsDeltaGetResponse(),
    }
    return m
}
// CreateFileStorageDeletedContainersItemDriveListItemsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageDeletedContainersItemDriveListItemsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageDeletedContainersItemDriveListItemsDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageDeletedContainersItemDriveListItemsDeltaGetResponseable instead.
type FileStorageDeletedContainersItemDriveListItemsDeltaResponseable interface {
    FileStorageDeletedContainersItemDriveListItemsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
