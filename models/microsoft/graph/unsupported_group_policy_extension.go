package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new unsupportedGroupPolicyExtension and sets the default values.
func NewUnsupportedGroupPolicyExtension()(*UnsupportedGroupPolicyExtension) {
    m := &UnsupportedGroupPolicyExtension{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the extensionType property value. ExtensionType of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) GetExtensionType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.extensionType
    }
}
// Gets the namespaceUrl property value. Namespace Url of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) GetNamespaceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.namespaceUrl
    }
}
// Gets the nodeName property value. Node name of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) GetNodeName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.nodeName
    }
}
// Gets the settingScope property value. Setting Scope of the unsupported extension. Possible values are: unknown, device, user.
func (m *UnsupportedGroupPolicyExtension) GetSettingScope()(*GroupPolicySettingScope) {
    if m == nil {
        return nil
    } else {
        return m.settingScope
    }
}
// The deserialization information for the current model
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
            cast := val.(GroupPolicySettingScope)
            m.SetSettingScope(&cast)
        }
        return nil
    }
    return res
}
func (m *UnsupportedGroupPolicyExtension) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the extensionType property value. ExtensionType of the unsupported extension.
// Parameters:
//  - value : Value to set for the extensionType property.
func (m *UnsupportedGroupPolicyExtension) SetExtensionType(value *string)() {
    m.extensionType = value
}
// Sets the namespaceUrl property value. Namespace Url of the unsupported extension.
// Parameters:
//  - value : Value to set for the namespaceUrl property.
func (m *UnsupportedGroupPolicyExtension) SetNamespaceUrl(value *string)() {
    m.namespaceUrl = value
}
// Sets the nodeName property value. Node name of the unsupported extension.
// Parameters:
//  - value : Value to set for the nodeName property.
func (m *UnsupportedGroupPolicyExtension) SetNodeName(value *string)() {
    m.nodeName = value
}
// Sets the settingScope property value. Setting Scope of the unsupported extension. Possible values are: unknown, device, user.
// Parameters:
//  - value : Value to set for the settingScope property.
func (m *UnsupportedGroupPolicyExtension) SetSettingScope(value *GroupPolicySettingScope)() {
    m.settingScope = value
}
