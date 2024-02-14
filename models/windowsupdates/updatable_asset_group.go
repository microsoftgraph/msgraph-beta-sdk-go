package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UpdatableAssetGroup struct {
    UpdatableAsset
}
// NewUpdatableAssetGroup instantiates a new UpdatableAssetGroup and sets the default values.
func NewUpdatableAssetGroup()(*UpdatableAssetGroup) {
    m := &UpdatableAssetGroup{
        UpdatableAsset: *NewUpdatableAsset(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.updatableAssetGroup"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUpdatableAssetGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUpdatableAssetGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdatableAssetGroup(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UpdatableAssetGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UpdatableAsset.GetFieldDeserializers()
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatableAssetable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(UpdatableAssetable)
                }
            }
            m.SetMembers(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. Members of the group. Read-only.
// returns a []UpdatableAssetable when successful
func (m *UpdatableAssetGroup) GetMembers()([]UpdatableAssetable) {
    val, err := m.GetBackingStore().Get("members")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]UpdatableAssetable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *UpdatableAssetGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UpdatableAsset.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMembers sets the members property value. Members of the group. Read-only.
func (m *UpdatableAssetGroup) SetMembers(value []UpdatableAssetable)() {
    err := m.GetBackingStore().Set("members", value)
    if err != nil {
        panic(err)
    }
}
type UpdatableAssetGroupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UpdatableAssetable
    GetMembers()([]UpdatableAssetable)
    SetMembers(value []UpdatableAssetable)()
}
