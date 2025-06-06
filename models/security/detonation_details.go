// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DetonationDetails struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDetonationDetails instantiates a new DetonationDetails and sets the default values.
func NewDetonationDetails()(*DetonationDetails) {
    m := &DetonationDetails{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDetonationDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDetonationDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDetonationDetails(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
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
// GetAnalysisDateTime gets the analysisDateTime property value. The time of detonation.
// returns a *Time when successful
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
// returns a BackingStore when successful
func (m *DetonationDetails) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCompromiseIndicators gets the compromiseIndicators property value. Represents indicators and its associated verdict that suggests whether an email is compromised.
// returns a []CompromiseIndicatorable when successful
func (m *DetonationDetails) GetCompromiseIndicators()([]CompromiseIndicatorable) {
    val, err := m.GetBackingStore().Get("compromiseIndicators")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CompromiseIndicatorable)
    }
    return nil
}
// GetDetonationBehaviourDetails gets the detonationBehaviourDetails property value. Shows the exact events that took place during detonation, and problematic or benign observations that contain URLs, IPs, domains, and files that were found during detonation
// returns a DetonationBehaviourDetailsable when successful
func (m *DetonationDetails) GetDetonationBehaviourDetails()(DetonationBehaviourDetailsable) {
    val, err := m.GetBackingStore().Get("detonationBehaviourDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DetonationBehaviourDetailsable)
    }
    return nil
}
// GetDetonationChain gets the detonationChain property value. The chain of detonation.
// returns a DetonationChainable when successful
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
// GetDetonationObservables gets the detonationObservables property value. All observables in the detonation tree.
// returns a DetonationObservablesable when successful
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
// GetDetonationScreenshotUri gets the detonationScreenshotUri property value. Show any screenshots that were captured during detonation. No screenshots are captured if the URL opens into a link that directly downloads a file. However, you see the downloaded file in the detonation chain.
// returns a *string when successful
func (m *DetonationDetails) GetDetonationScreenshotUri()(*string) {
    val, err := m.GetBackingStore().Get("detonationScreenshotUri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDetonationVerdict gets the detonationVerdict property value. The verdict of the detonation.
// returns a *string when successful
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
// GetDetonationVerdictReason gets the detonationVerdictReason property value. The reason for the verdict of the detonation.
// returns a *string when successful
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
    res["compromiseIndicators"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCompromiseIndicatorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CompromiseIndicatorable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CompromiseIndicatorable)
                }
            }
            m.SetCompromiseIndicators(res)
        }
        return nil
    }
    res["detonationBehaviourDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDetonationBehaviourDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetonationBehaviourDetails(val.(DetonationBehaviourDetailsable))
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
    res["detonationScreenshotUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetonationScreenshotUri(val)
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
// returns a *string when successful
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
    if m.GetCompromiseIndicators() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompromiseIndicators()))
        for i, v := range m.GetCompromiseIndicators() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("compromiseIndicators", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("detonationBehaviourDetails", m.GetDetonationBehaviourDetails())
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
        err := writer.WriteStringValue("detonationScreenshotUri", m.GetDetonationScreenshotUri())
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
// SetAnalysisDateTime sets the analysisDateTime property value. The time of detonation.
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
// SetCompromiseIndicators sets the compromiseIndicators property value. Represents indicators and its associated verdict that suggests whether an email is compromised.
func (m *DetonationDetails) SetCompromiseIndicators(value []CompromiseIndicatorable)() {
    err := m.GetBackingStore().Set("compromiseIndicators", value)
    if err != nil {
        panic(err)
    }
}
// SetDetonationBehaviourDetails sets the detonationBehaviourDetails property value. Shows the exact events that took place during detonation, and problematic or benign observations that contain URLs, IPs, domains, and files that were found during detonation
func (m *DetonationDetails) SetDetonationBehaviourDetails(value DetonationBehaviourDetailsable)() {
    err := m.GetBackingStore().Set("detonationBehaviourDetails", value)
    if err != nil {
        panic(err)
    }
}
// SetDetonationChain sets the detonationChain property value. The chain of detonation.
func (m *DetonationDetails) SetDetonationChain(value DetonationChainable)() {
    err := m.GetBackingStore().Set("detonationChain", value)
    if err != nil {
        panic(err)
    }
}
// SetDetonationObservables sets the detonationObservables property value. All observables in the detonation tree.
func (m *DetonationDetails) SetDetonationObservables(value DetonationObservablesable)() {
    err := m.GetBackingStore().Set("detonationObservables", value)
    if err != nil {
        panic(err)
    }
}
// SetDetonationScreenshotUri sets the detonationScreenshotUri property value. Show any screenshots that were captured during detonation. No screenshots are captured if the URL opens into a link that directly downloads a file. However, you see the downloaded file in the detonation chain.
func (m *DetonationDetails) SetDetonationScreenshotUri(value *string)() {
    err := m.GetBackingStore().Set("detonationScreenshotUri", value)
    if err != nil {
        panic(err)
    }
}
// SetDetonationVerdict sets the detonationVerdict property value. The verdict of the detonation.
func (m *DetonationDetails) SetDetonationVerdict(value *string)() {
    err := m.GetBackingStore().Set("detonationVerdict", value)
    if err != nil {
        panic(err)
    }
}
// SetDetonationVerdictReason sets the detonationVerdictReason property value. The reason for the verdict of the detonation.
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
type DetonationDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalysisDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCompromiseIndicators()([]CompromiseIndicatorable)
    GetDetonationBehaviourDetails()(DetonationBehaviourDetailsable)
    GetDetonationChain()(DetonationChainable)
    GetDetonationObservables()(DetonationObservablesable)
    GetDetonationScreenshotUri()(*string)
    GetDetonationVerdict()(*string)
    GetDetonationVerdictReason()(*string)
    GetOdataType()(*string)
    SetAnalysisDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCompromiseIndicators(value []CompromiseIndicatorable)()
    SetDetonationBehaviourDetails(value DetonationBehaviourDetailsable)()
    SetDetonationChain(value DetonationChainable)()
    SetDetonationObservables(value DetonationObservablesable)()
    SetDetonationScreenshotUri(value *string)()
    SetDetonationVerdict(value *string)()
    SetDetonationVerdictReason(value *string)()
    SetOdataType(value *string)()
}
