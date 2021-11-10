package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Updates struct {
    Entity
    // Catalog of content that can be approved for deployment by the deployment service. Read-only.
    catalog *Catalog;
    // Deployments created using the deployment service. Read-only.
    deployments []Deployment;
    // Assets registered with the deployment service that can receive updates. Read-only.
    updatableAssets []UpdatableAsset;
}
// Instantiates a new updates and sets the default values.
func NewUpdates()(*Updates) {
    m := &Updates{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the catalog property value. Catalog of content that can be approved for deployment by the deployment service. Read-only.
func (m *Updates) GetCatalog()(*Catalog) {
    if m == nil {
        return nil
    } else {
        return m.catalog
    }
}
// Gets the deployments property value. Deployments created using the deployment service. Read-only.
func (m *Updates) GetDeployments()([]Deployment) {
    if m == nil {
        return nil
    } else {
        return m.deployments
    }
}
// Gets the updatableAssets property value. Assets registered with the deployment service that can receive updates. Read-only.
func (m *Updates) GetUpdatableAssets()([]UpdatableAsset) {
    if m == nil {
        return nil
    } else {
        return m.updatableAssets
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
    {
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
    {
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
// Sets the catalog property value. Catalog of content that can be approved for deployment by the deployment service. Read-only.
// Parameters:
//  - value : Value to set for the catalog property.
func (m *Updates) SetCatalog(value *Catalog)() {
    m.catalog = value
}
// Sets the deployments property value. Deployments created using the deployment service. Read-only.
// Parameters:
//  - value : Value to set for the deployments property.
func (m *Updates) SetDeployments(value []Deployment)() {
    m.deployments = value
}
// Sets the updatableAssets property value. Assets registered with the deployment service that can receive updates. Read-only.
// Parameters:
//  - value : Value to set for the updatableAssets property.
func (m *Updates) SetUpdatableAssets(value []UpdatableAsset)() {
    m.updatableAssets = value
}
