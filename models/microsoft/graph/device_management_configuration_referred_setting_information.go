package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementConfigurationReferredSettingInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Setting definition id that is being referred to a setting. Applicable for reusable setting
    settingDefinitionId *string;
}
// Instantiates a new deviceManagementConfigurationReferredSettingInformation and sets the default values.
func NewDeviceManagementConfigurationReferredSettingInformation()(*DeviceManagementConfigurationReferredSettingInformation) {
    m := &DeviceManagementConfigurationReferredSettingInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationReferredSettingInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the settingDefinitionId property value. Setting definition id that is being referred to a setting. Applicable for reusable setting
func (m *DeviceManagementConfigurationReferredSettingInformation) GetSettingDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitionId
    }
}
// The deserialization information for the current model
func (m *DeviceManagementConfigurationReferredSettingInformation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["settingDefinitionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDefinitionId(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationReferredSettingInformation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementConfigurationReferredSettingInformation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("settingDefinitionId", m.GetSettingDefinitionId())
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
func (m *DeviceManagementConfigurationReferredSettingInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the settingDefinitionId property value. Setting definition id that is being referred to a setting. Applicable for reusable setting
// Parameters:
//  - value : Value to set for the settingDefinitionId property.
func (m *DeviceManagementConfigurationReferredSettingInformation) SetSettingDefinitionId(value *string)() {
    m.settingDefinitionId = value
}
