package models

import (
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
    odataTypeValue := "#microsoft.graph.personNamePronounciation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreatePersonNamePronounciationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonNamePronounciationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonNamePronounciation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PersonNamePronounciation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *PersonNamePronounciation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonNamePronounciation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["first"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirst(val)
        }
        return nil
    }
    res["last"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLast(val)
        }
        return nil
    }
    res["maiden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaiden(val)
        }
        return nil
    }
    res["middle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddle(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetFirst gets the first property value. The first property
func (m *PersonNamePronounciation) GetFirst()(*string) {
    if m == nil {
        return nil
    } else {
        return m.first
    }
}
// GetLast gets the last property value. The last property
func (m *PersonNamePronounciation) GetLast()(*string) {
    if m == nil {
        return nil
    } else {
        return m.last
    }
}
// GetMaiden gets the maiden property value. The maiden property
func (m *PersonNamePronounciation) GetMaiden()(*string) {
    if m == nil {
        return nil
    } else {
        return m.maiden
    }
}
// GetMiddle gets the middle property value. The middle property
func (m *PersonNamePronounciation) GetMiddle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.middle
    }
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PersonNamePronounciation) GetOdataType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.odataType
    }
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
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *PersonNamePronounciation) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetFirst sets the first property value. The first property
func (m *PersonNamePronounciation) SetFirst(value *string)() {
    if m != nil {
        m.first = value
    }
}
// SetLast sets the last property value. The last property
func (m *PersonNamePronounciation) SetLast(value *string)() {
    if m != nil {
        m.last = value
    }
}
// SetMaiden sets the maiden property value. The maiden property
func (m *PersonNamePronounciation) SetMaiden(value *string)() {
    if m != nil {
        m.maiden = value
    }
}
// SetMiddle sets the middle property value. The middle property
func (m *PersonNamePronounciation) SetMiddle(value *string)() {
    if m != nil {
        m.middle = value
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PersonNamePronounciation) SetOdataType(value *string)() {
    if m != nil {
        m.odataType = value
    }
}
