package windowsupdates

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// Updates provides operations to manage the admin singleton.
type Updates struct {
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Entity
    // Catalog of content that can be approved for deployment by the deployment service. Read-only.
    catalog Catalogable;
    // Deployments created using the deployment service. Read-only.
    deployments []Deploymentable;
    // Assets registered with the deployment service that can receive updates. Read-only.
    updatableAssets []UpdatableAssetable;
}
// NewUpdates instantiates a new updates and sets the default values.
func NewUpdates()(*Updates) {
    m := &Updates{
        Entity: *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewEntity(),
    }
    return m
}
// CreateUpdatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatesFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
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
func (m *Updates) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["catalog"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalog(val.(Catalogable))
        }
        return nil
    }
    res["deployments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["updatableAssets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
// GetUpdatableAssets gets the updatableAssets property value. Assets registered with the deployment service that can receive updates. Read-only.
func (m *Updates) GetUpdatableAssets()([]UpdatableAssetable) {
    if m == nil {
        return nil
    } else {
        return m.updatableAssets
    }
}
func (m *Updates) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Updates) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeployments()))
        for i, v := range m.GetDeployments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deployments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdatableAssets() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUpdatableAssets()))
        for i, v := range m.GetUpdatableAssets() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
// SetUpdatableAssets sets the updatableAssets property value. Assets registered with the deployment service that can receive updates. Read-only.
func (m *Updates) SetUpdatableAssets(value []UpdatableAssetable)() {
    if m != nil {
        m.updatableAssets = value
    }
}
