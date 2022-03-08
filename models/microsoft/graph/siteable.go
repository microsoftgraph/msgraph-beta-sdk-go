package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Siteable 
type Siteable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    BaseItemable
    GetAnalytics()(ItemAnalyticsable)
    GetColumns()([]ColumnDefinitionable)
    GetContentTypes()([]ContentTypeable)
    GetDeleted()(Deletedable)
    GetDisplayName()(*string)
    GetDrive()(Driveable)
    GetDrives()([]Driveable)
    GetExternalColumns()([]ColumnDefinitionable)
    GetItems()([]BaseItemable)
    GetLists()([]Listable)
    GetOnenote()(Onenoteable)
    GetOperations()([]RichLongRunningOperationable)
    GetPages()([]SitePageable)
    GetPermissions()([]Permissionable)
    GetRoot()(Rootable)
    GetSettings()(SiteSettingsable)
    GetSharepointIds()(SharepointIdsable)
    GetSiteCollection()(SiteCollectionable)
    GetSites()([]Siteable)
    GetTermStore()(Storeable)
    SetAnalytics(value ItemAnalyticsable)()
    SetColumns(value []ColumnDefinitionable)()
    SetContentTypes(value []ContentTypeable)()
    SetDeleted(value Deletedable)()
    SetDisplayName(value *string)()
    SetDrive(value Driveable)()
    SetDrives(value []Driveable)()
    SetExternalColumns(value []ColumnDefinitionable)()
    SetItems(value []BaseItemable)()
    SetLists(value []Listable)()
    SetOnenote(value Onenoteable)()
    SetOperations(value []RichLongRunningOperationable)()
    SetPages(value []SitePageable)()
    SetPermissions(value []Permissionable)()
    SetRoot(value Rootable)()
    SetSettings(value SiteSettingsable)()
    SetSharepointIds(value SharepointIdsable)()
    SetSiteCollection(value SiteCollectionable)()
    SetSites(value []Siteable)()
    SetTermStore(value Storeable)()
}
