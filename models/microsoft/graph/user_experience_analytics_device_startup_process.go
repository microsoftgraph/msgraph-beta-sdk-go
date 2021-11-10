package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsDeviceStartupProcess struct {
    Entity
    // The user experience analytics device id.
    managedDeviceId *string;
    // User experience analytics device startup process name.
    processName *string;
    // The user experience analytics device startup process product name.
    productName *string;
    // The User experience analytics device startup process publisher.
    publisher *string;
    // User experience analytics device startup process impact in milliseconds.
    startupImpactInMs *int32;
}
// Instantiates a new userExperienceAnalyticsDeviceStartupProcess and sets the default values.
func NewUserExperienceAnalyticsDeviceStartupProcess()(*UserExperienceAnalyticsDeviceStartupProcess) {
    m := &UserExperienceAnalyticsDeviceStartupProcess{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the managedDeviceId property value. The user experience analytics device id.
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// Gets the processName property value. User experience analytics device startup process name.
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetProcessName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.processName
    }
}
// Gets the productName property value. The user experience analytics device startup process product name.
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetProductName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productName
    }
}
// Gets the publisher property value. The User experience analytics device startup process publisher.
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// Gets the startupImpactInMs property value. User experience analytics device startup process impact in milliseconds.
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetStartupImpactInMs()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.startupImpactInMs
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceStartupProcess) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
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
    res["startupImpactInMs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartupImpactInMs(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsDeviceStartupProcess) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsDeviceStartupProcess) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
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
        err = writer.WriteInt32Value("startupImpactInMs", m.GetStartupImpactInMs())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the managedDeviceId property value. The user experience analytics device id.
// Parameters:
//  - value : Value to set for the managedDeviceId property.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// Sets the processName property value. User experience analytics device startup process name.
// Parameters:
//  - value : Value to set for the processName property.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetProcessName(value *string)() {
    m.processName = value
}
// Sets the productName property value. The user experience analytics device startup process product name.
// Parameters:
//  - value : Value to set for the productName property.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetProductName(value *string)() {
    m.productName = value
}
// Sets the publisher property value. The User experience analytics device startup process publisher.
// Parameters:
//  - value : Value to set for the publisher property.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetPublisher(value *string)() {
    m.publisher = value
}
// Sets the startupImpactInMs property value. User experience analytics device startup process impact in milliseconds.
// Parameters:
//  - value : Value to set for the startupImpactInMs property.
func (m *UserExperienceAnalyticsDeviceStartupProcess) SetStartupImpactInMs(value *int32)() {
    m.startupImpactInMs = value
}
