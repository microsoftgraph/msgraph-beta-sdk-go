package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/externalconnectors"
)

// 
type ExternalGroupMember struct {
    Entity
    // The identity source that the member belongs to. Possible values are: azureActiveDirectory, external.
    identitySource *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType;
    // The type of member added to the external group. Possible values are: user or group when the identitySource is azureActiveDirectory and just group when the identitySource is external.
    type_escaped *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalGroupMemberType;
}
// Instantiates a new externalGroupMember and sets the default values.
func NewExternalGroupMember()(*ExternalGroupMember) {
    m := &ExternalGroupMember{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the identitySource property value. The identity source that the member belongs to. Possible values are: azureActiveDirectory, external.
func (m *ExternalGroupMember) GetIdentitySource()(*i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType) {
    if m == nil {
        return nil
    } else {
        return m.identitySource
    }
}
// Gets the type_escaped property value. The type of member added to the external group. Possible values are: user or group when the identitySource is azureActiveDirectory and just group when the identitySource is external.
func (m *ExternalGroupMember) GetType_escaped()(*i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalGroupMemberType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
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
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ParseExternalGroupMemberType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalGroupMemberType)
            m.SetType_escaped(&cast)
        }
        return nil
    }
    return res
}
func (m *ExternalGroupMember) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the identitySource property value. The identity source that the member belongs to. Possible values are: azureActiveDirectory, external.
// Parameters:
//  - value : Value to set for the identitySource property.
func (m *ExternalGroupMember) SetIdentitySource(value *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType)() {
    m.identitySource = value
}
// Sets the type_escaped property value. The type of member added to the external group. Possible values are: user or group when the identitySource is azureActiveDirectory and just group when the identitySource is external.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *ExternalGroupMember) SetType_escaped(value *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ExternalGroupMemberType)() {
    m.type_escaped = value
}
