package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupPolicyCategoryable 
type GroupPolicyCategoryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetChildren()([]GroupPolicyCategoryable)
    GetDefinitionFile()(GroupPolicyDefinitionFileable)
    GetDefinitions()([]GroupPolicyDefinitionable)
    GetDisplayName()(*string)
    GetIsRoot()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetParent()(GroupPolicyCategoryable)
    SetChildren(value []GroupPolicyCategoryable)()
    SetDefinitionFile(value GroupPolicyDefinitionFileable)()
    SetDefinitions(value []GroupPolicyDefinitionable)()
    SetDisplayName(value *string)()
    SetIsRoot(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetParent(value GroupPolicyCategoryable)()
}
