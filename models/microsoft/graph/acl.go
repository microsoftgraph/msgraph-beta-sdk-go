package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/externalconnectors"
)

// 
type Acl struct {
    // The access granted to the identity. Possible values are: grant, deny, unknownFutureValue.
    accessType *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.AccessType;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The source of identity. Possible values are azureActiveDirectory or external.
    identitySource *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType;
    // The type of identity. Possible values are: user, group, everyone, everyoneExceptGuests, externalGroup, unknownFutureValue.
    type_escaped *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.AclType;
    // The unique identifer of the identity. In case of Azure Active Directory identities, value is set to the object identifier of the user, group or tenant for types user, group and everyone (and everyoneExceptGuests) respectively. In case of external groups value is set to the ID of the externalGroup
    value *string;
}
// Instantiates a new acl and sets the default values.
func NewAcl()(*Acl) {
    m := &Acl{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the accessType property value. The access granted to the identity. Possible values are: grant, deny, unknownFutureValue.
func (m *Acl) GetAccessType()(*i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.AccessType) {
    if m == nil {
        return nil
    } else {
        return m.accessType
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Acl) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the identitySource property value. The source of identity. Possible values are azureActiveDirectory or external.
func (m *Acl) GetIdentitySource()(*i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType) {
    if m == nil {
        return nil
    } else {
        return m.identitySource
    }
}
// Gets the type_escaped property value. The type of identity. Possible values are: user, group, everyone, everyoneExceptGuests, externalGroup, unknownFutureValue.
func (m *Acl) GetType_escaped()(*i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.AclType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// Gets the value property value. The unique identifer of the identity. In case of Azure Active Directory identities, value is set to the object identifier of the user, group or tenant for types user, group and everyone (and everyoneExceptGuests) respectively. In case of external groups value is set to the ID of the externalGroup
func (m *Acl) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// The deserialization information for the current model
func (m *Acl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ParseAccessType)
        if err != nil {
            return err
        }
        cast := val.(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.AccessType)
        m.SetAccessType(&cast)
        return nil
    }
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
        val, err := n.GetEnumValue(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.ParseAclType)
        if err != nil {
            return err
        }
        cast := val.(i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.AclType)
        m.SetType_escaped(&cast)
        return nil
    }
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *Acl) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Acl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAccessType() != nil {
        cast := m.GetAccessType().String()
        err := writer.WriteStringValue("accessType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIdentitySource() != nil {
        cast := m.GetIdentitySource().String()
        err := writer.WriteStringValue("identitySource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the accessType property value. The access granted to the identity. Possible values are: grant, deny, unknownFutureValue.
// Parameters:
//  - value : Value to set for the accessType property.
func (m *Acl) SetAccessType(value *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.AccessType)() {
    m.accessType = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *Acl) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the identitySource property value. The source of identity. Possible values are azureActiveDirectory or external.
// Parameters:
//  - value : Value to set for the identitySource property.
func (m *Acl) SetIdentitySource(value *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.IdentitySourceType)() {
    m.identitySource = value
}
// Sets the type_escaped property value. The type of identity. Possible values are: user, group, everyone, everyoneExceptGuests, externalGroup, unknownFutureValue.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *Acl) SetType_escaped(value *i3af76fce9a0d8c03f22ff90ccd64c93d01bbef0102a1c4e80376e26d2e22a367.AclType)() {
    m.type_escaped = value
}
// Sets the value property value. The unique identifer of the identity. In case of Azure Active Directory identities, value is set to the object identifier of the user, group or tenant for types user, group and everyone (and everyoneExceptGuests) respectively. In case of external groups value is set to the ID of the externalGroup
// Parameters:
//  - value : Value to set for the value property.
func (m *Acl) SetValue(value *string)() {
    m.value = value
}
