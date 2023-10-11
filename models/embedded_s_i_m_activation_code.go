package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// EmbeddedSIMActivationCode the embedded SIM activation code as provided by the mobile operator.
type EmbeddedSIMActivationCode struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewEmbeddedSIMActivationCode instantiates a new embeddedSIMActivationCode and sets the default values.
func NewEmbeddedSIMActivationCode()(*EmbeddedSIMActivationCode) {
    m := &EmbeddedSIMActivationCode{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEmbeddedSIMActivationCodeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmbeddedSIMActivationCodeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmbeddedSIMActivationCode(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmbeddedSIMActivationCode) GetAdditionalData()(map[string]any) {
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
func (m *EmbeddedSIMActivationCode) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmbeddedSIMActivationCode) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["integratedCircuitCardIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegratedCircuitCardIdentifier(val)
        }
        return nil
    }
    res["matchingIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMatchingIdentifier(val)
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
    res["smdpPlusServerAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmdpPlusServerAddress(val)
        }
        return nil
    }
    return res
}
// GetIntegratedCircuitCardIdentifier gets the integratedCircuitCardIdentifier property value. The Integrated Circuit Card Identifier (ICCID) for this embedded SIM activation code as provided by the mobile operator.
func (m *EmbeddedSIMActivationCode) GetIntegratedCircuitCardIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("integratedCircuitCardIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMatchingIdentifier gets the matchingIdentifier property value. The MatchingIdentifier (MatchingID) as specified in the GSMA Association SGP.22 RSP Technical Specification section 4.1.
func (m *EmbeddedSIMActivationCode) GetMatchingIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("matchingIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EmbeddedSIMActivationCode) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSmdpPlusServerAddress gets the smdpPlusServerAddress property value. The fully qualified domain name of the SM-DP+ server as specified in the GSM Association SPG .22 RSP Technical Specification.
func (m *EmbeddedSIMActivationCode) GetSmdpPlusServerAddress()(*string) {
    val, err := m.GetBackingStore().Get("smdpPlusServerAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EmbeddedSIMActivationCode) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("integratedCircuitCardIdentifier", m.GetIntegratedCircuitCardIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("matchingIdentifier", m.GetMatchingIdentifier())
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
        err := writer.WriteStringValue("smdpPlusServerAddress", m.GetSmdpPlusServerAddress())
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
func (m *EmbeddedSIMActivationCode) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *EmbeddedSIMActivationCode) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIntegratedCircuitCardIdentifier sets the integratedCircuitCardIdentifier property value. The Integrated Circuit Card Identifier (ICCID) for this embedded SIM activation code as provided by the mobile operator.
func (m *EmbeddedSIMActivationCode) SetIntegratedCircuitCardIdentifier(value *string)() {
    err := m.GetBackingStore().Set("integratedCircuitCardIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetMatchingIdentifier sets the matchingIdentifier property value. The MatchingIdentifier (MatchingID) as specified in the GSMA Association SGP.22 RSP Technical Specification section 4.1.
func (m *EmbeddedSIMActivationCode) SetMatchingIdentifier(value *string)() {
    err := m.GetBackingStore().Set("matchingIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EmbeddedSIMActivationCode) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSmdpPlusServerAddress sets the smdpPlusServerAddress property value. The fully qualified domain name of the SM-DP+ server as specified in the GSM Association SPG .22 RSP Technical Specification.
func (m *EmbeddedSIMActivationCode) SetSmdpPlusServerAddress(value *string)() {
    err := m.GetBackingStore().Set("smdpPlusServerAddress", value)
    if err != nil {
        panic(err)
    }
}
// EmbeddedSIMActivationCodeable 
type EmbeddedSIMActivationCodeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIntegratedCircuitCardIdentifier()(*string)
    GetMatchingIdentifier()(*string)
    GetOdataType()(*string)
    GetSmdpPlusServerAddress()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIntegratedCircuitCardIdentifier(value *string)()
    SetMatchingIdentifier(value *string)()
    SetOdataType(value *string)()
    SetSmdpPlusServerAddress(value *string)()
}
