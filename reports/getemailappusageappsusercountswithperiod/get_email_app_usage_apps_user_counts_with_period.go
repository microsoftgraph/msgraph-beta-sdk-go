package getemailappusageappsusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetEmailAppUsageAppsUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    imap4App *int64;
    mailForMac *int64;
    otherForMobile *int64;
    outlookForMac *int64;
    outlookForMobile *int64;
    outlookForWeb *int64;
    outlookForWindows *int64;
    pop3App *int64;
    reportPeriod *string;
    reportRefreshDate *string;
    smtpApp *int64;
}
func NewGetEmailAppUsageAppsUserCountsWithPeriod()(*GetEmailAppUsageAppsUserCountsWithPeriod) {
    m := &GetEmailAppUsageAppsUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetImap4App()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.imap4App
    }
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetMailForMac()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.mailForMac
    }
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetOtherForMobile()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.otherForMobile
    }
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetOutlookForMac()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlookForMac
    }
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetOutlookForMobile()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlookForMobile
    }
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetOutlookForWeb()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlookForWeb
    }
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetOutlookForWindows()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlookForWindows
    }
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetPop3App()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.pop3App
    }
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetSmtpApp()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.smtpApp
    }
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["imap4App"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetImap4App(val)
        return nil
    }
    res["mailForMac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetMailForMac(val)
        return nil
    }
    res["otherForMobile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOtherForMobile(val)
        return nil
    }
    res["outlookForMac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOutlookForMac(val)
        return nil
    }
    res["outlookForMobile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOutlookForMobile(val)
        return nil
    }
    res["outlookForWeb"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOutlookForWeb(val)
        return nil
    }
    res["outlookForWindows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetOutlookForWindows(val)
        return nil
    }
    res["pop3App"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetPop3App(val)
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
    res["smtpApp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetSmtpApp(val)
        return nil
    }
    return res
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("imap4App", m.GetImap4App())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("mailForMac", m.GetMailForMac())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("otherForMobile", m.GetOtherForMobile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("outlookForMac", m.GetOutlookForMac())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("outlookForMobile", m.GetOutlookForMobile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("outlookForWeb", m.GetOutlookForWeb())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("outlookForWindows", m.GetOutlookForWindows())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("pop3App", m.GetPop3App())
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
        err = writer.WriteInt64Value("smtpApp", m.GetSmtpApp())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetImap4App(value *int64)() {
    m.imap4App = value
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetMailForMac(value *int64)() {
    m.mailForMac = value
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetOtherForMobile(value *int64)() {
    m.otherForMobile = value
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetOutlookForMac(value *int64)() {
    m.outlookForMac = value
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetOutlookForMobile(value *int64)() {
    m.outlookForMobile = value
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetOutlookForWeb(value *int64)() {
    m.outlookForWeb = value
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetOutlookForWindows(value *int64)() {
    m.outlookForWindows = value
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetPop3App(value *int64)() {
    m.pop3App = value
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetSmtpApp(value *int64)() {
    m.smtpApp = value
}
