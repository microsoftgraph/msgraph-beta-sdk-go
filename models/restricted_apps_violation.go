package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RestrictedAppsViolation violation of restricted apps configuration profile per device per user
type RestrictedAppsViolation struct {
    Entity
    // Device configuration profile unique identifier, must be Guid
    deviceConfigurationId *string
    // Device configuration profile name
    deviceConfigurationName *string
    // Device name
    deviceName *string
    // Managed device unique identifier, must be Guid
    managedDeviceId *string
    // Supported platform types for policies.
    platformType *PolicyPlatformType
    // List of violated restricted apps
    restrictedApps []ManagedDeviceReportedAppable
    // Restricted apps state
    restrictedAppsState *RestrictedAppsState
    // User unique identifier, must be Guid
    userId *string
    // User name
    userName *string
}
// NewRestrictedAppsViolation instantiates a new restrictedAppsViolation and sets the default values.
func NewRestrictedAppsViolation()(*RestrictedAppsViolation) {
    m := &RestrictedAppsViolation{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.restrictedAppsViolation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRestrictedAppsViolationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRestrictedAppsViolationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRestrictedAppsViolation(), nil
}
// GetDeviceConfigurationId gets the deviceConfigurationId property value. Device configuration profile unique identifier, must be Guid
func (m *RestrictedAppsViolation) GetDeviceConfigurationId()(*string) {
    return m.deviceConfigurationId
}
// GetDeviceConfigurationName gets the deviceConfigurationName property value. Device configuration profile name
func (m *RestrictedAppsViolation) GetDeviceConfigurationName()(*string) {
    return m.deviceConfigurationName
}
// GetDeviceName gets the deviceName property value. Device name
func (m *RestrictedAppsViolation) GetDeviceName()(*string) {
    return m.deviceName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RestrictedAppsViolation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceConfigurationId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceConfigurationId)
    res["deviceConfigurationName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceConfigurationName)
    res["deviceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDeviceName)
    res["managedDeviceId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagedDeviceId)
    res["platformType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParsePolicyPlatformType , m.SetPlatformType)
    res["restrictedApps"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagedDeviceReportedAppFromDiscriminatorValue , m.SetRestrictedApps)
    res["restrictedAppsState"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRestrictedAppsState , m.SetRestrictedAppsState)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    res["userName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserName)
    return res
}
// GetManagedDeviceId gets the managedDeviceId property value. Managed device unique identifier, must be Guid
func (m *RestrictedAppsViolation) GetManagedDeviceId()(*string) {
    return m.managedDeviceId
}
// GetPlatformType gets the platformType property value. Supported platform types for policies.
func (m *RestrictedAppsViolation) GetPlatformType()(*PolicyPlatformType) {
    return m.platformType
}
// GetRestrictedApps gets the restrictedApps property value. List of violated restricted apps
func (m *RestrictedAppsViolation) GetRestrictedApps()([]ManagedDeviceReportedAppable) {
    return m.restrictedApps
}
// GetRestrictedAppsState gets the restrictedAppsState property value. Restricted apps state
func (m *RestrictedAppsViolation) GetRestrictedAppsState()(*RestrictedAppsState) {
    return m.restrictedAppsState
}
// GetUserId gets the userId property value. User unique identifier, must be Guid
func (m *RestrictedAppsViolation) GetUserId()(*string) {
    return m.userId
}
// GetUserName gets the userName property value. User name
func (m *RestrictedAppsViolation) GetUserName()(*string) {
    return m.userName
}
// Serialize serializes information the current object
func (m *RestrictedAppsViolation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceConfigurationId", m.GetDeviceConfigurationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceConfigurationName", m.GetDeviceConfigurationName())
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
        err = writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    if m.GetPlatformType() != nil {
        cast := (*m.GetPlatformType()).String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRestrictedApps() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRestrictedApps())
        err = writer.WriteCollectionOfObjectValues("restrictedApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRestrictedAppsState() != nil {
        cast := (*m.GetRestrictedAppsState()).String()
        err = writer.WriteStringValue("restrictedAppsState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceConfigurationId sets the deviceConfigurationId property value. Device configuration profile unique identifier, must be Guid
func (m *RestrictedAppsViolation) SetDeviceConfigurationId(value *string)() {
    m.deviceConfigurationId = value
}
// SetDeviceConfigurationName sets the deviceConfigurationName property value. Device configuration profile name
func (m *RestrictedAppsViolation) SetDeviceConfigurationName(value *string)() {
    m.deviceConfigurationName = value
}
// SetDeviceName sets the deviceName property value. Device name
func (m *RestrictedAppsViolation) SetDeviceName(value *string)() {
    m.deviceName = value
}
// SetManagedDeviceId sets the managedDeviceId property value. Managed device unique identifier, must be Guid
func (m *RestrictedAppsViolation) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// SetPlatformType sets the platformType property value. Supported platform types for policies.
func (m *RestrictedAppsViolation) SetPlatformType(value *PolicyPlatformType)() {
    m.platformType = value
}
// SetRestrictedApps sets the restrictedApps property value. List of violated restricted apps
func (m *RestrictedAppsViolation) SetRestrictedApps(value []ManagedDeviceReportedAppable)() {
    m.restrictedApps = value
}
// SetRestrictedAppsState sets the restrictedAppsState property value. Restricted apps state
func (m *RestrictedAppsViolation) SetRestrictedAppsState(value *RestrictedAppsState)() {
    m.restrictedAppsState = value
}
// SetUserId sets the userId property value. User unique identifier, must be Guid
func (m *RestrictedAppsViolation) SetUserId(value *string)() {
    m.userId = value
}
// SetUserName sets the userName property value. User name
func (m *RestrictedAppsViolation) SetUserName(value *string)() {
    m.userName = value
}
