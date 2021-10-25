package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UserExperienceAnalyticsRemoteConnection struct {
    Entity
    cloudPcFailurePercentage *float64;
    cloudPcRoundTripTime *float64;
    cloudPcSignInTime *float64;
    coreBootTime *float64;
    coreSignInTime *float64;
    deviceCount *int32;
    deviceId *string;
    deviceName *string;
    model *string;
    remoteSignInTime *float64;
    virtualNetwork *string;
}
func NewUserExperienceAnalyticsRemoteConnection()(*UserExperienceAnalyticsRemoteConnection) {
    m := &UserExperienceAnalyticsRemoteConnection{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcFailurePercentage()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcFailurePercentage
    }
}
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcRoundTripTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcRoundTripTime
    }
}
func (m *UserExperienceAnalyticsRemoteConnection) GetCloudPcSignInTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcSignInTime
    }
}
func (m *UserExperienceAnalyticsRemoteConnection) GetCoreBootTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.coreBootTime
    }
}
func (m *UserExperienceAnalyticsRemoteConnection) GetCoreSignInTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.coreSignInTime
    }
}
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deviceCount
    }
}
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
func (m *UserExperienceAnalyticsRemoteConnection) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
func (m *UserExperienceAnalyticsRemoteConnection) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
func (m *UserExperienceAnalyticsRemoteConnection) GetRemoteSignInTime()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.remoteSignInTime
    }
}
func (m *UserExperienceAnalyticsRemoteConnection) GetVirtualNetwork()(*string) {
    if m == nil {
        return nil
    } else {
        return m.virtualNetwork
    }
}
func (m *UserExperienceAnalyticsRemoteConnection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cloudPcFailurePercentage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCloudPcFailurePercentage(val)
        return nil
    }
    res["cloudPcRoundTripTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCloudPcRoundTripTime(val)
        return nil
    }
    res["cloudPcSignInTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCloudPcSignInTime(val)
        return nil
    }
    res["coreBootTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCoreBootTime(val)
        return nil
    }
    res["coreSignInTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetCoreSignInTime(val)
        return nil
    }
    res["deviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeviceCount(val)
        return nil
    }
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
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
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModel(val)
        return nil
    }
    res["remoteSignInTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetRemoteSignInTime(val)
        return nil
    }
    res["virtualNetwork"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVirtualNetwork(val)
        return nil
    }
    return res
}
func (m *UserExperienceAnalyticsRemoteConnection) IsNil()(bool) {
    return m == nil
}
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
        err = writer.WriteStringValue("virtualNetwork", m.GetVirtualNetwork())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcFailurePercentage(value *float64)() {
    m.cloudPcFailurePercentage = value
}
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcRoundTripTime(value *float64)() {
    m.cloudPcRoundTripTime = value
}
func (m *UserExperienceAnalyticsRemoteConnection) SetCloudPcSignInTime(value *float64)() {
    m.cloudPcSignInTime = value
}
func (m *UserExperienceAnalyticsRemoteConnection) SetCoreBootTime(value *float64)() {
    m.coreBootTime = value
}
func (m *UserExperienceAnalyticsRemoteConnection) SetCoreSignInTime(value *float64)() {
    m.coreSignInTime = value
}
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceCount(value *int32)() {
    m.deviceCount = value
}
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceId(value *string)() {
    m.deviceId = value
}
func (m *UserExperienceAnalyticsRemoteConnection) SetDeviceName(value *string)() {
    m.deviceName = value
}
func (m *UserExperienceAnalyticsRemoteConnection) SetModel(value *string)() {
    m.model = value
}
func (m *UserExperienceAnalyticsRemoteConnection) SetRemoteSignInTime(value *float64)() {
    m.remoteSignInTime = value
}
func (m *UserExperienceAnalyticsRemoteConnection) SetVirtualNetwork(value *string)() {
    m.virtualNetwork = value
}
