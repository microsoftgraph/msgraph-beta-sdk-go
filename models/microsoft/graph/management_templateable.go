package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// ManagementTemplateable 
type ManagementTemplateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCategory()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory)
    GetCreatedByUserId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetInformationLinks()([]ActionUrlable)
    GetLastActionByUserId()(*string)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagementTemplateCollections()([]ManagementTemplateCollectionable)
    GetManagementTemplateSteps()([]ManagementTemplateStepable)
    GetParameters()([]TemplateParameterable)
    GetPriority()(*int32)
    GetProvider()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementProvider)
    GetUserImpact()(*string)
    GetVersion()(*int32)
    GetWorkloadActions()([]WorkloadActionable)
    SetCategory(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory)()
    SetCreatedByUserId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetInformationLinks(value []ActionUrlable)()
    SetLastActionByUserId(value *string)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagementTemplateCollections(value []ManagementTemplateCollectionable)()
    SetManagementTemplateSteps(value []ManagementTemplateStepable)()
    SetParameters(value []TemplateParameterable)()
    SetPriority(value *int32)()
    SetProvider(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementProvider)()
    SetUserImpact(value *string)()
    SetVersion(value *int32)()
    SetWorkloadActions(value []WorkloadActionable)()
}
