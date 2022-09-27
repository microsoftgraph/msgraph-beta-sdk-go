package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethodModeDetail 
type AuthenticationMethodModeDetail struct {
    Entity
    // The authenticationMethod property
    authenticationMethod *BaseAuthenticationMethod
    // The displayName property
    displayName *string
}
// NewAuthenticationMethodModeDetail instantiates a new AuthenticationMethodModeDetail and sets the default values.
func NewAuthenticationMethodModeDetail()(*AuthenticationMethodModeDetail) {
    m := &AuthenticationMethodModeDetail{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.authenticationMethodModeDetail";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthenticationMethodModeDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationMethodModeDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationMethodModeDetail(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. The authenticationMethod property
func (m *AuthenticationMethodModeDetail) GetAuthenticationMethod()(*BaseAuthenticationMethod) {
    return m.authenticationMethod
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *AuthenticationMethodModeDetail) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationMethodModeDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationMethod"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseBaseAuthenticationMethod , m.SetAuthenticationMethod)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    return res
}
// Serialize serializes information the current object
func (m *AuthenticationMethodModeDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAuthenticationMethod() != nil {
        cast := (*m.GetAuthenticationMethod()).String()
        err = writer.WriteStringValue("authenticationMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. The authenticationMethod property
func (m *AuthenticationMethodModeDetail) SetAuthenticationMethod(value *BaseAuthenticationMethod)() {
    m.authenticationMethod = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *AuthenticationMethodModeDetail) SetDisplayName(value *string)() {
    m.displayName = value
}
