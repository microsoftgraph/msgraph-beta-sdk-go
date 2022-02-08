package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/search"
)

// Bookmark 
type Bookmark struct {
    SearchAnswer
    // 
    availabilityEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    availabilityStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    categories []string;
    // 
    groupIds []string;
    // 
    isSuggested *bool;
    // 
    keywords *AnswerKeyword;
    // 
    languageTags []string;
    // 
    platforms []DevicePlatformType;
    // 
    powerAppIds []string;
    // 
    state *id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState;
    // 
    targetedVariations []AnswerVariant;
}
// NewBookmark instantiates a new bookmark and sets the default values.
func NewBookmark()(*Bookmark) {
    m := &Bookmark{
        SearchAnswer: *NewSearchAnswer(),
    }
    return m
}
// GetAvailabilityEndDateTime gets the availabilityEndDateTime property value. 
func (m *Bookmark) GetAvailabilityEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.availabilityEndDateTime
    }
}
// GetAvailabilityStartDateTime gets the availabilityStartDateTime property value. 
func (m *Bookmark) GetAvailabilityStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.availabilityStartDateTime
    }
}
// GetCategories gets the categories property value. 
func (m *Bookmark) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// GetGroupIds gets the groupIds property value. 
func (m *Bookmark) GetGroupIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.groupIds
    }
}
// GetIsSuggested gets the isSuggested property value. 
func (m *Bookmark) GetIsSuggested()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSuggested
    }
}
// GetKeywords gets the keywords property value. 
func (m *Bookmark) GetKeywords()(*AnswerKeyword) {
    if m == nil {
        return nil
    } else {
        return m.keywords
    }
}
// GetLanguageTags gets the languageTags property value. 
func (m *Bookmark) GetLanguageTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.languageTags
    }
}
// GetPlatforms gets the platforms property value. 
func (m *Bookmark) GetPlatforms()([]DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platforms
    }
}
// GetPowerAppIds gets the powerAppIds property value. 
func (m *Bookmark) GetPowerAppIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.powerAppIds
    }
}
// GetState gets the state property value. 
func (m *Bookmark) GetState()(*id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetTargetedVariations gets the targetedVariations property value. 
func (m *Bookmark) GetTargetedVariations()([]AnswerVariant) {
    if m == nil {
        return nil
    } else {
        return m.targetedVariations
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Bookmark) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.SearchAnswer.GetFieldDeserializers()
    res["availabilityEndDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityEndDateTime(val)
        }
        return nil
    }
    res["availabilityStartDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityStartDateTime(val)
        }
        return nil
    }
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
    res["groupIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupIds(res)
        }
        return nil
    }
    res["isSuggested"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSuggested(val)
        }
        return nil
    }
    res["keywords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAnswerKeyword() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeywords(val.(*AnswerKeyword))
        }
        return nil
    }
    res["languageTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetLanguageTags(res)
        }
        return nil
    }
    res["platforms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DevicePlatformType, len(val))
            for i, v := range val {
                res[i] = *(v.(*DevicePlatformType))
            }
            m.SetPlatforms(res)
        }
        return nil
    }
    res["powerAppIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPowerAppIds(res)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.ParseAnswerState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState))
        }
        return nil
    }
    res["targetedVariations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAnswerVariant() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AnswerVariant, len(val))
            for i, v := range val {
                res[i] = *(v.(*AnswerVariant))
            }
            m.SetTargetedVariations(res)
        }
        return nil
    }
    return res
}
func (m *Bookmark) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Bookmark) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.SearchAnswer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("availabilityEndDateTime", m.GetAvailabilityEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("availabilityStartDateTime", m.GetAvailabilityStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        err = writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    if m.GetGroupIds() != nil {
        err = writer.WriteCollectionOfStringValues("groupIds", m.GetGroupIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSuggested", m.GetIsSuggested())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("keywords", m.GetKeywords())
        if err != nil {
            return err
        }
    }
    if m.GetLanguageTags() != nil {
        err = writer.WriteCollectionOfStringValues("languageTags", m.GetLanguageTags())
        if err != nil {
            return err
        }
    }
    if m.GetPlatforms() != nil {
        err = writer.WriteCollectionOfStringValues("platforms", SerializeDevicePlatformType(m.GetPlatforms()))
        if err != nil {
            return err
        }
    }
    if m.GetPowerAppIds() != nil {
        err = writer.WriteCollectionOfStringValues("powerAppIds", m.GetPowerAppIds())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetedVariations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTargetedVariations()))
        for i, v := range m.GetTargetedVariations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("targetedVariations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAvailabilityEndDateTime sets the availabilityEndDateTime property value. 
func (m *Bookmark) SetAvailabilityEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.availabilityEndDateTime = value
    }
}
// SetAvailabilityStartDateTime sets the availabilityStartDateTime property value. 
func (m *Bookmark) SetAvailabilityStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.availabilityStartDateTime = value
    }
}
// SetCategories sets the categories property value. 
func (m *Bookmark) SetCategories(value []string)() {
    if m != nil {
        m.categories = value
    }
}
// SetGroupIds sets the groupIds property value. 
func (m *Bookmark) SetGroupIds(value []string)() {
    if m != nil {
        m.groupIds = value
    }
}
// SetIsSuggested sets the isSuggested property value. 
func (m *Bookmark) SetIsSuggested(value *bool)() {
    if m != nil {
        m.isSuggested = value
    }
}
// SetKeywords sets the keywords property value. 
func (m *Bookmark) SetKeywords(value *AnswerKeyword)() {
    if m != nil {
        m.keywords = value
    }
}
// SetLanguageTags sets the languageTags property value. 
func (m *Bookmark) SetLanguageTags(value []string)() {
    if m != nil {
        m.languageTags = value
    }
}
// SetPlatforms sets the platforms property value. 
func (m *Bookmark) SetPlatforms(value []DevicePlatformType)() {
    if m != nil {
        m.platforms = value
    }
}
// SetPowerAppIds sets the powerAppIds property value. 
func (m *Bookmark) SetPowerAppIds(value []string)() {
    if m != nil {
        m.powerAppIds = value
    }
}
// SetState sets the state property value. 
func (m *Bookmark) SetState(value *id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState)() {
    if m != nil {
        m.state = value
    }
}
// SetTargetedVariations sets the targetedVariations property value. 
func (m *Bookmark) SetTargetedVariations(value []AnswerVariant)() {
    if m != nil {
        m.targetedVariations = value
    }
}
