package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ContentApproval 
type ContentApproval struct {
    ComplianceChange
}
// NewContentApproval instantiates a new contentApproval and sets the default values.
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
    val, err := m.GetBackingStore().Get("content")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeployableContentable)
    }
    return nil
}
// GetDeployments gets the deployments property value. Deployments created as a result of applying the approval.
func (m *ContentApproval) GetDeployments()([]Deploymentable) {
    val, err := m.GetBackingStore().Get("deployments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Deploymentable)
    }
    return nil
}
// GetDeploymentSettings gets the deploymentSettings property value. Settings for governing how to deploy content.
func (m *ContentApproval) GetDeploymentSettings()(DeploymentSettingsable) {
    val, err := m.GetBackingStore().Get("deploymentSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeploymentSettingsable)
    }
    return nil
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
                if v != nil {
                    res[i] = v.(Deploymentable)
                }
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ContentApproval) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
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
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content property
func (m *ContentApproval) SetContent(value DeployableContentable)() {
    err := m.GetBackingStore().Set("content", value)
    if err != nil {
        panic(err)
    }
}
// SetDeployments sets the deployments property value. Deployments created as a result of applying the approval.
func (m *ContentApproval) SetDeployments(value []Deploymentable)() {
    err := m.GetBackingStore().Set("deployments", value)
    if err != nil {
        panic(err)
    }
}
// SetDeploymentSettings sets the deploymentSettings property value. Settings for governing how to deploy content.
func (m *ContentApproval) SetDeploymentSettings(value DeploymentSettingsable)() {
    err := m.GetBackingStore().Set("deploymentSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ContentApproval) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// ContentApprovalable 
type ContentApprovalable interface {
    ComplianceChangeable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()(DeployableContentable)
    GetDeployments()([]Deploymentable)
    GetDeploymentSettings()(DeploymentSettingsable)
    GetOdataType()(*string)
    SetContent(value DeployableContentable)()
    SetDeployments(value []Deploymentable)()
    SetDeploymentSettings(value DeploymentSettingsable)()
    SetOdataType(value *string)()
}
