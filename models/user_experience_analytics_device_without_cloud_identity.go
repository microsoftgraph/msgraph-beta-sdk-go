package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceWithoutCloudIdentity 
type UserExperienceAnalyticsDeviceWithoutCloudIdentity struct {
    Entity
}
// NewUserExperienceAnalyticsDeviceWithoutCloudIdentity instantiates a new UserExperienceAnalyticsDeviceWithoutCloudIdentity and sets the default values.
func NewUserExperienceAnalyticsDeviceWithoutCloudIdentity()(*UserExperienceAnalyticsDeviceWithoutCloudIdentity) {
    m := &UserExperienceAnalyticsDeviceWithoutCloudIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsDeviceWithoutCloudIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDeviceWithoutCloudIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceWithoutCloudIdentity(), nil
}
// GetAzureAdDeviceId gets the azureAdDeviceId property value. Azure Active Directory Device Id
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) GetAzureAdDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("azureAdDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The tenant attach device's name.
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["azureAdDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    err := m.GetBackingStore().Set("azureAdDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The tenant attach device's name.
func (m *UserExperienceAnalyticsDeviceWithoutCloudIdentity) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsDeviceWithoutCloudIdentityable 
type UserExperienceAnalyticsDeviceWithoutCloudIdentityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureAdDeviceId()(*string)
    GetDeviceName()(*string)
    SetAzureAdDeviceId(value *string)()
    SetDeviceName(value *string)()
}
