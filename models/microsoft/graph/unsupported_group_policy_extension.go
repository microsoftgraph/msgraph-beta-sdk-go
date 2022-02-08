package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnsupportedGroupPolicyExtension 
type UnsupportedGroupPolicyExtension struct {
    Entity
    // ExtensionType of the unsupported extension.
    extensionType *string;
    // Namespace Url of the unsupported extension.
    namespaceUrl *string;
    // Node name of the unsupported extension.
    nodeName *string;
    // Setting Scope of the unsupported extension. Possible values are: unknown, device, user.
    settingScope *GroupPolicySettingScope;
}
// NewUnsupportedGroupPolicyExtension instantiates a new unsupportedGroupPolicyExtension and sets the default values.
func NewUnsupportedGroupPolicyExtension()(*UnsupportedGroupPolicyExtension) {
    m := &UnsupportedGroupPolicyExtension{
        Entity: *NewEntity(),
    }
    return m
}
// GetExtensionType gets the extensionType property value. ExtensionType of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) GetExtensionType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionType
    }
}
// GetNamespaceUrl gets the namespaceUrl property value. Namespace Url of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) GetNamespaceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.namespaceUrl
    }
}
// GetNodeName gets the nodeName property value. Node name of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) GetNodeName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nodeName
    }
}
// GetSettingScope gets the settingScope property value. Setting Scope of the unsupported extension. Possible values are: unknown, device, user.
func (m *UnsupportedGroupPolicyExtension) GetSettingScope()(*GroupPolicySettingScope) {
    if m == nil {
        return nil
    } else {
        return m.settingScope
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnsupportedGroupPolicyExtension) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["extensionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionType(val)
        }
        return nil
    }
    res["namespaceUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNamespaceUrl(val)
        }
        return nil
    }
    res["nodeName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeName(val)
        }
        return nil
    }
    res["settingScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupPolicySettingScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingScope(val.(*GroupPolicySettingScope))
        }
        return nil
    }
    return res
}
func (m *UnsupportedGroupPolicyExtension) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetSettingScope()).String()
        err = writer.WriteStringValue("settingScope", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExtensionType sets the extensionType property value. ExtensionType of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) SetExtensionType(value *string)() {
    if m != nil {
        m.extensionType = value
    }
}
// SetNamespaceUrl sets the namespaceUrl property value. Namespace Url of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) SetNamespaceUrl(value *string)() {
    if m != nil {
        m.namespaceUrl = value
    }
}
// SetNodeName sets the nodeName property value. Node name of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) SetNodeName(value *string)() {
    if m != nil {
        m.nodeName = value
    }
}
// SetSettingScope sets the settingScope property value. Setting Scope of the unsupported extension. Possible values are: unknown, device, user.
func (m *UnsupportedGroupPolicyExtension) SetSettingScope(value *GroupPolicySettingScope)() {
    if m != nil {
        m.settingScope = value
    }
}
