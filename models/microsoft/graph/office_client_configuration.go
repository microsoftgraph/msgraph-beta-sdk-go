package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OfficeClientConfiguration struct {
    Entity
    // The list of group assignments for the policy.
    assignments []OfficeClientConfigurationAssignment;
    // List of office Client check-in status.
    checkinStatuses []OfficeClientCheckinStatus;
    // Not yet documented
    description *string;
    // Admin provided description of the office client configuration policy.
    displayName *string;
    // Policy settings JSON string in binary format, these values cannot be changed by the user.
    policyPayload []byte;
    // Priority value should be unique value for each policy under a tenant and will be used for conflict resolution, lower values mean priority is high.
    priority *int32;
    // User check-in summary for the policy.
    userCheckinSummary *OfficeUserCheckinSummary;
    // Preference settings JSON string in binary format, these values can be overridden by the user.
    userPreferencePayload []byte;
}
// Instantiates a new officeClientConfiguration and sets the default values.
func NewOfficeClientConfiguration()(*OfficeClientConfiguration) {
    m := &OfficeClientConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignments property value. The list of group assignments for the policy.
func (m *OfficeClientConfiguration) GetAssignments()([]OfficeClientConfigurationAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the checkinStatuses property value. List of office Client check-in status.
func (m *OfficeClientConfiguration) GetCheckinStatuses()([]OfficeClientCheckinStatus) {
    if m == nil {
        return nil
    } else {
        return m.checkinStatuses
    }
}
// Gets the description property value. Not yet documented
func (m *OfficeClientConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Admin provided description of the office client configuration policy.
func (m *OfficeClientConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the policyPayload property value. Policy settings JSON string in binary format, these values cannot be changed by the user.
func (m *OfficeClientConfiguration) GetPolicyPayload()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.policyPayload
    }
}
// Gets the priority property value. Priority value should be unique value for each policy under a tenant and will be used for conflict resolution, lower values mean priority is high.
func (m *OfficeClientConfiguration) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// Gets the userCheckinSummary property value. User check-in summary for the policy.
func (m *OfficeClientConfiguration) GetUserCheckinSummary()(*OfficeUserCheckinSummary) {
    if m == nil {
        return nil
    } else {
        return m.userCheckinSummary
    }
}
// Gets the userPreferencePayload property value. Preference settings JSON string in binary format, these values can be overridden by the user.
func (m *OfficeClientConfiguration) GetUserPreferencePayload()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.userPreferencePayload
    }
}
// The deserialization information for the current model
func (m *OfficeClientConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeClientConfigurationAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfficeClientConfigurationAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*OfficeClientConfigurationAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["checkinStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeClientCheckinStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfficeClientCheckinStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*OfficeClientCheckinStatus))
            }
            m.SetCheckinStatuses(res)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["policyPayload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyPayload(val)
        }
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["userCheckinSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeUserCheckinSummary() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserCheckinSummary(val.(*OfficeUserCheckinSummary))
        }
        return nil
    }
    res["userPreferencePayload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPreferencePayload(val)
        }
        return nil
    }
    return res
}
func (m *OfficeClientConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OfficeClientConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCheckinStatuses()))
        for i, v := range m.GetCheckinStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("checkinStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteByteArrayValue("policyPayload", m.GetPolicyPayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userCheckinSummary", m.GetUserCheckinSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("userPreferencePayload", m.GetUserPreferencePayload())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignments property value. The list of group assignments for the policy.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *OfficeClientConfiguration) SetAssignments(value []OfficeClientConfigurationAssignment)() {
    m.assignments = value
}
// Sets the checkinStatuses property value. List of office Client check-in status.
// Parameters:
//  - value : Value to set for the checkinStatuses property.
func (m *OfficeClientConfiguration) SetCheckinStatuses(value []OfficeClientCheckinStatus)() {
    m.checkinStatuses = value
}
// Sets the description property value. Not yet documented
// Parameters:
//  - value : Value to set for the description property.
func (m *OfficeClientConfiguration) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Admin provided description of the office client configuration policy.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *OfficeClientConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the policyPayload property value. Policy settings JSON string in binary format, these values cannot be changed by the user.
// Parameters:
//  - value : Value to set for the policyPayload property.
func (m *OfficeClientConfiguration) SetPolicyPayload(value []byte)() {
    m.policyPayload = value
}
// Sets the priority property value. Priority value should be unique value for each policy under a tenant and will be used for conflict resolution, lower values mean priority is high.
// Parameters:
//  - value : Value to set for the priority property.
func (m *OfficeClientConfiguration) SetPriority(value *int32)() {
    m.priority = value
}
// Sets the userCheckinSummary property value. User check-in summary for the policy.
// Parameters:
//  - value : Value to set for the userCheckinSummary property.
func (m *OfficeClientConfiguration) SetUserCheckinSummary(value *OfficeUserCheckinSummary)() {
    m.userCheckinSummary = value
}
// Sets the userPreferencePayload property value. Preference settings JSON string in binary format, these values can be overridden by the user.
// Parameters:
//  - value : Value to set for the userPreferencePayload property.
func (m *OfficeClientConfiguration) SetUserPreferencePayload(value []byte)() {
    m.userPreferencePayload = value
}
