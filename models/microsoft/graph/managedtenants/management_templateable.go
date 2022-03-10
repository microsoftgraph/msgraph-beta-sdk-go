package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ManagementTemplateable 
type ManagementTemplateable interface {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCategory()(*ManagementCategory)
    GetCreatedByUserId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetInformationLinks()([]i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ActionUrlable)
    GetLastActionByUserId()(*string)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagementTemplateCollections()([]ManagementTemplateCollectionable)
    GetManagementTemplateSteps()([]ManagementTemplateStepable)
    GetParameters()([]TemplateParameterable)
    GetPriority()(*int32)
    GetProvider()(*ManagementProvider)
    GetUserImpact()(*string)
    GetVersion()(*int32)
    GetWorkloadActions()([]WorkloadActionable)
    SetCategory(value *ManagementCategory)()
    SetCreatedByUserId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetInformationLinks(value []i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ActionUrlable)()
    SetLastActionByUserId(value *string)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagementTemplateCollections(value []ManagementTemplateCollectionable)()
    SetManagementTemplateSteps(value []ManagementTemplateStepable)()
    SetParameters(value []TemplateParameterable)()
    SetPriority(value *int32)()
    SetProvider(value *ManagementProvider)()
    SetUserImpact(value *string)()
    SetVersion(value *int32)()
    SetWorkloadActions(value []WorkloadActionable)()
}
