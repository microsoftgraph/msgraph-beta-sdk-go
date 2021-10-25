package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ManagementIntent struct {
    Entity
    displayName *string;
    isGlobal *bool;
    managementTemplates []ManagementTemplateDetailedInfo;
}
func NewManagementIntent()(*ManagementIntent) {
    m := &ManagementIntent{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ManagementIntent) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *ManagementIntent) GetIsGlobal()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isGlobal
    }
}
func (m *ManagementIntent) GetManagementTemplates()([]ManagementTemplateDetailedInfo) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplates
    }
}
func (m *ManagementIntent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["isGlobal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsGlobal(val)
        return nil
    }
    res["managementTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagementTemplateDetailedInfo() })
        if err != nil {
            return err
        }
        res := make([]ManagementTemplateDetailedInfo, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagementTemplateDetailedInfo))
        }
        m.SetManagementTemplates(res)
        return nil
    }
    return res
}
func (m *ManagementIntent) IsNil()(bool) {
    return m == nil
}
func (m *ManagementIntent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isGlobal", m.GetIsGlobal())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagementTemplates()))
        for i, v := range m.GetManagementTemplates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managementTemplates", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ManagementIntent) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *ManagementIntent) SetIsGlobal(value *bool)() {
    m.isGlobal = value
}
func (m *ManagementIntent) SetManagementTemplates(value []ManagementTemplateDetailedInfo)() {
    m.managementTemplates = value
}
