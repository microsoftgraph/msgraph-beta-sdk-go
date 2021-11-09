package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new userExperienceAnalyticsImpactingProcess and sets the default values.
func NewUserExperienceAnalyticsImpactingProcess()(*UserExperienceAnalyticsImpactingProcess) {
    m := &UserExperienceAnalyticsImpactingProcess{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the category property value. The category of impacting process.
func (m *UserExperienceAnalyticsImpactingProcess) GetCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// Gets the description property value. The description of process.
func (m *UserExperienceAnalyticsImpactingProcess) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the deviceId property value. The unique identifier of the impacted device.
func (m *UserExperienceAnalyticsImpactingProcess) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the impactValue property value. The impact value of the process. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsImpactingProcess) GetImpactValue()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.impactValue
    }
}
// Gets the processName property value. The process name.
func (m *UserExperienceAnalyticsImpactingProcess) GetProcessName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.processName
    }
}
// Gets the publisher property value. The publisher of the process.
func (m *UserExperienceAnalyticsImpactingProcess) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the category property value. The category of impacting process.
// Parameters:
//  - value : Value to set for the category property.
func (m *UserExperienceAnalyticsImpactingProcess) SetCategory(value *string)() {
    m.category = value
}
// Sets the description property value. The description of process.
// Parameters:
//  - value : Value to set for the description property.
func (m *UserExperienceAnalyticsImpactingProcess) SetDescription(value *string)() {
    m.description = value
}
// Sets the deviceId property value. The unique identifier of the impacted device.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *UserExperienceAnalyticsImpactingProcess) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the impactValue property value. The impact value of the process. Valid values 0 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the impactValue property.
func (m *UserExperienceAnalyticsImpactingProcess) SetImpactValue(value *float64)() {
    m.impactValue = value
}
// Sets the processName property value. The process name.
// Parameters:
//  - value : Value to set for the processName property.
func (m *UserExperienceAnalyticsImpactingProcess) SetProcessName(value *string)() {
    m.processName = value
}
// Sets the publisher property value. The publisher of the process.
// Parameters:
//  - value : Value to set for the publisher property.
func (m *UserExperienceAnalyticsImpactingProcess) SetPublisher(value *string)() {
    m.publisher = value
}
