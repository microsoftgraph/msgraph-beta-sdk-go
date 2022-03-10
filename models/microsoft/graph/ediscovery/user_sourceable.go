package ediscovery

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserSourceable 
type UserSourceable interface {
    DataSourceable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetEmail()(*string)
    GetIncludedSources()(*SourceType)
    GetSiteWebUrl()(*string)
    SetEmail(value *string)()
    SetIncludedSources(value *SourceType)()
    SetSiteWebUrl(value *string)()
}
