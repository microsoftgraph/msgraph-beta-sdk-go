package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InboundActivityResults 
type InboundActivityResults struct {
    IndustryDataActivityStatistics
}
// NewInboundActivityResults instantiates a new InboundActivityResults and sets the default values.
func NewInboundActivityResults()(*InboundActivityResults) {
    m := &InboundActivityResults{
        IndustryDataActivityStatistics: *NewIndustryDataActivityStatistics(),
    }
    odataTypeValue := "#microsoft.graph.industryData.inboundActivityResults"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateInboundActivityResultsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInboundActivityResultsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInboundActivityResults(), nil
}
// GetErrors gets the errors property value. Number of errors encountered while processing the inbound flow.
func (m *InboundActivityResults) GetErrors()(*int32) {
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
func (m *InboundActivityResults) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IndustryDataActivityStatistics.GetFieldDeserializers()
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
        val, err := n.GetObjectValue(CreateIndustryDataRunEntityCountMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroups(val.(IndustryDataRunEntityCountMetricable))
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
                if v != nil {
                    res[i] = v.(IndustryDataRunRoleCountMetricable)
                }
            }
            m.SetMatchedPeopleByRole(res)
        }
        return nil
    }
    res["memberships"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIndustryDataRunEntityCountMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberships(val.(IndustryDataRunEntityCountMetricable))
        }
        return nil
    }
    res["organizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIndustryDataRunEntityCountMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizations(val.(IndustryDataRunEntityCountMetricable))
        }
        return nil
    }
    res["people"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIndustryDataRunEntityCountMetricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeople(val.(IndustryDataRunEntityCountMetricable))
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
                if v != nil {
                    res[i] = v.(IndustryDataRunRoleCountMetricable)
                }
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
// GetGroups gets the groups property value. Counts of active and inactive groups processed by the inbound flow.
func (m *InboundActivityResults) GetGroups()(IndustryDataRunEntityCountMetricable) {
    val, err := m.GetBackingStore().Get("groups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IndustryDataRunEntityCountMetricable)
    }
    return nil
}
// GetMatchedPeopleByRole gets the matchedPeopleByRole property value. Number of people matched to an Azure Active Directory user, by role.
func (m *InboundActivityResults) GetMatchedPeopleByRole()([]IndustryDataRunRoleCountMetricable) {
    val, err := m.GetBackingStore().Get("matchedPeopleByRole")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IndustryDataRunRoleCountMetricable)
    }
    return nil
}
// GetMemberships gets the memberships property value. Counts of active and inactive memberships processed by the inbound flow.
func (m *InboundActivityResults) GetMemberships()(IndustryDataRunEntityCountMetricable) {
    val, err := m.GetBackingStore().Get("memberships")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IndustryDataRunEntityCountMetricable)
    }
    return nil
}
// GetOrganizations gets the organizations property value. Counts of active and inactive organizations processed by the inbound flow.
func (m *InboundActivityResults) GetOrganizations()(IndustryDataRunEntityCountMetricable) {
    val, err := m.GetBackingStore().Get("organizations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IndustryDataRunEntityCountMetricable)
    }
    return nil
}
// GetPeople gets the people property value. Counts of active and inactive people processed by the inbound flow.
func (m *InboundActivityResults) GetPeople()(IndustryDataRunEntityCountMetricable) {
    val, err := m.GetBackingStore().Get("people")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IndustryDataRunEntityCountMetricable)
    }
    return nil
}
// GetUnmatchedPeopleByRole gets the unmatchedPeopleByRole property value. Number of people not matched to an Azure Active Directory user, by role.
func (m *InboundActivityResults) GetUnmatchedPeopleByRole()([]IndustryDataRunRoleCountMetricable) {
    val, err := m.GetBackingStore().Get("unmatchedPeopleByRole")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IndustryDataRunRoleCountMetricable)
    }
    return nil
}
// GetWarnings gets the warnings property value. Number of warnings encountered while processing the inbound flow.
func (m *InboundActivityResults) GetWarnings()(*int32) {
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
func (m *InboundActivityResults) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IndustryDataActivityStatistics.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SetErrors sets the errors property value. Number of errors encountered while processing the inbound flow.
func (m *InboundActivityResults) SetErrors(value *int32)() {
    err := m.GetBackingStore().Set("errors", value)
    if err != nil {
        panic(err)
    }
}
// SetGroups sets the groups property value. Counts of active and inactive groups processed by the inbound flow.
func (m *InboundActivityResults) SetGroups(value IndustryDataRunEntityCountMetricable)() {
    err := m.GetBackingStore().Set("groups", value)
    if err != nil {
        panic(err)
    }
}
// SetMatchedPeopleByRole sets the matchedPeopleByRole property value. Number of people matched to an Azure Active Directory user, by role.
func (m *InboundActivityResults) SetMatchedPeopleByRole(value []IndustryDataRunRoleCountMetricable)() {
    err := m.GetBackingStore().Set("matchedPeopleByRole", value)
    if err != nil {
        panic(err)
    }
}
// SetMemberships sets the memberships property value. Counts of active and inactive memberships processed by the inbound flow.
func (m *InboundActivityResults) SetMemberships(value IndustryDataRunEntityCountMetricable)() {
    err := m.GetBackingStore().Set("memberships", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizations sets the organizations property value. Counts of active and inactive organizations processed by the inbound flow.
func (m *InboundActivityResults) SetOrganizations(value IndustryDataRunEntityCountMetricable)() {
    err := m.GetBackingStore().Set("organizations", value)
    if err != nil {
        panic(err)
    }
}
// SetPeople sets the people property value. Counts of active and inactive people processed by the inbound flow.
func (m *InboundActivityResults) SetPeople(value IndustryDataRunEntityCountMetricable)() {
    err := m.GetBackingStore().Set("people", value)
    if err != nil {
        panic(err)
    }
}
// SetUnmatchedPeopleByRole sets the unmatchedPeopleByRole property value. Number of people not matched to an Azure Active Directory user, by role.
func (m *InboundActivityResults) SetUnmatchedPeopleByRole(value []IndustryDataRunRoleCountMetricable)() {
    err := m.GetBackingStore().Set("unmatchedPeopleByRole", value)
    if err != nil {
        panic(err)
    }
}
// SetWarnings sets the warnings property value. Number of warnings encountered while processing the inbound flow.
func (m *InboundActivityResults) SetWarnings(value *int32)() {
    err := m.GetBackingStore().Set("warnings", value)
    if err != nil {
        panic(err)
    }
}
// InboundActivityResultsable 
type InboundActivityResultsable interface {
    IndustryDataActivityStatisticsable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetErrors()(*int32)
    GetGroups()(IndustryDataRunEntityCountMetricable)
    GetMatchedPeopleByRole()([]IndustryDataRunRoleCountMetricable)
    GetMemberships()(IndustryDataRunEntityCountMetricable)
    GetOrganizations()(IndustryDataRunEntityCountMetricable)
    GetPeople()(IndustryDataRunEntityCountMetricable)
    GetUnmatchedPeopleByRole()([]IndustryDataRunRoleCountMetricable)
    GetWarnings()(*int32)
    SetErrors(value *int32)()
    SetGroups(value IndustryDataRunEntityCountMetricable)()
    SetMatchedPeopleByRole(value []IndustryDataRunRoleCountMetricable)()
    SetMemberships(value IndustryDataRunEntityCountMetricable)()
    SetOrganizations(value IndustryDataRunEntityCountMetricable)()
    SetPeople(value IndustryDataRunEntityCountMetricable)()
    SetUnmatchedPeopleByRole(value []IndustryDataRunRoleCountMetricable)()
    SetWarnings(value *int32)()
}
