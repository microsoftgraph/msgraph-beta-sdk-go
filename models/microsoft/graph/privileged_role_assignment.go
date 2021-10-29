package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PrivilegedRoleAssignment struct {
    Entity
    // The UTC DateTime when the temporary privileged role assignment will be expired. For permanent role assignment, the value is null.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // true if the role assignment is activated. false if the role assignment is deactivated.
    isElevated *bool;
    // Result message set by the service.
    resultMessage *string;
    // Role identifier. In GUID string format.
    roleId *string;
    // Read-only. Nullable. The associated role information.
    roleInfo *PrivilegedRole;
    // User identifier. In GUID string format.
    userId *string;
}
// Instantiates a new privilegedRoleAssignment and sets the default values.
func NewPrivilegedRoleAssignment()(*PrivilegedRoleAssignment) {
    m := &PrivilegedRoleAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the expirationDateTime property value. The UTC DateTime when the temporary privileged role assignment will be expired. For permanent role assignment, the value is null.
func (m *PrivilegedRoleAssignment) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the isElevated property value. true if the role assignment is activated. false if the role assignment is deactivated.
func (m *PrivilegedRoleAssignment) GetIsElevated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isElevated
    }
}
// Gets the resultMessage property value. Result message set by the service.
func (m *PrivilegedRoleAssignment) GetResultMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resultMessage
    }
}
// Gets the roleId property value. Role identifier. In GUID string format.
func (m *PrivilegedRoleAssignment) GetRoleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.roleId
    }
}
// Gets the roleInfo property value. Read-only. Nullable. The associated role information.
func (m *PrivilegedRoleAssignment) GetRoleInfo()(*PrivilegedRole) {
    if m == nil {
        return nil
    } else {
        return m.roleInfo
    }
}
// Gets the userId property value. User identifier. In GUID string format.
func (m *PrivilegedRoleAssignment) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// The deserialization information for the current model
func (m *PrivilegedRoleAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["isElevated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsElevated(val)
        return nil
    }
    res["resultMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResultMessage(val)
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
func (m *PrivilegedRoleAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PrivilegedRoleAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isElevated", m.GetIsElevated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resultMessage", m.GetResultMessage())
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
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the expirationDateTime property value. The UTC DateTime when the temporary privileged role assignment will be expired. For permanent role assignment, the value is null.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *PrivilegedRoleAssignment) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the isElevated property value. true if the role assignment is activated. false if the role assignment is deactivated.
// Parameters:
//  - value : Value to set for the isElevated property.
func (m *PrivilegedRoleAssignment) SetIsElevated(value *bool)() {
    m.isElevated = value
}
// Sets the resultMessage property value. Result message set by the service.
// Parameters:
//  - value : Value to set for the resultMessage property.
func (m *PrivilegedRoleAssignment) SetResultMessage(value *string)() {
    m.resultMessage = value
}
// Sets the roleId property value. Role identifier. In GUID string format.
// Parameters:
//  - value : Value to set for the roleId property.
func (m *PrivilegedRoleAssignment) SetRoleId(value *string)() {
    m.roleId = value
}
// Sets the roleInfo property value. Read-only. Nullable. The associated role information.
// Parameters:
//  - value : Value to set for the roleInfo property.
func (m *PrivilegedRoleAssignment) SetRoleInfo(value *PrivilegedRole)() {
    m.roleInfo = value
}
// Sets the userId property value. User identifier. In GUID string format.
// Parameters:
//  - value : Value to set for the userId property.
func (m *PrivilegedRoleAssignment) SetUserId(value *string)() {
    m.userId = value
}
