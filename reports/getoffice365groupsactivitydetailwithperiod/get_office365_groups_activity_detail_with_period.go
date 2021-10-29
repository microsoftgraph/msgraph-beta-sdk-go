package getoffice365groupsactivitydetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetOffice365GroupsActivityDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
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
    lastActivityDate *string;
    // The group member count.
    memberCount *int64;
    // The group owner principal name.
    ownerPrincipalName *string;
    // The number of days the report covers.
    reportPeriod *string;
    // The latest date of the content.
    reportRefreshDate *string;
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
// Instantiates a new getOffice365GroupsActivityDetailWithPeriod and sets the default values.
func NewGetOffice365GroupsActivityDetailWithPeriod()(*GetOffice365GroupsActivityDetailWithPeriod) {
    m := &GetOffice365GroupsActivityDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the exchangeMailboxStorageUsedInBytes property value. The storage used of the group mailbox.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetExchangeMailboxStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeMailboxStorageUsedInBytes
    }
}
// Gets the exchangeMailboxTotalItemCount property value. The number of items in the group mailbox.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetExchangeMailboxTotalItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeMailboxTotalItemCount
    }
}
// Gets the exchangeReceivedEmailCount property value. The number of email that the group mailbox received.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetExchangeReceivedEmailCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeReceivedEmailCount
    }
}
// Gets the externalMemberCount property value. The group external member count.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetExternalMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.externalMemberCount
    }
}
// Gets the groupDisplayName property value. The display name of the group.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetGroupDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupDisplayName
    }
}
// Gets the groupId property value. The group id.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupId
    }
}
// Gets the groupType property value. The group type. Possible values are: Public or Private.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetGroupType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupType
    }
}
// Gets the isDeleted property value. Whether this user has been deleted or soft deleted.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// Gets the lastActivityDate property value. The last activity date for the following scenarios:  group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the memberCount property value. The group member count.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.memberCount
    }
}
// Gets the ownerPrincipalName property value. The group owner principal name.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetOwnerPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerPrincipalName
    }
}
// Gets the reportPeriod property value. The number of days the report covers.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the sharePointActiveFileCount property value. The number of active files in SharePoint Group site.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetSharePointActiveFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointActiveFileCount
    }
}
// Gets the sharePointSiteStorageUsedInBytes property value. The storage used by SharePoint Group site.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetSharePointSiteStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointSiteStorageUsedInBytes
    }
}
// Gets the sharePointTotalFileCount property value. The total number of files in SharePoint Group site.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetSharePointTotalFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointTotalFileCount
    }
}
// Gets the yammerLikedMessageCount property value. The number of messages liked in Yammer groups.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetYammerLikedMessageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerLikedMessageCount
    }
}
// Gets the yammerPostedMessageCount property value. The number of messages posted to Yammer groups.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetYammerPostedMessageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerPostedMessageCount
    }
}
// Gets the yammerReadMessageCount property value. The number of messages read in Yammer groups.
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetYammerReadMessageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerReadMessageCount
    }
}
// The deserialization information for the current model
func (m *GetOffice365GroupsActivityDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchangeMailboxStorageUsedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetExchangeMailboxStorageUsedInBytes(val)
        return nil
    }
    res["exchangeMailboxTotalItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetExchangeMailboxTotalItemCount(val)
        return nil
    }
    res["exchangeReceivedEmailCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetExchangeReceivedEmailCount(val)
        return nil
    }
    res["externalMemberCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetExternalMemberCount(val)
        return nil
    }
    res["groupDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupDisplayName(val)
        return nil
    }
    res["groupId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupId(val)
        return nil
    }
    res["groupType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupType(val)
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDeleted(val)
        return nil
    }
    res["lastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastActivityDate(val)
        return nil
    }
    res["memberCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetMemberCount(val)
        return nil
    }
    res["ownerPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOwnerPrincipalName(val)
        return nil
    }
    res["reportPeriod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportPeriod(val)
        return nil
    }
    res["reportRefreshDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportRefreshDate(val)
        return nil
    }
    res["sharePointActiveFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSharePointActiveFileCount(val)
        return nil
    }
    res["sharePointSiteStorageUsedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSharePointSiteStorageUsedInBytes(val)
        return nil
    }
    res["sharePointTotalFileCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSharePointTotalFileCount(val)
        return nil
    }
    res["yammerLikedMessageCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetYammerLikedMessageCount(val)
        return nil
    }
    res["yammerPostedMessageCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetYammerPostedMessageCount(val)
        return nil
    }
    res["yammerReadMessageCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetYammerReadMessageCount(val)
        return nil
    }
    return res
}
func (m *GetOffice365GroupsActivityDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetOffice365GroupsActivityDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("lastActivityDate", m.GetLastActivityDate())
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
        err = writer.WriteStringValue("reportRefreshDate", m.GetReportRefreshDate())
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
// Sets the exchangeMailboxStorageUsedInBytes property value. The storage used of the group mailbox.
// Parameters:
//  - value : Value to set for the exchangeMailboxStorageUsedInBytes property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetExchangeMailboxStorageUsedInBytes(value *int64)() {
    m.exchangeMailboxStorageUsedInBytes = value
}
// Sets the exchangeMailboxTotalItemCount property value. The number of items in the group mailbox.
// Parameters:
//  - value : Value to set for the exchangeMailboxTotalItemCount property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetExchangeMailboxTotalItemCount(value *int64)() {
    m.exchangeMailboxTotalItemCount = value
}
// Sets the exchangeReceivedEmailCount property value. The number of email that the group mailbox received.
// Parameters:
//  - value : Value to set for the exchangeReceivedEmailCount property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetExchangeReceivedEmailCount(value *int64)() {
    m.exchangeReceivedEmailCount = value
}
// Sets the externalMemberCount property value. The group external member count.
// Parameters:
//  - value : Value to set for the externalMemberCount property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetExternalMemberCount(value *int64)() {
    m.externalMemberCount = value
}
// Sets the groupDisplayName property value. The display name of the group.
// Parameters:
//  - value : Value to set for the groupDisplayName property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetGroupDisplayName(value *string)() {
    m.groupDisplayName = value
}
// Sets the groupId property value. The group id.
// Parameters:
//  - value : Value to set for the groupId property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetGroupId(value *string)() {
    m.groupId = value
}
// Sets the groupType property value. The group type. Possible values are: Public or Private.
// Parameters:
//  - value : Value to set for the groupType property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetGroupType(value *string)() {
    m.groupType = value
}
// Sets the isDeleted property value. Whether this user has been deleted or soft deleted.
// Parameters:
//  - value : Value to set for the isDeleted property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// Sets the lastActivityDate property value. The last activity date for the following scenarios:  group mailbox received email; user viewed, edited, shared, or synced files in SharePoint document library; user viewed SharePoint pages; user posted, read, or liked messages in Yammer groups.
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the memberCount property value. The group member count.
// Parameters:
//  - value : Value to set for the memberCount property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetMemberCount(value *int64)() {
    m.memberCount = value
}
// Sets the ownerPrincipalName property value. The group owner principal name.
// Parameters:
//  - value : Value to set for the ownerPrincipalName property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetOwnerPrincipalName(value *string)() {
    m.ownerPrincipalName = value
}
// Sets the reportPeriod property value. The number of days the report covers.
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. The latest date of the content.
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the sharePointActiveFileCount property value. The number of active files in SharePoint Group site.
// Parameters:
//  - value : Value to set for the sharePointActiveFileCount property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetSharePointActiveFileCount(value *int64)() {
    m.sharePointActiveFileCount = value
}
// Sets the sharePointSiteStorageUsedInBytes property value. The storage used by SharePoint Group site.
// Parameters:
//  - value : Value to set for the sharePointSiteStorageUsedInBytes property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetSharePointSiteStorageUsedInBytes(value *int64)() {
    m.sharePointSiteStorageUsedInBytes = value
}
// Sets the sharePointTotalFileCount property value. The total number of files in SharePoint Group site.
// Parameters:
//  - value : Value to set for the sharePointTotalFileCount property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetSharePointTotalFileCount(value *int64)() {
    m.sharePointTotalFileCount = value
}
// Sets the yammerLikedMessageCount property value. The number of messages liked in Yammer groups.
// Parameters:
//  - value : Value to set for the yammerLikedMessageCount property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetYammerLikedMessageCount(value *int64)() {
    m.yammerLikedMessageCount = value
}
// Sets the yammerPostedMessageCount property value. The number of messages posted to Yammer groups.
// Parameters:
//  - value : Value to set for the yammerPostedMessageCount property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetYammerPostedMessageCount(value *int64)() {
    m.yammerPostedMessageCount = value
}
// Sets the yammerReadMessageCount property value. The number of messages read in Yammer groups.
// Parameters:
//  - value : Value to set for the yammerReadMessageCount property.
func (m *GetOffice365GroupsActivityDetailWithPeriod) SetYammerReadMessageCount(value *int64)() {
    m.yammerReadMessageCount = value
}
