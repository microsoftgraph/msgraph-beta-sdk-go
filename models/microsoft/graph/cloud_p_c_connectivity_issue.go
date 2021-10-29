package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPCConnectivityIssue struct {
    Entity
    // The Intune DeviceId of the device the connection is associated with.
    deviceId *string;
    // The error code of the connectivity issue.
    errorCode *string;
    // The time that the connection initiated. The time is shown in ISO 8601 format and Coordinated Universal Time (UTC) time.
    errorDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The detailed description of what went wrong.
    errorDescription *string;
    // The recommended action to fix the corresponding error.
    recommendedAction *string;
    // The unique id of user who initialize the connection.
    userId *string;
}
// Instantiates a new cloudPCConnectivityIssue and sets the default values.
func NewCloudPCConnectivityIssue()(*CloudPCConnectivityIssue) {
    m := &CloudPCConnectivityIssue{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceId property value. The Intune DeviceId of the device the connection is associated with.
func (m *CloudPCConnectivityIssue) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the errorCode property value. The error code of the connectivity issue.
func (m *CloudPCConnectivityIssue) GetErrorCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// Gets the errorDateTime property value. The time that the connection initiated. The time is shown in ISO 8601 format and Coordinated Universal Time (UTC) time.
func (m *CloudPCConnectivityIssue) GetErrorDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.errorDateTime
    }
}
// Gets the errorDescription property value. The detailed description of what went wrong.
func (m *CloudPCConnectivityIssue) GetErrorDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorDescription
    }
}
// Gets the recommendedAction property value. The recommended action to fix the corresponding error.
func (m *CloudPCConnectivityIssue) GetRecommendedAction()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recommendedAction
    }
}
// Gets the userId property value. The unique id of user who initialize the connection.
func (m *CloudPCConnectivityIssue) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// The deserialization information for the current model
func (m *CloudPCConnectivityIssue) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceId(val)
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetErrorCode(val)
        return nil
    }
    res["errorDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetErrorDateTime(val)
        return nil
    }
    res["errorDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetErrorDescription(val)
        return nil
    }
    res["recommendedAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRecommendedAction(val)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    return res
}
func (m *CloudPCConnectivityIssue) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CloudPCConnectivityIssue) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("errorDateTime", m.GetErrorDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorDescription", m.GetErrorDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recommendedAction", m.GetRecommendedAction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deviceId property value. The Intune DeviceId of the device the connection is associated with.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *CloudPCConnectivityIssue) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the errorCode property value. The error code of the connectivity issue.
// Parameters:
//  - value : Value to set for the errorCode property.
func (m *CloudPCConnectivityIssue) SetErrorCode(value *string)() {
    m.errorCode = value
}
// Sets the errorDateTime property value. The time that the connection initiated. The time is shown in ISO 8601 format and Coordinated Universal Time (UTC) time.
// Parameters:
//  - value : Value to set for the errorDateTime property.
func (m *CloudPCConnectivityIssue) SetErrorDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.errorDateTime = value
}
// Sets the errorDescription property value. The detailed description of what went wrong.
// Parameters:
//  - value : Value to set for the errorDescription property.
func (m *CloudPCConnectivityIssue) SetErrorDescription(value *string)() {
    m.errorDescription = value
}
// Sets the recommendedAction property value. The recommended action to fix the corresponding error.
// Parameters:
//  - value : Value to set for the recommendedAction property.
func (m *CloudPCConnectivityIssue) SetRecommendedAction(value *string)() {
    m.recommendedAction = value
}
// Sets the userId property value. The unique id of user who initialize the connection.
// Parameters:
//  - value : Value to set for the userId property.
func (m *CloudPCConnectivityIssue) SetUserId(value *string)() {
    m.userId = value
}
