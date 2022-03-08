package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserExperienceAnalyticsNotAutopilotReadyDeviceable 
type UserExperienceAnalyticsNotAutopilotReadyDeviceable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAutoPilotProfileAssigned()(*bool)
    GetAutoPilotRegistered()(*bool)
    GetAzureAdJoinType()(*string)
    GetAzureAdRegistered()(*bool)
    GetDeviceName()(*string)
    GetManagedBy()(*string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetSerialNumber()(*string)
    SetAutoPilotProfileAssigned(value *bool)()
    SetAutoPilotRegistered(value *bool)()
    SetAzureAdJoinType(value *string)()
    SetAzureAdRegistered(value *bool)()
    SetDeviceName(value *string)()
    SetManagedBy(value *string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetSerialNumber(value *string)()
}
