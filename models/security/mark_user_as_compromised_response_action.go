package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MarkUserAsCompromisedResponseAction 
type MarkUserAsCompromisedResponseAction struct {
    ResponseAction
}
// NewMarkUserAsCompromisedResponseAction instantiates a new markUserAsCompromisedResponseAction and sets the default values.
func NewMarkUserAsCompromisedResponseAction()(*MarkUserAsCompromisedResponseAction) {
    m := &MarkUserAsCompromisedResponseAction{
        ResponseAction: *NewResponseAction(),
    }
    odataTypeValue := "#microsoft.graph.security.markUserAsCompromisedResponseAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMarkUserAsCompromisedResponseActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMarkUserAsCompromisedResponseActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMarkUserAsCompromisedResponseAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MarkUserAsCompromisedResponseAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ResponseAction.GetFieldDeserializers()
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMarkUserAsCompromisedEntityIdentifier)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val.(*MarkUserAsCompromisedEntityIdentifier))
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. The identifier property
func (m *MarkUserAsCompromisedResponseAction) GetIdentifier()(*MarkUserAsCompromisedEntityIdentifier) {
    val, err := m.GetBackingStore().Get("identifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MarkUserAsCompromisedEntityIdentifier)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MarkUserAsCompromisedResponseAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ResponseAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIdentifier() != nil {
        cast := (*m.GetIdentifier()).String()
        err = writer.WriteStringValue("identifier", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentifier sets the identifier property value. The identifier property
func (m *MarkUserAsCompromisedResponseAction) SetIdentifier(value *MarkUserAsCompromisedEntityIdentifier)() {
    err := m.GetBackingStore().Set("identifier", value)
    if err != nil {
        panic(err)
    }
}
// MarkUserAsCompromisedResponseActionable 
type MarkUserAsCompromisedResponseActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResponseActionable
    GetIdentifier()(*MarkUserAsCompromisedEntityIdentifier)
    SetIdentifier(value *MarkUserAsCompromisedEntityIdentifier)()
}
