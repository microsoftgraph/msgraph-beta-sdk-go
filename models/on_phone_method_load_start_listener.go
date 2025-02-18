package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OnPhoneMethodLoadStartListener struct {
    AuthenticationEventListener
}
// NewOnPhoneMethodLoadStartListener instantiates a new OnPhoneMethodLoadStartListener and sets the default values.
func NewOnPhoneMethodLoadStartListener()(*OnPhoneMethodLoadStartListener) {
    m := &OnPhoneMethodLoadStartListener{
        AuthenticationEventListener: *NewAuthenticationEventListener(),
    }
    odataTypeValue := "#microsoft.graph.onPhoneMethodLoadStartListener"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnPhoneMethodLoadStartListenerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOnPhoneMethodLoadStartListenerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPhoneMethodLoadStartListener(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OnPhoneMethodLoadStartListener) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationEventListener.GetFieldDeserializers()
    res["handler"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnPhoneMethodLoadStartHandlerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHandler(val.(OnPhoneMethodLoadStartHandlerable))
        }
        return nil
    }
    return res
}
// GetHandler gets the handler property value. The handler property
// returns a OnPhoneMethodLoadStartHandlerable when successful
func (m *OnPhoneMethodLoadStartListener) GetHandler()(OnPhoneMethodLoadStartHandlerable) {
    val, err := m.GetBackingStore().Get("handler")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OnPhoneMethodLoadStartHandlerable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnPhoneMethodLoadStartListener) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetHandler sets the handler property value. The handler property
func (m *OnPhoneMethodLoadStartListener) SetHandler(value OnPhoneMethodLoadStartHandlerable)() {
    err := m.GetBackingStore().Set("handler", value)
    if err != nil {
        panic(err)
    }
}
type OnPhoneMethodLoadStartListenerable interface {
    AuthenticationEventListenerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHandler()(OnPhoneMethodLoadStartHandlerable)
    SetHandler(value OnPhoneMethodLoadStartHandlerable)()
}
