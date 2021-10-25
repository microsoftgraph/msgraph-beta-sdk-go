package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Catalog struct {
    Entity
    entries []CatalogEntry;
}
func NewCatalog()(*Catalog) {
    m := &Catalog{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Catalog) GetEntries()([]CatalogEntry) {
    if m == nil {
        return nil
    } else {
        return m.entries
    }
}
func (m *Catalog) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["entries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCatalogEntry() })
        if err != nil {
            return err
        }
        res := make([]CatalogEntry, len(val))
        for i, v := range val {
            res[i] = *(v.(*CatalogEntry))
        }
        m.SetEntries(res)
        return nil
    }
    return res
}
func (m *Catalog) IsNil()(bool) {
    return m == nil
}
func (m *Catalog) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEntries()))
        for i, v := range m.GetEntries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("entries", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Catalog) SetEntries(value []CatalogEntry)() {
    m.entries = value
}
