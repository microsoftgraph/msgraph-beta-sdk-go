package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProgramControl 
type ProgramControl struct {
    Entity
    // The controlId of the control, in particular the identifier of an access review. Required on create.
    controlId *string;
    // The programControlType identifies the type of program control - for example, a control linking to guest access reviews. Required on create.
    controlTypeId *string;
    // The creation date and time of the program control.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The name of the control.
    displayName *string;
    // The user who created the program control.
    owner *UserIdentity;
    // The program this control is part of.
    program *Program;
    // The programId of the program this control is a part of. Required on create.
    programId *string;
    // The resource, a group or an app, targeted by this program control's access review.
    resource *ProgramResource;
    // The life cycle status of the control.
    status *string;
}
// NewProgramControl instantiates a new programControl and sets the default values.
func NewProgramControl()(*ProgramControl) {
    m := &ProgramControl{
        Entity: *NewEntity(),
    }
    return m
}
// GetControlId gets the controlId property value. The controlId of the control, in particular the identifier of an access review. Required on create.
func (m *ProgramControl) GetControlId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.controlId
    }
}
// GetControlTypeId gets the controlTypeId property value. The programControlType identifies the type of program control - for example, a control linking to guest access reviews. Required on create.
func (m *ProgramControl) GetControlTypeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.controlTypeId
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The creation date and time of the program control.
func (m *ProgramControl) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDisplayName gets the displayName property value. The name of the control.
func (m *ProgramControl) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetOwner gets the owner property value. The user who created the program control.
func (m *ProgramControl) GetOwner()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
// GetProgram gets the program property value. The program this control is part of.
func (m *ProgramControl) GetProgram()(*Program) {
    if m == nil {
        return nil
    } else {
        return m.program
    }
}
// GetProgramId gets the programId property value. The programId of the program this control is a part of. Required on create.
func (m *ProgramControl) GetProgramId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.programId
    }
}
// GetResource gets the resource property value. The resource, a group or an app, targeted by this program control's access review.
func (m *ProgramControl) GetResource()(*ProgramResource) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
// GetStatus gets the status property value. The life cycle status of the control.
func (m *ProgramControl) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProgramControl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["controlId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControlId(val)
        }
        return nil
    }
    res["controlTypeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControlTypeId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(*UserIdentity))
        }
        return nil
    }
    res["program"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProgram() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgram(val.(*Program))
        }
        return nil
    }
    res["programId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgramId(val)
        }
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProgramResource() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(*ProgramResource))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *ProgramControl) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ProgramControl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.controlId = value
    }
}
// SetControlTypeId sets the controlTypeId property value. The programControlType identifies the type of program control - for example, a control linking to guest access reviews. Required on create.
func (m *ProgramControl) SetControlTypeId(value *string)() {
    if m != nil {
        m.controlTypeId = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The creation date and time of the program control.
func (m *ProgramControl) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDisplayName sets the displayName property value. The name of the control.
func (m *ProgramControl) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetOwner sets the owner property value. The user who created the program control.
func (m *ProgramControl) SetOwner(value *UserIdentity)() {
    if m != nil {
        m.owner = value
    }
}
// SetProgram sets the program property value. The program this control is part of.
func (m *ProgramControl) SetProgram(value *Program)() {
    if m != nil {
        m.program = value
    }
}
// SetProgramId sets the programId property value. The programId of the program this control is a part of. Required on create.
func (m *ProgramControl) SetProgramId(value *string)() {
    if m != nil {
        m.programId = value
    }
}
// SetResource sets the resource property value. The resource, a group or an app, targeted by this program control's access review.
func (m *ProgramControl) SetResource(value *ProgramResource)() {
    if m != nil {
        m.resource = value
    }
}
// SetStatus sets the status property value. The life cycle status of the control.
func (m *ProgramControl) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
