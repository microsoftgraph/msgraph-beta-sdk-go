package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ProgramControl struct {
    Entity
    controlId *string;
    controlTypeId *string;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    displayName *string;
    owner *UserIdentity;
    program *Program;
    programId *string;
    resource *ProgramResource;
    status *string;
}
func NewProgramControl()(*ProgramControl) {
    m := &ProgramControl{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ProgramControl) GetControlId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.controlId
    }
}
func (m *ProgramControl) GetControlTypeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.controlTypeId
    }
}
func (m *ProgramControl) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *ProgramControl) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *ProgramControl) GetOwner()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
func (m *ProgramControl) GetProgram()(*Program) {
    if m == nil {
        return nil
    } else {
        return m.program
    }
}
func (m *ProgramControl) GetProgramId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.programId
    }
}
func (m *ProgramControl) GetResource()(*ProgramResource) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
func (m *ProgramControl) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *ProgramControl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["controlId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetControlId(val)
        return nil
    }
    res["controlTypeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetControlTypeId(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        m.SetOwner(val.(*UserIdentity))
        return nil
    }
    res["program"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProgram() })
        if err != nil {
            return err
        }
        m.SetProgram(val.(*Program))
        return nil
    }
    res["programId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProgramId(val)
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProgramResource() })
        if err != nil {
            return err
        }
        m.SetResource(val.(*ProgramResource))
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    return res
}
func (m *ProgramControl) IsNil()(bool) {
    return m == nil
}
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
func (m *ProgramControl) SetControlId(value *string)() {
    m.controlId = value
}
func (m *ProgramControl) SetControlTypeId(value *string)() {
    m.controlTypeId = value
}
func (m *ProgramControl) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *ProgramControl) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *ProgramControl) SetOwner(value *UserIdentity)() {
    m.owner = value
}
func (m *ProgramControl) SetProgram(value *Program)() {
    m.program = value
}
func (m *ProgramControl) SetProgramId(value *string)() {
    m.programId = value
}
func (m *ProgramControl) SetResource(value *ProgramResource)() {
    m.resource = value
}
func (m *ProgramControl) SetStatus(value *string)() {
    m.status = value
}
