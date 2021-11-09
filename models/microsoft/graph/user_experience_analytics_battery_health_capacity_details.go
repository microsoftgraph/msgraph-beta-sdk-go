package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsBatteryHealthCapacityDetails struct {
    Entity
    // Number of active devices within the tenant. Valid values -2147483648 to 2147483647
    activeDevices *int32;
    // Number of devices whose battery maximum capacity is greater than 50% but lesser than 80%. Valid values -2147483648 to 2147483647
    batteryCapacityFair *int32;
    // Number of devices whose battery maximum capacity is greater than 80%. Valid values -2147483648 to 2147483647
    batteryCapacityGood *int32;
    // Number of devices whose battery maximum capacity is lesser than 50%. Valid values -2147483648 to 2147483647
    batteryCapacityPoor *int32;
}
// Instantiates a new userExperienceAnalyticsBatteryHealthCapacityDetails and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthCapacityDetails()(*UserExperienceAnalyticsBatteryHealthCapacityDetails) {
    m := &UserExperienceAnalyticsBatteryHealthCapacityDetails{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activeDevices property value. Number of active devices within the tenant. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) GetActiveDevices()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeDevices
    }
}
// Gets the batteryCapacityFair property value. Number of devices whose battery maximum capacity is greater than 50% but lesser than 80%. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) GetBatteryCapacityFair()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryCapacityFair
    }
}
// Gets the batteryCapacityGood property value. Number of devices whose battery maximum capacity is greater than 80%. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) GetBatteryCapacityGood()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryCapacityGood
    }
}
// Gets the batteryCapacityPoor property value. Number of devices whose battery maximum capacity is lesser than 50%. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) GetBatteryCapacityPoor()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryCapacityPoor
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetActiveDevices(val)
        return nil
    }
    res["batteryCapacityFair"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetBatteryCapacityFair(val)
        return nil
    }
    res["batteryCapacityGood"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetBatteryCapacityGood(val)
        return nil
    }
    res["batteryCapacityPoor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetBatteryCapacityPoor(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("batteryCapacityFair", m.GetBatteryCapacityFair())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("batteryCapacityGood", m.GetBatteryCapacityGood())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("batteryCapacityPoor", m.GetBatteryCapacityPoor())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the activeDevices property value. Number of active devices within the tenant. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the activeDevices property.
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) SetActiveDevices(value *int32)() {
    m.activeDevices = value
}
// Sets the batteryCapacityFair property value. Number of devices whose battery maximum capacity is greater than 50% but lesser than 80%. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the batteryCapacityFair property.
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) SetBatteryCapacityFair(value *int32)() {
    m.batteryCapacityFair = value
}
// Sets the batteryCapacityGood property value. Number of devices whose battery maximum capacity is greater than 80%. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the batteryCapacityGood property.
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) SetBatteryCapacityGood(value *int32)() {
    m.batteryCapacityGood = value
}
// Sets the batteryCapacityPoor property value. Number of devices whose battery maximum capacity is lesser than 50%. Valid values -2147483648 to 2147483647
// Parameters:
//  - value : Value to set for the batteryCapacityPoor property.
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) SetBatteryCapacityPoor(value *int32)() {
    m.batteryCapacityPoor = value
}
