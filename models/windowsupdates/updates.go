package windowsupdates

import (
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
    // The resourceConnections property
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
    if m == nil {
        return nil
    } else {
        return m.catalog
    }
}
// GetDeployments gets the deployments property value. Deployments created using the deployment service. Read-only.
func (m *Updates) GetDeployments()([]Deploymentable) {
    if m == nil {
        return nil
    } else {
        return m.deployments
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Updates) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["catalog"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalog(val.(Catalogable))
        }
        return nil
    }
    res["deployments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeploymentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Deploymentable, len(val))
            for i, v := range val {
                res[i] = v.(Deploymentable)
            }
            m.SetDeployments(res)
        }
        return nil
    }
    res["resourceConnections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateResourceConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ResourceConnectionable, len(val))
            for i, v := range val {
                res[i] = v.(ResourceConnectionable)
            }
            m.SetResourceConnections(res)
        }
        return nil
    }
    res["updatableAssets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatableAssetable, len(val))
            for i, v := range val {
                res[i] = v.(UpdatableAssetable)
            }
            m.SetUpdatableAssets(res)
        }
        return nil
    }
    return res
}
// GetResourceConnections gets the resourceConnections property value. The resourceConnections property
func (m *Updates) GetResourceConnections()([]ResourceConnectionable) {
    if m == nil {
        return nil
    } else {
        return m.resourceConnections
    }
}
// GetUpdatableAssets gets the updatableAssets property value. Assets registered with the deployment service that can receive updates. Read-only.
func (m *Updates) GetUpdatableAssets()([]UpdatableAssetable) {
    if m == nil {
        return nil
    } else {
        return m.updatableAssets
    }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeployments()))
        for i, v := range m.GetDeployments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deployments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetResourceConnections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResourceConnections()))
        for i, v := range m.GetResourceConnections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resourceConnections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdatableAssets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUpdatableAssets()))
        for i, v := range m.GetUpdatableAssets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("updatableAssets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCatalog sets the catalog property value. Catalog of content that can be approved for deployment by the deployment service. Read-only.
func (m *Updates) SetCatalog(value Catalogable)() {
    if m != nil {
        m.catalog = value
    }
}
// SetDeployments sets the deployments property value. Deployments created using the deployment service. Read-only.
func (m *Updates) SetDeployments(value []Deploymentable)() {
    if m != nil {
        m.deployments = value
    }
}
// SetResourceConnections sets the resourceConnections property value. The resourceConnections property
func (m *Updates) SetResourceConnections(value []ResourceConnectionable)() {
    if m != nil {
        m.resourceConnections = value
    }
}
// SetUpdatableAssets sets the updatableAssets property value. Assets registered with the deployment service that can receive updates. Read-only.
func (m *Updates) SetUpdatableAssets(value []UpdatableAssetable)() {
    if m != nil {
        m.updatableAssets = value
    }
}
