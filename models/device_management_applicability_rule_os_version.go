package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementApplicabilityRuleOsVersion 
type DeviceManagementApplicabilityRuleOsVersion struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Max OS version for Applicability Rule.
    maxOSVersion *string
    // Min OS version for Applicability Rule.
    minOSVersion *string
    // Name for object.
    name *string
    // The OdataType property
    odataType *string
    // Supported Applicability rule types for Device Configuration
    ruleType *DeviceManagementApplicabilityRuleType
}
// NewDeviceManagementApplicabilityRuleOsVersion instantiates a new deviceManagementApplicabilityRuleOsVersion and sets the default values.
func NewDeviceManagementApplicabilityRuleOsVersion()(*DeviceManagementApplicabilityRuleOsVersion) {
    m := &DeviceManagementApplicabilityRuleOsVersion{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.deviceManagementApplicabilityRuleOsVersion";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementApplicabilityRuleOsVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementApplicabilityRuleOsVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementApplicabilityRuleOsVersion(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementApplicabilityRuleOsVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maxOSVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMaxOSVersion)
    res["minOSVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMinOSVersion)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["ruleType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceManagementApplicabilityRuleType , m.SetRuleType)
    return res
}
// GetMaxOSVersion gets the maxOSVersion property value. Max OS version for Applicability Rule.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetMaxOSVersion()(*string) {
    return m.maxOSVersion
}
// GetMinOSVersion gets the minOSVersion property value. Min OS version for Applicability Rule.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetMinOSVersion()(*string) {
    return m.minOSVersion
}
// GetName gets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleOsVersion) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementApplicabilityRuleOsVersion) GetOdataType()(*string) {
    return m.odataType
}
// GetRuleType gets the ruleType property value. Supported Applicability rule types for Device Configuration
func (m *DeviceManagementApplicabilityRuleOsVersion) GetRuleType()(*DeviceManagementApplicabilityRuleType) {
    return m.ruleType
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMaxOSVersion sets the maxOSVersion property value. Max OS version for Applicability Rule.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetMaxOSVersion(value *string)() {
    m.maxOSVersion = value
}
// SetMinOSVersion sets the minOSVersion property value. Min OS version for Applicability Rule.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetMinOSVersion(value *string)() {
    m.minOSVersion = value
}
// SetName sets the name property value. Name for object.
func (m *DeviceManagementApplicabilityRuleOsVersion) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementApplicabilityRuleOsVersion) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRuleType sets the ruleType property value. Supported Applicability rule types for Device Configuration
func (m *DeviceManagementApplicabilityRuleOsVersion) SetRuleType(value *DeviceManagementApplicabilityRuleType)() {
    m.ruleType = value
}
