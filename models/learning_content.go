package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LearningContent 
type LearningContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The additionalTags property
    additionalTags []string
    // The contentWebUrl property
    contentWebUrl *string
    // The contributor property
    contributor *string
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // The duration property
    duration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The externalId property
    externalId *string
    // The format property
    format *string
    // The isActive property
    isActive *bool
    // The isPremium property
    isPremium *bool
    // The isSearchable property
    isSearchable *bool
    // The languageTag property
    languageTag *string
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The numberOfPages property
    numberOfPages *int32
    // The OdataType property
    odataType *string
    // The skillTags property
    skillTags []string
    // The sourceName property
    sourceName *string
    // The thumbnailWebUrl property
    thumbnailWebUrl *string
    // The title property
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
// GetAdditionalTags gets the additionalTags property value. The additionalTags property
func (m *LearningContent) GetAdditionalTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.additionalTags
    }
}
// GetContentWebUrl gets the contentWebUrl property value. The contentWebUrl property
func (m *LearningContent) GetContentWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentWebUrl
    }
}
// GetContributor gets the contributor property value. The contributor property
func (m *LearningContent) GetContributor()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contributor
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *LearningContent) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. The description property
func (m *LearningContent) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDuration gets the duration property value. The duration property
func (m *LearningContent) GetDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// GetExternalId gets the externalId property value. The externalId property
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
// GetFormat gets the format property value. The format property
func (m *LearningContent) GetFormat()(*string) {
    if m == nil {
        return nil
    } else {
        return m.format
    }
}
// GetIsActive gets the isActive property value. The isActive property
func (m *LearningContent) GetIsActive()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isActive
    }
}
// GetIsPremium gets the isPremium property value. The isPremium property
func (m *LearningContent) GetIsPremium()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPremium
    }
}
// GetIsSearchable gets the isSearchable property value. The isSearchable property
func (m *LearningContent) GetIsSearchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSearchable
    }
}
// GetLanguageTag gets the languageTag property value. The languageTag property
func (m *LearningContent) GetLanguageTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.languageTag
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *LearningContent) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetNumberOfPages gets the numberOfPages property value. The numberOfPages property
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
// GetSkillTags gets the skillTags property value. The skillTags property
func (m *LearningContent) GetSkillTags()([]string) {
    if m == nil {
        return nil
    } else {
        return m.skillTags
    }
}
// GetSourceName gets the sourceName property value. The sourceName property
func (m *LearningContent) GetSourceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sourceName
    }
}
// GetThumbnailWebUrl gets the thumbnailWebUrl property value. The thumbnailWebUrl property
func (m *LearningContent) GetThumbnailWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.thumbnailWebUrl
    }
}
// GetTitle gets the title property value. The title property
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
// SetAdditionalTags sets the additionalTags property value. The additionalTags property
func (m *LearningContent) SetAdditionalTags(value []string)() {
    if m != nil {
        m.additionalTags = value
    }
}
// SetContentWebUrl sets the contentWebUrl property value. The contentWebUrl property
func (m *LearningContent) SetContentWebUrl(value *string)() {
    if m != nil {
        m.contentWebUrl = value
    }
}
// SetContributor sets the contributor property value. The contributor property
func (m *LearningContent) SetContributor(value *string)() {
    if m != nil {
        m.contributor = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *LearningContent) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. The description property
func (m *LearningContent) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDuration sets the duration property value. The duration property
func (m *LearningContent) SetDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    if m != nil {
        m.duration = value
    }
}
// SetExternalId sets the externalId property value. The externalId property
func (m *LearningContent) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetFormat sets the format property value. The format property
func (m *LearningContent) SetFormat(value *string)() {
    if m != nil {
        m.format = value
    }
}
// SetIsActive sets the isActive property value. The isActive property
func (m *LearningContent) SetIsActive(value *bool)() {
    if m != nil {
        m.isActive = value
    }
}
// SetIsPremium sets the isPremium property value. The isPremium property
func (m *LearningContent) SetIsPremium(value *bool)() {
    if m != nil {
        m.isPremium = value
    }
}
// SetIsSearchable sets the isSearchable property value. The isSearchable property
func (m *LearningContent) SetIsSearchable(value *bool)() {
    if m != nil {
        m.isSearchable = value
    }
}
// SetLanguageTag sets the languageTag property value. The languageTag property
func (m *LearningContent) SetLanguageTag(value *string)() {
    if m != nil {
        m.languageTag = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *LearningContent) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetNumberOfPages sets the numberOfPages property value. The numberOfPages property
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
// SetSkillTags sets the skillTags property value. The skillTags property
func (m *LearningContent) SetSkillTags(value []string)() {
    if m != nil {
        m.skillTags = value
    }
}
// SetSourceName sets the sourceName property value. The sourceName property
func (m *LearningContent) SetSourceName(value *string)() {
    if m != nil {
        m.sourceName = value
    }
}
// SetThumbnailWebUrl sets the thumbnailWebUrl property value. The thumbnailWebUrl property
func (m *LearningContent) SetThumbnailWebUrl(value *string)() {
    if m != nil {
        m.thumbnailWebUrl = value
    }
}
// SetTitle sets the title property value. The title property
func (m *LearningContent) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
