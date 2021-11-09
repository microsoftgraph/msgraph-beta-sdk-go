package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new userExperienceAnalyticsRemoteConnection and sets the default values.
func NewUserExperienceAnalyticsRemoteConnection()(*UserExperienceAnalyticsRemoteConnection) {
    m := &UserExperienceAnalyticsRemoteConnection{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the cloudPcFailurePercentage property value. The sign in failure percentage of Cloud PC Device. Valid values 0 to 100
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcFailurePercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcFailurePercentage
    }
}
// Gets the cloudPcRoundTripTime property value. The round tip time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcRoundTripTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcRoundTripTime
    }
}
// Gets the cloudPcSignInTime property value. The sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcSignInTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcSignInTime
    }
}
// Gets the coreBootTime property value. The core boot time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCoreBootTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.coreBootTime
    }
}
// Gets the coreSignInTime property value. The core sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetCoreSignInTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.coreSignInTime
    }
}
// Gets the deviceCount property value. The count of remote connection. Valid values 0 to 2147483647
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
// Gets the deviceId property value. The id of the device.
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the manufacturer property value. The user experience analytics manufacturer.
func (m *UserExperienceAnalyticsRemoteConnection) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// Gets the model property value. The user experience analytics device model.
func (m *UserExperienceAnalyticsRemoteConnection) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// Gets the remoteSignInTime property value. The remote sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsRemoteConnection) GetRemoteSignInTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.remoteSignInTime
    }
}
// Gets the userPrincipalName property value. The user experience analytics userPrincipalName.
func (m *UserExperienceAnalyticsRemoteConnection) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Gets the virtualNetwork property value. The user experience analytics virtual network.
func (m *UserExperienceAnalyticsRemoteConnection) GetVirtualNetwork()(*string) {
    if m == nil {
        return nil
    } else {
        return m.virtualNetwork
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the cloudPcFailurePercentage property value. The sign in failure percentage of Cloud PC Device. Valid values 0 to 100
// Parameters:
//  - value : Value to set for the cloudPcFailurePercentage property.
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcFailurePercentage(value *float64)() {
    m.cloudPcFailurePercentage = value
}
// Sets the cloudPcRoundTripTime property value. The round tip time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the cloudPcRoundTripTime property.
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcRoundTripTime(value *float64)() {
    m.cloudPcRoundTripTime = value
}
// Sets the cloudPcSignInTime property value. The sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the cloudPcSignInTime property.
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcSignInTime(value *float64)() {
    m.cloudPcSignInTime = value
}
// Sets the coreBootTime property value. The core boot time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the coreBootTime property.
func (m *UserExperienceAnalyticsRemoteConnection) SetCoreBootTime(value *float64)() {
    m.coreBootTime = value
}
// Sets the coreSignInTime property value. The core sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the coreSignInTime property.
func (m *UserExperienceAnalyticsRemoteConnection) SetCoreSignInTime(value *float64)() {
    m.coreSignInTime = value
}
// Sets the deviceCount property value. The count of remote connection. Valid values 0 to 2147483647
// Parameters:
//  - value : Value to set for the deviceCount property.
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceCount(value *int32)() {
    m.deviceCount = value
}
// Sets the deviceId property value. The id of the device.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the deviceName property value. The name of the device.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the manufacturer property value. The user experience analytics manufacturer.
// Parameters:
//  - value : Value to set for the manufacturer property.
func (m *UserExperienceAnalyticsRemoteConnection) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// Sets the model property value. The user experience analytics device model.
// Parameters:
//  - value : Value to set for the model property.
func (m *UserExperienceAnalyticsRemoteConnection) SetModel(value *string)() {
    m.model = value
}
// Sets the remoteSignInTime property value. The remote sign in time of Cloud PC Device. Valid values 0 to 1.79769313486232E+308
// Parameters:
//  - value : Value to set for the remoteSignInTime property.
func (m *UserExperienceAnalyticsRemoteConnection) SetRemoteSignInTime(value *float64)() {
    m.remoteSignInTime = value
}
// Sets the userPrincipalName property value. The user experience analytics userPrincipalName.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *UserExperienceAnalyticsRemoteConnection) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// Sets the virtualNetwork property value. The user experience analytics virtual network.
// Parameters:
//  - value : Value to set for the virtualNetwork property.
func (m *UserExperienceAnalyticsRemoteConnection) SetVirtualNetwork(value *string)() {
    m.virtualNetwork = value
}
