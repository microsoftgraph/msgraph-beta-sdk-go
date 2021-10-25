package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type GovernanceRoleDefinition struct {
    Entity
    displayName *string;
    externalId *string;
    resource *GovernanceResource;
    resourceId *string;
    roleSetting *GovernanceRoleSetting;
    templateId *string;
}
func NewGovernanceRoleDefinition()(*GovernanceRoleDefinition) {
    m := &GovernanceRoleDefinition{
        Entity: *NewEntity(),
    }
    return m
}
func (m *GovernanceRoleDefinition) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *GovernanceRoleDefinition) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
func (m *GovernanceRoleDefinition) GetResource()(*GovernanceResource) {
    if m == nil {
        return nil
    } else {
        return m.resource
    }
}
func (m *GovernanceRoleDefinition) GetResourceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceId
    }
}
func (m *GovernanceRoleDefinition) GetRoleSetting()(*GovernanceRoleSetting) {
    if m == nil {
        return nil
    } else {
        return m.roleSetting
    }
}
func (m *GovernanceRoleDefinition) GetTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.templateId
    }
}
func (m *GovernanceRoleDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
        return nil
    }
    res["resource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceResource() })
        if err != nil {
            return err
        }
        m.SetResource(val.(*GovernanceResource))
        return nil
    }
    res["resourceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceId(val)
        return nil
    }
    res["roleSetting"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewGovernanceRoleSetting() })
        if err != nil {
            return err
        }
        m.SetRoleSetting(val.(*GovernanceRoleSetting))
        return nil
    }
    res["templateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTemplateId(val)
        return nil
    }
    return res
}
func (m *GovernanceRoleDefinition) IsNil()(bool) {
    return m == nil
}
func (m *GovernanceRoleDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleSetting", m.GetRoleSetting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *GovernanceRoleDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *GovernanceRoleDefinition) SetExternalId(value *string)() {
    m.externalId = value
}
func (m *GovernanceRoleDefinition) SetResource(value *GovernanceResource)() {
    m.resource = value
}
func (m *GovernanceRoleDefinition) SetResourceId(value *string)() {
    m.resourceId = value
}
func (m *GovernanceRoleDefinition) SetRoleSetting(value *GovernanceRoleSetting)() {
    m.roleSetting = value
}
func (m *GovernanceRoleDefinition) SetTemplateId(value *string)() {
    m.templateId = value
}
