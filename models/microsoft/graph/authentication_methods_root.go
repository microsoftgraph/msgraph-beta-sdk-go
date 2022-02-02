package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuthenticationMethodsRoot 
type AuthenticationMethodsRoot struct {
    Entity
    // Represents the state of a user's authentication methods, including which methods are registered and which features the user is registered and capable of (such as multi-factor authentication, self-service password reset, and passwordless authentication).
    userRegistrationDetails []UserRegistrationDetails;
}
// NewAuthenticationMethodsRoot instantiates a new authenticationMethodsRoot and sets the default values.
func NewAuthenticationMethodsRoot()(*AuthenticationMethodsRoot) {
    m := &AuthenticationMethodsRoot{
        Entity: *NewEntity(),
    }
    return m
}
// GetUserRegistrationDetails gets the userRegistrationDetails property value. Represents the state of a user's authentication methods, including which methods are registered and which features the user is registered and capable of (such as multi-factor authentication, self-service password reset, and passwordless authentication).
func (m *AuthenticationMethodsRoot) GetUserRegistrationDetails()([]UserRegistrationDetails) {
    if m == nil {
        return nil
    } else {
        return m.userRegistrationDetails
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationMethodsRoot) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["userRegistrationDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserRegistrationDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserRegistrationDetails, len(val))
            for i, v := range val {
                res[i] = *(v.(*UserRegistrationDetails))
            }
            m.SetUserRegistrationDetails(res)
        }
        return nil
    }
    return res
}
func (m *AuthenticationMethodsRoot) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AuthenticationMethodsRoot) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUserRegistrationDetails() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserRegistrationDetails()))
        for i, v := range m.GetUserRegistrationDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userRegistrationDetails", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserRegistrationDetails sets the userRegistrationDetails property value. Represents the state of a user's authentication methods, including which methods are registered and which features the user is registered and capable of (such as multi-factor authentication, self-service password reset, and passwordless authentication).
func (m *AuthenticationMethodsRoot) SetUserRegistrationDetails(value []UserRegistrationDetails)() {
    if m != nil {
        m.userRegistrationDetails = value
    }
}
