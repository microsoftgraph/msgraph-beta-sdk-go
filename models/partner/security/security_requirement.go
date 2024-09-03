package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type SecurityRequirement struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewSecurityRequirement instantiates a new SecurityRequirement and sets the default values.
func NewSecurityRequirement()(*SecurityRequirement) {
    m := &SecurityRequirement{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateSecurityRequirementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecurityRequirementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.partner.security.adminsMfaEnforcedSecurityRequirement":
                        return NewAdminsMfaEnforcedSecurityRequirement(), nil
                    case "#microsoft.graph.partner.security.customersMfaEnforcedSecurityRequirement":
                        return NewCustomersMfaEnforcedSecurityRequirement(), nil
                    case "#microsoft.graph.partner.security.customersSpendingBudgetSecurityRequirement":
                        return NewCustomersSpendingBudgetSecurityRequirement(), nil
                    case "#microsoft.graph.partner.security.responseTimeSecurityRequirement":
                        return NewResponseTimeSecurityRequirement(), nil
                }
            }
        }
    }
    return NewSecurityRequirement(), nil
}
// GetActionUrl gets the actionUrl property value. The link to the site where the admin can take action on the requirement.
// returns a *string when successful
func (m *SecurityRequirement) GetActionUrl()(*string) {
    val, err := m.GetBackingStore().Get("actionUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetComplianceStatus gets the complianceStatus property value. The complianceStatus property
// returns a *ComplianceStatus when successful
func (m *SecurityRequirement) GetComplianceStatus()(*ComplianceStatus) {
    val, err := m.GetBackingStore().Get("complianceStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ComplianceStatus)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecurityRequirement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionUrl(val)
        }
        return nil
    }
    res["complianceStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceStatus(val.(*ComplianceStatus))
        }
        return nil
    }
    res["helpUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHelpUrl(val)
        }
        return nil
    }
    res["maxScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxScore(val)
        }
        return nil
    }
    res["requirementType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityRequirementType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequirementType(val.(*SecurityRequirementType))
        }
        return nil
    }
    res["score"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScore(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSecurityRequirementState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*SecurityRequirementState))
        }
        return nil
    }
    res["updatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedDateTime(val)
        }
        return nil
    }
    return res
}
// GetHelpUrl gets the helpUrl property value. The link to documentation for the requirement.
// returns a *string when successful
func (m *SecurityRequirement) GetHelpUrl()(*string) {
    val, err := m.GetBackingStore().Get("helpUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetMaxScore gets the maxScore property value. The maximum score possible for the requirement.
// returns a *int64 when successful
func (m *SecurityRequirement) GetMaxScore()(*int64) {
    val, err := m.GetBackingStore().Get("maxScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetRequirementType gets the requirementType property value. The requirementType property
// returns a *SecurityRequirementType when successful
func (m *SecurityRequirement) GetRequirementType()(*SecurityRequirementType) {
    val, err := m.GetBackingStore().Get("requirementType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SecurityRequirementType)
    }
    return nil
}
// GetScore gets the score property value. The score received for this requirement.
// returns a *int64 when successful
func (m *SecurityRequirement) GetScore()(*int64) {
    val, err := m.GetBackingStore().Get("score")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetState gets the state property value. The state property
// returns a *SecurityRequirementState when successful
func (m *SecurityRequirement) GetState()(*SecurityRequirementState) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SecurityRequirementState)
    }
    return nil
}
// GetUpdatedDateTime gets the updatedDateTime property value. The date the requirement properties were last updated.
// returns a *Time when successful
func (m *SecurityRequirement) GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("updatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SecurityRequirement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("actionUrl", m.GetActionUrl())
        if err != nil {
            return err
        }
    }
    if m.GetComplianceStatus() != nil {
        cast := (*m.GetComplianceStatus()).String()
        err = writer.WriteStringValue("complianceStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("helpUrl", m.GetHelpUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("maxScore", m.GetMaxScore())
        if err != nil {
            return err
        }
    }
    if m.GetRequirementType() != nil {
        cast := (*m.GetRequirementType()).String()
        err = writer.WriteStringValue("requirementType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("score", m.GetScore())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("updatedDateTime", m.GetUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionUrl sets the actionUrl property value. The link to the site where the admin can take action on the requirement.
func (m *SecurityRequirement) SetActionUrl(value *string)() {
    err := m.GetBackingStore().Set("actionUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetComplianceStatus sets the complianceStatus property value. The complianceStatus property
func (m *SecurityRequirement) SetComplianceStatus(value *ComplianceStatus)() {
    err := m.GetBackingStore().Set("complianceStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetHelpUrl sets the helpUrl property value. The link to documentation for the requirement.
func (m *SecurityRequirement) SetHelpUrl(value *string)() {
    err := m.GetBackingStore().Set("helpUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxScore sets the maxScore property value. The maximum score possible for the requirement.
func (m *SecurityRequirement) SetMaxScore(value *int64)() {
    err := m.GetBackingStore().Set("maxScore", value)
    if err != nil {
        panic(err)
    }
}
// SetRequirementType sets the requirementType property value. The requirementType property
func (m *SecurityRequirement) SetRequirementType(value *SecurityRequirementType)() {
    err := m.GetBackingStore().Set("requirementType", value)
    if err != nil {
        panic(err)
    }
}
// SetScore sets the score property value. The score received for this requirement.
func (m *SecurityRequirement) SetScore(value *int64)() {
    err := m.GetBackingStore().Set("score", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. The state property
func (m *SecurityRequirement) SetState(value *SecurityRequirementState)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdatedDateTime sets the updatedDateTime property value. The date the requirement properties were last updated.
func (m *SecurityRequirement) SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("updatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
type SecurityRequirementable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionUrl()(*string)
    GetComplianceStatus()(*ComplianceStatus)
    GetHelpUrl()(*string)
    GetMaxScore()(*int64)
    GetRequirementType()(*SecurityRequirementType)
    GetScore()(*int64)
    GetState()(*SecurityRequirementState)
    GetUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetActionUrl(value *string)()
    SetComplianceStatus(value *ComplianceStatus)()
    SetHelpUrl(value *string)()
    SetMaxScore(value *int64)()
    SetRequirementType(value *SecurityRequirementType)()
    SetScore(value *int64)()
    SetState(value *SecurityRequirementState)()
    SetUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
