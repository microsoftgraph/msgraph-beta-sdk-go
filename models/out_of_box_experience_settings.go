package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// OutOfBoxExperienceSettings out of box experience setting
type OutOfBoxExperienceSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOutOfBoxExperienceSettings instantiates a new outOfBoxExperienceSettings and sets the default values.
func NewOutOfBoxExperienceSettings()(*OutOfBoxExperienceSettings) {
    m := &OutOfBoxExperienceSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOutOfBoxExperienceSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutOfBoxExperienceSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutOfBoxExperienceSettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OutOfBoxExperienceSettings) GetAdditionalData()(map[string]any) {
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
func (m *OutOfBoxExperienceSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceUsageType gets the deviceUsageType property value. The deviceUsageType property
func (m *OutOfBoxExperienceSettings) GetDeviceUsageType()(*WindowsDeviceUsageType) {
    val, err := m.GetBackingStore().Get("deviceUsageType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsDeviceUsageType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutOfBoxExperienceSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceUsageType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDeviceUsageType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceUsageType(val.(*WindowsDeviceUsageType))
        }
        return nil
    }
    res["hideEscapeLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideEscapeLink(val)
        }
        return nil
    }
    res["hideEULA"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideEULA(val)
        }
        return nil
    }
    res["hidePrivacySettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHidePrivacySettings(val)
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
    res["skipKeyboardSelectionPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkipKeyboardSelectionPage(val)
        }
        return nil
    }
    res["userType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsUserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserType(val.(*WindowsUserType))
        }
        return nil
    }
    return res
}
// GetHideEscapeLink gets the hideEscapeLink property value. If set to true, then the user can't start over with different account, on company sign-in
func (m *OutOfBoxExperienceSettings) GetHideEscapeLink()(*bool) {
    val, err := m.GetBackingStore().Get("hideEscapeLink")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetHideEULA gets the hideEULA property value. Show or hide EULA to user
func (m *OutOfBoxExperienceSettings) GetHideEULA()(*bool) {
    val, err := m.GetBackingStore().Get("hideEULA")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetHidePrivacySettings gets the hidePrivacySettings property value. Show or hide privacy settings to user
func (m *OutOfBoxExperienceSettings) GetHidePrivacySettings()(*bool) {
    val, err := m.GetBackingStore().Get("hidePrivacySettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OutOfBoxExperienceSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSkipKeyboardSelectionPage gets the skipKeyboardSelectionPage property value. If set, then skip the keyboard selection page if Language and Region are set
func (m *OutOfBoxExperienceSettings) GetSkipKeyboardSelectionPage()(*bool) {
    val, err := m.GetBackingStore().Get("skipKeyboardSelectionPage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetUserType gets the userType property value. The userType property
func (m *OutOfBoxExperienceSettings) GetUserType()(*WindowsUserType) {
    val, err := m.GetBackingStore().Get("userType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsUserType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OutOfBoxExperienceSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceUsageType() != nil {
        cast := (*m.GetDeviceUsageType()).String()
        err := writer.WriteStringValue("deviceUsageType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideEscapeLink", m.GetHideEscapeLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hideEULA", m.GetHideEULA())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hidePrivacySettings", m.GetHidePrivacySettings())
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
        err := writer.WriteBoolValue("skipKeyboardSelectionPage", m.GetSkipKeyboardSelectionPage())
        if err != nil {
            return err
        }
    }
    if m.GetUserType() != nil {
        cast := (*m.GetUserType()).String()
        err := writer.WriteStringValue("userType", &cast)
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
func (m *OutOfBoxExperienceSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *OutOfBoxExperienceSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceUsageType sets the deviceUsageType property value. The deviceUsageType property
func (m *OutOfBoxExperienceSettings) SetDeviceUsageType(value *WindowsDeviceUsageType)() {
    err := m.GetBackingStore().Set("deviceUsageType", value)
    if err != nil {
        panic(err)
    }
}
// SetHideEscapeLink sets the hideEscapeLink property value. If set to true, then the user can't start over with different account, on company sign-in
func (m *OutOfBoxExperienceSettings) SetHideEscapeLink(value *bool)() {
    err := m.GetBackingStore().Set("hideEscapeLink", value)
    if err != nil {
        panic(err)
    }
}
// SetHideEULA sets the hideEULA property value. Show or hide EULA to user
func (m *OutOfBoxExperienceSettings) SetHideEULA(value *bool)() {
    err := m.GetBackingStore().Set("hideEULA", value)
    if err != nil {
        panic(err)
    }
}
// SetHidePrivacySettings sets the hidePrivacySettings property value. Show or hide privacy settings to user
func (m *OutOfBoxExperienceSettings) SetHidePrivacySettings(value *bool)() {
    err := m.GetBackingStore().Set("hidePrivacySettings", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OutOfBoxExperienceSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSkipKeyboardSelectionPage sets the skipKeyboardSelectionPage property value. If set, then skip the keyboard selection page if Language and Region are set
func (m *OutOfBoxExperienceSettings) SetSkipKeyboardSelectionPage(value *bool)() {
    err := m.GetBackingStore().Set("skipKeyboardSelectionPage", value)
    if err != nil {
        panic(err)
    }
}
// SetUserType sets the userType property value. The userType property
func (m *OutOfBoxExperienceSettings) SetUserType(value *WindowsUserType)() {
    err := m.GetBackingStore().Set("userType", value)
    if err != nil {
        panic(err)
    }
}
// OutOfBoxExperienceSettingsable 
type OutOfBoxExperienceSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceUsageType()(*WindowsDeviceUsageType)
    GetHideEscapeLink()(*bool)
    GetHideEULA()(*bool)
    GetHidePrivacySettings()(*bool)
    GetOdataType()(*string)
    GetSkipKeyboardSelectionPage()(*bool)
    GetUserType()(*WindowsUserType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceUsageType(value *WindowsDeviceUsageType)()
    SetHideEscapeLink(value *bool)()
    SetHideEULA(value *bool)()
    SetHidePrivacySettings(value *bool)()
    SetOdataType(value *string)()
    SetSkipKeyboardSelectionPage(value *bool)()
    SetUserType(value *WindowsUserType)()
}
