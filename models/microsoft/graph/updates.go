package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Updates 
type Updates struct {
    Entity
    // Catalog of content that can be approved for deployment by the deployment service. Read-only.
    catalog *Catalog;
    // Deployments created using the deployment service. Read-only.
    deployments []Deployment;
    // Assets registered with the deployment service that can receive updates. Read-only.
    updatableAssets []UpdatableAsset;
}
// NewUpdates instantiates a new updates and sets the default values.
func NewUpdates()(*Updates) {
    m := &Updates{
        Entity: *NewEntity(),
    }
    return m
}
// GetCatalog gets the catalog property value. Catalog of content that can be approved for deployment by the deployment service. Read-only.
func (m *Updates) GetCatalog()(*Catalog) {
    if m == nil {
        return nil
    } else {
        return m.catalog
    }
}
// GetDeployments gets the deployments property value. Deployments created using the deployment service. Read-only.
func (m *Updates) GetDeployments()([]Deployment) {
    if m == nil {
        return nil
    } else {
        return m.deployments
    }
}
// GetUpdatableAssets gets the updatableAssets property value. Assets registered with the deployment service that can receive updates. Read-only.
func (m *Updates) GetUpdatableAssets()([]UpdatableAsset) {
    if m == nil {
        return nil
    } else {
        return m.updatableAssets
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Updates) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["catalog"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCatalog() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalog(val.(*Catalog))
        }
        return nil
    }
    res["deployments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeployment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Deployment, len(val))
            for i, v := range val {
                res[i] = *(v.(*Deployment))
            }
            m.SetDeployments(res)
        }
        return nil
    }
    res["updatableAssets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdatableAsset() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UpdatableAsset, len(val))
            for i, v := range val {
                res[i] = *(v.(*UpdatableAsset))
            }
            m.SetUpdatableAssets(res)
        }
        return nil
    }
    return res
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
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deployments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdatableAssets() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUpdatableAssets()))
        for i, v := range m.GetUpdatableAssets() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("updatableAssets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCatalog sets the catalog property value. Catalog of content that can be approved for deployment by the deployment service. Read-only.
func (m *Updates) SetCatalog(value *Catalog)() {
    if m != nil {
        m.catalog = value
    }
}
// SetDeployments sets the deployments property value. Deployments created using the deployment service. Read-only.
func (m *Updates) SetDeployments(value []Deployment)() {
    if m != nil {
        m.deployments = value
    }
}
// SetUpdatableAssets sets the updatableAssets property value. Assets registered with the deployment service that can receive updates. Read-only.
func (m *Updates) SetUpdatableAssets(value []UpdatableAsset)() {
    if m != nil {
        m.updatableAssets = value
    }
}
