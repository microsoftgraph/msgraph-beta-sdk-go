package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsImpactingProcess 
type UserExperienceAnalyticsImpactingProcess struct {
    Entity
    // The category of impacting process.
    category *string;
    // The description of process.
    description *string;
    // The unique identifier of the impacted device.
    deviceId *string;
    // The impact value of the process. Valid values 0 to 1.79769313486232E+308
    impactValue *float64;
    // The process name.
    processName *string;
    // The publisher of the process.
    publisher *string;
}
// NewUserExperienceAnalyticsImpactingProcess instantiates a new userExperienceAnalyticsImpactingProcess and sets the default values.
func NewUserExperienceAnalyticsImpactingProcess()(*UserExperienceAnalyticsImpactingProcess) {
    m := &UserExperienceAnalyticsImpactingProcess{
        Entity: *NewEntity(),
    }
    return m
}
// GetCategory gets the category property value. The category of impacting process.
func (m *UserExperienceAnalyticsImpactingProcess) GetCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetDescription gets the description property value. The description of process.
func (m *UserExperienceAnalyticsImpactingProcess) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDeviceId gets the deviceId property value. The unique identifier of the impacted device.
func (m *UserExperienceAnalyticsImpactingProcess) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetImpactValue gets the impactValue property value. The impact value of the process. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsImpactingProcess) GetImpactValue()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.impactValue
    }
}
// GetProcessName gets the processName property value. The process name.
func (m *UserExperienceAnalyticsImpactingProcess) GetProcessName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.processName
    }
}
// GetPublisher gets the publisher property value. The publisher of the process.
func (m *UserExperienceAnalyticsImpactingProcess) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsImpactingProcess) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["impactValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImpactValue(val)
        }
        return nil
    }
    res["processName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessName(val)
        }
        return nil
    }
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsImpactingProcess) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetCategory sets the category property value. The category of impacting process.
func (m *UserExperienceAnalyticsImpactingProcess) SetCategory(value *string)() {
    if m != nil {
        m.category = value
    }
}
// SetDescription sets the description property value. The description of process.
func (m *UserExperienceAnalyticsImpactingProcess) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDeviceId sets the deviceId property value. The unique identifier of the impacted device.
func (m *UserExperienceAnalyticsImpactingProcess) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetImpactValue sets the impactValue property value. The impact value of the process. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsImpactingProcess) SetImpactValue(value *float64)() {
    if m != nil {
        m.impactValue = value
    }
}
// SetProcessName sets the processName property value. The process name.
func (m *UserExperienceAnalyticsImpactingProcess) SetProcessName(value *string)() {
    if m != nil {
        m.processName = value
    }
}
// SetPublisher sets the publisher property value. The publisher of the process.
func (m *UserExperienceAnalyticsImpactingProcess) SetPublisher(value *string)() {
    if m != nil {
        m.publisher = value
    }
}
