package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentable 
type AccessPackageAssignmentable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessPackage()(AccessPackageable)
    GetAccessPackageAssignmentPolicy()(AccessPackageAssignmentPolicyable)
    GetAccessPackageAssignmentRequests()([]AccessPackageAssignmentRequestable)
    GetAccessPackageAssignmentResourceRoles()([]AccessPackageAssignmentResourceRoleable)
    GetAccessPackageId()(*string)
    GetAssignmentPolicyId()(*string)
    GetAssignmentState()(*string)
    GetAssignmentStatus()(*string)
    GetCatalogId()(*string)
    GetExpiredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsExtended()(*bool)
    GetSchedule()(RequestScheduleable)
    GetTarget()(AccessPackageSubjectable)
    GetTargetId()(*string)
    SetAccessPackage(value AccessPackageable)()
    SetAccessPackageAssignmentPolicy(value AccessPackageAssignmentPolicyable)()
    SetAccessPackageAssignmentRequests(value []AccessPackageAssignmentRequestable)()
    SetAccessPackageAssignmentResourceRoles(value []AccessPackageAssignmentResourceRoleable)()
    SetAccessPackageId(value *string)()
    SetAssignmentPolicyId(value *string)()
    SetAssignmentState(value *string)()
    SetAssignmentStatus(value *string)()
    SetCatalogId(value *string)()
    SetExpiredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsExtended(value *bool)()
    SetSchedule(value RequestScheduleable)()
    SetTarget(value AccessPackageSubjectable)()
    SetTargetId(value *string)()
}
