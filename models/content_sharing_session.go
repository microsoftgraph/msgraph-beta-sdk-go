package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContentSharingSession struct {
    Entity
}
// NewContentSharingSession instantiates a new ContentSharingSession and sets the default values.
func NewContentSharingSession()(*ContentSharingSession) {
    m := &ContentSharingSession{
        Entity: *NewEntity(),
    }
    return m
}
// CreateContentSharingSessionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContentSharingSessionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentSharingSession(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContentSharingSession) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["pngOfCurrentSlide"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPngOfCurrentSlide(val)
        }
        return nil
    }
    res["presenterParticipantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPresenterParticipantId(val)
        }
        return nil
    }
    return res
}
// GetPngOfCurrentSlide gets the pngOfCurrentSlide property value. The pngOfCurrentSlide property
// returns a []byte when successful
func (m *ContentSharingSession) GetPngOfCurrentSlide()([]byte) {
    val, err := m.GetBackingStore().Get("pngOfCurrentSlide")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetPresenterParticipantId gets the presenterParticipantId property value. The presenterParticipantId property
// returns a *string when successful
func (m *ContentSharingSession) GetPresenterParticipantId()(*string) {
    val, err := m.GetBackingStore().Get("presenterParticipantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ContentSharingSession) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("pngOfCurrentSlide", m.GetPngOfCurrentSlide())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("presenterParticipantId", m.GetPresenterParticipantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPngOfCurrentSlide sets the pngOfCurrentSlide property value. The pngOfCurrentSlide property
func (m *ContentSharingSession) SetPngOfCurrentSlide(value []byte)() {
    err := m.GetBackingStore().Set("pngOfCurrentSlide", value)
    if err != nil {
        panic(err)
    }
}
// SetPresenterParticipantId sets the presenterParticipantId property value. The presenterParticipantId property
func (m *ContentSharingSession) SetPresenterParticipantId(value *string)() {
    err := m.GetBackingStore().Set("presenterParticipantId", value)
    if err != nil {
        panic(err)
    }
}
type ContentSharingSessionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPngOfCurrentSlide()([]byte)
    GetPresenterParticipantId()(*string)
    SetPngOfCurrentSlide(value []byte)()
    SetPresenterParticipantId(value *string)()
}
