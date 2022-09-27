package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OperationalInsightsConnection 
type OperationalInsightsConnection struct {
    ResourceConnection
    // The name of the Azure resource group that contains the Log Analytics workspace.
    azureResourceGroupName *string
    // The Azure subscription ID that contains the Log Analytics workspace.
    azureSubscriptionId *string
    // The name of the Log Analytics workspace.
    workspaceName *string
}
// NewOperationalInsightsConnection instantiates a new OperationalInsightsConnection and sets the default values.
func NewOperationalInsightsConnection()(*OperationalInsightsConnection) {
    m := &OperationalInsightsConnection{
        ResourceConnection: *NewResourceConnection(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.operationalInsightsConnection";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOperationalInsightsConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOperationalInsightsConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationalInsightsConnection(), nil
}
// GetAzureResourceGroupName gets the azureResourceGroupName property value. The name of the Azure resource group that contains the Log Analytics workspace.
func (m *OperationalInsightsConnection) GetAzureResourceGroupName()(*string) {
    return m.azureResourceGroupName
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. The Azure subscription ID that contains the Log Analytics workspace.
func (m *OperationalInsightsConnection) GetAzureSubscriptionId()(*string) {
    return m.azureSubscriptionId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OperationalInsightsConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ResourceConnection.GetFieldDeserializers()
    res["azureResourceGroupName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureResourceGroupName)
    res["azureSubscriptionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureSubscriptionId)
    res["workspaceName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetWorkspaceName)
    return res
}
// GetWorkspaceName gets the workspaceName property value. The name of the Log Analytics workspace.
func (m *OperationalInsightsConnection) GetWorkspaceName()(*string) {
    return m.workspaceName
}
// Serialize serializes information the current object
func (m *OperationalInsightsConnection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ResourceConnection.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureResourceGroupName", m.GetAzureResourceGroupName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureSubscriptionId", m.GetAzureSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("workspaceName", m.GetWorkspaceName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureResourceGroupName sets the azureResourceGroupName property value. The name of the Azure resource group that contains the Log Analytics workspace.
func (m *OperationalInsightsConnection) SetAzureResourceGroupName(value *string)() {
    m.azureResourceGroupName = value
}
// SetAzureSubscriptionId sets the azureSubscriptionId property value. The Azure subscription ID that contains the Log Analytics workspace.
func (m *OperationalInsightsConnection) SetAzureSubscriptionId(value *string)() {
    m.azureSubscriptionId = value
}
// SetWorkspaceName sets the workspaceName property value. The name of the Log Analytics workspace.
func (m *OperationalInsightsConnection) SetWorkspaceName(value *string)() {
    m.workspaceName = value
}
