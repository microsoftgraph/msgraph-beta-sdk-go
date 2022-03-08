package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AssignmentFilterEvaluationSummaryable 
type AssignmentFilterEvaluationSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignmentFilterDisplayName()(*string)
    GetAssignmentFilterId()(*string)
    GetAssignmentFilterLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAssignmentFilterPlatform()(*DevicePlatformType)
    GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType)
    GetAssignmentFilterTypeAndEvaluationResults()([]AssignmentFilterTypeAndEvaluationResultable)
    GetEvaluationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEvaluationResult()(*AssignmentFilterEvaluationResult)
    SetAssignmentFilterDisplayName(value *string)()
    SetAssignmentFilterId(value *string)()
    SetAssignmentFilterLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAssignmentFilterPlatform(value *DevicePlatformType)()
    SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)()
    SetAssignmentFilterTypeAndEvaluationResults(value []AssignmentFilterTypeAndEvaluationResultable)()
    SetEvaluationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEvaluationResult(value *AssignmentFilterEvaluationResult)()
}
