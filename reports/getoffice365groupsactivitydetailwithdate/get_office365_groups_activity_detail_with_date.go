package getoffice365groupsactivitydetailwithdate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetOffice365GroupsActivityDetailWithDate struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    exchangeMailboxStorageUsedInBytes *int64;
    exchangeMailboxTotalItemCount *int64;
    exchangeReceivedEmailCount *int64;
    externalMemberCount *int64;
    groupDisplayName *string;
    groupId *string;
    groupType *string;
    isDeleted *bool;
    lastActivityDate *string;
    memberCount *int64;
    ownerPrincipalName *string;
    reportPeriod *string;
    reportRefreshDate *string;
    sharePointActiveFileCount *int64;
    sharePointSiteStorageUsedInBytes *int64;
    sharePointTotalFileCount *int64;
    yammerLikedMessageCount *int64;
    yammerPostedMessageCount *int64;
    yammerReadMessageCount *int64;
}
func NewGetOffice365GroupsActivityDetailWithDate()(*GetOffice365GroupsActivityDetailWithDate) {
    m := &GetOffice365GroupsActivityDetailWithDate{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetExchangeMailboxStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeMailboxStorageUsedInBytes
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetExchangeMailboxTotalItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeMailboxTotalItemCount
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetExchangeReceivedEmailCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeReceivedEmailCount
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetExternalMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.externalMemberCount
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetGroupDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupDisplayName
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetGroupId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupId
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetGroupType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupType
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetMemberCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.memberCount
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetOwnerPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ownerPrincipalName
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetSharePointActiveFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointActiveFileCount
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetSharePointSiteStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointSiteStorageUsedInBytes
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetSharePointTotalFileCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sharePointTotalFileCount
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetYammerLikedMessageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerLikedMessageCount
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetYammerPostedMessageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerPostedMessageCount
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetYammerReadMessageCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerReadMessageCount
    }
}
func (m *GetOffice365GroupsActivityDetailWithDate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *GetOffice365GroupsActivityDetailWithDate) IsNil()(bool) {
    return m == nil
}
func (m *GetOffice365GroupsActivityDetailWithDate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *GetOffice365GroupsActivityDetailWithDate) SetExchangeMailboxStorageUsedInBytes(value *int64)() {
    m.exchangeMailboxStorageUsedInBytes = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetExchangeMailboxTotalItemCount(value *int64)() {
    m.exchangeMailboxTotalItemCount = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetExchangeReceivedEmailCount(value *int64)() {
    m.exchangeReceivedEmailCount = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetExternalMemberCount(value *int64)() {
    m.externalMemberCount = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetGroupDisplayName(value *string)() {
    m.groupDisplayName = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetGroupId(value *string)() {
    m.groupId = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetGroupType(value *string)() {
    m.groupType = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetMemberCount(value *int64)() {
    m.memberCount = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetOwnerPrincipalName(value *string)() {
    m.ownerPrincipalName = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetSharePointActiveFileCount(value *int64)() {
    m.sharePointActiveFileCount = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetSharePointSiteStorageUsedInBytes(value *int64)() {
    m.sharePointSiteStorageUsedInBytes = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetSharePointTotalFileCount(value *int64)() {
    m.sharePointTotalFileCount = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetYammerLikedMessageCount(value *int64)() {
    m.yammerLikedMessageCount = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetYammerPostedMessageCount(value *int64)() {
    m.yammerPostedMessageCount = value
}
func (m *GetOffice365GroupsActivityDetailWithDate) SetYammerReadMessageCount(value *int64)() {
    m.yammerReadMessageCount = value
}
