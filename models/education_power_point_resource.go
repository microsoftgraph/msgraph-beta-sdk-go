package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationPowerPointResource 
type EducationPowerPointResource struct {
    EducationResource
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Location of the file on disk.
    fileUrl *string
}
// NewEducationPowerPointResource instantiates a new EducationPowerPointResource and sets the default values.
func NewEducationPowerPointResource()(*EducationPowerPointResource) {
    m := &EducationPowerPointResource{
        EducationResource: *NewEducationResource(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationPowerPointResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationPowerPointResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationPowerPointResource(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationPowerPointResource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationPowerPointResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationResource.GetFieldDeserializers()
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
// GetFileUrl gets the fileUrl property value. Location of the file on disk.
func (m *EducationPowerPointResource) GetFileUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileUrl
    }
}
// Serialize serializes information the current object
func (m *EducationPowerPointResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationResource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("fileUrl", m.GetFileUrl())
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
func (m *EducationPowerPointResource) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFileUrl sets the fileUrl property value. Location of the file on disk.
func (m *EducationPowerPointResource) SetFileUrl(value *string)() {
    if m != nil {
        m.fileUrl = value
    }
}
