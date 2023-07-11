package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnsupportedGroupPolicyExtension unsupported Group Policy Extension.
type UnsupportedGroupPolicyExtension struct {
    Entity
}
// NewUnsupportedGroupPolicyExtension instantiates a new unsupportedGroupPolicyExtension and sets the default values.
func NewUnsupportedGroupPolicyExtension()(*UnsupportedGroupPolicyExtension) {
    m := &UnsupportedGroupPolicyExtension{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnsupportedGroupPolicyExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnsupportedGroupPolicyExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnsupportedGroupPolicyExtension(), nil
}
// GetExtensionType gets the extensionType property value. ExtensionType of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) GetExtensionType()(*string) {
    val, err := m.GetBackingStore().Get("extensionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnsupportedGroupPolicyExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["extensionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionType(val)
        }
        return nil
    }
    res["namespaceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNamespaceUrl(val)
        }
        return nil
    }
    res["nodeName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNodeName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["settingScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetNamespaceUrl gets the namespaceUrl property value. Namespace Url of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) GetNamespaceUrl()(*string) {
    val, err := m.GetBackingStore().Get("namespaceUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetNodeName gets the nodeName property value. Node name of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) GetNodeName()(*string) {
    val, err := m.GetBackingStore().Get("nodeName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UnsupportedGroupPolicyExtension) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSettingScope gets the settingScope property value. Scope of the group policy setting.
func (m *UnsupportedGroupPolicyExtension) GetSettingScope()(*GroupPolicySettingScope) {
    val, err := m.GetBackingStore().Get("settingScope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*GroupPolicySettingScope)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UnsupportedGroupPolicyExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    err := m.GetBackingStore().Set("extensionType", value)
    if err != nil {
        panic(err)
    }
}
// SetNamespaceUrl sets the namespaceUrl property value. Namespace Url of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) SetNamespaceUrl(value *string)() {
    err := m.GetBackingStore().Set("namespaceUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetNodeName sets the nodeName property value. Node name of the unsupported extension.
func (m *UnsupportedGroupPolicyExtension) SetNodeName(value *string)() {
    err := m.GetBackingStore().Set("nodeName", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UnsupportedGroupPolicyExtension) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSettingScope sets the settingScope property value. Scope of the group policy setting.
func (m *UnsupportedGroupPolicyExtension) SetSettingScope(value *GroupPolicySettingScope)() {
    err := m.GetBackingStore().Set("settingScope", value)
    if err != nil {
        panic(err)
    }
}
// UnsupportedGroupPolicyExtensionable 
type UnsupportedGroupPolicyExtensionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetExtensionType()(*string)
    GetNamespaceUrl()(*string)
    GetNodeName()(*string)
    GetOdataType()(*string)
    GetSettingScope()(*GroupPolicySettingScope)
    SetExtensionType(value *string)()
    SetNamespaceUrl(value *string)()
    SetNodeName(value *string)()
    SetOdataType(value *string)()
    SetSettingScope(value *GroupPolicySettingScope)()
}
