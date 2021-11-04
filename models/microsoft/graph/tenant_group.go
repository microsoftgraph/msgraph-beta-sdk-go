package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TenantGroup struct {
    Entity
    // A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
    allTenantsIncluded *bool;
    // The display name for the tenant group. Optional. Read-only.
    displayName *string;
    // The collection of management action associated with the tenant group. Optional. Read-only.
    managementActions []ManagementActionInfo;
    // The collection of management intents associated with the tenant group. Optional. Read-only.
    managementIntents []ManagementIntentInfo;
    // The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
    tenantIds []string;
}
// Instantiates a new tenantGroup and sets the default values.
func NewTenantGroup()(*TenantGroup) {
    m := &TenantGroup{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the allTenantsIncluded property value. A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
func (m *TenantGroup) GetAllTenantsIncluded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allTenantsIncluded
    }
}
// Gets the displayName property value. The display name for the tenant group. Optional. Read-only.
func (m *TenantGroup) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the managementActions property value. The collection of management action associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) GetManagementActions()([]ManagementActionInfo) {
    if m == nil {
        return nil
    } else {
        return m.managementActions
    }
}
// Gets the managementIntents property value. The collection of management intents associated with the tenant group. Optional. Read-only.
func (m *TenantGroup) GetManagementIntents()([]ManagementIntentInfo) {
    if m == nil {
        return nil
    } else {
        return m.managementIntents
    }
}
// Gets the tenantIds property value. The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
func (m *TenantGroup) GetTenantIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tenantIds
    }
}
// The deserialization information for the current model
func (m *TenantGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allTenantsIncluded"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllTenantsIncluded(val)
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
    res["managementActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementActionInfo() })
        if err != nil {
            return err
        }
        res := make([]ManagementActionInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagementActionInfo))
        }
        m.SetManagementActions(res)
        return nil
    }
    res["managementIntents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementIntentInfo() })
        if err != nil {
            return err
        }
        res := make([]ManagementIntentInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagementIntentInfo))
        }
        m.SetManagementIntents(res)
        return nil
    }
    res["tenantIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetTenantIds(res)
        return nil
    }
    return res
}
func (m *TenantGroup) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TenantGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allTenantsIncluded", m.GetAllTenantsIncluded())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementActions()))
        for i, v := range m.GetManagementActions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementActions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementIntents()))
        for i, v := range m.GetManagementIntents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementIntents", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("tenantIds", m.GetTenantIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the allTenantsIncluded property value. A flag indicating whether all managed tenant are included in the tenant group. Required. Read-only.
// Parameters:
//  - value : Value to set for the allTenantsIncluded property.
func (m *TenantGroup) SetAllTenantsIncluded(value *bool)() {
    m.allTenantsIncluded = value
}
// Sets the displayName property value. The display name for the tenant group. Optional. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *TenantGroup) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the managementActions property value. The collection of management action associated with the tenant group. Optional. Read-only.
// Parameters:
//  - value : Value to set for the managementActions property.
func (m *TenantGroup) SetManagementActions(value []ManagementActionInfo)() {
    m.managementActions = value
}
// Sets the managementIntents property value. The collection of management intents associated with the tenant group. Optional. Read-only.
// Parameters:
//  - value : Value to set for the managementIntents property.
func (m *TenantGroup) SetManagementIntents(value []ManagementIntentInfo)() {
    m.managementIntents = value
}
// Sets the tenantIds property value. The collection of managed tenant identifiers include in the tenant group. Optional. Read-only.
// Parameters:
//  - value : Value to set for the tenantIds property.
func (m *TenantGroup) SetTenantIds(value []string)() {
    m.tenantIds = value
}
