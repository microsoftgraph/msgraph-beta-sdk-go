package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsDefenderApplicationControlSupplementalPolicyable 
type WindowsDefenderApplicationControlSupplementalPolicyable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignments()([]WindowsDefenderApplicationControlSupplementalPolicyAssignmentable)
    GetContent()([]byte)
    GetContentFileName()(*string)
    GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeploySummary()(WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable)
    GetDescription()(*string)
    GetDeviceStatuses()([]WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRoleScopeTagIds()([]string)
    GetVersion()(*string)
    SetAssignments(value []WindowsDefenderApplicationControlSupplementalPolicyAssignmentable)()
    SetContent(value []byte)()
    SetContentFileName(value *string)()
    SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeploySummary(value WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummaryable)()
    SetDescription(value *string)()
    SetDeviceStatuses(value []WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatusable)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRoleScopeTagIds(value []string)()
    SetVersion(value *string)()
}
