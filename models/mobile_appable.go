package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppable 
type MobileAppable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignments()([]MobileAppAssignmentable)
    GetCategories()([]MobileAppCategoryable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDependentAppCount()(*int32)
    GetDescription()(*string)
    GetDeveloper()(*string)
    GetDeviceStatuses()([]MobileAppInstallStatusable)
    GetDisplayName()(*string)
    GetInformationUrl()(*string)
    GetInstallSummary()(MobileAppInstallSummaryable)
    GetIsAssigned()(*bool)
    GetIsFeatured()(*bool)
    GetLargeIcon()(MimeContentable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetNotes()(*string)
    GetOwner()(*string)
    GetPrivacyInformationUrl()(*string)
    GetPublisher()(*string)
    GetPublishingState()(*MobileAppPublishingState)
    GetRelationships()([]MobileAppRelationshipable)
    GetRoleScopeTagIds()([]string)
    GetSupersededAppCount()(*int32)
    GetSupersedingAppCount()(*int32)
    GetUploadState()(*int32)
    GetUserStatuses()([]UserAppInstallStatusable)
    SetAssignments(value []MobileAppAssignmentable)()
    SetCategories(value []MobileAppCategoryable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDependentAppCount(value *int32)()
    SetDescription(value *string)()
    SetDeveloper(value *string)()
    SetDeviceStatuses(value []MobileAppInstallStatusable)()
    SetDisplayName(value *string)()
    SetInformationUrl(value *string)()
    SetInstallSummary(value MobileAppInstallSummaryable)()
    SetIsAssigned(value *bool)()
    SetIsFeatured(value *bool)()
    SetLargeIcon(value MimeContentable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetNotes(value *string)()
    SetOwner(value *string)()
    SetPrivacyInformationUrl(value *string)()
    SetPublisher(value *string)()
    SetPublishingState(value *MobileAppPublishingState)()
    SetRelationships(value []MobileAppRelationshipable)()
    SetRoleScopeTagIds(value []string)()
    SetSupersededAppCount(value *int32)()
    SetSupersedingAppCount(value *int32)()
    SetUploadState(value *int32)()
    SetUserStatuses(value []UserAppInstallStatusable)()
}
