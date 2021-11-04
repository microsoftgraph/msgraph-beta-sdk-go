package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PersonInterest struct {
    ItemFacet
    // Contains categories a user has associated with the interest (for example, personal, recipies).
    categories []string;
    // Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
    collaborationTags []string;
    // Contains a description of the interest.
    description *string;
    // Contains a friendly name for the interest.
    displayName *string;
    // 
    thumbnailUrl *string;
    // Contains a link to a web page or resource about the interest.
    webUrl *string;
}
// Instantiates a new personInterest and sets the default values.
func NewPersonInterest()(*PersonInterest) {
    m := &PersonInterest{
        ItemFacet: *NewItemFacet(),
    }
    return m
}
// Gets the categories property value. Contains categories a user has associated with the interest (for example, personal, recipies).
func (m *PersonInterest) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// Gets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
func (m *PersonInterest) GetCollaborationTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.collaborationTags
    }
}
// Gets the description property value. Contains a description of the interest.
func (m *PersonInterest) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Contains a friendly name for the interest.
func (m *PersonInterest) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the thumbnailUrl property value. 
func (m *PersonInterest) GetThumbnailUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailUrl
    }
}
// Gets the webUrl property value. Contains a link to a web page or resource about the interest.
func (m *PersonInterest) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *PersonInterest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetCategories(res)
        return nil
    }
    res["collaborationTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetCollaborationTags(res)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
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
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebUrl(val)
        return nil
    }
    return res
}
func (m *PersonInterest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PersonInterest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("description", m.GetDescription())
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
// Sets the categories property value. Contains categories a user has associated with the interest (for example, personal, recipies).
// Parameters:
//  - value : Value to set for the categories property.
func (m *PersonInterest) SetCategories(value []string)() {
    m.categories = value
}
// Sets the collaborationTags property value. Contains experience scenario tags a user has associated with the interest. Allowed values in the collection are: askMeAbout, ableToMentor, wantsToLearn, wantsToImprove.
// Parameters:
//  - value : Value to set for the collaborationTags property.
func (m *PersonInterest) SetCollaborationTags(value []string)() {
    m.collaborationTags = value
}
// Sets the description property value. Contains a description of the interest.
// Parameters:
//  - value : Value to set for the description property.
func (m *PersonInterest) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Contains a friendly name for the interest.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *PersonInterest) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the thumbnailUrl property value. 
// Parameters:
//  - value : Value to set for the thumbnailUrl property.
func (m *PersonInterest) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// Sets the webUrl property value. Contains a link to a web page or resource about the interest.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *PersonInterest) SetWebUrl(value *string)() {
    m.webUrl = value
}
