package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementAutopilotPolicyStatusDetail 
type DeviceManagementAutopilotPolicyStatusDetail struct {
    Entity
    // The policy compliance status. Possible values are: unknown, compliant, installed, notCompliant, notInstalled, error.
    complianceStatus *DeviceManagementAutopilotPolicyComplianceStatus;
    // The friendly name of the policy.
    displayName *string;
    // The errorode associated with the compliance or enforcement status of the policy. Error code for enforcement status takes precedence if it exists.
    errorCode *int32;
    // Timestamp of the reported policy status
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The type of policy. Possible values are: unknown, application, appModel, configurationPolicy.
    policyType *DeviceManagementAutopilotPolicyType;
    // Indicates if this prolicy was tracked as part of the autopilot bootstrap enrollment sync session
    trackedOnEnrollmentStatus *bool;
}
// NewDeviceManagementAutopilotPolicyStatusDetail instantiates a new deviceManagementAutopilotPolicyStatusDetail and sets the default values.
func NewDeviceManagementAutopilotPolicyStatusDetail()(*DeviceManagementAutopilotPolicyStatusDetail) {
    m := &DeviceManagementAutopilotPolicyStatusDetail{
        Entity: *NewEntity(),
    }
    return m
}
// GetComplianceStatus gets the complianceStatus property value. The policy compliance status. Possible values are: unknown, compliant, installed, notCompliant, notInstalled, error.
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetComplianceStatus()(*DeviceManagementAutopilotPolicyComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.complianceStatus
    }
}
// GetDisplayName gets the displayName property value. The friendly name of the policy.
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetErrorCode gets the errorCode property value. The errorode associated with the compliance or enforcement status of the policy. Error code for enforcement status takes precedence if it exists.
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// GetLastReportedDateTime gets the lastReportedDateTime property value. Timestamp of the reported policy status
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastReportedDateTime
    }
}
// GetPolicyType gets the policyType property value. The type of policy. Possible values are: unknown, application, appModel, configurationPolicy.
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetPolicyType()(*DeviceManagementAutopilotPolicyType) {
    if m == nil {
        return nil
    } else {
        return m.policyType
    }
}
// GetTrackedOnEnrollmentStatus gets the trackedOnEnrollmentStatus property value. Indicates if this prolicy was tracked as part of the autopilot bootstrap enrollment sync session
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetTrackedOnEnrollmentStatus()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.trackedOnEnrollmentStatus
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["complianceStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementAutopilotPolicyComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceStatus(val.(*DeviceManagementAutopilotPolicyComplianceStatus))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["lastReportedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastReportedDateTime(val)
        }
        return nil
    }
    res["policyType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementAutopilotPolicyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyType(val.(*DeviceManagementAutopilotPolicyType))
        }
        return nil
    }
    res["trackedOnEnrollmentStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrackedOnEnrollmentStatus(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementAutopilotPolicyStatusDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetComplianceStatus() != nil {
        cast := (*m.GetComplianceStatus()).String()
        err = writer.WriteStringValue("complianceStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastReportedDateTime", m.GetLastReportedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPolicyType() != nil {
        cast := (*m.GetPolicyType()).String()
        err = writer.WriteStringValue("policyType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("trackedOnEnrollmentStatus", m.GetTrackedOnEnrollmentStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComplianceStatus sets the complianceStatus property value. The policy compliance status. Possible values are: unknown, compliant, installed, notCompliant, notInstalled, error.
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetComplianceStatus(value *DeviceManagementAutopilotPolicyComplianceStatus)() {
    if m != nil {
        m.complianceStatus = value
    }
}
// SetDisplayName sets the displayName property value. The friendly name of the policy.
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetErrorCode sets the errorCode property value. The errorode associated with the compliance or enforcement status of the policy. Error code for enforcement status takes precedence if it exists.
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetErrorCode(value *int32)() {
    if m != nil {
        m.errorCode = value
    }
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. Timestamp of the reported policy status
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastReportedDateTime = value
    }
}
// SetPolicyType sets the policyType property value. The type of policy. Possible values are: unknown, application, appModel, configurationPolicy.
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetPolicyType(value *DeviceManagementAutopilotPolicyType)() {
    if m != nil {
        m.policyType = value
    }
}
// SetTrackedOnEnrollmentStatus sets the trackedOnEnrollmentStatus property value. Indicates if this prolicy was tracked as part of the autopilot bootstrap enrollment sync session
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetTrackedOnEnrollmentStatus(value *bool)() {
    if m != nil {
        m.trackedOnEnrollmentStatus = value
    }
}
