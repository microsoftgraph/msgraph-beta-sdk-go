package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentPolicyable 
type AccessPackageAssignmentPolicyable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessPackage()(AccessPackageable)
    GetAccessPackageCatalog()(AccessPackageCatalogable)
    GetAccessPackageId()(*string)
    GetAccessReviewSettings()(AssignmentReviewSettingsable)
    GetCanExtend()(*bool)
    GetCreatedBy()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomExtensionHandlers()([]CustomExtensionHandlerable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDurationInDays()(*int32)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetModifiedBy()(*string)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetQuestions()([]AccessPackageQuestionable)
    GetRequestApprovalSettings()(ApprovalSettingsable)
    GetRequestorSettings()(RequestorSettingsable)
    SetAccessPackage(value AccessPackageable)()
    SetAccessPackageCatalog(value AccessPackageCatalogable)()
    SetAccessPackageId(value *string)()
    SetAccessReviewSettings(value AssignmentReviewSettingsable)()
    SetCanExtend(value *bool)()
    SetCreatedBy(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomExtensionHandlers(value []CustomExtensionHandlerable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDurationInDays(value *int32)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetModifiedBy(value *string)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetQuestions(value []AccessPackageQuestionable)()
    SetRequestApprovalSettings(value ApprovalSettingsable)()
    SetRequestorSettings(value RequestorSettingsable)()
}
