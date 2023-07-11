package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerKioskModeManagedFolderReference a reference to folder containing apps and weblinks on the Managed Home Screen
type AndroidDeviceOwnerKioskModeManagedFolderReference struct {
    AndroidDeviceOwnerKioskModeHomeScreenItem
}
// NewAndroidDeviceOwnerKioskModeManagedFolderReference instantiates a new androidDeviceOwnerKioskModeManagedFolderReference and sets the default values.
func NewAndroidDeviceOwnerKioskModeManagedFolderReference()(*AndroidDeviceOwnerKioskModeManagedFolderReference) {
    m := &AndroidDeviceOwnerKioskModeManagedFolderReference{
        AndroidDeviceOwnerKioskModeHomeScreenItem: *NewAndroidDeviceOwnerKioskModeHomeScreenItem(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerKioskModeManagedFolderReference"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidDeviceOwnerKioskModeManagedFolderReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerKioskModeManagedFolderReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerKioskModeManagedFolderReference(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AndroidDeviceOwnerKioskModeHomeScreenItem.GetFieldDeserializers()
    res["folderIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolderIdentifier(val)
        }
        return nil
    }
    res["folderName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolderName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetFolderIdentifier gets the folderIdentifier property value. Unique identifier for the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) GetFolderIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("folderIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFolderName gets the folderName property value. Name of the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) GetFolderName()(*string) {
    val, err := m.GetBackingStore().Get("folderName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AndroidDeviceOwnerKioskModeHomeScreenItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("folderIdentifier", m.GetFolderIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("folderName", m.GetFolderName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFolderIdentifier sets the folderIdentifier property value. Unique identifier for the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) SetFolderIdentifier(value *string)() {
    err := m.GetBackingStore().Set("folderIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetFolderName sets the folderName property value. Name of the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) SetFolderName(value *string)() {
    err := m.GetBackingStore().Set("folderName", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerKioskModeManagedFolderReference) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// AndroidDeviceOwnerKioskModeManagedFolderReferenceable 
type AndroidDeviceOwnerKioskModeManagedFolderReferenceable interface {
    AndroidDeviceOwnerKioskModeHomeScreenItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFolderIdentifier()(*string)
    GetFolderName()(*string)
    GetOdataType()(*string)
    SetFolderIdentifier(value *string)()
    SetFolderName(value *string)()
    SetOdataType(value *string)()
}
