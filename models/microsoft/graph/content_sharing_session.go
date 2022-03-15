package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ContentSharingSession provides operations to manage the commsApplication singleton.
type ContentSharingSession struct {
    Entity
}
// NewContentSharingSession instantiates a new contentSharingSession and sets the default values.
func NewContentSharingSession()(*ContentSharingSession) {
    m := &ContentSharingSession{
        Entity: *NewEntity(),
    }
    return m
}
// CreateContentSharingSessionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentSharingSessionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewContentSharingSession(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentSharingSession) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
func (m *ContentSharingSession) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ContentSharingSession) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
