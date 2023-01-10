package devicemanagement

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody 
type WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody struct {
    // An enum type to represent approval actions of single or list of drivers.
    actionName *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriverApprovalAction
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The deploymentDate property
    deploymentDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The driverIds property
    driverIds []string
}
// NewWindowsDriverUpdateProfilesItemExecuteActionPostRequestBody instantiates a new WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody and sets the default values.
func NewWindowsDriverUpdateProfilesItemExecuteActionPostRequestBody()(*WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody) {
    m := &WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindowsDriverUpdateProfilesItemExecuteActionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDriverUpdateProfilesItemExecuteActionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDriverUpdateProfilesItemExecuteActionPostRequestBody(), nil
}
// GetActionName gets the actionName property value. An enum type to represent approval actions of single or list of drivers.
func (m *WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody) GetActionName()(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriverApprovalAction) {
    return m.actionName
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDeploymentDate gets the deploymentDate property value. The deploymentDate property
func (m *WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody) GetDeploymentDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.deploymentDate
}
// GetDriverIds gets the driverIds property value. The driverIds property
func (m *WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody) GetDriverIds()([]string) {
    return m.driverIds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["actionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ParseDriverApprovalAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionName(val.(*ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriverApprovalAction))
        }
        return nil
    }
    res["deploymentDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentDate(val)
        }
        return nil
    }
    res["driverIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDriverIds(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetActionName() != nil {
        cast := (*m.GetActionName()).String()
        err := writer.WriteStringValue("actionName", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("deploymentDate", m.GetDeploymentDate())
        if err != nil {
            return err
        }
    }
    if m.GetDriverIds() != nil {
        err := writer.WriteCollectionOfStringValues("driverIds", m.GetDriverIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionName sets the actionName property value. An enum type to represent approval actions of single or list of drivers.
func (m *WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody) SetActionName(value *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriverApprovalAction)() {
    m.actionName = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDeploymentDate sets the deploymentDate property value. The deploymentDate property
func (m *WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody) SetDeploymentDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deploymentDate = value
}
// SetDriverIds sets the driverIds property value. The driverIds property
func (m *WindowsDriverUpdateProfilesItemExecuteActionPostRequestBody) SetDriverIds(value []string)() {
    m.driverIds = value
}
