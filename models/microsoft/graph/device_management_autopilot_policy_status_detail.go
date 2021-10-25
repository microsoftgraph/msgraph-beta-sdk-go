package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagementAutopilotPolicyStatusDetail struct {
    Entity
    complianceStatus *DeviceManagementAutopilotPolicyComplianceStatus;
    displayName *string;
    errorCode *int32;
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    policyType *DeviceManagementAutopilotPolicyType;
    trackedOnEnrollmentStatus *bool;
}
func NewDeviceManagementAutopilotPolicyStatusDetail()(*DeviceManagementAutopilotPolicyStatusDetail) {
    m := &DeviceManagementAutopilotPolicyStatusDetail{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetComplianceStatus()(*DeviceManagementAutopilotPolicyComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.complianceStatus
    }
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetErrorCode()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastReportedDateTime
    }
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetPolicyType()(*DeviceManagementAutopilotPolicyType) {
    if m == nil {
        return nil
    } else {
        return m.policyType
    }
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetTrackedOnEnrollmentStatus()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.trackedOnEnrollmentStatus
    }
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["complianceStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementAutopilotPolicyComplianceStatus)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementAutopilotPolicyComplianceStatus)
        m.SetComplianceStatus(&cast)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetErrorCode(val)
        return nil
    }
    res["lastReportedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastReportedDateTime(val)
        return nil
    }
    res["policyType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementAutopilotPolicyType)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementAutopilotPolicyType)
        m.SetPolicyType(&cast)
        return nil
    }
    res["trackedOnEnrollmentStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetTrackedOnEnrollmentStatus(val)
        return nil
    }
    return res
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) IsNil()(bool) {
    return m == nil
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetComplianceStatus() != nil {
        cast := m.GetComplianceStatus().String()
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
        cast := m.GetPolicyType().String()
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
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetComplianceStatus(value *DeviceManagementAutopilotPolicyComplianceStatus)() {
    m.complianceStatus = value
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetErrorCode(value *int32)() {
    m.errorCode = value
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetPolicyType(value *DeviceManagementAutopilotPolicyType)() {
    m.policyType = value
}
func (m *DeviceManagementAutopilotPolicyStatusDetail) SetTrackedOnEnrollmentStatus(value *bool)() {
    m.trackedOnEnrollmentStatus = value
}
