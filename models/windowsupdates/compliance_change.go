package windowsupdates

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type ComplianceChange struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewComplianceChange instantiates a new ComplianceChange and sets the default values.
func NewComplianceChange()(*ComplianceChange) {
    m := &ComplianceChange{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateComplianceChangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
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
// GetCreatedDateTime gets the createdDateTime property value. The date and time when a compliance change was created.
// returns a *Time when successful
func (m *ComplianceChange) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
// GetIsRevoked gets the isRevoked property value. True indicates that a compliance change is revoked, preventing further application. Revoking a compliance change is a final action.
// returns a *bool when successful
func (m *ComplianceChange) GetIsRevoked()(*bool) {
    val, err := m.GetBackingStore().Get("isRevoked")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRevokedDateTime gets the revokedDateTime property value. The date and time when the compliance change was revoked.
// returns a *Time when successful
func (m *ComplianceChange) GetRevokedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("revokedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetUpdatePolicy gets the updatePolicy property value. The policy this compliance change is a member of.
// returns a UpdatePolicyable when successful
func (m *ComplianceChange) GetUpdatePolicy()(UpdatePolicyable) {
    val, err := m.GetBackingStore().Get("updatePolicy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UpdatePolicyable)
    }
    return nil
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
// SetCreatedDateTime sets the createdDateTime property value. The date and time when a compliance change was created.
func (m *ComplianceChange) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIsRevoked sets the isRevoked property value. True indicates that a compliance change is revoked, preventing further application. Revoking a compliance change is a final action.
func (m *ComplianceChange) SetIsRevoked(value *bool)() {
    err := m.GetBackingStore().Set("isRevoked", value)
    if err != nil {
        panic(err)
    }
}
// SetRevokedDateTime sets the revokedDateTime property value. The date and time when the compliance change was revoked.
func (m *ComplianceChange) SetRevokedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("revokedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdatePolicy sets the updatePolicy property value. The policy this compliance change is a member of.
func (m *ComplianceChange) SetUpdatePolicy(value UpdatePolicyable)() {
    err := m.GetBackingStore().Set("updatePolicy", value)
    if err != nil {
        panic(err)
    }
}
type ComplianceChangeable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsRevoked()(*bool)
    GetRevokedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUpdatePolicy()(UpdatePolicyable)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsRevoked(value *bool)()
    SetRevokedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUpdatePolicy(value UpdatePolicyable)()
}
