package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrivilegedApproval struct {
    Entity
    // 
    approvalDuration *string;
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
// Instantiates a new privilegedApproval and sets the default values.
func NewPrivilegedApproval()(*PrivilegedApproval) {
    m := &PrivilegedApproval{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the approvalDuration property value. 
func (m *PrivilegedApproval) GetApprovalDuration()(*string) {
    if m == nil {
        return nil
    } else {
        return m.approvalDuration
    }
}
// Gets the approvalState property value. Possible values are: pending, approved, denied, aborted, canceled.
func (m *PrivilegedApproval) GetApprovalState()(*ApprovalState) {
    if m == nil {
        return nil
    } else {
        return m.approvalState
    }
}
// Gets the approvalType property value. 
func (m *PrivilegedApproval) GetApprovalType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.approvalType
    }
}
// Gets the approverReason property value. 
func (m *PrivilegedApproval) GetApproverReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.approverReason
    }
}
// Gets the endDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PrivilegedApproval) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// Gets the request property value. Read-only. The role assignment request for this approval object
func (m *PrivilegedApproval) GetRequest()(*PrivilegedRoleAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.request
    }
}
// Gets the requestorReason property value. 
func (m *PrivilegedApproval) GetRequestorReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestorReason
    }
}
// Gets the roleId property value. 
func (m *PrivilegedApproval) GetRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleId
    }
}
// Gets the roleInfo property value. Read-only. Nullable.
func (m *PrivilegedApproval) GetRoleInfo()(*PrivilegedRole) {
    if m == nil {
        return nil
    } else {
        return m.roleInfo
    }
}
// Gets the startDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *PrivilegedApproval) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the userId property value. 
func (m *PrivilegedApproval) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// The deserialization information for the current model
func (m *PrivilegedApproval) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["approvalDuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApprovalDuration(val)
        return nil
    }
    res["approvalState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseApprovalState)
        if err != nil {
            return err
        }
        cast := val.(ApprovalState)
        m.SetApprovalState(&cast)
        return nil
    }
    res["approvalType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApprovalType(val)
        return nil
    }
    res["approverReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetApproverReason(val)
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEndDateTime(val)
        return nil
    }
    res["request"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedRoleAssignmentRequest() })
        if err != nil {
            return err
        }
        m.SetRequest(val.(*PrivilegedRoleAssignmentRequest))
        return nil
    }
    res["requestorReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRequestorReason(val)
        return nil
    }
    res["roleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRoleId(val)
        return nil
    }
    res["roleInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrivilegedRole() })
        if err != nil {
            return err
        }
        m.SetRoleInfo(val.(*PrivilegedRole))
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
func (m *PrivilegedApproval) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrivilegedApproval) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("approvalDuration", m.GetApprovalDuration())
        if err != nil {
            return err
        }
    }
    if m.GetApprovalState() != nil {
        cast := m.GetApprovalState().String()
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
// Sets the approvalDuration property value. 
// Parameters:
//  - value : Value to set for the approvalDuration property.
func (m *PrivilegedApproval) SetApprovalDuration(value *string)() {
    m.approvalDuration = value
}
// Sets the approvalState property value. Possible values are: pending, approved, denied, aborted, canceled.
// Parameters:
//  - value : Value to set for the approvalState property.
func (m *PrivilegedApproval) SetApprovalState(value *ApprovalState)() {
    m.approvalState = value
}
// Sets the approvalType property value. 
// Parameters:
//  - value : Value to set for the approvalType property.
func (m *PrivilegedApproval) SetApprovalType(value *string)() {
    m.approvalType = value
}
// Sets the approverReason property value. 
// Parameters:
//  - value : Value to set for the approverReason property.
func (m *PrivilegedApproval) SetApproverReason(value *string)() {
    m.approverReason = value
}
// Sets the endDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the endDateTime property.
func (m *PrivilegedApproval) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// Sets the request property value. Read-only. The role assignment request for this approval object
// Parameters:
//  - value : Value to set for the request property.
func (m *PrivilegedApproval) SetRequest(value *PrivilegedRoleAssignmentRequest)() {
    m.request = value
}
// Sets the requestorReason property value. 
// Parameters:
//  - value : Value to set for the requestorReason property.
func (m *PrivilegedApproval) SetRequestorReason(value *string)() {
    m.requestorReason = value
}
// Sets the roleId property value. 
// Parameters:
//  - value : Value to set for the roleId property.
func (m *PrivilegedApproval) SetRoleId(value *string)() {
    m.roleId = value
}
// Sets the roleInfo property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the roleInfo property.
func (m *PrivilegedApproval) SetRoleInfo(value *PrivilegedRole)() {
    m.roleInfo = value
}
// Sets the startDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *PrivilegedApproval) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// Sets the userId property value. 
// Parameters:
//  - value : Value to set for the userId property.
func (m *PrivilegedApproval) SetUserId(value *string)() {
    m.userId = value
}
