package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// SubmissionResult 
type SubmissionResult struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewSubmissionResult instantiates a new submissionResult and sets the default values.
func NewSubmissionResult()(*SubmissionResult) {
    m := &SubmissionResult{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSubmissionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubmissionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubmissionResult(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubmissionResult) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *SubmissionResult) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCategory gets the category property value. The submission result category. The possible values are: notJunk, spam, phishing, malware, allowedByPolicy, blockedByPolicy, spoof, unknown, noResultAvailable and unkownFutureValue.
func (m *SubmissionResult) GetCategory()(*SubmissionResult_category) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubmissionResult_category)
    }
    return nil
}
// GetDetail gets the detail property value. Specifies the extra details provided by Microsoft to substantiate their analysis result.
func (m *SubmissionResult) GetDetail()(*SubmissionResult_detail) {
    val, err := m.GetBackingStore().Get("detail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubmissionResult_detail)
    }
    return nil
}
// GetDetectedFiles gets the detectedFiles property value. Specifies the files detected by Microsoft in the submitted emails.
func (m *SubmissionResult) GetDetectedFiles()([]SubmissionDetectedFileable) {
    val, err := m.GetBackingStore().Get("detectedFiles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SubmissionDetectedFileable)
    }
    return nil
}
// GetDetectedUrls gets the detectedUrls property value. Specifies the URLs detected by Microsoft in the submitted email.
func (m *SubmissionResult) GetDetectedUrls()([]string) {
    val, err := m.GetBackingStore().Get("detectedUrls")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubmissionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubmissionResult_category)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*SubmissionResult_category))
        }
        return nil
    }
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubmissionResult_detail)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(*SubmissionResult_detail))
        }
        return nil
    }
    res["detectedFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubmissionDetectedFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubmissionDetectedFileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SubmissionDetectedFileable)
                }
            }
            m.SetDetectedFiles(res)
        }
        return nil
    }
    res["detectedUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetDetectedUrls(res)
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
    res["userMailboxSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubmissionResult_userMailboxSetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserMailboxSetting(val.(*SubmissionResult_userMailboxSetting))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SubmissionResult) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserMailboxSetting gets the userMailboxSetting property value. Specifies the setting for user mailbox denoted by a comma-separated string.
func (m *SubmissionResult) GetUserMailboxSetting()(*SubmissionResult_userMailboxSetting) {
    val, err := m.GetBackingStore().Get("userMailboxSetting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubmissionResult_userMailboxSetting)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetectedFiles()))
        for i, v := range m.GetDetectedFiles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubmissionResult) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *SubmissionResult) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCategory sets the category property value. The submission result category. The possible values are: notJunk, spam, phishing, malware, allowedByPolicy, blockedByPolicy, spoof, unknown, noResultAvailable and unkownFutureValue.
func (m *SubmissionResult) SetCategory(value *SubmissionResult_category)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetDetail sets the detail property value. Specifies the extra details provided by Microsoft to substantiate their analysis result.
func (m *SubmissionResult) SetDetail(value *SubmissionResult_detail)() {
    err := m.GetBackingStore().Set("detail", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectedFiles sets the detectedFiles property value. Specifies the files detected by Microsoft in the submitted emails.
func (m *SubmissionResult) SetDetectedFiles(value []SubmissionDetectedFileable)() {
    err := m.GetBackingStore().Set("detectedFiles", value)
    if err != nil {
        panic(err)
    }
}
// SetDetectedUrls sets the detectedUrls property value. Specifies the URLs detected by Microsoft in the submitted email.
func (m *SubmissionResult) SetDetectedUrls(value []string)() {
    err := m.GetBackingStore().Set("detectedUrls", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SubmissionResult) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetUserMailboxSetting sets the userMailboxSetting property value. Specifies the setting for user mailbox denoted by a comma-separated string.
func (m *SubmissionResult) SetUserMailboxSetting(value *SubmissionResult_userMailboxSetting)() {
    err := m.GetBackingStore().Set("userMailboxSetting", value)
    if err != nil {
        panic(err)
    }
}
// SubmissionResultable 
type SubmissionResultable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCategory()(*SubmissionResult_category)
    GetDetail()(*SubmissionResult_detail)
    GetDetectedFiles()([]SubmissionDetectedFileable)
    GetDetectedUrls()([]string)
    GetOdataType()(*string)
    GetUserMailboxSetting()(*SubmissionResult_userMailboxSetting)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCategory(value *SubmissionResult_category)()
    SetDetail(value *SubmissionResult_detail)()
    SetDetectedFiles(value []SubmissionDetectedFileable)()
    SetDetectedUrls(value []string)()
    SetOdataType(value *string)()
    SetUserMailboxSetting(value *SubmissionResult_userMailboxSetting)()
}
