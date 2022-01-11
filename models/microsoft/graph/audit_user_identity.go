package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuditUserIdentity 
type AuditUserIdentity struct {
    UserIdentity
    // For user sign ins, the identifier of the tenant that the user is a member of.
    homeTenantId *string;
    // For user sign ins, the name of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
    homeTenantName *string;
}
// NewAuditUserIdentity instantiates a new auditUserIdentity and sets the default values.
func NewAuditUserIdentity()(*AuditUserIdentity) {
    m := &AuditUserIdentity{
        UserIdentity: *NewUserIdentity(),
    }
    return m
}
// GetHomeTenantId gets the homeTenantId property value. For user sign ins, the identifier of the tenant that the user is a member of.
func (m *AuditUserIdentity) GetHomeTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantId
    }
}
// GetHomeTenantName gets the homeTenantName property value. For user sign ins, the name of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
func (m *AuditUserIdentity) GetHomeTenantName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.homeTenantName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuditUserIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.UserIdentity.GetFieldDeserializers()
    res["homeTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantId(val)
        }
        return nil
    }
    res["homeTenantName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeTenantName(val)
        }
        return nil
    }
    return res
}
func (m *AuditUserIdentity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetHomeTenantId sets the homeTenantId property value. For user sign ins, the identifier of the tenant that the user is a member of.
func (m *AuditUserIdentity) SetHomeTenantId(value *string)() {
    if m != nil {
        m.homeTenantId = value
    }
}
// SetHomeTenantName sets the homeTenantName property value. For user sign ins, the name of the tenant that the user is a member of. Only populated in cases where the home tenant has provided affirmative consent to Azure AD to show the tenant content.
func (m *AuditUserIdentity) SetHomeTenantName(value *string)() {
    if m != nil {
        m.homeTenantName = value
    }
}
