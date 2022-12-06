package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubmissionResult 
type SubmissionResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The submission result category. The possible values are: notJunk, spam, phishing, malware, allowedByPolicy, blockedByPolicy, spoof, unknown, noResultAvailable and unkownFutureValue.
    category *SubmissionResultCategory
    // Specifies the additional details provided by Microsoft to substantiate their analysis result.
    detail *SubmissionResultDetail
    // Specifies the files detected by Microsoft in the submitted emails.
    detectedFiles []SubmissionDetectedFileable
    // Specifes the URLs detected by Microsoft in the submitted email.
    detectedUrls []string
    // The OdataType property
    odataType *string
    // Specifies the setting for user mailbox denoted by a comma-separated string.
    userMailboxSetting *UserMailboxSetting
}
// NewSubmissionResult instantiates a new submissionResult and sets the default values.
func NewSubmissionResult()(*SubmissionResult) {
    m := &SubmissionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSubmissionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubmissionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubmissionResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubmissionResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCategory gets the category property value. The submission result category. The possible values are: notJunk, spam, phishing, malware, allowedByPolicy, blockedByPolicy, spoof, unknown, noResultAvailable and unkownFutureValue.
func (m *SubmissionResult) GetCategory()(*SubmissionResultCategory) {
    return m.category
}
// GetDetail gets the detail property value. Specifies the additional details provided by Microsoft to substantiate their analysis result.
func (m *SubmissionResult) GetDetail()(*SubmissionResultDetail) {
    return m.detail
}
// GetDetectedFiles gets the detectedFiles property value. Specifies the files detected by Microsoft in the submitted emails.
func (m *SubmissionResult) GetDetectedFiles()([]SubmissionDetectedFileable) {
    return m.detectedFiles
}
// GetDetectedUrls gets the detectedUrls property value. Specifes the URLs detected by Microsoft in the submitted email.
func (m *SubmissionResult) GetDetectedUrls()([]string) {
    return m.detectedUrls
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubmissionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["category"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSubmissionResultCategory , m.SetCategory)
    res["detail"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSubmissionResultDetail , m.SetDetail)
    res["detectedFiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSubmissionDetectedFileFromDiscriminatorValue , m.SetDetectedFiles)
    res["detectedUrls"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetDetectedUrls)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["userMailboxSetting"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseUserMailboxSetting , m.SetUserMailboxSetting)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SubmissionResult) GetOdataType()(*string) {
    return m.odataType
}
// GetUserMailboxSetting gets the userMailboxSetting property value. Specifies the setting for user mailbox denoted by a comma-separated string.
func (m *SubmissionResult) GetUserMailboxSetting()(*UserMailboxSetting) {
    return m.userMailboxSetting
}
// Serialize serializes information the current object
func (m *SubmissionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err := writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDetail() != nil {
        cast := (*m.GetDetail()).String()
        err := writer.WriteStringValue("detail", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDetectedFiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetDetectedFiles())
        err := writer.WriteCollectionOfObjectValues("detectedFiles", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDetectedUrls() != nil {
        err := writer.WriteCollectionOfStringValues("detectedUrls", m.GetDetectedUrls())
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
    if m.GetUserMailboxSetting() != nil {
        cast := (*m.GetUserMailboxSetting()).String()
        err := writer.WriteStringValue("userMailboxSetting", &cast)
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
func (m *SubmissionResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCategory sets the category property value. The submission result category. The possible values are: notJunk, spam, phishing, malware, allowedByPolicy, blockedByPolicy, spoof, unknown, noResultAvailable and unkownFutureValue.
func (m *SubmissionResult) SetCategory(value *SubmissionResultCategory)() {
    m.category = value
}
// SetDetail sets the detail property value. Specifies the additional details provided by Microsoft to substantiate their analysis result.
func (m *SubmissionResult) SetDetail(value *SubmissionResultDetail)() {
    m.detail = value
}
// SetDetectedFiles sets the detectedFiles property value. Specifies the files detected by Microsoft in the submitted emails.
func (m *SubmissionResult) SetDetectedFiles(value []SubmissionDetectedFileable)() {
    m.detectedFiles = value
}
// SetDetectedUrls sets the detectedUrls property value. Specifes the URLs detected by Microsoft in the submitted email.
func (m *SubmissionResult) SetDetectedUrls(value []string)() {
    m.detectedUrls = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SubmissionResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserMailboxSetting sets the userMailboxSetting property value. Specifies the setting for user mailbox denoted by a comma-separated string.
func (m *SubmissionResult) SetUserMailboxSetting(value *UserMailboxSetting)() {
    m.userMailboxSetting = value
}
