package search

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Qna 
type Qna struct {
    SearchAnswer
}
// NewQna instantiates a new Qna and sets the default values.
func NewQna()(*Qna) {
    m := &Qna{
        SearchAnswer: *NewSearchAnswer(),
    }
    return m
}
// CreateQnaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateQnaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewQna(), nil
}
// GetAvailabilityEndDateTime gets the availabilityEndDateTime property value. Timestamp of when the qna will stop to appear as a search result. Set as null for always available.
func (m *Qna) GetAvailabilityEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("availabilityEndDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetAvailabilityStartDateTime gets the availabilityStartDateTime property value. Timestamp of when the qna will start to appear as a search result. Set as null for always available.
func (m *Qna) GetAvailabilityStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("availabilityStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Qna) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SearchAnswer.GetFieldDeserializers()
    res["availabilityEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityEndDateTime(val)
        }
        return nil
    }
    res["availabilityStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailabilityStartDateTime(val)
        }
        return nil
    }
    res["groupIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["isSuggested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSuggested(val)
        }
        return nil
    }
    res["keywords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAnswerKeywordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeywords(val.(AnswerKeywordable))
        }
        return nil
    }
    res["languageTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
    res["platforms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DevicePlatformType, len(val))
            for i, v := range val {
                res[i] = *(v.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DevicePlatformType))
            }
            m.SetPlatforms(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAnswerState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AnswerState))
        }
        return nil
    }
    res["targetedVariations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAnswerVariantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AnswerVariantable, len(val))
            for i, v := range val {
                res[i] = v.(AnswerVariantable)
            }
            m.SetTargetedVariations(res)
        }
        return nil
    }
    return res
}
// GetGroupIds gets the groupIds property value. List of security groups able to view this qna.
func (m *Qna) GetGroupIds()([]string) {
    val, err := m.GetBackingStore().Get("groupIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetIsSuggested gets the isSuggested property value. True if this qna was suggested to the admin by a user or was mined and suggested by Microsoft. Read-only.
func (m *Qna) GetIsSuggested()(*bool) {
    val, err := m.GetBackingStore().Get("isSuggested")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetKeywords gets the keywords property value. Keywords that trigger this qna to appear in search results.
func (m *Qna) GetKeywords()(AnswerKeywordable) {
    val, err := m.GetBackingStore().Get("keywords")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AnswerKeywordable)
    }
    return nil
}
// GetLanguageTags gets the languageTags property value. A list of language names that are geographically specific and that this QnA can be viewed in. Each language tag value follows the pattern {language}-{region}. As an example, en-us is English as used in the United States. For the list of possible values, see supported language tags.
func (m *Qna) GetLanguageTags()([]string) {
    val, err := m.GetBackingStore().Get("languageTags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetPlatforms gets the platforms property value. List of devices and operating systems able to view this qna. Possible values are: unknown, android, androidForWork, ios, macOS, windowsPhone81, windowsPhone81AndLater, windows10AndLater, androidWorkProfile, androidASOP.
func (m *Qna) GetPlatforms()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DevicePlatformType) {
    val, err := m.GetBackingStore().Get("platforms")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DevicePlatformType)
    }
    return nil
}
// GetState gets the state property value. The state property
func (m *Qna) GetState()(*AnswerState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AnswerState)
    }
    return nil
}
// GetTargetedVariations gets the targetedVariations property value. Variations of a qna for different countries or devices. Use when you need to show different content to users based on their device, country/region, or both. The date and group settings will apply to all variations.
func (m *Qna) GetTargetedVariations()([]AnswerVariantable) {
    val, err := m.GetBackingStore().Get("targetedVariations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AnswerVariantable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Qna) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteCollectionOfStringValues("platforms", ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SerializeDevicePlatformType(m.GetPlatforms()))
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargetedVariations()))
        for i, v := range m.GetTargetedVariations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("targetedVariations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAvailabilityEndDateTime sets the availabilityEndDateTime property value. Timestamp of when the qna will stop to appear as a search result. Set as null for always available.
func (m *Qna) SetAvailabilityEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("availabilityEndDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetAvailabilityStartDateTime sets the availabilityStartDateTime property value. Timestamp of when the qna will start to appear as a search result. Set as null for always available.
func (m *Qna) SetAvailabilityStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("availabilityStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetGroupIds sets the groupIds property value. List of security groups able to view this qna.
func (m *Qna) SetGroupIds(value []string)() {
    err := m.GetBackingStore().Set("groupIds", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSuggested sets the isSuggested property value. True if this qna was suggested to the admin by a user or was mined and suggested by Microsoft. Read-only.
func (m *Qna) SetIsSuggested(value *bool)() {
    err := m.GetBackingStore().Set("isSuggested", value)
    if err != nil {
        panic(err)
    }
}
// SetKeywords sets the keywords property value. Keywords that trigger this qna to appear in search results.
func (m *Qna) SetKeywords(value AnswerKeywordable)() {
    err := m.GetBackingStore().Set("keywords", value)
    if err != nil {
        panic(err)
    }
}
// SetLanguageTags sets the languageTags property value. A list of language names that are geographically specific and that this QnA can be viewed in. Each language tag value follows the pattern {language}-{region}. As an example, en-us is English as used in the United States. For the list of possible values, see supported language tags.
func (m *Qna) SetLanguageTags(value []string)() {
    err := m.GetBackingStore().Set("languageTags", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatforms sets the platforms property value. List of devices and operating systems able to view this qna. Possible values are: unknown, android, androidForWork, ios, macOS, windowsPhone81, windowsPhone81AndLater, windows10AndLater, androidWorkProfile, androidASOP.
func (m *Qna) SetPlatforms(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DevicePlatformType)() {
    err := m.GetBackingStore().Set("platforms", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. The state property
func (m *Qna) SetState(value *AnswerState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetedVariations sets the targetedVariations property value. Variations of a qna for different countries or devices. Use when you need to show different content to users based on their device, country/region, or both. The date and group settings will apply to all variations.
func (m *Qna) SetTargetedVariations(value []AnswerVariantable)() {
    err := m.GetBackingStore().Set("targetedVariations", value)
    if err != nil {
        panic(err)
    }
}
// Qnaable 
type Qnaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SearchAnswerable
    GetAvailabilityEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAvailabilityStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGroupIds()([]string)
    GetIsSuggested()(*bool)
    GetKeywords()(AnswerKeywordable)
    GetLanguageTags()([]string)
    GetPlatforms()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DevicePlatformType)
    GetState()(*AnswerState)
    GetTargetedVariations()([]AnswerVariantable)
    SetAvailabilityEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAvailabilityStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGroupIds(value []string)()
    SetIsSuggested(value *bool)()
    SetKeywords(value AnswerKeywordable)()
    SetLanguageTags(value []string)()
    SetPlatforms(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DevicePlatformType)()
    SetState(value *AnswerState)()
    SetTargetedVariations(value []AnswerVariantable)()
}
