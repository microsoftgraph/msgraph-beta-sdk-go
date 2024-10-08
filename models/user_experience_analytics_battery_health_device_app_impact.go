package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// UserExperienceAnalyticsBatteryHealthDeviceAppImpact the user experience analytics battery health device app impact entity contains battery usage related information at an app level for a given device.
type UserExperienceAnalyticsBatteryHealthDeviceAppImpact struct {
    Entity
}
// UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage composed type wrapper for classes float64, ReferenceNumeric, string
type UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage instantiates a new UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage()(*UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage) {
    m := &UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage) GetReferenceNumeric()(*ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentageable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
// NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact instantiates a new UserExperienceAnalyticsBatteryHealthDeviceAppImpact and sets the default values.
func NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact()(*UserExperienceAnalyticsBatteryHealthDeviceAppImpact) {
    m := &UserExperienceAnalyticsBatteryHealthDeviceAppImpact{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsBatteryHealthDeviceAppImpact(), nil
}
// GetAppDisplayName gets the appDisplayName property value. User friendly display name for the app. Eg: Outlook
// returns a *string when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetAppDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("appDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppName gets the appName property value. App name. Eg: oltk.exe
// returns a *string when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetAppName()(*string) {
    val, err := m.GetBackingStore().Get("appName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAppPublisher gets the appPublisher property value. App publisher. Eg: Microsoft Corporation
// returns a *string when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetAppPublisher()(*string) {
    val, err := m.GetBackingStore().Get("appPublisher")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBatteryUsagePercentage gets the batteryUsagePercentage property value. The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days. Unit in percentage. Valid values 0 to 1.79769313486232E+308
// returns a UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentageable when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetBatteryUsagePercentage()(UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentageable) {
    val, err := m.GetBackingStore().Get("batteryUsagePercentage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentageable)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The unique identifier of the device, Intune DeviceID or SCCM device id.
// returns a *string when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
        }
        return nil
    }
    res["appName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppName(val)
        }
        return nil
    }
    res["appPublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppPublisher(val)
        }
        return nil
    }
    res["batteryUsagePercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBatteryUsagePercentage(val.(*UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentage))
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
    res["isForegroundApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsForegroundApp(val)
        }
        return nil
    }
    return res
}
// GetIsForegroundApp gets the isForegroundApp property value. true if the user had active interaction with the app.
// returns a *bool when successful
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) GetIsForegroundApp()(*bool) {
    val, err := m.GetBackingStore().Get("isForegroundApp")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appName", m.GetAppName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appPublisher", m.GetAppPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("batteryUsagePercentage", m.GetBatteryUsagePercentage())
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
        err = writer.WriteBoolValue("isForegroundApp", m.GetIsForegroundApp())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDisplayName sets the appDisplayName property value. User friendly display name for the app. Eg: Outlook
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetAppDisplayName(value *string)() {
    err := m.GetBackingStore().Set("appDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetAppName sets the appName property value. App name. Eg: oltk.exe
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetAppName(value *string)() {
    err := m.GetBackingStore().Set("appName", value)
    if err != nil {
        panic(err)
    }
}
// SetAppPublisher sets the appPublisher property value. App publisher. Eg: Microsoft Corporation
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetAppPublisher(value *string)() {
    err := m.GetBackingStore().Set("appPublisher", value)
    if err != nil {
        panic(err)
    }
}
// SetBatteryUsagePercentage sets the batteryUsagePercentage property value. The percent of total battery power used by this application when the device was not plugged into AC power, over 14 days. Unit in percentage. Valid values 0 to 1.79769313486232E+308
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetBatteryUsagePercentage(value UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentageable)() {
    err := m.GetBackingStore().Set("batteryUsagePercentage", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The unique identifier of the device, Intune DeviceID or SCCM device id.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetIsForegroundApp sets the isForegroundApp property value. true if the user had active interaction with the app.
func (m *UserExperienceAnalyticsBatteryHealthDeviceAppImpact) SetIsForegroundApp(value *bool)() {
    err := m.GetBackingStore().Set("isForegroundApp", value)
    if err != nil {
        panic(err)
    }
}
type UserExperienceAnalyticsBatteryHealthDeviceAppImpactable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppDisplayName()(*string)
    GetAppName()(*string)
    GetAppPublisher()(*string)
    GetBatteryUsagePercentage()(UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentageable)
    GetDeviceId()(*string)
    GetIsForegroundApp()(*bool)
    SetAppDisplayName(value *string)()
    SetAppName(value *string)()
    SetAppPublisher(value *string)()
    SetBatteryUsagePercentage(value UserExperienceAnalyticsBatteryHealthDeviceAppImpact_UserExperienceAnalyticsBatteryHealthDeviceAppImpact_batteryUsagePercentageable)()
    SetDeviceId(value *string)()
    SetIsForegroundApp(value *bool)()
}
