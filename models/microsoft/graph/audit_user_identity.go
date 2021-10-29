package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AuditUserIdentity struct {
    UserIdentity
    // 
    homeTenantId *string;
    // 
    homeTenantName *string;
}
// Instantiates a new auditUserIdentity and sets the default values.
func NewAuditUserIdentity()(*AuditUserIdentity) {
    m := &AuditUserIdentity{
        UserIdentity: *NewUserIdentity(),
    }
    return m
}
// Gets the homeTenantId property value. 
func (m *AuditUserIdentity) GetHomeTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantId
    }
}
// Gets the homeTenantName property value. 
func (m *AuditUserIdentity) GetHomeTenantName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantName
    }
}
// The deserialization information for the current model
func (m *AuditUserIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.UserIdentity.GetFieldDeserializers()
    res["homeTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHomeTenantId(val)
        return nil
    }
    res["homeTenantName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHomeTenantName(val)
        return nil
    }
    return res
}
func (m *AuditUserIdentity) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AuditUserIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.UserIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("homeTenantId", m.GetHomeTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("homeTenantName", m.GetHomeTenantName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the homeTenantId property value. 
// Parameters:
//  - value : Value to set for the homeTenantId property.
func (m *AuditUserIdentity) SetHomeTenantId(value *string)() {
    m.homeTenantId = value
}
// Sets the homeTenantName property value. 
// Parameters:
//  - value : Value to set for the homeTenantName property.
func (m *AuditUserIdentity) SetHomeTenantName(value *string)() {
    m.homeTenantName = value
}
