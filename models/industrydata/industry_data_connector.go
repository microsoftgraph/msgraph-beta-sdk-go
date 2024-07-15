package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type IndustryDataConnector struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewIndustryDataConnector instantiates a new IndustryDataConnector and sets the default values.
func NewIndustryDataConnector()(*IndustryDataConnector) {
    m := &IndustryDataConnector{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateIndustryDataConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIndustryDataConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.industryData.apiDataConnector":
                        return NewApiDataConnector(), nil
                    case "#microsoft.graph.industryData.azureDataLakeConnector":
                        return NewAzureDataLakeConnector(), nil
                    case "#microsoft.graph.industryData.fileDataConnector":
                        return NewFileDataConnector(), nil
                    case "#microsoft.graph.industryData.oneRosterApiDataConnector":
                        return NewOneRosterApiDataConnector(), nil
                }
            }
        }
    }
    return NewIndustryDataConnector(), nil
}
// GetDisplayName gets the displayName property value. The name of the data connector. Maximum supported length is 100 characters.
// returns a *string when successful
func (m *IndustryDataConnector) GetDisplayName()(*string) {
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
func (m *IndustryDataConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["sourceSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSourceSystemDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceSystem(val.(SourceSystemDefinitionable))
        }
        return nil
    }
    return res
}
// GetSourceSystem gets the sourceSystem property value. The sourceSystem property
// returns a SourceSystemDefinitionable when successful
func (m *IndustryDataConnector) GetSourceSystem()(SourceSystemDefinitionable) {
    val, err := m.GetBackingStore().Get("sourceSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SourceSystemDefinitionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *IndustryDataConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteObjectValue("sourceSystem", m.GetSourceSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the data connector. Maximum supported length is 100 characters.
func (m *IndustryDataConnector) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceSystem sets the sourceSystem property value. The sourceSystem property
func (m *IndustryDataConnector) SetSourceSystem(value SourceSystemDefinitionable)() {
    err := m.GetBackingStore().Set("sourceSystem", value)
    if err != nil {
        panic(err)
    }
}
type IndustryDataConnectorable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetSourceSystem()(SourceSystemDefinitionable)
    SetDisplayName(value *string)()
    SetSourceSystem(value SourceSystemDefinitionable)()
}
