package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/search"
)

// 
type Qna struct {
    SearchAnswer
    // 
    availabilityEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    availabilityStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
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
    state *id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState;
    // 
    targetedVariations []AnswerVariant;
}
// Instantiates a new qna and sets the default values.
func NewQna()(*Qna) {
    m := &Qna{
        SearchAnswer: *NewSearchAnswer(),
    }
    return m
}
// Gets the availabilityEndDateTime property value. 
func (m *Qna) GetAvailabilityEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.availabilityEndDateTime
    }
}
// Gets the availabilityStartDateTime property value. 
func (m *Qna) GetAvailabilityStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.availabilityStartDateTime
    }
}
// Gets the groupIds property value. 
func (m *Qna) GetGroupIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.groupIds
    }
}
// Gets the isSuggested property value. 
func (m *Qna) GetIsSuggested()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSuggested
    }
}
// Gets the keywords property value. 
func (m *Qna) GetKeywords()(*AnswerKeyword) {
    if m == nil {
        return nil
    } else {
        return m.keywords
    }
}
// Gets the languageTags property value. 
func (m *Qna) GetLanguageTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.languageTags
    }
}
// Gets the platforms property value. 
func (m *Qna) GetPlatforms()([]DevicePlatformType) {
    if m == nil {
        return nil
    } else {
        return m.platforms
    }
}
// Gets the state property value. 
func (m *Qna) GetState()(*id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the targetedVariations property value. 
func (m *Qna) GetTargetedVariations()([]AnswerVariant) {
    if m == nil {
        return nil
    } else {
        return m.targetedVariations
    }
}
// The deserialization information for the current model
func (m *Qna) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.SearchAnswer.GetFieldDeserializers()
    res["availabilityEndDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetAvailabilityEndDateTime(val)
        return nil
    }
    res["availabilityStartDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetAvailabilityStartDateTime(val)
        return nil
    }
    res["groupIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetGroupIds(res)
        return nil
    }
    res["isSuggested"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSuggested(val)
        return nil
    }
    res["keywords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAnswerKeyword() })
        if err != nil {
            return err
        }
        m.SetKeywords(val.(*AnswerKeyword))
        return nil
    }
    res["languageTags"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetLanguageTags(res)
        return nil
    }
    res["platforms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        res := make([]DevicePlatformType, len(val))
        for i, v := range val {
            res[i] = *(v.(*DevicePlatformType))
        }
        m.SetPlatforms(res)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.ParseAnswerState)
        if err != nil {
            return err
        }
        cast := val.(id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState)
        m.SetState(&cast)
        return nil
    }
    res["targetedVariations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAnswerVariant() })
        if err != nil {
            return err
        }
        res := make([]AnswerVariant, len(val))
        for i, v := range val {
            res[i] = *(v.(*AnswerVariant))
        }
        m.SetTargetedVariations(res)
        return nil
    }
    return res
}
func (m *Qna) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Qna) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
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
    {
        err = writer.WriteCollectionOfStringValues("languageTags", m.GetLanguageTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("platforms", SerializeDevicePlatformType(m.GetPlatforms()))
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
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
// Sets the availabilityEndDateTime property value. 
// Parameters:
//  - value : Value to set for the availabilityEndDateTime property.
func (m *Qna) SetAvailabilityEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.availabilityEndDateTime = value
}
// Sets the availabilityStartDateTime property value. 
// Parameters:
//  - value : Value to set for the availabilityStartDateTime property.
func (m *Qna) SetAvailabilityStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.availabilityStartDateTime = value
}
// Sets the groupIds property value. 
// Parameters:
//  - value : Value to set for the groupIds property.
func (m *Qna) SetGroupIds(value []string)() {
    m.groupIds = value
}
// Sets the isSuggested property value. 
// Parameters:
//  - value : Value to set for the isSuggested property.
func (m *Qna) SetIsSuggested(value *bool)() {
    m.isSuggested = value
}
// Sets the keywords property value. 
// Parameters:
//  - value : Value to set for the keywords property.
func (m *Qna) SetKeywords(value *AnswerKeyword)() {
    m.keywords = value
}
// Sets the languageTags property value. 
// Parameters:
//  - value : Value to set for the languageTags property.
func (m *Qna) SetLanguageTags(value []string)() {
    m.languageTags = value
}
// Sets the platforms property value. 
// Parameters:
//  - value : Value to set for the platforms property.
func (m *Qna) SetPlatforms(value []DevicePlatformType)() {
    m.platforms = value
}
// Sets the state property value. 
// Parameters:
//  - value : Value to set for the state property.
func (m *Qna) SetState(value *id2242e0abfe0270d8d02377d5aa406c0b4e2307a32628cf8b4c8c6d7176530e8.AnswerState)() {
    m.state = value
}
// Sets the targetedVariations property value. 
// Parameters:
//  - value : Value to set for the targetedVariations property.
func (m *Qna) SetTargetedVariations(value []AnswerVariant)() {
    m.targetedVariations = value
}
