package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsNotAutopilotReadyDevice struct {
    Entity
    autoPilotProfileAssigned *bool;
    autoPilotRegistered *bool;
    azureAdJoinType *string;
    azureAdRegistered *bool;
    deviceName *string;
    managedBy *string;
    manufacturer *string;
    model *string;
    serialNumber *string;
}
func NewUserExperienceAnalyticsNotAutopilotReadyDevice()(*UserExperienceAnalyticsNotAutopilotReadyDevice) {
    m := &UserExperienceAnalyticsNotAutopilotReadyDevice{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAutoPilotProfileAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoPilotProfileAssigned
    }
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAutoPilotRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoPilotRegistered
    }
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAzureAdJoinType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureAdJoinType
    }
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetAzureAdRegistered()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.azureAdRegistered
    }
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetManagedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedBy
    }
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["autoPilotProfileAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAutoPilotProfileAssigned(val)
        return nil
    }
    res["autoPilotRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAutoPilotRegistered(val)
        return nil
    }
    res["azureAdJoinType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureAdJoinType(val)
        return nil
    }
    res["azureAdRegistered"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAzureAdRegistered(val)
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceName(val)
        return nil
    }
    res["managedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedBy(val)
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManufacturer(val)
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModel(val)
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSerialNumber(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) IsNil()(bool) {
    return m == nil
}
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
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAutoPilotProfileAssigned(value *bool)() {
    m.autoPilotProfileAssigned = value
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAutoPilotRegistered(value *bool)() {
    m.autoPilotRegistered = value
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAzureAdJoinType(value *string)() {
    m.azureAdJoinType = value
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetAzureAdRegistered(value *bool)() {
    m.azureAdRegistered = value
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetDeviceName(value *string)() {
    m.deviceName = value
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetManagedBy(value *string)() {
    m.managedBy = value
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetManufacturer(value *string)() {
    m.manufacturer = value
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetModel(value *string)() {
    m.model = value
}
func (m *UserExperienceAnalyticsNotAutopilotReadyDevice) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
