package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AuthenticationDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The type of authentication method used to perform this step of authentication. Possible values: Password, SMS, Voice, Authenticator App, Software OATH token, Satisfied by token, Previously satisfied.
    authenticationMethod *string;
    // Details about the authentication method used to perform this authentication step. For example, phone number (for SMS and voice), device name (for Authenticator app), and password source (e.g. cloud, AD FS, PTA, PHS).
    authenticationMethodDetail *string;
    // Represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    authenticationStepDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The step of authentication that this satisfied. For example, primary authentication, or multi-factor authentication.
    authenticationStepRequirement *string;
    // Details about why the step succeeded or failed. For examples, user is blocked, fraud code entered, no phone input - timed out, phone unreachable, or claim in token.
    authenticationStepResultDetail *string;
    // Indicates the status of the authentication step. Possible values: succeeded, failed.
    succeeded *bool;
}
// Instantiates a new authenticationDetail and sets the default values.
func NewAuthenticationDetail()(*AuthenticationDetail) {
    m := &AuthenticationDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the authenticationMethod property value. The type of authentication method used to perform this step of authentication. Possible values: Password, SMS, Voice, Authenticator App, Software OATH token, Satisfied by token, Previously satisfied.
func (m *AuthenticationDetail) GetAuthenticationMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethod
    }
}
// Gets the authenticationMethodDetail property value. Details about the authentication method used to perform this authentication step. For example, phone number (for SMS and voice), device name (for Authenticator app), and password source (e.g. cloud, AD FS, PTA, PHS).
func (m *AuthenticationDetail) GetAuthenticationMethodDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodDetail
    }
}
// Gets the authenticationStepDateTime property value. Represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AuthenticationDetail) GetAuthenticationStepDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.authenticationStepDateTime
    }
}
// Gets the authenticationStepRequirement property value. The step of authentication that this satisfied. For example, primary authentication, or multi-factor authentication.
func (m *AuthenticationDetail) GetAuthenticationStepRequirement()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationStepRequirement
    }
}
// Gets the authenticationStepResultDetail property value. Details about why the step succeeded or failed. For examples, user is blocked, fraud code entered, no phone input - timed out, phone unreachable, or claim in token.
func (m *AuthenticationDetail) GetAuthenticationStepResultDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationStepResultDetail
    }
}
// Gets the succeeded property value. Indicates the status of the authentication step. Possible values: succeeded, failed.
func (m *AuthenticationDetail) GetSucceeded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.succeeded
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *AuthenticationDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the authenticationMethod property value. The type of authentication method used to perform this step of authentication. Possible values: Password, SMS, Voice, Authenticator App, Software OATH token, Satisfied by token, Previously satisfied.
// Parameters:
//  - value : Value to set for the authenticationMethod property.
func (m *AuthenticationDetail) SetAuthenticationMethod(value *string)() {
    m.authenticationMethod = value
}
// Sets the authenticationMethodDetail property value. Details about the authentication method used to perform this authentication step. For example, phone number (for SMS and voice), device name (for Authenticator app), and password source (e.g. cloud, AD FS, PTA, PHS).
// Parameters:
//  - value : Value to set for the authenticationMethodDetail property.
func (m *AuthenticationDetail) SetAuthenticationMethodDetail(value *string)() {
    m.authenticationMethodDetail = value
}
// Sets the authenticationStepDateTime property value. Represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the authenticationStepDateTime property.
func (m *AuthenticationDetail) SetAuthenticationStepDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.authenticationStepDateTime = value
}
// Sets the authenticationStepRequirement property value. The step of authentication that this satisfied. For example, primary authentication, or multi-factor authentication.
// Parameters:
//  - value : Value to set for the authenticationStepRequirement property.
func (m *AuthenticationDetail) SetAuthenticationStepRequirement(value *string)() {
    m.authenticationStepRequirement = value
}
// Sets the authenticationStepResultDetail property value. Details about why the step succeeded or failed. For examples, user is blocked, fraud code entered, no phone input - timed out, phone unreachable, or claim in token.
// Parameters:
//  - value : Value to set for the authenticationStepResultDetail property.
func (m *AuthenticationDetail) SetAuthenticationStepResultDetail(value *string)() {
    m.authenticationStepResultDetail = value
}
// Sets the succeeded property value. Indicates the status of the authentication step. Possible values: succeeded, failed.
// Parameters:
//  - value : Value to set for the succeeded property.
func (m *AuthenticationDetail) SetSucceeded(value *bool)() {
    m.succeeded = value
}
