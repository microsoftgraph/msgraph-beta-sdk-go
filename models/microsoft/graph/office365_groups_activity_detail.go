package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Office365GroupsActivityDetail 
type Office365GroupsActivityDetail struct {
    Entity
    // The storage used of the group mailbox.
    exchangeMailboxStorageUsedInBytes *int64;
    // The number of items in the group mailbox.
    exchangeMailboxTotalItemCount *int64;
    // The number of email that the group mailbox received.
    exchangeReceivedEmailCount *int64;
    // The group external member count.
    externalMemberCount *int64;
    // The display name of the group.
    groupDisplayName *string;
    // The group id.
    groupId *string;
    // The group type. Possible values are: Public or Private.
    groupType *string;
    // Whether this user has been deleted or soft deleted.
    isDeleted *bool;
    // The last activity date for the following scenarios:  group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
    lastActivityDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The group member count.
    memberCount *int64;
    // The group owner principal name.
    ownerPrincipalName *string;
    // The number of days the report covers.
    reportPeriod *string;
    // The latest date of the content.
    reportRefreshDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The number of active files in SharePoint Group site.
    sharePointActiveFileCount *int64;
    // The storage used by SharePoint Group site.
    sharePointSiteStorageUsedInBytes *int64;
    // The total number of files in SharePoint Group site.
    sharePointTotalFileCount *int64;
    // The number of messages liked in Yammer groups.
    yammerLikedMessageCount *int64;
    // The number of messages posted to Yammer groups.
    yammerPostedMessageCount *int64;
    // The number of messages read in Yammer groups.
    yammerReadMessageCount *int64;
}
// NewOffice365GroupsActivityDetail instantiates a new Office365GroupsActivityDetail and sets the default values.
func NewOffice365GroupsActivityDetail()(*Office365GroupsActivityDetail) {
    m := &Office365GroupsActivityDetail{
        Entity: *NewEntity(),
    }
    return m
}
// GetExchangeMailboxStorageUsedInBytes gets the exchangeMailboxStorageUsedInBytes property value. The storage used of the group mailbox.
func (m *Office365GroupsActivityDetail) GetExchangeMailboxStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeMailboxStorageUsedInBytes
    }
}
// GetExchangeMailboxTotalItemCount gets the exchangeMailboxTotalItemCount property value. The number of items in the group mailbox.
func (m *Office365GroupsActivityDetail) GetExchangeMailboxTotalItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeMailboxTotalItemCount
    }
}
// GetExchangeReceivedEmailCount gets the exchangeReceivedEmailCount property value. The number of email that the group mailbox received.
func (m *Office365GroupsActivityDetail) GetExchangeReceivedEmailCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeReceivedEmailCount
    }
}
// GetExternalMemberCount gets the externalMemberCount property value. The group external member count.
func (m *Office365GroupsActivityDetail) GetExternalMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.externalMemberCount
    }
}
// GetGroupDisplayName gets the groupDisplayName property value. The display name of the group.
func (m *Office365GroupsActivityDetail) GetGroupDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupDisplayName
    }
}
// GetGroupId gets the groupId property value. The group id.
func (m *Office365GroupsActivityDetail) GetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupId
    }
}
// GetGroupType gets the groupType property value. The group type. Possible values are: Public or Private.
func (m *Office365GroupsActivityDetail) GetGroupType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupType
    }
}
// GetIsDeleted gets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *Office365GroupsActivityDetail) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// GetLastActivityDate gets the lastActivityDate property value. The last activity date for the following scenarios:  group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
func (m *Office365GroupsActivityDetail) GetLastActivityDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// GetMemberCount gets the memberCount property value. The group member count.
func (m *Office365GroupsActivityDetail) GetMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.memberCount
    }
}
// GetOwnerPrincipalName gets the ownerPrincipalName property value. The group owner principal name.
func (m *Office365GroupsActivityDetail) GetOwnerPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerPrincipalName
    }
}
// GetReportPeriod gets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityDetail) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityDetail) GetReportRefreshDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// GetSharePointActiveFileCount gets the sharePointActiveFileCount property value. The number of active files in SharePoint Group site.
func (m *Office365GroupsActivityDetail) GetSharePointActiveFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointActiveFileCount
    }
}
// GetSharePointSiteStorageUsedInBytes gets the sharePointSiteStorageUsedInBytes property value. The storage used by SharePoint Group site.
func (m *Office365GroupsActivityDetail) GetSharePointSiteStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointSiteStorageUsedInBytes
    }
}
// GetSharePointTotalFileCount gets the sharePointTotalFileCount property value. The total number of files in SharePoint Group site.
func (m *Office365GroupsActivityDetail) GetSharePointTotalFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointTotalFileCount
    }
}
// GetYammerLikedMessageCount gets the yammerLikedMessageCount property value. The number of messages liked in Yammer groups.
func (m *Office365GroupsActivityDetail) GetYammerLikedMessageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerLikedMessageCount
    }
}
// GetYammerPostedMessageCount gets the yammerPostedMessageCount property value. The number of messages posted to Yammer groups.
func (m *Office365GroupsActivityDetail) GetYammerPostedMessageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerPostedMessageCount
    }
}
// GetYammerReadMessageCount gets the yammerReadMessageCount property value. The number of messages read in Yammer groups.
func (m *Office365GroupsActivityDetail) GetYammerReadMessageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerReadMessageCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365GroupsActivityDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchangeMailboxStorageUsedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeMailboxStorageUsedInBytes(val)
        }
        return nil
    }
    res["exchangeMailboxTotalItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeMailboxTotalItemCount(val)
        }
        return nil
    }
    res["exchangeReceivedEmailCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeReceivedEmailCount(val)
        }
        return nil
    }
    res["externalMemberCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalMemberCount(val)
        }
        return nil
    }
    res["groupDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupDisplayName(val)
        }
        return nil
    }
    res["groupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["groupType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupType(val)
        }
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["lastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivityDate(val)
        }
        return nil
    }
    res["memberCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberCount(val)
        }
        return nil
    }
    res["ownerPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerPrincipalName(val)
        }
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportPeriod(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["sharePointActiveFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointActiveFileCount(val)
        }
        return nil
    }
    res["sharePointSiteStorageUsedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointSiteStorageUsedInBytes(val)
        }
        return nil
    }
    res["sharePointTotalFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointTotalFileCount(val)
        }
        return nil
    }
    res["yammerLikedMessageCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerLikedMessageCount(val)
        }
        return nil
    }
    res["yammerPostedMessageCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerPostedMessageCount(val)
        }
        return nil
    }
    res["yammerReadMessageCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerReadMessageCount(val)
        }
        return nil
    }
    return res
}
func (m *Office365GroupsActivityDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Office365GroupsActivityDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("exchangeMailboxStorageUsedInBytes", m.GetExchangeMailboxStorageUsedInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("exchangeMailboxTotalItemCount", m.GetExchangeMailboxTotalItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("exchangeReceivedEmailCount", m.GetExchangeReceivedEmailCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("externalMemberCount", m.GetExternalMemberCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupDisplayName", m.GetGroupDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupId", m.GetGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupType", m.GetGroupType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("lastActivityDate", m.GetLastActivityDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("memberCount", m.GetMemberCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerPrincipalName", m.GetOwnerPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportPeriod", m.GetReportPeriod())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("reportRefreshDate", m.GetReportRefreshDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sharePointActiveFileCount", m.GetSharePointActiveFileCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sharePointSiteStorageUsedInBytes", m.GetSharePointSiteStorageUsedInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sharePointTotalFileCount", m.GetSharePointTotalFileCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerLikedMessageCount", m.GetYammerLikedMessageCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerPostedMessageCount", m.GetYammerPostedMessageCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerReadMessageCount", m.GetYammerReadMessageCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExchangeMailboxStorageUsedInBytes sets the exchangeMailboxStorageUsedInBytes property value. The storage used of the group mailbox.
func (m *Office365GroupsActivityDetail) SetExchangeMailboxStorageUsedInBytes(value *int64)() {
    if m != nil {
        m.exchangeMailboxStorageUsedInBytes = value
    }
}
// SetExchangeMailboxTotalItemCount sets the exchangeMailboxTotalItemCount property value. The number of items in the group mailbox.
func (m *Office365GroupsActivityDetail) SetExchangeMailboxTotalItemCount(value *int64)() {
    if m != nil {
        m.exchangeMailboxTotalItemCount = value
    }
}
// SetExchangeReceivedEmailCount sets the exchangeReceivedEmailCount property value. The number of email that the group mailbox received.
func (m *Office365GroupsActivityDetail) SetExchangeReceivedEmailCount(value *int64)() {
    if m != nil {
        m.exchangeReceivedEmailCount = value
    }
}
// SetExternalMemberCount sets the externalMemberCount property value. The group external member count.
func (m *Office365GroupsActivityDetail) SetExternalMemberCount(value *int64)() {
    if m != nil {
        m.externalMemberCount = value
    }
}
// SetGroupDisplayName sets the groupDisplayName property value. The display name of the group.
func (m *Office365GroupsActivityDetail) SetGroupDisplayName(value *string)() {
    if m != nil {
        m.groupDisplayName = value
    }
}
// SetGroupId sets the groupId property value. The group id.
func (m *Office365GroupsActivityDetail) SetGroupId(value *string)() {
    if m != nil {
        m.groupId = value
    }
}
// SetGroupType sets the groupType property value. The group type. Possible values are: Public or Private.
func (m *Office365GroupsActivityDetail) SetGroupType(value *string)() {
    if m != nil {
        m.groupType = value
    }
}
// SetIsDeleted sets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *Office365GroupsActivityDetail) SetIsDeleted(value *bool)() {
    if m != nil {
        m.isDeleted = value
    }
}
// SetLastActivityDate sets the lastActivityDate property value. The last activity date for the following scenarios:  group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
func (m *Office365GroupsActivityDetail) SetLastActivityDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.lastActivityDate = value
    }
}
// SetMemberCount sets the memberCount property value. The group member count.
func (m *Office365GroupsActivityDetail) SetMemberCount(value *int64)() {
    if m != nil {
        m.memberCount = value
    }
}
// SetOwnerPrincipalName sets the ownerPrincipalName property value. The group owner principal name.
func (m *Office365GroupsActivityDetail) SetOwnerPrincipalName(value *string)() {
    if m != nil {
        m.ownerPrincipalName = value
    }
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityDetail) SetReportPeriod(value *string)() {
    if m != nil {
        m.reportPeriod = value
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityDetail) SetReportRefreshDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.reportRefreshDate = value
    }
}
// SetSharePointActiveFileCount sets the sharePointActiveFileCount property value. The number of active files in SharePoint Group site.
func (m *Office365GroupsActivityDetail) SetSharePointActiveFileCount(value *int64)() {
    if m != nil {
        m.sharePointActiveFileCount = value
    }
}
// SetSharePointSiteStorageUsedInBytes sets the sharePointSiteStorageUsedInBytes property value. The storage used by SharePoint Group site.
func (m *Office365GroupsActivityDetail) SetSharePointSiteStorageUsedInBytes(value *int64)() {
    if m != nil {
        m.sharePointSiteStorageUsedInBytes = value
    }
}
// SetSharePointTotalFileCount sets the sharePointTotalFileCount property value. The total number of files in SharePoint Group site.
func (m *Office365GroupsActivityDetail) SetSharePointTotalFileCount(value *int64)() {
    if m != nil {
        m.sharePointTotalFileCount = value
    }
}
// SetYammerLikedMessageCount sets the yammerLikedMessageCount property value. The number of messages liked in Yammer groups.
func (m *Office365GroupsActivityDetail) SetYammerLikedMessageCount(value *int64)() {
    if m != nil {
        m.yammerLikedMessageCount = value
    }
}
// SetYammerPostedMessageCount sets the yammerPostedMessageCount property value. The number of messages posted to Yammer groups.
func (m *Office365GroupsActivityDetail) SetYammerPostedMessageCount(value *int64)() {
    if m != nil {
        m.yammerPostedMessageCount = value
    }
}
// SetYammerReadMessageCount sets the yammerReadMessageCount property value. The number of messages read in Yammer groups.
func (m *Office365GroupsActivityDetail) SetYammerReadMessageCount(value *int64)() {
    if m != nil {
        m.yammerReadMessageCount = value
    }
}
