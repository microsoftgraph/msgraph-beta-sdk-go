package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedDeviceCleanupSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Number of days when the device has not contacted Intune.
    deviceInactivityBeforeRetirementInDays *string;
}
// Instantiates a new managedDeviceCleanupSettings and sets the default values.
func NewManagedDeviceCleanupSettings()(*ManagedDeviceCleanupSettings) {
    m := &ManagedDeviceCleanupSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDeviceCleanupSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the deviceInactivityBeforeRetirementInDays property value. Number of days when the device has not contacted Intune.
func (m *ManagedDeviceCleanupSettings) GetDeviceInactivityBeforeRetirementInDays()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceInactivityBeforeRetirementInDays
    }
}
// The deserialization information for the current model
func (m *ManagedDeviceCleanupSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["deviceInactivityBeforeRetirementInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceInactivityBeforeRetirementInDays(val)
        return nil
    }
    return res
}
func (m *ManagedDeviceCleanupSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ManagedDeviceCleanupSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the deviceInactivityBeforeRetirementInDays property value. Number of days when the device has not contacted Intune.
// Parameters:
//  - value : Value to set for the deviceInactivityBeforeRetirementInDays property.
func (m *ManagedDeviceCleanupSettings) SetDeviceInactivityBeforeRetirementInDays(value *string)() {
    m.deviceInactivityBeforeRetirementInDays = value
}
