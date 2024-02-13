package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceNowConnection serviceNow properties including the ServiceNow instanceUrl, connection credentials and other metadata.
type ServiceNowConnection struct {
    Entity
}
// NewServiceNowConnection instantiates a new ServiceNowConnection and sets the default values.
func NewServiceNowConnection()(*ServiceNowConnection) {
    m := &ServiceNowConnection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateServiceNowConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceNowConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceNowConnection(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Indicates the method used by Intune to authenticate with ServiceNow. Currently supports only web authentication with ServiceNow using the specified app id.
// returns a ServiceNowAuthenticationMethodable when successful
func (m *ServiceNowConnection) GetAuthenticationMethod()(ServiceNowAuthenticationMethodable) {
    val, err := m.GetBackingStore().Get("authenticationMethod")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ServiceNowAuthenticationMethodable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Date Time when connection properties were created. The value cannot be modified and is automatically populated when the connection properties were entered.
// returns a *Time when successful
func (m *ServiceNowConnection) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServiceNowConnection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceNowAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethod(val.(ServiceNowAuthenticationMethodable))
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
    res["incidentApiUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncidentApiUrl(val)
        }
        return nil
    }
    res["instanceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstanceUrl(val)
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
    res["lastQueriedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastQueriedDateTime(val)
        }
        return nil
    }
    res["serviceNowConnectionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceNowConnectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceNowConnectionStatus(val.(*ServiceNowConnectionStatus))
        }
        return nil
    }
    return res
}
// GetIncidentApiUrl gets the incidentApiUrl property value. Indicates the ServiceNow incident API URL that Intune will use the fetch incidents. Saved in the format of /api/now/table/incident
// returns a *string when successful
func (m *ServiceNowConnection) GetIncidentApiUrl()(*string) {
    val, err := m.GetBackingStore().Get("incidentApiUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetInstanceUrl gets the instanceUrl property value. Indicates the ServiceNow instance URL that Intune will connect to. Saved in the format of https://.service-now.com
// returns a *string when successful
func (m *ServiceNowConnection) GetInstanceUrl()(*string) {
    val, err := m.GetBackingStore().Get("instanceUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date Time when connection properties were last updated. The value cannot be modified and is automatically populated when the connection properties were updated.
// returns a *Time when successful
func (m *ServiceNowConnection) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastQueriedDateTime gets the lastQueriedDateTime property value. Date Time when incidents from ServiceNow were last queried
// returns a *Time when successful
func (m *ServiceNowConnection) GetLastQueriedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastQueriedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetServiceNowConnectionStatus gets the serviceNowConnectionStatus property value. Status of ServiceNow Connection
// returns a *ServiceNowConnectionStatus when successful
func (m *ServiceNowConnection) GetServiceNowConnectionStatus()(*ServiceNowConnectionStatus) {
    val, err := m.GetBackingStore().Get("serviceNowConnectionStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ServiceNowConnectionStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ServiceNowConnection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("authenticationMethod", m.GetAuthenticationMethod())
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
        err = writer.WriteStringValue("incidentApiUrl", m.GetIncidentApiUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("instanceUrl", m.GetInstanceUrl())
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
        err = writer.WriteTimeValue("lastQueriedDateTime", m.GetLastQueriedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetServiceNowConnectionStatus() != nil {
        cast := (*m.GetServiceNowConnectionStatus()).String()
        err = writer.WriteStringValue("serviceNowConnectionStatus", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationMethod sets the authenticationMethod property value. Indicates the method used by Intune to authenticate with ServiceNow. Currently supports only web authentication with ServiceNow using the specified app id.
func (m *ServiceNowConnection) SetAuthenticationMethod(value ServiceNowAuthenticationMethodable)() {
    err := m.GetBackingStore().Set("authenticationMethod", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Date Time when connection properties were created. The value cannot be modified and is automatically populated when the connection properties were entered.
func (m *ServiceNowConnection) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIncidentApiUrl sets the incidentApiUrl property value. Indicates the ServiceNow incident API URL that Intune will use the fetch incidents. Saved in the format of /api/now/table/incident
func (m *ServiceNowConnection) SetIncidentApiUrl(value *string)() {
    err := m.GetBackingStore().Set("incidentApiUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetInstanceUrl sets the instanceUrl property value. Indicates the ServiceNow instance URL that Intune will connect to. Saved in the format of https://.service-now.com
func (m *ServiceNowConnection) SetInstanceUrl(value *string)() {
    err := m.GetBackingStore().Set("instanceUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date Time when connection properties were last updated. The value cannot be modified and is automatically populated when the connection properties were updated.
func (m *ServiceNowConnection) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastQueriedDateTime sets the lastQueriedDateTime property value. Date Time when incidents from ServiceNow were last queried
func (m *ServiceNowConnection) SetLastQueriedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastQueriedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetServiceNowConnectionStatus sets the serviceNowConnectionStatus property value. Status of ServiceNow Connection
func (m *ServiceNowConnection) SetServiceNowConnectionStatus(value *ServiceNowConnectionStatus)() {
    err := m.GetBackingStore().Set("serviceNowConnectionStatus", value)
    if err != nil {
        panic(err)
    }
}
type ServiceNowConnectionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(ServiceNowAuthenticationMethodable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIncidentApiUrl()(*string)
    GetInstanceUrl()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastQueriedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetServiceNowConnectionStatus()(*ServiceNowConnectionStatus)
    SetAuthenticationMethod(value ServiceNowAuthenticationMethodable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIncidentApiUrl(value *string)()
    SetInstanceUrl(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastQueriedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetServiceNowConnectionStatus(value *ServiceNowConnectionStatus)()
}
