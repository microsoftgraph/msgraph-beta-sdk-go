package getoffice365groupsactivitycountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetOffice365GroupsActivityCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // The number of emails received by Group mailboxes.
    exchangeEmailsReceived *int64;
    // The date on which a number of emails were sent to a group mailbox or a number of messages were posted, read, or liked in a Yammer group
    reportDate *string;
    // The number of days the report covers.
    reportPeriod *string;
    // The latest date of the content.
    reportRefreshDate *string;
    // The number of messages liked in Yammer groups.
    yammerMessagesLiked *int64;
    // The number of messages posted to Yammer groups.
    yammerMessagesPosted *int64;
    // The number of messages read in Yammer groups.
    yammerMessagesRead *int64;
}
// Instantiates a new getOffice365GroupsActivityCountsWithPeriod and sets the default values.
func NewGetOffice365GroupsActivityCountsWithPeriod()(*GetOffice365GroupsActivityCountsWithPeriod) {
    m := &GetOffice365GroupsActivityCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the exchangeEmailsReceived property value. The number of emails received by Group mailboxes.
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetExchangeEmailsReceived()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeEmailsReceived
    }
}
// Gets the reportDate property value. The date on which a number of emails were sent to a group mailbox or a number of messages were posted, read, or liked in a Yammer group
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
// Gets the reportPeriod property value. The number of days the report covers.
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. The latest date of the content.
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the yammerMessagesLiked property value. The number of messages liked in Yammer groups.
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetYammerMessagesLiked()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerMessagesLiked
    }
}
// Gets the yammerMessagesPosted property value. The number of messages posted to Yammer groups.
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetYammerMessagesPosted()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerMessagesPosted
    }
}
// Gets the yammerMessagesRead property value. The number of messages read in Yammer groups.
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetYammerMessagesRead()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerMessagesRead
    }
}
// The deserialization information for the current model
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchangeEmailsReceived"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeEmailsReceived(val)
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
    res["yammerMessagesLiked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerMessagesLiked(val)
        }
        return nil
    }
    res["yammerMessagesPosted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerMessagesPosted(val)
        }
        return nil
    }
    res["yammerMessagesRead"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYammerMessagesRead(val)
        }
        return nil
    }
    return res
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *GetOffice365GroupsActivityCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("exchangeEmailsReceived", m.GetExchangeEmailsReceived())
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
        err = writer.WriteInt64Value("yammerMessagesLiked", m.GetYammerMessagesLiked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerMessagesPosted", m.GetYammerMessagesPosted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("yammerMessagesRead", m.GetYammerMessagesRead())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the exchangeEmailsReceived property value. The number of emails received by Group mailboxes.
// Parameters:
//  - value : Value to set for the exchangeEmailsReceived property.
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetExchangeEmailsReceived(value *int64)() {
    m.exchangeEmailsReceived = value
}
// Sets the reportDate property value. The date on which a number of emails were sent to a group mailbox or a number of messages were posted, read, or liked in a Yammer group
// Parameters:
//  - value : Value to set for the reportDate property.
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
// Sets the reportPeriod property value. The number of days the report covers.
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. The latest date of the content.
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the yammerMessagesLiked property value. The number of messages liked in Yammer groups.
// Parameters:
//  - value : Value to set for the yammerMessagesLiked property.
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetYammerMessagesLiked(value *int64)() {
    m.yammerMessagesLiked = value
}
// Sets the yammerMessagesPosted property value. The number of messages posted to Yammer groups.
// Parameters:
//  - value : Value to set for the yammerMessagesPosted property.
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetYammerMessagesPosted(value *int64)() {
    m.yammerMessagesPosted = value
}
// Sets the yammerMessagesRead property value. The number of messages read in Yammer groups.
// Parameters:
//  - value : Value to set for the yammerMessagesRead property.
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetYammerMessagesRead(value *int64)() {
    m.yammerMessagesRead = value
}
