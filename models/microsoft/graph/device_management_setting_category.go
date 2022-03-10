package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementSettingCategory provides operations to manage the deviceManagement singleton.
type DeviceManagementSettingCategory struct {
    Entity
    // The category name
    displayName *string;
    // The category contains top level required setting
    hasRequiredSetting *bool;
    // The setting definitions this category contains
    settingDefinitions []DeviceManagementSettingDefinitionable;
}
// NewDeviceManagementSettingCategory instantiates a new deviceManagementSettingCategory and sets the default values.
func NewDeviceManagementSettingCategory()(*DeviceManagementSettingCategory) {
    m := &DeviceManagementSettingCategory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementSettingCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingCategoryFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceManagementSettingCategory(), nil
}
// GetDisplayName gets the displayName property value. The category name
func (m *DeviceManagementSettingCategory) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["hasRequiredSetting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasRequiredSetting(val)
        }
        return nil
    }
    res["settingDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementSettingDefinitionable)
            }
            m.SetSettingDefinitions(res)
        }
        return nil
    }
    return res
}
// GetHasRequiredSetting gets the hasRequiredSetting property value. The category contains top level required setting
func (m *DeviceManagementSettingCategory) GetHasRequiredSetting()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasRequiredSetting
    }
}
// GetSettingDefinitions gets the settingDefinitions property value. The setting definitions this category contains
func (m *DeviceManagementSettingCategory) GetSettingDefinitions()([]DeviceManagementSettingDefinitionable) {
    if m == nil {
        return nil
    } else {
        return m.settingDefinitions
    }
}
func (m *DeviceManagementSettingCategory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasRequiredSetting", m.GetHasRequiredSetting())
        if err != nil {
            return err
        }
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
    return nil
}
// SetDisplayName sets the displayName property value. The category name
func (m *DeviceManagementSettingCategory) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetHasRequiredSetting sets the hasRequiredSetting property value. The category contains top level required setting
func (m *DeviceManagementSettingCategory) SetHasRequiredSetting(value *bool)() {
    if m != nil {
        m.hasRequiredSetting = value
    }
}
// SetSettingDefinitions sets the settingDefinitions property value. The setting definitions this category contains
func (m *DeviceManagementSettingCategory) SetSettingDefinitions(value []DeviceManagementSettingDefinitionable)() {
    if m != nil {
        m.settingDefinitions = value
    }
}
