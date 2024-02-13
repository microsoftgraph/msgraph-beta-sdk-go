package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PrivilegedRoleSettings struct {
    Entity
}
// NewPrivilegedRoleSettings instantiates a new PrivilegedRoleSettings and sets the default values.
func NewPrivilegedRoleSettings()(*PrivilegedRoleSettings) {
    m := &PrivilegedRoleSettings{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedRoleSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrivilegedRoleSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedRoleSettings(), nil
}
// GetApprovalOnElevation gets the approvalOnElevation property value. The approvalOnElevation property
// returns a *bool when successful
func (m *PrivilegedRoleSettings) GetApprovalOnElevation()(*bool) {
    val, err := m.GetBackingStore().Get("approvalOnElevation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetApproverIds gets the approverIds property value. The approverIds property
// returns a []string when successful
func (m *PrivilegedRoleSettings) GetApproverIds()([]string) {
    val, err := m.GetBackingStore().Get("approverIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetElevationDuration gets the elevationDuration property value. The elevationDuration property
// returns a *ISODuration when successful
func (m *PrivilegedRoleSettings) GetElevationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("elevationDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PrivilegedRoleSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvalOnElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalOnElevation(val)
        }
        return nil
    }
    res["approverIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetApproverIds(res)
        }
        return nil
    }
    res["elevationDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetElevationDuration(val)
        }
        return nil
    }
    res["isMfaOnElevationConfigurable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMfaOnElevationConfigurable(val)
        }
        return nil
    }
    res["lastGlobalAdmin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastGlobalAdmin(val)
        }
        return nil
    }
    res["maxElavationDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxElavationDuration(val)
        }
        return nil
    }
    res["mfaOnElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMfaOnElevation(val)
        }
        return nil
    }
    res["minElevationDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinElevationDuration(val)
        }
        return nil
    }
    res["notificationToUserOnElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationToUserOnElevation(val)
        }
        return nil
    }
    res["ticketingInfoOnElevation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketingInfoOnElevation(val)
        }
        return nil
    }
    return res
}
// GetIsMfaOnElevationConfigurable gets the isMfaOnElevationConfigurable property value. The isMfaOnElevationConfigurable property
// returns a *bool when successful
func (m *PrivilegedRoleSettings) GetIsMfaOnElevationConfigurable()(*bool) {
    val, err := m.GetBackingStore().Get("isMfaOnElevationConfigurable")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastGlobalAdmin gets the lastGlobalAdmin property value. The lastGlobalAdmin property
// returns a *bool when successful
func (m *PrivilegedRoleSettings) GetLastGlobalAdmin()(*bool) {
    val, err := m.GetBackingStore().Get("lastGlobalAdmin")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMaxElavationDuration gets the maxElavationDuration property value. The maxElavationDuration property
// returns a *ISODuration when successful
func (m *PrivilegedRoleSettings) GetMaxElavationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("maxElavationDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetMfaOnElevation gets the mfaOnElevation property value. The mfaOnElevation property
// returns a *bool when successful
func (m *PrivilegedRoleSettings) GetMfaOnElevation()(*bool) {
    val, err := m.GetBackingStore().Get("mfaOnElevation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMinElevationDuration gets the minElevationDuration property value. The minElevationDuration property
// returns a *ISODuration when successful
func (m *PrivilegedRoleSettings) GetMinElevationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("minElevationDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetNotificationToUserOnElevation gets the notificationToUserOnElevation property value. The notificationToUserOnElevation property
// returns a *bool when successful
func (m *PrivilegedRoleSettings) GetNotificationToUserOnElevation()(*bool) {
    val, err := m.GetBackingStore().Get("notificationToUserOnElevation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTicketingInfoOnElevation gets the ticketingInfoOnElevation property value. The ticketingInfoOnElevation property
// returns a *bool when successful
func (m *PrivilegedRoleSettings) GetTicketingInfoOnElevation()(*bool) {
    val, err := m.GetBackingStore().Get("ticketingInfoOnElevation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrivilegedRoleSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("approvalOnElevation", m.GetApprovalOnElevation())
        if err != nil {
            return err
        }
    }
    if m.GetApproverIds() != nil {
        err = writer.WriteCollectionOfStringValues("approverIds", m.GetApproverIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("elevationDuration", m.GetElevationDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMfaOnElevationConfigurable", m.GetIsMfaOnElevationConfigurable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lastGlobalAdmin", m.GetLastGlobalAdmin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("maxElavationDuration", m.GetMaxElavationDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("mfaOnElevation", m.GetMfaOnElevation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("minElevationDuration", m.GetMinElevationDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("notificationToUserOnElevation", m.GetNotificationToUserOnElevation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("ticketingInfoOnElevation", m.GetTicketingInfoOnElevation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApprovalOnElevation sets the approvalOnElevation property value. The approvalOnElevation property
func (m *PrivilegedRoleSettings) SetApprovalOnElevation(value *bool)() {
    err := m.GetBackingStore().Set("approvalOnElevation", value)
    if err != nil {
        panic(err)
    }
}
// SetApproverIds sets the approverIds property value. The approverIds property
func (m *PrivilegedRoleSettings) SetApproverIds(value []string)() {
    err := m.GetBackingStore().Set("approverIds", value)
    if err != nil {
        panic(err)
    }
}
// SetElevationDuration sets the elevationDuration property value. The elevationDuration property
func (m *PrivilegedRoleSettings) SetElevationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("elevationDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMfaOnElevationConfigurable sets the isMfaOnElevationConfigurable property value. The isMfaOnElevationConfigurable property
func (m *PrivilegedRoleSettings) SetIsMfaOnElevationConfigurable(value *bool)() {
    err := m.GetBackingStore().Set("isMfaOnElevationConfigurable", value)
    if err != nil {
        panic(err)
    }
}
// SetLastGlobalAdmin sets the lastGlobalAdmin property value. The lastGlobalAdmin property
func (m *PrivilegedRoleSettings) SetLastGlobalAdmin(value *bool)() {
    err := m.GetBackingStore().Set("lastGlobalAdmin", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxElavationDuration sets the maxElavationDuration property value. The maxElavationDuration property
func (m *PrivilegedRoleSettings) SetMaxElavationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("maxElavationDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetMfaOnElevation sets the mfaOnElevation property value. The mfaOnElevation property
func (m *PrivilegedRoleSettings) SetMfaOnElevation(value *bool)() {
    err := m.GetBackingStore().Set("mfaOnElevation", value)
    if err != nil {
        panic(err)
    }
}
// SetMinElevationDuration sets the minElevationDuration property value. The minElevationDuration property
func (m *PrivilegedRoleSettings) SetMinElevationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("minElevationDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetNotificationToUserOnElevation sets the notificationToUserOnElevation property value. The notificationToUserOnElevation property
func (m *PrivilegedRoleSettings) SetNotificationToUserOnElevation(value *bool)() {
    err := m.GetBackingStore().Set("notificationToUserOnElevation", value)
    if err != nil {
        panic(err)
    }
}
// SetTicketingInfoOnElevation sets the ticketingInfoOnElevation property value. The ticketingInfoOnElevation property
func (m *PrivilegedRoleSettings) SetTicketingInfoOnElevation(value *bool)() {
    err := m.GetBackingStore().Set("ticketingInfoOnElevation", value)
    if err != nil {
        panic(err)
    }
}
type PrivilegedRoleSettingsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApprovalOnElevation()(*bool)
    GetApproverIds()([]string)
    GetElevationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetIsMfaOnElevationConfigurable()(*bool)
    GetLastGlobalAdmin()(*bool)
    GetMaxElavationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetMfaOnElevation()(*bool)
    GetMinElevationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetNotificationToUserOnElevation()(*bool)
    GetTicketingInfoOnElevation()(*bool)
    SetApprovalOnElevation(value *bool)()
    SetApproverIds(value []string)()
    SetElevationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetIsMfaOnElevationConfigurable(value *bool)()
    SetLastGlobalAdmin(value *bool)()
    SetMaxElavationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetMfaOnElevation(value *bool)()
    SetMinElevationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetNotificationToUserOnElevation(value *bool)()
    SetTicketingInfoOnElevation(value *bool)()
}
