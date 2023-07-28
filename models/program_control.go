package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProgramControl 
type ProgramControl struct {
    Entity
}
// NewProgramControl instantiates a new programControl and sets the default values.
func NewProgramControl()(*ProgramControl) {
    m := &ProgramControl{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProgramControlFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProgramControlFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProgramControl(), nil
}
// GetControlId gets the controlId property value. The controlId of the control, in particular the identifier of an access review. Required on create.
func (m *ProgramControl) GetControlId()(*string) {
    val, err := m.GetBackingStore().Get("controlId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetControlTypeId gets the controlTypeId property value. The programControlType identifies the type of program control - for example, a control linking to guest access reviews. Required on create.
func (m *ProgramControl) GetControlTypeId()(*string) {
    val, err := m.GetBackingStore().Get("controlTypeId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The creation date and time of the program control.
func (m *ProgramControl) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The name of the control.
func (m *ProgramControl) GetDisplayName()(*string) {
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
func (m *ProgramControl) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["controlId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControlId(val)
        }
        return nil
    }
    res["controlTypeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControlTypeId(val)
        }
        return nil
    }
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
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(UserIdentityable))
        }
        return nil
    }
    res["program"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProgramFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgram(val.(Programable))
        }
        return nil
    }
    res["programId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgramId(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProgramResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(ProgramResourceable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetOwner gets the owner property value. The user who created the program control.
func (m *ProgramControl) GetOwner()(UserIdentityable) {
    val, err := m.GetBackingStore().Get("owner")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserIdentityable)
    }
    return nil
}
// GetProgram gets the program property value. The program this control is part of.
func (m *ProgramControl) GetProgram()(Programable) {
    val, err := m.GetBackingStore().Get("program")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Programable)
    }
    return nil
}
// GetProgramId gets the programId property value. The programId of the program this control is a part of. Required on create.
func (m *ProgramControl) GetProgramId()(*string) {
    val, err := m.GetBackingStore().Get("programId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetResource gets the resource property value. The resource, a group or an app, targeted by this program control's access review.
func (m *ProgramControl) GetResource()(ProgramResourceable) {
    val, err := m.GetBackingStore().Get("resource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ProgramResourceable)
    }
    return nil
}
// GetStatus gets the status property value. The life cycle status of the control.
func (m *ProgramControl) GetStatus()(*string) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ProgramControl) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("controlId", m.GetControlId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("controlTypeId", m.GetControlTypeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("program", m.GetProgram())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("programId", m.GetProgramId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetControlId sets the controlId property value. The controlId of the control, in particular the identifier of an access review. Required on create.
func (m *ProgramControl) SetControlId(value *string)() {
    err := m.GetBackingStore().Set("controlId", value)
    if err != nil {
        panic(err)
    }
}
// SetControlTypeId sets the controlTypeId property value. The programControlType identifies the type of program control - for example, a control linking to guest access reviews. Required on create.
func (m *ProgramControl) SetControlTypeId(value *string)() {
    err := m.GetBackingStore().Set("controlTypeId", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The creation date and time of the program control.
func (m *ProgramControl) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name of the control.
func (m *ProgramControl) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetOwner sets the owner property value. The user who created the program control.
func (m *ProgramControl) SetOwner(value UserIdentityable)() {
    err := m.GetBackingStore().Set("owner", value)
    if err != nil {
        panic(err)
    }
}
// SetProgram sets the program property value. The program this control is part of.
func (m *ProgramControl) SetProgram(value Programable)() {
    err := m.GetBackingStore().Set("program", value)
    if err != nil {
        panic(err)
    }
}
// SetProgramId sets the programId property value. The programId of the program this control is a part of. Required on create.
func (m *ProgramControl) SetProgramId(value *string)() {
    err := m.GetBackingStore().Set("programId", value)
    if err != nil {
        panic(err)
    }
}
// SetResource sets the resource property value. The resource, a group or an app, targeted by this program control's access review.
func (m *ProgramControl) SetResource(value ProgramResourceable)() {
    err := m.GetBackingStore().Set("resource", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The life cycle status of the control.
func (m *ProgramControl) SetStatus(value *string)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// ProgramControlable 
type ProgramControlable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetControlId()(*string)
    GetControlTypeId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetOwner()(UserIdentityable)
    GetProgram()(Programable)
    GetProgramId()(*string)
    GetResource()(ProgramResourceable)
    GetStatus()(*string)
    SetControlId(value *string)()
    SetControlTypeId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetOwner(value UserIdentityable)()
    SetProgram(value Programable)()
    SetProgramId(value *string)()
    SetResource(value ProgramResourceable)()
    SetStatus(value *string)()
}
