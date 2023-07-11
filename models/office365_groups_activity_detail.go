package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Office365GroupsActivityDetail 
type Office365GroupsActivityDetail struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewOffice365GroupsActivityDetail instantiates a new office365GroupsActivityDetail and sets the default values.
func NewOffice365GroupsActivityDetail()(*Office365GroupsActivityDetail) {
    m := &Office365GroupsActivityDetail{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOffice365GroupsActivityDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOffice365GroupsActivityDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOffice365GroupsActivityDetail(), nil
}
// GetExchangeMailboxStorageUsedInBytes gets the exchangeMailboxStorageUsedInBytes property value. The storage used of the group mailbox.
func (m *Office365GroupsActivityDetail) GetExchangeMailboxStorageUsedInBytes()(*int64) {
    val, err := m.GetBackingStore().Get("exchangeMailboxStorageUsedInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetExchangeMailboxTotalItemCount gets the exchangeMailboxTotalItemCount property value. The number of items in the group mailbox.
func (m *Office365GroupsActivityDetail) GetExchangeMailboxTotalItemCount()(*int64) {
    val, err := m.GetBackingStore().Get("exchangeMailboxTotalItemCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetExchangeReceivedEmailCount gets the exchangeReceivedEmailCount property value. The number of email that the group mailbox received.
func (m *Office365GroupsActivityDetail) GetExchangeReceivedEmailCount()(*int64) {
    val, err := m.GetBackingStore().Get("exchangeReceivedEmailCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetExternalMemberCount gets the externalMemberCount property value. The group external member count.
func (m *Office365GroupsActivityDetail) GetExternalMemberCount()(*int64) {
    val, err := m.GetBackingStore().Get("externalMemberCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Office365GroupsActivityDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchangeMailboxStorageUsedInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeMailboxStorageUsedInBytes(val)
        }
        return nil
    }
    res["exchangeMailboxTotalItemCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeMailboxTotalItemCount(val)
        }
        return nil
    }
    res["exchangeReceivedEmailCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeReceivedEmailCount(val)
        }
        return nil
    }
    res["externalMemberCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalMemberCount(val)
        }
        return nil
    }
    res["groupDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupDisplayName(val)
        }
        return nil
    }
    res["groupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupId(val)
        }
        return nil
    }
    res["groupType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupType(val)
        }
        return nil
    }
    res["isDeleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["lastActivityDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivityDate(val)
        }
        return nil
    }
    res["memberCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberCount(val)
        }
        return nil
    }
    res["ownerPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerPrincipalName(val)
        }
        return nil
    }
    res["reportPeriod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportPeriod(val)
        }
        return nil
    }
    res["reportRefreshDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["sharePointActiveFileCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointActiveFileCount(val)
        }
        return nil
    }
    res["sharePointSiteStorageUsedInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointSiteStorageUsedInBytes(val)
        }
        return nil
    }
    res["sharePointTotalFileCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointTotalFileCount(val)
        }
        return nil
    }
    res["teamsChannelMessagesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsChannelMessagesCount(val)
        }
        return nil
    }
    res["teamsMeetingsOrganizedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsMeetingsOrganizedCount(val)
        }
        return nil
    }
    res["yammerLikedMessageCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerLikedMessageCount(val)
        }
        return nil
    }
    res["yammerPostedMessageCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerPostedMessageCount(val)
        }
        return nil
    }
    res["yammerReadMessageCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetGroupDisplayName gets the groupDisplayName property value. The display name of the group.
func (m *Office365GroupsActivityDetail) GetGroupDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("groupDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetGroupId gets the groupId property value. The group id.
func (m *Office365GroupsActivityDetail) GetGroupId()(*string) {
    val, err := m.GetBackingStore().Get("groupId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetGroupType gets the groupType property value. The group type. Possible values are: Public or Private.
func (m *Office365GroupsActivityDetail) GetGroupType()(*string) {
    val, err := m.GetBackingStore().Get("groupType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIsDeleted gets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *Office365GroupsActivityDetail) GetIsDeleted()(*bool) {
    val, err := m.GetBackingStore().Get("isDeleted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastActivityDate gets the lastActivityDate property value. The last activity date for the following scenarios:  group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
func (m *Office365GroupsActivityDetail) GetLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("lastActivityDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetMemberCount gets the memberCount property value. The group member count.
func (m *Office365GroupsActivityDetail) GetMemberCount()(*int64) {
    val, err := m.GetBackingStore().Get("memberCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetOwnerPrincipalName gets the ownerPrincipalName property value. The group owner principal name.
func (m *Office365GroupsActivityDetail) GetOwnerPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("ownerPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReportPeriod gets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityDetail) GetReportPeriod()(*string) {
    val, err := m.GetBackingStore().Get("reportPeriod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReportRefreshDate gets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityDetail) GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("reportRefreshDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetSharePointActiveFileCount gets the sharePointActiveFileCount property value. The number of active files in SharePoint Group site.
func (m *Office365GroupsActivityDetail) GetSharePointActiveFileCount()(*int64) {
    val, err := m.GetBackingStore().Get("sharePointActiveFileCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSharePointSiteStorageUsedInBytes gets the sharePointSiteStorageUsedInBytes property value. The storage used by SharePoint Group site.
func (m *Office365GroupsActivityDetail) GetSharePointSiteStorageUsedInBytes()(*int64) {
    val, err := m.GetBackingStore().Get("sharePointSiteStorageUsedInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSharePointTotalFileCount gets the sharePointTotalFileCount property value. The total number of files in SharePoint Group site.
func (m *Office365GroupsActivityDetail) GetSharePointTotalFileCount()(*int64) {
    val, err := m.GetBackingStore().Get("sharePointTotalFileCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTeamsChannelMessagesCount gets the teamsChannelMessagesCount property value. The number of channel messages in Teams team.
func (m *Office365GroupsActivityDetail) GetTeamsChannelMessagesCount()(*int64) {
    val, err := m.GetBackingStore().Get("teamsChannelMessagesCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetTeamsMeetingsOrganizedCount gets the teamsMeetingsOrganizedCount property value. The number of meetings organized in Teams team.
func (m *Office365GroupsActivityDetail) GetTeamsMeetingsOrganizedCount()(*int64) {
    val, err := m.GetBackingStore().Get("teamsMeetingsOrganizedCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetYammerLikedMessageCount gets the yammerLikedMessageCount property value. The number of messages liked in Yammer groups.
func (m *Office365GroupsActivityDetail) GetYammerLikedMessageCount()(*int64) {
    val, err := m.GetBackingStore().Get("yammerLikedMessageCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetYammerPostedMessageCount gets the yammerPostedMessageCount property value. The number of messages posted to Yammer groups.
func (m *Office365GroupsActivityDetail) GetYammerPostedMessageCount()(*int64) {
    val, err := m.GetBackingStore().Get("yammerPostedMessageCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetYammerReadMessageCount gets the yammerReadMessageCount property value. The number of messages read in Yammer groups.
func (m *Office365GroupsActivityDetail) GetYammerReadMessageCount()(*int64) {
    val, err := m.GetBackingStore().Get("yammerReadMessageCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Office365GroupsActivityDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteInt64Value("teamsChannelMessagesCount", m.GetTeamsChannelMessagesCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("teamsMeetingsOrganizedCount", m.GetTeamsMeetingsOrganizedCount())
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
    err := m.GetBackingStore().Set("exchangeMailboxStorageUsedInBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetExchangeMailboxTotalItemCount sets the exchangeMailboxTotalItemCount property value. The number of items in the group mailbox.
func (m *Office365GroupsActivityDetail) SetExchangeMailboxTotalItemCount(value *int64)() {
    err := m.GetBackingStore().Set("exchangeMailboxTotalItemCount", value)
    if err != nil {
        panic(err)
    }
}
// SetExchangeReceivedEmailCount sets the exchangeReceivedEmailCount property value. The number of email that the group mailbox received.
func (m *Office365GroupsActivityDetail) SetExchangeReceivedEmailCount(value *int64)() {
    err := m.GetBackingStore().Set("exchangeReceivedEmailCount", value)
    if err != nil {
        panic(err)
    }
}
// SetExternalMemberCount sets the externalMemberCount property value. The group external member count.
func (m *Office365GroupsActivityDetail) SetExternalMemberCount(value *int64)() {
    err := m.GetBackingStore().Set("externalMemberCount", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupDisplayName sets the groupDisplayName property value. The display name of the group.
func (m *Office365GroupsActivityDetail) SetGroupDisplayName(value *string)() {
    err := m.GetBackingStore().Set("groupDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupId sets the groupId property value. The group id.
func (m *Office365GroupsActivityDetail) SetGroupId(value *string)() {
    err := m.GetBackingStore().Set("groupId", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupType sets the groupType property value. The group type. Possible values are: Public or Private.
func (m *Office365GroupsActivityDetail) SetGroupType(value *string)() {
    err := m.GetBackingStore().Set("groupType", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDeleted sets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *Office365GroupsActivityDetail) SetIsDeleted(value *bool)() {
    err := m.GetBackingStore().Set("isDeleted", value)
    if err != nil {
        panic(err)
    }
}
// SetLastActivityDate sets the lastActivityDate property value. The last activity date for the following scenarios:  group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
func (m *Office365GroupsActivityDetail) SetLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("lastActivityDate", value)
    if err != nil {
        panic(err)
    }
}
// SetMemberCount sets the memberCount property value. The group member count.
func (m *Office365GroupsActivityDetail) SetMemberCount(value *int64)() {
    err := m.GetBackingStore().Set("memberCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnerPrincipalName sets the ownerPrincipalName property value. The group owner principal name.
func (m *Office365GroupsActivityDetail) SetOwnerPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("ownerPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetReportPeriod sets the reportPeriod property value. The number of days the report covers.
func (m *Office365GroupsActivityDetail) SetReportPeriod(value *string)() {
    err := m.GetBackingStore().Set("reportPeriod", value)
    if err != nil {
        panic(err)
    }
}
// SetReportRefreshDate sets the reportRefreshDate property value. The latest date of the content.
func (m *Office365GroupsActivityDetail) SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("reportRefreshDate", value)
    if err != nil {
        panic(err)
    }
}
// SetSharePointActiveFileCount sets the sharePointActiveFileCount property value. The number of active files in SharePoint Group site.
func (m *Office365GroupsActivityDetail) SetSharePointActiveFileCount(value *int64)() {
    err := m.GetBackingStore().Set("sharePointActiveFileCount", value)
    if err != nil {
        panic(err)
    }
}
// SetSharePointSiteStorageUsedInBytes sets the sharePointSiteStorageUsedInBytes property value. The storage used by SharePoint Group site.
func (m *Office365GroupsActivityDetail) SetSharePointSiteStorageUsedInBytes(value *int64)() {
    err := m.GetBackingStore().Set("sharePointSiteStorageUsedInBytes", value)
    if err != nil {
        panic(err)
    }
}
// SetSharePointTotalFileCount sets the sharePointTotalFileCount property value. The total number of files in SharePoint Group site.
func (m *Office365GroupsActivityDetail) SetSharePointTotalFileCount(value *int64)() {
    err := m.GetBackingStore().Set("sharePointTotalFileCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsChannelMessagesCount sets the teamsChannelMessagesCount property value. The number of channel messages in Teams team.
func (m *Office365GroupsActivityDetail) SetTeamsChannelMessagesCount(value *int64)() {
    err := m.GetBackingStore().Set("teamsChannelMessagesCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsMeetingsOrganizedCount sets the teamsMeetingsOrganizedCount property value. The number of meetings organized in Teams team.
func (m *Office365GroupsActivityDetail) SetTeamsMeetingsOrganizedCount(value *int64)() {
    err := m.GetBackingStore().Set("teamsMeetingsOrganizedCount", value)
    if err != nil {
        panic(err)
    }
}
// SetYammerLikedMessageCount sets the yammerLikedMessageCount property value. The number of messages liked in Yammer groups.
func (m *Office365GroupsActivityDetail) SetYammerLikedMessageCount(value *int64)() {
    err := m.GetBackingStore().Set("yammerLikedMessageCount", value)
    if err != nil {
        panic(err)
    }
}
// SetYammerPostedMessageCount sets the yammerPostedMessageCount property value. The number of messages posted to Yammer groups.
func (m *Office365GroupsActivityDetail) SetYammerPostedMessageCount(value *int64)() {
    err := m.GetBackingStore().Set("yammerPostedMessageCount", value)
    if err != nil {
        panic(err)
    }
}
// SetYammerReadMessageCount sets the yammerReadMessageCount property value. The number of messages read in Yammer groups.
func (m *Office365GroupsActivityDetail) SetYammerReadMessageCount(value *int64)() {
    err := m.GetBackingStore().Set("yammerReadMessageCount", value)
    if err != nil {
        panic(err)
    }
}
// Office365GroupsActivityDetailable 
type Office365GroupsActivityDetailable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExchangeMailboxStorageUsedInBytes()(*int64)
    GetExchangeMailboxTotalItemCount()(*int64)
    GetExchangeReceivedEmailCount()(*int64)
    GetExternalMemberCount()(*int64)
    GetGroupDisplayName()(*string)
    GetGroupId()(*string)
    GetGroupType()(*string)
    GetIsDeleted()(*bool)
    GetLastActivityDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetMemberCount()(*int64)
    GetOwnerPrincipalName()(*string)
    GetReportPeriod()(*string)
    GetReportRefreshDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetSharePointActiveFileCount()(*int64)
    GetSharePointSiteStorageUsedInBytes()(*int64)
    GetSharePointTotalFileCount()(*int64)
    GetTeamsChannelMessagesCount()(*int64)
    GetTeamsMeetingsOrganizedCount()(*int64)
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
    SetLastActivityDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetMemberCount(value *int64)()
    SetOwnerPrincipalName(value *string)()
    SetReportPeriod(value *string)()
    SetReportRefreshDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetSharePointActiveFileCount(value *int64)()
    SetSharePointSiteStorageUsedInBytes(value *int64)()
    SetSharePointTotalFileCount(value *int64)()
    SetTeamsChannelMessagesCount(value *int64)()
    SetTeamsMeetingsOrganizedCount(value *int64)()
    SetYammerLikedMessageCount(value *int64)()
    SetYammerPostedMessageCount(value *int64)()
    SetYammerReadMessageCount(value *int64)()
}
