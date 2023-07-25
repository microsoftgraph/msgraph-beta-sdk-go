package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HardDeleteResponseAction 
type HardDeleteResponseAction struct {
    ResponseAction
}
// NewHardDeleteResponseAction instantiates a new hardDeleteResponseAction and sets the default values.
func NewHardDeleteResponseAction()(*HardDeleteResponseAction) {
    m := &HardDeleteResponseAction{
        ResponseAction: *NewResponseAction(),
    }
    odataTypeValue := "#microsoft.graph.security.hardDeleteResponseAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateHardDeleteResponseActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateHardDeleteResponseActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewHardDeleteResponseAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *HardDeleteResponseAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *HardDeleteResponseAction) GetIdentifier()(*EmailEntityIdentifier) {
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
func (m *HardDeleteResponseAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *HardDeleteResponseAction) SetIdentifier(value *EmailEntityIdentifier)() {
    err := m.GetBackingStore().Set("identifier", value)
    if err != nil {
        panic(err)
    }
}
// HardDeleteResponseActionable 
type HardDeleteResponseActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResponseActionable
    GetIdentifier()(*EmailEntityIdentifier)
    SetIdentifier(value *EmailEntityIdentifier)()
}
