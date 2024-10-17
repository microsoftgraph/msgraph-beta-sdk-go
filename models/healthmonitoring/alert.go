package healthmonitoring

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type Alert struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewAlert instantiates a new Alert and sets the default values.
func NewAlert()(*Alert) {
    m := &Alert{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateAlertFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAlertFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlert(), nil
}
// GetAlertType gets the alertType property value. The alertType property
// returns a *AlertType when successful
func (m *Alert) GetAlertType()(*AlertType) {
    val, err := m.GetBackingStore().Get("alertType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AlertType)
    }
    return nil
}
// GetCategory gets the category property value. The category property
// returns a *Category when successful
func (m *Alert) GetCategory()(*Category) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Category)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The time when Microsoft Entra Health monitoring generated the alert. Supports $orderby.
// returns a *Time when successful
func (m *Alert) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDocumentation gets the documentation property value. A key-value pair that contains the name of and link to the documentation to aid in investigation of the alert.
// returns a Documentationable when successful
func (m *Alert) GetDocumentation()(Documentationable) {
    val, err := m.GetBackingStore().Get("documentation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Documentationable)
    }
    return nil
}
// GetEnrichment gets the enrichment property value. Investigative information on the alert. This information typically includes counts of impacted objects, which include directory objects such as users, groups, and devices, and a pointer to supporting data.
// returns a Enrichmentable when successful
func (m *Alert) GetEnrichment()(Enrichmentable) {
    val, err := m.GetBackingStore().Get("enrichment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Enrichmentable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Alert) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlertType(val.(*AlertType))
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*Category))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["documentation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDocumentationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDocumentation(val.(Documentationable))
        }
        return nil
    }
    res["enrichment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEnrichmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrichment(val.(Enrichmentable))
        }
        return nil
    }
    res["scenario"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseScenario)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScenario(val.(*Scenario))
        }
        return nil
    }
    res["signals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignalsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignals(val.(Signalsable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAlertState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AlertState))
        }
        return nil
    }
    return res
}
// GetScenario gets the scenario property value. The scenario property
// returns a *Scenario when successful
func (m *Alert) GetScenario()(*Scenario) {
    val, err := m.GetBackingStore().Get("scenario")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Scenario)
    }
    return nil
}
// GetSignals gets the signals property value. The collection of signals that were used in the generation of the alert. These signals are sourced from serviceActivity APIs and are added to the alert as key-value pairs.
// returns a Signalsable when successful
func (m *Alert) GetSignals()(Signalsable) {
    val, err := m.GetBackingStore().Get("signals")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Signalsable)
    }
    return nil
}
// GetState gets the state property value. The state property
// returns a *AlertState when successful
func (m *Alert) GetState()(*AlertState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AlertState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Alert) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAlertType() != nil {
        cast := (*m.GetAlertType()).String()
        err = writer.WriteStringValue("alertType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("documentation", m.GetDocumentation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("enrichment", m.GetEnrichment())
        if err != nil {
            return err
        }
    }
    if m.GetScenario() != nil {
        cast := (*m.GetScenario()).String()
        err = writer.WriteStringValue("scenario", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("signals", m.GetSignals())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAlertType sets the alertType property value. The alertType property
func (m *Alert) SetAlertType(value *AlertType)() {
    err := m.GetBackingStore().Set("alertType", value)
    if err != nil {
        panic(err)
    }
}
// SetCategory sets the category property value. The category property
func (m *Alert) SetCategory(value *Category)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The time when Microsoft Entra Health monitoring generated the alert. Supports $orderby.
func (m *Alert) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDocumentation sets the documentation property value. A key-value pair that contains the name of and link to the documentation to aid in investigation of the alert.
func (m *Alert) SetDocumentation(value Documentationable)() {
    err := m.GetBackingStore().Set("documentation", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrichment sets the enrichment property value. Investigative information on the alert. This information typically includes counts of impacted objects, which include directory objects such as users, groups, and devices, and a pointer to supporting data.
func (m *Alert) SetEnrichment(value Enrichmentable)() {
    err := m.GetBackingStore().Set("enrichment", value)
    if err != nil {
        panic(err)
    }
}
// SetScenario sets the scenario property value. The scenario property
func (m *Alert) SetScenario(value *Scenario)() {
    err := m.GetBackingStore().Set("scenario", value)
    if err != nil {
        panic(err)
    }
}
// SetSignals sets the signals property value. The collection of signals that were used in the generation of the alert. These signals are sourced from serviceActivity APIs and are added to the alert as key-value pairs.
func (m *Alert) SetSignals(value Signalsable)() {
    err := m.GetBackingStore().Set("signals", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. The state property
func (m *Alert) SetState(value *AlertState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
type Alertable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertType()(*AlertType)
    GetCategory()(*Category)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDocumentation()(Documentationable)
    GetEnrichment()(Enrichmentable)
    GetScenario()(*Scenario)
    GetSignals()(Signalsable)
    GetState()(*AlertState)
    SetAlertType(value *AlertType)()
    SetCategory(value *Category)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDocumentation(value Documentationable)()
    SetEnrichment(value Enrichmentable)()
    SetScenario(value *Scenario)()
    SetSignals(value Signalsable)()
    SetState(value *AlertState)()
}
