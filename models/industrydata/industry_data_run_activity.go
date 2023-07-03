package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// IndustryDataRunActivity 
type IndustryDataRunActivity struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewIndustryDataRunActivity instantiates a new IndustryDataRunActivity and sets the default values.
func NewIndustryDataRunActivity()(*IndustryDataRunActivity) {
    m := &IndustryDataRunActivity{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateIndustryDataRunActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIndustryDataRunActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.industryData.inboundFlowActivity":
                        return NewInboundFlowActivity(), nil
                    case "#microsoft.graph.industryData.outboundFlowActivity":
                        return NewOutboundFlowActivity(), nil
                }
            }
        }
    }
    return NewIndustryDataRunActivity(), nil
}
// GetActivity gets the activity property value. The flow that was run by this activity.
func (m *IndustryDataRunActivity) GetActivity()(IndustryDataActivityable) {
    val, err := m.GetBackingStore().Get("activity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IndustryDataActivityable)
    }
    return nil
}
// GetBlockingError gets the blockingError property value. An error object to diagnose critical failures in an activity.
func (m *IndustryDataRunActivity) GetBlockingError()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicErrorable) {
    val, err := m.GetBackingStore().Get("blockingError")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicErrorable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The name of the running flow.
func (m *IndustryDataRunActivity) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IndustryDataRunActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIndustryDataActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val.(IndustryDataActivityable))
        }
        return nil
    }
    res["blockingError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockingError(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicErrorable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIndustryDataActivityStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*IndustryDataActivityStatus))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. The status property
func (m *IndustryDataRunActivity) GetStatus()(*IndustryDataActivityStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*IndustryDataActivityStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IndustryDataRunActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. The flow that was run by this activity.
func (m *IndustryDataRunActivity) SetActivity(value IndustryDataActivityable)() {
    err := m.GetBackingStore().Set("activity", value)
    if err != nil {
        panic(err)
    }
}
// SetBlockingError sets the blockingError property value. An error object to diagnose critical failures in an activity.
func (m *IndustryDataRunActivity) SetBlockingError(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicErrorable)() {
    err := m.GetBackingStore().Set("blockingError", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name of the running flow.
func (m *IndustryDataRunActivity) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *IndustryDataRunActivity) SetStatus(value *IndustryDataActivityStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// IndustryDataRunActivityable 
type IndustryDataRunActivityable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivity()(IndustryDataActivityable)
    GetBlockingError()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicErrorable)
    GetDisplayName()(*string)
    GetStatus()(*IndustryDataActivityStatus)
    SetActivity(value IndustryDataActivityable)()
    SetBlockingError(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PublicErrorable)()
    SetDisplayName(value *string)()
    SetStatus(value *IndustryDataActivityStatus)()
}
