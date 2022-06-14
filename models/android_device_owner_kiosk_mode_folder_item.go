package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerKioskModeFolderItem represents an item that can be added to Android Device Owner folder (application or weblink)
type AndroidDeviceOwnerKioskModeFolderItem struct {
    AndroidDeviceOwnerKioskModeHomeScreenItem
}
// NewAndroidDeviceOwnerKioskModeFolderItem instantiates a new androidDeviceOwnerKioskModeFolderItem and sets the default values.
func NewAndroidDeviceOwnerKioskModeFolderItem()(*AndroidDeviceOwnerKioskModeFolderItem) {
    m := &AndroidDeviceOwnerKioskModeFolderItem{
        AndroidDeviceOwnerKioskModeHomeScreenItem: *NewAndroidDeviceOwnerKioskModeHomeScreenItem(),
    }
    return m
}
// CreateAndroidDeviceOwnerKioskModeFolderItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerKioskModeFolderItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerKioskModeFolderItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerKioskModeFolderItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceOwnerKioskModeHomeScreenItem.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerKioskModeFolderItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidDeviceOwnerKioskModeHomeScreenItem.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
