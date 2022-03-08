package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationSettingTemplate provides operations to manage the deviceManagement singleton.
type DeviceManagementConfigurationSettingTemplate struct {
    Entity
    // List of related Setting Definitions
    settingDefinitions []DeviceManagementConfigurationSettingDefinitionable;
    // Setting Instance Template
    settingInstanceTemplate DeviceManagementConfigurationSettingInstanceTemplateable;
}
// NewDeviceManagementConfigurationSettingTemplate instantiates a new deviceManagementConfigurationSettingTemplate and sets the default values.
func NewDeviceManagementConfigurationSettingTemplate()(*DeviceManagementConfigurationSettingTemplate) {
    m := &DeviceManagementConfigurationSettingTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementConfigurationSettingTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSettingTemplateFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceManagementConfigurationSettingTemplate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["settingDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationSettingDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationSettingDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationSettingDefinitionable)
            }
            m.SetSettingDefinitions(res)
        }
        return nil
    }
    res["settingInstanceTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSettingInstanceTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingInstanceTemplate(val.(DeviceManagementConfigurationSettingInstanceTemplateable))
        }
        return nil
    }
    return res
}
// GetSettingDefinitions gets the settingDefinitions property value. List of related Setting Definitions
func (m *DeviceManagementConfigurationSettingTemplate) GetSettingDefinitions()([]DeviceManagementConfigurationSettingDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitions
    }
}
// GetSettingInstanceTemplate gets the settingInstanceTemplate property value. Setting Instance Template
func (m *DeviceManagementConfigurationSettingTemplate) GetSettingInstanceTemplate()(DeviceManagementConfigurationSettingInstanceTemplateable) {
    if m == nil {
        return nil
    } else {
        return m.settingInstanceTemplate
    }
}
func (m *DeviceManagementConfigurationSettingTemplate) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSettingDefinitions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettingDefinitions()))
        for i, v := range m.GetSettingDefinitions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("settingDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settingInstanceTemplate", m.GetSettingInstanceTemplate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSettingDefinitions sets the settingDefinitions property value. List of related Setting Definitions
func (m *DeviceManagementConfigurationSettingTemplate) SetSettingDefinitions(value []DeviceManagementConfigurationSettingDefinitionable)() {
    if m != nil {
        m.settingDefinitions = value
    }
}
// SetSettingInstanceTemplate sets the settingInstanceTemplate property value. Setting Instance Template
func (m *DeviceManagementConfigurationSettingTemplate) SetSettingInstanceTemplate(value DeviceManagementConfigurationSettingInstanceTemplateable)() {
    if m != nil {
        m.settingInstanceTemplate = value
    }
}
