package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/externalconnectors"
)

type ExternalGroupMember struct {
    Entity
    identitySource *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType;
    type_escaped *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalGroupMemberType;
}
func NewExternalGroupMember()(*ExternalGroupMember) {
    m := &ExternalGroupMember{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ExternalGroupMember) GetIdentitySource()(*i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType) {
    if m == nil {
        return nil
    } else {
        return m.identitySource
    }
}
func (m *ExternalGroupMember) GetType_escaped()(*i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalGroupMemberType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
func (m *ExternalGroupMember) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["identitySource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ParseIdentitySourceType)
        if err != nil {
            return err
        }
        cast := val.(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType)
        m.SetIdentitySource(&cast)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ParseExternalGroupMemberType)
        if err != nil {
            return err
        }
        cast := val.(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalGroupMemberType)
        m.SetType_escaped(&cast)
        return nil
    }
    return res
}
func (m *ExternalGroupMember) IsNil()(bool) {
    return m == nil
}
func (m *ExternalGroupMember) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIdentitySource() != nil {
        cast := m.GetIdentitySource().String()
        err = writer.WriteStringValue("identitySource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err = writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ExternalGroupMember) SetIdentitySource(value *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType)() {
    m.identitySource = value
}
func (m *ExternalGroupMember) SetType_escaped(value *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalGroupMemberType)() {
    m.type_escaped = value
}
