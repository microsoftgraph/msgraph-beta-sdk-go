package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementConfigurationSettingApplicability struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // description of the setting
    description *string;
    // Device Mode that setting can be applied on. Possible values are: none, kiosk.
    deviceMode *DeviceManagementConfigurationDeviceMode;
    // Platform setting can be applied on. Possible values are: none, android, iOS, macOS, windows10X, windows10.
    platform *DeviceManagementConfigurationPlatforms;
    // Which technology channels this setting can be deployed through. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
    technologies *DeviceManagementConfigurationTechnologies;
}
// Instantiates a new deviceManagementConfigurationSettingApplicability and sets the default values.
func NewDeviceManagementConfigurationSettingApplicability()(*DeviceManagementConfigurationSettingApplicability) {
    m := &DeviceManagementConfigurationSettingApplicability{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingApplicability) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the description property value. description of the setting
func (m *DeviceManagementConfigurationSettingApplicability) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the deviceMode property value. Device Mode that setting can be applied on. Possible values are: none, kiosk.
func (m *DeviceManagementConfigurationSettingApplicability) GetDeviceMode()(*DeviceManagementConfigurationDeviceMode) {
    if m == nil {
        return nil
    } else {
        return m.deviceMode
    }
}
// Gets the platform property value. Platform setting can be applied on. Possible values are: none, android, iOS, macOS, windows10X, windows10.
func (m *DeviceManagementConfigurationSettingApplicability) GetPlatform()(*DeviceManagementConfigurationPlatforms) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// Gets the technologies property value. Which technology channels this setting can be deployed through. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
func (m *DeviceManagementConfigurationSettingApplicability) GetTechnologies()(*DeviceManagementConfigurationTechnologies) {
    if m == nil {
        return nil
    } else {
        return m.technologies
    }
}
// The deserialization information for the current model
func (m *DeviceManagementConfigurationSettingApplicability) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["deviceMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationDeviceMode)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementConfigurationDeviceMode)
        m.SetDeviceMode(&cast)
        return nil
    }
    res["platform"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationPlatforms)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementConfigurationPlatforms)
        m.SetPlatform(&cast)
        return nil
    }
    res["technologies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationTechnologies)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementConfigurationTechnologies)
        m.SetTechnologies(&cast)
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationSettingApplicability) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementConfigurationSettingApplicability) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceMode() != nil {
        cast := m.GetDeviceMode().String()
        err := writer.WriteStringValue("deviceMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := m.GetPlatform().String()
        err := writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTechnologies() != nil {
        cast := m.GetTechnologies().String()
        err := writer.WriteStringValue("technologies", &cast)
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeviceManagementConfigurationSettingApplicability) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the description property value. description of the setting
// Parameters:
//  - value : Value to set for the description property.
func (m *DeviceManagementConfigurationSettingApplicability) SetDescription(value *string)() {
    m.description = value
}
// Sets the deviceMode property value. Device Mode that setting can be applied on. Possible values are: none, kiosk.
// Parameters:
//  - value : Value to set for the deviceMode property.
func (m *DeviceManagementConfigurationSettingApplicability) SetDeviceMode(value *DeviceManagementConfigurationDeviceMode)() {
    m.deviceMode = value
}
// Sets the platform property value. Platform setting can be applied on. Possible values are: none, android, iOS, macOS, windows10X, windows10.
// Parameters:
//  - value : Value to set for the platform property.
func (m *DeviceManagementConfigurationSettingApplicability) SetPlatform(value *DeviceManagementConfigurationPlatforms)() {
    m.platform = value
}
// Sets the technologies property value. Which technology channels this setting can be deployed through. Possible values are: none, mdm, windows10XManagement, configManager, microsoftSense, exchangeOnline, linuxMdm, unknownFutureValue.
// Parameters:
//  - value : Value to set for the technologies property.
func (m *DeviceManagementConfigurationSettingApplicability) SetTechnologies(value *DeviceManagementConfigurationTechnologies)() {
    m.technologies = value
}
