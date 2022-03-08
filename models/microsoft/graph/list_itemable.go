package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ListItemable 
type ListItemable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    BaseItemable
    GetActivities()([]ItemActivityOLDable)
    GetAnalytics()(ItemAnalyticsable)
    GetContentType()(ContentTypeInfoable)
    GetDeleted()(Deletedable)
    GetDriveItem()(DriveItemable)
    GetFields()(FieldValueSetable)
    GetSharepointIds()(SharepointIdsable)
    GetVersions()([]ListItemVersionable)
    SetActivities(value []ItemActivityOLDable)()
    SetAnalytics(value ItemAnalyticsable)()
    SetContentType(value ContentTypeInfoable)()
    SetDeleted(value Deletedable)()
    SetDriveItem(value DriveItemable)()
    SetFields(value FieldValueSetable)()
    SetSharepointIds(value SharepointIdsable)()
    SetVersions(value []ListItemVersionable)()
}
