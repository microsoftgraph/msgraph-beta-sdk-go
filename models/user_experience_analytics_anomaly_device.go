package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAnomalyDevice the user experience analytics anomaly entity contains device details.
type UserExperienceAnalyticsAnomalyDevice struct {
    Entity
}
// NewUserExperienceAnalyticsAnomalyDevice instantiates a new UserExperienceAnalyticsAnomalyDevice and sets the default values.
func NewUserExperienceAnalyticsAnomalyDevice()(*UserExperienceAnalyticsAnomalyDevice) {
    m := &UserExperienceAnalyticsAnomalyDevice{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsAnomalyDeviceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsAnomalyDeviceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsAnomalyDevice(), nil
}
// GetAnomalyId gets the anomalyId property value. The unique identifier of the anomaly.
// returns a *string when successful
func (m *UserExperienceAnalyticsAnomalyDevice) GetAnomalyId()(*string) {
    val, err := m.GetBackingStore().Get("anomalyId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAnomalyOnDeviceFirstOccurrenceDateTime gets the anomalyOnDeviceFirstOccurrenceDateTime property value. Indicates the first occurance date and time for the anomaly on the device.
// returns a *Time when successful
func (m *UserExperienceAnalyticsAnomalyDevice) GetAnomalyOnDeviceFirstOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("anomalyOnDeviceFirstOccurrenceDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetAnomalyOnDeviceLatestOccurrenceDateTime gets the anomalyOnDeviceLatestOccurrenceDateTime property value. Indicates the latest occurance date and time for the anomaly on the device.
// returns a *Time when successful
func (m *UserExperienceAnalyticsAnomalyDevice) GetAnomalyOnDeviceLatestOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("anomalyOnDeviceLatestOccurrenceDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCorrelationGroupId gets the correlationGroupId property value. The unique identifier of the correlation group.
// returns a *string when successful
func (m *UserExperienceAnalyticsAnomalyDevice) GetCorrelationGroupId()(*string) {
    val, err := m.GetBackingStore().Get("correlationGroupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The unique identifier of the device.
// returns a *string when successful
func (m *UserExperienceAnalyticsAnomalyDevice) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceManufacturer gets the deviceManufacturer property value. The manufacturer name of the device.
// returns a *string when successful
func (m *UserExperienceAnalyticsAnomalyDevice) GetDeviceManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("deviceManufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceModel gets the deviceModel property value. The model name of the device.
// returns a *string when successful
func (m *UserExperienceAnalyticsAnomalyDevice) GetDeviceModel()(*string) {
    val, err := m.GetBackingStore().Get("deviceModel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The name of the device.
// returns a *string when successful
func (m *UserExperienceAnalyticsAnomalyDevice) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceStatus gets the deviceStatus property value. Indicates the status of the device in the correlation group. Eg: Device status can be anomalous, affected, at risk.
// returns a *UserExperienceAnalyticsDeviceStatus when successful
func (m *UserExperienceAnalyticsAnomalyDevice) GetDeviceStatus()(*UserExperienceAnalyticsDeviceStatus) {
    val, err := m.GetBackingStore().Get("deviceStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserExperienceAnalyticsDeviceStatus)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsAnomalyDevice) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["anomalyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyId(val)
        }
        return nil
    }
    res["anomalyOnDeviceFirstOccurrenceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyOnDeviceFirstOccurrenceDateTime(val)
        }
        return nil
    }
    res["anomalyOnDeviceLatestOccurrenceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnomalyOnDeviceLatestOccurrenceDateTime(val)
        }
        return nil
    }
    res["correlationGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationGroupId(val)
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
    res["deviceManufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManufacturer(val)
        }
        return nil
    }
    res["deviceModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
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
    res["deviceStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserExperienceAnalyticsDeviceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceStatus(val.(*UserExperienceAnalyticsDeviceStatus))
        }
        return nil
    }
    res["osName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsName(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    return res
}
// GetOsName gets the osName property value. The name of the OS installed on the device.
// returns a *string when successful
func (m *UserExperienceAnalyticsAnomalyDevice) GetOsName()(*string) {
    val, err := m.GetBackingStore().Get("osName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOsVersion gets the osVersion property value. The OS version installed on the device.
// returns a *string when successful
func (m *UserExperienceAnalyticsAnomalyDevice) GetOsVersion()(*string) {
    val, err := m.GetBackingStore().Get("osVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsAnomalyDevice) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("anomalyId", m.GetAnomalyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("anomalyOnDeviceFirstOccurrenceDateTime", m.GetAnomalyOnDeviceFirstOccurrenceDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("anomalyOnDeviceLatestOccurrenceDateTime", m.GetAnomalyOnDeviceLatestOccurrenceDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("correlationGroupId", m.GetCorrelationGroupId())
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
        err = writer.WriteStringValue("deviceManufacturer", m.GetDeviceManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModel", m.GetDeviceModel())
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
    if m.GetDeviceStatus() != nil {
        cast := (*m.GetDeviceStatus()).String()
        err = writer.WriteStringValue("deviceStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osName", m.GetOsName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnomalyId sets the anomalyId property value. The unique identifier of the anomaly.
func (m *UserExperienceAnalyticsAnomalyDevice) SetAnomalyId(value *string)() {
    err := m.GetBackingStore().Set("anomalyId", value)
    if err != nil {
        panic(err)
    }
}
// SetAnomalyOnDeviceFirstOccurrenceDateTime sets the anomalyOnDeviceFirstOccurrenceDateTime property value. Indicates the first occurance date and time for the anomaly on the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetAnomalyOnDeviceFirstOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("anomalyOnDeviceFirstOccurrenceDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetAnomalyOnDeviceLatestOccurrenceDateTime sets the anomalyOnDeviceLatestOccurrenceDateTime property value. Indicates the latest occurance date and time for the anomaly on the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetAnomalyOnDeviceLatestOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("anomalyOnDeviceLatestOccurrenceDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCorrelationGroupId sets the correlationGroupId property value. The unique identifier of the correlation group.
func (m *UserExperienceAnalyticsAnomalyDevice) SetCorrelationGroupId(value *string)() {
    err := m.GetBackingStore().Set("correlationGroupId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The unique identifier of the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceManufacturer sets the deviceManufacturer property value. The manufacturer name of the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetDeviceManufacturer(value *string)() {
    err := m.GetBackingStore().Set("deviceManufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceModel sets the deviceModel property value. The model name of the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetDeviceModel(value *string)() {
    err := m.GetBackingStore().Set("deviceModel", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The name of the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceStatus sets the deviceStatus property value. Indicates the status of the device in the correlation group. Eg: Device status can be anomalous, affected, at risk.
func (m *UserExperienceAnalyticsAnomalyDevice) SetDeviceStatus(value *UserExperienceAnalyticsDeviceStatus)() {
    err := m.GetBackingStore().Set("deviceStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetOsName sets the osName property value. The name of the OS installed on the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetOsName(value *string)() {
    err := m.GetBackingStore().Set("osName", value)
    if err != nil {
        panic(err)
    }
}
// SetOsVersion sets the osVersion property value. The OS version installed on the device.
func (m *UserExperienceAnalyticsAnomalyDevice) SetOsVersion(value *string)() {
    err := m.GetBackingStore().Set("osVersion", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsAnomalyDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnomalyId()(*string)
    GetAnomalyOnDeviceFirstOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAnomalyOnDeviceLatestOccurrenceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCorrelationGroupId()(*string)
    GetDeviceId()(*string)
    GetDeviceManufacturer()(*string)
    GetDeviceModel()(*string)
    GetDeviceName()(*string)
    GetDeviceStatus()(*UserExperienceAnalyticsDeviceStatus)
    GetOsName()(*string)
    GetOsVersion()(*string)
    SetAnomalyId(value *string)()
    SetAnomalyOnDeviceFirstOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAnomalyOnDeviceLatestOccurrenceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCorrelationGroupId(value *string)()
    SetDeviceId(value *string)()
    SetDeviceManufacturer(value *string)()
    SetDeviceModel(value *string)()
    SetDeviceName(value *string)()
    SetDeviceStatus(value *UserExperienceAnalyticsDeviceStatus)()
    SetOsName(value *string)()
    SetOsVersion(value *string)()
}
