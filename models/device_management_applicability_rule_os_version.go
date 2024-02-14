package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DeviceManagementApplicabilityRuleOsVersion struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDeviceManagementApplicabilityRuleOsVersion instantiates a new DeviceManagementApplicabilityRuleOsVersion and sets the default values.
func NewDeviceManagementApplicabilityRuleOsVersion()(*DeviceManagementApplicabilityRuleOsVersion) {
    m := &DeviceManagementApplicabilityRuleOsVersion{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceManagementApplicabilityRuleOsVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeviceManagementApplicabilityRuleOsVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementApplicabilityRuleOsVersion(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DeviceManagementApplicabilityRuleOsVersion) GetAdditionalData()(map[string]any) {
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
// returns a BackingStore when successful
func (m *DeviceManagementApplicabilityRuleOsVersion) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DeviceManagementApplicabilityRuleOsVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maxOSVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxOSVersion(val)
        }
        return nil
    }
    res["minOSVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinOSVersion(val)
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
// GetMaxOSVersion gets the maxOSVersion property value. Max OS version for Applicability Rule.
// returns a *string when successful
func (m *DeviceManagementApplicabilityRuleOsVersion) GetMaxOSVersion()(*string) {
    val, err := m.GetBackingStore().Get("maxOSVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMinOSVersion gets the minOSVersion property value. Min OS version for Applicability Rule.
// returns a *string when successful
func (m *DeviceManagementApplicabilityRuleOsVersion) GetMinOSVersion()(*string) {
    val, err := m.GetBackingStore().Get("minOSVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetName gets the name property value. Name for object.
// returns a *string when successful
func (m *DeviceManagementApplicabilityRuleOsVersion) GetName()(*string) {
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
// returns a *string when successful
func (m *DeviceManagementApplicabilityRuleOsVersion) GetOdataType()(*string) {
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
// returns a *DeviceManagementApplicabilityRuleType when successful
func (m *DeviceManagementApplicabilityRuleOsVersion) GetRuleType()(*DeviceManagementApplicabilityRuleType) {
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
func (m *DeviceManagementApplicabilityRuleOsVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("maxOSVersion", m.GetMaxOSVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("minOSVersion", m.GetMinOSVersion())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetMaxOSVersion sets the maxOSVersion property value. Max OS version for Applicability Rule.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetMaxOSVersion(value *string)() {
    err := m.GetBackingStore().Set("maxOSVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetMinOSVersion sets the minOSVersion property value. Min OS version for Applicability Rule.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetMinOSVersion(value *string)() {
    err := m.GetBackingStore().Set("minOSVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementApplicabilityRuleOsVersion) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRuleType sets the ruleType property value. Supported Applicability rule types for Device Configuration
func (m *DeviceManagementApplicabilityRuleOsVersion) SetRuleType(value *DeviceManagementApplicabilityRuleType)() {
    err := m.GetBackingStore().Set("ruleType", value)
    if err != nil {
        panic(err)
    }
}
type DeviceManagementApplicabilityRuleOsVersionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetMaxOSVersion()(*string)
    GetMinOSVersion()(*string)
    GetName()(*string)
    GetOdataType()(*string)
    GetRuleType()(*DeviceManagementApplicabilityRuleType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetMaxOSVersion(value *string)()
    SetMinOSVersion(value *string)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetRuleType(value *DeviceManagementApplicabilityRuleType)()
}
