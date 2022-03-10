package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcGalleryImageable 
type CloudPcGalleryImageable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDisplayName()(*string)
    GetEndDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetExpirationDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetOffer()(*string)
    GetOfferDisplayName()(*string)
    GetPublisher()(*string)
    GetRecommendedSku()(*string)
    GetSizeInGB()(*int32)
    GetSku()(*string)
    GetSkuDisplayName()(*string)
    GetStartDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetStatus()(*CloudPcGalleryImageStatus)
    SetDisplayName(value *string)()
    SetEndDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetExpirationDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetOffer(value *string)()
    SetOfferDisplayName(value *string)()
    SetPublisher(value *string)()
    SetRecommendedSku(value *string)()
    SetSizeInGB(value *int32)()
    SetSku(value *string)()
    SetSkuDisplayName(value *string)()
    SetStartDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetStatus(value *CloudPcGalleryImageStatus)()
}
