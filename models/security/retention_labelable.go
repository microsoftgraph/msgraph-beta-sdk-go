package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// RetentionLabelable 
type RetentionLabelable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionAfterRetentionPeriod()(*ActionAfterRetentionPeriod)
    GetBehaviorDuringRetentionPeriod()(*BehaviorDuringRetentionPeriod)
    GetCreatedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDefaultRecordBehavior()(*DefaultRecordBehavior)
    GetDescriptionForAdmins()(*string)
    GetDescriptionForUsers()(*string)
    GetDisplayName()(*string)
    GetDispositionReviewStages()([]DispositionReviewStageable)
    GetIsInUse()(*bool)
    GetLabelToBeApplied()(*string)
    GetLastModifiedBy()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRetentionDuration()(RetentionDurationable)
    GetRetentionEventType()(RetentionEventTypeable)
    GetRetentionTrigger()(*RetentionTrigger)
    SetActionAfterRetentionPeriod(value *ActionAfterRetentionPeriod)()
    SetBehaviorDuringRetentionPeriod(value *BehaviorDuringRetentionPeriod)()
    SetCreatedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDefaultRecordBehavior(value *DefaultRecordBehavior)()
    SetDescriptionForAdmins(value *string)()
    SetDescriptionForUsers(value *string)()
    SetDisplayName(value *string)()
    SetDispositionReviewStages(value []DispositionReviewStageable)()
    SetIsInUse(value *bool)()
    SetLabelToBeApplied(value *string)()
    SetLastModifiedBy(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRetentionDuration(value RetentionDurationable)()
    SetRetentionEventType(value RetentionEventTypeable)()
    SetRetentionTrigger(value *RetentionTrigger)()
}
