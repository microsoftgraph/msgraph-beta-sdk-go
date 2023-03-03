package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SkillProficiency 
type SkillProficiency struct {
    ItemFacet
}
// NewSkillProficiency instantiates a new SkillProficiency and sets the default values.
func NewSkillProficiency()(*SkillProficiency) {
    m := &SkillProficiency{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.skillProficiency"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSkillProficiencyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSkillProficiencyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSkillProficiency(), nil
}
// GetCategories gets the categories property value. Contains categories a user has associated with the skill (for example, personal, professional, hobby).
func (m *SkillProficiency) GetCategories()([]string) {
    val, err := m.GetBackingStore().Get("categories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetCollaborationTags gets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *SkillProficiency) GetCollaborationTags()([]string) {
    val, err := m.GetBackingStore().Get("collaborationTags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Contains a friendly name for the skill.
func (m *SkillProficiency) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SkillProficiency) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["collaborationTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["proficiency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSkillProficiencyLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProficiency(val.(*SkillProficiencyLevel))
        }
        return nil
    }
    res["thumbnailUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetProficiency gets the proficiency property value. Detail of the users proficiency with this skill. Possible values are: elementary, limitedWorking, generalProfessional, advancedProfessional, expert, unknownFutureValue.
func (m *SkillProficiency) GetProficiency()(*SkillProficiencyLevel) {
    val, err := m.GetBackingStore().Get("proficiency")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SkillProficiencyLevel)
    }
    return nil
}
// GetThumbnailUrl gets the thumbnailUrl property value. The thumbnailUrl property
func (m *SkillProficiency) GetThumbnailUrl()(*string) {
    val, err := m.GetBackingStore().Get("thumbnailUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWebUrl gets the webUrl property value. Contains a link to an information source about the skill.
func (m *SkillProficiency) GetWebUrl()(*string) {
    val, err := m.GetBackingStore().Get("webUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SkillProficiency) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategories() != nil {
        err = writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    if m.GetCollaborationTags() != nil {
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
        cast := (*m.GetProficiency()).String()
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
// SetCategories sets the categories property value. Contains categories a user has associated with the skill (for example, personal, professional, hobby).
func (m *SkillProficiency) SetCategories(value []string)() {
    err := m.GetBackingStore().Set("categories", value)
    if err != nil {
        panic(err)
    }
}
// SetCollaborationTags sets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *SkillProficiency) SetCollaborationTags(value []string)() {
    err := m.GetBackingStore().Set("collaborationTags", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Contains a friendly name for the skill.
func (m *SkillProficiency) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetProficiency sets the proficiency property value. Detail of the users proficiency with this skill. Possible values are: elementary, limitedWorking, generalProfessional, advancedProfessional, expert, unknownFutureValue.
func (m *SkillProficiency) SetProficiency(value *SkillProficiencyLevel)() {
    err := m.GetBackingStore().Set("proficiency", value)
    if err != nil {
        panic(err)
    }
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *SkillProficiency) SetThumbnailUrl(value *string)() {
    err := m.GetBackingStore().Set("thumbnailUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetWebUrl sets the webUrl property value. Contains a link to an information source about the skill.
func (m *SkillProficiency) SetWebUrl(value *string)() {
    err := m.GetBackingStore().Set("webUrl", value)
    if err != nil {
        panic(err)
    }
}
// SkillProficiencyable 
type SkillProficiencyable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategories()([]string)
    GetCollaborationTags()([]string)
    GetDisplayName()(*string)
    GetProficiency()(*SkillProficiencyLevel)
    GetThumbnailUrl()(*string)
    GetWebUrl()(*string)
    SetCategories(value []string)()
    SetCollaborationTags(value []string)()
    SetDisplayName(value *string)()
    SetProficiency(value *SkillProficiencyLevel)()
    SetThumbnailUrl(value *string)()
    SetWebUrl(value *string)()
}
