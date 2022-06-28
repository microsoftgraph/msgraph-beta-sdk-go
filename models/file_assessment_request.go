package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileAssessmentRequest 
type FileAssessmentRequest struct {
    ThreatAssessmentRequest
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Base64 encoded file content. The file content cannot fetch back because it isn't stored.
    contentData *string
    // The file name.
    fileName *string
}
// NewFileAssessmentRequest instantiates a new FileAssessmentRequest and sets the default values.
func NewFileAssessmentRequest()(*FileAssessmentRequest) {
    m := &FileAssessmentRequest{
        ThreatAssessmentRequest: *NewThreatAssessmentRequest(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFileAssessmentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFileAssessmentRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFileAssessmentRequest(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileAssessmentRequest) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentData gets the contentData property value. Base64 encoded file content. The file content cannot fetch back because it isn't stored.
func (m *FileAssessmentRequest) GetContentData()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FileAssessmentRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ThreatAssessmentRequest.GetFieldDeserializers()
    res["contentData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentData(val)
        }
        return nil
    }
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. The file name.
func (m *FileAssessmentRequest) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
// Serialize serializes information the current object
func (m *FileAssessmentRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ThreatAssessmentRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contentData", m.GetContentData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FileAssessmentRequest) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentData sets the contentData property value. Base64 encoded file content. The file content cannot fetch back because it isn't stored.
func (m *FileAssessmentRequest) SetContentData(value *string)() {
    if m != nil {
        m.contentData = value
    }
}
// SetFileName sets the fileName property value. The file name.
func (m *FileAssessmentRequest) SetFileName(value *string)() {
    if m != nil {
        m.fileName = value
    }
}
