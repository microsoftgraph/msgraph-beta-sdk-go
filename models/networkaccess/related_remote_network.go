package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RelatedRemoteNetwork struct {
    RelatedResource
}
// NewRelatedRemoteNetwork instantiates a new RelatedRemoteNetwork and sets the default values.
func NewRelatedRemoteNetwork()(*RelatedRemoteNetwork) {
    m := &RelatedRemoteNetwork{
        RelatedResource: *NewRelatedResource(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.relatedRemoteNetwork"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRelatedRemoteNetworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRelatedRemoteNetworkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelatedRemoteNetwork(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RelatedRemoteNetwork) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RelatedResource.GetFieldDeserializers()
    res["remoteNetworkId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteNetworkId(val)
        }
        return nil
    }
    return res
}
// GetRemoteNetworkId gets the remoteNetworkId property value. The remoteNetworkId property
// returns a *string when successful
func (m *RelatedRemoteNetwork) GetRemoteNetworkId()(*string) {
    val, err := m.GetBackingStore().Get("remoteNetworkId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RelatedRemoteNetwork) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RelatedResource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("remoteNetworkId", m.GetRemoteNetworkId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRemoteNetworkId sets the remoteNetworkId property value. The remoteNetworkId property
func (m *RelatedRemoteNetwork) SetRemoteNetworkId(value *string)() {
    err := m.GetBackingStore().Set("remoteNetworkId", value)
    if err != nil {
        panic(err)
    }
}
type RelatedRemoteNetworkable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RelatedResourceable
    GetRemoteNetworkId()(*string)
    SetRemoteNetworkId(value *string)()
}
