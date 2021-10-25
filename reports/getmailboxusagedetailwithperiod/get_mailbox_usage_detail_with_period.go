package getmailboxusagedetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetMailboxUsageDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    createdDate *string;
    deletedDate *string;
    deletedItemCount *int64;
    deletedItemSizeInBytes *int64;
    displayName *string;
    isDeleted *bool;
    issueWarningQuotaInBytes *int64;
    itemCount *int64;
    lastActivityDate *string;
    prohibitSendQuotaInBytes *int64;
    prohibitSendReceiveQuotaInBytes *int64;
    reportPeriod *string;
    reportRefreshDate *string;
    storageUsedInBytes *int64;
    userPrincipalName *string;
}
func NewGetMailboxUsageDetailWithPeriod()(*GetMailboxUsageDetailWithPeriod) {
    m := &GetMailboxUsageDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetMailboxUsageDetailWithPeriod) GetCreatedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdDate
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetDeletedItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.deletedItemCount
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetDeletedItemSizeInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.deletedItemSizeInBytes
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetIssueWarningQuotaInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.issueWarningQuotaInBytes
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetItemCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.itemCount
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetProhibitSendQuotaInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.prohibitSendQuotaInBytes
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetProhibitSendReceiveQuotaInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.prohibitSendReceiveQuotaInBytes
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetStorageUsedInBytes()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.storageUsedInBytes
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *GetMailboxUsageDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCreatedDate(val)
        return nil
    }
    res["deletedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeletedDate(val)
        return nil
    }
    res["deletedItemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDeletedItemCount(val)
        return nil
    }
    res["deletedItemSizeInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetDeletedItemSizeInBytes(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
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
    res["issueWarningQuotaInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetIssueWarningQuotaInBytes(val)
        return nil
    }
    res["itemCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetItemCount(val)
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
    res["prohibitSendQuotaInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetProhibitSendQuotaInBytes(val)
        return nil
    }
    res["prohibitSendReceiveQuotaInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetProhibitSendReceiveQuotaInBytes(val)
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
    res["storageUsedInBytes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetStorageUsedInBytes(val)
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserPrincipalName(val)
        return nil
    }
    return res
}
func (m *GetMailboxUsageDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetMailboxUsageDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("createdDate", m.GetCreatedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deletedDate", m.GetDeletedDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("deletedItemCount", m.GetDeletedItemCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("deletedItemSizeInBytes", m.GetDeletedItemSizeInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err = writer.WriteInt64Value("issueWarningQuotaInBytes", m.GetIssueWarningQuotaInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("itemCount", m.GetItemCount())
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
        err = writer.WriteInt64Value("prohibitSendQuotaInBytes", m.GetProhibitSendQuotaInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("prohibitSendReceiveQuotaInBytes", m.GetProhibitSendReceiveQuotaInBytes())
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
        err = writer.WriteInt64Value("storageUsedInBytes", m.GetStorageUsedInBytes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetMailboxUsageDetailWithPeriod) SetCreatedDate(value *string)() {
    m.createdDate = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetDeletedItemCount(value *int64)() {
    m.deletedItemCount = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetDeletedItemSizeInBytes(value *int64)() {
    m.deletedItemSizeInBytes = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetIssueWarningQuotaInBytes(value *int64)() {
    m.issueWarningQuotaInBytes = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetItemCount(value *int64)() {
    m.itemCount = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetProhibitSendQuotaInBytes(value *int64)() {
    m.prohibitSendQuotaInBytes = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetProhibitSendReceiveQuotaInBytes(value *int64)() {
    m.prohibitSendReceiveQuotaInBytes = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetStorageUsedInBytes(value *int64)() {
    m.storageUsedInBytes = value
}
func (m *GetMailboxUsageDetailWithPeriod) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
