package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageDeletedContainersItemDriveSearchWithQGetResponseable instead.
type FileStorageDeletedContainersItemDriveSearchWithQResponse struct {
    FileStorageDeletedContainersItemDriveSearchWithQGetResponse
}
// NewFileStorageDeletedContainersItemDriveSearchWithQResponse instantiates a new FileStorageDeletedContainersItemDriveSearchWithQResponse and sets the default values.
func NewFileStorageDeletedContainersItemDriveSearchWithQResponse()(*FileStorageDeletedContainersItemDriveSearchWithQResponse) {
    m := &FileStorageDeletedContainersItemDriveSearchWithQResponse{
        FileStorageDeletedContainersItemDriveSearchWithQGetResponse: *NewFileStorageDeletedContainersItemDriveSearchWithQGetResponse(),
    }
    return m
}
// CreateFileStorageDeletedContainersItemDriveSearchWithQResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageDeletedContainersItemDriveSearchWithQResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageDeletedContainersItemDriveSearchWithQResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageDeletedContainersItemDriveSearchWithQGetResponseable instead.
type FileStorageDeletedContainersItemDriveSearchWithQResponseable interface {
    FileStorageDeletedContainersItemDriveSearchWithQGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
