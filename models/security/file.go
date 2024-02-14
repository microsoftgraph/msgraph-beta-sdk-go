package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type File struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewFile instantiates a new File and sets the default values.
func NewFile()(*File) {
    m := &File{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
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
// returns a []byte when successful
func (m *File) GetContent()([]byte) {
    val, err := m.GetBackingStore().Get("content")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetDateTime gets the dateTime property value. The dateTime property
// returns a *Time when successful
func (m *File) GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("dateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetExtension gets the extension property value. The extension property
// returns a *string when successful
func (m *File) GetExtension()(*string) {
    val, err := m.GetBackingStore().Get("extension")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExtractedTextContent gets the extractedTextContent property value. The extractedTextContent property
// returns a []byte when successful
func (m *File) GetExtractedTextContent()([]byte) {
    val, err := m.GetBackingStore().Get("extractedTextContent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *File) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["dateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateTime(val)
        }
        return nil
    }
    res["extension"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtension(val)
        }
        return nil
    }
    res["extractedTextContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtractedTextContent(val)
        }
        return nil
    }
    res["mediaType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaType(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["otherProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateStringValueDictionaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOtherProperties(val.(StringValueDictionaryable))
        }
        return nil
    }
    res["processingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFileProcessingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingStatus(val.(*FileProcessingStatus))
        }
        return nil
    }
    res["senderOrAuthors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSenderOrAuthors(res)
        }
        return nil
    }
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    res["sourceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceType(val.(*SourceType))
        }
        return nil
    }
    res["subjectTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectTitle(val)
        }
        return nil
    }
    return res
}
// GetMediaType gets the mediaType property value. The mediaType property
// returns a *string when successful
func (m *File) GetMediaType()(*string) {
    val, err := m.GetBackingStore().Get("mediaType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *File) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOtherProperties gets the otherProperties property value. The otherProperties property
// returns a StringValueDictionaryable when successful
func (m *File) GetOtherProperties()(StringValueDictionaryable) {
    val, err := m.GetBackingStore().Get("otherProperties")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(StringValueDictionaryable)
    }
    return nil
}
// GetProcessingStatus gets the processingStatus property value. The processingStatus property
// returns a *FileProcessingStatus when successful
func (m *File) GetProcessingStatus()(*FileProcessingStatus) {
    val, err := m.GetBackingStore().Get("processingStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*FileProcessingStatus)
    }
    return nil
}
// GetSenderOrAuthors gets the senderOrAuthors property value. The senderOrAuthors property
// returns a []string when successful
func (m *File) GetSenderOrAuthors()([]string) {
    val, err := m.GetBackingStore().Get("senderOrAuthors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSize gets the size property value. The size property
// returns a *int64 when successful
func (m *File) GetSize()(*int64) {
    val, err := m.GetBackingStore().Get("size")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetSourceType gets the sourceType property value. The sourceType property
// returns a *SourceType when successful
func (m *File) GetSourceType()(*SourceType) {
    val, err := m.GetBackingStore().Get("sourceType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SourceType)
    }
    return nil
}
// GetSubjectTitle gets the subjectTitle property value. The subjectTitle property
// returns a *string when successful
func (m *File) GetSubjectTitle()(*string) {
    val, err := m.GetBackingStore().Get("subjectTitle")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    err := m.GetBackingStore().Set("content", value)
    if err != nil {
        panic(err)
    }
}
// SetDateTime sets the dateTime property value. The dateTime property
func (m *File) SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("dateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetExtension sets the extension property value. The extension property
func (m *File) SetExtension(value *string)() {
    err := m.GetBackingStore().Set("extension", value)
    if err != nil {
        panic(err)
    }
}
// SetExtractedTextContent sets the extractedTextContent property value. The extractedTextContent property
func (m *File) SetExtractedTextContent(value []byte)() {
    err := m.GetBackingStore().Set("extractedTextContent", value)
    if err != nil {
        panic(err)
    }
}
// SetMediaType sets the mediaType property value. The mediaType property
func (m *File) SetMediaType(value *string)() {
    err := m.GetBackingStore().Set("mediaType", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name property
func (m *File) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOtherProperties sets the otherProperties property value. The otherProperties property
func (m *File) SetOtherProperties(value StringValueDictionaryable)() {
    err := m.GetBackingStore().Set("otherProperties", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessingStatus sets the processingStatus property value. The processingStatus property
func (m *File) SetProcessingStatus(value *FileProcessingStatus)() {
    err := m.GetBackingStore().Set("processingStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetSenderOrAuthors sets the senderOrAuthors property value. The senderOrAuthors property
func (m *File) SetSenderOrAuthors(value []string)() {
    err := m.GetBackingStore().Set("senderOrAuthors", value)
    if err != nil {
        panic(err)
    }
}
// SetSize sets the size property value. The size property
func (m *File) SetSize(value *int64)() {
    err := m.GetBackingStore().Set("size", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceType sets the sourceType property value. The sourceType property
func (m *File) SetSourceType(value *SourceType)() {
    err := m.GetBackingStore().Set("sourceType", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectTitle sets the subjectTitle property value. The subjectTitle property
func (m *File) SetSubjectTitle(value *string)() {
    err := m.GetBackingStore().Set("subjectTitle", value)
    if err != nil {
        panic(err)
    }
}
type Fileable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]byte)
    GetDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExtension()(*string)
    GetExtractedTextContent()([]byte)
    GetMediaType()(*string)
    GetName()(*string)
    GetOtherProperties()(StringValueDictionaryable)
    GetProcessingStatus()(*FileProcessingStatus)
    GetSenderOrAuthors()([]string)
    GetSize()(*int64)
    GetSourceType()(*SourceType)
    GetSubjectTitle()(*string)
    SetContent(value []byte)()
    SetDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExtension(value *string)()
    SetExtractedTextContent(value []byte)()
    SetMediaType(value *string)()
    SetName(value *string)()
    SetOtherProperties(value StringValueDictionaryable)()
    SetProcessingStatus(value *FileProcessingStatus)()
    SetSenderOrAuthors(value []string)()
    SetSize(value *int64)()
    SetSourceType(value *SourceType)()
    SetSubjectTitle(value *string)()
}
