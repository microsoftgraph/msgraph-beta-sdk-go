package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsBatteryHealthRuntimeDetails struct {
    Entity
    // Number of active devices within the tenant. Valid values -2147483648 to 2147483647
    activeDevices *int32;
    // Number of devices whose active runtime is greater than 3 hours but lesser than 5 hours. Valid values -2147483648 to 2147483647
    batteryRuntimeFair *int32;
    // Number of devices  whose active runtime is greater than 5 hours. Valid values -2147483648 to 2147483647
    batteryRuntimeGood *int32;
    // Number of devices whose active runtime is lesser than 3 hours. Valid values -2147483648 to 2147483647
    batteryRuntimePoor *int32;
}
// Instantiates a new userExperienceAnalyticsBatteryHealthRuntimeDetails and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthRuntimeDetails()(*UserExperienceAnalyticsBatteryHealthRuntimeDetails) {
    m := &UserExperienceAnalyticsBatteryHealthRuntimeDetails{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activeDevices property value. Number of active devices within the tenant. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) GetActiveDevices()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeDevices
    }
}
// Gets the batteryRuntimeFair property value. Number of devices whose active runtime is greater than 3 hours but lesser than 5 hours. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) GetBatteryRuntimeFair()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryRuntimeFair
    }
}
// Gets the batteryRuntimeGood property value. Number of devices  whose active runtime is greater than 5 hours. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) GetBatteryRuntimeGood()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryRuntimeGood
    }
}
// Gets the batteryRuntimePoor property value. Number of devices whose active runtime is lesser than 3 hours. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) GetBatteryRuntimePoor()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryRuntimePoor
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveDevices(val)
        }
        return nil
    }
    res["batteryRuntimeFair"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryRuntimeFair(val)
        }
        return nil
    }
    res["batteryRuntimeGood"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryRuntimeGood(val)
        }
        return nil
    }
    res["batteryRuntimePoor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryRuntimePoor(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("activeDevices", m.GetActiveDevices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("batteryRuntimeFair", m.GetBatteryRuntimeFair())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("batteryRuntimeGood", m.GetBatteryRuntimeGood())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("batteryRuntimePoor", m.GetBatteryRuntimePoor())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the activeDevices property value. Number of active devices within the tenant. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the activeDevices property.
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) SetActiveDevices(value *int32)() {
    m.activeDevices = value
}
// Sets the batteryRuntimeFair property value. Number of devices whose active runtime is greater than 3 hours but lesser than 5 hours. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the batteryRuntimeFair property.
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) SetBatteryRuntimeFair(value *int32)() {
    m.batteryRuntimeFair = value
}
// Sets the batteryRuntimeGood property value. Number of devices  whose active runtime is greater than 5 hours. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the batteryRuntimeGood property.
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) SetBatteryRuntimeGood(value *int32)() {
    m.batteryRuntimeGood = value
}
// Sets the batteryRuntimePoor property value. Number of devices whose active runtime is lesser than 3 hours. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the batteryRuntimePoor property.
func (m *UserExperienceAnalyticsBatteryHealthRuntimeDetails) SetBatteryRuntimePoor(value *int32)() {
    m.batteryRuntimePoor = value
}
