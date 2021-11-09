package setandroiddeviceownerfullymanagedenrollmentstate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    enabled *bool;
}
// Instantiates a new setAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody and sets the default values.
func NewSetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody()(*SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) {
    m := &SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the enabled property value. 
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// The deserialization information for the current model
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    return res
}
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the enabled property value. 
// Parameters:
//  - value : Value to set for the enabled property.
func (m *SetAndroidDeviceOwnerFullyManagedEnrollmentStateRequestBody) SetEnabled(value *bool)() {
    m.enabled = value
}
