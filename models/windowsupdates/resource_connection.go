package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ResourceConnection 
type ResourceConnection struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The state property
    state *ResourceConnectionState
}
// NewResourceConnection instantiates a new ResourceConnection and sets the default values.
func NewResourceConnection()(*ResourceConnection) {
    m := &ResourceConnection{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    typeValue := "#microsoft.graph.windowsUpdates.resourceConnection";
    m.SetType(&typeValue);
    return m
}
// CreateResourceConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResourceConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.windowsUpdates.operationalInsightsConnection":
                        return NewOperationalInsightsConnection(), nil
                }
            }
        }
    }
    return NewResourceConnection(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseResourceConnectionState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ResourceConnectionState))
        }
        return nil
    }
    return res
}
// GetState gets the state property value. The state property
func (m *ResourceConnection) GetState()(*ResourceConnectionState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Serialize serializes information the current object
func (m *ResourceConnection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetState sets the state property value. The state property
func (m *ResourceConnection) SetState(value *ResourceConnectionState)() {
    if m != nil {
        m.state = value
    }
}
