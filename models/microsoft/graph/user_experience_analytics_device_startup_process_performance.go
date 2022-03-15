package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsDeviceStartupProcessPerformance provides operations to manage the deviceManagement singleton.
type UserExperienceAnalyticsDeviceStartupProcessPerformance struct {
    Entity
    // User experience analytics device startup process summarized count.
    deviceCount *int32;
    // User experience analytics device startup process median impact in milliseconds.
    medianImpactInMs *int32;
    // User experience analytics device startup process median impact in milliseconds.
    medianImpactInMs2 *int32;
    // User experience analytics device startup process name.
    processName *string;
    // The user experience analytics device startup process product name.
    productName *string;
    // The User experience analytics device startup process publisher.
    publisher *string;
    // User experience analytics device startup process total impact in milliseconds.
    totalImpactInMs *int32;
    // User experience analytics device startup process total impact in milliseconds.
    totalImpactInMs2 *int32;
}
// NewUserExperienceAnalyticsDeviceStartupProcessPerformance instantiates a new userExperienceAnalyticsDeviceStartupProcessPerformance and sets the default values.
func NewUserExperienceAnalyticsDeviceStartupProcessPerformance()(*UserExperienceAnalyticsDeviceStartupProcessPerformance) {
    m := &UserExperienceAnalyticsDeviceStartupProcessPerformance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDeviceStartupProcessPerformanceFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceStartupProcessPerformance(), nil
}
// GetDeviceCount gets the deviceCount property value. User experience analytics device startup process summarized count.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    res["medianImpactInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedianImpactInMs(val)
        }
        return nil
    }
    res["medianImpactInMs2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMedianImpactInMs2(val)
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
    res["productName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductName(val)
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
    res["totalImpactInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalImpactInMs(val)
        }
        return nil
    }
    res["totalImpactInMs2"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalImpactInMs2(val)
        }
        return nil
    }
    return res
}
// GetMedianImpactInMs gets the medianImpactInMs property value. User experience analytics device startup process median impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetMedianImpactInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.medianImpactInMs
    }
}
// GetMedianImpactInMs2 gets the medianImpactInMs2 property value. User experience analytics device startup process median impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetMedianImpactInMs2()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.medianImpactInMs2
    }
}
// GetProcessName gets the processName property value. User experience analytics device startup process name.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetProcessName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.processName
    }
}
// GetProductName gets the productName property value. The user experience analytics device startup process product name.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetProductName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productName
    }
}
// GetPublisher gets the publisher property value. The User experience analytics device startup process publisher.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// GetTotalImpactInMs gets the totalImpactInMs property value. User experience analytics device startup process total impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetTotalImpactInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalImpactInMs
    }
}
// GetTotalImpactInMs2 gets the totalImpactInMs2 property value. User experience analytics device startup process total impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) GetTotalImpactInMs2()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalImpactInMs2
    }
}
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
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
        err = writer.WriteInt32Value("medianImpactInMs2", m.GetMedianImpactInMs2())
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
        err = writer.WriteInt32Value("totalImpactInMs2", m.GetTotalImpactInMs2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceCount sets the deviceCount property value. User experience analytics device startup process summarized count.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetDeviceCount(value *int32)() {
    if m != nil {
        m.deviceCount = value
    }
}
// SetMedianImpactInMs sets the medianImpactInMs property value. User experience analytics device startup process median impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetMedianImpactInMs(value *int32)() {
    if m != nil {
        m.medianImpactInMs = value
    }
}
// SetMedianImpactInMs2 sets the medianImpactInMs2 property value. User experience analytics device startup process median impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetMedianImpactInMs2(value *int32)() {
    if m != nil {
        m.medianImpactInMs2 = value
    }
}
// SetProcessName sets the processName property value. User experience analytics device startup process name.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetProcessName(value *string)() {
    if m != nil {
        m.processName = value
    }
}
// SetProductName sets the productName property value. The user experience analytics device startup process product name.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetProductName(value *string)() {
    if m != nil {
        m.productName = value
    }
}
// SetPublisher sets the publisher property value. The User experience analytics device startup process publisher.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetPublisher(value *string)() {
    if m != nil {
        m.publisher = value
    }
}
// SetTotalImpactInMs sets the totalImpactInMs property value. User experience analytics device startup process total impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetTotalImpactInMs(value *int32)() {
    if m != nil {
        m.totalImpactInMs = value
    }
}
// SetTotalImpactInMs2 sets the totalImpactInMs2 property value. User experience analytics device startup process total impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcessPerformance) SetTotalImpactInMs2(value *int32)() {
    if m != nil {
        m.totalImpactInMs2 = value
    }
}
