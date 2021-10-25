package getemailappusageuserdetailwithdate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

type GetEmailAppUsageUserDetailWithDate struct {
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
func NewGetEmailAppUsageUserDetailWithDate()(*GetEmailAppUsageUserDetailWithDate) {
    m := &GetEmailAppUsageUserDetailWithDate{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
func (m *GetEmailAppUsageUserDetailWithDate) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetImap4App()([]string) {
    if m == nil {
        return nil
    } else {
        return m.imap4App
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetMailForMac()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mailForMac
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetOtherForMobile()([]string) {
    if m == nil {
        return nil
    } else {
        return m.otherForMobile
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetOutlookForMac()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outlookForMac
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetOutlookForMobile()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outlookForMobile
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetOutlookForWeb()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outlookForWeb
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetOutlookForWindows()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outlookForWindows
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetPop3App()([]string) {
    if m == nil {
        return nil
    } else {
        return m.pop3App
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetSmtpApp()([]string) {
    if m == nil {
        return nil
    } else {
        return m.smtpApp
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
func (m *GetEmailAppUsageUserDetailWithDate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *GetEmailAppUsageUserDetailWithDate) IsNil()(bool) {
    return m == nil
}
func (m *GetEmailAppUsageUserDetailWithDate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *GetEmailAppUsageUserDetailWithDate) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetImap4App(value []string)() {
    m.imap4App = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetMailForMac(value []string)() {
    m.mailForMac = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetOtherForMobile(value []string)() {
    m.otherForMobile = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetOutlookForMac(value []string)() {
    m.outlookForMac = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetOutlookForMobile(value []string)() {
    m.outlookForMobile = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetOutlookForWeb(value []string)() {
    m.outlookForWeb = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetOutlookForWindows(value []string)() {
    m.outlookForWindows = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetPop3App(value []string)() {
    m.pop3App = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetSmtpApp(value []string)() {
    m.smtpApp = value
}
func (m *GetEmailAppUsageUserDetailWithDate) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
