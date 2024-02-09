package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type ForwardingProfile struct {
    Profile
}
// NewForwardingProfile instantiates a new ForwardingProfile and sets the default values.
func NewForwardingProfile()(*ForwardingProfile) {
    m := &ForwardingProfile{
        Profile: *NewProfile(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.forwardingProfile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateForwardingProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateForwardingProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewForwardingProfile(), nil
}
// GetAssociations gets the associations property value. Specifies the users, groups, devices, and branch locations whose traffic is associated with the given traffic forwarding profile.
// returns a []Associationable when successful
func (m *ForwardingProfile) GetAssociations()([]Associationable) {
    val, err := m.GetBackingStore().Get("associations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Associationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ForwardingProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Profile.GetFieldDeserializers()
    res["associations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssociationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Associationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Associationable)
                }
            }
            m.SetAssociations(res)
        }
        return nil
    }
    res["priority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriority(val)
        }
        return nil
    }
    res["servicePrincipal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateServicePrincipalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipal(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable))
        }
        return nil
    }
    res["trafficForwardingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTrafficForwardingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrafficForwardingType(val.(*TrafficForwardingType))
        }
        return nil
    }
    return res
}
// GetPriority gets the priority property value. Profile priority.
// returns a *int32 when successful
func (m *ForwardingProfile) GetPriority()(*int32) {
    val, err := m.GetBackingStore().Get("priority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetServicePrincipal gets the servicePrincipal property value. The servicePrincipal property
// returns a ServicePrincipalable when successful
func (m *ForwardingProfile) GetServicePrincipal()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable) {
    val, err := m.GetBackingStore().Get("servicePrincipal")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable)
    }
    return nil
}
// GetTrafficForwardingType gets the trafficForwardingType property value. The trafficForwardingType property
// returns a *TrafficForwardingType when successful
func (m *ForwardingProfile) GetTrafficForwardingType()(*TrafficForwardingType) {
    val, err := m.GetBackingStore().Get("trafficForwardingType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TrafficForwardingType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ForwardingProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Profile.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssociations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssociations()))
        for i, v := range m.GetAssociations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("associations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("priority", m.GetPriority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("servicePrincipal", m.GetServicePrincipal())
        if err != nil {
            return err
        }
    }
    if m.GetTrafficForwardingType() != nil {
        cast := (*m.GetTrafficForwardingType()).String()
        err = writer.WriteStringValue("trafficForwardingType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssociations sets the associations property value. Specifies the users, groups, devices, and branch locations whose traffic is associated with the given traffic forwarding profile.
func (m *ForwardingProfile) SetAssociations(value []Associationable)() {
    err := m.GetBackingStore().Set("associations", value)
    if err != nil {
        panic(err)
    }
}
// SetPriority sets the priority property value. Profile priority.
func (m *ForwardingProfile) SetPriority(value *int32)() {
    err := m.GetBackingStore().Set("priority", value)
    if err != nil {
        panic(err)
    }
}
// SetServicePrincipal sets the servicePrincipal property value. The servicePrincipal property
func (m *ForwardingProfile) SetServicePrincipal(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable)() {
    err := m.GetBackingStore().Set("servicePrincipal", value)
    if err != nil {
        panic(err)
    }
}
// SetTrafficForwardingType sets the trafficForwardingType property value. The trafficForwardingType property
func (m *ForwardingProfile) SetTrafficForwardingType(value *TrafficForwardingType)() {
    err := m.GetBackingStore().Set("trafficForwardingType", value)
    if err != nil {
        panic(err)
    }
}
type ForwardingProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Profileable
    GetAssociations()([]Associationable)
    GetPriority()(*int32)
    GetServicePrincipal()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable)
    GetTrafficForwardingType()(*TrafficForwardingType)
    SetAssociations(value []Associationable)()
    SetPriority(value *int32)()
    SetServicePrincipal(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ServicePrincipalable)()
    SetTrafficForwardingType(value *TrafficForwardingType)()
}
