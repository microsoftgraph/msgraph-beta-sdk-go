package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Office365GroupsActivityDetailable 
type Office365GroupsActivityDetailable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExchangeMailboxStorageUsedInBytes()(*int32)
    GetExchangeMailboxTotalItemCount()(*int32)
    GetExchangeReceivedEmailCount()(*int32)
    GetExternalMemberCount()(*int32)
    GetGroupDisplayName()(*string)
    GetGroupId()(*string)
    GetGroupType()(*string)
    GetIsDeleted()(*bool)
    GetLastActivityDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetMemberCount()(*int32)
    GetOwnerPrincipalName()(*string)
    GetReportPeriod()(*string)
    GetReportRefreshDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetSharePointActiveFileCount()(*int32)
    GetSharePointSiteStorageUsedInBytes()(*int32)
    GetSharePointTotalFileCount()(*int32)
    GetYammerLikedMessageCount()(*int32)
    GetYammerPostedMessageCount()(*int32)
    GetYammerReadMessageCount()(*int32)
    SetExchangeMailboxStorageUsedInBytes(value *int32)()
    SetExchangeMailboxTotalItemCount(value *int32)()
    SetExchangeReceivedEmailCount(value *int32)()
    SetExternalMemberCount(value *int32)()
    SetGroupDisplayName(value *string)()
    SetGroupId(value *string)()
    SetGroupType(value *string)()
    SetIsDeleted(value *bool)()
    SetLastActivityDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetMemberCount(value *int32)()
    SetOwnerPrincipalName(value *string)()
    SetReportPeriod(value *string)()
    SetReportRefreshDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetSharePointActiveFileCount(value *int32)()
    SetSharePointSiteStorageUsedInBytes(value *int32)()
    SetSharePointTotalFileCount(value *int32)()
    SetYammerLikedMessageCount(value *int32)()
    SetYammerPostedMessageCount(value *int32)()
    SetYammerReadMessageCount(value *int32)()
}
