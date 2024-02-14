package industrydata

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type InboundFlow struct {
    IndustryDataActivity
}
// NewInboundFlow instantiates a new InboundFlow and sets the default values.
func NewInboundFlow()(*InboundFlow) {
    m := &InboundFlow{
        IndustryDataActivity: *NewIndustryDataActivity(),
    }
    odataTypeValue := "#microsoft.graph.industryData.inboundFlow"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateInboundFlowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateInboundFlowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.industryData.inboundFileFlow":
                        return NewInboundFileFlow(), nil
                }
            }
        }
    }
    return NewInboundFlow(), nil
}
// GetDataConnector gets the dataConnector property value. The dataConnector property
// returns a IndustryDataConnectorable when successful
func (m *InboundFlow) GetDataConnector()(IndustryDataConnectorable) {
    val, err := m.GetBackingStore().Get("dataConnector")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IndustryDataConnectorable)
    }
    return nil
}
// GetDataDomain gets the dataDomain property value. The dataDomain property
// returns a *InboundDomain when successful
func (m *InboundFlow) GetDataDomain()(*InboundDomain) {
    val, err := m.GetBackingStore().Get("dataDomain")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*InboundDomain)
    }
    return nil
}
// GetEffectiveDateTime gets the effectiveDateTime property value. The start of the time window when the flow is allowed to run. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *InboundFlow) GetEffectiveDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("effectiveDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The end of the time window when the flow is allowed to run. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *InboundFlow) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *InboundFlow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IndustryDataActivity.GetFieldDeserializers()
    res["dataConnector"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIndustryDataConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataConnector(val.(IndustryDataConnectorable))
        }
        return nil
    }
    res["dataDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInboundDomain)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataDomain(val.(*InboundDomain))
        }
        return nil
    }
    res["effectiveDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEffectiveDateTime(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["year"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateYearTimePeriodDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYear(val.(YearTimePeriodDefinitionable))
        }
        return nil
    }
    return res
}
// GetYear gets the year property value. The year property
// returns a YearTimePeriodDefinitionable when successful
func (m *InboundFlow) GetYear()(YearTimePeriodDefinitionable) {
    val, err := m.GetBackingStore().Get("year")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(YearTimePeriodDefinitionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *InboundFlow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IndustryDataActivity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("dataConnector", m.GetDataConnector())
        if err != nil {
            return err
        }
    }
    if m.GetDataDomain() != nil {
        cast := (*m.GetDataDomain()).String()
        err = writer.WriteStringValue("dataDomain", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("effectiveDateTime", m.GetEffectiveDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("year", m.GetYear())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDataConnector sets the dataConnector property value. The dataConnector property
func (m *InboundFlow) SetDataConnector(value IndustryDataConnectorable)() {
    err := m.GetBackingStore().Set("dataConnector", value)
    if err != nil {
        panic(err)
    }
}
// SetDataDomain sets the dataDomain property value. The dataDomain property
func (m *InboundFlow) SetDataDomain(value *InboundDomain)() {
    err := m.GetBackingStore().Set("dataDomain", value)
    if err != nil {
        panic(err)
    }
}
// SetEffectiveDateTime sets the effectiveDateTime property value. The start of the time window when the flow is allowed to run. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *InboundFlow) SetEffectiveDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("effectiveDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The end of the time window when the flow is allowed to run. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *InboundFlow) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetYear sets the year property value. The year property
func (m *InboundFlow) SetYear(value YearTimePeriodDefinitionable)() {
    err := m.GetBackingStore().Set("year", value)
    if err != nil {
        panic(err)
    }
}
type InboundFlowable interface {
    IndustryDataActivityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataConnector()(IndustryDataConnectorable)
    GetDataDomain()(*InboundDomain)
    GetEffectiveDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetYear()(YearTimePeriodDefinitionable)
    SetDataConnector(value IndustryDataConnectorable)()
    SetDataDomain(value *InboundDomain)()
    SetEffectiveDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetYear(value YearTimePeriodDefinitionable)()
}
