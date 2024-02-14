package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type Logs struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewLogs instantiates a new Logs and sets the default values.
func NewLogs()(*Logs) {
    m := &Logs{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateLogsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLogsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLogs(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Logs) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["remoteNetworks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRemoteNetworkHealthEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RemoteNetworkHealthEventable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RemoteNetworkHealthEventable)
                }
            }
            m.SetRemoteNetworks(res)
        }
        return nil
    }
    res["traffic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateNetworkAccessTrafficFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]NetworkAccessTrafficable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(NetworkAccessTrafficable)
                }
            }
            m.SetTraffic(res)
        }
        return nil
    }
    return res
}
// GetRemoteNetworks gets the remoteNetworks property value. A collection of remote network health events.
// returns a []RemoteNetworkHealthEventable when successful
func (m *Logs) GetRemoteNetworks()([]RemoteNetworkHealthEventable) {
    val, err := m.GetBackingStore().Get("remoteNetworks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RemoteNetworkHealthEventable)
    }
    return nil
}
// GetTraffic gets the traffic property value. A network access traffic log entry that contains comprehensive information about network traffic events.
// returns a []NetworkAccessTrafficable when successful
func (m *Logs) GetTraffic()([]NetworkAccessTrafficable) {
    val, err := m.GetBackingStore().Get("traffic")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]NetworkAccessTrafficable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Logs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRemoteNetworks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRemoteNetworks()))
        for i, v := range m.GetRemoteNetworks() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("remoteNetworks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTraffic() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTraffic()))
        for i, v := range m.GetTraffic() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("traffic", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRemoteNetworks sets the remoteNetworks property value. A collection of remote network health events.
func (m *Logs) SetRemoteNetworks(value []RemoteNetworkHealthEventable)() {
    err := m.GetBackingStore().Set("remoteNetworks", value)
    if err != nil {
        panic(err)
    }
}
// SetTraffic sets the traffic property value. A network access traffic log entry that contains comprehensive information about network traffic events.
func (m *Logs) SetTraffic(value []NetworkAccessTrafficable)() {
    err := m.GetBackingStore().Set("traffic", value)
    if err != nil {
        panic(err)
    }
}
type Logsable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRemoteNetworks()([]RemoteNetworkHealthEventable)
    GetTraffic()([]NetworkAccessTrafficable)
    SetRemoteNetworks(value []RemoteNetworkHealthEventable)()
    SetTraffic(value []NetworkAccessTrafficable)()
}
