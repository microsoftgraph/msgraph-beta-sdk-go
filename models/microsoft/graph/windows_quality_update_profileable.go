package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsQualityUpdateProfileable 
type WindowsQualityUpdateProfileable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignments()([]WindowsQualityUpdateProfileAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeployableContentDisplayName()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExpeditedUpdateSettings()(ExpeditedWindowsQualityUpdateSettingsable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReleaseDateDisplayName()(*string)
    GetRoleScopeTagIds()([]string)
    SetAssignments(value []WindowsQualityUpdateProfileAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeployableContentDisplayName(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExpeditedUpdateSettings(value ExpeditedWindowsQualityUpdateSettingsable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReleaseDateDisplayName(value *string)()
    SetRoleScopeTagIds(value []string)()
}
