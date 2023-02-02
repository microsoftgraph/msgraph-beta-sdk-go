package windowsupdates

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ComplianceChange 
type ComplianceChange struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The isRevoked property
    isRevoked *bool
    // The revokedDateTime property
    revokedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The updatePolicy property
    updatePolicy UpdatePolicyable
}
// NewComplianceChange instantiates a new complianceChange and sets the default values.
func NewComplianceChange()(*ComplianceChange) {
    m := &ComplianceChange{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateComplianceChangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComplianceChangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windowsUpdates.contentApproval":
                        return NewContentApproval(), nil
                }
            }
        }
    }
    return NewComplianceChange(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ComplianceChange) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComplianceChange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["isRevoked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRevoked(val)
        }
        return nil
    }
    res["revokedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRevokedDateTime(val)
        }
        return nil
    }
    res["updatePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUpdatePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatePolicy(val.(UpdatePolicyable))
        }
        return nil
    }
    return res
}
// GetIsRevoked gets the isRevoked property value. The isRevoked property
func (m *ComplianceChange) GetIsRevoked()(*bool) {
    return m.isRevoked
}
// GetRevokedDateTime gets the revokedDateTime property value. The revokedDateTime property
func (m *ComplianceChange) GetRevokedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.revokedDateTime
}
// GetUpdatePolicy gets the updatePolicy property value. The updatePolicy property
func (m *ComplianceChange) GetUpdatePolicy()(UpdatePolicyable) {
    return m.updatePolicy
}
// Serialize serializes information the current object
func (m *ComplianceChange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRevoked", m.GetIsRevoked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("revokedDateTime", m.GetRevokedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("updatePolicy", m.GetUpdatePolicy())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ComplianceChange) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetIsRevoked sets the isRevoked property value. The isRevoked property
func (m *ComplianceChange) SetIsRevoked(value *bool)() {
    m.isRevoked = value
}
// SetRevokedDateTime sets the revokedDateTime property value. The revokedDateTime property
func (m *ComplianceChange) SetRevokedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.revokedDateTime = value
}
// SetUpdatePolicy sets the updatePolicy property value. The updatePolicy property
func (m *ComplianceChange) SetUpdatePolicy(value UpdatePolicyable)() {
    m.updatePolicy = value
}
