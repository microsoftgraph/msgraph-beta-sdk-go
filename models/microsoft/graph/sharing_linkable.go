package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharingLinkable 
type SharingLinkable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetApplication()(Identityable)
    GetConfiguratorUrl()(*string)
    GetPreventsDownload()(*bool)
    GetScope()(*string)
    GetType()(*string)
    GetWebHtml()(*string)
    GetWebUrl()(*string)
    SetApplication(value Identityable)()
    SetConfiguratorUrl(value *string)()
    SetPreventsDownload(value *bool)()
    SetScope(value *string)()
    SetType(value *string)()
    SetWebHtml(value *string)()
    SetWebUrl(value *string)()
}
