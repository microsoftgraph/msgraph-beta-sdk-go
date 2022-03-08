package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamsAppIconable 
type TeamsAppIconable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetHostedContent()(TeamworkHostedContentable)
    GetWebUrl()(*string)
    SetHostedContent(value TeamworkHostedContentable)()
    SetWebUrl(value *string)()
}
