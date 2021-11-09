package getemailappusageuserdetailwithdate

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// 
type GetEmailAppUsageUserDetailWithDate struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // 
    deletedDate *string;
    // 
    displayName *string;
    // 
    imap4App []string;
    // 
    isDeleted *bool;
    // 
    lastActivityDate *string;
    // 
    mailForMac []string;
    // 
    otherForMobile []string;
    // 
    outlookForMac []string;
    // 
    outlookForMobile []string;
    // 
    outlookForWeb []string;
    // 
    outlookForWindows []string;
    // 
    pop3App []string;
    // 
    reportPeriod *string;
    // 
    reportRefreshDate *string;
    // 
    smtpApp []string;
    // 
    userPrincipalName *string;
}
// Instantiates a new getEmailAppUsageUserDetailWithDate and sets the default values.
func NewGetEmailAppUsageUserDetailWithDate()(*GetEmailAppUsageUserDetailWithDate) {
    m := &GetEmailAppUsageUserDetailWithDate{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// Gets the deletedDate property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetDeletedDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deletedDate
    }
}
// Gets the displayName property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the imap4App property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetImap4App()([]string) {
    if m == nil {
        return nil
    } else {
        return m.imap4App
    }
}
// Gets the isDeleted property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetIsDeleted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeleted
    }
}
// Gets the lastActivityDate property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetLastActivityDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastActivityDate
    }
}
// Gets the mailForMac property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetMailForMac()([]string) {
    if m == nil {
        return nil
    } else {
        return m.mailForMac
    }
}
// Gets the otherForMobile property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetOtherForMobile()([]string) {
    if m == nil {
        return nil
    } else {
        return m.otherForMobile
    }
}
// Gets the outlookForMac property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetOutlookForMac()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outlookForMac
    }
}
// Gets the outlookForMobile property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetOutlookForMobile()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outlookForMobile
    }
}
// Gets the outlookForWeb property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetOutlookForWeb()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outlookForWeb
    }
}
// Gets the outlookForWindows property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetOutlookForWindows()([]string) {
    if m == nil {
        return nil
    } else {
        return m.outlookForWindows
    }
}
// Gets the pop3App property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetPop3App()([]string) {
    if m == nil {
        return nil
    } else {
        return m.pop3App
    }
}
// Gets the reportPeriod property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetReportPeriod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportPeriod
    }
}
// Gets the reportRefreshDate property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetReportRefreshDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportRefreshDate
    }
}
// Gets the smtpApp property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetSmtpApp()([]string) {
    if m == nil {
        return nil
    } else {
        return m.smtpApp
    }
}
// Gets the userPrincipalName property value. 
func (m *GetEmailAppUsageUserDetailWithDate) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *GetEmailAppUsageUserDetailWithDate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deletedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedDate(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["imap4App"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetImap4App(res)
        }
        return nil
    }
    res["isDeleted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["lastActivityDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActivityDate(val)
        }
        return nil
    }
    res["mailForMac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMailForMac(res)
        }
        return nil
    }
    res["otherForMobile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOtherForMobile(res)
        }
        return nil
    }
    res["outlookForMac"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOutlookForMac(res)
        }
        return nil
    }
    res["outlookForMobile"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOutlookForMobile(res)
        }
        return nil
    }
    res["outlookForWeb"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOutlookForWeb(res)
        }
        return nil
    }
    res["outlookForWindows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOutlookForWindows(res)
        }
        return nil
    }
    res["pop3App"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPop3App(res)
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
    res["smtpApp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSmtpApp(res)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
func (m *GetEmailAppUsageUserDetailWithDate) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the deletedDate property value. 
// Parameters:
//  - value : Value to set for the deletedDate property.
func (m *GetEmailAppUsageUserDetailWithDate) SetDeletedDate(value *string)() {
    m.deletedDate = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *GetEmailAppUsageUserDetailWithDate) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the imap4App property value. 
// Parameters:
//  - value : Value to set for the imap4App property.
func (m *GetEmailAppUsageUserDetailWithDate) SetImap4App(value []string)() {
    m.imap4App = value
}
// Sets the isDeleted property value. 
// Parameters:
//  - value : Value to set for the isDeleted property.
func (m *GetEmailAppUsageUserDetailWithDate) SetIsDeleted(value *bool)() {
    m.isDeleted = value
}
// Sets the lastActivityDate property value. 
// Parameters:
//  - value : Value to set for the lastActivityDate property.
func (m *GetEmailAppUsageUserDetailWithDate) SetLastActivityDate(value *string)() {
    m.lastActivityDate = value
}
// Sets the mailForMac property value. 
// Parameters:
//  - value : Value to set for the mailForMac property.
func (m *GetEmailAppUsageUserDetailWithDate) SetMailForMac(value []string)() {
    m.mailForMac = value
}
// Sets the otherForMobile property value. 
// Parameters:
//  - value : Value to set for the otherForMobile property.
func (m *GetEmailAppUsageUserDetailWithDate) SetOtherForMobile(value []string)() {
    m.otherForMobile = value
}
// Sets the outlookForMac property value. 
// Parameters:
//  - value : Value to set for the outlookForMac property.
func (m *GetEmailAppUsageUserDetailWithDate) SetOutlookForMac(value []string)() {
    m.outlookForMac = value
}
// Sets the outlookForMobile property value. 
// Parameters:
//  - value : Value to set for the outlookForMobile property.
func (m *GetEmailAppUsageUserDetailWithDate) SetOutlookForMobile(value []string)() {
    m.outlookForMobile = value
}
// Sets the outlookForWeb property value. 
// Parameters:
//  - value : Value to set for the outlookForWeb property.
func (m *GetEmailAppUsageUserDetailWithDate) SetOutlookForWeb(value []string)() {
    m.outlookForWeb = value
}
// Sets the outlookForWindows property value. 
// Parameters:
//  - value : Value to set for the outlookForWindows property.
func (m *GetEmailAppUsageUserDetailWithDate) SetOutlookForWindows(value []string)() {
    m.outlookForWindows = value
}
// Sets the pop3App property value. 
// Parameters:
//  - value : Value to set for the pop3App property.
func (m *GetEmailAppUsageUserDetailWithDate) SetPop3App(value []string)() {
    m.pop3App = value
}
// Sets the reportPeriod property value. 
// Parameters:
//  - value : Value to set for the reportPeriod property.
func (m *GetEmailAppUsageUserDetailWithDate) SetReportPeriod(value *string)() {
    m.reportPeriod = value
}
// Sets the reportRefreshDate property value. 
// Parameters:
//  - value : Value to set for the reportRefreshDate property.
func (m *GetEmailAppUsageUserDetailWithDate) SetReportRefreshDate(value *string)() {
    m.reportRefreshDate = value
}
// Sets the smtpApp property value. 
// Parameters:
//  - value : Value to set for the smtpApp property.
func (m *GetEmailAppUsageUserDetailWithDate) SetSmtpApp(value []string)() {
    m.smtpApp = value
}
// Sets the userPrincipalName property value. 
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *GetEmailAppUsageUserDetailWithDate) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
