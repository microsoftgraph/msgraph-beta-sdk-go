package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceCleanupRule define the rule when the admin wants the devices to be cleaned up.
type ManagedDeviceCleanupRule struct {
    Entity
}
// NewManagedDeviceCleanupRule instantiates a new ManagedDeviceCleanupRule and sets the default values.
func NewManagedDeviceCleanupRule()(*ManagedDeviceCleanupRule) {
    m := &ManagedDeviceCleanupRule{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedDeviceCleanupRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManagedDeviceCleanupRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceCleanupRule(), nil
}
// GetDescription gets the description property value. Indicates the description for the device clean up rule.
// returns a *string when successful
func (m *ManagedDeviceCleanupRule) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceCleanupRulePlatformType gets the deviceCleanupRulePlatformType property value. Define the platform type for which the admin wants to create the device clean up rule
// returns a *DeviceCleanupRulePlatformType when successful
func (m *ManagedDeviceCleanupRule) GetDeviceCleanupRulePlatformType()(*DeviceCleanupRulePlatformType) {
    val, err := m.GetBackingStore().Get("deviceCleanupRulePlatformType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceCleanupRulePlatformType)
    }
    return nil
}
// GetDeviceInactivityBeforeRetirementInDays gets the deviceInactivityBeforeRetirementInDays property value. Indicates the number of days when the device has not contacted Intune. Valid values 0 to 2147483647
// returns a *int32 when successful
func (m *ManagedDeviceCleanupRule) GetDeviceInactivityBeforeRetirementInDays()(*int32) {
    val, err := m.GetBackingStore().Get("deviceInactivityBeforeRetirementInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Indicates the display name of the device cleanup rule.
// returns a *string when successful
func (m *ManagedDeviceCleanupRule) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
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
func (m *ManagedDeviceCleanupRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceCleanupRulePlatformType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceCleanupRulePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCleanupRulePlatformType(val.(*DeviceCleanupRulePlatformType))
        }
        return nil
    }
    res["deviceInactivityBeforeRetirementInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceInactivityBeforeRetirementInDays(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Indicates the date and time when the device cleanup rule was last modified. This property is read-only.
// returns a *Time when successful
func (m *ManagedDeviceCleanupRule) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceCleanupRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceCleanupRulePlatformType() != nil {
        cast := (*m.GetDeviceCleanupRulePlatformType()).String()
        err = writer.WriteStringValue("deviceCleanupRulePlatformType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceInactivityBeforeRetirementInDays", m.GetDeviceInactivityBeforeRetirementInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Indicates the description for the device clean up rule.
func (m *ManagedDeviceCleanupRule) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceCleanupRulePlatformType sets the deviceCleanupRulePlatformType property value. Define the platform type for which the admin wants to create the device clean up rule
func (m *ManagedDeviceCleanupRule) SetDeviceCleanupRulePlatformType(value *DeviceCleanupRulePlatformType)() {
    err := m.GetBackingStore().Set("deviceCleanupRulePlatformType", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceInactivityBeforeRetirementInDays sets the deviceInactivityBeforeRetirementInDays property value. Indicates the number of days when the device has not contacted Intune. Valid values 0 to 2147483647
func (m *ManagedDeviceCleanupRule) SetDeviceInactivityBeforeRetirementInDays(value *int32)() {
    err := m.GetBackingStore().Set("deviceInactivityBeforeRetirementInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Indicates the display name of the device cleanup rule.
func (m *ManagedDeviceCleanupRule) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Indicates the date and time when the device cleanup rule was last modified. This property is read-only.
func (m *ManagedDeviceCleanupRule) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
type ManagedDeviceCleanupRuleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDeviceCleanupRulePlatformType()(*DeviceCleanupRulePlatformType)
    GetDeviceInactivityBeforeRetirementInDays()(*int32)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetDescription(value *string)()
    SetDeviceCleanupRulePlatformType(value *DeviceCleanupRulePlatformType)()
    SetDeviceInactivityBeforeRetirementInDays(value *int32)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
