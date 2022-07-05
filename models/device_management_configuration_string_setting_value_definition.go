package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationStringSettingValueDefinition 
type DeviceManagementConfigurationStringSettingValueDefinition struct {
    DeviceManagementConfigurationSettingValueDefinition
    // Pre-defined format of the string. Possible values are: none, email, guid, ip, base64, url, version, xml, date, time, binary, regEx, json, dateTime, surfaceHub.
    format *DeviceManagementConfigurationStringFormat
    // Regular expression or any xml or json schema that the input string should match
    inputValidationSchema *string
    // Specifies whether the setting needs to be treated as a secret. Settings marked as yes will be encrypted in transit and at rest and will be displayed as asterisks when represented in the UX.
    isSecret *bool
    // Maximum length of string. Valid values 0 to 87516
    maximumLength *int64
    // Minimum length of string. Valid values 0 to 87516
    minimumLength *int64
}
// NewDeviceManagementConfigurationStringSettingValueDefinition instantiates a new DeviceManagementConfigurationStringSettingValueDefinition and sets the default values.
func NewDeviceManagementConfigurationStringSettingValueDefinition()(*DeviceManagementConfigurationStringSettingValueDefinition) {
    m := &DeviceManagementConfigurationStringSettingValueDefinition{
        DeviceManagementConfigurationSettingValueDefinition: *NewDeviceManagementConfigurationSettingValueDefinition(),
    }
    return m
}
// CreateDeviceManagementConfigurationStringSettingValueDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationStringSettingValueDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationStringSettingValueDefinition(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationStringSettingValueDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingValueDefinition.GetFieldDeserializers()
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationStringFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*DeviceManagementConfigurationStringFormat))
        }
        return nil
    }
    res["inputValidationSchema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInputValidationSchema(val)
        }
        return nil
    }
    res["isSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSecret(val)
        }
        return nil
    }
    res["maximumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumLength(val)
        }
        return nil
    }
    res["minimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumLength(val)
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. Pre-defined format of the string. Possible values are: none, email, guid, ip, base64, url, version, xml, date, time, binary, regEx, json, dateTime, surfaceHub.
func (m *DeviceManagementConfigurationStringSettingValueDefinition) GetFormat()(*DeviceManagementConfigurationStringFormat) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetInputValidationSchema gets the inputValidationSchema property value. Regular expression or any xml or json schema that the input string should match
func (m *DeviceManagementConfigurationStringSettingValueDefinition) GetInputValidationSchema()(*string) {
    if m == nil {
        return nil
    } else {
        return m.inputValidationSchema
    }
}
// GetIsSecret gets the isSecret property value. Specifies whether the setting needs to be treated as a secret. Settings marked as yes will be encrypted in transit and at rest and will be displayed as asterisks when represented in the UX.
func (m *DeviceManagementConfigurationStringSettingValueDefinition) GetIsSecret()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSecret
    }
}
// GetMaximumLength gets the maximumLength property value. Maximum length of string. Valid values 0 to 87516
func (m *DeviceManagementConfigurationStringSettingValueDefinition) GetMaximumLength()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.maximumLength
    }
}
// GetMinimumLength gets the minimumLength property value. Minimum length of string. Valid values 0 to 87516
func (m *DeviceManagementConfigurationStringSettingValueDefinition) GetMinimumLength()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.minimumLength
    }
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationStringSettingValueDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingValueDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetFormat() != nil {
        cast := (*m.GetFormat()).String()
        err = writer.WriteStringValue("format", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("inputValidationSchema", m.GetInputValidationSchema())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSecret", m.GetIsSecret())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("maximumLength", m.GetMaximumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("minimumLength", m.GetMinimumLength())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFormat sets the format property value. Pre-defined format of the string. Possible values are: none, email, guid, ip, base64, url, version, xml, date, time, binary, regEx, json, dateTime, surfaceHub.
func (m *DeviceManagementConfigurationStringSettingValueDefinition) SetFormat(value *DeviceManagementConfigurationStringFormat)() {
    if m != nil {
        m.format = value
    }
}
// SetInputValidationSchema sets the inputValidationSchema property value. Regular expression or any xml or json schema that the input string should match
func (m *DeviceManagementConfigurationStringSettingValueDefinition) SetInputValidationSchema(value *string)() {
    if m != nil {
        m.inputValidationSchema = value
    }
}
// SetIsSecret sets the isSecret property value. Specifies whether the setting needs to be treated as a secret. Settings marked as yes will be encrypted in transit and at rest and will be displayed as asterisks when represented in the UX.
func (m *DeviceManagementConfigurationStringSettingValueDefinition) SetIsSecret(value *bool)() {
    if m != nil {
        m.isSecret = value
    }
}
// SetMaximumLength sets the maximumLength property value. Maximum length of string. Valid values 0 to 87516
func (m *DeviceManagementConfigurationStringSettingValueDefinition) SetMaximumLength(value *int64)() {
    if m != nil {
        m.maximumLength = value
    }
}
// SetMinimumLength sets the minimumLength property value. Minimum length of string. Valid values 0 to 87516
func (m *DeviceManagementConfigurationStringSettingValueDefinition) SetMinimumLength(value *int64)() {
    if m != nil {
        m.minimumLength = value
    }
}
