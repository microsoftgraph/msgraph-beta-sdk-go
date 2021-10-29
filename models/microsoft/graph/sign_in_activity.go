package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SignInActivity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The last non-interactive sign-in date for a specific user. You can use this field to calculate the last time a client signed in to the directory on behalf of a user. Because some users may use clients to access tenant resources rather than signing into your tenant directly, you can use the non-interactive sign-in date to along with lastSignInDateTime to identify inactive users. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is: '2014-01-01T00:00:00Z'. For more information about using the value of this property, see Manage inactive user accounts in Azure AD.
    lastNonInteractiveSignInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Request identifier of the last non-interactive sign-in performed by this user.
    lastNonInteractiveSignInRequestId *string;
    // The last interactive sign-in date for a specific user. You can use this field to calculate the last time a user signed in to the directory with an interactive authentication method. This field can be used to build reports, such as inactive users. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is: '2014-01-01T00:00:00Z'. For more information about using the value of this property, see Manage inactive user accounts in Azure AD.
    lastSignInDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Request identifier of the last interactive sign-in performed by this user.
    lastSignInRequestId *string;
}
// Instantiates a new signInActivity and sets the default values.
func NewSignInActivity()(*SignInActivity) {
    m := &SignInActivity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SignInActivity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the lastNonInteractiveSignInDateTime property value. The last non-interactive sign-in date for a specific user. You can use this field to calculate the last time a client signed in to the directory on behalf of a user. Because some users may use clients to access tenant resources rather than signing into your tenant directly, you can use the non-interactive sign-in date to along with lastSignInDateTime to identify inactive users. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is: '2014-01-01T00:00:00Z'. For more information about using the value of this property, see Manage inactive user accounts in Azure AD.
func (m *SignInActivity) GetLastNonInteractiveSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastNonInteractiveSignInDateTime
    }
}
// Gets the lastNonInteractiveSignInRequestId property value. Request identifier of the last non-interactive sign-in performed by this user.
func (m *SignInActivity) GetLastNonInteractiveSignInRequestId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastNonInteractiveSignInRequestId
    }
}
// Gets the lastSignInDateTime property value. The last interactive sign-in date for a specific user. You can use this field to calculate the last time a user signed in to the directory with an interactive authentication method. This field can be used to build reports, such as inactive users. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is: '2014-01-01T00:00:00Z'. For more information about using the value of this property, see Manage inactive user accounts in Azure AD.
func (m *SignInActivity) GetLastSignInDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSignInDateTime
    }
}
// Gets the lastSignInRequestId property value. Request identifier of the last interactive sign-in performed by this user.
func (m *SignInActivity) GetLastSignInRequestId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastSignInRequestId
    }
}
// The deserialization information for the current model
func (m *SignInActivity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["lastNonInteractiveSignInDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastNonInteractiveSignInDateTime(val)
        return nil
    }
    res["lastNonInteractiveSignInRequestId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastNonInteractiveSignInRequestId(val)
        return nil
    }
    res["lastSignInDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastSignInDateTime(val)
        return nil
    }
    res["lastSignInRequestId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastSignInRequestId(val)
        return nil
    }
    return res
}
func (m *SignInActivity) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SignInActivity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("lastNonInteractiveSignInDateTime", m.GetLastNonInteractiveSignInDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastNonInteractiveSignInRequestId", m.GetLastNonInteractiveSignInRequestId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastSignInDateTime", m.GetLastSignInDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastSignInRequestId", m.GetLastSignInRequestId())
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
func (m *SignInActivity) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the lastNonInteractiveSignInDateTime property value. The last non-interactive sign-in date for a specific user. You can use this field to calculate the last time a client signed in to the directory on behalf of a user. Because some users may use clients to access tenant resources rather than signing into your tenant directly, you can use the non-interactive sign-in date to along with lastSignInDateTime to identify inactive users. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is: '2014-01-01T00:00:00Z'. For more information about using the value of this property, see Manage inactive user accounts in Azure AD.
// Parameters:
//  - value : Value to set for the lastNonInteractiveSignInDateTime property.
func (m *SignInActivity) SetLastNonInteractiveSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastNonInteractiveSignInDateTime = value
}
// Sets the lastNonInteractiveSignInRequestId property value. Request identifier of the last non-interactive sign-in performed by this user.
// Parameters:
//  - value : Value to set for the lastNonInteractiveSignInRequestId property.
func (m *SignInActivity) SetLastNonInteractiveSignInRequestId(value *string)() {
    m.lastNonInteractiveSignInRequestId = value
}
// Sets the lastSignInDateTime property value. The last interactive sign-in date for a specific user. You can use this field to calculate the last time a user signed in to the directory with an interactive authentication method. This field can be used to build reports, such as inactive users. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is: '2014-01-01T00:00:00Z'. For more information about using the value of this property, see Manage inactive user accounts in Azure AD.
// Parameters:
//  - value : Value to set for the lastSignInDateTime property.
func (m *SignInActivity) SetLastSignInDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSignInDateTime = value
}
// Sets the lastSignInRequestId property value. Request identifier of the last interactive sign-in performed by this user.
// Parameters:
//  - value : Value to set for the lastSignInRequestId property.
func (m *SignInActivity) SetLastSignInRequestId(value *string)() {
    m.lastSignInRequestId = value
}
