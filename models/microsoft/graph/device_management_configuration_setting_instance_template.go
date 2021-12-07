package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationSettingInstanceTemplate 
type DeviceManagementConfigurationSettingInstanceTemplate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates if a policy must specify this setting.
    isRequired *bool;
    // Setting Definition Id
    settingDefinitionId *string;
    // Setting Instance Template Id
    settingInstanceTemplateId *string;
}
// NewDeviceManagementConfigurationSettingInstanceTemplate instantiates a new deviceManagementConfigurationSettingInstanceTemplate and sets the default values.
func NewDeviceManagementConfigurationSettingInstanceTemplate()(*DeviceManagementConfigurationSettingInstanceTemplate) {
    m := &DeviceManagementConfigurationSettingInstanceTemplate{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsRequired gets the isRequired property value. Indicates if a policy must specify this setting.
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetIsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRequired
    }
}
// GetSettingDefinitionId gets the settingDefinitionId property value. Setting Definition Id
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetSettingDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitionId
    }
}
// GetSettingInstanceTemplateId gets the settingInstanceTemplateId property value. Setting Instance Template Id
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetSettingInstanceTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingInstanceTemplateId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingInstanceTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRequired(val)
        }
        return nil
    }
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
func (m *DeviceManagementConfigurationSettingInstanceTemplate) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingInstanceTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isRequired", m.GetIsRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingDefinitionId", m.GetSettingDefinitionId())
        if err != nil {
            return err
        }
    }
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingInstanceTemplate) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsRequired sets the isRequired property value. Indicates if a policy must specify this setting.
func (m *DeviceManagementConfigurationSettingInstanceTemplate) SetIsRequired(value *bool)() {
    if m != nil {
        m.isRequired = value
    }
}
// SetSettingDefinitionId sets the settingDefinitionId property value. Setting Definition Id
func (m *DeviceManagementConfigurationSettingInstanceTemplate) SetSettingDefinitionId(value *string)() {
    if m != nil {
        m.settingDefinitionId = value
    }
}
// SetSettingInstanceTemplateId sets the settingInstanceTemplateId property value. Setting Instance Template Id
func (m *DeviceManagementConfigurationSettingInstanceTemplate) SetSettingInstanceTemplateId(value *string)() {
    if m != nil {
        m.settingInstanceTemplateId = value
    }
}
