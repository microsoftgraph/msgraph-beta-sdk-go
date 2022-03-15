package ediscovery

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EstimateStatisticsOperationable 
type EstimateStatisticsOperationable interface {
    CaseOperationable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetIndexedItemCount()(*int32)
    GetIndexedItemsSize()(*int32)
    GetMailboxCount()(*int32)
    GetSiteCount()(*int32)
    GetSourceCollection()(SourceCollectionable)
    GetUnindexedItemCount()(*int32)
    GetUnindexedItemsSize()(*int32)
    SetIndexedItemCount(value *int32)()
    SetIndexedItemsSize(value *int32)()
    SetMailboxCount(value *int32)()
    SetSiteCount(value *int32)()
    SetSourceCollection(value SourceCollectionable)()
    SetUnindexedItemCount(value *int32)()
    SetUnindexedItemsSize(value *int32)()
}
