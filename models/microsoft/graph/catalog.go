package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Catalog struct {
    Entity
    // Lists the content that you can approve for deployment. Read-only.
    entries []CatalogEntry;
}
// Instantiates a new catalog and sets the default values.
func NewCatalog()(*Catalog) {
    m := &Catalog{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the entries property value. Lists the content that you can approve for deployment. Read-only.
func (m *Catalog) GetEntries()([]CatalogEntry) {
    if m == nil {
        return nil
    } else {
        return m.entries
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the entries property value. Lists the content that you can approve for deployment. Read-only.
// Parameters:
//  - value : Value to set for the entries property.
func (m *Catalog) SetEntries(value []CatalogEntry)() {
    m.entries = value
}
