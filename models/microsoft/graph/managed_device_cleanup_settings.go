package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedDeviceCleanupSettings 
type ManagedDeviceCleanupSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Number of days when the device has not contacted Intune.
    deviceInactivityBeforeRetirementInDays *string;
}
// NewManagedDeviceCleanupSettings instantiates a new managedDeviceCleanupSettings and sets the default values.
func NewManagedDeviceCleanupSettings()(*ManagedDeviceCleanupSettings) {
    m := &ManagedDeviceCleanupSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDeviceCleanupSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDeviceInactivityBeforeRetirementInDays gets the deviceInactivityBeforeRetirementInDays property value. Number of days when the device has not contacted Intune.
func (m *ManagedDeviceCleanupSettings) GetDeviceInactivityBeforeRetirementInDays()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceInactivityBeforeRetirementInDays
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceCleanupSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceInactivityBeforeRetirementInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceInactivityBeforeRetirementInDays(val)
        }
        return nil
    }
    return res
}
func (m *ManagedDeviceCleanupSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceCleanupSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceInactivityBeforeRetirementInDays", m.GetDeviceInactivityBeforeRetirementInDays())
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
func (m *ManagedDeviceCleanupSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDeviceInactivityBeforeRetirementInDays sets the deviceInactivityBeforeRetirementInDays property value. Number of days when the device has not contacted Intune.
func (m *ManagedDeviceCleanupSettings) SetDeviceInactivityBeforeRetirementInDays(value *string)() {
    if m != nil {
        m.deviceInactivityBeforeRetirementInDays = value
    }
}
