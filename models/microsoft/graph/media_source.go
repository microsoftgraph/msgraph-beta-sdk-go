package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MediaSource 
type MediaSource struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Enumeration value that indicates the media content category.
    contentCategory *MediaSourceContentCategory;
}
// NewMediaSource instantiates a new mediaSource and sets the default values.
func NewMediaSource()(*MediaSource) {
    m := &MediaSource{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaSource) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContentCategory gets the contentCategory property value. Enumeration value that indicates the media content category.
func (m *MediaSource) GetContentCategory()(*MediaSourceContentCategory) {
    if m == nil {
        return nil
    } else {
        return m.contentCategory
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaSource) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMediaSourceContentCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCategory(val.(*MediaSourceContentCategory))
        }
        return nil
    }
    return res
}
func (m *MediaSource) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MediaSource) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetContentCategory() != nil {
        cast := (*m.GetContentCategory()).String()
        err := writer.WriteStringValue("contentCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaSource) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetContentCategory sets the contentCategory property value. Enumeration value that indicates the media content category.
func (m *MediaSource) SetContentCategory(value *MediaSourceContentCategory)() {
    if m != nil {
        m.contentCategory = value
    }
}
