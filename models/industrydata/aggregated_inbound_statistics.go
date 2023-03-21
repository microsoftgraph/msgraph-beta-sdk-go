package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AggregatedInboundStatistics 
type AggregatedInboundStatistics struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAggregatedInboundStatistics instantiates a new aggregatedInboundStatistics and sets the default values.
func NewAggregatedInboundStatistics()(*AggregatedInboundStatistics) {
    m := &AggregatedInboundStatistics{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAggregatedInboundStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAggregatedInboundStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAggregatedInboundStatistics(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AggregatedInboundStatistics) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *AggregatedInboundStatistics) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetErrors gets the errors property value. The aggregate count of errors encountered by activities during this run.
func (m *AggregatedInboundStatistics) GetErrors()(*int32) {
    val, err := m.GetBackingStore().Get("errors")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AggregatedInboundStatistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["errors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrors(val)
        }
        return nil
    }
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroups(val)
        }
        return nil
    }
    res["matchedPeopleByRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIndustryDataRunRoleCountMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IndustryDataRunRoleCountMetricable, len(val))
            for i, v := range val {
                res[i] = v.(IndustryDataRunRoleCountMetricable)
            }
            m.SetMatchedPeopleByRole(res)
        }
        return nil
    }
    res["memberships"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberships(val)
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
    res["organizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizations(val)
        }
        return nil
    }
    res["people"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeople(val)
        }
        return nil
    }
    res["unmatchedPeopleByRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIndustryDataRunRoleCountMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IndustryDataRunRoleCountMetricable, len(val))
            for i, v := range val {
                res[i] = v.(IndustryDataRunRoleCountMetricable)
            }
            m.SetUnmatchedPeopleByRole(res)
        }
        return nil
    }
    res["warnings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWarnings(val)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. The aggregate count of active inbound groups processed during the run.
func (m *AggregatedInboundStatistics) GetGroups()(*int32) {
    val, err := m.GetBackingStore().Get("groups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetMatchedPeopleByRole gets the matchedPeopleByRole property value. The aggregate count of active people matched to an Azure Active Directory user, by role.
func (m *AggregatedInboundStatistics) GetMatchedPeopleByRole()([]IndustryDataRunRoleCountMetricable) {
    val, err := m.GetBackingStore().Get("matchedPeopleByRole")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IndustryDataRunRoleCountMetricable)
    }
    return nil
}
// GetMemberships gets the memberships property value. The aggregate count of active inbound memberships processed during the run.
func (m *AggregatedInboundStatistics) GetMemberships()(*int32) {
    val, err := m.GetBackingStore().Get("memberships")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AggregatedInboundStatistics) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOrganizations gets the organizations property value. The aggregate count of active inbound organizations processed during the run.
func (m *AggregatedInboundStatistics) GetOrganizations()(*int32) {
    val, err := m.GetBackingStore().Get("organizations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetPeople gets the people property value. The aggregate count of active inbound people processed during the run.
func (m *AggregatedInboundStatistics) GetPeople()(*int32) {
    val, err := m.GetBackingStore().Get("people")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUnmatchedPeopleByRole gets the unmatchedPeopleByRole property value. The aggregate count of active people not matched to an Azure Active Directory user, by role.
func (m *AggregatedInboundStatistics) GetUnmatchedPeopleByRole()([]IndustryDataRunRoleCountMetricable) {
    val, err := m.GetBackingStore().Get("unmatchedPeopleByRole")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IndustryDataRunRoleCountMetricable)
    }
    return nil
}
// GetWarnings gets the warnings property value. The aggregate count of warnings generated by activities during this run.
func (m *AggregatedInboundStatistics) GetWarnings()(*int32) {
    val, err := m.GetBackingStore().Get("warnings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AggregatedInboundStatistics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AggregatedInboundStatistics) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *AggregatedInboundStatistics) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetErrors sets the errors property value. The aggregate count of errors encountered by activities during this run.
func (m *AggregatedInboundStatistics) SetErrors(value *int32)() {
    err := m.GetBackingStore().Set("errors", value)
    if err != nil {
        panic(err)
    }
}
// SetGroups sets the groups property value. The aggregate count of active inbound groups processed during the run.
func (m *AggregatedInboundStatistics) SetGroups(value *int32)() {
    err := m.GetBackingStore().Set("groups", value)
    if err != nil {
        panic(err)
    }
}
// SetMatchedPeopleByRole sets the matchedPeopleByRole property value. The aggregate count of active people matched to an Azure Active Directory user, by role.
func (m *AggregatedInboundStatistics) SetMatchedPeopleByRole(value []IndustryDataRunRoleCountMetricable)() {
    err := m.GetBackingStore().Set("matchedPeopleByRole", value)
    if err != nil {
        panic(err)
    }
}
// SetMemberships sets the memberships property value. The aggregate count of active inbound memberships processed during the run.
func (m *AggregatedInboundStatistics) SetMemberships(value *int32)() {
    err := m.GetBackingStore().Set("memberships", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AggregatedInboundStatistics) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizations sets the organizations property value. The aggregate count of active inbound organizations processed during the run.
func (m *AggregatedInboundStatistics) SetOrganizations(value *int32)() {
    err := m.GetBackingStore().Set("organizations", value)
    if err != nil {
        panic(err)
    }
}
// SetPeople sets the people property value. The aggregate count of active inbound people processed during the run.
func (m *AggregatedInboundStatistics) SetPeople(value *int32)() {
    err := m.GetBackingStore().Set("people", value)
    if err != nil {
        panic(err)
    }
}
// SetUnmatchedPeopleByRole sets the unmatchedPeopleByRole property value. The aggregate count of active people not matched to an Azure Active Directory user, by role.
func (m *AggregatedInboundStatistics) SetUnmatchedPeopleByRole(value []IndustryDataRunRoleCountMetricable)() {
    err := m.GetBackingStore().Set("unmatchedPeopleByRole", value)
    if err != nil {
        panic(err)
    }
}
// SetWarnings sets the warnings property value. The aggregate count of warnings generated by activities during this run.
func (m *AggregatedInboundStatistics) SetWarnings(value *int32)() {
    err := m.GetBackingStore().Set("warnings", value)
    if err != nil {
        panic(err)
    }
}
// AggregatedInboundStatisticsable 
type AggregatedInboundStatisticsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetErrors()(*int32)
    GetGroups()(*int32)
    GetMatchedPeopleByRole()([]IndustryDataRunRoleCountMetricable)
    GetMemberships()(*int32)
    GetOdataType()(*string)
    GetOrganizations()(*int32)
    GetPeople()(*int32)
    GetUnmatchedPeopleByRole()([]IndustryDataRunRoleCountMetricable)
    GetWarnings()(*int32)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetErrors(value *int32)()
    SetGroups(value *int32)()
    SetMatchedPeopleByRole(value []IndustryDataRunRoleCountMetricable)()
    SetMemberships(value *int32)()
    SetOdataType(value *string)()
    SetOrganizations(value *int32)()
    SetPeople(value *int32)()
    SetUnmatchedPeopleByRole(value []IndustryDataRunRoleCountMetricable)()
    SetWarnings(value *int32)()
}
