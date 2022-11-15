package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// File provides operations to manage the collection of accessReviewDecision entities.
type File struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The content property
    content []byte
    // The dateTime property
    dateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The extension property
    extension *string
    // The extractedTextContent property
    extractedTextContent []byte
    // The mediaType property
    mediaType *string
    // The name property
    name *string
    // The otherProperties property
    otherProperties StringValueDictionaryable
    // The processingStatus property
    processingStatus *FileProcessingStatus
    // The senderOrAuthors property
    senderOrAuthors []string
    // The size property
    size *int64
    // The sourceType property
    sourceType *SourceType
    // The subjectTitle property
    subjectTitle *string
}
// NewFile instantiates a new file and sets the default values.
func NewFile()(*File) {
    m := &File{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.security.file";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.security.ediscoveryFile":
                        return NewEdiscoveryFile(), nil
                }
            }
        }
    }
    return NewFile(), nil
}
// GetContent gets the content property value. The content property
func (m *File) GetContent()([]byte) {
    return m.content
}
// GetDateTime gets the dateTime property value. The dateTime property
func (m *File) GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateTime
}
// GetExtension gets the extension property value. The extension property
func (m *File) GetExtension()(*string) {
    return m.extension
}
// GetExtractedTextContent gets the extractedTextContent property value. The extractedTextContent property
func (m *File) GetExtractedTextContent()([]byte) {
    return m.extractedTextContent
}
// GetFieldDeserializers the deserialization information for the current model
func (m *File) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetContent)
    res["dateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetDateTime)
    res["extension"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetExtension)
    res["extractedTextContent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetByteArrayValue(m.SetExtractedTextContent)
    res["mediaType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMediaType)
    res["name"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetName)
    res["otherProperties"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateStringValueDictionaryFromDiscriminatorValue , m.SetOtherProperties)
    res["processingStatus"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseFileProcessingStatus , m.SetProcessingStatus)
    res["senderOrAuthors"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSenderOrAuthors)
    res["size"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetSize)
    res["sourceType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSourceType , m.SetSourceType)
    res["subjectTitle"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubjectTitle)
    return res
}
// GetMediaType gets the mediaType property value. The mediaType property
func (m *File) GetMediaType()(*string) {
    return m.mediaType
}
// GetName gets the name property value. The name property
func (m *File) GetName()(*string) {
    return m.name
}
// GetOtherProperties gets the otherProperties property value. The otherProperties property
func (m *File) GetOtherProperties()(StringValueDictionaryable) {
    return m.otherProperties
}
// GetProcessingStatus gets the processingStatus property value. The processingStatus property
func (m *File) GetProcessingStatus()(*FileProcessingStatus) {
    return m.processingStatus
}
// GetSenderOrAuthors gets the senderOrAuthors property value. The senderOrAuthors property
func (m *File) GetSenderOrAuthors()([]string) {
    return m.senderOrAuthors
}
// GetSize gets the size property value. The size property
func (m *File) GetSize()(*int64) {
    return m.size
}
// GetSourceType gets the sourceType property value. The sourceType property
func (m *File) GetSourceType()(*SourceType) {
    return m.sourceType
}
// GetSubjectTitle gets the subjectTitle property value. The subjectTitle property
func (m *File) GetSubjectTitle()(*string) {
    return m.subjectTitle
}
// Serialize serializes information the current object
func (m *File) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("dateTime", m.GetDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("extension", m.GetExtension())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("extractedTextContent", m.GetExtractedTextContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mediaType", m.GetMediaType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("otherProperties", m.GetOtherProperties())
        if err != nil {
            return err
        }
    }
    if m.GetProcessingStatus() != nil {
        cast := (*m.GetProcessingStatus()).String()
        err = writer.WriteStringValue("processingStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSenderOrAuthors() != nil {
        err = writer.WriteCollectionOfStringValues("senderOrAuthors", m.GetSenderOrAuthors())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    if m.GetSourceType() != nil {
        cast := (*m.GetSourceType()).String()
        err = writer.WriteStringValue("sourceType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectTitle", m.GetSubjectTitle())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content property
func (m *File) SetContent(value []byte)() {
    m.content = value
}
// SetDateTime sets the dateTime property value. The dateTime property
func (m *File) SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateTime = value
}
// SetExtension sets the extension property value. The extension property
func (m *File) SetExtension(value *string)() {
    m.extension = value
}
// SetExtractedTextContent sets the extractedTextContent property value. The extractedTextContent property
func (m *File) SetExtractedTextContent(value []byte)() {
    m.extractedTextContent = value
}
// SetMediaType sets the mediaType property value. The mediaType property
func (m *File) SetMediaType(value *string)() {
    m.mediaType = value
}
// SetName sets the name property value. The name property
func (m *File) SetName(value *string)() {
    m.name = value
}
// SetOtherProperties sets the otherProperties property value. The otherProperties property
func (m *File) SetOtherProperties(value StringValueDictionaryable)() {
    m.otherProperties = value
}
// SetProcessingStatus sets the processingStatus property value. The processingStatus property
func (m *File) SetProcessingStatus(value *FileProcessingStatus)() {
    m.processingStatus = value
}
// SetSenderOrAuthors sets the senderOrAuthors property value. The senderOrAuthors property
func (m *File) SetSenderOrAuthors(value []string)() {
    m.senderOrAuthors = value
}
// SetSize sets the size property value. The size property
func (m *File) SetSize(value *int64)() {
    m.size = value
}
// SetSourceType sets the sourceType property value. The sourceType property
func (m *File) SetSourceType(value *SourceType)() {
    m.sourceType = value
}
// SetSubjectTitle sets the subjectTitle property value. The subjectTitle property
func (m *File) SetSubjectTitle(value *string)() {
    m.subjectTitle = value
}
