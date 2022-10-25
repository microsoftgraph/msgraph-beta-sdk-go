package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceInsight provides operations to manage the collection of accessReview entities.
type GovernanceInsight struct {
    Entity
    // Indicates when the insight was created.
    insightCreatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewGovernanceInsight instantiates a new governanceInsight and sets the default values.
func NewGovernanceInsight()(*GovernanceInsight) {
    m := &GovernanceInsight{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.governanceInsight";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateGovernanceInsightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceInsightFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.membershipOutlierInsight":
                        return NewMembershipOutlierInsight(), nil
                    case "#microsoft.graph.userSignInInsight":
                        return NewUserSignInInsight(), nil
                }
            }
        }
    }
    return NewGovernanceInsight(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceInsight) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["insightCreatedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetInsightCreatedDateTime)
    return res
}
// GetInsightCreatedDateTime gets the insightCreatedDateTime property value. Indicates when the insight was created.
func (m *GovernanceInsight) GetInsightCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.insightCreatedDateTime
}
// Serialize serializes information the current object
func (m *GovernanceInsight) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("insightCreatedDateTime", m.GetInsightCreatedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInsightCreatedDateTime sets the insightCreatedDateTime property value. Indicates when the insight was created.
func (m *GovernanceInsight) SetInsightCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.insightCreatedDateTime = value
}
