package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GovernancePolicyTemplateable 
type GovernancePolicyTemplateable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetPolicy()(GovernancePolicyable)
    GetSettings()(BusinessFlowSettingsable)
    SetDisplayName(value *string)()
    SetPolicy(value GovernancePolicyable)()
    SetSettings(value BusinessFlowSettingsable)()
}
