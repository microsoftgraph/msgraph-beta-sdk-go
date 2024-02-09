package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ImpactedUserAsset struct {
    ImpactedAsset
}
// NewImpactedUserAsset instantiates a new ImpactedUserAsset and sets the default values.
func NewImpactedUserAsset()(*ImpactedUserAsset) {
    m := &ImpactedUserAsset{
        ImpactedAsset: *NewImpactedAsset(),
    }
    odataTypeValue := "#microsoft.graph.security.impactedUserAsset"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateImpactedUserAssetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateImpactedUserAssetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImpactedUserAsset(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ImpactedUserAsset) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ImpactedAsset.GetFieldDeserializers()
    res["identifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserAssetIdentifier)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentifier(val.(*UserAssetIdentifier))
        }
        return nil
    }
    return res
}
// GetIdentifier gets the identifier property value. The identifier property
// returns a *UserAssetIdentifier when successful
func (m *ImpactedUserAsset) GetIdentifier()(*UserAssetIdentifier) {
    val, err := m.GetBackingStore().Get("identifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*UserAssetIdentifier)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ImpactedUserAsset) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ImpactedUserAsset) SetIdentifier(value *UserAssetIdentifier)() {
    err := m.GetBackingStore().Set("identifier", value)
    if err != nil {
        panic(err)
    }
}
type ImpactedUserAssetable interface {
    ImpactedAssetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIdentifier()(*UserAssetIdentifier)
    SetIdentifier(value *UserAssetIdentifier)()
}
