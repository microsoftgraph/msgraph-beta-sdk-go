package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DisplayTemplateable 
type DisplayTemplateable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetId()(*string)
    GetLayout()(Jsonable)
    GetPriority()(*int32)
    GetRules()([]PropertyRuleable)
    SetId(value *string)()
    SetLayout(value Jsonable)()
    SetPriority(value *int32)()
    SetRules(value []PropertyRuleable)()
}
