package addtoreviewset

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/ediscovery"
)

// AddToReviewSetRequestBodyable 
type AddToReviewSetRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdditionalDataOptions()(*i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.AdditionalDataOptions)
    GetSourceCollection()(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollectionable)
    SetAdditionalDataOptions(value *i2756dc8c91c60abdde0aa43bf23ca1c0a6ac9b630146e89b7184e174a72c2de3.AdditionalDataOptions)()
    SetSourceCollection(value i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.SourceCollectionable)()
}
