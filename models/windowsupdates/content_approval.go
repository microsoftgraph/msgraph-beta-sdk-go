package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentApproval 
type ContentApproval struct {
    ComplianceChange
    // The content property
    content DeployableContentable
    // The deployments property
    deployments []Deploymentable
    // The deploymentSettings property
    deploymentSettings DeploymentSettingsable
}
// NewContentApproval instantiates a new ContentApproval and sets the default values.
func NewContentApproval()(*ContentApproval) {
    m := &ContentApproval{
        ComplianceChange: *NewComplianceChange(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.contentApproval"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateContentApprovalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateContentApprovalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewContentApproval(), nil
}
// GetContent gets the content property value. The content property
func (m *ContentApproval) GetContent()(DeployableContentable) {
    return m.content
}
// GetDeployments gets the deployments property value. The deployments property
func (m *ContentApproval) GetDeployments()([]Deploymentable) {
    return m.deployments
}
// GetDeploymentSettings gets the deploymentSettings property value. The deploymentSettings property
func (m *ContentApproval) GetDeploymentSettings()(DeploymentSettingsable) {
    return m.deploymentSettings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ContentApproval) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ComplianceChange.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeployableContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val.(DeployableContentable))
        }
        return nil
    }
    res["deployments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeploymentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Deploymentable, len(val))
            for i, v := range val {
                res[i] = v.(Deploymentable)
            }
            m.SetDeployments(res)
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
func (m *ContentApproval) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ComplianceChange.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    if m.GetDeployments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeployments()))
        for i, v := range m.GetDeployments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deployments", cast)
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
// SetContent sets the content property value. The content property
func (m *ContentApproval) SetContent(value DeployableContentable)() {
    m.content = value
}
// SetDeployments sets the deployments property value. The deployments property
func (m *ContentApproval) SetDeployments(value []Deploymentable)() {
    m.deployments = value
}
// SetDeploymentSettings sets the deploymentSettings property value. The deploymentSettings property
func (m *ContentApproval) SetDeploymentSettings(value DeploymentSettingsable)() {
    m.deploymentSettings = value
}
