package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ForwardingOptions 
type ForwardingOptions struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The OdataType property
    OdataType *string
}
// NewForwardingOptions instantiates a new forwardingOptions and sets the default values.
func NewForwardingOptions()(*ForwardingOptions) {
    m := &ForwardingOptions{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateForwardingOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateForwardingOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewForwardingOptions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ForwardingOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["skipDnsLookupState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkipDnsLookupState(val.(*Status))
        }
        return nil
    }
    return res
}
// GetSkipDnsLookupState gets the skipDnsLookupState property value. The skipDnsLookupState property
func (m *ForwardingOptions) GetSkipDnsLookupState()(*Status) {
    val, err := m.GetBackingStore().Get("skipDnsLookupState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Status)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ForwardingOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSkipDnsLookupState() != nil {
        cast := (*m.GetSkipDnsLookupState()).String()
        err = writer.WriteStringValue("skipDnsLookupState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSkipDnsLookupState sets the skipDnsLookupState property value. The skipDnsLookupState property
func (m *ForwardingOptions) SetSkipDnsLookupState(value *Status)() {
    err := m.GetBackingStore().Set("skipDnsLookupState", value)
    if err != nil {
        panic(err)
    }
}
// ForwardingOptionsable 
type ForwardingOptionsable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSkipDnsLookupState()(*Status)
    SetSkipDnsLookupState(value *Status)()
}
