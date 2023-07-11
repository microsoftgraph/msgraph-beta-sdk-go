package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnTokenIssuanceStartListener 
type OnTokenIssuanceStartListener struct {
    AuthenticationEventListener
}
// NewOnTokenIssuanceStartListener instantiates a new onTokenIssuanceStartListener and sets the default values.
func NewOnTokenIssuanceStartListener()(*OnTokenIssuanceStartListener) {
    m := &OnTokenIssuanceStartListener{
        AuthenticationEventListener: *NewAuthenticationEventListener(),
    }
    odataTypeValue := "#microsoft.graph.onTokenIssuanceStartListener"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnTokenIssuanceStartListenerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnTokenIssuanceStartListenerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnTokenIssuanceStartListener(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnTokenIssuanceStartListener) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationEventListener.GetFieldDeserializers()
    res["handler"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnTokenIssuanceStartHandlerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHandler(val.(OnTokenIssuanceStartHandlerable))
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
// GetHandler gets the handler property value. The handler to invoke when conditions are met for this onTokenIssuanceStartListener.
func (m *OnTokenIssuanceStartListener) GetHandler()(OnTokenIssuanceStartHandlerable) {
    val, err := m.GetBackingStore().Get("handler")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OnTokenIssuanceStartHandlerable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnTokenIssuanceStartListener) GetOdataType()(*string) {
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
func (m *OnTokenIssuanceStartListener) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationEventListener.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("handler", m.GetHandler())
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
// SetHandler sets the handler property value. The handler to invoke when conditions are met for this onTokenIssuanceStartListener.
func (m *OnTokenIssuanceStartListener) SetHandler(value OnTokenIssuanceStartHandlerable)() {
    err := m.GetBackingStore().Set("handler", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnTokenIssuanceStartListener) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// OnTokenIssuanceStartListenerable 
type OnTokenIssuanceStartListenerable interface {
    AuthenticationEventListenerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHandler()(OnTokenIssuanceStartHandlerable)
    GetOdataType()(*string)
    SetHandler(value OnTokenIssuanceStartHandlerable)()
    SetOdataType(value *string)()
}
