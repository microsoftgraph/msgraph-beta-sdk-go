package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ImpactedDeviceAsset struct {
    ImpactedAsset
}
// NewImpactedDeviceAsset instantiates a new ImpactedDeviceAsset and sets the default values.
func NewImpactedDeviceAsset()(*ImpactedDeviceAsset) {
    m := &ImpactedDeviceAsset{
        ImpactedAsset: *NewImpactedAsset(),
    }
    odataTypeValue := "#microsoft.graph.security.impactedDeviceAsset"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateImpactedDeviceAssetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateImpactedDeviceAssetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImpactedDeviceAsset(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ImpactedDeviceAsset) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ImpactedAsset.GetFieldDeserializers()
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAssetIdentifier)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val.(*DeviceAssetIdentifier))
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. The identifier property
// returns a *DeviceAssetIdentifier when successful
func (m *ImpactedDeviceAsset) GetIdentifier()(*DeviceAssetIdentifier) {
    val, err := m.GetBackingStore().Get("identifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DeviceAssetIdentifier)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ImpactedDeviceAsset) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ImpactedAsset.Serialize(writer)
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
func (m *ImpactedDeviceAsset) SetIdentifier(value *DeviceAssetIdentifier)() {
    err := m.GetBackingStore().Set("identifier", value)
    if err != nil {
        panic(err)
    }
}
type ImpactedDeviceAssetable interface {
    ImpactedAssetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentifier()(*DeviceAssetIdentifier)
    SetIdentifier(value *DeviceAssetIdentifier)()
}
