package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementConfigurationSettingOccurrence 
type DeviceManagementConfigurationSettingOccurrence struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Maximum times setting can be set on device.
    maxDeviceOccurrence *int32;
    // Minimum times setting can be set on device. A MinDeviceOccurrence of 0 means setting is optional
    minDeviceOccurrence *int32;
}
// NewDeviceManagementConfigurationSettingOccurrence instantiates a new deviceManagementConfigurationSettingOccurrence and sets the default values.
func NewDeviceManagementConfigurationSettingOccurrence()(*DeviceManagementConfigurationSettingOccurrence) {
    m := &DeviceManagementConfigurationSettingOccurrence{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationSettingOccurrence) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetMaxDeviceOccurrence gets the maxDeviceOccurrence property value. Maximum times setting can be set on device.
func (m *DeviceManagementConfigurationSettingOccurrence) GetMaxDeviceOccurrence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maxDeviceOccurrence
    }
}
// GetMinDeviceOccurrence gets the minDeviceOccurrence property value. Minimum times setting can be set on device. A MinDeviceOccurrence of 0 means setting is optional
func (m *DeviceManagementConfigurationSettingOccurrence) GetMinDeviceOccurrence()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minDeviceOccurrence
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSettingOccurrence) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["maxDeviceOccurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxDeviceOccurrence(val)
        }
        return nil
    }
    res["minDeviceOccurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinDeviceOccurrence(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementConfigurationSettingOccurrence) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSettingOccurrence) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("maxDeviceOccurrence", m.GetMaxDeviceOccurrence())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minDeviceOccurrence", m.GetMinDeviceOccurrence())
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
func (m *DeviceManagementConfigurationSettingOccurrence) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMaxDeviceOccurrence sets the maxDeviceOccurrence property value. Maximum times setting can be set on device.
func (m *DeviceManagementConfigurationSettingOccurrence) SetMaxDeviceOccurrence(value *int32)() {
    if m != nil {
        m.maxDeviceOccurrence = value
    }
}
// SetMinDeviceOccurrence sets the minDeviceOccurrence property value. Minimum times setting can be set on device. A MinDeviceOccurrence of 0 means setting is optional
func (m *DeviceManagementConfigurationSettingOccurrence) SetMinDeviceOccurrence(value *int32)() {
    if m != nil {
        m.minDeviceOccurrence = value
    }
}
