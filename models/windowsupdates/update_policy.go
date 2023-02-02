package windowsupdates

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// UpdatePolicy 
type UpdatePolicy struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The audience property
    audience DeploymentAudienceable
    // The complianceChangeRules property
    complianceChangeRules []ComplianceChangeRuleable
    // The complianceChanges property
    complianceChanges []ComplianceChangeable
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The deploymentSettings property
    deploymentSettings DeploymentSettingsable
}
// NewUpdatePolicy instantiates a new updatePolicy and sets the default values.
func NewUpdatePolicy()(*UpdatePolicy) {
    m := &UpdatePolicy{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateUpdatePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdatePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdatePolicy(), nil
}
// GetAudience gets the audience property value. The audience property
func (m *UpdatePolicy) GetAudience()(DeploymentAudienceable) {
    return m.audience
}
// GetComplianceChangeRules gets the complianceChangeRules property value. The complianceChangeRules property
func (m *UpdatePolicy) GetComplianceChangeRules()([]ComplianceChangeRuleable) {
    return m.complianceChangeRules
}
// GetComplianceChanges gets the complianceChanges property value. The complianceChanges property
func (m *UpdatePolicy) GetComplianceChanges()([]ComplianceChangeable) {
    return m.complianceChanges
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *UpdatePolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeploymentSettings gets the deploymentSettings property value. The deploymentSettings property
func (m *UpdatePolicy) GetDeploymentSettings()(DeploymentSettingsable) {
    return m.deploymentSettings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdatePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["audience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeploymentAudienceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudience(val.(DeploymentAudienceable))
        }
        return nil
    }
    res["complianceChangeRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComplianceChangeRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceChangeRuleable, len(val))
            for i, v := range val {
                res[i] = v.(ComplianceChangeRuleable)
            }
            m.SetComplianceChangeRules(res)
        }
        return nil
    }
    res["complianceChanges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComplianceChangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceChangeable, len(val))
            for i, v := range val {
                res[i] = v.(ComplianceChangeable)
            }
            m.SetComplianceChanges(res)
        }
        return nil
    }
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
    res["deploymentSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeploymentSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentSettings(val.(DeploymentSettingsable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UpdatePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("audience", m.GetAudience())
        if err != nil {
            return err
        }
    }
    if m.GetComplianceChangeRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComplianceChangeRules()))
        for i, v := range m.GetComplianceChangeRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("complianceChangeRules", cast)
        if err != nil {
            return err
        }
    }
    if m.GetComplianceChanges() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComplianceChanges()))
        for i, v := range m.GetComplianceChanges() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("complianceChanges", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploymentSettings", m.GetDeploymentSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAudience sets the audience property value. The audience property
func (m *UpdatePolicy) SetAudience(value DeploymentAudienceable)() {
    m.audience = value
}
// SetComplianceChangeRules sets the complianceChangeRules property value. The complianceChangeRules property
func (m *UpdatePolicy) SetComplianceChangeRules(value []ComplianceChangeRuleable)() {
    m.complianceChangeRules = value
}
// SetComplianceChanges sets the complianceChanges property value. The complianceChanges property
func (m *UpdatePolicy) SetComplianceChanges(value []ComplianceChangeable)() {
    m.complianceChanges = value
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *UpdatePolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeploymentSettings sets the deploymentSettings property value. The deploymentSettings property
func (m *UpdatePolicy) SetDeploymentSettings(value DeploymentSettingsable)() {
    m.deploymentSettings = value
}
