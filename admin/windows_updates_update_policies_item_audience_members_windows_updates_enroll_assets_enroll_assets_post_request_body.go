package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
)

// WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody 
type WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody instantiates a new WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody()(*WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody) {
    m := &WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody) GetAdditionalData()(map[string]any) {
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
// GetAssets gets the assets property value. The assets property
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody) GetAssets()([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable) {
    val, err := m.GetBackingStore().Get("assets")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, len(val))
            for i, v := range val {
                res[i] = v.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)
            }
            m.SetAssets(res)
        }
        return nil
    }
    res["updateCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ParseUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateCategory(val.(*i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory))
        }
        return nil
    }
    return res
}
// GetUpdateCategory gets the updateCategory property value. The updateCategory property
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody) GetUpdateCategory()(*i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory) {
    val, err := m.GetBackingStore().Get("updateCategory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssets()))
        for i, v := range m.GetAssets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assets", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdateCategory() != nil {
        cast := (*m.GetUpdateCategory()).String()
        err := writer.WriteStringValue("updateCategory", &cast)
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
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAssets sets the assets property value. The assets property
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody) SetAssets(value []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)() {
    err := m.GetBackingStore().Set("assets", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetUpdateCategory sets the updateCategory property value. The updateCategory property
func (m *WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBody) SetUpdateCategory(value *i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory)() {
    err := m.GetBackingStore().Set("updateCategory", value)
    if err != nil {
        panic(err)
    }
}
// WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBodyable 
type WindowsUpdatesUpdatePoliciesItemAudienceMembersWindowsUpdatesEnrollAssetsEnrollAssetsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssets()([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetUpdateCategory()(*i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory)
    SetAssets(value []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetUpdateCategory(value *i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory)()
}
