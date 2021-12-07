package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsWindows10DevicesSummary 
type UserExperienceAnalyticsWindows10DevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The count of Windows 10 devices that have unsupported OS versions.
    unsupportedOSversionDeviceCount *int32;
}
// NewUserExperienceAnalyticsWindows10DevicesSummary instantiates a new userExperienceAnalyticsWindows10DevicesSummary and sets the default values.
func NewUserExperienceAnalyticsWindows10DevicesSummary()(*UserExperienceAnalyticsWindows10DevicesSummary) {
    m := &UserExperienceAnalyticsWindows10DevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsWindows10DevicesSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetUnsupportedOSversionDeviceCount gets the unsupportedOSversionDeviceCount property value. The count of Windows 10 devices that have unsupported OS versions.
func (m *UserExperienceAnalyticsWindows10DevicesSummary) GetUnsupportedOSversionDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.unsupportedOSversionDeviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsWindows10DevicesSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["unsupportedOSversionDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnsupportedOSversionDeviceCount(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsWindows10DevicesSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsWindows10DevicesSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("unsupportedOSversionDeviceCount", m.GetUnsupportedOSversionDeviceCount())
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
func (m *UserExperienceAnalyticsWindows10DevicesSummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUnsupportedOSversionDeviceCount sets the unsupportedOSversionDeviceCount property value. The count of Windows 10 devices that have unsupported OS versions.
func (m *UserExperienceAnalyticsWindows10DevicesSummary) SetUnsupportedOSversionDeviceCount(value *int32)() {
    if m != nil {
        m.unsupportedOSversionDeviceCount = value
    }
}
