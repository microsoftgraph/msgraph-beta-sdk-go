package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonNamePronounciation 
type PersonNamePronounciation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The displayName property
    displayName *string
    // The first property
    first *string
    // The last property
    last *string
    // The maiden property
    maiden *string
    // The middle property
    middle *string
    // The OdataType property
    odataType *string
}
// NewPersonNamePronounciation instantiates a new personNamePronounciation and sets the default values.
func NewPersonNamePronounciation()(*PersonNamePronounciation) {
    m := &PersonNamePronounciation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePersonNamePronounciationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonNamePronounciationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonNamePronounciation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PersonNamePronounciation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *PersonNamePronounciation) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonNamePronounciation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["first"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetFirst)
    res["last"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetLast)
    res["maiden"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMaiden)
    res["middle"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMiddle)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetFirst gets the first property value. The first property
func (m *PersonNamePronounciation) GetFirst()(*string) {
    return m.first
}
// GetLast gets the last property value. The last property
func (m *PersonNamePronounciation) GetLast()(*string) {
    return m.last
}
// GetMaiden gets the maiden property value. The maiden property
func (m *PersonNamePronounciation) GetMaiden()(*string) {
    return m.maiden
}
// GetMiddle gets the middle property value. The middle property
func (m *PersonNamePronounciation) GetMiddle()(*string) {
    return m.middle
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PersonNamePronounciation) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *PersonNamePronounciation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("first", m.GetFirst())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("last", m.GetLast())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("maiden", m.GetMaiden())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("middle", m.GetMiddle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *PersonNamePronounciation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *PersonNamePronounciation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFirst sets the first property value. The first property
func (m *PersonNamePronounciation) SetFirst(value *string)() {
    m.first = value
}
// SetLast sets the last property value. The last property
func (m *PersonNamePronounciation) SetLast(value *string)() {
    m.last = value
}
// SetMaiden sets the maiden property value. The maiden property
func (m *PersonNamePronounciation) SetMaiden(value *string)() {
    m.maiden = value
}
// SetMiddle sets the middle property value. The middle property
func (m *PersonNamePronounciation) SetMiddle(value *string)() {
    m.middle = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PersonNamePronounciation) SetOdataType(value *string)() {
    m.odataType = value
}
