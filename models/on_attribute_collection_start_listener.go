package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnAttributeCollectionStartListener 
type OnAttributeCollectionStartListener struct {
    AuthenticationEventListener
}
// NewOnAttributeCollectionStartListener instantiates a new onAttributeCollectionStartListener and sets the default values.
func NewOnAttributeCollectionStartListener()(*OnAttributeCollectionStartListener) {
    m := &OnAttributeCollectionStartListener{
        AuthenticationEventListener: *NewAuthenticationEventListener(),
    }
    odataTypeValue := "#microsoft.graph.onAttributeCollectionStartListener"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnAttributeCollectionStartListenerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnAttributeCollectionStartListenerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnAttributeCollectionStartListener(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnAttributeCollectionStartListener) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationEventListener.GetFieldDeserializers()
    res["handler"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnAttributeCollectionStartHandlerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHandler(val.(OnAttributeCollectionStartHandlerable))
        }
        return nil
    }
    return res
}
// GetHandler gets the handler property value. The handler property
func (m *OnAttributeCollectionStartListener) GetHandler()(OnAttributeCollectionStartHandlerable) {
    val, err := m.GetBackingStore().Get("handler")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OnAttributeCollectionStartHandlerable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnAttributeCollectionStartListener) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OnAttributeCollectionStartListener) SetHandler(value OnAttributeCollectionStartHandlerable)() {
    err := m.GetBackingStore().Set("handler", value)
    if err != nil {
        panic(err)
    }
}
// OnAttributeCollectionStartListenerable 
type OnAttributeCollectionStartListenerable interface {
    AuthenticationEventListenerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHandler()(OnAttributeCollectionStartHandlerable)
    SetHandler(value OnAttributeCollectionStartHandlerable)()
}
