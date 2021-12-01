package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsDeviceWithoutCloudIdentity 
type UserExperienceAnalyticsDeviceWithoutCloudIdentity struct {
    Entity
    // Azure Active Directory Device Id
    azureAdDeviceId *string;
    // The tenant attach device's name.
    deviceName *string;
}
// NewUserExperienceAnalyticsDeviceWithoutCloudIdentity instantiates a new userExperienceAnalyticsDeviceWithoutCloudIdentity and sets the default values.
func NewUserExperienceAnalyticsDeviceWithoutCloudIdentity()(*UserExperienceAnalyticsDeviceWithoutCloudIdentity) {
    m := &UserExperienceAnalyticsDeviceWithoutCloudIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// GetAzureAdDeviceId gets the azureAdDeviceId property value. Azure Active Directory Device Id
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) GetAzureAdDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureAdDeviceId
    }
}
// GetDeviceName gets the deviceName property value. The tenant attach device's name.
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureAdDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAzureAdDeviceId sets the azureAdDeviceId property value. Azure Active Directory Device Id
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) SetAzureAdDeviceId(value *string)() {
    if m != nil {
        m.azureAdDeviceId = value
    }
}
// SetDeviceName sets the deviceName property value. The tenant attach device's name.
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
