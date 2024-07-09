package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageDeletedContainersItemDriveSharedWithMeGetResponseable instead.
type FileStorageDeletedContainersItemDriveSharedWithMeResponse struct {
    FileStorageDeletedContainersItemDriveSharedWithMeGetResponse
}
// NewFileStorageDeletedContainersItemDriveSharedWithMeResponse instantiates a new FileStorageDeletedContainersItemDriveSharedWithMeResponse and sets the default values.
func NewFileStorageDeletedContainersItemDriveSharedWithMeResponse()(*FileStorageDeletedContainersItemDriveSharedWithMeResponse) {
    m := &FileStorageDeletedContainersItemDriveSharedWithMeResponse{
        FileStorageDeletedContainersItemDriveSharedWithMeGetResponse: *NewFileStorageDeletedContainersItemDriveSharedWithMeGetResponse(),
    }
    return m
}
// CreateFileStorageDeletedContainersItemDriveSharedWithMeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageDeletedContainersItemDriveSharedWithMeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageDeletedContainersItemDriveSharedWithMeResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageDeletedContainersItemDriveSharedWithMeGetResponseable instead.
type FileStorageDeletedContainersItemDriveSharedWithMeResponseable interface {
    FileStorageDeletedContainersItemDriveSharedWithMeGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
