package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type Connectivity struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewConnectivity instantiates a new Connectivity and sets the default values.
func NewConnectivity()(*Connectivity) {
    m := &Connectivity{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateConnectivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConnectivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnectivity(), nil
}
// GetBranches gets the branches property value. Branches represent locations for connectivity.
// returns a []BranchSiteable when successful
func (m *Connectivity) GetBranches()([]BranchSiteable) {
    val, err := m.GetBackingStore().Get("branches")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]BranchSiteable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Connectivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["branches"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBranchSiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BranchSiteable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(BranchSiteable)
                }
            }
            m.SetBranches(res)
        }
        return nil
    }
    res["remoteNetworks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRemoteNetworkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RemoteNetworkable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RemoteNetworkable)
                }
            }
            m.SetRemoteNetworks(res)
        }
        return nil
    }
    res["webCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWebCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WebCategoryable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(WebCategoryable)
                }
            }
            m.SetWebCategories(res)
        }
        return nil
    }
    return res
}
// GetRemoteNetworks gets the remoteNetworks property value. The remoteNetworks property
// returns a []RemoteNetworkable when successful
func (m *Connectivity) GetRemoteNetworks()([]RemoteNetworkable) {
    val, err := m.GetBackingStore().Get("remoteNetworks")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RemoteNetworkable)
    }
    return nil
}
// GetWebCategories gets the webCategories property value. The webCategories property
// returns a []WebCategoryable when successful
func (m *Connectivity) GetWebCategories()([]WebCategoryable) {
    val, err := m.GetBackingStore().Get("webCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]WebCategoryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Connectivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBranches() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBranches()))
        for i, v := range m.GetBranches() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("branches", cast)
        if err != nil {
            return err
        }
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
    if m.GetWebCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWebCategories()))
        for i, v := range m.GetWebCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("webCategories", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBranches sets the branches property value. Branches represent locations for connectivity.
func (m *Connectivity) SetBranches(value []BranchSiteable)() {
    err := m.GetBackingStore().Set("branches", value)
    if err != nil {
        panic(err)
    }
}
// SetRemoteNetworks sets the remoteNetworks property value. The remoteNetworks property
func (m *Connectivity) SetRemoteNetworks(value []RemoteNetworkable)() {
    err := m.GetBackingStore().Set("remoteNetworks", value)
    if err != nil {
        panic(err)
    }
}
// SetWebCategories sets the webCategories property value. The webCategories property
func (m *Connectivity) SetWebCategories(value []WebCategoryable)() {
    err := m.GetBackingStore().Set("webCategories", value)
    if err != nil {
        panic(err)
    }
}
type Connectivityable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBranches()([]BranchSiteable)
    GetRemoteNetworks()([]RemoteNetworkable)
    GetWebCategories()([]WebCategoryable)
    SetBranches(value []BranchSiteable)()
    SetRemoteNetworks(value []RemoteNetworkable)()
    SetWebCategories(value []WebCategoryable)()
}
