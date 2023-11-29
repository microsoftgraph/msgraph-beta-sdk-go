package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScheduledPermissionsRequest 
type ScheduledPermissionsRequest struct {
    Entity
}
// NewScheduledPermissionsRequest instantiates a new scheduledPermissionsRequest and sets the default values.
func NewScheduledPermissionsRequest()(*ScheduledPermissionsRequest) {
    m := &ScheduledPermissionsRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateScheduledPermissionsRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateScheduledPermissionsRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScheduledPermissionsRequest(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Defines when the identity created the request.
func (m *ScheduledPermissionsRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ScheduledPermissionsRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["justification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["requestedPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionsDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedPermissions(val.(PermissionsDefinitionable))
        }
        return nil
    }
    res["scheduleInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleInfo(val.(RequestScheduleable))
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
    res["ticketInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTicketInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketInfo(val.(TicketInfoable))
        }
        return nil
    }
    return res
}
// GetJustification gets the justification property value. The identity's justification for the request.
func (m *ScheduledPermissionsRequest) GetJustification()(*string) {
    val, err := m.GetBackingStore().Get("justification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNotes gets the notes property value. Additional context for the permissions request.
func (m *ScheduledPermissionsRequest) GetNotes()(*string) {
    val, err := m.GetBackingStore().Get("notes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestedPermissions gets the requestedPermissions property value. The requestedPermissions property
func (m *ScheduledPermissionsRequest) GetRequestedPermissions()(PermissionsDefinitionable) {
    val, err := m.GetBackingStore().Get("requestedPermissions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsDefinitionable)
    }
    return nil
}
// GetScheduleInfo gets the scheduleInfo property value. When to assign the requested permissions.
func (m *ScheduledPermissionsRequest) GetScheduleInfo()(RequestScheduleable) {
    val, err := m.GetBackingStore().Get("scheduleInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RequestScheduleable)
    }
    return nil
}
// GetStatusDetail gets the statusDetail property value. The statusDetail property
func (m *ScheduledPermissionsRequest) GetStatusDetail()(*StatusDetail) {
    val, err := m.GetBackingStore().Get("statusDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*StatusDetail)
    }
    return nil
}
// GetTicketInfo gets the ticketInfo property value. Ticketing-related metadata that you can use to correlate to the request.
func (m *ScheduledPermissionsRequest) GetTicketInfo()(TicketInfoable) {
    val, err := m.GetBackingStore().Get("ticketInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TicketInfoable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ScheduledPermissionsRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requestedPermissions", m.GetRequestedPermissions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scheduleInfo", m.GetScheduleInfo())
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
        err = writer.WriteObjectValue("ticketInfo", m.GetTicketInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Defines when the identity created the request.
func (m *ScheduledPermissionsRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetJustification sets the justification property value. The identity's justification for the request.
func (m *ScheduledPermissionsRequest) SetJustification(value *string)() {
    err := m.GetBackingStore().Set("justification", value)
    if err != nil {
        panic(err)
    }
}
// SetNotes sets the notes property value. Additional context for the permissions request.
func (m *ScheduledPermissionsRequest) SetNotes(value *string)() {
    err := m.GetBackingStore().Set("notes", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestedPermissions sets the requestedPermissions property value. The requestedPermissions property
func (m *ScheduledPermissionsRequest) SetRequestedPermissions(value PermissionsDefinitionable)() {
    err := m.GetBackingStore().Set("requestedPermissions", value)
    if err != nil {
        panic(err)
    }
}
// SetScheduleInfo sets the scheduleInfo property value. When to assign the requested permissions.
func (m *ScheduledPermissionsRequest) SetScheduleInfo(value RequestScheduleable)() {
    err := m.GetBackingStore().Set("scheduleInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetStatusDetail sets the statusDetail property value. The statusDetail property
func (m *ScheduledPermissionsRequest) SetStatusDetail(value *StatusDetail)() {
    err := m.GetBackingStore().Set("statusDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetTicketInfo sets the ticketInfo property value. Ticketing-related metadata that you can use to correlate to the request.
func (m *ScheduledPermissionsRequest) SetTicketInfo(value TicketInfoable)() {
    err := m.GetBackingStore().Set("ticketInfo", value)
    if err != nil {
        panic(err)
    }
}
// ScheduledPermissionsRequestable 
type ScheduledPermissionsRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetJustification()(*string)
    GetNotes()(*string)
    GetRequestedPermissions()(PermissionsDefinitionable)
    GetScheduleInfo()(RequestScheduleable)
    GetStatusDetail()(*StatusDetail)
    GetTicketInfo()(TicketInfoable)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetJustification(value *string)()
    SetNotes(value *string)()
    SetRequestedPermissions(value PermissionsDefinitionable)()
    SetScheduleInfo(value RequestScheduleable)()
    SetStatusDetail(value *StatusDetail)()
    SetTicketInfo(value TicketInfoable)()
}
