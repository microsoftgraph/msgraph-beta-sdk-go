package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceScriptRuleError 
type DeviceComplianceScriptRuleError struct {
    DeviceComplianceScriptError
    // Setting name for the rule with error.
    settingName *string
}
// NewDeviceComplianceScriptRuleError instantiates a new DeviceComplianceScriptRuleError and sets the default values.
func NewDeviceComplianceScriptRuleError()(*DeviceComplianceScriptRuleError) {
    m := &DeviceComplianceScriptRuleError{
        DeviceComplianceScriptError: *NewDeviceComplianceScriptError(),
    }
    odataTypeValue := "#microsoft.graph.deviceComplianceScriptRuleError";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceComplianceScriptRuleErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComplianceScriptRuleErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceComplianceScriptRuleError(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceScriptRuleError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceComplianceScriptError.GetFieldDeserializers()
    res["settingName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSettingName)
    return res
}
// GetSettingName gets the settingName property value. Setting name for the rule with error.
func (m *DeviceComplianceScriptRuleError) GetSettingName()(*string) {
    return m.settingName
}
// Serialize serializes information the current object
func (m *DeviceComplianceScriptRuleError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceComplianceScriptError.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSettingName sets the settingName property value. Setting name for the rule with error.
func (m *DeviceComplianceScriptRuleError) SetSettingName(value *string)() {
    m.settingName = value
}
