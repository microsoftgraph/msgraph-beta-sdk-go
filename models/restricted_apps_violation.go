package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RestrictedAppsViolation violation of restricted apps configuration profile per device per user
type RestrictedAppsViolation struct {
    Entity
}
// NewRestrictedAppsViolation instantiates a new restrictedAppsViolation and sets the default values.
func NewRestrictedAppsViolation()(*RestrictedAppsViolation) {
    m := &RestrictedAppsViolation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRestrictedAppsViolationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRestrictedAppsViolationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRestrictedAppsViolation(), nil
}
// GetDeviceConfigurationId gets the deviceConfigurationId property value. Device configuration profile unique identifier, must be Guid
func (m *RestrictedAppsViolation) GetDeviceConfigurationId()(*string) {
    val, err := m.GetBackingStore().Get("deviceConfigurationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceConfigurationName gets the deviceConfigurationName property value. Device configuration profile name
func (m *RestrictedAppsViolation) GetDeviceConfigurationName()(*string) {
    val, err := m.GetBackingStore().Get("deviceConfigurationName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. Device name
func (m *RestrictedAppsViolation) GetDeviceName()(*string) {
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
func (m *RestrictedAppsViolation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceConfigurationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigurationId(val)
        }
        return nil
    }
    res["deviceConfigurationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigurationName(val)
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
    res["managedDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceId(val)
        }
        return nil
    }
    res["platformType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*PolicyPlatformType))
        }
        return nil
    }
    res["restrictedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceReportedAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceReportedAppable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceReportedAppable)
            }
            m.SetRestrictedApps(res)
        }
        return nil
    }
    res["restrictedAppsState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRestrictedAppsState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrictedAppsState(val.(*RestrictedAppsState))
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    return res
}
// GetManagedDeviceId gets the managedDeviceId property value. Managed device unique identifier, must be Guid
func (m *RestrictedAppsViolation) GetManagedDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("managedDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPlatformType gets the platformType property value. Supported platform types for policies.
func (m *RestrictedAppsViolation) GetPlatformType()(*PolicyPlatformType) {
    val, err := m.GetBackingStore().Get("platformType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PolicyPlatformType)
    }
    return nil
}
// GetRestrictedApps gets the restrictedApps property value. List of violated restricted apps
func (m *RestrictedAppsViolation) GetRestrictedApps()([]ManagedDeviceReportedAppable) {
    val, err := m.GetBackingStore().Get("restrictedApps")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ManagedDeviceReportedAppable)
    }
    return nil
}
// GetRestrictedAppsState gets the restrictedAppsState property value. Restricted apps state
func (m *RestrictedAppsViolation) GetRestrictedAppsState()(*RestrictedAppsState) {
    val, err := m.GetBackingStore().Get("restrictedAppsState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*RestrictedAppsState)
    }
    return nil
}
// GetUserId gets the userId property value. User unique identifier, must be Guid
func (m *RestrictedAppsViolation) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserName gets the userName property value. User name
func (m *RestrictedAppsViolation) GetUserName()(*string) {
    val, err := m.GetBackingStore().Get("userName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRestrictedApps()))
        for i, v := range m.GetRestrictedApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
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
    err := m.GetBackingStore().Set("deviceConfigurationId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceConfigurationName sets the deviceConfigurationName property value. Device configuration profile name
func (m *RestrictedAppsViolation) SetDeviceConfigurationName(value *string)() {
    err := m.GetBackingStore().Set("deviceConfigurationName", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. Device name
func (m *RestrictedAppsViolation) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetManagedDeviceId sets the managedDeviceId property value. Managed device unique identifier, must be Guid
func (m *RestrictedAppsViolation) SetManagedDeviceId(value *string)() {
    err := m.GetBackingStore().Set("managedDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatformType sets the platformType property value. Supported platform types for policies.
func (m *RestrictedAppsViolation) SetPlatformType(value *PolicyPlatformType)() {
    err := m.GetBackingStore().Set("platformType", value)
    if err != nil {
        panic(err)
    }
}
// SetRestrictedApps sets the restrictedApps property value. List of violated restricted apps
func (m *RestrictedAppsViolation) SetRestrictedApps(value []ManagedDeviceReportedAppable)() {
    err := m.GetBackingStore().Set("restrictedApps", value)
    if err != nil {
        panic(err)
    }
}
// SetRestrictedAppsState sets the restrictedAppsState property value. Restricted apps state
func (m *RestrictedAppsViolation) SetRestrictedAppsState(value *RestrictedAppsState)() {
    err := m.GetBackingStore().Set("restrictedAppsState", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. User unique identifier, must be Guid
func (m *RestrictedAppsViolation) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserName sets the userName property value. User name
func (m *RestrictedAppsViolation) SetUserName(value *string)() {
    err := m.GetBackingStore().Set("userName", value)
    if err != nil {
        panic(err)
    }
}
// RestrictedAppsViolationable 
type RestrictedAppsViolationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceConfigurationId()(*string)
    GetDeviceConfigurationName()(*string)
    GetDeviceName()(*string)
    GetManagedDeviceId()(*string)
    GetPlatformType()(*PolicyPlatformType)
    GetRestrictedApps()([]ManagedDeviceReportedAppable)
    GetRestrictedAppsState()(*RestrictedAppsState)
    GetUserId()(*string)
    GetUserName()(*string)
    SetDeviceConfigurationId(value *string)()
    SetDeviceConfigurationName(value *string)()
    SetDeviceName(value *string)()
    SetManagedDeviceId(value *string)()
    SetPlatformType(value *PolicyPlatformType)()
    SetRestrictedApps(value []ManagedDeviceReportedAppable)()
    SetRestrictedAppsState(value *RestrictedAppsState)()
    SetUserId(value *string)()
    SetUserName(value *string)()
}
