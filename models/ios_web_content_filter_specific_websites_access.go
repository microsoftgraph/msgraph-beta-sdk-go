package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosWebContentFilterSpecificWebsitesAccess represents an iOS Web Content Filter setting type, which installs URL bookmarks into iOS built-in browser. An example scenario is in the classroom where teachers would like the students to navigate websites through browser bookmarks configured on their iOS devices, and no access to other sites.
type IosWebContentFilterSpecificWebsitesAccess struct {
    IosWebContentFilterBase
}
// NewIosWebContentFilterSpecificWebsitesAccess instantiates a new iosWebContentFilterSpecificWebsitesAccess and sets the default values.
func NewIosWebContentFilterSpecificWebsitesAccess()(*IosWebContentFilterSpecificWebsitesAccess) {
    m := &IosWebContentFilterSpecificWebsitesAccess{
        IosWebContentFilterBase: *NewIosWebContentFilterBase(),
    }
    odataTypeValue := "#microsoft.graph.iosWebContentFilterSpecificWebsitesAccess"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosWebContentFilterSpecificWebsitesAccessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosWebContentFilterSpecificWebsitesAccessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosWebContentFilterSpecificWebsitesAccess(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosWebContentFilterSpecificWebsitesAccess) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosWebContentFilterBase.GetFieldDeserializers()
    res["specificWebsitesOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosBookmarkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosBookmarkable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IosBookmarkable)
                }
            }
            m.SetSpecificWebsitesOnly(res)
        }
        return nil
    }
    res["websiteList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosBookmarkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosBookmarkable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(IosBookmarkable)
                }
            }
            m.SetWebsiteList(res)
        }
        return nil
    }
    return res
}
// GetSpecificWebsitesOnly gets the specificWebsitesOnly property value. URL bookmarks which will be installed into built-in browser and user is only allowed to access websites through bookmarks. This collection can contain a maximum of 500 elements.
func (m *IosWebContentFilterSpecificWebsitesAccess) GetSpecificWebsitesOnly()([]IosBookmarkable) {
    val, err := m.GetBackingStore().Get("specificWebsitesOnly")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IosBookmarkable)
    }
    return nil
}
// GetWebsiteList gets the websiteList property value. URL bookmarks which will be installed into built-in browser and user is only allowed to access websites through bookmarks. This collection can contain a maximum of 500 elements.
func (m *IosWebContentFilterSpecificWebsitesAccess) GetWebsiteList()([]IosBookmarkable) {
    val, err := m.GetBackingStore().Get("websiteList")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]IosBookmarkable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IosWebContentFilterSpecificWebsitesAccess) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosWebContentFilterBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSpecificWebsitesOnly() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSpecificWebsitesOnly()))
        for i, v := range m.GetSpecificWebsitesOnly() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("specificWebsitesOnly", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWebsiteList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWebsiteList()))
        for i, v := range m.GetWebsiteList() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("websiteList", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSpecificWebsitesOnly sets the specificWebsitesOnly property value. URL bookmarks which will be installed into built-in browser and user is only allowed to access websites through bookmarks. This collection can contain a maximum of 500 elements.
func (m *IosWebContentFilterSpecificWebsitesAccess) SetSpecificWebsitesOnly(value []IosBookmarkable)() {
    err := m.GetBackingStore().Set("specificWebsitesOnly", value)
    if err != nil {
        panic(err)
    }
}
// SetWebsiteList sets the websiteList property value. URL bookmarks which will be installed into built-in browser and user is only allowed to access websites through bookmarks. This collection can contain a maximum of 500 elements.
func (m *IosWebContentFilterSpecificWebsitesAccess) SetWebsiteList(value []IosBookmarkable)() {
    err := m.GetBackingStore().Set("websiteList", value)
    if err != nil {
        panic(err)
    }
}
// IosWebContentFilterSpecificWebsitesAccessable 
type IosWebContentFilterSpecificWebsitesAccessable interface {
    IosWebContentFilterBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSpecificWebsitesOnly()([]IosBookmarkable)
    GetWebsiteList()([]IosBookmarkable)
    SetSpecificWebsitesOnly(value []IosBookmarkable)()
    SetWebsiteList(value []IosBookmarkable)()
}
