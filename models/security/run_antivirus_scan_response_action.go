package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RunAntivirusScanResponseAction struct {
    ResponseAction
}
// NewRunAntivirusScanResponseAction instantiates a new RunAntivirusScanResponseAction and sets the default values.
func NewRunAntivirusScanResponseAction()(*RunAntivirusScanResponseAction) {
    m := &RunAntivirusScanResponseAction{
        ResponseAction: *NewResponseAction(),
    }
    odataTypeValue := "#microsoft.graph.security.runAntivirusScanResponseAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRunAntivirusScanResponseActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRunAntivirusScanResponseActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRunAntivirusScanResponseAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RunAntivirusScanResponseAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ResponseAction.GetFieldDeserializers()
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceIdEntityIdentifier)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val.(*DeviceIdEntityIdentifier))
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. The identifier property
// returns a *DeviceIdEntityIdentifier when successful
func (m *RunAntivirusScanResponseAction) GetIdentifier()(*DeviceIdEntityIdentifier) {
    val, err := m.GetBackingStore().Get("identifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceIdEntityIdentifier)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RunAntivirusScanResponseAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *RunAntivirusScanResponseAction) SetIdentifier(value *DeviceIdEntityIdentifier)() {
    err := m.GetBackingStore().Set("identifier", value)
    if err != nil {
        panic(err)
    }
}
type RunAntivirusScanResponseActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResponseActionable
    GetIdentifier()(*DeviceIdEntityIdentifier)
    SetIdentifier(value *DeviceIdEntityIdentifier)()
}
