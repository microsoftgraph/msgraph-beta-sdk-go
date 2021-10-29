package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SecurityActionState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The Application ID of the calling application that submitted an update (PATCH) to the action. The appId should be extracted from the auth token and not entered manually by the calling application.
    appId *string;
    // Status of the securityAction in this update. Possible values are: NotStarted, Running, Completed, Failed.
    status *OperationStatus;
    // Timestamp when the actionState was updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    updatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The user principal name of the signed-in user that submitted an update (PATCH) to the action. The user should be extracted from the auth token and not entered manually by the calling application.
    user *string;
}
// Instantiates a new securityActionState and sets the default values.
func NewSecurityActionState()(*SecurityActionState) {
    m := &SecurityActionState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SecurityActionState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the appId property value. The Application ID of the calling application that submitted an update (PATCH) to the action. The appId should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityActionState) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// Gets the status property value. Status of the securityAction in this update. Possible values are: NotStarted, Running, Completed, Failed.
func (m *SecurityActionState) GetStatus()(*OperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the updatedDateTime property value. Timestamp when the actionState was updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *SecurityActionState) GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.updatedDateTime
    }
}
// Gets the user property value. The user principal name of the signed-in user that submitted an update (PATCH) to the action. The user should be extracted from the auth token and not entered manually by the calling application.
func (m *SecurityActionState) GetUser()(*string) {
    if m == nil {
        return nil
    } else {
        return m.user
    }
}
// The deserialization information for the current model
func (m *SecurityActionState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppId(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationStatus)
        if err != nil {
            return err
        }
        cast := val.(OperationStatus)
        m.SetStatus(&cast)
        return nil
    }
    res["updatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetUpdatedDateTime(val)
        return nil
    }
    res["user"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUser(val)
        return nil
    }
    return res
}
func (m *SecurityActionState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SecurityActionState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updatedDateTime", m.GetUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("user", m.GetUser())
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
func (m *SecurityActionState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the appId property value. The Application ID of the calling application that submitted an update (PATCH) to the action. The appId should be extracted from the auth token and not entered manually by the calling application.
// Parameters:
//  - value : Value to set for the appId property.
func (m *SecurityActionState) SetAppId(value *string)() {
    m.appId = value
}
// Sets the status property value. Status of the securityAction in this update. Possible values are: NotStarted, Running, Completed, Failed.
// Parameters:
//  - value : Value to set for the status property.
func (m *SecurityActionState) SetStatus(value *OperationStatus)() {
    m.status = value
}
// Sets the updatedDateTime property value. Timestamp when the actionState was updated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the updatedDateTime property.
func (m *SecurityActionState) SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedDateTime = value
}
// Sets the user property value. The user principal name of the signed-in user that submitted an update (PATCH) to the action. The user should be extracted from the auth token and not entered manually by the calling application.
// Parameters:
//  - value : Value to set for the user property.
func (m *SecurityActionState) SetUser(value *string)() {
    m.user = value
}
