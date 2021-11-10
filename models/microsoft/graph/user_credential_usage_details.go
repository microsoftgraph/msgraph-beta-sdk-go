package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UserCredentialUsageDetails struct {
    Entity
    // Represents the authentication method that the user used. Possible values are:email, mobileSMS, mobileCall, officePhone, securityQuestion (only used for self-service password reset), appNotification, appCode, alternateMobileCall (supported only in registration), fido, appPassword,unknownFutureValue
    authMethod *UsageAuthMethod;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Provides the failure reason for the corresponding reset or registration workflow.
    failureReason *string;
    // Possible values are: registration, reset, unknownFutureValue.
    feature *FeatureType;
    // Indicates success or failure of the workflow.
    isSuccess *bool;
    // User name of the user performing the reset or registration workflow.
    userDisplayName *string;
    // User principal name of the user performing the reset or registration workflow.
    userPrincipalName *string;
}
// Instantiates a new userCredentialUsageDetails and sets the default values.
func NewUserCredentialUsageDetails()(*UserCredentialUsageDetails) {
    m := &UserCredentialUsageDetails{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the authMethod property value. Represents the authentication method that the user used. Possible values are:email, mobileSMS, mobileCall, officePhone, securityQuestion (only used for self-service password reset), appNotification, appCode, alternateMobileCall (supported only in registration), fido, appPassword,unknownFutureValue
func (m *UserCredentialUsageDetails) GetAuthMethod()(*UsageAuthMethod) {
    if m == nil {
        return nil
    } else {
        return m.authMethod
    }
}
// Gets the eventDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *UserCredentialUsageDetails) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
// Gets the failureReason property value. Provides the failure reason for the corresponding reset or registration workflow.
func (m *UserCredentialUsageDetails) GetFailureReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.failureReason
    }
}
// Gets the feature property value. Possible values are: registration, reset, unknownFutureValue.
func (m *UserCredentialUsageDetails) GetFeature()(*FeatureType) {
    if m == nil {
        return nil
    } else {
        return m.feature
    }
}
// Gets the isSuccess property value. Indicates success or failure of the workflow.
func (m *UserCredentialUsageDetails) GetIsSuccess()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSuccess
    }
}
// Gets the userDisplayName property value. User name of the user performing the reset or registration workflow.
func (m *UserCredentialUsageDetails) GetUserDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userDisplayName
    }
}
// Gets the userPrincipalName property value. User principal name of the user performing the reset or registration workflow.
func (m *UserCredentialUsageDetails) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// The deserialization information for the current model
func (m *UserCredentialUsageDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUsageAuthMethod)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(UsageAuthMethod)
            m.SetAuthMethod(&cast)
        }
        return nil
    }
    res["eventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
        }
        return nil
    }
    res["failureReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureReason(val)
        }
        return nil
    }
    res["feature"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFeatureType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(FeatureType)
            m.SetFeature(&cast)
        }
        return nil
    }
    res["isSuccess"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSuccess(val)
        }
        return nil
    }
    res["userDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
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
func (m *UserCredentialUsageDetails) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UserCredentialUsageDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthMethod() != nil {
        cast := m.GetAuthMethod().String()
        err = writer.WriteStringValue("authMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("failureReason", m.GetFailureReason())
        if err != nil {
            return err
        }
    }
    if m.GetFeature() != nil {
        cast := m.GetFeature().String()
        err = writer.WriteStringValue("feature", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSuccess", m.GetIsSuccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
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
// Sets the authMethod property value. Represents the authentication method that the user used. Possible values are:email, mobileSMS, mobileCall, officePhone, securityQuestion (only used for self-service password reset), appNotification, appCode, alternateMobileCall (supported only in registration), fido, appPassword,unknownFutureValue
// Parameters:
//  - value : Value to set for the authMethod property.
func (m *UserCredentialUsageDetails) SetAuthMethod(value *UsageAuthMethod)() {
    m.authMethod = value
}
// Sets the eventDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the eventDateTime property.
func (m *UserCredentialUsageDetails) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// Sets the failureReason property value. Provides the failure reason for the corresponding reset or registration workflow.
// Parameters:
//  - value : Value to set for the failureReason property.
func (m *UserCredentialUsageDetails) SetFailureReason(value *string)() {
    m.failureReason = value
}
// Sets the feature property value. Possible values are: registration, reset, unknownFutureValue.
// Parameters:
//  - value : Value to set for the feature property.
func (m *UserCredentialUsageDetails) SetFeature(value *FeatureType)() {
    m.feature = value
}
// Sets the isSuccess property value. Indicates success or failure of the workflow.
// Parameters:
//  - value : Value to set for the isSuccess property.
func (m *UserCredentialUsageDetails) SetIsSuccess(value *bool)() {
    m.isSuccess = value
}
// Sets the userDisplayName property value. User name of the user performing the reset or registration workflow.
// Parameters:
//  - value : Value to set for the userDisplayName property.
func (m *UserCredentialUsageDetails) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// Sets the userPrincipalName property value. User principal name of the user performing the reset or registration workflow.
// Parameters:
//  - value : Value to set for the userPrincipalName property.
func (m *UserCredentialUsageDetails) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
