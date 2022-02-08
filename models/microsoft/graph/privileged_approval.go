package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrivilegedApproval 
type PrivilegedApproval struct {
    Entity
    // 
    approvalDuration *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // Possible values are: pending, approved, denied, aborted, canceled.
    approvalState *ApprovalState;
    // 
    approvalType *string;
    // 
    approverReason *string;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only. The role assignment request for this approval object
    request *PrivilegedRoleAssignmentRequest;
    // 
    requestorReason *string;
    // 
    roleId *string;
    // Read-only. Nullable.
    roleInfo *PrivilegedRole;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    userId *string;
}
// NewPrivilegedApproval instantiates a new privilegedApproval and sets the default values.
func NewPrivilegedApproval()(*PrivilegedApproval) {
    m := &PrivilegedApproval{
        Entity: *NewEntity(),
    }
    return m
}
// GetApprovalDuration gets the approvalDuration property value. 
func (m *PrivilegedApproval) GetApprovalDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.approvalDuration
    }
}
// GetApprovalState gets the approvalState property value. Possible values are: pending, approved, denied, aborted, canceled.
func (m *PrivilegedApproval) GetApprovalState()(*ApprovalState) {
    if m == nil {
        return nil
    } else {
        return m.approvalState
    }
}
// GetApprovalType gets the approvalType property value. 
func (m *PrivilegedApproval) GetApprovalType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.approvalType
    }
}
// GetApproverReason gets the approverReason property value. 
func (m *PrivilegedApproval) GetApproverReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.approverReason
    }
}
// GetEndDateTime gets the endDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PrivilegedApproval) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetRequest gets the request property value. Read-only. The role assignment request for this approval object
func (m *PrivilegedApproval) GetRequest()(*PrivilegedRoleAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.request
    }
}
// GetRequestorReason gets the requestorReason property value. 
func (m *PrivilegedApproval) GetRequestorReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestorReason
    }
}
// GetRoleId gets the roleId property value. 
func (m *PrivilegedApproval) GetRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleId
    }
}
// GetRoleInfo gets the roleInfo property value. Read-only. Nullable.
func (m *PrivilegedApproval) GetRoleInfo()(*PrivilegedRole) {
    if m == nil {
        return nil
    } else {
        return m.roleInfo
    }
}
// GetStartDateTime gets the startDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PrivilegedApproval) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetUserId gets the userId property value. 
func (m *PrivilegedApproval) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedApproval) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvalDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalDuration(val)
        }
        return nil
    }
    res["approvalState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseApprovalState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalState(val.(*ApprovalState))
        }
        return nil
    }
    res["approvalType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalType(val)
        }
        return nil
    }
    res["approverReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApproverReason(val)
        }
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["request"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedRoleAssignmentRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequest(val.(*PrivilegedRoleAssignmentRequest))
        }
        return nil
    }
    res["requestorReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestorReason(val)
        }
        return nil
    }
    res["roleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleId(val)
        }
        return nil
    }
    res["roleInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedRole() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleInfo(val.(*PrivilegedRole))
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
func (m *PrivilegedApproval) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrivilegedApproval) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteISODurationValue("approvalDuration", m.GetApprovalDuration())
        if err != nil {
            return err
        }
    }
    if m.GetApprovalState() != nil {
        cast := (*m.GetApprovalState()).String()
        err = writer.WriteStringValue("approvalState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("approvalType", m.GetApprovalType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("approverReason", m.GetApproverReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("request", m.GetRequest())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestorReason", m.GetRequestorReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleId", m.GetRoleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleInfo", m.GetRoleInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
// SetApprovalDuration sets the approvalDuration property value. 
func (m *PrivilegedApproval) SetApprovalDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.approvalDuration = value
    }
}
// SetApprovalState sets the approvalState property value. Possible values are: pending, approved, denied, aborted, canceled.
func (m *PrivilegedApproval) SetApprovalState(value *ApprovalState)() {
    if m != nil {
        m.approvalState = value
    }
}
// SetApprovalType sets the approvalType property value. 
func (m *PrivilegedApproval) SetApprovalType(value *string)() {
    if m != nil {
        m.approvalType = value
    }
}
// SetApproverReason sets the approverReason property value. 
func (m *PrivilegedApproval) SetApproverReason(value *string)() {
    if m != nil {
        m.approverReason = value
    }
}
// SetEndDateTime sets the endDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PrivilegedApproval) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetRequest sets the request property value. Read-only. The role assignment request for this approval object
func (m *PrivilegedApproval) SetRequest(value *PrivilegedRoleAssignmentRequest)() {
    if m != nil {
        m.request = value
    }
}
// SetRequestorReason sets the requestorReason property value. 
func (m *PrivilegedApproval) SetRequestorReason(value *string)() {
    if m != nil {
        m.requestorReason = value
    }
}
// SetRoleId sets the roleId property value. 
func (m *PrivilegedApproval) SetRoleId(value *string)() {
    if m != nil {
        m.roleId = value
    }
}
// SetRoleInfo sets the roleInfo property value. Read-only. Nullable.
func (m *PrivilegedApproval) SetRoleInfo(value *PrivilegedRole)() {
    if m != nil {
        m.roleInfo = value
    }
}
// SetStartDateTime sets the startDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PrivilegedApproval) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.startDateTime = value
    }
}
// SetUserId sets the userId property value. 
func (m *PrivilegedApproval) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
