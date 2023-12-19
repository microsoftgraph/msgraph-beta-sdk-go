package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcBulkAction 
type CloudPcBulkAction struct {
    Entity
}
// NewCloudPcBulkAction instantiates a new cloudPcBulkAction and sets the default values.
func NewCloudPcBulkAction()(*CloudPcBulkAction) {
    m := &CloudPcBulkAction{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcBulkActionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcBulkActionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.cloudPcBulkModifyDiskEncryptionType":
                        return NewCloudPcBulkModifyDiskEncryptionType(), nil
                    case "#microsoft.graph.cloudPcBulkPowerOff":
                        return NewCloudPcBulkPowerOff(), nil
                    case "#microsoft.graph.cloudPcBulkPowerOn":
                        return NewCloudPcBulkPowerOn(), nil
                    case "#microsoft.graph.cloudPcBulkReprovision":
                        return NewCloudPcBulkReprovision(), nil
                    case "#microsoft.graph.cloudPcBulkResize":
                        return NewCloudPcBulkResize(), nil
                    case "#microsoft.graph.cloudPcBulkRestart":
                        return NewCloudPcBulkRestart(), nil
                    case "#microsoft.graph.cloudPcBulkRestore":
                        return NewCloudPcBulkRestore(), nil
                    case "#microsoft.graph.cloudPcBulkTroubleshoot":
                        return NewCloudPcBulkTroubleshoot(), nil
                }
            }
        }
    }
    return NewCloudPcBulkAction(), nil
}
// GetActionSummary gets the actionSummary property value. The actionSummary property
func (m *CloudPcBulkAction) GetActionSummary()(CloudPcBulkActionSummaryable) {
    val, err := m.GetBackingStore().Get("actionSummary")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CloudPcBulkActionSummaryable)
    }
    return nil
}
// GetCloudPcIds gets the cloudPcIds property value. The cloudPcIds property
func (m *CloudPcBulkAction) GetCloudPcIds()([]string) {
    val, err := m.GetBackingStore().Get("cloudPcIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *CloudPcBulkAction) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *CloudPcBulkAction) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcBulkAction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCloudPcBulkActionSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionSummary(val.(CloudPcBulkActionSummaryable))
        }
        return nil
    }
    res["cloudPcIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetCloudPcIds(res)
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CloudPcBulkAction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("actionSummary", m.GetActionSummary())
        if err != nil {
            return err
        }
    }
    if m.GetCloudPcIds() != nil {
        err = writer.WriteCollectionOfStringValues("cloudPcIds", m.GetCloudPcIds())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionSummary sets the actionSummary property value. The actionSummary property
func (m *CloudPcBulkAction) SetActionSummary(value CloudPcBulkActionSummaryable)() {
    err := m.GetBackingStore().Set("actionSummary", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudPcIds sets the cloudPcIds property value. The cloudPcIds property
func (m *CloudPcBulkAction) SetCloudPcIds(value []string)() {
    err := m.GetBackingStore().Set("cloudPcIds", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *CloudPcBulkAction) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CloudPcBulkAction) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcBulkActionable 
type CloudPcBulkActionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionSummary()(CloudPcBulkActionSummaryable)
    GetCloudPcIds()([]string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    SetActionSummary(value CloudPcBulkActionSummaryable)()
    SetCloudPcIds(value []string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
}
