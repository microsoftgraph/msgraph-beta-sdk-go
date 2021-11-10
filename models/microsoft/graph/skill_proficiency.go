package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SkillProficiency struct {
    ItemFacet
    // Contains categories a user has associated with the skill (for example, personal, professional, hobby).
    categories []string;
    // Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
    collaborationTags []string;
    // Contains a friendly name for the skill.
    displayName *string;
    // Detail of the users proficiency with this skill. Possible values are: elementary, limitedWorking, generalProfessional, advancedProfessional, expert, unknownFutureValue.
    proficiency *SkillProficiencyLevel;
    // 
    thumbnailUrl *string;
    // Contains a link to an information source about the skill.
    webUrl *string;
}
// Instantiates a new skillProficiency and sets the default values.
func NewSkillProficiency()(*SkillProficiency) {
    m := &SkillProficiency{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the categories property value. Contains categories a user has associated with the skill (for example, personal, professional, hobby).
func (m *SkillProficiency) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// Gets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *SkillProficiency) GetCollaborationTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.collaborationTags
    }
}
// Gets the displayName property value. Contains a friendly name for the skill.
func (m *SkillProficiency) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the proficiency property value. Detail of the users proficiency with this skill. Possible values are: elementary, limitedWorking, generalProfessional, advancedProfessional, expert, unknownFutureValue.
func (m *SkillProficiency) GetProficiency()(*SkillProficiencyLevel) {
    if m == nil {
        return nil
    } else {
        return m.proficiency
    }
}
// Gets the thumbnailUrl property value. 
func (m *SkillProficiency) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// Gets the webUrl property value. Contains a link to an information source about the skill.
func (m *SkillProficiency) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *SkillProficiency) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["collaborationTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCollaborationTags(res)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["proficiency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSkillProficiencyLevel)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SkillProficiencyLevel)
            m.SetProficiency(&cast)
        }
        return nil
    }
    res["thumbnailUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
func (m *SkillProficiency) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SkillProficiency) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("collaborationTags", m.GetCollaborationTags())
        if err != nil {
            return err
        }
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
    {
        err = writer.WriteStringValue("thumbnailUrl", m.GetThumbnailUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the categories property value. Contains categories a user has associated with the skill (for example, personal, professional, hobby).
// Parameters:
//  - value : Value to set for the categories property.
func (m *SkillProficiency) SetCategories(value []string)() {
    m.categories = value
}
// Sets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
// Parameters:
//  - value : Value to set for the collaborationTags property.
func (m *SkillProficiency) SetCollaborationTags(value []string)() {
    m.collaborationTags = value
}
// Sets the displayName property value. Contains a friendly name for the skill.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *SkillProficiency) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the proficiency property value. Detail of the users proficiency with this skill. Possible values are: elementary, limitedWorking, generalProfessional, advancedProfessional, expert, unknownFutureValue.
// Parameters:
//  - value : Value to set for the proficiency property.
func (m *SkillProficiency) SetProficiency(value *SkillProficiencyLevel)() {
    m.proficiency = value
}
// Sets the thumbnailUrl property value. 
// Parameters:
//  - value : Value to set for the thumbnailUrl property.
func (m *SkillProficiency) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// Sets the webUrl property value. Contains a link to an information source about the skill.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *SkillProficiency) SetWebUrl(value *string)() {
    m.webUrl = value
}
