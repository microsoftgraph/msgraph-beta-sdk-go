package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// LogonUser 
type LogonUser struct {
    // Domain of user account used to logon.
    accountDomain *string;
    // Account name of user account used to logon.
    accountName *string;
    // User Account type, per Windows definition. Possible values are: unknown, standard, power, administrator.
    accountType *UserAccountSecurityType;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // DateTime at which the earliest logon by this user account occurred (provider-determined period). The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // DateTime at which the latest logon by this user account occurred. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    lastSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // User logon ID.
    logonId *string;
    // Collection of the logon types observed for the logged on user from when first to last seen. Possible values are: unknown, interactive, remoteInteractive, network, batch, service.
    logonTypes []LogonType;
}
// NewLogonUser instantiates a new logonUser and sets the default values.
func NewLogonUser()(*LogonUser) {
    m := &LogonUser{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAccountDomain gets the accountDomain property value. Domain of user account used to logon.
func (m *LogonUser) GetAccountDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountDomain
    }
}
// GetAccountName gets the accountName property value. Account name of user account used to logon.
func (m *LogonUser) GetAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountName
    }
}
// GetAccountType gets the accountType property value. User Account type, per Windows definition. Possible values are: unknown, standard, power, administrator.
func (m *LogonUser) GetAccountType()(*UserAccountSecurityType) {
    if m == nil {
        return nil
    } else {
        return m.accountType
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LogonUser) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. DateTime at which the earliest logon by this user account occurred (provider-determined period). The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *LogonUser) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.firstSeenDateTime
    }
}
// GetLastSeenDateTime gets the lastSeenDateTime property value. DateTime at which the latest logon by this user account occurred. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *LogonUser) GetLastSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSeenDateTime
    }
}
// GetLogonId gets the logonId property value. User logon ID.
func (m *LogonUser) GetLogonId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.logonId
    }
}
// GetLogonTypes gets the logonTypes property value. Collection of the logon types observed for the logged on user from when first to last seen. Possible values are: unknown, interactive, remoteInteractive, network, batch, service.
func (m *LogonUser) GetLogonTypes()([]LogonType) {
    if m == nil {
        return nil
    } else {
        return m.logonTypes
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LogonUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accountDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountDomain(val)
        }
        return nil
    }
    res["accountName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountName(val)
        }
        return nil
    }
    res["accountType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserAccountSecurityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountType(val.(*UserAccountSecurityType))
        }
        return nil
    }
    res["firstSeenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstSeenDateTime(val)
        }
        return nil
    }
    res["lastSeenDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSeenDateTime(val)
        }
        return nil
    }
    res["logonId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogonId(val)
        }
        return nil
    }
    res["logonTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseLogonType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LogonType, len(val))
            for i, v := range val {
                res[i] = *(v.(*LogonType))
            }
            m.SetLogonTypes(res)
        }
        return nil
    }
    return res
}
func (m *LogonUser) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetAccountType()).String()
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
    if m.GetLogonTypes() != nil {
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
// SetAccountDomain sets the accountDomain property value. Domain of user account used to logon.
func (m *LogonUser) SetAccountDomain(value *string)() {
    if m != nil {
        m.accountDomain = value
    }
}
// SetAccountName sets the accountName property value. Account name of user account used to logon.
func (m *LogonUser) SetAccountName(value *string)() {
    if m != nil {
        m.accountName = value
    }
}
// SetAccountType sets the accountType property value. User Account type, per Windows definition. Possible values are: unknown, standard, power, administrator.
func (m *LogonUser) SetAccountType(value *UserAccountSecurityType)() {
    if m != nil {
        m.accountType = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LogonUser) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. DateTime at which the earliest logon by this user account occurred (provider-determined period). The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *LogonUser) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.firstSeenDateTime = value
    }
}
// SetLastSeenDateTime sets the lastSeenDateTime property value. DateTime at which the latest logon by this user account occurred. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *LogonUser) SetLastSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSeenDateTime = value
    }
}
// SetLogonId sets the logonId property value. User logon ID.
func (m *LogonUser) SetLogonId(value *string)() {
    if m != nil {
        m.logonId = value
    }
}
// SetLogonTypes sets the logonTypes property value. Collection of the logon types observed for the logged on user from when first to last seen. Possible values are: unknown, interactive, remoteInteractive, network, batch, service.
func (m *LogonUser) SetLogonTypes(value []LogonType)() {
    if m != nil {
        m.logonTypes = value
    }
}
