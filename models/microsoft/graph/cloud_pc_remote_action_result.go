package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPcRemoteActionResult struct {
    // The specified action. Supported values: Reprovision, Resize.
    actionName *string;
    // State of the action. Possible values are: None, pending, canceled, active, done, failed, notSupported. Read-only.
    actionState *ActionState;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The ID of the Cloud PC device on which the remote action is performed. Read-only.
    cloudPcId *string;
    // Last update time for action. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The ID of the Intune managed device on which the remote action is performed. Read-only.
    managedDeviceId *string;
    // Time the action was initiated. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The details of the Cloud PC status.
    statusDetails *CloudPcStatusDetails;
}
// Instantiates a new cloudPcRemoteActionResult and sets the default values.
func NewCloudPcRemoteActionResult()(*CloudPcRemoteActionResult) {
    m := &CloudPcRemoteActionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the actionName property value. The specified action. Supported values: Reprovision, Resize.
func (m *CloudPcRemoteActionResult) GetActionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionName
    }
}
// Gets the actionState property value. State of the action. Possible values are: None, pending, canceled, active, done, failed, notSupported. Read-only.
func (m *CloudPcRemoteActionResult) GetActionState()(*ActionState) {
    if m == nil {
        return nil
    } else {
        return m.actionState
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcRemoteActionResult) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the cloudPcId property value. The ID of the Cloud PC device on which the remote action is performed. Read-only.
func (m *CloudPcRemoteActionResult) GetCloudPcId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cloudPcId
    }
}
// Gets the lastUpdatedDateTime property value. Last update time for action. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
func (m *CloudPcRemoteActionResult) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// Gets the managedDeviceId property value. The ID of the Intune managed device on which the remote action is performed. Read-only.
func (m *CloudPcRemoteActionResult) GetManagedDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceId
    }
}
// Gets the startDateTime property value. Time the action was initiated. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
func (m *CloudPcRemoteActionResult) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the statusDetails property value. The details of the Cloud PC status.
func (m *CloudPcRemoteActionResult) GetStatusDetails()(*CloudPcStatusDetails) {
    if m == nil {
        return nil
    } else {
        return m.statusDetails
    }
}
// The deserialization information for the current model
func (m *CloudPcRemoteActionResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["actionName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActionName(val)
        return nil
    }
    res["actionState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseActionState)
        if err != nil {
            return err
        }
        cast := val.(ActionState)
        m.SetActionState(&cast)
        return nil
    }
    res["cloudPcId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCloudPcId(val)
        return nil
    }
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastUpdatedDateTime(val)
        return nil
    }
    res["managedDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetManagedDeviceId(val)
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetStartDateTime(val)
        return nil
    }
    res["statusDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudPcStatusDetails() })
        if err != nil {
            return err
        }
        m.SetStatusDetails(val.(*CloudPcStatusDetails))
        return nil
    }
    return res
}
func (m *CloudPcRemoteActionResult) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CloudPcRemoteActionResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("actionName", m.GetActionName())
        if err != nil {
            return err
        }
    }
    if m.GetActionState() != nil {
        cast := m.GetActionState().String()
        err := writer.WriteStringValue("actionState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudPcId", m.GetCloudPcId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managedDeviceId", m.GetManagedDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("statusDetails", m.GetStatusDetails())
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
// Sets the actionName property value. The specified action. Supported values: Reprovision, Resize.
// Parameters:
//  - value : Value to set for the actionName property.
func (m *CloudPcRemoteActionResult) SetActionName(value *string)() {
    m.actionName = value
}
// Sets the actionState property value. State of the action. Possible values are: None, pending, canceled, active, done, failed, notSupported. Read-only.
// Parameters:
//  - value : Value to set for the actionState property.
func (m *CloudPcRemoteActionResult) SetActionState(value *ActionState)() {
    m.actionState = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *CloudPcRemoteActionResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the cloudPcId property value. The ID of the Cloud PC device on which the remote action is performed. Read-only.
// Parameters:
//  - value : Value to set for the cloudPcId property.
func (m *CloudPcRemoteActionResult) SetCloudPcId(value *string)() {
    m.cloudPcId = value
}
// Sets the lastUpdatedDateTime property value. Last update time for action. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
// Parameters:
//  - value : Value to set for the lastUpdatedDateTime property.
func (m *CloudPcRemoteActionResult) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// Sets the managedDeviceId property value. The ID of the Intune managed device on which the remote action is performed. Read-only.
// Parameters:
//  - value : Value to set for the managedDeviceId property.
func (m *CloudPcRemoteActionResult) SetManagedDeviceId(value *string)() {
    m.managedDeviceId = value
}
// Sets the startDateTime property value. Time the action was initiated. The Timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as '2014-01-01T00:00:00Z'.
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *CloudPcRemoteActionResult) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// Sets the statusDetails property value. The details of the Cloud PC status.
// Parameters:
//  - value : Value to set for the statusDetails property.
func (m *CloudPcRemoteActionResult) SetStatusDetails(value *CloudPcStatusDetails)() {
    m.statusDetails = value
}
