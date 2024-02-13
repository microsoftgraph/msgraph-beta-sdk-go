package networkaccess

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FilteringProfile struct {
    Profile
}
// NewFilteringProfile instantiates a new FilteringProfile and sets the default values.
func NewFilteringProfile()(*FilteringProfile) {
    m := &FilteringProfile{
        Profile: *NewProfile(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.filteringProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateFilteringProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilteringProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFilteringProfile(), nil
}
// GetConditionalAccessPolicies gets the conditionalAccessPolicies property value. A set of associated policies defined to regulate access to resources or systems based on specific conditions. Automatically expanded.
// returns a []ConditionalAccessPolicyable when successful
func (m *FilteringProfile) GetConditionalAccessPolicies()([]ConditionalAccessPolicyable) {
    val, err := m.GetBackingStore().Get("conditionalAccessPolicies")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ConditionalAccessPolicyable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the filteringProfile was created.
// returns a *Time when successful
func (m *FilteringProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
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
func (m *FilteringProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Profile.GetFieldDeserializers()
    res["conditionalAccessPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConditionalAccessPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessPolicyable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ConditionalAccessPolicyable)
                }
            }
            m.SetConditionalAccessPolicies(res)
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
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    return res
}
// GetPriority gets the priority property value. The priority used to order the profile for processing within a list.
// returns a *int64 when successful
func (m *FilteringProfile) GetPriority()(*int64) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *FilteringProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Profile.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConditionalAccessPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetConditionalAccessPolicies()))
        for i, v := range m.GetConditionalAccessPolicies() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("conditionalAccessPolicies", cast)
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
        err = writer.WriteInt64Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConditionalAccessPolicies sets the conditionalAccessPolicies property value. A set of associated policies defined to regulate access to resources or systems based on specific conditions. Automatically expanded.
func (m *FilteringProfile) SetConditionalAccessPolicies(value []ConditionalAccessPolicyable)() {
    err := m.GetBackingStore().Set("conditionalAccessPolicies", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the filteringProfile was created.
func (m *FilteringProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. The priority used to order the profile for processing within a list.
func (m *FilteringProfile) SetPriority(value *int64)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
type FilteringProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Profileable
    GetConditionalAccessPolicies()([]ConditionalAccessPolicyable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPriority()(*int64)
    SetConditionalAccessPolicies(value []ConditionalAccessPolicyable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPriority(value *int64)()
}
