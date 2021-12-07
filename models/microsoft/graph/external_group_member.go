package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/externalconnectors"
)

// ExternalGroupMember 
type ExternalGroupMember struct {
    Entity
    // The identity source that the member belongs to. Possible values are: azureActiveDirectory, external.
    identitySource *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType;
    // The type of member added to the external group. Possible values are: user or group when the identitySource is azureActiveDirectory and just group when the identitySource is external.
    type_escaped *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalGroupMemberType;
}
// NewExternalGroupMember instantiates a new externalGroupMember and sets the default values.
func NewExternalGroupMember()(*ExternalGroupMember) {
    m := &ExternalGroupMember{
        Entity: *NewEntity(),
    }
    return m
}
// GetIdentitySource gets the identitySource property value. The identity source that the member belongs to. Possible values are: azureActiveDirectory, external.
func (m *ExternalGroupMember) GetIdentitySource()(*i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType) {
    if m == nil {
        return nil
    } else {
        return m.identitySource
    }
}
// GetType gets the type property value. The type of member added to the external group. Possible values are: user or group when the identitySource is azureActiveDirectory and just group when the identitySource is external.
func (m *ExternalGroupMember) GetType()(*i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalGroupMemberType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalGroupMember) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["identitySource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ParseIdentitySourceType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType)
            m.SetIdentitySource(&cast)
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ParseExternalGroupMemberType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalGroupMemberType)
            m.SetType(&cast)
        }
        return nil
    }
    return res
}
func (m *ExternalGroupMember) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetType() != nil {
        cast := m.GetType().String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentitySource sets the identitySource property value. The identity source that the member belongs to. Possible values are: azureActiveDirectory, external.
func (m *ExternalGroupMember) SetIdentitySource(value *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType)() {
    if m != nil {
        m.identitySource = value
    }
}
// SetType sets the type property value. The type of member added to the external group. Possible values are: user or group when the identitySource is azureActiveDirectory and just group when the identitySource is external.
func (m *ExternalGroupMember) SetType(value *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalGroupMemberType)() {
    if m != nil {
        m.type_escaped = value
    }
}
