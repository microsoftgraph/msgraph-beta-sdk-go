package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// DetonationDetails 
type DetonationDetails struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDetonationDetails instantiates a new detonationDetails and sets the default values.
func NewDetonationDetails()(*DetonationDetails) {
    m := &DetonationDetails{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDetonationDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDetonationDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDetonationDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DetonationDetails) GetAdditionalData()(map[string]any) {
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
// GetAnalysisDateTime gets the analysisDateTime property value. The analysisDateTime property
func (m *DetonationDetails) GetAnalysisDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("analysisDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *DetonationDetails) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDetonationChain gets the detonationChain property value. The detonationChain property
func (m *DetonationDetails) GetDetonationChain()(DetonationChainable) {
    val, err := m.GetBackingStore().Get("detonationChain")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DetonationChainable)
    }
    return nil
}
// GetDetonationObservables gets the detonationObservables property value. The detonationObservables property
func (m *DetonationDetails) GetDetonationObservables()(DetonationObservablesable) {
    val, err := m.GetBackingStore().Get("detonationObservables")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DetonationObservablesable)
    }
    return nil
}
// GetDetonationVerdict gets the detonationVerdict property value. The detonationVerdict property
func (m *DetonationDetails) GetDetonationVerdict()(*string) {
    val, err := m.GetBackingStore().Get("detonationVerdict")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDetonationVerdictReason gets the detonationVerdictReason property value. The detonationVerdictReason property
func (m *DetonationDetails) GetDetonationVerdictReason()(*string) {
    val, err := m.GetBackingStore().Get("detonationVerdictReason")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DetonationDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["analysisDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnalysisDateTime(val)
        }
        return nil
    }
    res["detonationChain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDetonationChainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetonationChain(val.(DetonationChainable))
        }
        return nil
    }
    res["detonationObservables"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDetonationObservablesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetonationObservables(val.(DetonationObservablesable))
        }
        return nil
    }
    res["detonationVerdict"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetonationVerdict(val)
        }
        return nil
    }
    res["detonationVerdictReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetonationVerdictReason(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DetonationDetails) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DetonationDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("analysisDateTime", m.GetAnalysisDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("detonationChain", m.GetDetonationChain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("detonationObservables", m.GetDetonationObservables())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("detonationVerdict", m.GetDetonationVerdict())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("detonationVerdictReason", m.GetDetonationVerdictReason())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DetonationDetails) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAnalysisDateTime sets the analysisDateTime property value. The analysisDateTime property
func (m *DetonationDetails) SetAnalysisDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("analysisDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DetonationDetails) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDetonationChain sets the detonationChain property value. The detonationChain property
func (m *DetonationDetails) SetDetonationChain(value DetonationChainable)() {
    err := m.GetBackingStore().Set("detonationChain", value)
    if err != nil {
        panic(err)
    }
}
// SetDetonationObservables sets the detonationObservables property value. The detonationObservables property
func (m *DetonationDetails) SetDetonationObservables(value DetonationObservablesable)() {
    err := m.GetBackingStore().Set("detonationObservables", value)
    if err != nil {
        panic(err)
    }
}
// SetDetonationVerdict sets the detonationVerdict property value. The detonationVerdict property
func (m *DetonationDetails) SetDetonationVerdict(value *string)() {
    err := m.GetBackingStore().Set("detonationVerdict", value)
    if err != nil {
        panic(err)
    }
}
// SetDetonationVerdictReason sets the detonationVerdictReason property value. The detonationVerdictReason property
func (m *DetonationDetails) SetDetonationVerdictReason(value *string)() {
    err := m.GetBackingStore().Set("detonationVerdictReason", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DetonationDetails) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// DetonationDetailsable 
type DetonationDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalysisDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDetonationChain()(DetonationChainable)
    GetDetonationObservables()(DetonationObservablesable)
    GetDetonationVerdict()(*string)
    GetDetonationVerdictReason()(*string)
    GetOdataType()(*string)
    SetAnalysisDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDetonationChain(value DetonationChainable)()
    SetDetonationObservables(value DetonationObservablesable)()
    SetDetonationVerdict(value *string)()
    SetDetonationVerdictReason(value *string)()
    SetOdataType(value *string)()
}
