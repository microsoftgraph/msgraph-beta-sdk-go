package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagementIntentSettingCategory struct {
    DeviceManagementSettingCategory
    // The settings this category contains
    settings []DeviceManagementSettingInstance;
}
// Instantiates a new deviceManagementIntentSettingCategory and sets the default values.
func NewDeviceManagementIntentSettingCategory()(*DeviceManagementIntentSettingCategory) {
    m := &DeviceManagementIntentSettingCategory{
        DeviceManagementSettingCategory: *NewDeviceManagementSettingCategory(),
    }
    return m
}
// Gets the settings property value. The settings this category contains
func (m *DeviceManagementIntentSettingCategory) GetSettings()([]DeviceManagementSettingInstance) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// The deserialization information for the current model
func (m *DeviceManagementIntentSettingCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DeviceManagementSettingCategory.GetFieldDeserializers()
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettingInstance() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInstance, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementSettingInstance))
            }
            m.SetSettings(res)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementIntentSettingCategory) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementIntentSettingCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DeviceManagementSettingCategory.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the settings property value. The settings this category contains
// Parameters:
//  - value : Value to set for the settings property.
func (m *DeviceManagementIntentSettingCategory) SetSettings(value []DeviceManagementSettingInstance)() {
    m.settings = value
}
