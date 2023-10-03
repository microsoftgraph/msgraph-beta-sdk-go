package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileUrlThreatSubmission 
type FileUrlThreatSubmission struct {
    FileThreatSubmission
}
// NewFileUrlThreatSubmission instantiates a new fileUrlThreatSubmission and sets the default values.
func NewFileUrlThreatSubmission()(*FileUrlThreatSubmission) {
    m := &FileUrlThreatSubmission{
        FileThreatSubmission: *NewFileThreatSubmission(),
    }
    odataTypeValue := "#microsoft.graph.security.fileUrlThreatSubmission"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateFileUrlThreatSubmissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileUrlThreatSubmissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileUrlThreatSubmission(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileUrlThreatSubmission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.FileThreatSubmission.GetFieldDeserializers()
    res["fileUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileUrl(val)
        }
        return nil
    }
    return res
}
// GetFileUrl gets the fileUrl property value. It specifies the URL of the file that needs to be submitted.
func (m *FileUrlThreatSubmission) GetFileUrl()(*string) {
    val, err := m.GetBackingStore().Get("fileUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *FileUrlThreatSubmission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.FileThreatSubmission.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("fileUrl", m.GetFileUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFileUrl sets the fileUrl property value. It specifies the URL of the file that needs to be submitted.
func (m *FileUrlThreatSubmission) SetFileUrl(value *string)() {
    err := m.GetBackingStore().Set("fileUrl", value)
    if err != nil {
        panic(err)
    }
}
// FileUrlThreatSubmissionable 
type FileUrlThreatSubmissionable interface {
    FileThreatSubmissionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFileUrl()(*string)
    SetFileUrl(value *string)()
}
