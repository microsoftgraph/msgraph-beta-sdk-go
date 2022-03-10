package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementSettingDefinitionable 
type DeviceManagementSettingDefinitionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConstraints()([]DeviceManagementConstraintable)
    GetDependencies()([]DeviceManagementSettingDependencyable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDocumentationUrl()(*string)
    GetHeaderSubtitle()(*string)
    GetHeaderTitle()(*string)
    GetIsTopLevel()(*bool)
    GetKeywords()([]string)
    GetPlaceholderText()(*string)
    GetValueType()(*DeviceManangementIntentValueType)
    SetConstraints(value []DeviceManagementConstraintable)()
    SetDependencies(value []DeviceManagementSettingDependencyable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDocumentationUrl(value *string)()
    SetHeaderSubtitle(value *string)()
    SetHeaderTitle(value *string)()
    SetIsTopLevel(value *bool)()
    SetKeywords(value []string)()
    SetPlaceholderText(value *string)()
    SetValueType(value *DeviceManangementIntentValueType)()
}
