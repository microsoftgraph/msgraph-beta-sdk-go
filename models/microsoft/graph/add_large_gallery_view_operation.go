package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AddLargeGalleryViewOperation 
type AddLargeGalleryViewOperation struct {
    CommsOperation
}
// NewAddLargeGalleryViewOperation instantiates a new addLargeGalleryViewOperation and sets the default values.
func NewAddLargeGalleryViewOperation()(*AddLargeGalleryViewOperation) {
    m := &AddLargeGalleryViewOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreateAddLargeGalleryViewOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAddLargeGalleryViewOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAddLargeGalleryViewOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AddLargeGalleryViewOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AddLargeGalleryViewOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
