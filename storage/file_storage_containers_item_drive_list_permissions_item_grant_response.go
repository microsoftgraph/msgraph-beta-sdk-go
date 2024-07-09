package storage

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveListPermissionsItemGrantPostResponseable instead.
type FileStorageContainersItemDriveListPermissionsItemGrantResponse struct {
    FileStorageContainersItemDriveListPermissionsItemGrantPostResponse
}
// NewFileStorageContainersItemDriveListPermissionsItemGrantResponse instantiates a new FileStorageContainersItemDriveListPermissionsItemGrantResponse and sets the default values.
func NewFileStorageContainersItemDriveListPermissionsItemGrantResponse()(*FileStorageContainersItemDriveListPermissionsItemGrantResponse) {
    m := &FileStorageContainersItemDriveListPermissionsItemGrantResponse{
        FileStorageContainersItemDriveListPermissionsItemGrantPostResponse: *NewFileStorageContainersItemDriveListPermissionsItemGrantPostResponse(),
    }
    return m
}
// CreateFileStorageContainersItemDriveListPermissionsItemGrantResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileStorageContainersItemDriveListPermissionsItemGrantResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileStorageContainersItemDriveListPermissionsItemGrantResponse(), nil
}
// Deprecated: This class is obsolete. Use FileStorageContainersItemDriveListPermissionsItemGrantPostResponseable instead.
type FileStorageContainersItemDriveListPermissionsItemGrantResponseable interface {
    FileStorageContainersItemDriveListPermissionsItemGrantPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
