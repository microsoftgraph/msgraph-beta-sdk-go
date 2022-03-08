package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchEntity provides operations to manage the searchEntity singleton.
type SearchEntity struct {
    Entity
    // Administrative answer in Microsoft Search results to define common acronyms in a organization.
    acronyms []Acronymable;
    // Administrative answer in Microsoft Search results for common search queries in an organization.
    bookmarks []Bookmarkable;
    // Administrative answer in Microsoft Search results which provide answers for specific search keywords in an organization.
    qnas []Qnaable;
}
// NewSearchEntity instantiates a new searchEntity and sets the default values.
func NewSearchEntity()(*SearchEntity) {
    m := &SearchEntity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSearchEntityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchEntityFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSearchEntity(), nil
}
// GetAcronyms gets the acronyms property value. Administrative answer in Microsoft Search results to define common acronyms in a organization.
func (m *SearchEntity) GetAcronyms()([]Acronymable) {
    if m == nil {
        return nil
    } else {
        return m.acronyms
    }
}
// GetBookmarks gets the bookmarks property value. Administrative answer in Microsoft Search results for common search queries in an organization.
func (m *SearchEntity) GetBookmarks()([]Bookmarkable) {
    if m == nil {
        return nil
    } else {
        return m.bookmarks
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchEntity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acronyms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAcronymFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Acronymable, len(val))
            for i, v := range val {
                res[i] = v.(Acronymable)
            }
            m.SetAcronyms(res)
        }
        return nil
    }
    res["bookmarks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookmarkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Bookmarkable, len(val))
            for i, v := range val {
                res[i] = v.(Bookmarkable)
            }
            m.SetBookmarks(res)
        }
        return nil
    }
    res["qnas"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateQnaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Qnaable, len(val))
            for i, v := range val {
                res[i] = v.(Qnaable)
            }
            m.SetQnas(res)
        }
        return nil
    }
    return res
}
// GetQnas gets the qnas property value. Administrative answer in Microsoft Search results which provide answers for specific search keywords in an organization.
func (m *SearchEntity) GetQnas()([]Qnaable) {
    if m == nil {
        return nil
    } else {
        return m.qnas
    }
}
func (m *SearchEntity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SearchEntity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAcronyms() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAcronyms()))
        for i, v := range m.GetAcronyms() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("acronyms", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBookmarks() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBookmarks()))
        for i, v := range m.GetBookmarks() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("bookmarks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetQnas() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetQnas()))
        for i, v := range m.GetQnas() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("qnas", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcronyms sets the acronyms property value. Administrative answer in Microsoft Search results to define common acronyms in a organization.
func (m *SearchEntity) SetAcronyms(value []Acronymable)() {
    if m != nil {
        m.acronyms = value
    }
}
// SetBookmarks sets the bookmarks property value. Administrative answer in Microsoft Search results for common search queries in an organization.
func (m *SearchEntity) SetBookmarks(value []Bookmarkable)() {
    if m != nil {
        m.bookmarks = value
    }
}
// SetQnas sets the qnas property value. Administrative answer in Microsoft Search results which provide answers for specific search keywords in an organization.
func (m *SearchEntity) SetQnas(value []Qnaable)() {
    if m != nil {
        m.qnas = value
    }
}
