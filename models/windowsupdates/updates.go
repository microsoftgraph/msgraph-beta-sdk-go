package windowsupdates

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Updates 
type Updates struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Catalog of content that can be approved for deployment by the deployment service. Read-only.
    catalog Catalogable
    // Deployments created using the deployment service. Read-only.
    deployments []Deploymentable
    // Service connections to external resources such as analytics workspaces.
    resourceConnections []ResourceConnectionable
    // Assets registered with the deployment service that can receive updates. Read-only.
    updatableAssets []UpdatableAssetable
}
// NewUpdates instantiates a new updates and sets the default values.
func NewUpdates()(*Updates) {
    m := &Updates{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateUpdatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdates(), nil
}
// GetCatalog gets the catalog property value. Catalog of content that can be approved for deployment by the deployment service. Read-only.
func (m *Updates) GetCatalog()(Catalogable) {
    return m.catalog
}
// GetDeployments gets the deployments property value. Deployments created using the deployment service. Read-only.
func (m *Updates) GetDeployments()([]Deploymentable) {
    return m.deployments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Updates) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["catalog"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateCatalogFromDiscriminatorValue , m.SetCatalog)
    res["deployments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeploymentFromDiscriminatorValue , m.SetDeployments)
    res["resourceConnections"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateResourceConnectionFromDiscriminatorValue , m.SetResourceConnections)
    res["updatableAssets"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUpdatableAssetFromDiscriminatorValue , m.SetUpdatableAssets)
    return res
}
// GetResourceConnections gets the resourceConnections property value. Service connections to external resources such as analytics workspaces.
func (m *Updates) GetResourceConnections()([]ResourceConnectionable) {
    return m.resourceConnections
}
// GetUpdatableAssets gets the updatableAssets property value. Assets registered with the deployment service that can receive updates. Read-only.
func (m *Updates) GetUpdatableAssets()([]UpdatableAssetable) {
    return m.updatableAssets
}
// Serialize serializes information the current object
func (m *Updates) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("catalog", m.GetCatalog())
        if err != nil {
            return err
        }
    }
    if m.GetDeployments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDeployments())
        err = writer.WriteCollectionOfObjectValues("deployments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResourceConnections() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetResourceConnections())
        err = writer.WriteCollectionOfObjectValues("resourceConnections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdatableAssets() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUpdatableAssets())
        err = writer.WriteCollectionOfObjectValues("updatableAssets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCatalog sets the catalog property value. Catalog of content that can be approved for deployment by the deployment service. Read-only.
func (m *Updates) SetCatalog(value Catalogable)() {
    m.catalog = value
}
// SetDeployments sets the deployments property value. Deployments created using the deployment service. Read-only.
func (m *Updates) SetDeployments(value []Deploymentable)() {
    m.deployments = value
}
// SetResourceConnections sets the resourceConnections property value. Service connections to external resources such as analytics workspaces.
func (m *Updates) SetResourceConnections(value []ResourceConnectionable)() {
    m.resourceConnections = value
}
// SetUpdatableAssets sets the updatableAssets property value. Assets registered with the deployment service that can receive updates. Read-only.
func (m *Updates) SetUpdatableAssets(value []UpdatableAssetable)() {
    m.updatableAssets = value
}
