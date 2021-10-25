package getemailappusageuserdetailwithperiod

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetEmailAppUsageUserDetailWithPeriod struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    deletedDate *string;
    displayName *string;
    imap4App []string;
    isDeleted *bool;
    lastActivityDate *string;
    mailForMac []string;
    otherForMobile []string;
    outlookForMac []string;
    outlookForMobile []string;
    outlookForWeb []string;
    outlookForWindows []string;
    pop3App []string;
    reportPeriod *string;
    reportRefreshDate *string;
    smtpApp []string;
    userPrincipalName *string;
}
func NewGetEmailAppUsageUserDetailWithPeriod()(*GetEmailAppUsageUserDetailWithPeriod) {
    m := &GetEmailAppUsageUserDetailWithPeriod{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetImap4App()([]string) {
    if m == nil {
        return nil
    } else {
        return m.imap4App
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetMailForMac()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mailForMac
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetOtherForMobile()([]string) {
    if m == nil {
        return nil
    } else {
        return m.otherForMobile
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetOutlookForMac()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outlookForMac
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetOutlookForMobile()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outlookForMobile
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetOutlookForWeb()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outlookForWeb
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetOutlookForWindows()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outlookForWindows
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetPop3App()([]string) {
    if m == nil {
        return nil
    } else {
        return m.pop3App
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetSmtpApp()([]string) {
    if m == nil {
        return nil
    } else {
        return m.smtpApp
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *GetEmailAppUsageUserDetailWithPeriod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deletedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeletedDate(val)
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
    res["imap4App"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetImap4App(res)
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
    res["mailForMac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetMailForMac(res)
        return nil
    }
    res["otherForMobile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetOtherForMobile(res)
        return nil
    }
    res["outlookForMac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetOutlookForMac(res)
        return nil
    }
    res["outlookForMobile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetOutlookForMobile(res)
        return nil
    }
    res["outlookForWeb"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetOutlookForWeb(res)
        return nil
    }
    res["outlookForWindows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetOutlookForWindows(res)
        return nil
    }
    res["pop3App"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetPop3App(res)
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
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSmtpApp(res)
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
func (m *GetEmailAppUsageUserDetailWithPeriod) IsNil()(bool) {
    return m == nil
}
func (m *GetEmailAppUsageUserDetailWithPeriod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deletedDate", m.GetDeletedDate())
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
        err = writer.WriteCollectionOfStringValues("imap4App", m.GetImap4App())
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
        err = writer.WriteCollectionOfStringValues("mailForMac", m.GetMailForMac())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("otherForMobile", m.GetOtherForMobile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("outlookForMac", m.GetOutlookForMac())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("outlookForMobile", m.GetOutlookForMobile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("outlookForWeb", m.GetOutlookForWeb())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("outlookForWindows", m.GetOutlookForWindows())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("pop3App", m.GetPop3App())
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
        err = writer.WriteCollectionOfStringValues("smtpApp", m.GetSmtpApp())
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
func (m *GetEmailAppUsageUserDetailWithPeriod) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetImap4App(value []string)() {
    m.imap4App = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetMailForMac(value []string)() {
    m.mailForMac = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetOtherForMobile(value []string)() {
    m.otherForMobile = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetOutlookForMac(value []string)() {
    m.outlookForMac = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetOutlookForMobile(value []string)() {
    m.outlookForMobile = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetOutlookForWeb(value []string)() {
    m.outlookForWeb = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetOutlookForWindows(value []string)() {
    m.outlookForWindows = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetPop3App(value []string)() {
    m.pop3App = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetSmtpApp(value []string)() {
    m.smtpApp = value
}
func (m *GetEmailAppUsageUserDetailWithPeriod) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
