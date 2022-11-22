package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceScope the user experience analytics device scope entity contains device scope configuration values use to apply filtering on the endpoint analytics reports.
type UserExperienceAnalyticsDeviceScope struct {
    Entity
    // Indicates the creation date and time for the custom device scope.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the user experience analytics device Scope configuration.
    deviceScopeName *string
    // Indicates whether a device scope is enabled or disabled. When TRUE, the device scope is enabled. When FALSE, the device scope is disabled. Default value is FALSE.
    enabled *bool
    // Indicates whether the device scope configuration is built-in or custom. When TRUE, the device scope configuration is built-in. When FALSE, the device scope configuration is custom. Default value is FALSE.
    isBuiltIn *bool
    // Indicates the last updated date and time for the custom device scope.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Device scope configuration query operator. Possible values are: equals, notEquals, contains, notContains, greaterThan, lessThan. Default value: equals.
    operator *DeviceScopeOperator
    // The unique identifier of the person (admin) who created the device scope configuration.
    ownerId *string
    // Device scope configuration parameter. It will be expend in future to add more parameter. Eg: device scope parameter can be OS version, Disk Type, Device manufacturer, device model or Scope tag. Default value: scopeTag.
    parameter *DeviceScopeParameter
    // Indicates the device scope status after the device scope has been enabled. Possible values are: none, computing, insufficientData or completed. Default value is none.
    status *DeviceScopeStatus
    // The device scope configuration query clause value.
    value *string
    // The unique identifier for a user device scope tag Id used for the creation of device scope configuration.
    valueObjectId *string
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
    return m.createdDateTime
}
// GetDeviceScopeName gets the deviceScopeName property value. The name of the user experience analytics device Scope configuration.
func (m *UserExperienceAnalyticsDeviceScope) GetDeviceScopeName()(*string) {
    return m.deviceScopeName
}
// GetEnabled gets the enabled property value. Indicates whether a device scope is enabled or disabled. When TRUE, the device scope is enabled. When FALSE, the device scope is disabled. Default value is FALSE.
func (m *UserExperienceAnalyticsDeviceScope) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsDeviceScope) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["deviceScopeName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceScopeName)
    res["enabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEnabled)
    res["isBuiltIn"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsBuiltIn)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["operator"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceScopeOperator , m.SetOperator)
    res["ownerId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOwnerId)
    res["parameter"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceScopeParameter , m.SetParameter)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseDeviceScopeStatus , m.SetStatus)
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetValue)
    res["valueObjectId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetValueObjectId)
    return res
}
// GetIsBuiltIn gets the isBuiltIn property value. Indicates whether the device scope configuration is built-in or custom. When TRUE, the device scope configuration is built-in. When FALSE, the device scope configuration is custom. Default value is FALSE.
func (m *UserExperienceAnalyticsDeviceScope) GetIsBuiltIn()(*bool) {
    return m.isBuiltIn
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Indicates the last updated date and time for the custom device scope.
func (m *UserExperienceAnalyticsDeviceScope) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOperator gets the operator property value. Device scope configuration query operator. Possible values are: equals, notEquals, contains, notContains, greaterThan, lessThan. Default value: equals.
func (m *UserExperienceAnalyticsDeviceScope) GetOperator()(*DeviceScopeOperator) {
    return m.operator
}
// GetOwnerId gets the ownerId property value. The unique identifier of the person (admin) who created the device scope configuration.
func (m *UserExperienceAnalyticsDeviceScope) GetOwnerId()(*string) {
    return m.ownerId
}
// GetParameter gets the parameter property value. Device scope configuration parameter. It will be expend in future to add more parameter. Eg: device scope parameter can be OS version, Disk Type, Device manufacturer, device model or Scope tag. Default value: scopeTag.
func (m *UserExperienceAnalyticsDeviceScope) GetParameter()(*DeviceScopeParameter) {
    return m.parameter
}
// GetStatus gets the status property value. Indicates the device scope status after the device scope has been enabled. Possible values are: none, computing, insufficientData or completed. Default value is none.
func (m *UserExperienceAnalyticsDeviceScope) GetStatus()(*DeviceScopeStatus) {
    return m.status
}
// GetValue gets the value property value. The device scope configuration query clause value.
func (m *UserExperienceAnalyticsDeviceScope) GetValue()(*string) {
    return m.value
}
// GetValueObjectId gets the valueObjectId property value. The unique identifier for a user device scope tag Id used for the creation of device scope configuration.
func (m *UserExperienceAnalyticsDeviceScope) GetValueObjectId()(*string) {
    return m.valueObjectId
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
    m.createdDateTime = value
}
// SetDeviceScopeName sets the deviceScopeName property value. The name of the user experience analytics device Scope configuration.
func (m *UserExperienceAnalyticsDeviceScope) SetDeviceScopeName(value *string)() {
    m.deviceScopeName = value
}
// SetEnabled sets the enabled property value. Indicates whether a device scope is enabled or disabled. When TRUE, the device scope is enabled. When FALSE, the device scope is disabled. Default value is FALSE.
func (m *UserExperienceAnalyticsDeviceScope) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetIsBuiltIn sets the isBuiltIn property value. Indicates whether the device scope configuration is built-in or custom. When TRUE, the device scope configuration is built-in. When FALSE, the device scope configuration is custom. Default value is FALSE.
func (m *UserExperienceAnalyticsDeviceScope) SetIsBuiltIn(value *bool)() {
    m.isBuiltIn = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Indicates the last updated date and time for the custom device scope.
func (m *UserExperienceAnalyticsDeviceScope) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOperator sets the operator property value. Device scope configuration query operator. Possible values are: equals, notEquals, contains, notContains, greaterThan, lessThan. Default value: equals.
func (m *UserExperienceAnalyticsDeviceScope) SetOperator(value *DeviceScopeOperator)() {
    m.operator = value
}
// SetOwnerId sets the ownerId property value. The unique identifier of the person (admin) who created the device scope configuration.
func (m *UserExperienceAnalyticsDeviceScope) SetOwnerId(value *string)() {
    m.ownerId = value
}
// SetParameter sets the parameter property value. Device scope configuration parameter. It will be expend in future to add more parameter. Eg: device scope parameter can be OS version, Disk Type, Device manufacturer, device model or Scope tag. Default value: scopeTag.
func (m *UserExperienceAnalyticsDeviceScope) SetParameter(value *DeviceScopeParameter)() {
    m.parameter = value
}
// SetStatus sets the status property value. Indicates the device scope status after the device scope has been enabled. Possible values are: none, computing, insufficientData or completed. Default value is none.
func (m *UserExperienceAnalyticsDeviceScope) SetStatus(value *DeviceScopeStatus)() {
    m.status = value
}
// SetValue sets the value property value. The device scope configuration query clause value.
func (m *UserExperienceAnalyticsDeviceScope) SetValue(value *string)() {
    m.value = value
}
// SetValueObjectId sets the valueObjectId property value. The unique identifier for a user device scope tag Id used for the creation of device scope configuration.
func (m *UserExperienceAnalyticsDeviceScope) SetValueObjectId(value *string)() {
    m.valueObjectId = value
}
