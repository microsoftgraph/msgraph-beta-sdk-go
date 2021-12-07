package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationSettingInstance 
type DeviceManagementConfigurationSettingInstance struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Setting Definition Id
    settingDefinitionId *string;
    // Setting Instance Template Reference
    settingInstanceTemplateReference *DeviceManagementConfigurationSettingInstanceTemplateReference;
}
// NewDeviceManagementConfigurationSettingInstance instantiates a new deviceManagementConfigurationSettingInstance and sets the default values.
func NewDeviceManagementConfigurationSettingInstance()(*DeviceManagementConfigurationSettingInstance) {
    m := &DeviceManagementConfigurationSettingInstance{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingInstance) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetSettingDefinitionId gets the settingDefinitionId property value. Setting Definition Id
func (m *DeviceManagementConfigurationSettingInstance) GetSettingDefinitionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitionId
    }
}
// GetSettingInstanceTemplateReference gets the settingInstanceTemplateReference property value. Setting Instance Template Reference
func (m *DeviceManagementConfigurationSettingInstance) GetSettingInstanceTemplateReference()(*DeviceManagementConfigurationSettingInstanceTemplateReference) {
    if m == nil {
        return nil
    } else {
        return m.settingInstanceTemplateReference
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingInstance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["settingInstanceTemplateReference"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementConfigurationSettingInstanceTemplateReference() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingInstanceTemplateReference(val.(*DeviceManagementConfigurationSettingInstanceTemplateReference))
        }
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationSettingInstance) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingInstance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("settingDefinitionId", m.GetSettingDefinitionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("settingInstanceTemplateReference", m.GetSettingInstanceTemplateReference())
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
func (m *DeviceManagementConfigurationSettingInstance) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSettingDefinitionId sets the settingDefinitionId property value. Setting Definition Id
func (m *DeviceManagementConfigurationSettingInstance) SetSettingDefinitionId(value *string)() {
    if m != nil {
        m.settingDefinitionId = value
    }
}
// SetSettingInstanceTemplateReference sets the settingInstanceTemplateReference property value. Setting Instance Template Reference
func (m *DeviceManagementConfigurationSettingInstance) SetSettingInstanceTemplateReference(value *DeviceManagementConfigurationSettingInstanceTemplateReference)() {
    if m != nil {
        m.settingInstanceTemplateReference = value
    }
}
