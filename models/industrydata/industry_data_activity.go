package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type IndustryDataActivity struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewIndustryDataActivity instantiates a new IndustryDataActivity and sets the default values.
func NewIndustryDataActivity()(*IndustryDataActivity) {
    m := &IndustryDataActivity{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateIndustryDataActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIndustryDataActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.industryData.inboundApiFlow":
                        return NewInboundApiFlow(), nil
                    case "#microsoft.graph.industryData.inboundFileFlow":
                        return NewInboundFileFlow(), nil
                    case "#microsoft.graph.industryData.inboundFlow":
                        return NewInboundFlow(), nil
                }
            }
        }
    }
    return NewIndustryDataActivity(), nil
}
// GetDisplayName gets the displayName property value. The name of the activity. Maximum supported length is 100 characters.
// returns a *string when successful
func (m *IndustryDataActivity) GetDisplayName()(*string) {
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
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *IndustryDataActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["readinessStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseReadinessStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadinessStatus(val.(*ReadinessStatus))
        }
        return nil
    }
    return res
}
// GetReadinessStatus gets the readinessStatus property value. The readinessStatus property
// returns a *ReadinessStatus when successful
func (m *IndustryDataActivity) GetReadinessStatus()(*ReadinessStatus) {
    val, err := m.GetBackingStore().Get("readinessStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReadinessStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IndustryDataActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetReadinessStatus() != nil {
        cast := (*m.GetReadinessStatus()).String()
        err = writer.WriteStringValue("readinessStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the activity. Maximum supported length is 100 characters.
func (m *IndustryDataActivity) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetReadinessStatus sets the readinessStatus property value. The readinessStatus property
func (m *IndustryDataActivity) SetReadinessStatus(value *ReadinessStatus)() {
    err := m.GetBackingStore().Set("readinessStatus", value)
    if err != nil {
        panic(err)
    }
}
type IndustryDataActivityable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetReadinessStatus()(*ReadinessStatus)
    SetDisplayName(value *string)()
    SetReadinessStatus(value *ReadinessStatus)()
}
