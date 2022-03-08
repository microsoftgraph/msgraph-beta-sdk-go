package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InformationProtectionPolicySettingable 
type InformationProtectionPolicySettingable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDefaultLabelId()(*string)
    GetIsDowngradeJustificationRequired()(*bool)
    GetIsMandatory()(*bool)
    GetMoreInfoUrl()(*string)
    SetDefaultLabelId(value *string)()
    SetIsDowngradeJustificationRequired(value *bool)()
    SetIsMandatory(value *bool)()
    SetMoreInfoUrl(value *string)()
}
