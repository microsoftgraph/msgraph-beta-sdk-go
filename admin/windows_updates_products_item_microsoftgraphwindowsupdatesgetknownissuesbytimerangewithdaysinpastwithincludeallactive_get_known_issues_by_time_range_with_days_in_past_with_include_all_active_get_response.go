package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
)

type WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponse
}
// NewWindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse instantiates a new WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse and sets the default values.
func NewWindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse()(*WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse) {
    m := &WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse{
        BaseCollectionPaginationCountResponse: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateWindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateKnownIssueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnownIssueable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnownIssueable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []KnownIssueable when successful
func (m *WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse) GetValue()([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnownIssueable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnownIssueable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponse) SetValue(value []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnownIssueable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type WindowsUpdatesProductsItemMicrosoftgraphwindowsupdatesgetknownissuesbytimerangewithdaysinpastwithincludeallactiveGetKnownIssuesByTimeRangeWithDaysInPastWithIncludeAllActiveGetResponseable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnownIssueable)
    SetValue(value []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.KnownIssueable)()
}
