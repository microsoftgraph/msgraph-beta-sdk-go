package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type IsolateDeviceResponseAction struct {
    ResponseAction
}
// NewIsolateDeviceResponseAction instantiates a new IsolateDeviceResponseAction and sets the default values.
func NewIsolateDeviceResponseAction()(*IsolateDeviceResponseAction) {
    m := &IsolateDeviceResponseAction{
        ResponseAction: *NewResponseAction(),
    }
    odataTypeValue := "#microsoft.graph.security.isolateDeviceResponseAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIsolateDeviceResponseActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIsolateDeviceResponseActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIsolateDeviceResponseAction(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IsolateDeviceResponseAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["isolationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIsolationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsolationType(val.(*IsolationType))
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. The identifier property
// returns a *DeviceIdEntityIdentifier when successful
func (m *IsolateDeviceResponseAction) GetIdentifier()(*DeviceIdEntityIdentifier) {
    val, err := m.GetBackingStore().Get("identifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceIdEntityIdentifier)
    }
    return nil
}
// GetIsolationType gets the isolationType property value. The isolationType property
// returns a *IsolationType when successful
func (m *IsolateDeviceResponseAction) GetIsolationType()(*IsolationType) {
    val, err := m.GetBackingStore().Get("isolationType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IsolationType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IsolateDeviceResponseAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetIsolationType() != nil {
        cast := (*m.GetIsolationType()).String()
        err = writer.WriteStringValue("isolationType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIdentifier sets the identifier property value. The identifier property
func (m *IsolateDeviceResponseAction) SetIdentifier(value *DeviceIdEntityIdentifier)() {
    err := m.GetBackingStore().Set("identifier", value)
    if err != nil {
        panic(err)
    }
}
// SetIsolationType sets the isolationType property value. The isolationType property
func (m *IsolateDeviceResponseAction) SetIsolationType(value *IsolationType)() {
    err := m.GetBackingStore().Set("isolationType", value)
    if err != nil {
        panic(err)
    }
}
type IsolateDeviceResponseActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResponseActionable
    GetIdentifier()(*DeviceIdEntityIdentifier)
    GetIsolationType()(*IsolationType)
    SetIdentifier(value *DeviceIdEntityIdentifier)()
    SetIsolationType(value *IsolationType)()
}
