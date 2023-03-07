package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceScope the user experience analytics device scope entity contains device scope configuration values use to apply filtering on the endpoint analytics reports.
type UserExperienceAnalyticsDeviceScope struct {
    Entity
}
// NewUserExperienceAnalyticsDeviceScope instantiates a new userExperienceAnalyticsDeviceScope and sets the default values.
func NewUserExperienceAnalyticsDeviceScope()(*UserExperienceAnalyticsDeviceScope) {
    m := &UserExperienceAnalyticsDeviceScope{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserExperienceAnalyticsDeviceScopeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsDeviceScopeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsDeviceScope(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Indicates the creation date and time for the custom device scope.
func (m *UserExperienceAnalyticsDeviceScope) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDeviceScopeName gets the deviceScopeName property value. The name of the user experience analytics device Scope configuration.
func (m *UserExperienceAnalyticsDeviceScope) GetDeviceScopeName()(*string) {
    val, err := m.GetBackingStore().Get("deviceScopeName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnabled gets the enabled property value. Indicates whether a device scope is enabled or disabled. When TRUE, the device scope is enabled. When FALSE, the device scope is disabled. Default value is FALSE.
func (m *UserExperienceAnalyticsDeviceScope) GetEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("enabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceScope) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deviceScopeName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceScopeName(val)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["isBuiltIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBuiltIn(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceScopeOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val.(*DeviceScopeOperator))
        }
        return nil
    }
    res["ownerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerId(val)
        }
        return nil
    }
    res["parameter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceScopeParameter)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameter(val.(*DeviceScopeParameter))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceScopeStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DeviceScopeStatus))
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    res["valueObjectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueObjectId(val)
        }
        return nil
    }
    return res
}
// GetIsBuiltIn gets the isBuiltIn property value. Indicates whether the device scope configuration is built-in or custom. When TRUE, the device scope configuration is built-in. When FALSE, the device scope configuration is custom. Default value is FALSE.
func (m *UserExperienceAnalyticsDeviceScope) GetIsBuiltIn()(*bool) {
    val, err := m.GetBackingStore().Get("isBuiltIn")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Indicates the last updated date and time for the custom device scope.
func (m *UserExperienceAnalyticsDeviceScope) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOperator gets the operator property value. Device scope configuration query operator. Possible values are: equals, notEquals, contains, notContains, greaterThan, lessThan. Default value: equals.
func (m *UserExperienceAnalyticsDeviceScope) GetOperator()(*DeviceScopeOperator) {
    val, err := m.GetBackingStore().Get("operator")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceScopeOperator)
    }
    return nil
}
// GetOwnerId gets the ownerId property value. The unique identifier of the person (admin) who created the device scope configuration.
func (m *UserExperienceAnalyticsDeviceScope) GetOwnerId()(*string) {
    val, err := m.GetBackingStore().Get("ownerId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetParameter gets the parameter property value. Device scope configuration parameter. It will be expend in future to add more parameter. Eg: device scope parameter can be OS version, Disk Type, Device manufacturer, device model or Scope tag. Default value: scopeTag.
func (m *UserExperienceAnalyticsDeviceScope) GetParameter()(*DeviceScopeParameter) {
    val, err := m.GetBackingStore().Get("parameter")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceScopeParameter)
    }
    return nil
}
// GetStatus gets the status property value. Indicates the device scope status after the device scope has been enabled. Possible values are: none, computing, insufficientData or completed. Default value is none.
func (m *UserExperienceAnalyticsDeviceScope) GetStatus()(*DeviceScopeStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceScopeStatus)
    }
    return nil
}
// GetValue gets the value property value. The device scope configuration query clause value.
func (m *UserExperienceAnalyticsDeviceScope) GetValue()(*string) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValueObjectId gets the valueObjectId property value. The unique identifier for a user device scope tag Id used for the creation of device scope configuration.
func (m *UserExperienceAnalyticsDeviceScope) GetValueObjectId()(*string) {
    val, err := m.GetBackingStore().Get("valueObjectId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsDeviceScope) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceScopeName", m.GetDeviceScopeName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBuiltIn", m.GetIsBuiltIn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOperator() != nil {
        cast := (*m.GetOperator()).String()
        err = writer.WriteStringValue("operator", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerId", m.GetOwnerId())
        if err != nil {
            return err
        }
    }
    if m.GetParameter() != nil {
        cast := (*m.GetParameter()).String()
        err = writer.WriteStringValue("parameter", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("valueObjectId", m.GetValueObjectId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Indicates the creation date and time for the custom device scope.
func (m *UserExperienceAnalyticsDeviceScope) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceScopeName sets the deviceScopeName property value. The name of the user experience analytics device Scope configuration.
func (m *UserExperienceAnalyticsDeviceScope) SetDeviceScopeName(value *string)() {
    err := m.GetBackingStore().Set("deviceScopeName", value)
    if err != nil {
        panic(err)
    }
}
// SetEnabled sets the enabled property value. Indicates whether a device scope is enabled or disabled. When TRUE, the device scope is enabled. When FALSE, the device scope is disabled. Default value is FALSE.
func (m *UserExperienceAnalyticsDeviceScope) SetEnabled(value *bool)() {
    err := m.GetBackingStore().Set("enabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsBuiltIn sets the isBuiltIn property value. Indicates whether the device scope configuration is built-in or custom. When TRUE, the device scope configuration is built-in. When FALSE, the device scope configuration is custom. Default value is FALSE.
func (m *UserExperienceAnalyticsDeviceScope) SetIsBuiltIn(value *bool)() {
    err := m.GetBackingStore().Set("isBuiltIn", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Indicates the last updated date and time for the custom device scope.
func (m *UserExperienceAnalyticsDeviceScope) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOperator sets the operator property value. Device scope configuration query operator. Possible values are: equals, notEquals, contains, notContains, greaterThan, lessThan. Default value: equals.
func (m *UserExperienceAnalyticsDeviceScope) SetOperator(value *DeviceScopeOperator)() {
    err := m.GetBackingStore().Set("operator", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnerId sets the ownerId property value. The unique identifier of the person (admin) who created the device scope configuration.
func (m *UserExperienceAnalyticsDeviceScope) SetOwnerId(value *string)() {
    err := m.GetBackingStore().Set("ownerId", value)
    if err != nil {
        panic(err)
    }
}
// SetParameter sets the parameter property value. Device scope configuration parameter. It will be expend in future to add more parameter. Eg: device scope parameter can be OS version, Disk Type, Device manufacturer, device model or Scope tag. Default value: scopeTag.
func (m *UserExperienceAnalyticsDeviceScope) SetParameter(value *DeviceScopeParameter)() {
    err := m.GetBackingStore().Set("parameter", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Indicates the device scope status after the device scope has been enabled. Possible values are: none, computing, insufficientData or completed. Default value is none.
func (m *UserExperienceAnalyticsDeviceScope) SetStatus(value *DeviceScopeStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetValue sets the value property value. The device scope configuration query clause value.
func (m *UserExperienceAnalyticsDeviceScope) SetValue(value *string)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// SetValueObjectId sets the valueObjectId property value. The unique identifier for a user device scope tag Id used for the creation of device scope configuration.
func (m *UserExperienceAnalyticsDeviceScope) SetValueObjectId(value *string)() {
    err := m.GetBackingStore().Set("valueObjectId", value)
    if err != nil {
        panic(err)
    }
}
// UserExperienceAnalyticsDeviceScopeable 
type UserExperienceAnalyticsDeviceScopeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceScopeName()(*string)
    GetEnabled()(*bool)
    GetIsBuiltIn()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOperator()(*DeviceScopeOperator)
    GetOwnerId()(*string)
    GetParameter()(*DeviceScopeParameter)
    GetStatus()(*DeviceScopeStatus)
    GetValue()(*string)
    GetValueObjectId()(*string)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceScopeName(value *string)()
    SetEnabled(value *bool)()
    SetIsBuiltIn(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOperator(value *DeviceScopeOperator)()
    SetOwnerId(value *string)()
    SetParameter(value *DeviceScopeParameter)()
    SetStatus(value *DeviceScopeStatus)()
    SetValue(value *string)()
    SetValueObjectId(value *string)()
}
