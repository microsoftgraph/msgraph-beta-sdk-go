package ediscovery

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// DataSourceContainer 
type DataSourceContainer struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewDataSourceContainer instantiates a new dataSourceContainer and sets the default values.
func NewDataSourceContainer()(*DataSourceContainer) {
    m := &DataSourceContainer{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateDataSourceContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataSourceContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.ediscovery.custodian":
                        return NewCustodian(), nil
                    case "#microsoft.graph.ediscovery.noncustodialDataSource":
                        return NewNoncustodialDataSource(), nil
                }
            }
        }
    }
    return NewDataSourceContainer(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Created date and time of the dataSourceContainer entity.
func (m *DataSourceContainer) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Display name of the dataSourceContainer entity.
func (m *DataSourceContainer) GetDisplayName()(*string) {
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
func (m *DataSourceContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["holdStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataSourceContainer_holdStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHoldStatus(val.(*DataSourceContainer_holdStatus))
        }
        return nil
    }
    res["lastIndexOperation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCaseIndexOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastIndexOperation(val.(CaseIndexOperationable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["releasedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReleasedDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataSourceContainer_status)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DataSourceContainer_status))
        }
        return nil
    }
    return res
}
// GetHoldStatus gets the holdStatus property value. The holdStatus property
func (m *DataSourceContainer) GetHoldStatus()(*DataSourceContainer_holdStatus) {
    val, err := m.GetBackingStore().Get("holdStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DataSourceContainer_holdStatus)
    }
    return nil
}
// GetLastIndexOperation gets the lastIndexOperation property value. The lastIndexOperation property
func (m *DataSourceContainer) GetLastIndexOperation()(CaseIndexOperationable) {
    val, err := m.GetBackingStore().Get("lastIndexOperation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CaseIndexOperationable)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modified date and time of the dataSourceContainer.
func (m *DataSourceContainer) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetReleasedDateTime gets the releasedDateTime property value. Date and time that the dataSourceContainer was released from the case.
func (m *DataSourceContainer) GetReleasedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("releasedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetStatus gets the status property value. Latest status of the dataSourceContainer. Possible values are: Active, Released.
func (m *DataSourceContainer) GetStatus()(*DataSourceContainer_status) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DataSourceContainer_status)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DataSourceContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetHoldStatus() != nil {
        cast := (*m.GetHoldStatus()).String()
        err = writer.WriteStringValue("holdStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastIndexOperation", m.GetLastIndexOperation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("releasedDateTime", m.GetReleasedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Created date and time of the dataSourceContainer entity.
func (m *DataSourceContainer) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Display name of the dataSourceContainer entity.
func (m *DataSourceContainer) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetHoldStatus sets the holdStatus property value. The holdStatus property
func (m *DataSourceContainer) SetHoldStatus(value *DataSourceContainer_holdStatus)() {
    err := m.GetBackingStore().Set("holdStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetLastIndexOperation sets the lastIndexOperation property value. The lastIndexOperation property
func (m *DataSourceContainer) SetLastIndexOperation(value CaseIndexOperationable)() {
    err := m.GetBackingStore().Set("lastIndexOperation", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modified date and time of the dataSourceContainer.
func (m *DataSourceContainer) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetReleasedDateTime sets the releasedDateTime property value. Date and time that the dataSourceContainer was released from the case.
func (m *DataSourceContainer) SetReleasedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("releasedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Latest status of the dataSourceContainer. Possible values are: Active, Released.
func (m *DataSourceContainer) SetStatus(value *DataSourceContainer_status)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// DataSourceContainerable 
type DataSourceContainerable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDisplayName()(*string)
    GetHoldStatus()(*DataSourceContainer_holdStatus)
    GetLastIndexOperation()(CaseIndexOperationable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReleasedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetStatus()(*DataSourceContainer_status)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDisplayName(value *string)()
    SetHoldStatus(value *DataSourceContainer_holdStatus)()
    SetLastIndexOperation(value CaseIndexOperationable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReleasedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetStatus(value *DataSourceContainer_status)()
}
