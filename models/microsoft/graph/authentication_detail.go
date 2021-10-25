package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AuthenticationDetail struct {
    additionalData map[string]interface{};
    authenticationMethod *string;
    authenticationMethodDetail *string;
    authenticationStepDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    authenticationStepRequirement *string;
    authenticationStepResultDetail *string;
    succeeded *bool;
}
func NewAuthenticationDetail()(*AuthenticationDetail) {
    m := &AuthenticationDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AuthenticationDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AuthenticationDetail) GetAuthenticationMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethod
    }
}
func (m *AuthenticationDetail) GetAuthenticationMethodDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodDetail
    }
}
func (m *AuthenticationDetail) GetAuthenticationStepDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.authenticationStepDateTime
    }
}
func (m *AuthenticationDetail) GetAuthenticationStepRequirement()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationStepRequirement
    }
}
func (m *AuthenticationDetail) GetAuthenticationStepResultDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationStepResultDetail
    }
}
func (m *AuthenticationDetail) GetSucceeded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.succeeded
    }
}
func (m *AuthenticationDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["authenticationMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthenticationMethod(val)
        return nil
    }
    res["authenticationMethodDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthenticationMethodDetail(val)
        return nil
    }
    res["authenticationStepDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetAuthenticationStepDateTime(val)
        return nil
    }
    res["authenticationStepRequirement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthenticationStepRequirement(val)
        return nil
    }
    res["authenticationStepResultDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthenticationStepResultDetail(val)
        return nil
    }
    res["succeeded"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSucceeded(val)
        return nil
    }
    return res
}
func (m *AuthenticationDetail) IsNil()(bool) {
    return m == nil
}
func (m *AuthenticationDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("authenticationMethod", m.GetAuthenticationMethod())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("authenticationMethodDetail", m.GetAuthenticationMethodDetail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("authenticationStepDateTime", m.GetAuthenticationStepDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("authenticationStepRequirement", m.GetAuthenticationStepRequirement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("authenticationStepResultDetail", m.GetAuthenticationStepResultDetail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("succeeded", m.GetSucceeded())
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
func (m *AuthenticationDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AuthenticationDetail) SetAuthenticationMethod(value *string)() {
    m.authenticationMethod = value
}
func (m *AuthenticationDetail) SetAuthenticationMethodDetail(value *string)() {
    m.authenticationMethodDetail = value
}
func (m *AuthenticationDetail) SetAuthenticationStepDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.authenticationStepDateTime = value
}
func (m *AuthenticationDetail) SetAuthenticationStepRequirement(value *string)() {
    m.authenticationStepRequirement = value
}
func (m *AuthenticationDetail) SetAuthenticationStepResultDetail(value *string)() {
    m.authenticationStepResultDetail = value
}
func (m *AuthenticationDetail) SetSucceeded(value *bool)() {
    m.succeeded = value
}
