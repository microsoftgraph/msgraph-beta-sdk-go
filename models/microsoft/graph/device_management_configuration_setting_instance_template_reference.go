package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementConfigurationSettingInstanceTemplateReference struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Setting instance template id
    settingInstanceTemplateId *string;
}
// Instantiates a new deviceManagementConfigurationSettingInstanceTemplateReference and sets the default values.
func NewDeviceManagementConfigurationSettingInstanceTemplateReference()(*DeviceManagementConfigurationSettingInstanceTemplateReference) {
    m := &DeviceManagementConfigurationSettingInstanceTemplateReference{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the settingInstanceTemplateId property value. Setting instance template id
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) GetSettingInstanceTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingInstanceTemplateId
    }
}
// The deserialization information for the current model
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["settingInstanceTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingInstanceTemplateId(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("settingInstanceTemplateId", m.GetSettingInstanceTemplateId())
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
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the settingInstanceTemplateId property value. Setting instance template id
// Parameters:
//  - value : Value to set for the settingInstanceTemplateId property.
func (m *DeviceManagementConfigurationSettingInstanceTemplateReference) SetSettingInstanceTemplateId(value *string)() {
    m.settingInstanceTemplateId = value
}
