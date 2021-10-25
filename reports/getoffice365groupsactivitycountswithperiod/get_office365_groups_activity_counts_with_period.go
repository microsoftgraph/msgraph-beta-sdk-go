package getoffice365groupsactivitycountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetOffice365GroupsActivityCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    exchangeEmailsReceived *int64;
    reportDate *string;
    reportPeriod *string;
    reportRefreshDate *string;
    yammerMessagesLiked *int64;
    yammerMessagesPosted *int64;
    yammerMessagesRead *int64;
}
func NewGetOffice365GroupsActivityCountsWithPeriod()(*GetOffice365GroupsActivityCountsWithPeriod) {
    m := &GetOffice365GroupsActivityCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetExchangeEmailsReceived()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.exchangeEmailsReceived
    }
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetReportDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportDate
    }
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetYammerMessagesLiked()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerMessagesLiked
    }
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetYammerMessagesPosted()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerMessagesPosted
    }
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetYammerMessagesRead()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.yammerMessagesRead
    }
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exchangeEmailsReceived"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetExchangeEmailsReceived(val)
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
    res["yammerMessagesLiked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetYammerMessagesLiked(val)
        return nil
    }
    res["yammerMessagesPosted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetYammerMessagesPosted(val)
        return nil
    }
    res["yammerMessagesRead"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetYammerMessagesRead(val)
        return nil
    }
    return res
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
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
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetExchangeEmailsReceived(value *int64)() {
    m.exchangeEmailsReceived = value
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetReportDate(value *string)() {
    m.reportDate = value
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetYammerMessagesLiked(value *int64)() {
    m.yammerMessagesLiked = value
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetYammerMessagesPosted(value *int64)() {
    m.yammerMessagesPosted = value
}
func (m *GetOffice365GroupsActivityCountsWithPeriod) SetYammerMessagesRead(value *int64)() {
    m.yammerMessagesRead = value
}
