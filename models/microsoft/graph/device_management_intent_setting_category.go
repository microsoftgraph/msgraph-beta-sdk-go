package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementIntentSettingCategory provides operations to manage the deviceManagement singleton.
type DeviceManagementIntentSettingCategory struct {
    DeviceManagementSettingCategory
    // The settings this category contains
    settings []DeviceManagementSettingInstanceable;
}
// NewDeviceManagementIntentSettingCategory instantiates a new deviceManagementIntentSettingCategory and sets the default values.
func NewDeviceManagementIntentSettingCategory()(*DeviceManagementIntentSettingCategory) {
    m := &DeviceManagementIntentSettingCategory{
        DeviceManagementSettingCategory: *NewDeviceManagementSettingCategory(),
    }
    return m
}
// CreateDeviceManagementIntentSettingCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementIntentSettingCategoryFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDeviceManagementIntentSettingCategory(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntentSettingCategory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.DeviceManagementSettingCategory.GetFieldDeserializers()
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInstanceable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementSettingInstanceable)
            }
            m.SetSettings(res)
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. The settings this category contains
func (m *DeviceManagementIntentSettingCategory) GetSettings()([]DeviceManagementSettingInstanceable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *DeviceManagementIntentSettingCategory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementIntentSettingCategory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.DeviceManagementSettingCategory.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSettings()))
        for i, v := range m.GetSettings() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("settings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSettings sets the settings property value. The settings this category contains
func (m *DeviceManagementIntentSettingCategory) SetSettings(value []DeviceManagementSettingInstanceable)() {
    if m != nil {
        m.settings = value
    }
}
