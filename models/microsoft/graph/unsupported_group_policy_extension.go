package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UnsupportedGroupPolicyExtension struct {
    Entity
    extensionType *string;
    namespaceUrl *string;
    nodeName *string;
    settingScope *GroupPolicySettingScope;
}
func NewUnsupportedGroupPolicyExtension()(*UnsupportedGroupPolicyExtension) {
    m := &UnsupportedGroupPolicyExtension{
        Entity: *NewEntity(),
    }
    return m
}
func (m *UnsupportedGroupPolicyExtension) GetExtensionType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionType
    }
}
func (m *UnsupportedGroupPolicyExtension) GetNamespaceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.namespaceUrl
    }
}
func (m *UnsupportedGroupPolicyExtension) GetNodeName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nodeName
    }
}
func (m *UnsupportedGroupPolicyExtension) GetSettingScope()(*GroupPolicySettingScope) {
    if m == nil {
        return nil
    } else {
        return m.settingScope
    }
}
func (m *UnsupportedGroupPolicyExtension) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["extensionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExtensionType(val)
        return nil
    }
    res["namespaceUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNamespaceUrl(val)
        return nil
    }
    res["nodeName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNodeName(val)
        return nil
    }
    res["settingScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicySettingScope)
        if err != nil {
            return err
        }
        cast := val.(GroupPolicySettingScope)
        m.SetSettingScope(&cast)
        return nil
    }
    return res
}
func (m *UnsupportedGroupPolicyExtension) IsNil()(bool) {
    return m == nil
}
func (m *UnsupportedGroupPolicyExtension) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("extensionType", m.GetExtensionType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("namespaceUrl", m.GetNamespaceUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nodeName", m.GetNodeName())
        if err != nil {
            return err
        }
    }
    if m.GetSettingScope() != nil {
        cast := m.GetSettingScope().String()
        err = writer.WriteStringValue("settingScope", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UnsupportedGroupPolicyExtension) SetExtensionType(value *string)() {
    m.extensionType = value
}
func (m *UnsupportedGroupPolicyExtension) SetNamespaceUrl(value *string)() {
    m.namespaceUrl = value
}
func (m *UnsupportedGroupPolicyExtension) SetNodeName(value *string)() {
    m.nodeName = value
}
func (m *UnsupportedGroupPolicyExtension) SetSettingScope(value *GroupPolicySettingScope)() {
    m.settingScope = value
}
