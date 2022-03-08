package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementSettingComparisonable 
type DeviceManagementSettingComparisonable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetComparisonResult()(*DeviceManagementComparisonResult)
    GetCurrentValueJson()(*string)
    GetDefinitionId()(*string)
    GetDisplayName()(*string)
    GetId()(*string)
    GetNewValueJson()(*string)
    SetComparisonResult(value *DeviceManagementComparisonResult)()
    SetCurrentValueJson(value *string)()
    SetDefinitionId(value *string)()
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetNewValueJson(value *string)()
}
