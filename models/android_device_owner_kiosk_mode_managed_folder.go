package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AndroidDeviceOwnerKioskModeManagedFolder a folder containing pages of apps and weblinks on the Managed Home Screen
type AndroidDeviceOwnerKioskModeManagedFolder struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAndroidDeviceOwnerKioskModeManagedFolder instantiates a new androidDeviceOwnerKioskModeManagedFolder and sets the default values.
func NewAndroidDeviceOwnerKioskModeManagedFolder()(*AndroidDeviceOwnerKioskModeManagedFolder) {
    m := &AndroidDeviceOwnerKioskModeManagedFolder{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAndroidDeviceOwnerKioskModeManagedFolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerKioskModeManagedFolderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerKioskModeManagedFolder(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidDeviceOwnerKioskModeManagedFolder) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *AndroidDeviceOwnerKioskModeManagedFolder) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerKioskModeManagedFolder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidDeviceOwnerKioskModeFolderItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidDeviceOwnerKioskModeFolderItemable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AndroidDeviceOwnerKioskModeFolderItemable)
                }
            }
            m.SetItems(res)
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
func (m *AndroidDeviceOwnerKioskModeManagedFolder) GetFolderIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("folderIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFolderName gets the folderName property value. Display name for the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolder) GetFolderName()(*string) {
    val, err := m.GetBackingStore().Get("folderName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetItems gets the items property value. Items to be added to managed folder. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerKioskModeManagedFolder) GetItems()([]AndroidDeviceOwnerKioskModeFolderItemable) {
    val, err := m.GetBackingStore().Get("items")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AndroidDeviceOwnerKioskModeFolderItemable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerKioskModeManagedFolder) GetOdataType()(*string) {
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
func (m *AndroidDeviceOwnerKioskModeManagedFolder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("folderIdentifier", m.GetFolderIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("folderName", m.GetFolderName())
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidDeviceOwnerKioskModeManagedFolder) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AndroidDeviceOwnerKioskModeManagedFolder) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFolderIdentifier sets the folderIdentifier property value. Unique identifier for the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolder) SetFolderIdentifier(value *string)() {
    err := m.GetBackingStore().Set("folderIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetFolderName sets the folderName property value. Display name for the folder
func (m *AndroidDeviceOwnerKioskModeManagedFolder) SetFolderName(value *string)() {
    err := m.GetBackingStore().Set("folderName", value)
    if err != nil {
        panic(err)
    }
}
// SetItems sets the items property value. Items to be added to managed folder. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerKioskModeManagedFolder) SetItems(value []AndroidDeviceOwnerKioskModeFolderItemable)() {
    err := m.GetBackingStore().Set("items", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidDeviceOwnerKioskModeManagedFolder) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// AndroidDeviceOwnerKioskModeManagedFolderable 
type AndroidDeviceOwnerKioskModeManagedFolderable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFolderIdentifier()(*string)
    GetFolderName()(*string)
    GetItems()([]AndroidDeviceOwnerKioskModeFolderItemable)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFolderIdentifier(value *string)()
    SetFolderName(value *string)()
    SetItems(value []AndroidDeviceOwnerKioskModeFolderItemable)()
    SetOdataType(value *string)()
}
