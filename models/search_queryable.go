package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchQueryable 
type SearchQueryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetQuery_string()(SearchQueryStringable)
    GetQueryString()(*string)
    GetQueryTemplate()(*string)
    SetQuery_string(value SearchQueryStringable)()
    SetQueryString(value *string)()
    SetQueryTemplate(value *string)()
}
