package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Office365GroupsActivityDetailable 
type Office365GroupsActivityDetailable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExchangeMailboxStorageUsedInBytes()(*int64)
    GetExchangeMailboxTotalItemCount()(*int64)
    GetExchangeReceivedEmailCount()(*int64)
    GetExternalMemberCount()(*int64)
    GetGroupDisplayName()(*string)
    GetGroupId()(*string)
    GetGroupType()(*string)
    GetIsDeleted()(*bool)
    GetLastActivityDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetMemberCount()(*int64)
    GetOwnerPrincipalName()(*string)
    GetReportPeriod()(*string)
    GetReportRefreshDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)
    GetSharePointActiveFileCount()(*int64)
    GetSharePointSiteStorageUsedInBytes()(*int64)
    GetSharePointTotalFileCount()(*int64)
    GetYammerLikedMessageCount()(*int64)
    GetYammerPostedMessageCount()(*int64)
    GetYammerReadMessageCount()(*int64)
    SetExchangeMailboxStorageUsedInBytes(value *int64)()
    SetExchangeMailboxTotalItemCount(value *int64)()
    SetExchangeReceivedEmailCount(value *int64)()
    SetExternalMemberCount(value *int64)()
    SetGroupDisplayName(value *string)()
    SetGroupId(value *string)()
    SetGroupType(value *string)()
    SetIsDeleted(value *bool)()
    SetLastActivityDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetMemberCount(value *int64)()
    SetOwnerPrincipalName(value *string)()
    SetReportPeriod(value *string)()
    SetReportRefreshDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)()
    SetSharePointActiveFileCount(value *int64)()
    SetSharePointSiteStorageUsedInBytes(value *int64)()
    SetSharePointTotalFileCount(value *int64)()
    SetYammerLikedMessageCount(value *int64)()
    SetYammerPostedMessageCount(value *int64)()
    SetYammerReadMessageCount(value *int64)()
}
