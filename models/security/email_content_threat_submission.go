package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmailContentThreatSubmission 
type EmailContentThreatSubmission struct {
    EmailThreatSubmission
}
// NewEmailContentThreatSubmission instantiates a new emailContentThreatSubmission and sets the default values.
func NewEmailContentThreatSubmission()(*EmailContentThreatSubmission) {
    m := &EmailContentThreatSubmission{
        EmailThreatSubmission: *NewEmailThreatSubmission(),
    }
    odataTypeValue := "#microsoft.graph.security.emailContentThreatSubmission"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEmailContentThreatSubmissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailContentThreatSubmissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmailContentThreatSubmission(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailContentThreatSubmission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EmailThreatSubmission.GetFieldDeserializers()
    res["fileContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileContent(val)
        }
        return nil
    }
    return res
}
// GetFileContent gets the fileContent property value. Base64 encoded file content.
func (m *EmailContentThreatSubmission) GetFileContent()(*string) {
    val, err := m.GetBackingStore().Get("fileContent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EmailContentThreatSubmission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EmailThreatSubmission.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("fileContent", m.GetFileContent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFileContent sets the fileContent property value. Base64 encoded file content.
func (m *EmailContentThreatSubmission) SetFileContent(value *string)() {
    err := m.GetBackingStore().Set("fileContent", value)
    if err != nil {
        panic(err)
    }
}
// EmailContentThreatSubmissionable 
type EmailContentThreatSubmissionable interface {
    EmailThreatSubmissionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFileContent()(*string)
    SetFileContent(value *string)()
}
