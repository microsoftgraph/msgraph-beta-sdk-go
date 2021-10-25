package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TenantRelationships struct {
    additionalData map[string]interface{};
    managedTenants *ManagedTenant;
}
func NewTenantRelationships()(*TenantRelationships) {
    m := &TenantRelationships{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *TenantRelationships) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *TenantRelationships) GetManagedTenants()(*ManagedTenant) {
    if m == nil {
        return nil
    } else {
        return m.managedTenants
    }
}
func (m *TenantRelationships) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["managedTenants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedTenant() })
        if err != nil {
            return err
        }
        m.SetManagedTenants(val.(*ManagedTenant))
        return nil
    }
    return res
}
func (m *TenantRelationships) IsNil()(bool) {
    return m == nil
}
func (m *TenantRelationships) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("managedTenants", m.GetManagedTenants())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *TenantRelationships) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *TenantRelationships) SetManagedTenants(value *ManagedTenant)() {
    m.managedTenants = value
}
