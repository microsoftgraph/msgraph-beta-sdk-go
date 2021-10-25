package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type LanguageProficiency struct {
    ItemFacet
    displayName *string;
    proficiency *LanguageProficiencyLevel;
    reading *LanguageProficiencyLevel;
    spoken *LanguageProficiencyLevel;
    tag *string;
    thumbnailUrl *string;
    written *LanguageProficiencyLevel;
}
func NewLanguageProficiency()(*LanguageProficiency) {
    m := &LanguageProficiency{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
func (m *LanguageProficiency) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *LanguageProficiency) GetProficiency()(*LanguageProficiencyLevel) {
    if m == nil {
        return nil
    } else {
        return m.proficiency
    }
}
func (m *LanguageProficiency) GetReading()(*LanguageProficiencyLevel) {
    if m == nil {
        return nil
    } else {
        return m.reading
    }
}
func (m *LanguageProficiency) GetSpoken()(*LanguageProficiencyLevel) {
    if m == nil {
        return nil
    } else {
        return m.spoken
    }
}
func (m *LanguageProficiency) GetTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tag
    }
}
func (m *LanguageProficiency) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
func (m *LanguageProficiency) GetWritten()(*LanguageProficiencyLevel) {
    if m == nil {
        return nil
    } else {
        return m.written
    }
}
func (m *LanguageProficiency) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["proficiency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        cast := val.(LanguageProficiencyLevel)
        m.SetProficiency(&cast)
        return nil
    }
    res["reading"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        cast := val.(LanguageProficiencyLevel)
        m.SetReading(&cast)
        return nil
    }
    res["spoken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        cast := val.(LanguageProficiencyLevel)
        m.SetSpoken(&cast)
        return nil
    }
    res["tag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTag(val)
        return nil
    }
    res["thumbnailUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetThumbnailUrl(val)
        return nil
    }
    res["written"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseLanguageProficiencyLevel)
        if err != nil {
            return err
        }
        cast := val.(LanguageProficiencyLevel)
        m.SetWritten(&cast)
        return nil
    }
    return res
}
func (m *LanguageProficiency) IsNil()(bool) {
    return m == nil
}
func (m *LanguageProficiency) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetProficiency() != nil {
        cast := m.GetProficiency().String()
        err = writer.WriteStringValue("proficiency", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetReading() != nil {
        cast := m.GetReading().String()
        err = writer.WriteStringValue("reading", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSpoken() != nil {
        cast := m.GetSpoken().String()
        err = writer.WriteStringValue("spoken", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tag", m.GetTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("thumbnailUrl", m.GetThumbnailUrl())
        if err != nil {
            return err
        }
    }
    if m.GetWritten() != nil {
        cast := m.GetWritten().String()
        err = writer.WriteStringValue("written", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *LanguageProficiency) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *LanguageProficiency) SetProficiency(value *LanguageProficiencyLevel)() {
    m.proficiency = value
}
func (m *LanguageProficiency) SetReading(value *LanguageProficiencyLevel)() {
    m.reading = value
}
func (m *LanguageProficiency) SetSpoken(value *LanguageProficiencyLevel)() {
    m.spoken = value
}
func (m *LanguageProficiency) SetTag(value *string)() {
    m.tag = value
}
func (m *LanguageProficiency) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
func (m *LanguageProficiency) SetWritten(value *LanguageProficiencyLevel)() {
    m.written = value
}
