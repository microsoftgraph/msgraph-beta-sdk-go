package getmailboxusagequotastatusmailboxcountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetMailboxUsageQuotaStatusMailboxCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    indeterminate *int64;
    reportDate *string;
    reportPeriod *string;
    reportRefreshDate *string;
    sendProhibited *int64;
    sendReceiveProhibited *int64;
    underLimit *int64;
    warningIssued *int64;
}
func NewGetMailboxUsageQuotaStatusMailboxCountsWithPeriod()(*GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) {
    m := &GetMailboxUsageQuotaStatusMailboxCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetIndeterminate()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.indeterminate
    }
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetSendProhibited()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sendProhibited
    }
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetSendReceiveProhibited()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sendReceiveProhibited
    }
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetUnderLimit()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.underLimit
    }
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetWarningIssued()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.warningIssued
    }
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["indeterminate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetIndeterminate(val)
        return nil
    }
    res["reportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportDate(val)
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
    res["sendProhibited"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSendProhibited(val)
        return nil
    }
    res["sendReceiveProhibited"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSendReceiveProhibited(val)
        return nil
    }
    res["underLimit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetUnderLimit(val)
        return nil
    }
    res["warningIssued"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetWarningIssued(val)
        return nil
    }
    return res
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("indeterminate", m.GetIndeterminate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportDate", m.GetReportDate())
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
        err = writer.WriteInt64Value("sendProhibited", m.GetSendProhibited())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sendReceiveProhibited", m.GetSendReceiveProhibited())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("underLimit", m.GetUnderLimit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("warningIssued", m.GetWarningIssued())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetIndeterminate(value *int64)() {
    m.indeterminate = value
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetSendProhibited(value *int64)() {
    m.sendProhibited = value
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetSendReceiveProhibited(value *int64)() {
    m.sendReceiveProhibited = value
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetUnderLimit(value *int64)() {
    m.underLimit = value
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetWarningIssued(value *int64)() {
    m.warningIssued = value
}
