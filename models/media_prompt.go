package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MediaPrompt 
type MediaPrompt struct {
    Prompt
    // The OdataType property
    OdataType *string
}
// NewMediaPrompt instantiates a new mediaPrompt and sets the default values.
func NewMediaPrompt()(*MediaPrompt) {
    m := &MediaPrompt{
        Prompt: *NewPrompt(),
    }
    odataTypeValue := "#microsoft.graph.mediaPrompt"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMediaPromptFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaPromptFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMediaPrompt(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaPrompt) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Prompt.GetFieldDeserializers()
    res["loop"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoop(val)
        }
        return nil
    }
    res["mediaInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaInfo(val.(MediaInfoable))
        }
        return nil
    }
    return res
}
// GetLoop gets the loop property value. The loop property
func (m *MediaPrompt) GetLoop()(*int32) {
    val, err := m.GetBackingStore().Get("loop")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMediaInfo gets the mediaInfo property value. The mediaInfo property
func (m *MediaPrompt) GetMediaInfo()(MediaInfoable) {
    val, err := m.GetBackingStore().Get("mediaInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MediaInfoable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MediaPrompt) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Prompt.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("loop", m.GetLoop())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaInfo", m.GetMediaInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLoop sets the loop property value. The loop property
func (m *MediaPrompt) SetLoop(value *int32)() {
    err := m.GetBackingStore().Set("loop", value)
    if err != nil {
        panic(err)
    }
}
// SetMediaInfo sets the mediaInfo property value. The mediaInfo property
func (m *MediaPrompt) SetMediaInfo(value MediaInfoable)() {
    err := m.GetBackingStore().Set("mediaInfo", value)
    if err != nil {
        panic(err)
    }
}
// MediaPromptable 
type MediaPromptable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Promptable
    GetLoop()(*int32)
    GetMediaInfo()(MediaInfoable)
    SetLoop(value *int32)()
    SetMediaInfo(value MediaInfoable)()
}
