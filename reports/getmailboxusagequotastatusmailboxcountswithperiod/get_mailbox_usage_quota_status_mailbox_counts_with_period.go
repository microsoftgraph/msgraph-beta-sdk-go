package getmailboxusagequotastatusmailboxcountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetMailboxUsageQuotaStatusMailboxCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    indeterminate *int64;
    // 
    reportDate *string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    sendProhibited *int64;
    // 
    sendReceiveProhibited *int64;
    // 
    underLimit *int64;
    // 
    warningIssued *int64;
}
// Instantiates a new getMailboxUsageQuotaStatusMailboxCountsWithPeriod and sets the default values.
func NewGetMailboxUsageQuotaStatusMailboxCountsWithPeriod()(*GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) {
    m := &GetMailboxUsageQuotaStatusMailboxCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the indeterminate property value. 
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetIndeterminate()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.indeterminate
    }
}
// Gets the reportDate property value. 
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// Gets the reportPeriod property value. 
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the sendProhibited property value. 
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetSendProhibited()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sendProhibited
    }
}
// Gets the sendReceiveProhibited property value. 
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetSendReceiveProhibited()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.sendReceiveProhibited
    }
}
// Gets the underLimit property value. 
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetUnderLimit()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.underLimit
    }
}
// Gets the warningIssued property value. 
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetWarningIssued()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.warningIssued
    }
}
// The deserialization information for the current model
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["indeterminate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndeterminate(val)
        }
        return nil
    }
    res["reportDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportDate(val)
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportRefreshDate(val)
        }
        return nil
    }
    res["sendProhibited"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendProhibited(val)
        }
        return nil
    }
    res["sendReceiveProhibited"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendReceiveProhibited(val)
        }
        return nil
    }
    res["underLimit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnderLimit(val)
        }
        return nil
    }
    res["warningIssued"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWarningIssued(val)
        }
        return nil
    }
    return res
}
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the indeterminate property value. 
// Parameters:
//  - value : Value to set for the indeterminate property.
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetIndeterminate(value *int64)() {
    m.indeterminate = value
}
// Sets the reportDate property value. 
// Parameters:
//  - value : Value to set for the reportDate property.
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the sendProhibited property value. 
// Parameters:
//  - value : Value to set for the sendProhibited property.
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetSendProhibited(value *int64)() {
    m.sendProhibited = value
}
// Sets the sendReceiveProhibited property value. 
// Parameters:
//  - value : Value to set for the sendReceiveProhibited property.
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetSendReceiveProhibited(value *int64)() {
    m.sendReceiveProhibited = value
}
// Sets the underLimit property value. 
// Parameters:
//  - value : Value to set for the underLimit property.
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetUnderLimit(value *int64)() {
    m.underLimit = value
}
// Sets the warningIssued property value. 
// Parameters:
//  - value : Value to set for the warningIssued property.
func (m *GetMailboxUsageQuotaStatusMailboxCountsWithPeriod) SetWarningIssued(value *int64)() {
    m.warningIssued = value
}
