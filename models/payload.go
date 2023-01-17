package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Payload 
type Payload struct {
    Entity
    // The brand property
    brand *PayloadBrand
    // The complexity property
    complexity *PayloadComplexity
    // The createdBy property
    createdBy EmailIdentityable
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description property
    description *string
    // The detail property
    detail PayloadDetailable
    // The displayName property
    displayName *string
    // The industry property
    industry *PayloadIndustry
    // The isAutomated property
    isAutomated *bool
    // The isControversial property
    isControversial *bool
    // The isCurrentEvent property
    isCurrentEvent *bool
    // The language property
    language *string
    // The lastModifiedBy property
    lastModifiedBy EmailIdentityable
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The payloadTags property
    payloadTags []string
    // The platform property
    platform *PayloadDeliveryPlatform
    // The predictedCompromiseRate property
    predictedCompromiseRate *float64
    // The simulationAttackType property
    simulationAttackType *SimulationAttackType
    // The source property
    source *SimulationContentSource
    // The status property
    status *SimulationContentStatus
    // The technique property
    technique *SimulationAttackTechnique
    // The theme property
    theme *PayloadTheme
}
// NewPayload instantiates a new payload and sets the default values.
func NewPayload()(*Payload) {
    m := &Payload{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePayloadFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePayloadFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPayload(), nil
}
// GetBrand gets the brand property value. The brand property
func (m *Payload) GetBrand()(*PayloadBrand) {
    return m.brand
}
// GetComplexity gets the complexity property value. The complexity property
func (m *Payload) GetComplexity()(*PayloadComplexity) {
    return m.complexity
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *Payload) GetCreatedBy()(EmailIdentityable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *Payload) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description property
func (m *Payload) GetDescription()(*string) {
    return m.description
}
// GetDetail gets the detail property value. The detail property
func (m *Payload) GetDetail()(PayloadDetailable) {
    return m.detail
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *Payload) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Payload) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["brand"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadBrand)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrand(val.(*PayloadBrand))
        }
        return nil
    }
    res["complexity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadComplexity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplexity(val.(*PayloadComplexity))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(EmailIdentityable))
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePayloadDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(PayloadDetailable))
        }
        return nil
    }
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
    res["industry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadIndustry)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndustry(val.(*PayloadIndustry))
        }
        return nil
    }
    res["isAutomated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAutomated(val)
        }
        return nil
    }
    res["isControversial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsControversial(val)
        }
        return nil
    }
    res["isCurrentEvent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCurrentEvent(val)
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(EmailIdentityable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["payloadTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPayloadTags(res)
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadDeliveryPlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*PayloadDeliveryPlatform))
        }
        return nil
    }
    res["predictedCompromiseRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPredictedCompromiseRate(val)
        }
        return nil
    }
    res["simulationAttackType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationAttackType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimulationAttackType(val.(*SimulationAttackType))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationContentSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*SimulationContentSource))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationContentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SimulationContentStatus))
        }
        return nil
    }
    res["technique"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationAttackTechnique)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnique(val.(*SimulationAttackTechnique))
        }
        return nil
    }
    res["theme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePayloadTheme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTheme(val.(*PayloadTheme))
        }
        return nil
    }
    return res
}
// GetIndustry gets the industry property value. The industry property
func (m *Payload) GetIndustry()(*PayloadIndustry) {
    return m.industry
}
// GetIsAutomated gets the isAutomated property value. The isAutomated property
func (m *Payload) GetIsAutomated()(*bool) {
    return m.isAutomated
}
// GetIsControversial gets the isControversial property value. The isControversial property
func (m *Payload) GetIsControversial()(*bool) {
    return m.isControversial
}
// GetIsCurrentEvent gets the isCurrentEvent property value. The isCurrentEvent property
func (m *Payload) GetIsCurrentEvent()(*bool) {
    return m.isCurrentEvent
}
// GetLanguage gets the language property value. The language property
func (m *Payload) GetLanguage()(*string) {
    return m.language
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
func (m *Payload) GetLastModifiedBy()(EmailIdentityable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Payload) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPayloadTags gets the payloadTags property value. The payloadTags property
func (m *Payload) GetPayloadTags()([]string) {
    return m.payloadTags
}
// GetPlatform gets the platform property value. The platform property
func (m *Payload) GetPlatform()(*PayloadDeliveryPlatform) {
    return m.platform
}
// GetPredictedCompromiseRate gets the predictedCompromiseRate property value. The predictedCompromiseRate property
func (m *Payload) GetPredictedCompromiseRate()(*float64) {
    return m.predictedCompromiseRate
}
// GetSimulationAttackType gets the simulationAttackType property value. The simulationAttackType property
func (m *Payload) GetSimulationAttackType()(*SimulationAttackType) {
    return m.simulationAttackType
}
// GetSource gets the source property value. The source property
func (m *Payload) GetSource()(*SimulationContentSource) {
    return m.source
}
// GetStatus gets the status property value. The status property
func (m *Payload) GetStatus()(*SimulationContentStatus) {
    return m.status
}
// GetTechnique gets the technique property value. The technique property
func (m *Payload) GetTechnique()(*SimulationAttackTechnique) {
    return m.technique
}
// GetTheme gets the theme property value. The theme property
func (m *Payload) GetTheme()(*PayloadTheme) {
    return m.theme
}
// Serialize serializes information the current object
func (m *Payload) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBrand() != nil {
        cast := (*m.GetBrand()).String()
        err = writer.WriteStringValue("brand", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplexity() != nil {
        cast := (*m.GetComplexity()).String()
        err = writer.WriteStringValue("complexity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetIndustry() != nil {
        cast := (*m.GetIndustry()).String()
        err = writer.WriteStringValue("industry", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAutomated", m.GetIsAutomated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isControversial", m.GetIsControversial())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCurrentEvent", m.GetIsCurrentEvent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPayloadTags() != nil {
        err = writer.WriteCollectionOfStringValues("payloadTags", m.GetPayloadTags())
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("predictedCompromiseRate", m.GetPredictedCompromiseRate())
        if err != nil {
            return err
        }
    }
    if m.GetSimulationAttackType() != nil {
        cast := (*m.GetSimulationAttackType()).String()
        err = writer.WriteStringValue("simulationAttackType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTechnique() != nil {
        cast := (*m.GetTechnique()).String()
        err = writer.WriteStringValue("technique", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTheme() != nil {
        cast := (*m.GetTheme()).String()
        err = writer.WriteStringValue("theme", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBrand sets the brand property value. The brand property
func (m *Payload) SetBrand(value *PayloadBrand)() {
    m.brand = value
}
// SetComplexity sets the complexity property value. The complexity property
func (m *Payload) SetComplexity(value *PayloadComplexity)() {
    m.complexity = value
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *Payload) SetCreatedBy(value EmailIdentityable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *Payload) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description property
func (m *Payload) SetDescription(value *string)() {
    m.description = value
}
// SetDetail sets the detail property value. The detail property
func (m *Payload) SetDetail(value PayloadDetailable)() {
    m.detail = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *Payload) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIndustry sets the industry property value. The industry property
func (m *Payload) SetIndustry(value *PayloadIndustry)() {
    m.industry = value
}
// SetIsAutomated sets the isAutomated property value. The isAutomated property
func (m *Payload) SetIsAutomated(value *bool)() {
    m.isAutomated = value
}
// SetIsControversial sets the isControversial property value. The isControversial property
func (m *Payload) SetIsControversial(value *bool)() {
    m.isControversial = value
}
// SetIsCurrentEvent sets the isCurrentEvent property value. The isCurrentEvent property
func (m *Payload) SetIsCurrentEvent(value *bool)() {
    m.isCurrentEvent = value
}
// SetLanguage sets the language property value. The language property
func (m *Payload) SetLanguage(value *string)() {
    m.language = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *Payload) SetLastModifiedBy(value EmailIdentityable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Payload) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPayloadTags sets the payloadTags property value. The payloadTags property
func (m *Payload) SetPayloadTags(value []string)() {
    m.payloadTags = value
}
// SetPlatform sets the platform property value. The platform property
func (m *Payload) SetPlatform(value *PayloadDeliveryPlatform)() {
    m.platform = value
}
// SetPredictedCompromiseRate sets the predictedCompromiseRate property value. The predictedCompromiseRate property
func (m *Payload) SetPredictedCompromiseRate(value *float64)() {
    m.predictedCompromiseRate = value
}
// SetSimulationAttackType sets the simulationAttackType property value. The simulationAttackType property
func (m *Payload) SetSimulationAttackType(value *SimulationAttackType)() {
    m.simulationAttackType = value
}
// SetSource sets the source property value. The source property
func (m *Payload) SetSource(value *SimulationContentSource)() {
    m.source = value
}
// SetStatus sets the status property value. The status property
func (m *Payload) SetStatus(value *SimulationContentStatus)() {
    m.status = value
}
// SetTechnique sets the technique property value. The technique property
func (m *Payload) SetTechnique(value *SimulationAttackTechnique)() {
    m.technique = value
}
// SetTheme sets the theme property value. The theme property
func (m *Payload) SetTheme(value *PayloadTheme)() {
    m.theme = value
}
