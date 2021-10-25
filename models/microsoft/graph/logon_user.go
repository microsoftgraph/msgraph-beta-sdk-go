package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type LogonUser struct {
    accountDomain *string;
    accountName *string;
    accountType *UserAccountSecurityType;
    additionalData map[string]interface{};
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    logonId *string;
    logonTypes []LogonType;
}
func NewLogonUser()(*LogonUser) {
    m := &LogonUser{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *LogonUser) GetAccountDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountDomain
    }
}
func (m *LogonUser) GetAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountName
    }
}
func (m *LogonUser) GetAccountType()(*UserAccountSecurityType) {
    if m == nil {
        return nil
    } else {
        return m.accountType
    }
}
func (m *LogonUser) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *LogonUser) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.firstSeenDateTime
    }
}
func (m *LogonUser) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSeenDateTime
    }
}
func (m *LogonUser) GetLogonId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logonId
    }
}
func (m *LogonUser) GetLogonTypes()([]LogonType) {
    if m == nil {
        return nil
    } else {
        return m.logonTypes
    }
}
func (m *LogonUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accountDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAccountDomain(val)
        return nil
    }
    res["accountName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAccountName(val)
        return nil
    }
    res["accountType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserAccountSecurityType)
        if err != nil {
            return err
        }
        cast := val.(UserAccountSecurityType)
        m.SetAccountType(&cast)
        return nil
    }
    res["firstSeenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetFirstSeenDateTime(val)
        return nil
    }
    res["lastSeenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSeenDateTime(val)
        return nil
    }
    res["logonId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLogonId(val)
        return nil
    }
    res["logonTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseLogonType)
        if err != nil {
            return err
        }
        res := make([]LogonType, len(val))
        for i, v := range val {
            res[i] = *(v.(*LogonType))
        }
        m.SetLogonTypes(res)
        return nil
    }
    return res
}
func (m *LogonUser) IsNil()(bool) {
    return m == nil
}
func (m *LogonUser) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("accountDomain", m.GetAccountDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("accountName", m.GetAccountName())
        if err != nil {
            return err
        }
    }
    if m.GetAccountType() != nil {
        cast := m.GetAccountType().String()
        err := writer.WriteStringValue("accountType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("firstSeenDateTime", m.GetFirstSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastSeenDateTime", m.GetLastSeenDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logonId", m.GetLogonId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("logonTypes", SerializeLogonType(m.GetLogonTypes()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *LogonUser) SetAccountDomain(value *string)() {
    m.accountDomain = value
}
func (m *LogonUser) SetAccountName(value *string)() {
    m.accountName = value
}
func (m *LogonUser) SetAccountType(value *UserAccountSecurityType)() {
    m.accountType = value
}
func (m *LogonUser) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *LogonUser) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
func (m *LogonUser) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSeenDateTime = value
}
func (m *LogonUser) SetLogonId(value *string)() {
    m.logonId = value
}
func (m *LogonUser) SetLogonTypes(value []LogonType)() {
    m.logonTypes = value
}
