package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type TenantGroup struct {
    Entity
    allTenantsIncluded *bool;
    displayName *string;
    managementActions []ManagementActionInfo;
    managementIntents []ManagementIntentInfo;
    tenantIds []string;
}
func NewTenantGroup()(*TenantGroup) {
    m := &TenantGroup{
        Entity: *NewEntity(),
    }
    return m
}
func (m *TenantGroup) GetAllTenantsIncluded()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allTenantsIncluded
    }
}
func (m *TenantGroup) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *TenantGroup) GetManagementActions()([]ManagementActionInfo) {
    if m == nil {
        return nil
    } else {
        return m.managementActions
    }
}
func (m *TenantGroup) GetManagementIntents()([]ManagementIntentInfo) {
    if m == nil {
        return nil
    } else {
        return m.managementIntents
    }
}
func (m *TenantGroup) GetTenantIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.tenantIds
    }
}
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
            res[i] = v.(string)
        }
        m.SetTenantIds(res)
        return nil
    }
    return res
}
func (m *TenantGroup) IsNil()(bool) {
    return m == nil
}
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
func (m *TenantGroup) SetAllTenantsIncluded(value *bool)() {
    m.allTenantsIncluded = value
}
func (m *TenantGroup) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *TenantGroup) SetManagementActions(value []ManagementActionInfo)() {
    m.managementActions = value
}
func (m *TenantGroup) SetManagementIntents(value []ManagementIntentInfo)() {
    m.managementIntents = value
}
func (m *TenantGroup) SetTenantIds(value []string)() {
    m.tenantIds = value
}
