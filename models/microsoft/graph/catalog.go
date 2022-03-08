package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Catalog provides operations to manage the admin singleton.
type Catalog struct {
    Entity
    // Lists the content that you can approve for deployment. Read-only.
    entries []CatalogEntryable;
}
// NewCatalog instantiates a new catalog and sets the default values.
func NewCatalog()(*Catalog) {
    m := &Catalog{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCatalogFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCatalogFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewCatalog(), nil
}
// GetEntries gets the entries property value. Lists the content that you can approve for deployment. Read-only.
func (m *Catalog) GetEntries()([]CatalogEntryable) {
    if m == nil {
        return nil
    } else {
        return m.entries
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Catalog) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["entries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCatalogEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CatalogEntryable, len(val))
            for i, v := range val {
                res[i] = v.(CatalogEntryable)
            }
            m.SetEntries(res)
        }
        return nil
    }
    return res
}
func (m *Catalog) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Catalog) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEntries() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetEntries()))
        for i, v := range m.GetEntries() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("entries", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEntries sets the entries property value. Lists the content that you can approve for deployment. Read-only.
func (m *Catalog) SetEntries(value []CatalogEntryable)() {
    if m != nil {
        m.entries = value
    }
}
