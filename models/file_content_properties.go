package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileContentProperties 
type FileContentProperties struct {
    ContentProperties
}
// NewFileContentProperties instantiates a new FileContentProperties and sets the default values.
func NewFileContentProperties()(*FileContentProperties) {
    m := &FileContentProperties{
        ContentProperties: *NewContentProperties(),
    }
    odataTypeValue := "#microsoft.graph.fileContentProperties"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateFileContentPropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileContentPropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileContentProperties(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileContentProperties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ContentProperties.GetFieldDeserializers()
    res["isVisibleOnlyToOneDriveOwner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVisibleOnlyToOneDriveOwner(val)
        }
        return nil
    }
    return res
}
// GetIsVisibleOnlyToOneDriveOwner gets the isVisibleOnlyToOneDriveOwner property value. The isVisibleOnlyToOneDriveOwner property
func (m *FileContentProperties) GetIsVisibleOnlyToOneDriveOwner()(*bool) {
    val, err := m.GetBackingStore().Get("isVisibleOnlyToOneDriveOwner")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *FileContentProperties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ContentProperties.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isVisibleOnlyToOneDriveOwner", m.GetIsVisibleOnlyToOneDriveOwner())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsVisibleOnlyToOneDriveOwner sets the isVisibleOnlyToOneDriveOwner property value. The isVisibleOnlyToOneDriveOwner property
func (m *FileContentProperties) SetIsVisibleOnlyToOneDriveOwner(value *bool)() {
    err := m.GetBackingStore().Set("isVisibleOnlyToOneDriveOwner", value)
    if err != nil {
        panic(err)
    }
}
// FileContentPropertiesable 
type FileContentPropertiesable interface {
    ContentPropertiesable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsVisibleOnlyToOneDriveOwner()(*bool)
    SetIsVisibleOnlyToOneDriveOwner(value *bool)()
}
