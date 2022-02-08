package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamworkDeviceOperation 
type TeamworkDeviceOperation struct {
    Entity
    // Time at which the operation reached a final state (for example, Successful, Failed, and Cancelled).
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Identity of the user who created the device operation.
    createdBy *IdentitySet;
    // The UTC date and time when the device operation was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Error details are available only in case of a failed status.
    error *OperationError;
    // Identity of the user who last modified the device operation.
    lastActionBy *IdentitySet;
    // The UTC date and time when the device operation was last modified.
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Type of async operation on a device. The possible values are: deviceRestart, configUpdate, deviceDiagnostics, softwareUpdate, deviceManagementAgentConfigUpdate, remoteLogin, remoteLogout, unknownFutureValue.
    operationType *TeamworkDeviceOperationType;
    // Time at which the operation was started.
    startedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The current status of the async operation, for example, Queued, Scheduled, InProgress,  Successful, Cancelled, and Failed.
    status *string;
}
// NewTeamworkDeviceOperation instantiates a new teamworkDeviceOperation and sets the default values.
func NewTeamworkDeviceOperation()(*TeamworkDeviceOperation) {
    m := &TeamworkDeviceOperation{
        Entity: *NewEntity(),
    }
    return m
}
// GetCompletedDateTime gets the completedDateTime property value. Time at which the operation reached a final state (for example, Successful, Failed, and Cancelled).
func (m *TeamworkDeviceOperation) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the device operation.
func (m *TeamworkDeviceOperation) GetCreatedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The UTC date and time when the device operation was created.
func (m *TeamworkDeviceOperation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetError gets the error property value. Error details are available only in case of a failed status.
func (m *TeamworkDeviceOperation) GetError()(*OperationError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetLastActionBy gets the lastActionBy property value. Identity of the user who last modified the device operation.
func (m *TeamworkDeviceOperation) GetLastActionBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.lastActionBy
    }
}
// GetLastActionDateTime gets the lastActionDateTime property value. The UTC date and time when the device operation was last modified.
func (m *TeamworkDeviceOperation) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastActionDateTime
    }
}
// GetOperationType gets the operationType property value. Type of async operation on a device. The possible values are: deviceRestart, configUpdate, deviceDiagnostics, softwareUpdate, deviceManagementAgentConfigUpdate, remoteLogin, remoteLogout, unknownFutureValue.
func (m *TeamworkDeviceOperation) GetOperationType()(*TeamworkDeviceOperationType) {
    if m == nil {
        return nil
    } else {
        return m.operationType
    }
}
// GetStartedDateTime gets the startedDateTime property value. Time at which the operation was started.
func (m *TeamworkDeviceOperation) GetStartedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startedDateTime
    }
}
// GetStatus gets the status property value. The current status of the async operation, for example, Queued, Scheduled, InProgress,  Successful, Cancelled, and Failed.
func (m *TeamworkDeviceOperation) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkDeviceOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val)
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOperationError() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(*OperationError))
        }
        return nil
    }
    res["lastActionBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionBy(val.(*IdentitySet))
        }
        return nil
    }
    res["lastActionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["operationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkDeviceOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationType(val.(*TeamworkDeviceOperationType))
        }
        return nil
    }
    res["startedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartedDateTime(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
func (m *TeamworkDeviceOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkDeviceOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastActionBy", m.GetLastActionBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOperationType() != nil {
        cast := (*m.GetOperationType()).String()
        err = writer.WriteStringValue("operationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startedDateTime", m.GetStartedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletedDateTime sets the completedDateTime property value. Time at which the operation reached a final state (for example, Successful, Failed, and Cancelled).
func (m *TeamworkDeviceOperation) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.completedDateTime = value
    }
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the device operation.
func (m *TeamworkDeviceOperation) SetCreatedBy(value *IdentitySet)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The UTC date and time when the device operation was created.
func (m *TeamworkDeviceOperation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetError sets the error property value. Error details are available only in case of a failed status.
func (m *TeamworkDeviceOperation) SetError(value *OperationError)() {
    if m != nil {
        m.error = value
    }
}
// SetLastActionBy sets the lastActionBy property value. Identity of the user who last modified the device operation.
func (m *TeamworkDeviceOperation) SetLastActionBy(value *IdentitySet)() {
    if m != nil {
        m.lastActionBy = value
    }
}
// SetLastActionDateTime sets the lastActionDateTime property value. The UTC date and time when the device operation was last modified.
func (m *TeamworkDeviceOperation) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastActionDateTime = value
    }
}
// SetOperationType sets the operationType property value. Type of async operation on a device. The possible values are: deviceRestart, configUpdate, deviceDiagnostics, softwareUpdate, deviceManagementAgentConfigUpdate, remoteLogin, remoteLogout, unknownFutureValue.
func (m *TeamworkDeviceOperation) SetOperationType(value *TeamworkDeviceOperationType)() {
    if m != nil {
        m.operationType = value
    }
}
// SetStartedDateTime sets the startedDateTime property value. Time at which the operation was started.
func (m *TeamworkDeviceOperation) SetStartedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startedDateTime = value
    }
}
// SetStatus sets the status property value. The current status of the async operation, for example, Queued, Scheduled, InProgress,  Successful, Cancelled, and Failed.
func (m *TeamworkDeviceOperation) SetStatus(value *string)() {
    if m != nil {
        m.status = value
    }
}
