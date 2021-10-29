package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsDeviceStartupProcessPerformance struct {
    Entity
    // User experience analytics device startup process summarized count.
    deviceCount *int64;
    // User experience analytics device startup process median impact in milliseconds.
    medianImpactInMs *int32;
    // User experience analytics device startup process median impact in milliseconds.
    medianImpactInMs2 *int64;
    // User experience analytics device startup process name.
    processName *string;
    // The user experience analytics device startup process product name.
    productName *string;
    // The User experience analytics device startup process publisher.
    publisher *string;
    // User experience analytics device startup process total impact in milliseconds.
    totalImpactInMs *int32;
    // User experience analytics device startup process total impact in milliseconds.
    totalImpactInMs2 *int64;
}
// Instantiates a new userExperienceAnalyticsDeviceStartupProcessPerformance and sets the default values.
func NewUserExperienceAnalyticsDeviceStartupProcessPerformance()(*UserExperienceAnalyticsDeviceStartupProcessPerformance) {
    m := &UserExperienceAnalyticsDeviceStartupProcessPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceCount property value. User experience analytics device startup process summarized count.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetDeviceCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
// Gets the medianImpactInMs property value. User experience analytics device startup process median impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetMedianImpactInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.medianImpactInMs
    }
}
// Gets the medianImpactInMs2 property value. User experience analytics device startup process median impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetMedianImpactInMs2()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.medianImpactInMs2
    }
}
// Gets the processName property value. User experience analytics device startup process name.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetProcessName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.processName
    }
}
// Gets the productName property value. The user experience analytics device startup process product name.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetProductName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productName
    }
}
// Gets the publisher property value. The User experience analytics device startup process publisher.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// Gets the totalImpactInMs property value. User experience analytics device startup process total impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetTotalImpactInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalImpactInMs
    }
}
// Gets the totalImpactInMs2 property value. User experience analytics device startup process total impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetTotalImpactInMs2()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalImpactInMs2
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDeviceCount(val)
        return nil
    }
    res["medianImpactInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMedianImpactInMs(val)
        return nil
    }
    res["medianImpactInMs2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetMedianImpactInMs2(val)
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
    res["productName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProductName(val)
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
    res["totalImpactInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotalImpactInMs(val)
        return nil
    }
    res["totalImpactInMs2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetTotalImpactInMs2(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("deviceCount", m.GetDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("medianImpactInMs", m.GetMedianImpactInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("medianImpactInMs2", m.GetMedianImpactInMs2())
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
        err = writer.WriteStringValue("productName", m.GetProductName())
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
    {
        err = writer.WriteInt32Value("totalImpactInMs", m.GetTotalImpactInMs())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("totalImpactInMs2", m.GetTotalImpactInMs2())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deviceCount property value. User experience analytics device startup process summarized count.
// Parameters:
//  - value : Value to set for the deviceCount property.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetDeviceCount(value *int64)() {
    m.deviceCount = value
}
// Sets the medianImpactInMs property value. User experience analytics device startup process median impact in milliseconds.
// Parameters:
//  - value : Value to set for the medianImpactInMs property.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetMedianImpactInMs(value *int32)() {
    m.medianImpactInMs = value
}
// Sets the medianImpactInMs2 property value. User experience analytics device startup process median impact in milliseconds.
// Parameters:
//  - value : Value to set for the medianImpactInMs2 property.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetMedianImpactInMs2(value *int64)() {
    m.medianImpactInMs2 = value
}
// Sets the processName property value. User experience analytics device startup process name.
// Parameters:
//  - value : Value to set for the processName property.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetProcessName(value *string)() {
    m.processName = value
}
// Sets the productName property value. The user experience analytics device startup process product name.
// Parameters:
//  - value : Value to set for the productName property.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetProductName(value *string)() {
    m.productName = value
}
// Sets the publisher property value. The User experience analytics device startup process publisher.
// Parameters:
//  - value : Value to set for the publisher property.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetPublisher(value *string)() {
    m.publisher = value
}
// Sets the totalImpactInMs property value. User experience analytics device startup process total impact in milliseconds.
// Parameters:
//  - value : Value to set for the totalImpactInMs property.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetTotalImpactInMs(value *int32)() {
    m.totalImpactInMs = value
}
// Sets the totalImpactInMs2 property value. User experience analytics device startup process total impact in milliseconds.
// Parameters:
//  - value : Value to set for the totalImpactInMs2 property.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetTotalImpactInMs2(value *int64)() {
    m.totalImpactInMs2 = value
}
