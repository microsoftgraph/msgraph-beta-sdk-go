package ediscovery

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EstimateStatisticsOperationable 
type EstimateStatisticsOperationable interface {
    CaseOperationable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetIndexedItemCount()(*int64)
    GetIndexedItemsSize()(*int64)
    GetMailboxCount()(*int32)
    GetSiteCount()(*int32)
    GetSourceCollection()(SourceCollectionable)
    GetUnindexedItemCount()(*int64)
    GetUnindexedItemsSize()(*int64)
    SetIndexedItemCount(value *int64)()
    SetIndexedItemsSize(value *int64)()
    SetMailboxCount(value *int32)()
    SetSiteCount(value *int32)()
    SetSourceCollection(value SourceCollectionable)()
    SetUnindexedItemCount(value *int64)()
    SetUnindexedItemsSize(value *int64)()
}
