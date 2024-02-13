package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type PrivilegedApproval struct {
    Entity
}
// NewPrivilegedApproval instantiates a new PrivilegedApproval and sets the default values.
func NewPrivilegedApproval()(*PrivilegedApproval) {
    m := &PrivilegedApproval{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedApprovalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePrivilegedApprovalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedApproval(), nil
}
// GetApprovalDuration gets the approvalDuration property value. The approvalDuration property
// returns a *ISODuration when successful
func (m *PrivilegedApproval) GetApprovalDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    val, err := m.GetBackingStore().Get("approvalDuration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    }
    return nil
}
// GetApprovalState gets the approvalState property value. The approvalState property
// returns a *ApprovalState when successful
func (m *PrivilegedApproval) GetApprovalState()(*ApprovalState) {
    val, err := m.GetBackingStore().Get("approvalState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ApprovalState)
    }
    return nil
}
// GetApprovalType gets the approvalType property value. The approvalType property
// returns a *string when successful
func (m *PrivilegedApproval) GetApprovalType()(*string) {
    val, err := m.GetBackingStore().Get("approvalType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetApproverReason gets the approverReason property value. The approverReason property
// returns a *string when successful
func (m *PrivilegedApproval) GetApproverReason()(*string) {
    val, err := m.GetBackingStore().Get("approverReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
// returns a *Time when successful
func (m *PrivilegedApproval) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("endDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *PrivilegedApproval) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvalDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalDuration(val)
        }
        return nil
    }
    res["approvalState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApprovalState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalState(val.(*ApprovalState))
        }
        return nil
    }
    res["approvalType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalType(val)
        }
        return nil
    }
    res["approverReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApproverReason(val)
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["request"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivilegedRoleAssignmentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequest(val.(PrivilegedRoleAssignmentRequestable))
        }
        return nil
    }
    res["requestorReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestorReason(val)
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
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
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
// GetRequest gets the request property value. The request property
// returns a PrivilegedRoleAssignmentRequestable when successful
func (m *PrivilegedApproval) GetRequest()(PrivilegedRoleAssignmentRequestable) {
    val, err := m.GetBackingStore().Get("request")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PrivilegedRoleAssignmentRequestable)
    }
    return nil
}
// GetRequestorReason gets the requestorReason property value. The requestorReason property
// returns a *string when successful
func (m *PrivilegedApproval) GetRequestorReason()(*string) {
    val, err := m.GetBackingStore().Get("requestorReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleId gets the roleId property value. The roleId property
// returns a *string when successful
func (m *PrivilegedApproval) GetRoleId()(*string) {
    val, err := m.GetBackingStore().Get("roleId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleInfo gets the roleInfo property value. The roleInfo property
// returns a PrivilegedRoleable when successful
func (m *PrivilegedApproval) GetRoleInfo()(PrivilegedRoleable) {
    val, err := m.GetBackingStore().Get("roleInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PrivilegedRoleable)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
// returns a *Time when successful
func (m *PrivilegedApproval) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetUserId gets the userId property value. The userId property
// returns a *string when successful
func (m *PrivilegedApproval) GetUserId()(*string) {
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
func (m *PrivilegedApproval) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteISODurationValue("approvalDuration", m.GetApprovalDuration())
        if err != nil {
            return err
        }
    }
    if m.GetApprovalState() != nil {
        cast := (*m.GetApprovalState()).String()
        err = writer.WriteStringValue("approvalState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("approvalType", m.GetApprovalType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("approverReason", m.GetApproverReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("request", m.GetRequest())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestorReason", m.GetRequestorReason())
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
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
// SetApprovalDuration sets the approvalDuration property value. The approvalDuration property
func (m *PrivilegedApproval) SetApprovalDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    err := m.GetBackingStore().Set("approvalDuration", value)
    if err != nil {
        panic(err)
    }
}
// SetApprovalState sets the approvalState property value. The approvalState property
func (m *PrivilegedApproval) SetApprovalState(value *ApprovalState)() {
    err := m.GetBackingStore().Set("approvalState", value)
    if err != nil {
        panic(err)
    }
}
// SetApprovalType sets the approvalType property value. The approvalType property
func (m *PrivilegedApproval) SetApprovalType(value *string)() {
    err := m.GetBackingStore().Set("approvalType", value)
    if err != nil {
        panic(err)
    }
}
// SetApproverReason sets the approverReason property value. The approverReason property
func (m *PrivilegedApproval) SetApproverReason(value *string)() {
    err := m.GetBackingStore().Set("approverReason", value)
    if err != nil {
        panic(err)
    }
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *PrivilegedApproval) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("endDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRequest sets the request property value. The request property
func (m *PrivilegedApproval) SetRequest(value PrivilegedRoleAssignmentRequestable)() {
    err := m.GetBackingStore().Set("request", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestorReason sets the requestorReason property value. The requestorReason property
func (m *PrivilegedApproval) SetRequestorReason(value *string)() {
    err := m.GetBackingStore().Set("requestorReason", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleId sets the roleId property value. The roleId property
func (m *PrivilegedApproval) SetRoleId(value *string)() {
    err := m.GetBackingStore().Set("roleId", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleInfo sets the roleInfo property value. The roleInfo property
func (m *PrivilegedApproval) SetRoleInfo(value PrivilegedRoleable)() {
    err := m.GetBackingStore().Set("roleInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *PrivilegedApproval) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The userId property
func (m *PrivilegedApproval) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
type PrivilegedApprovalable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApprovalDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetApprovalState()(*ApprovalState)
    GetApprovalType()(*string)
    GetApproverReason()(*string)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequest()(PrivilegedRoleAssignmentRequestable)
    GetRequestorReason()(*string)
    GetRoleId()(*string)
    GetRoleInfo()(PrivilegedRoleable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUserId()(*string)
    SetApprovalDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetApprovalState(value *ApprovalState)()
    SetApprovalType(value *string)()
    SetApproverReason(value *string)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequest(value PrivilegedRoleAssignmentRequestable)()
    SetRequestorReason(value *string)()
    SetRoleId(value *string)()
    SetRoleInfo(value PrivilegedRoleable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUserId(value *string)()
}
