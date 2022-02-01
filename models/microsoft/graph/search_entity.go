package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchEntity 
type SearchEntity struct {
    Entity
    // 
    acronyms []Acronym;
    // 
    bookmarks []Bookmark;
    // 
    qnas []Qna;
}
// NewSearchEntity instantiates a new searchEntity and sets the default values.
func NewSearchEntity()(*SearchEntity) {
    m := &SearchEntity{
        Entity: *NewEntity(),
    }
    return m
}
// GetAcronyms gets the acronyms property value. 
func (m *SearchEntity) GetAcronyms()([]Acronym) {
    if m == nil {
        return nil
    } else {
        return m.acronyms
    }
}
// GetBookmarks gets the bookmarks property value. 
func (m *SearchEntity) GetBookmarks()([]Bookmark) {
    if m == nil {
        return nil
    } else {
        return m.bookmarks
    }
}
// GetQnas gets the qnas property value. 
func (m *SearchEntity) GetQnas()([]Qna) {
    if m == nil {
        return nil
    } else {
        return m.qnas
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchEntity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["acronyms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAcronym() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Acronym, len(val))
            for i, v := range val {
                res[i] = *(v.(*Acronym))
            }
            m.SetAcronyms(res)
        }
        return nil
    }
    res["bookmarks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBookmark() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Bookmark, len(val))
            for i, v := range val {
                res[i] = *(v.(*Bookmark))
            }
            m.SetBookmarks(res)
        }
        return nil
    }
    res["qnas"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewQna() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Qna, len(val))
            for i, v := range val {
                res[i] = *(v.(*Qna))
            }
            m.SetQnas(res)
        }
        return nil
    }
    return res
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
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("acronyms", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBookmarks() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBookmarks()))
        for i, v := range m.GetBookmarks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("bookmarks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetQnas() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetQnas()))
        for i, v := range m.GetQnas() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("qnas", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcronyms sets the acronyms property value. 
func (m *SearchEntity) SetAcronyms(value []Acronym)() {
    if m != nil {
        m.acronyms = value
    }
}
// SetBookmarks sets the bookmarks property value. 
func (m *SearchEntity) SetBookmarks(value []Bookmark)() {
    if m != nil {
        m.bookmarks = value
    }
}
// SetQnas sets the qnas property value. 
func (m *SearchEntity) SetQnas(value []Qna)() {
    if m != nil {
        m.qnas = value
    }
}
