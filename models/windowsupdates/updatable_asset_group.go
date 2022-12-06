package windowsupdates

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UpdatableAssetGroup 
type UpdatableAssetGroup struct {
    UpdatableAsset
    // Members of the group. Read-only.
    members []UpdatableAssetable
}
// NewUpdatableAssetGroup instantiates a new UpdatableAssetGroup and sets the default values.
func NewUpdatableAssetGroup()(*UpdatableAssetGroup) {
    m := &UpdatableAssetGroup{
        UpdatableAsset: *NewUpdatableAsset(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.updatableAssetGroup";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUpdatableAssetGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatableAssetGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdatableAssetGroup(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatableAssetGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UpdatableAsset.GetFieldDeserializers()
    res["members"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUpdatableAssetFromDiscriminatorValue , m.SetMembers)
    return res
}
// GetMembers gets the members property value. Members of the group. Read-only.
func (m *UpdatableAssetGroup) GetMembers()([]UpdatableAssetable) {
    return m.members
}
// Serialize serializes information the current object
func (m *UpdatableAssetGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UpdatableAsset.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMembers() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetMembers())
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMembers sets the members property value. Members of the group. Read-only.
func (m *UpdatableAssetGroup) SetMembers(value []UpdatableAssetable)() {
    m.members = value
}
