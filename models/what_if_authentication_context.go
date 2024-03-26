package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WhatIfAuthenticationContext struct {
    ConditionalAccessContext
}
// NewWhatIfAuthenticationContext instantiates a new WhatIfAuthenticationContext and sets the default values.
func NewWhatIfAuthenticationContext()(*WhatIfAuthenticationContext) {
    m := &WhatIfAuthenticationContext{
        ConditionalAccessContext: *NewConditionalAccessContext(),
    }
    odataTypeValue := "#microsoft.graph.whatIfAuthenticationContext"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWhatIfAuthenticationContextFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWhatIfAuthenticationContextFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWhatIfAuthenticationContext(), nil
}
// GetAuthenticationContext gets the authenticationContext property value. The authenticationContext property
// returns a *string when successful
func (m *WhatIfAuthenticationContext) GetAuthenticationContext()(*string) {
    val, err := m.GetBackingStore().Get("authenticationContext")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WhatIfAuthenticationContext) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ConditionalAccessContext.GetFieldDeserializers()
    res["authenticationContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationContext(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WhatIfAuthenticationContext) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ConditionalAccessContext.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("authenticationContext", m.GetAuthenticationContext())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationContext sets the authenticationContext property value. The authenticationContext property
func (m *WhatIfAuthenticationContext) SetAuthenticationContext(value *string)() {
    err := m.GetBackingStore().Set("authenticationContext", value)
    if err != nil {
        panic(err)
    }
}
type WhatIfAuthenticationContextable interface {
    ConditionalAccessContextable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationContext()(*string)
    SetAuthenticationContext(value *string)()
}
