package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OperationalInsightsConnection struct {
    ResourceConnection
}
// NewOperationalInsightsConnection instantiates a new OperationalInsightsConnection and sets the default values.
func NewOperationalInsightsConnection()(*OperationalInsightsConnection) {
    m := &OperationalInsightsConnection{
        ResourceConnection: *NewResourceConnection(),
    }
    odataTypeValue := "#microsoft.graph.windowsUpdates.operationalInsightsConnection"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOperationalInsightsConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOperationalInsightsConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationalInsightsConnection(), nil
}
// GetAzureResourceGroupName gets the azureResourceGroupName property value. The name of the Azure resource group that contains the Log Analytics workspace.
// returns a *string when successful
func (m *OperationalInsightsConnection) GetAzureResourceGroupName()(*string) {
    val, err := m.GetBackingStore().Get("azureResourceGroupName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. The Azure subscription ID that contains the Log Analytics workspace.
// returns a *string when successful
func (m *OperationalInsightsConnection) GetAzureSubscriptionId()(*string) {
    val, err := m.GetBackingStore().Get("azureSubscriptionId")
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
func (m *OperationalInsightsConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ResourceConnection.GetFieldDeserializers()
    res["azureResourceGroupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureResourceGroupName(val)
        }
        return nil
    }
    res["azureSubscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureSubscriptionId(val)
        }
        return nil
    }
    res["workspaceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkspaceName(val)
        }
        return nil
    }
    return res
}
// GetWorkspaceName gets the workspaceName property value. The name of the Log Analytics workspace.
// returns a *string when successful
func (m *OperationalInsightsConnection) GetWorkspaceName()(*string) {
    val, err := m.GetBackingStore().Get("workspaceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    err := m.GetBackingStore().Set("azureResourceGroupName", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureSubscriptionId sets the azureSubscriptionId property value. The Azure subscription ID that contains the Log Analytics workspace.
func (m *OperationalInsightsConnection) SetAzureSubscriptionId(value *string)() {
    err := m.GetBackingStore().Set("azureSubscriptionId", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkspaceName sets the workspaceName property value. The name of the Log Analytics workspace.
func (m *OperationalInsightsConnection) SetWorkspaceName(value *string)() {
    err := m.GetBackingStore().Set("workspaceName", value)
    if err != nil {
        panic(err)
    }
}
type OperationalInsightsConnectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ResourceConnectionable
    GetAzureResourceGroupName()(*string)
    GetAzureSubscriptionId()(*string)
    GetWorkspaceName()(*string)
    SetAzureResourceGroupName(value *string)()
    SetAzureSubscriptionId(value *string)()
    SetWorkspaceName(value *string)()
}
