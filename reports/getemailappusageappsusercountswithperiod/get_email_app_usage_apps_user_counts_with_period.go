package getemailappusageappsusercountswithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetEmailAppUsageAppsUserCountsWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    imap4App *int64;
    // 
    mailForMac *int64;
    // 
    otherForMobile *int64;
    // 
    outlookForMac *int64;
    // 
    outlookForMobile *int64;
    // 
    outlookForWeb *int64;
    // 
    outlookForWindows *int64;
    // 
    pop3App *int64;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    smtpApp *int64;
}
// Instantiates a new getEmailAppUsageAppsUserCountsWithPeriod and sets the default values.
func NewGetEmailAppUsageAppsUserCountsWithPeriod()(*GetEmailAppUsageAppsUserCountsWithPeriod) {
    m := &GetEmailAppUsageAppsUserCountsWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the imap4App property value. 
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetImap4App()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.imap4App
    }
}
// Gets the mailForMac property value. 
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetMailForMac()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.mailForMac
    }
}
// Gets the otherForMobile property value. 
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetOtherForMobile()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.otherForMobile
    }
}
// Gets the outlookForMac property value. 
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetOutlookForMac()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlookForMac
    }
}
// Gets the outlookForMobile property value. 
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetOutlookForMobile()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlookForMobile
    }
}
// Gets the outlookForWeb property value. 
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetOutlookForWeb()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlookForWeb
    }
}
// Gets the outlookForWindows property value. 
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetOutlookForWindows()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.outlookForWindows
    }
}
// Gets the pop3App property value. 
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetPop3App()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.pop3App
    }
}
// Gets the reportPeriod property value. 
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the smtpApp property value. 
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) GetSmtpApp()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.smtpApp
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the imap4App property value. 
// Parameters:
//  - value : Value to set for the imap4App property.
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetImap4App(value *int64)() {
    m.imap4App = value
}
// Sets the mailForMac property value. 
// Parameters:
//  - value : Value to set for the mailForMac property.
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetMailForMac(value *int64)() {
    m.mailForMac = value
}
// Sets the otherForMobile property value. 
// Parameters:
//  - value : Value to set for the otherForMobile property.
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetOtherForMobile(value *int64)() {
    m.otherForMobile = value
}
// Sets the outlookForMac property value. 
// Parameters:
//  - value : Value to set for the outlookForMac property.
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetOutlookForMac(value *int64)() {
    m.outlookForMac = value
}
// Sets the outlookForMobile property value. 
// Parameters:
//  - value : Value to set for the outlookForMobile property.
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetOutlookForMobile(value *int64)() {
    m.outlookForMobile = value
}
// Sets the outlookForWeb property value. 
// Parameters:
//  - value : Value to set for the outlookForWeb property.
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetOutlookForWeb(value *int64)() {
    m.outlookForWeb = value
}
// Sets the outlookForWindows property value. 
// Parameters:
//  - value : Value to set for the outlookForWindows property.
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetOutlookForWindows(value *int64)() {
    m.outlookForWindows = value
}
// Sets the pop3App property value. 
// Parameters:
//  - value : Value to set for the pop3App property.
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetPop3App(value *int64)() {
    m.pop3App = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the smtpApp property value. 
// Parameters:
//  - value : Value to set for the smtpApp property.
func (m *GetEmailAppUsageAppsUserCountsWithPeriod) SetSmtpApp(value *int64)() {
    m.smtpApp = value
}
