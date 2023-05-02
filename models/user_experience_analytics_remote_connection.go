package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsRemoteConnection 
type UserExperienceAnalyticsRemoteConnection struct {
    Entity
}
// NewUserExperienceAnalyticsRemoteConnection instantiates a new UserExperienceAnalyticsRemoteConnection and sets the default values.
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
    val, err := m.GetBackingStore().Get("cloudPcFailurePercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetCloudPcRoundTripTime gets the cloudPcRoundTripTime property value. The round tip time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcRoundTripTime()(*float64) {
    val, err := m.GetBackingStore().Get("cloudPcRoundTripTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetCloudPcSignInTime gets the cloudPcSignInTime property value. The sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcSignInTime()(*float64) {
    val, err := m.GetBackingStore().Get("cloudPcSignInTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetCoreBootTime gets the coreBootTime property value. The core boot time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCoreBootTime()(*float64) {
    val, err := m.GetBackingStore().Get("coreBootTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetCoreSignInTime gets the coreSignInTime property value. The core sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCoreSignInTime()(*float64) {
    val, err := m.GetBackingStore().Get("coreSignInTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetDeviceCount gets the deviceCount property value. The count of remote connection. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("deviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceName()(*string) {
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
func (m *UserExperienceAnalyticsRemoteConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cloudPcFailurePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcFailurePercentage(val)
        }
        return nil
    }
    res["cloudPcRoundTripTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcRoundTripTime(val)
        }
        return nil
    }
    res["cloudPcSignInTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcSignInTime(val)
        }
        return nil
    }
    res["coreBootTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreBootTime(val)
        }
        return nil
    }
    res["coreSignInTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreSignInTime(val)
        }
        return nil
    }
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
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
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["remoteSignInTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteSignInTime(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["virtualNetwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualNetwork(val)
        }
        return nil
    }
    return res
}
// GetManufacturer gets the manufacturer property value. The user experience analytics manufacturer.
func (m *UserExperienceAnalyticsRemoteConnection) GetManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("manufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetModel gets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsRemoteConnection) GetModel()(*string) {
    val, err := m.GetBackingStore().Get("model")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemoteSignInTime gets the remoteSignInTime property value. The remote sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetRemoteSignInTime()(*float64) {
    val, err := m.GetBackingStore().Get("remoteSignInTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. The user experience analytics userPrincipalName.
func (m *UserExperienceAnalyticsRemoteConnection) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVirtualNetwork gets the virtualNetwork property value. The user experience analytics virtual network.
func (m *UserExperienceAnalyticsRemoteConnection) GetVirtualNetwork()(*string) {
    val, err := m.GetBackingStore().Get("virtualNetwork")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    err := m.GetBackingStore().Set("cloudPcFailurePercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcRoundTripTime sets the cloudPcRoundTripTime property value. The round tip time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcRoundTripTime(value *float64)() {
    err := m.GetBackingStore().Set("cloudPcRoundTripTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcSignInTime sets the cloudPcSignInTime property value. The sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcSignInTime(value *float64)() {
    err := m.GetBackingStore().Set("cloudPcSignInTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCoreBootTime sets the coreBootTime property value. The core boot time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCoreBootTime(value *float64)() {
    err := m.GetBackingStore().Set("coreBootTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCoreSignInTime sets the coreSignInTime property value. The core sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCoreSignInTime(value *float64)() {
    err := m.GetBackingStore().Set("coreSignInTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCount sets the deviceCount property value. The count of remote connection. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("deviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetManufacturer sets the manufacturer property value. The user experience analytics manufacturer.
func (m *UserExperienceAnalyticsRemoteConnection) SetManufacturer(value *string)() {
    err := m.GetBackingStore().Set("manufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetModel sets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsRemoteConnection) SetModel(value *string)() {
    err := m.GetBackingStore().Set("model", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteSignInTime sets the remoteSignInTime property value. The remote sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetRemoteSignInTime(value *float64)() {
    err := m.GetBackingStore().Set("remoteSignInTime", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user experience analytics userPrincipalName.
func (m *UserExperienceAnalyticsRemoteConnection) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetVirtualNetwork sets the virtualNetwork property value. The user experience analytics virtual network.
func (m *UserExperienceAnalyticsRemoteConnection) SetVirtualNetwork(value *string)() {
    err := m.GetBackingStore().Set("virtualNetwork", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsRemoteConnectionable 
type UserExperienceAnalyticsRemoteConnectionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudPcFailurePercentage()(*float64)
    GetCloudPcRoundTripTime()(*float64)
    GetCloudPcSignInTime()(*float64)
    GetCoreBootTime()(*float64)
    GetCoreSignInTime()(*float64)
    GetDeviceCount()(*int32)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetRemoteSignInTime()(*float64)
    GetUserPrincipalName()(*string)
    GetVirtualNetwork()(*string)
    SetCloudPcFailurePercentage(value *float64)()
    SetCloudPcRoundTripTime(value *float64)()
    SetCloudPcSignInTime(value *float64)()
    SetCoreBootTime(value *float64)()
    SetCoreSignInTime(value *float64)()
    SetDeviceCount(value *int32)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetRemoteSignInTime(value *float64)()
    SetUserPrincipalName(value *string)()
    SetVirtualNetwork(value *string)()
}
