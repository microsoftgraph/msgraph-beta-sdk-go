package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupPolicyDefinitionFileable 
type GroupPolicyDefinitionFileable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDefinitions()([]GroupPolicyDefinitionable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetFileName()(*string)
    GetLanguageCodes()([]string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPolicyType()(*GroupPolicyType)
    GetRevision()(*string)
    GetTargetNamespace()(*string)
    GetTargetPrefix()(*string)
    SetDefinitions(value []GroupPolicyDefinitionable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetFileName(value *string)()
    SetLanguageCodes(value []string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPolicyType(value *GroupPolicyType)()
    SetRevision(value *string)()
    SetTargetNamespace(value *string)()
    SetTargetPrefix(value *string)()
}
