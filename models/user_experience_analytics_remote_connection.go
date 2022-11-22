package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsRemoteConnection the user experience analyte remote connection entity.
type UserExperienceAnalyticsRemoteConnection struct {
    Entity
    // The sign in failure percentage of Cloud PC Device. Valid values 0 to 100
    cloudPcFailurePercentage *float64
    // The round tip time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
    cloudPcRoundTripTime *float64
    // The sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
    cloudPcSignInTime *float64
    // The core boot time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
    coreBootTime *float64
    // The core sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
    coreSignInTime *float64
    // The count of remote connection. Valid values 0 to 2147483647
    deviceCount *int32
    // The id of the device.
    deviceId *string
    // The name of the device.
    deviceName *string
    // The user experience analytics manufacturer.
    manufacturer *string
    // The user experience analytics device model.
    model *string
    // The remote sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
    remoteSignInTime *float64
    // The user experience analytics userPrincipalName.
    userPrincipalName *string
    // The user experience analytics virtual network.
    virtualNetwork *string
}
// NewUserExperienceAnalyticsRemoteConnection instantiates a new userExperienceAnalyticsRemoteConnection and sets the default values.
func NewUserExperienceAnalyticsRemoteConnection()(*UserExperienceAnalyticsRemoteConnection) {
    m := &UserExperienceAnalyticsRemoteConnection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsRemoteConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsRemoteConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsRemoteConnection(), nil
}
// GetCloudPcFailurePercentage gets the cloudPcFailurePercentage property value. The sign in failure percentage of Cloud PC Device. Valid values 0 to 100
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcFailurePercentage()(*float64) {
    return m.cloudPcFailurePercentage
}
// GetCloudPcRoundTripTime gets the cloudPcRoundTripTime property value. The round tip time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcRoundTripTime()(*float64) {
    return m.cloudPcRoundTripTime
}
// GetCloudPcSignInTime gets the cloudPcSignInTime property value. The sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcSignInTime()(*float64) {
    return m.cloudPcSignInTime
}
// GetCoreBootTime gets the coreBootTime property value. The core boot time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCoreBootTime()(*float64) {
    return m.coreBootTime
}
// GetCoreSignInTime gets the coreSignInTime property value. The core sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCoreSignInTime()(*float64) {
    return m.coreSignInTime
}
// GetDeviceCount gets the deviceCount property value. The count of remote connection. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceCount()(*int32) {
    return m.deviceCount
}
// GetDeviceId gets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceId()(*string) {
    return m.deviceId
}
// GetDeviceName gets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsRemoteConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cloudPcFailurePercentage"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetCloudPcFailurePercentage)
    res["cloudPcRoundTripTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetCloudPcRoundTripTime)
    res["cloudPcSignInTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetCloudPcSignInTime)
    res["coreBootTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetCoreBootTime)
    res["coreSignInTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetCoreSignInTime)
    res["deviceCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDeviceCount)
    res["deviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceId)
    res["deviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceName)
    res["manufacturer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManufacturer)
    res["model"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetModel)
    res["remoteSignInTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetFloat64Value(m.SetRemoteSignInTime)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    res["virtualNetwork"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVirtualNetwork)
    return res
}
// GetManufacturer gets the manufacturer property value. The user experience analytics manufacturer.
func (m *UserExperienceAnalyticsRemoteConnection) GetManufacturer()(*string) {
    return m.manufacturer
}
// GetModel gets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsRemoteConnection) GetModel()(*string) {
    return m.model
}
// GetRemoteSignInTime gets the remoteSignInTime property value. The remote sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetRemoteSignInTime()(*float64) {
    return m.remoteSignInTime
}
// GetUserPrincipalName gets the userPrincipalName property value. The user experience analytics userPrincipalName.
func (m *UserExperienceAnalyticsRemoteConnection) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetVirtualNetwork gets the virtualNetwork property value. The user experience analytics virtual network.
func (m *UserExperienceAnalyticsRemoteConnection) GetVirtualNetwork()(*string) {
    return m.virtualNetwork
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsRemoteConnection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("cloudPcFailurePercentage", m.GetCloudPcFailurePercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("cloudPcRoundTripTime", m.GetCloudPcRoundTripTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("cloudPcSignInTime", m.GetCloudPcSignInTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("coreBootTime", m.GetCoreBootTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("coreSignInTime", m.GetCoreSignInTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
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
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("remoteSignInTime", m.GetRemoteSignInTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("virtualNetwork", m.GetVirtualNetwork())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCloudPcFailurePercentage sets the cloudPcFailurePercentage property value. The sign in failure percentage of Cloud PC Device. Valid values 0 to 100
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcFailurePercentage(value *float64)() {
    m.cloudPcFailurePercentage = value
}
// SetCloudPcRoundTripTime sets the cloudPcRoundTripTime property value. The round tip time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcRoundTripTime(value *float64)() {
    m.cloudPcRoundTripTime = value
}
// SetCloudPcSignInTime sets the cloudPcSignInTime property value. The sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcSignInTime(value *float64)() {
    m.cloudPcSignInTime = value
}
// SetCoreBootTime sets the coreBootTime property value. The core boot time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCoreBootTime(value *float64)() {
    m.coreBootTime = value
}
// SetCoreSignInTime sets the coreSignInTime property value. The core sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCoreSignInTime(value *float64)() {
    m.coreSignInTime = value
}
// SetDeviceCount sets the deviceCount property value. The count of remote connection. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceCount(value *int32)() {
    m.deviceCount = value
}
// SetDeviceId sets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetDeviceName sets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetManufacturer sets the manufacturer property value. The user experience analytics manufacturer.
func (m *UserExperienceAnalyticsRemoteConnection) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// SetModel sets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsRemoteConnection) SetModel(value *string)() {
    m.model = value
}
// SetRemoteSignInTime sets the remoteSignInTime property value. The remote sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetRemoteSignInTime(value *float64)() {
    m.remoteSignInTime = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user experience analytics userPrincipalName.
func (m *UserExperienceAnalyticsRemoteConnection) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetVirtualNetwork sets the virtualNetwork property value. The user experience analytics virtual network.
func (m *UserExperienceAnalyticsRemoteConnection) SetVirtualNetwork(value *string)() {
    m.virtualNetwork = value
}
