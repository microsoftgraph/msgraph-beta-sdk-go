package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsBatteryHealthCapacityDetails 
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
    // Recorded date time of this capacity details instance.
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewUserExperienceAnalyticsBatteryHealthCapacityDetails instantiates a new userExperienceAnalyticsBatteryHealthCapacityDetails and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthCapacityDetails()(*UserExperienceAnalyticsBatteryHealthCapacityDetails) {
    m := &UserExperienceAnalyticsBatteryHealthCapacityDetails{
        Entity: *NewEntity(),
    }
    return m
}
// GetActiveDevices gets the activeDevices property value. Number of active devices within the tenant. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) GetActiveDevices()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activeDevices
    }
}
// GetBatteryCapacityFair gets the batteryCapacityFair property value. Number of devices whose battery maximum capacity is greater than 50% but lesser than 80%. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) GetBatteryCapacityFair()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryCapacityFair
    }
}
// GetBatteryCapacityGood gets the batteryCapacityGood property value. Number of devices whose battery maximum capacity is greater than 80%. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) GetBatteryCapacityGood()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryCapacityGood
    }
}
// GetBatteryCapacityPoor gets the batteryCapacityPoor property value. Number of devices whose battery maximum capacity is lesser than 50%. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) GetBatteryCapacityPoor()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.batteryCapacityPoor
    }
}
// GetLastRefreshedDateTime gets the lastRefreshedDateTime property value. Recorded date time of this capacity details instance.
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["batteryCapacityFair"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryCapacityFair(val)
        }
        return nil
    }
    res["batteryCapacityGood"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryCapacityGood(val)
        }
        return nil
    }
    res["batteryCapacityPoor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryCapacityPoor(val)
        }
        return nil
    }
    res["lastRefreshedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRefreshedDateTime(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    {
        err = writer.WriteTimeValue("lastRefreshedDateTime", m.GetLastRefreshedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveDevices sets the activeDevices property value. Number of active devices within the tenant. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) SetActiveDevices(value *int32)() {
    if m != nil {
        m.activeDevices = value
    }
}
// SetBatteryCapacityFair sets the batteryCapacityFair property value. Number of devices whose battery maximum capacity is greater than 50% but lesser than 80%. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) SetBatteryCapacityFair(value *int32)() {
    if m != nil {
        m.batteryCapacityFair = value
    }
}
// SetBatteryCapacityGood sets the batteryCapacityGood property value. Number of devices whose battery maximum capacity is greater than 80%. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) SetBatteryCapacityGood(value *int32)() {
    if m != nil {
        m.batteryCapacityGood = value
    }
}
// SetBatteryCapacityPoor sets the batteryCapacityPoor property value. Number of devices whose battery maximum capacity is lesser than 50%. Valid values -2147483648 to 2147483647
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) SetBatteryCapacityPoor(value *int32)() {
    if m != nil {
        m.batteryCapacityPoor = value
    }
}
// SetLastRefreshedDateTime sets the lastRefreshedDateTime property value. Recorded date time of this capacity details instance.
func (m *UserExperienceAnalyticsBatteryHealthCapacityDetails) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastRefreshedDateTime = value
    }
}
