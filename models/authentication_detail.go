package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationDetail 
type AuthenticationDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The type of authentication method used to perform this step of authentication. Possible values: Password, SMS, Voice, Authenticator App, Software OATH token, Satisfied by token, Previously satisfied.
    authenticationMethod *string
    // Details about the authentication method used to perform this authentication step. For example, phone number (for SMS and voice), device name (for Authenticator app), and password source (e.g. cloud, AD FS, PTA, PHS).
    authenticationMethodDetail *string
    // Represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    authenticationStepDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The step of authentication that this satisfied. For example, primary authentication, or multi-factor authentication.
    authenticationStepRequirement *string
    // Details about why the step succeeded or failed. For examples, user is blocked, fraud code entered, no phone input - timed out, phone unreachable, or claim in token.
    authenticationStepResultDetail *string
    // Indicates the status of the authentication step. Possible values: succeeded, failed.
    succeeded *bool
}
// NewAuthenticationDetail instantiates a new authenticationDetail and sets the default values.
func NewAuthenticationDetail()(*AuthenticationDetail) {
    m := &AuthenticationDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAuthenticationDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAuthenticationMethod gets the authenticationMethod property value. The type of authentication method used to perform this step of authentication. Possible values: Password, SMS, Voice, Authenticator App, Software OATH token, Satisfied by token, Previously satisfied.
func (m *AuthenticationDetail) GetAuthenticationMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethod
    }
}
// GetAuthenticationMethodDetail gets the authenticationMethodDetail property value. Details about the authentication method used to perform this authentication step. For example, phone number (for SMS and voice), device name (for Authenticator app), and password source (e.g. cloud, AD FS, PTA, PHS).
func (m *AuthenticationDetail) GetAuthenticationMethodDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodDetail
    }
}
// GetAuthenticationStepDateTime gets the authenticationStepDateTime property value. Represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AuthenticationDetail) GetAuthenticationStepDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.authenticationStepDateTime
    }
}
// GetAuthenticationStepRequirement gets the authenticationStepRequirement property value. The step of authentication that this satisfied. For example, primary authentication, or multi-factor authentication.
func (m *AuthenticationDetail) GetAuthenticationStepRequirement()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationStepRequirement
    }
}
// GetAuthenticationStepResultDetail gets the authenticationStepResultDetail property value. Details about why the step succeeded or failed. For examples, user is blocked, fraud code entered, no phone input - timed out, phone unreachable, or claim in token.
func (m *AuthenticationDetail) GetAuthenticationStepResultDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationStepResultDetail
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val)
        }
        return nil
    }
    res["authenticationMethodDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethodDetail(val)
        }
        return nil
    }
    res["authenticationStepDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationStepDateTime(val)
        }
        return nil
    }
    res["authenticationStepRequirement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationStepRequirement(val)
        }
        return nil
    }
    res["authenticationStepResultDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationStepResultDetail(val)
        }
        return nil
    }
    res["succeeded"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSucceeded(val)
        }
        return nil
    }
    return res
}
// GetSucceeded gets the succeeded property value. Indicates the status of the authentication step. Possible values: succeeded, failed.
func (m *AuthenticationDetail) GetSucceeded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.succeeded
    }
}
// Serialize serializes information the current object
func (m *AuthenticationDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAuthenticationMethod sets the authenticationMethod property value. The type of authentication method used to perform this step of authentication. Possible values: Password, SMS, Voice, Authenticator App, Software OATH token, Satisfied by token, Previously satisfied.
func (m *AuthenticationDetail) SetAuthenticationMethod(value *string)() {
    if m != nil {
        m.authenticationMethod = value
    }
}
// SetAuthenticationMethodDetail sets the authenticationMethodDetail property value. Details about the authentication method used to perform this authentication step. For example, phone number (for SMS and voice), device name (for Authenticator app), and password source (e.g. cloud, AD FS, PTA, PHS).
func (m *AuthenticationDetail) SetAuthenticationMethodDetail(value *string)() {
    if m != nil {
        m.authenticationMethodDetail = value
    }
}
// SetAuthenticationStepDateTime sets the authenticationStepDateTime property value. Represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AuthenticationDetail) SetAuthenticationStepDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.authenticationStepDateTime = value
    }
}
// SetAuthenticationStepRequirement sets the authenticationStepRequirement property value. The step of authentication that this satisfied. For example, primary authentication, or multi-factor authentication.
func (m *AuthenticationDetail) SetAuthenticationStepRequirement(value *string)() {
    if m != nil {
        m.authenticationStepRequirement = value
    }
}
// SetAuthenticationStepResultDetail sets the authenticationStepResultDetail property value. Details about why the step succeeded or failed. For examples, user is blocked, fraud code entered, no phone input - timed out, phone unreachable, or claim in token.
func (m *AuthenticationDetail) SetAuthenticationStepResultDetail(value *string)() {
    if m != nil {
        m.authenticationStepResultDetail = value
    }
}
// SetSucceeded sets the succeeded property value. Indicates the status of the authentication step. Possible values: succeeded, failed.
func (m *AuthenticationDetail) SetSucceeded(value *bool)() {
    if m != nil {
        m.succeeded = value
    }
}
