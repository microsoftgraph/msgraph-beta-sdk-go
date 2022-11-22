package managedtenants

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ManagementTemplateStepTenantSummaryable 
type ManagementTemplateStepTenantSummaryable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedTenantsCount()(*int32)
    GetCompliantTenantsCount()(*int32)
    GetCreatedByUserId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDismissedTenantsCount()(*int32)
    GetIneligibleTenantsCount()(*int32)
    GetLastActionByUserId()(*string)
    GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagementTemplateCollectionDisplayName()(*string)
    GetManagementTemplateCollectionId()(*string)
    GetManagementTemplateDisplayName()(*string)
    GetManagementTemplateId()(*string)
    GetManagementTemplateStepDisplayName()(*string)
    GetManagementTemplateStepId()(*string)
    GetNotCompliantTenantsCount()(*int32)
    SetAssignedTenantsCount(value *int32)()
    SetCompliantTenantsCount(value *int32)()
    SetCreatedByUserId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDismissedTenantsCount(value *int32)()
    SetIneligibleTenantsCount(value *int32)()
    SetLastActionByUserId(value *string)()
    SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagementTemplateCollectionDisplayName(value *string)()
    SetManagementTemplateCollectionId(value *string)()
    SetManagementTemplateDisplayName(value *string)()
    SetManagementTemplateId(value *string)()
    SetManagementTemplateStepDisplayName(value *string)()
    SetManagementTemplateStepId(value *string)()
    SetNotCompliantTenantsCount(value *int32)()
}
