package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BlockFileResponseAction struct {
    ResponseAction
}
// NewBlockFileResponseAction instantiates a new BlockFileResponseAction and sets the default values.
func NewBlockFileResponseAction()(*BlockFileResponseAction) {
    m := &BlockFileResponseAction{
        ResponseAction: *NewResponseAction(),
    }
    odataTypeValue := "#microsoft.graph.security.blockFileResponseAction"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateBlockFileResponseActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBlockFileResponseActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBlockFileResponseAction(), nil
}
// GetDeviceGroupNames gets the deviceGroupNames property value. Device groups to which the actions set in the custom detection rule are applied. More information
// returns a []string when successful
func (m *BlockFileResponseAction) GetDeviceGroupNames()([]string) {
    val, err := m.GetBackingStore().Get("deviceGroupNames")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BlockFileResponseAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ResponseAction.GetFieldDeserializers()
    res["deviceGroupNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDeviceGroupNames(res)
        }
        return nil
    }
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFileEntityIdentifier)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val.(*FileEntityIdentifier))
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. The identifier property
// returns a *FileEntityIdentifier when successful
func (m *BlockFileResponseAction) GetIdentifier()(*FileEntityIdentifier) {
    val, err := m.GetBackingStore().Get("identifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*FileEntityIdentifier)
    }
    return nil
}
// Serialize serializes information the current object
func (m *BlockFileResponseAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ResponseAction.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeviceGroupNames() != nil {
        err = writer.WriteCollectionOfStringValues("deviceGroupNames", m.GetDeviceGroupNames())
        if err != nil {
            return err
        }
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
// SetDeviceGroupNames sets the deviceGroupNames property value. Device groups to which the actions set in the custom detection rule are applied. More information
func (m *BlockFileResponseAction) SetDeviceGroupNames(value []string)() {
    err := m.GetBackingStore().Set("deviceGroupNames", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentifier sets the identifier property value. The identifier property
func (m *BlockFileResponseAction) SetIdentifier(value *FileEntityIdentifier)() {
    err := m.GetBackingStore().Set("identifier", value)
    if err != nil {
        panic(err)
    }
}
type BlockFileResponseActionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResponseActionable
    GetDeviceGroupNames()([]string)
    GetIdentifier()(*FileEntityIdentifier)
    SetDeviceGroupNames(value []string)()
    SetIdentifier(value *FileEntityIdentifier)()
}
