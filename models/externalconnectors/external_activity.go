package externalconnectors

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ExternalActivity provides operations to manage the collection of activityStatistics entities.
type ExternalActivity struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Represents an identity used to identify who is responsible for the activity.
    performedBy Identityable
    // When the particular activity occurred.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type property
    type_escaped *ExternalActivityType
}
// NewExternalActivity instantiates a new externalActivity and sets the default values.
func NewExternalActivity()(*ExternalActivity) {
    m := &ExternalActivity{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.externalConnectors.externalActivity";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateExternalActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.externalConnectors.externalActivityResult":
                        return NewExternalActivityResult(), nil
                }
            }
        }
    }
    return NewExternalActivity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["performedBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateIdentityFromDiscriminatorValue , m.SetPerformedBy)
    res["startDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartDateTime)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseExternalActivityType , m.SetType)
    return res
}
// GetPerformedBy gets the performedBy property value. Represents an identity used to identify who is responsible for the activity.
func (m *ExternalActivity) GetPerformedBy()(Identityable) {
    return m.performedBy
}
// GetStartDateTime gets the startDateTime property value. When the particular activity occurred.
func (m *ExternalActivity) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetType gets the type property value. The type property
func (m *ExternalActivity) GetType()(*ExternalActivityType) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *ExternalActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("performedBy", m.GetPerformedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPerformedBy sets the performedBy property value. Represents an identity used to identify who is responsible for the activity.
func (m *ExternalActivity) SetPerformedBy(value Identityable)() {
    m.performedBy = value
}
// SetStartDateTime sets the startDateTime property value. When the particular activity occurred.
func (m *ExternalActivity) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetType sets the type property value. The type property
func (m *ExternalActivity) SetType(value *ExternalActivityType)() {
    m.type_escaped = value
}
