package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DeviceManagementApplicabilityRuleDeviceMode 
type DeviceManagementApplicabilityRuleDeviceMode struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceManagementApplicabilityRuleDeviceMode instantiates a new deviceManagementApplicabilityRuleDeviceMode and sets the default values.
func NewDeviceManagementApplicabilityRuleDeviceMode()(*DeviceManagementApplicabilityRuleDeviceMode) {
    m := &DeviceManagementApplicabilityRuleDeviceMode{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceManagementApplicabilityRuleDeviceModeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementApplicabilityRuleDeviceModeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementApplicabilityRuleDeviceMode(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDeviceMode gets the deviceMode property value. Windows 10 Device Mode type.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetDeviceMode()(*Windows10DeviceModeType) {
    val, err := m.GetBackingStore().Get("deviceMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Windows10DeviceModeType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10DeviceModeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceMode(val.(*Windows10DeviceModeType))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["ruleType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementApplicabilityRuleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRuleType(val.(*DeviceManagementApplicabilityRuleType))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRuleType gets the ruleType property value. Supported Applicability rule types for Device Configuration
func (m *DeviceManagementApplicabilityRuleDeviceMode) GetRuleType()(*DeviceManagementApplicabilityRuleType) {
    val, err := m.GetBackingStore().Get("ruleType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceManagementApplicabilityRuleType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceManagementApplicabilityRuleDeviceMode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceMode() != nil {
        cast := (*m.GetDeviceMode()).String()
        err := writer.WriteStringValue("deviceMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
    if m.GetRuleType() != nil {
        cast := (*m.GetRuleType()).String()
        err := writer.WriteStringValue("ruleType", &cast)
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
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDeviceMode sets the deviceMode property value. Windows 10 Device Mode type.
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetDeviceMode(value *Windows10DeviceModeType)() {
    err := m.GetBackingStore().Set("deviceMode", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRuleType sets the ruleType property value. Supported Applicability rule types for Device Configuration
func (m *DeviceManagementApplicabilityRuleDeviceMode) SetRuleType(value *DeviceManagementApplicabilityRuleType)() {
    err := m.GetBackingStore().Set("ruleType", value)
    if err != nil {
        panic(err)
    }
}
// DeviceManagementApplicabilityRuleDeviceModeable 
type DeviceManagementApplicabilityRuleDeviceModeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDeviceMode()(*Windows10DeviceModeType)
    GetName()(*string)
    GetOdataType()(*string)
    GetRuleType()(*DeviceManagementApplicabilityRuleType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDeviceMode(value *Windows10DeviceModeType)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetRuleType(value *DeviceManagementApplicabilityRuleType)()
}
