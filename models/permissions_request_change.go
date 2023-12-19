package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PermissionsRequestChange 
type PermissionsRequestChange struct {
    Entity
}
// NewPermissionsRequestChange instantiates a new permissionsRequestChange and sets the default values.
func NewPermissionsRequestChange()(*PermissionsRequestChange) {
    m := &PermissionsRequestChange{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePermissionsRequestChangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePermissionsRequestChangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPermissionsRequestChange(), nil
}
// GetActiveOccurrenceStatus gets the activeOccurrenceStatus property value. The status of the active occurence of the schedule if one exists. The possible values are: grantingFailed, granted, granting, revoked, revoking, revokingFailed, unknownFutureValue.
func (m *PermissionsRequestChange) GetActiveOccurrenceStatus()(*PermissionsRequestChange_activeOccurrenceStatus) {
    val, err := m.GetBackingStore().Get("activeOccurrenceStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PermissionsRequestChange_activeOccurrenceStatus)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PermissionsRequestChange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activeOccurrenceStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePermissionsRequestChange_activeOccurrenceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveOccurrenceStatus(val.(*PermissionsRequestChange_activeOccurrenceStatus))
        }
        return nil
    }
    res["modificationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModificationDateTime(val)
        }
        return nil
    }
    res["permissionsRequestId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionsRequestId(val)
        }
        return nil
    }
    res["statusDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStatusDetail)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatusDetail(val.(*StatusDetail))
        }
        return nil
    }
    res["ticketId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketId(val)
        }
        return nil
    }
    return res
}
// GetModificationDateTime gets the modificationDateTime property value. Time when the change occurred.
func (m *PermissionsRequestChange) GetModificationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("modificationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPermissionsRequestId gets the permissionsRequestId property value. The ID of the scheduledPermissionsRequest object.
func (m *PermissionsRequestChange) GetPermissionsRequestId()(*string) {
    val, err := m.GetBackingStore().Get("permissionsRequestId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatusDetail gets the statusDetail property value. The statusDetail property
func (m *PermissionsRequestChange) GetStatusDetail()(*StatusDetail) {
    val, err := m.GetBackingStore().Get("statusDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*StatusDetail)
    }
    return nil
}
// GetTicketId gets the ticketId property value. Represents the ticketing system identifier.
func (m *PermissionsRequestChange) GetTicketId()(*string) {
    val, err := m.GetBackingStore().Get("ticketId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PermissionsRequestChange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActiveOccurrenceStatus() != nil {
        cast := (*m.GetActiveOccurrenceStatus()).String()
        err = writer.WriteStringValue("activeOccurrenceStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modificationDateTime", m.GetModificationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionsRequestId", m.GetPermissionsRequestId())
        if err != nil {
            return err
        }
    }
    if m.GetStatusDetail() != nil {
        cast := (*m.GetStatusDetail()).String()
        err = writer.WriteStringValue("statusDetail", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ticketId", m.GetTicketId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveOccurrenceStatus sets the activeOccurrenceStatus property value. The status of the active occurence of the schedule if one exists. The possible values are: grantingFailed, granted, granting, revoked, revoking, revokingFailed, unknownFutureValue.
func (m *PermissionsRequestChange) SetActiveOccurrenceStatus(value *PermissionsRequestChange_activeOccurrenceStatus)() {
    err := m.GetBackingStore().Set("activeOccurrenceStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetModificationDateTime sets the modificationDateTime property value. Time when the change occurred.
func (m *PermissionsRequestChange) SetModificationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("modificationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionsRequestId sets the permissionsRequestId property value. The ID of the scheduledPermissionsRequest object.
func (m *PermissionsRequestChange) SetPermissionsRequestId(value *string)() {
    err := m.GetBackingStore().Set("permissionsRequestId", value)
    if err != nil {
        panic(err)
    }
}
// SetStatusDetail sets the statusDetail property value. The statusDetail property
func (m *PermissionsRequestChange) SetStatusDetail(value *StatusDetail)() {
    err := m.GetBackingStore().Set("statusDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetTicketId sets the ticketId property value. Represents the ticketing system identifier.
func (m *PermissionsRequestChange) SetTicketId(value *string)() {
    err := m.GetBackingStore().Set("ticketId", value)
    if err != nil {
        panic(err)
    }
}
// PermissionsRequestChangeable 
type PermissionsRequestChangeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveOccurrenceStatus()(*PermissionsRequestChange_activeOccurrenceStatus)
    GetModificationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPermissionsRequestId()(*string)
    GetStatusDetail()(*StatusDetail)
    GetTicketId()(*string)
    SetActiveOccurrenceStatus(value *PermissionsRequestChange_activeOccurrenceStatus)()
    SetModificationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPermissionsRequestId(value *string)()
    SetStatusDetail(value *StatusDetail)()
    SetTicketId(value *string)()
}
