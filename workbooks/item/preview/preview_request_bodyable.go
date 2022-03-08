package preview

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PreviewRequestBodyable 
type PreviewRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowEdit()(*bool)
    GetChromeless()(*bool)
    GetPage()(*string)
    GetViewer()(*string)
    GetZoom()(*float64)
    SetAllowEdit(value *bool)()
    SetChromeless(value *bool)()
    SetPage(value *string)()
    SetViewer(value *string)()
    SetZoom(value *float64)()
}
