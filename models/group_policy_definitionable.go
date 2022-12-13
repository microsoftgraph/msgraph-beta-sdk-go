package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyDefinitionable 
type GroupPolicyDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(GroupPolicyCategoryable)
    GetCategoryPath()(*string)
    GetClassType()(*GroupPolicyDefinitionClassType)
    GetDefinitionFile()(GroupPolicyDefinitionFileable)
    GetDisplayName()(*string)
    GetExplainText()(*string)
    GetGroupPolicyCategoryId()(*UUID)
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
    SetGroupPolicyCategoryId(value *UUID)()
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
