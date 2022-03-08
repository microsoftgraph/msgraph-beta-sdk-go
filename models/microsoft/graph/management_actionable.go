package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/managedtenants"
)

// ManagementActionable 
type ManagementActionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCategory()(*i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetReferenceTemplateId()(*string)
    GetReferenceTemplateVersion()(*int32)
    GetWorkloadActions()([]WorkloadActionable)
    SetCategory(value *i5c2592132064055aae424492b066923068e6d9a29d4565707b3591c21983fe01.ManagementCategory)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetReferenceTemplateId(value *string)()
    SetReferenceTemplateVersion(value *int32)()
    SetWorkloadActions(value []WorkloadActionable)()
}
