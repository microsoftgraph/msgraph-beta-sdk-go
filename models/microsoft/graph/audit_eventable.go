package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuditEventable 
type AuditEventable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetActivity()(*string)
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetActivityOperationType()(*string)
    GetActivityResult()(*string)
    GetActivityType()(*string)
    GetActor()(AuditActorable)
    GetCategory()(*string)
    GetComponentName()(*string)
    GetCorrelationId()(*string)
    GetDisplayName()(*string)
    GetResources()([]AuditResourceable)
    SetActivity(value *string)()
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetActivityOperationType(value *string)()
    SetActivityResult(value *string)()
    SetActivityType(value *string)()
    SetActor(value AuditActorable)()
    SetCategory(value *string)()
    SetComponentName(value *string)()
    SetCorrelationId(value *string)()
    SetDisplayName(value *string)()
    SetResources(value []AuditResourceable)()
}
