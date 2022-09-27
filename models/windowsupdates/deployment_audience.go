package windowsupdates

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DeploymentAudience 
type DeploymentAudience struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Specifies the assets to exclude from the audience.
    exclusions []UpdatableAssetable
    // Specifies the assets to include in the audience.
    members []UpdatableAssetable
}
// NewDeploymentAudience instantiates a new deploymentAudience and sets the default values.
func NewDeploymentAudience()(*DeploymentAudience) {
    m := &DeploymentAudience{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.deploymentAudience";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeploymentAudienceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeploymentAudienceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeploymentAudience(), nil
}
// GetExclusions gets the exclusions property value. Specifies the assets to exclude from the audience.
func (m *DeploymentAudience) GetExclusions()([]UpdatableAssetable) {
    return m.exclusions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeploymentAudience) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["exclusions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUpdatableAssetFromDiscriminatorValue , m.SetExclusions)
    res["members"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUpdatableAssetFromDiscriminatorValue , m.SetMembers)
    return res
}
// GetMembers gets the members property value. Specifies the assets to include in the audience.
func (m *DeploymentAudience) GetMembers()([]UpdatableAssetable) {
    return m.members
}
// Serialize serializes information the current object
func (m *DeploymentAudience) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetExclusions() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetExclusions())
        err = writer.WriteCollectionOfObjectValues("exclusions", cast)
        if err != nil {
            return err
        }
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
// SetExclusions sets the exclusions property value. Specifies the assets to exclude from the audience.
func (m *DeploymentAudience) SetExclusions(value []UpdatableAssetable)() {
    m.exclusions = value
}
// SetMembers sets the members property value. Specifies the assets to include in the audience.
func (m *DeploymentAudience) SetMembers(value []UpdatableAssetable)() {
    m.members = value
}
