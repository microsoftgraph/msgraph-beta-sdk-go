package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserExperienceAnalyticsDeviceWithoutCloudIdentity struct {
    Entity
    // Azure Active Directory Device Id
    azureAdDeviceId *string;
    // The tenant attach device's name.
    deviceName *string;
}
// Instantiates a new userExperienceAnalyticsDeviceWithoutCloudIdentity and sets the default values.
func NewUserExperienceAnalyticsDeviceWithoutCloudIdentity()(*UserExperienceAnalyticsDeviceWithoutCloudIdentity) {
    m := &UserExperienceAnalyticsDeviceWithoutCloudIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the azureAdDeviceId property value. Azure Active Directory Device Id
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) GetAzureAdDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureAdDeviceId
    }
}
// Gets the deviceName property value. The tenant attach device's name.
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// The deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureAdDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureAdDeviceId(val)
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
    return res
}
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureAdDeviceId", m.GetAzureAdDeviceId())
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
    return nil
}
// Sets the azureAdDeviceId property value. Azure Active Directory Device Id
// Parameters:
//  - value : Value to set for the azureAdDeviceId property.
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) SetAzureAdDeviceId(value *string)() {
    m.azureAdDeviceId = value
}
// Sets the deviceName property value. The tenant attach device's name.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) SetDeviceName(value *string)() {
    m.deviceName = value
}
