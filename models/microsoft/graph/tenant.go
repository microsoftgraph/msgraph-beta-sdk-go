package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Tenant 
type Tenant struct {
    Entity
    // The relationship details for the tenant with the managing entity.
    contract *TenantContract;
    // The date and time the tenant was created in the multi-tenant management platform. Optional. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The display name for the tenant. Required. Read-only.
    displayName *string;
    // The date and time the tenant was last updated within the multi-tenant management platform. Optional. Read-only.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
    tenantId *string;
    // The onboarding status information for the tenant. Optional. Read-only.
    tenantStatusInformation *TenantStatusInformation;
}
// NewTenant instantiates a new tenant and sets the default values.
func NewTenant()(*Tenant) {
    m := &Tenant{
        Entity: *NewEntity(),
    }
    return m
}
// GetContract gets the contract property value. The relationship details for the tenant with the managing entity.
func (m *Tenant) GetContract()(*TenantContract) {
    if m == nil {
        return nil
    } else {
        return m.contract
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the tenant was created in the multi-tenant management platform. Optional. Read-only.
func (m *Tenant) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDisplayName gets the displayName property value. The display name for the tenant. Required. Read-only.
func (m *Tenant) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The date and time the tenant was last updated within the multi-tenant management platform. Optional. Read-only.
func (m *Tenant) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdatedDateTime
    }
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *Tenant) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetTenantStatusInformation gets the tenantStatusInformation property value. The onboarding status information for the tenant. Optional. Read-only.
func (m *Tenant) GetTenantStatusInformation()(*TenantStatusInformation) {
    if m == nil {
        return nil
    } else {
        return m.tenantStatusInformation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Tenant) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contract"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTenantContract() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContract(val.(*TenantContract))
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
    res["lastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["tenantStatusInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTenantStatusInformation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantStatusInformation(val.(*TenantStatusInformation))
        }
        return nil
    }
    return res
}
func (m *Tenant) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Tenant) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("contract", m.GetContract())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("tenantStatusInformation", m.GetTenantStatusInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContract sets the contract property value. The relationship details for the tenant with the managing entity.
func (m *Tenant) SetContract(value *TenantContract)() {
    if m != nil {
        m.contract = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the tenant was created in the multi-tenant management platform. Optional. Read-only.
func (m *Tenant) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the tenant. Required. Read-only.
func (m *Tenant) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The date and time the tenant was last updated within the multi-tenant management platform. Optional. Read-only.
func (m *Tenant) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastUpdatedDateTime = value
    }
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant. Optional. Read-only.
func (m *Tenant) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
// SetTenantStatusInformation sets the tenantStatusInformation property value. The onboarding status information for the tenant. Optional. Read-only.
func (m *Tenant) SetTenantStatusInformation(value *TenantStatusInformation)() {
    if m != nil {
        m.tenantStatusInformation = value
    }
}
