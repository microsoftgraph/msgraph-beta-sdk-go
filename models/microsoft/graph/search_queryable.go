package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchQueryable 
type SearchQueryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetQuery_string()(SearchQueryStringable)
    GetQueryString()(*string)
    GetQueryTemplate()(*string)
    SetQuery_string(value SearchQueryStringable)()
    SetQueryString(value *string)()
    SetQueryTemplate(value *string)()
}
