package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GroupPolicyDefinitionable 
type GroupPolicyDefinitionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCategory()(GroupPolicyCategoryable)
    GetCategoryPath()(*string)
    GetClassType()(*GroupPolicyDefinitionClassType)
    GetDefinitionFile()(GroupPolicyDefinitionFileable)
    GetDisplayName()(*string)
    GetExplainText()(*string)
    GetGroupPolicyCategoryId()(*string)
    GetHasRelatedDefinitions()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMinDeviceCspVersion()(*string)
    GetMinUserCspVersion()(*string)
    GetNextVersionDefinition()(GroupPolicyDefinitionable)
    GetPolicyType()(*GroupPolicyType)
    GetPresentations()([]GroupPolicyPresentationable)
    GetPreviousVersionDefinition()(GroupPolicyDefinitionable)
    GetSupportedOn()(*string)
    GetVersion()(*string)
    SetCategory(value GroupPolicyCategoryable)()
    SetCategoryPath(value *string)()
    SetClassType(value *GroupPolicyDefinitionClassType)()
    SetDefinitionFile(value GroupPolicyDefinitionFileable)()
    SetDisplayName(value *string)()
    SetExplainText(value *string)()
    SetGroupPolicyCategoryId(value *string)()
    SetHasRelatedDefinitions(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMinDeviceCspVersion(value *string)()
    SetMinUserCspVersion(value *string)()
    SetNextVersionDefinition(value GroupPolicyDefinitionable)()
    SetPolicyType(value *GroupPolicyType)()
    SetPresentations(value []GroupPolicyPresentationable)()
    SetPreviousVersionDefinition(value GroupPolicyDefinitionable)()
    SetSupportedOn(value *string)()
    SetVersion(value *string)()
}
