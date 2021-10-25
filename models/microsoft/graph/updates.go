package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Updates struct {
    Entity
    catalog *Catalog;
    deployments []Deployment;
    updatableAssets []UpdatableAsset;
}
func NewUpdates()(*Updates) {
    m := &Updates{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Updates) GetCatalog()(*Catalog) {
    if m == nil {
        return nil
    } else {
        return m.catalog
    }
}
func (m *Updates) GetDeployments()([]Deployment) {
    if m == nil {
        return nil
    } else {
        return m.deployments
    }
}
func (m *Updates) GetUpdatableAssets()([]UpdatableAsset) {
    if m == nil {
        return nil
    } else {
        return m.updatableAssets
    }
}
func (m *Updates) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["catalog"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCatalog() })
        if err != nil {
            return err
        }
        m.SetCatalog(val.(*Catalog))
        return nil
    }
    res["deployments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeployment() })
        if err != nil {
            return err
        }
        res := make([]Deployment, len(val))
        for i, v := range val {
            res[i] = *(v.(*Deployment))
        }
        m.SetDeployments(res)
        return nil
    }
    res["updatableAssets"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUpdatableAsset() })
        if err != nil {
            return err
        }
        res := make([]UpdatableAsset, len(val))
        for i, v := range val {
            res[i] = *(v.(*UpdatableAsset))
        }
        m.SetUpdatableAssets(res)
        return nil
    }
    return res
}
func (m *Updates) IsNil()(bool) {
    return m == nil
}
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
func (m *Updates) SetCatalog(value *Catalog)() {
    m.catalog = value
}
func (m *Updates) SetDeployments(value []Deployment)() {
    m.deployments = value
}
func (m *Updates) SetUpdatableAssets(value []UpdatableAsset)() {
    m.updatableAssets = value
}
