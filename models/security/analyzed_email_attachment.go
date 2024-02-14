package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type AnalyzedEmailAttachment struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAnalyzedEmailAttachment instantiates a new AnalyzedEmailAttachment and sets the default values.
func NewAnalyzedEmailAttachment()(*AnalyzedEmailAttachment) {
    m := &AnalyzedEmailAttachment{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAnalyzedEmailAttachmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAnalyzedEmailAttachmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAnalyzedEmailAttachment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AnalyzedEmailAttachment) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *AnalyzedEmailAttachment) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDetonationDetails gets the detonationDetails property value. The detonationDetails property
// returns a DetonationDetailsable when successful
func (m *AnalyzedEmailAttachment) GetDetonationDetails()(DetonationDetailsable) {
    val, err := m.GetBackingStore().Get("detonationDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DetonationDetailsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AnalyzedEmailAttachment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["detonationDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDetonationDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetonationDetails(val.(DetonationDetailsable))
        }
        return nil
    }
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["fileType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileType(val)
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
    res["sha256"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha256(val)
        }
        return nil
    }
    res["threatName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatName(val)
        }
        return nil
    }
    res["threatType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreatType(val.(*ThreatType))
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. The fileName property
// returns a *string when successful
func (m *AnalyzedEmailAttachment) GetFileName()(*string) {
    val, err := m.GetBackingStore().Get("fileName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFileType gets the fileType property value. The fileType property
// returns a *string when successful
func (m *AnalyzedEmailAttachment) GetFileType()(*string) {
    val, err := m.GetBackingStore().Get("fileType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *AnalyzedEmailAttachment) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSha256 gets the sha256 property value. The sha256 property
// returns a *string when successful
func (m *AnalyzedEmailAttachment) GetSha256()(*string) {
    val, err := m.GetBackingStore().Get("sha256")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetThreatName gets the threatName property value. The threatName property
// returns a *string when successful
func (m *AnalyzedEmailAttachment) GetThreatName()(*string) {
    val, err := m.GetBackingStore().Get("threatName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetThreatType gets the threatType property value. The threatType property
// returns a *ThreatType when successful
func (m *AnalyzedEmailAttachment) GetThreatType()(*ThreatType) {
    val, err := m.GetBackingStore().Get("threatType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ThreatType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AnalyzedEmailAttachment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("detonationDetails", m.GetDetonationDetails())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileType", m.GetFileType())
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
        err := writer.WriteStringValue("sha256", m.GetSha256())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("threatName", m.GetThreatName())
        if err != nil {
            return err
        }
    }
    if m.GetThreatType() != nil {
        cast := (*m.GetThreatType()).String()
        err := writer.WriteStringValue("threatType", &cast)
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AnalyzedEmailAttachment) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AnalyzedEmailAttachment) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDetonationDetails sets the detonationDetails property value. The detonationDetails property
func (m *AnalyzedEmailAttachment) SetDetonationDetails(value DetonationDetailsable)() {
    err := m.GetBackingStore().Set("detonationDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetFileName sets the fileName property value. The fileName property
func (m *AnalyzedEmailAttachment) SetFileName(value *string)() {
    err := m.GetBackingStore().Set("fileName", value)
    if err != nil {
        panic(err)
    }
}
// SetFileType sets the fileType property value. The fileType property
func (m *AnalyzedEmailAttachment) SetFileType(value *string)() {
    err := m.GetBackingStore().Set("fileType", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AnalyzedEmailAttachment) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSha256 sets the sha256 property value. The sha256 property
func (m *AnalyzedEmailAttachment) SetSha256(value *string)() {
    err := m.GetBackingStore().Set("sha256", value)
    if err != nil {
        panic(err)
    }
}
// SetThreatName sets the threatName property value. The threatName property
func (m *AnalyzedEmailAttachment) SetThreatName(value *string)() {
    err := m.GetBackingStore().Set("threatName", value)
    if err != nil {
        panic(err)
    }
}
// SetThreatType sets the threatType property value. The threatType property
func (m *AnalyzedEmailAttachment) SetThreatType(value *ThreatType)() {
    err := m.GetBackingStore().Set("threatType", value)
    if err != nil {
        panic(err)
    }
}
type AnalyzedEmailAttachmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDetonationDetails()(DetonationDetailsable)
    GetFileName()(*string)
    GetFileType()(*string)
    GetOdataType()(*string)
    GetSha256()(*string)
    GetThreatName()(*string)
    GetThreatType()(*ThreatType)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDetonationDetails(value DetonationDetailsable)()
    SetFileName(value *string)()
    SetFileType(value *string)()
    SetOdataType(value *string)()
    SetSha256(value *string)()
    SetThreatName(value *string)()
    SetThreatType(value *ThreatType)()
}
