package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedRoleAssignment 
type PrivilegedRoleAssignment struct {
    Entity
}
// NewPrivilegedRoleAssignment instantiates a new privilegedRoleAssignment and sets the default values.
func NewPrivilegedRoleAssignment()(*PrivilegedRoleAssignment) {
    m := &PrivilegedRoleAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedRoleAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedRoleAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedRoleAssignment(), nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The UTC DateTime when the temporary privileged role assignment will be expired. For permanent role assignment, the value is null.
func (m *PrivilegedRoleAssignment) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedRoleAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["isElevated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsElevated(val)
        }
        return nil
    }
    res["resultMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultMessage(val)
        }
        return nil
    }
    res["roleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleId(val)
        }
        return nil
    }
    res["roleInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleInfo(val.(PrivilegedRoleable))
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
    return res
}
// GetIsElevated gets the isElevated property value. true if the role assignment is activated. false if the role assignment is deactivated.
func (m *PrivilegedRoleAssignment) GetIsElevated()(*bool) {
    val, err := m.GetBackingStore().Get("isElevated")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetResultMessage gets the resultMessage property value. Result message set by the service.
func (m *PrivilegedRoleAssignment) GetResultMessage()(*string) {
    val, err := m.GetBackingStore().Get("resultMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleId gets the roleId property value. Role identifier. In GUID string format.
func (m *PrivilegedRoleAssignment) GetRoleId()(*string) {
    val, err := m.GetBackingStore().Get("roleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleInfo gets the roleInfo property value. Read-only. Nullable. The associated role information.
func (m *PrivilegedRoleAssignment) GetRoleInfo()(PrivilegedRoleable) {
    val, err := m.GetBackingStore().Get("roleInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PrivilegedRoleable)
    }
    return nil
}
// GetUserId gets the userId property value. User identifier. In GUID string format.
func (m *PrivilegedRoleAssignment) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrivilegedRoleAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isElevated", m.GetIsElevated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resultMessage", m.GetResultMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleId", m.GetRoleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleInfo", m.GetRoleInfo())
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
    return nil
}
// SetExpirationDateTime sets the expirationDateTime property value. The UTC DateTime when the temporary privileged role assignment will be expired. For permanent role assignment, the value is null.
func (m *PrivilegedRoleAssignment) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIsElevated sets the isElevated property value. true if the role assignment is activated. false if the role assignment is deactivated.
func (m *PrivilegedRoleAssignment) SetIsElevated(value *bool)() {
    err := m.GetBackingStore().Set("isElevated", value)
    if err != nil {
        panic(err)
    }
}
// SetResultMessage sets the resultMessage property value. Result message set by the service.
func (m *PrivilegedRoleAssignment) SetResultMessage(value *string)() {
    err := m.GetBackingStore().Set("resultMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleId sets the roleId property value. Role identifier. In GUID string format.
func (m *PrivilegedRoleAssignment) SetRoleId(value *string)() {
    err := m.GetBackingStore().Set("roleId", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleInfo sets the roleInfo property value. Read-only. Nullable. The associated role information.
func (m *PrivilegedRoleAssignment) SetRoleInfo(value PrivilegedRoleable)() {
    err := m.GetBackingStore().Set("roleInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. User identifier. In GUID string format.
func (m *PrivilegedRoleAssignment) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// PrivilegedRoleAssignmentable 
type PrivilegedRoleAssignmentable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsElevated()(*bool)
    GetResultMessage()(*string)
    GetRoleId()(*string)
    GetRoleInfo()(PrivilegedRoleable)
    GetUserId()(*string)
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsElevated(value *bool)()
    SetResultMessage(value *string)()
    SetRoleId(value *string)()
    SetRoleInfo(value PrivilegedRoleable)()
    SetUserId(value *string)()
}
