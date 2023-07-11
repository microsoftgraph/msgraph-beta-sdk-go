package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp 
type OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp struct {
    OnInteractiveAuthFlowStartHandler
}
// NewOnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp instantiates a new onInteractiveAuthFlowStartExternalUsersSelfServiceSignUp and sets the default values.
func NewOnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp()(*OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp) {
    m := &OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp{
        OnInteractiveAuthFlowStartHandler: *NewOnInteractiveAuthFlowStartHandler(),
    }
    odataTypeValue := "#microsoft.graph.onInteractiveAuthFlowStartExternalUsersSelfServiceSignUp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnInteractiveAuthFlowStartExternalUsersSelfServiceSignUpFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnInteractiveAuthFlowStartExternalUsersSelfServiceSignUpFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnInteractiveAuthFlowStartHandler.GetFieldDeserializers()
    res["isSignUpAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSignUpAllowed(val)
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
    return res
}
// GetIsSignUpAllowed gets the isSignUpAllowed property value. Optional. Specifes whether the authentication flow includes an option to sign up (create account) as well as sign in. Default value is false meaning only sign in is enabled.
func (m *OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp) GetIsSignUpAllowed()(*bool) {
    val, err := m.GetBackingStore().Get("isSignUpAllowed")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnInteractiveAuthFlowStartHandler.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isSignUpAllowed", m.GetIsSignUpAllowed())
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
    return nil
}
// SetIsSignUpAllowed sets the isSignUpAllowed property value. Optional. Specifes whether the authentication flow includes an option to sign up (create account) as well as sign in. Default value is false meaning only sign in is enabled.
func (m *OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp) SetIsSignUpAllowed(value *bool)() {
    err := m.GetBackingStore().Set("isSignUpAllowed", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUpable 
type OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUpable interface {
    OnInteractiveAuthFlowStartHandlerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsSignUpAllowed()(*bool)
    GetOdataType()(*string)
    SetIsSignUpAllowed(value *bool)()
    SetOdataType(value *string)()
}
