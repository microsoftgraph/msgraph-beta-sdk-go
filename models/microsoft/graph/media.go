package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Media 
type Media struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // If a file has a transcript, this setting controls if the closed captions / transcription for the media file should be shown to people during viewing. Read-Write.
    isTranscriptionShown *bool;
    // Information about the source of media. Read-only.
    mediaSource *MediaSource;
}
// NewMedia instantiates a new media and sets the default values.
func NewMedia()(*Media) {
    m := &Media{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Media) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetIsTranscriptionShown gets the isTranscriptionShown property value. If a file has a transcript, this setting controls if the closed captions / transcription for the media file should be shown to people during viewing. Read-Write.
func (m *Media) GetIsTranscriptionShown()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isTranscriptionShown
    }
}
// GetMediaSource gets the mediaSource property value. Information about the source of media. Read-only.
func (m *Media) GetMediaSource()(*MediaSource) {
    if m == nil {
        return nil
    } else {
        return m.mediaSource
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Media) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isTranscriptionShown"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTranscriptionShown(val)
        }
        return nil
    }
    res["mediaSource"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMediaSource() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaSource(val.(*MediaSource))
        }
        return nil
    }
    return res
}
func (m *Media) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Media) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isTranscriptionShown", m.GetIsTranscriptionShown())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mediaSource", m.GetMediaSource())
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
func (m *Media) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsTranscriptionShown sets the isTranscriptionShown property value. If a file has a transcript, this setting controls if the closed captions / transcription for the media file should be shown to people during viewing. Read-Write.
func (m *Media) SetIsTranscriptionShown(value *bool)() {
    if m != nil {
        m.isTranscriptionShown = value
    }
}
// SetMediaSource sets the mediaSource property value. Information about the source of media. Read-only.
func (m *Media) SetMediaSource(value *MediaSource)() {
    if m != nil {
        m.mediaSource = value
    }
}
