package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// Fileable 
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
