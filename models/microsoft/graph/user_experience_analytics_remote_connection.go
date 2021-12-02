package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsRemoteConnection 
type UserExperienceAnalyticsRemoteConnection struct {
    Entity
    // The sign in failure percentage of Cloud PC Device. Valid values 0 to 100
    cloudPcFailurePercentage *float64;
    // The round tip time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
    cloudPcRoundTripTime *float64;
    // The sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
    cloudPcSignInTime *float64;
    // The core boot time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
    coreBootTime *float64;
    // The core sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
    coreSignInTime *float64;
    // The count of remote connection. Valid values 0 to 2147483647
    deviceCount *int32;
    // The id of the device.
    deviceId *string;
    // The name of the device.
    deviceName *string;
    // The user experience analytics manufacturer.
    manufacturer *string;
    // The user experience analytics device model.
    model *string;
    // The remote sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
    remoteSignInTime *float64;
    // The user experience analytics userPrincipalName.
    userPrincipalName *string;
    // The user experience analytics virtual network.
    virtualNetwork *string;
}
// NewUserExperienceAnalyticsRemoteConnection instantiates a new userExperienceAnalyticsRemoteConnection and sets the default values.
func NewUserExperienceAnalyticsRemoteConnection()(*UserExperienceAnalyticsRemoteConnection) {
    m := &UserExperienceAnalyticsRemoteConnection{
        Entity: *NewEntity(),
    }
    return m
}
// GetCloudPcFailurePercentage gets the cloudPcFailurePercentage property value. The sign in failure percentage of Cloud PC Device. Valid values 0 to 100
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcFailurePercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcFailurePercentage
    }
}
// GetCloudPcRoundTripTime gets the cloudPcRoundTripTime property value. The round tip time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcRoundTripTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcRoundTripTime
    }
}
// GetCloudPcSignInTime gets the cloudPcSignInTime property value. The sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcSignInTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcSignInTime
    }
}
// GetCoreBootTime gets the coreBootTime property value. The core boot time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCoreBootTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.coreBootTime
    }
}
// GetCoreSignInTime gets the coreSignInTime property value. The core sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCoreSignInTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.coreSignInTime
    }
}
// GetDeviceCount gets the deviceCount property value. The count of remote connection. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
// GetDeviceId gets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetDeviceName gets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetManufacturer gets the manufacturer property value. The user experience analytics manufacturer.
func (m *UserExperienceAnalyticsRemoteConnection) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetModel gets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsRemoteConnection) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetRemoteSignInTime gets the remoteSignInTime property value. The remote sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetRemoteSignInTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.remoteSignInTime
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. The user experience analytics userPrincipalName.
func (m *UserExperienceAnalyticsRemoteConnection) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// GetVirtualNetwork gets the virtualNetwork property value. The user experience analytics virtual network.
func (m *UserExperienceAnalyticsRemoteConnection) GetVirtualNetwork()(*string) {
    if m == nil {
        return nil
    } else {
        return m.virtualNetwork
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsRemoteConnection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cloudPcFailurePercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcFailurePercentage(val)
        }
        return nil
    }
    res["cloudPcRoundTripTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcRoundTripTime(val)
        }
        return nil
    }
    res["cloudPcSignInTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudPcSignInTime(val)
        }
        return nil
    }
    res["coreBootTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreBootTime(val)
        }
        return nil
    }
    res["coreSignInTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoreSignInTime(val)
        }
        return nil
    }
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
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
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
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["remoteSignInTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteSignInTime(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["virtualNetwork"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *UserExperienceAnalyticsRemoteConnection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsRemoteConnection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.cloudPcFailurePercentage = value
    }
}
// SetCloudPcRoundTripTime sets the cloudPcRoundTripTime property value. The round tip time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcRoundTripTime(value *float64)() {
    if m != nil {
        m.cloudPcRoundTripTime = value
    }
}
// SetCloudPcSignInTime sets the cloudPcSignInTime property value. The sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcSignInTime(value *float64)() {
    if m != nil {
        m.cloudPcSignInTime = value
    }
}
// SetCoreBootTime sets the coreBootTime property value. The core boot time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCoreBootTime(value *float64)() {
    if m != nil {
        m.coreBootTime = value
    }
}
// SetCoreSignInTime sets the coreSignInTime property value. The core sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetCoreSignInTime(value *float64)() {
    if m != nil {
        m.coreSignInTime = value
    }
}
// SetDeviceCount sets the deviceCount property value. The count of remote connection. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceCount(value *int32)() {
    if m != nil {
        m.deviceCount = value
    }
}
// SetDeviceId sets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetDeviceName sets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetManufacturer sets the manufacturer property value. The user experience analytics manufacturer.
func (m *UserExperienceAnalyticsRemoteConnection) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetModel sets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsRemoteConnection) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetRemoteSignInTime sets the remoteSignInTime property value. The remote sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) SetRemoteSignInTime(value *float64)() {
    if m != nil {
        m.remoteSignInTime = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The user experience analytics userPrincipalName.
func (m *UserExperienceAnalyticsRemoteConnection) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
// SetVirtualNetwork sets the virtualNetwork property value. The user experience analytics virtual network.
func (m *UserExperienceAnalyticsRemoteConnection) SetVirtualNetwork(value *string)() {
    if m != nil {
        m.virtualNetwork = value
    }
}
