package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsNotAutopilotReadyDevice struct {
    Entity
    // The intune device's autopilotProfileAssigned.
    autoPilotProfileAssigned *bool;
    // The intune device's autopilotRegistered.
    autoPilotRegistered *bool;
    // The intune device's azure Ad joinType.
    azureAdJoinType *string;
    // The intune device's azureAdRegistered.
    azureAdRegistered *bool;
    // The intune device's name.
    deviceName *string;
    // The intune device's managed by.
    managedBy *string;
    // The intune device's manufacturer.
    manufacturer *string;
    // The intune device's model.
    model *string;
    // The intune device's serial number.
    serialNumber *string;
}
// Instantiates a new userExperienceAnalyticsNotAutopilotReadyDevice and sets the default values.
func NewUserExperienceAnalyticsNotAutopilotReadyDevice()(*UserExperienceAnalyticsNotAutopilotReadyDevice) {
    m := &UserExperienceAnalyticsNotAutopilotReadyDevice{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the autoPilotProfileAssigned property value. The intune device's autopilotProfileAssigned.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAutoPilotProfileAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoPilotProfileAssigned
    }
}
// Gets the autoPilotRegistered property value. The intune device's autopilotRegistered.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAutoPilotRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoPilotRegistered
    }
}
// Gets the azureAdJoinType property value. The intune device's azure Ad joinType.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAzureAdJoinType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureAdJoinType
    }
}
// Gets the azureAdRegistered property value. The intune device's azureAdRegistered.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAzureAdRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.azureAdRegistered
    }
}
// Gets the deviceName property value. The intune device's name.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the managedBy property value. The intune device's managed by.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetManagedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedBy
    }
}
// Gets the manufacturer property value. The intune device's manufacturer.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// Gets the model property value. The intune device's model.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// Gets the serialNumber property value. The intune device's serial number.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["autoPilotProfileAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoPilotProfileAssigned(val)
        }
        return nil
    }
    res["autoPilotRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoPilotRegistered(val)
        }
        return nil
    }
    res["azureAdJoinType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdJoinType(val)
        }
        return nil
    }
    res["azureAdRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdRegistered(val)
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["managedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBy(val)
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("autoPilotProfileAssigned", m.GetAutoPilotProfileAssigned())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("autoPilotRegistered", m.GetAutoPilotRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureAdJoinType", m.GetAzureAdJoinType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("azureAdRegistered", m.GetAzureAdRegistered())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedBy", m.GetManagedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the autoPilotProfileAssigned property value. The intune device's autopilotProfileAssigned.
// Parameters:
//  - value : Value to set for the autoPilotProfileAssigned property.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAutoPilotProfileAssigned(value *bool)() {
    m.autoPilotProfileAssigned = value
}
// Sets the autoPilotRegistered property value. The intune device's autopilotRegistered.
// Parameters:
//  - value : Value to set for the autoPilotRegistered property.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAutoPilotRegistered(value *bool)() {
    m.autoPilotRegistered = value
}
// Sets the azureAdJoinType property value. The intune device's azure Ad joinType.
// Parameters:
//  - value : Value to set for the azureAdJoinType property.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAzureAdJoinType(value *string)() {
    m.azureAdJoinType = value
}
// Sets the azureAdRegistered property value. The intune device's azureAdRegistered.
// Parameters:
//  - value : Value to set for the azureAdRegistered property.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAzureAdRegistered(value *bool)() {
    m.azureAdRegistered = value
}
// Sets the deviceName property value. The intune device's name.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the managedBy property value. The intune device's managed by.
// Parameters:
//  - value : Value to set for the managedBy property.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetManagedBy(value *string)() {
    m.managedBy = value
}
// Sets the manufacturer property value. The intune device's manufacturer.
// Parameters:
//  - value : Value to set for the manufacturer property.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// Sets the model property value. The intune device's model.
// Parameters:
//  - value : Value to set for the model property.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetModel(value *string)() {
    m.model = value
}
// Sets the serialNumber property value. The intune device's serial number.
// Parameters:
//  - value : Value to set for the serialNumber property.
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
