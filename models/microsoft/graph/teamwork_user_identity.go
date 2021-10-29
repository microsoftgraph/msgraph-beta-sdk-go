package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TeamworkUserIdentity struct {
    Identity
    // Type of user. Possible values are: aadUser, onPremiseAadUser, anonymousGuest, federatedUser, personalMicrosoftAccountUser, skypeUser, phoneUser, and unknownFutureValue.
    userIdentityType *TeamworkUserIdentityType;
}
// Instantiates a new teamworkUserIdentity and sets the default values.
func NewTeamworkUserIdentity()(*TeamworkUserIdentity) {
    m := &TeamworkUserIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
// Gets the userIdentityType property value. Type of user. Possible values are: aadUser, onPremiseAadUser, anonymousGuest, federatedUser, personalMicrosoftAccountUser, skypeUser, phoneUser, and unknownFutureValue.
func (m *TeamworkUserIdentity) GetUserIdentityType()(*TeamworkUserIdentityType) {
    if m == nil {
        return nil
    } else {
        return m.userIdentityType
    }
}
// The deserialization information for the current model
func (m *TeamworkUserIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["userIdentityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkUserIdentityType)
        if err != nil {
            return err
        }
        cast := val.(TeamworkUserIdentityType)
        m.SetUserIdentityType(&cast)
        return nil
    }
    return res
}
func (m *TeamworkUserIdentity) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TeamworkUserIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUserIdentityType() != nil {
        cast := m.GetUserIdentityType().String()
        err = writer.WriteStringValue("userIdentityType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the userIdentityType property value. Type of user. Possible values are: aadUser, onPremiseAadUser, anonymousGuest, federatedUser, personalMicrosoftAccountUser, skypeUser, phoneUser, and unknownFutureValue.
// Parameters:
//  - value : Value to set for the userIdentityType property.
func (m *TeamworkUserIdentity) SetUserIdentityType(value *TeamworkUserIdentityType)() {
    m.userIdentityType = value
}
