package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AdminConsentable 
type AdminConsentable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetShareAPNSData()(*AdminConsentState)
    GetShareUserExperienceAnalyticsData()(*AdminConsentState)
    SetShareAPNSData(value *AdminConsentState)()
    SetShareUserExperienceAnalyticsData(value *AdminConsentState)()
}
