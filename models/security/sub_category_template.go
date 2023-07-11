package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubCategoryTemplate 
type SubCategoryTemplate struct {
    FilePlanDescriptorTemplate
    // The OdataType property
    OdataType *string
}
// NewSubCategoryTemplate instantiates a new subCategoryTemplate and sets the default values.
func NewSubCategoryTemplate()(*SubCategoryTemplate) {
    m := &SubCategoryTemplate{
        FilePlanDescriptorTemplate: *NewFilePlanDescriptorTemplate(),
    }
    return m
}
// CreateSubCategoryTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubCategoryTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubCategoryTemplate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubCategoryTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.FilePlanDescriptorTemplate.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SubCategoryTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.FilePlanDescriptorTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SubCategoryTemplateable 
type SubCategoryTemplateable interface {
    FilePlanDescriptorTemplateable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
