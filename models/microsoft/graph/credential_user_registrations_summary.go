package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CredentialUserRegistrationsSummary struct {
    Entity
    lastRefreshedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    mfaAndSsprCapableUserCount *int32;
    mfaConditionalAccessPolicyState *string;
    mfaRegisteredUserCount *int32;
    securityDefaultsEnabled *bool;
    ssprEnabledUserCount *int32;
    ssprRegisteredUserCount *int32;
    tenantDisplayName *string;
    tenantId *string;
    totalUserCount *int32;
}
func NewCredentialUserRegistrationsSummary()(*CredentialUserRegistrationsSummary) {
    m := &CredentialUserRegistrationsSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *CredentialUserRegistrationsSummary) GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastRefreshedDateTime
    }
}
func (m *CredentialUserRegistrationsSummary) GetMfaAndSsprCapableUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mfaAndSsprCapableUserCount
    }
}
func (m *CredentialUserRegistrationsSummary) GetMfaConditionalAccessPolicyState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mfaConditionalAccessPolicyState
    }
}
func (m *CredentialUserRegistrationsSummary) GetMfaRegisteredUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.mfaRegisteredUserCount
    }
}
func (m *CredentialUserRegistrationsSummary) GetSecurityDefaultsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.securityDefaultsEnabled
    }
}
func (m *CredentialUserRegistrationsSummary) GetSsprEnabledUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.ssprEnabledUserCount
    }
}
func (m *CredentialUserRegistrationsSummary) GetSsprRegisteredUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.ssprRegisteredUserCount
    }
}
func (m *CredentialUserRegistrationsSummary) GetTenantDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantDisplayName
    }
}
func (m *CredentialUserRegistrationsSummary) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
func (m *CredentialUserRegistrationsSummary) GetTotalUserCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalUserCount
    }
}
func (m *CredentialUserRegistrationsSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastRefreshedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastRefreshedDateTime(val)
        return nil
    }
    res["mfaAndSsprCapableUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMfaAndSsprCapableUserCount(val)
        return nil
    }
    res["mfaConditionalAccessPolicyState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMfaConditionalAccessPolicyState(val)
        return nil
    }
    res["mfaRegisteredUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMfaRegisteredUserCount(val)
        return nil
    }
    res["securityDefaultsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSecurityDefaultsEnabled(val)
        return nil
    }
    res["ssprEnabledUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSsprEnabledUserCount(val)
        return nil
    }
    res["ssprRegisteredUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetSsprRegisteredUserCount(val)
        return nil
    }
    res["tenantDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantDisplayName(val)
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
        return nil
    }
    res["totalUserCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotalUserCount(val)
        return nil
    }
    return res
}
func (m *CredentialUserRegistrationsSummary) IsNil()(bool) {
    return m == nil
}
func (m *CredentialUserRegistrationsSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastRefreshedDateTime", m.GetLastRefreshedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mfaAndSsprCapableUserCount", m.GetMfaAndSsprCapableUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mfaConditionalAccessPolicyState", m.GetMfaConditionalAccessPolicyState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("mfaRegisteredUserCount", m.GetMfaRegisteredUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityDefaultsEnabled", m.GetSecurityDefaultsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ssprEnabledUserCount", m.GetSsprEnabledUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("ssprRegisteredUserCount", m.GetSsprRegisteredUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantDisplayName", m.GetTenantDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalUserCount", m.GetTotalUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *CredentialUserRegistrationsSummary) SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRefreshedDateTime = value
}
func (m *CredentialUserRegistrationsSummary) SetMfaAndSsprCapableUserCount(value *int32)() {
    m.mfaAndSsprCapableUserCount = value
}
func (m *CredentialUserRegistrationsSummary) SetMfaConditionalAccessPolicyState(value *string)() {
    m.mfaConditionalAccessPolicyState = value
}
func (m *CredentialUserRegistrationsSummary) SetMfaRegisteredUserCount(value *int32)() {
    m.mfaRegisteredUserCount = value
}
func (m *CredentialUserRegistrationsSummary) SetSecurityDefaultsEnabled(value *bool)() {
    m.securityDefaultsEnabled = value
}
func (m *CredentialUserRegistrationsSummary) SetSsprEnabledUserCount(value *int32)() {
    m.ssprEnabledUserCount = value
}
func (m *CredentialUserRegistrationsSummary) SetSsprRegisteredUserCount(value *int32)() {
    m.ssprRegisteredUserCount = value
}
func (m *CredentialUserRegistrationsSummary) SetTenantDisplayName(value *string)() {
    m.tenantDisplayName = value
}
func (m *CredentialUserRegistrationsSummary) SetTenantId(value *string)() {
    m.tenantId = value
}
func (m *CredentialUserRegistrationsSummary) SetTotalUserCount(value *int32)() {
    m.totalUserCount = value
}
