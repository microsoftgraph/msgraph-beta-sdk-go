package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppLogCollectionDownloadDetailsable 
type AppLogCollectionDownloadDetailsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppLogDecryptionAlgorithm()(*AppLogDecryptionAlgorithm)
    GetDecryptionKey()(*string)
    GetDownloadUrl()(*string)
    SetAppLogDecryptionAlgorithm(value *AppLogDecryptionAlgorithm)()
    SetDecryptionKey(value *string)()
    SetDownloadUrl(value *string)()
}
