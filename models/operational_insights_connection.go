package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OperationalInsightsConnection 
type OperationalInsightsConnection struct {
    ResourceConnection
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The azureResourceGroupName property
    azureResourceGroupName *string
    // The azureSubscriptionId property
    azureSubscriptionId *string
    // The workspaceName property
    workspaceName *string
}
// NewOperationalInsightsConnection instantiates a new OperationalInsightsConnection and sets the default values.
func NewOperationalInsightsConnection()(*OperationalInsightsConnection) {
    m := &OperationalInsightsConnection{
        ResourceConnection: *NewResourceConnection(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOperationalInsightsConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOperationalInsightsConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOperationalInsightsConnection(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OperationalInsightsConnection) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAzureResourceGroupName gets the azureResourceGroupName property value. The azureResourceGroupName property
func (m *OperationalInsightsConnection) GetAzureResourceGroupName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureResourceGroupName
    }
}
// GetAzureSubscriptionId gets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *OperationalInsightsConnection) GetAzureSubscriptionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureSubscriptionId
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// GetWorkspaceName gets the workspaceName property value. The workspaceName property
func (m *OperationalInsightsConnection) GetWorkspaceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.workspaceName
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OperationalInsightsConnection) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAzureResourceGroupName sets the azureResourceGroupName property value. The azureResourceGroupName property
func (m *OperationalInsightsConnection) SetAzureResourceGroupName(value *string)() {
    if m != nil {
        m.azureResourceGroupName = value
    }
}
// SetAzureSubscriptionId sets the azureSubscriptionId property value. The azureSubscriptionId property
func (m *OperationalInsightsConnection) SetAzureSubscriptionId(value *string)() {
    if m != nil {
        m.azureSubscriptionId = value
    }
}
// SetWorkspaceName sets the workspaceName property value. The workspaceName property
func (m *OperationalInsightsConnection) SetWorkspaceName(value *string)() {
    if m != nil {
        m.workspaceName = value
    }
}
