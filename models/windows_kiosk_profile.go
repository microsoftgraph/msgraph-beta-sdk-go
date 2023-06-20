package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WindowsKioskProfile 
type WindowsKioskProfile struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsKioskProfile instantiates a new windowsKioskProfile and sets the default values.
func NewWindowsKioskProfile()(*WindowsKioskProfile) {
    m := &WindowsKioskProfile{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsKioskProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsKioskProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsKioskProfile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsKioskProfile) GetAdditionalData()(map[string]any) {
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
// GetAppConfiguration gets the appConfiguration property value. The app base class used to identify the application info for the kiosk configuration
func (m *WindowsKioskProfile) GetAppConfiguration()(WindowsKioskAppConfigurationable) {
    val, err := m.GetBackingStore().Get("appConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(WindowsKioskAppConfigurationable)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *WindowsKioskProfile) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsKioskProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsKioskAppConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppConfiguration(val.(WindowsKioskAppConfigurationable))
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
    res["profileId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileId(val)
        }
        return nil
    }
    res["profileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfileName(val)
        }
        return nil
    }
    res["userAccountsConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsKioskUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsKioskUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WindowsKioskUserable)
                }
            }
            m.SetUserAccountsConfiguration(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsKioskProfile) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProfileId gets the profileId property value. Key of the entity.
func (m *WindowsKioskProfile) GetProfileId()(*string) {
    val, err := m.GetBackingStore().Get("profileId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetProfileName gets the profileName property value. This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the users to whom this kiosk configuration is assigned.
func (m *WindowsKioskProfile) GetProfileName()(*string) {
    val, err := m.GetBackingStore().Get("profileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserAccountsConfiguration gets the userAccountsConfiguration property value. The user accounts that will be locked to this kiosk configuration. This collection can contain a maximum of 100 elements.
func (m *WindowsKioskProfile) GetUserAccountsConfiguration()([]WindowsKioskUserable) {
    val, err := m.GetBackingStore().Get("userAccountsConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WindowsKioskUserable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsKioskProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("appConfiguration", m.GetAppConfiguration())
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
        err := writer.WriteStringValue("profileId", m.GetProfileId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("profileName", m.GetProfileName())
        if err != nil {
            return err
        }
    }
    if m.GetUserAccountsConfiguration() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserAccountsConfiguration()))
        for i, v := range m.GetUserAccountsConfiguration() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("userAccountsConfiguration", cast)
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsKioskProfile) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAppConfiguration sets the appConfiguration property value. The app base class used to identify the application info for the kiosk configuration
func (m *WindowsKioskProfile) SetAppConfiguration(value WindowsKioskAppConfigurationable)() {
    err := m.GetBackingStore().Set("appConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *WindowsKioskProfile) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsKioskProfile) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileId sets the profileId property value. Key of the entity.
func (m *WindowsKioskProfile) SetProfileId(value *string)() {
    err := m.GetBackingStore().Set("profileId", value)
    if err != nil {
        panic(err)
    }
}
// SetProfileName sets the profileName property value. This is a friendly name used to identify a group of applications, the layout of these apps on the start menu and the users to whom this kiosk configuration is assigned.
func (m *WindowsKioskProfile) SetProfileName(value *string)() {
    err := m.GetBackingStore().Set("profileName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserAccountsConfiguration sets the userAccountsConfiguration property value. The user accounts that will be locked to this kiosk configuration. This collection can contain a maximum of 100 elements.
func (m *WindowsKioskProfile) SetUserAccountsConfiguration(value []WindowsKioskUserable)() {
    err := m.GetBackingStore().Set("userAccountsConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// WindowsKioskProfileable 
type WindowsKioskProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppConfiguration()(WindowsKioskAppConfigurationable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetProfileId()(*string)
    GetProfileName()(*string)
    GetUserAccountsConfiguration()([]WindowsKioskUserable)
    SetAppConfiguration(value WindowsKioskAppConfigurationable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetProfileId(value *string)()
    SetProfileName(value *string)()
    SetUserAccountsConfiguration(value []WindowsKioskUserable)()
}
