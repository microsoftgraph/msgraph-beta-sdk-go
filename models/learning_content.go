package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LearningContent 
type LearningContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Keywords, topics, and other tags associated with the learning content. Optional.
    additionalTags []string
    // The content web URL for the learning content. Required.
    contentWebUrl *string
    // The author, creator, or contributor of the learning content. Optional.
    contributor *string
    // The date when the learning content was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description or summary for the learning content. Optional.
    description *string
    // The duration of the learning content in seconds. Optional.
    duration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // Unique external content ID for the learning content. Required.
    externalId *string
    // The format of the learning content. For example, Course, Video, Book, Book Summary, Audiobook Summary. Optional.
    format *string
    // Indicates whether the content is active or not. Inactive content will not show up in the UI. The default value is true. Optional.
    isActive *bool
    // Indicates whether the learning content requires the user to sign-in on the learning provider platform or not. The default value is false. Optional.
    isPremium *bool
    // Indicates whether the learning content is searchable or not. The default value is true. Optional.
    isSearchable *bool
    // The language of the learning content, for example, en-us or fr-fr. Required.
    languageTag *string
    // The date when the learning content was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of pages of the learning content, for example, 9. Optional.
    numberOfPages *int32
    // The OdataType property
    odataType *string
    // The skills tags associated with the learning content. Optional.
    skillTags []string
    // The source name of the learning content, such as LinkedIn Learning or Coursera. Optional.
    sourceName *string
    // The URL of learning content thumbnail image. Optional.
    thumbnailWebUrl *string
    // The title of the learning content. Required.
    title *string
}
// NewLearningContent instantiates a new learningContent and sets the default values.
func NewLearningContent()(*LearningContent) {
    m := &LearningContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.learningContent";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateLearningContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLearningContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLearningContent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LearningContent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdditionalTags gets the additionalTags property value. Keywords, topics, and other tags associated with the learning content. Optional.
func (m *LearningContent) GetAdditionalTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.additionalTags
    }
}
// GetContentWebUrl gets the contentWebUrl property value. The content web URL for the learning content. Required.
func (m *LearningContent) GetContentWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentWebUrl
    }
}
// GetContributor gets the contributor property value. The author, creator, or contributor of the learning content. Optional.
func (m *LearningContent) GetContributor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contributor
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date when the learning content was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
func (m *LearningContent) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The description or summary for the learning content. Optional.
func (m *LearningContent) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDuration gets the duration property value. The duration of the learning content in seconds. Optional.
func (m *LearningContent) GetDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// GetExternalId gets the externalId property value. Unique external content ID for the learning content. Required.
func (m *LearningContent) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LearningContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAdditionalTags(res)
        }
        return nil
    }
    res["contentWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentWebUrl(val)
        }
        return nil
    }
    res["contributor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContributor(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["duration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val)
        }
        return nil
    }
    res["isActive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    res["isPremium"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPremium(val)
        }
        return nil
    }
    res["isSearchable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSearchable(val)
        }
        return nil
    }
    res["languageTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguageTag(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["numberOfPages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfPages(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["skillTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSkillTags(res)
        }
        return nil
    }
    res["sourceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceName(val)
        }
        return nil
    }
    res["thumbnailWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailWebUrl(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    return res
}
// GetFormat gets the format property value. The format of the learning content. For example, Course, Video, Book, Book Summary, Audiobook Summary. Optional.
func (m *LearningContent) GetFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetIsActive gets the isActive property value. Indicates whether the content is active or not. Inactive content will not show up in the UI. The default value is true. Optional.
func (m *LearningContent) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// GetIsPremium gets the isPremium property value. Indicates whether the learning content requires the user to sign-in on the learning provider platform or not. The default value is false. Optional.
func (m *LearningContent) GetIsPremium()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPremium
    }
}
// GetIsSearchable gets the isSearchable property value. Indicates whether the learning content is searchable or not. The default value is true. Optional.
func (m *LearningContent) GetIsSearchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSearchable
    }
}
// GetLanguageTag gets the languageTag property value. The language of the learning content, for example, en-us or fr-fr. Required.
func (m *LearningContent) GetLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageTag
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date when the learning content was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
func (m *LearningContent) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNumberOfPages gets the numberOfPages property value. The number of pages of the learning content, for example, 9. Optional.
func (m *LearningContent) GetNumberOfPages()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.numberOfPages
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *LearningContent) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
}
// GetSkillTags gets the skillTags property value. The skills tags associated with the learning content. Optional.
func (m *LearningContent) GetSkillTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.skillTags
    }
}
// GetSourceName gets the sourceName property value. The source name of the learning content, such as LinkedIn Learning or Coursera. Optional.
func (m *LearningContent) GetSourceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceName
    }
}
// GetThumbnailWebUrl gets the thumbnailWebUrl property value. The URL of learning content thumbnail image. Optional.
func (m *LearningContent) GetThumbnailWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailWebUrl
    }
}
// GetTitle gets the title property value. The title of the learning content. Required.
func (m *LearningContent) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// Serialize serializes information the current object
func (m *LearningContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdditionalTags() != nil {
        err := writer.WriteCollectionOfStringValues("additionalTags", m.GetAdditionalTags())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contentWebUrl", m.GetContentWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contributor", m.GetContributor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("format", m.GetFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isActive", m.GetIsActive())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPremium", m.GetIsPremium())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSearchable", m.GetIsSearchable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("languageTag", m.GetLanguageTag())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("numberOfPages", m.GetNumberOfPages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetSkillTags() != nil {
        err := writer.WriteCollectionOfStringValues("skillTags", m.GetSkillTags())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceName", m.GetSourceName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("thumbnailWebUrl", m.GetThumbnailWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LearningContent) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdditionalTags sets the additionalTags property value. Keywords, topics, and other tags associated with the learning content. Optional.
func (m *LearningContent) SetAdditionalTags(value []string)() {
    if m != nil {
        m.additionalTags = value
    }
}
// SetContentWebUrl sets the contentWebUrl property value. The content web URL for the learning content. Required.
func (m *LearningContent) SetContentWebUrl(value *string)() {
    if m != nil {
        m.contentWebUrl = value
    }
}
// SetContributor sets the contributor property value. The author, creator, or contributor of the learning content. Optional.
func (m *LearningContent) SetContributor(value *string)() {
    if m != nil {
        m.contributor = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date when the learning content was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
func (m *LearningContent) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. The description or summary for the learning content. Optional.
func (m *LearningContent) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDuration sets the duration property value. The duration of the learning content in seconds. Optional.
func (m *LearningContent) SetDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.duration = value
    }
}
// SetExternalId sets the externalId property value. Unique external content ID for the learning content. Required.
func (m *LearningContent) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetFormat sets the format property value. The format of the learning content. For example, Course, Video, Book, Book Summary, Audiobook Summary. Optional.
func (m *LearningContent) SetFormat(value *string)() {
    if m != nil {
        m.format = value
    }
}
// SetIsActive sets the isActive property value. Indicates whether the content is active or not. Inactive content will not show up in the UI. The default value is true. Optional.
func (m *LearningContent) SetIsActive(value *bool)() {
    if m != nil {
        m.isActive = value
    }
}
// SetIsPremium sets the isPremium property value. Indicates whether the learning content requires the user to sign-in on the learning provider platform or not. The default value is false. Optional.
func (m *LearningContent) SetIsPremium(value *bool)() {
    if m != nil {
        m.isPremium = value
    }
}
// SetIsSearchable sets the isSearchable property value. Indicates whether the learning content is searchable or not. The default value is true. Optional.
func (m *LearningContent) SetIsSearchable(value *bool)() {
    if m != nil {
        m.isSearchable = value
    }
}
// SetLanguageTag sets the languageTag property value. The language of the learning content, for example, en-us or fr-fr. Required.
func (m *LearningContent) SetLanguageTag(value *string)() {
    if m != nil {
        m.languageTag = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date when the learning content was last modified. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Optional.
func (m *LearningContent) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNumberOfPages sets the numberOfPages property value. The number of pages of the learning content, for example, 9. Optional.
func (m *LearningContent) SetNumberOfPages(value *int32)() {
    if m != nil {
        m.numberOfPages = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LearningContent) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
// SetSkillTags sets the skillTags property value. The skills tags associated with the learning content. Optional.
func (m *LearningContent) SetSkillTags(value []string)() {
    if m != nil {
        m.skillTags = value
    }
}
// SetSourceName sets the sourceName property value. The source name of the learning content, such as LinkedIn Learning or Coursera. Optional.
func (m *LearningContent) SetSourceName(value *string)() {
    if m != nil {
        m.sourceName = value
    }
}
// SetThumbnailWebUrl sets the thumbnailWebUrl property value. The URL of learning content thumbnail image. Optional.
func (m *LearningContent) SetThumbnailWebUrl(value *string)() {
    if m != nil {
        m.thumbnailWebUrl = value
    }
}
// SetTitle sets the title property value. The title of the learning content. Required.
func (m *LearningContent) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
