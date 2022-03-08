package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LinkedResource_v2able 
type LinkedResource_v2able interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplicationName()(*string)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetWebUrl()(*string)
    SetApplicationName(value *string)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetWebUrl(value *string)()
}
