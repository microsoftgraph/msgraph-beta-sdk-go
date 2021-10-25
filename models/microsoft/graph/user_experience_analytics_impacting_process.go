package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsImpactingProcess struct {
    Entity
    category *string;
    description *string;
    deviceId *string;
    impactValue *float64;
    processName *string;
    publisher *string;
}
func NewUserExperienceAnalyticsImpactingProcess()(*UserExperienceAnalyticsImpactingProcess) {
    m := &UserExperienceAnalyticsImpactingProcess{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserExperienceAnalyticsImpactingProcess) GetCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
func (m *UserExperienceAnalyticsImpactingProcess) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *UserExperienceAnalyticsImpactingProcess) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
func (m *UserExperienceAnalyticsImpactingProcess) GetImpactValue()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.impactValue
    }
}
func (m *UserExperienceAnalyticsImpactingProcess) GetProcessName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.processName
    }
}
func (m *UserExperienceAnalyticsImpactingProcess) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
func (m *UserExperienceAnalyticsImpactingProcess) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
        return nil
    }
    res["impactValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetImpactValue(val)
        return nil
    }
    res["processName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProcessName(val)
        return nil
    }
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPublisher(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsImpactingProcess) IsNil()(bool) {
    return m == nil
}
func (m *UserExperienceAnalyticsImpactingProcess) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("impactValue", m.GetImpactValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("processName", m.GetProcessName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UserExperienceAnalyticsImpactingProcess) SetCategory(value *string)() {
    m.category = value
}
func (m *UserExperienceAnalyticsImpactingProcess) SetDescription(value *string)() {
    m.description = value
}
func (m *UserExperienceAnalyticsImpactingProcess) SetDeviceId(value *string)() {
    m.deviceId = value
}
func (m *UserExperienceAnalyticsImpactingProcess) SetImpactValue(value *float64)() {
    m.impactValue = value
}
func (m *UserExperienceAnalyticsImpactingProcess) SetProcessName(value *string)() {
    m.processName = value
}
func (m *UserExperienceAnalyticsImpactingProcess) SetPublisher(value *string)() {
    m.publisher = value
}
