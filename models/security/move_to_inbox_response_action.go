package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MoveToInboxResponseAction struct {
    ResponseAction
}
// NewMoveToInboxResponseAction instantiates a new MoveToInboxResponseAction and sets the default values.
func NewMoveToInboxResponseAction()(*MoveToInboxResponseAction) {
    m := &MoveToInboxResponseAction{
        ResponseAction: *NewResponseAction(),
    }
    odataTypeValue := "#microsoft.graph.security.moveToInboxResponseAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMoveToInboxResponseActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMoveToInboxResponseActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMoveToInboxResponseAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MoveToInboxResponseAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ResponseAction.GetFieldDeserializers()
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailEntityIdentifier)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val.(*EmailEntityIdentifier))
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. The identifier property
// returns a *EmailEntityIdentifier when successful
func (m *MoveToInboxResponseAction) GetIdentifier()(*EmailEntityIdentifier) {
    val, err := m.GetBackingStore().Get("identifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EmailEntityIdentifier)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MoveToInboxResponseAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *MoveToInboxResponseAction) SetIdentifier(value *EmailEntityIdentifier)() {
    err := m.GetBackingStore().Set("identifier", value)
    if err != nil {
        panic(err)
    }
}
type MoveToInboxResponseActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResponseActionable
    GetIdentifier()(*EmailEntityIdentifier)
    SetIdentifier(value *EmailEntityIdentifier)()
}
