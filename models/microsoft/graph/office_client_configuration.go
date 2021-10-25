package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OfficeClientConfiguration struct {
    Entity
    assignments []OfficeClientConfigurationAssignment;
    checkinStatuses []OfficeClientCheckinStatus;
    description *string;
    displayName *string;
    policyPayload []byte;
    priority *int32;
    userCheckinSummary *OfficeUserCheckinSummary;
    userPreferencePayload []byte;
}
func NewOfficeClientConfiguration()(*OfficeClientConfiguration) {
    m := &OfficeClientConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
func (m *OfficeClientConfiguration) GetAssignments()([]OfficeClientConfigurationAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
func (m *OfficeClientConfiguration) GetCheckinStatuses()([]OfficeClientCheckinStatus) {
    if m == nil {
        return nil
    } else {
        return m.checkinStatuses
    }
}
func (m *OfficeClientConfiguration) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *OfficeClientConfiguration) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *OfficeClientConfiguration) GetPolicyPayload()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.policyPayload
    }
}
func (m *OfficeClientConfiguration) GetPriority()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
func (m *OfficeClientConfiguration) GetUserCheckinSummary()(*OfficeUserCheckinSummary) {
    if m == nil {
        return nil
    } else {
        return m.userCheckinSummary
    }
}
func (m *OfficeClientConfiguration) GetUserPreferencePayload()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.userPreferencePayload
    }
}
func (m *OfficeClientConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeClientConfigurationAssignment() })
        if err != nil {
            return err
        }
        res := make([]OfficeClientConfigurationAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*OfficeClientConfigurationAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["checkinStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeClientCheckinStatus() })
        if err != nil {
            return err
        }
        res := make([]OfficeClientCheckinStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*OfficeClientCheckinStatus))
        }
        m.SetCheckinStatuses(res)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
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
    res["policyPayload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetPolicyPayload(val)
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetPriority(val)
        return nil
    }
    res["userCheckinSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfficeUserCheckinSummary() })
        if err != nil {
            return err
        }
        m.SetUserCheckinSummary(val.(*OfficeUserCheckinSummary))
        return nil
    }
    res["userPreferencePayload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetUserPreferencePayload(val)
        return nil
    }
    return res
}
func (m *OfficeClientConfiguration) IsNil()(bool) {
    return m == nil
}
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
func (m *OfficeClientConfiguration) SetAssignments(value []OfficeClientConfigurationAssignment)() {
    m.assignments = value
}
func (m *OfficeClientConfiguration) SetCheckinStatuses(value []OfficeClientCheckinStatus)() {
    m.checkinStatuses = value
}
func (m *OfficeClientConfiguration) SetDescription(value *string)() {
    m.description = value
}
func (m *OfficeClientConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *OfficeClientConfiguration) SetPolicyPayload(value []byte)() {
    m.policyPayload = value
}
func (m *OfficeClientConfiguration) SetPriority(value *int32)() {
    m.priority = value
}
func (m *OfficeClientConfiguration) SetUserCheckinSummary(value *OfficeUserCheckinSummary)() {
    m.userCheckinSummary = value
}
func (m *OfficeClientConfiguration) SetUserPreferencePayload(value []byte)() {
    m.userPreferencePayload = value
}
