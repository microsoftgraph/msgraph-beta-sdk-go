package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SitePageable 
type SitePageable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentType()(ContentTypeInfoable)
    GetPageLayoutType()(*string)
    GetPublishingState()(PublicationFacetable)
    GetTitle()(*string)
    GetWebParts()([]WebPartable)
    SetContentType(value ContentTypeInfoable)()
    SetPageLayoutType(value *string)()
    SetPublishingState(value PublicationFacetable)()
    SetTitle(value *string)()
    SetWebParts(value []WebPartable)()
}
