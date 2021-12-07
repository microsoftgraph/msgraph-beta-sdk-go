package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsCloudIdentityDevicesSummary 
type UserExperienceAnalyticsCloudIdentityDevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The count of devices that are not cloud identity.
    deviceWithoutCloudIdentityCount *int32;
}
// NewUserExperienceAnalyticsCloudIdentityDevicesSummary instantiates a new userExperienceAnalyticsCloudIdentityDevicesSummary and sets the default values.
func NewUserExperienceAnalyticsCloudIdentityDevicesSummary()(*UserExperienceAnalyticsCloudIdentityDevicesSummary) {
    m := &UserExperienceAnalyticsCloudIdentityDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceWithoutCloudIdentityCount gets the deviceWithoutCloudIdentityCount property value. The count of devices that are not cloud identity.
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) GetDeviceWithoutCloudIdentityCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceWithoutCloudIdentityCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceWithoutCloudIdentityCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceWithoutCloudIdentityCount(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("deviceWithoutCloudIdentityCount", m.GetDeviceWithoutCloudIdentityCount())
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
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceWithoutCloudIdentityCount sets the deviceWithoutCloudIdentityCount property value. The count of devices that are not cloud identity.
func (m *UserExperienceAnalyticsCloudIdentityDevicesSummary) SetDeviceWithoutCloudIdentityCount(value *int32)() {
    if m != nil {
        m.deviceWithoutCloudIdentityCount = value
    }
}
