package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchRequestable 
type SearchRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAggregationFilters()([]string)
    GetAggregations()([]AggregationOptionable)
    GetContentSources()([]string)
    GetEnableTopResults()(*bool)
    GetEntityTypes()([]string)
    GetFields()([]string)
    GetFrom()(*int32)
    GetQuery()(SearchQueryable)
    GetQueryAlterationOptions()(SearchAlterationOptionsable)
    GetRegion()(*string)
    GetResultTemplateOptions()(ResultTemplateOptionable)
    GetSize()(*int32)
    GetSortProperties()([]SortPropertyable)
    GetStored_fields()([]string)
    GetTrimDuplicates()(*bool)
    SetAggregationFilters(value []string)()
    SetAggregations(value []AggregationOptionable)()
    SetContentSources(value []string)()
    SetEnableTopResults(value *bool)()
    SetEntityTypes(value []string)()
    SetFields(value []string)()
    SetFrom(value *int32)()
    SetQuery(value SearchQueryable)()
    SetQueryAlterationOptions(value SearchAlterationOptionsable)()
    SetRegion(value *string)()
    SetResultTemplateOptions(value ResultTemplateOptionable)()
    SetSize(value *int32)()
    SetSortProperties(value []SortPropertyable)()
    SetStored_fields(value []string)()
    SetTrimDuplicates(value *bool)()
}
