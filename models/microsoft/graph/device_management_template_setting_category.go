package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementTemplateSettingCategory 
type DeviceManagementTemplateSettingCategory struct {
    DeviceManagementSettingCategory
    // The settings this category contains
    recommendedSettings []DeviceManagementSettingInstance;
}
// NewDeviceManagementTemplateSettingCategory instantiates a new deviceManagementTemplateSettingCategory and sets the default values.
func NewDeviceManagementTemplateSettingCategory()(*DeviceManagementTemplateSettingCategory) {
    m := &DeviceManagementTemplateSettingCategory{
        DeviceManagementSettingCategory: *NewDeviceManagementSettingCategory(),
    }
    return m
}
// GetRecommendedSettings gets the recommendedSettings property value. The settings this category contains
func (m *DeviceManagementTemplateSettingCategory) GetRecommendedSettings()([]DeviceManagementSettingInstance) {
    if m == nil {
        return nil
    } else {
        return m.recommendedSettings
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementTemplateSettingCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DeviceManagementSettingCategory.GetFieldDeserializers()
    res["recommendedSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettingInstance() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInstance, len(val))
            for i, v := range val {
                res[i] = *(v.(*DeviceManagementSettingInstance))
            }
            m.SetRecommendedSettings(res)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementTemplateSettingCategory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementTemplateSettingCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DeviceManagementSettingCategory.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRecommendedSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRecommendedSettings()))
        for i, v := range m.GetRecommendedSettings() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("recommendedSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRecommendedSettings sets the recommendedSettings property value. The settings this category contains
func (m *DeviceManagementTemplateSettingCategory) SetRecommendedSettings(value []DeviceManagementSettingInstance)() {
    if m != nil {
        m.recommendedSettings = value
    }
}
