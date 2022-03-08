package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProgramControlable 
type ProgramControlable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
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
