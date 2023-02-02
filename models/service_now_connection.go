package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceNowConnection 
type ServiceNowConnection struct {
    Entity
    // Indicates the method used by Intune to authenticate with ServiceNow. Currently supports only web authentication with ServiceNow using the specified app id.
    authenticationMethod ServiceNowAuthenticationMethodable
    // Date Time when connection properties were created. The value cannot be modified and is automatically populated when the connection properties were entered.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates the ServiceNow incident API URL that Intune will use the fetch incidents. Saved in the format of /api/now/table/incident
    incidentApiUrl *string
    // Indicates the ServiceNow instance URL that Intune will connect to. Saved in the format of https://<instance>.service-now.com
    instanceUrl *string
    // Date Time when connection properties were last updated. The value cannot be modified and is automatically populated when the connection properties were updated.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date Time when incidents from ServiceNow were last queried
    lastQueriedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Status of ServiceNow Connection
    serviceNowConnectionStatus *ServiceNowConnectionStatus
}
// NewServiceNowConnection instantiates a new ServiceNowConnection and sets the default values.
func NewServiceNowConnection()(*ServiceNowConnection) {
    m := &ServiceNowConnection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateServiceNowConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceNowConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceNowConnection(), nil
}
// GetAuthenticationMethod gets the authenticationMethod property value. Indicates the method used by Intune to authenticate with ServiceNow. Currently supports only web authentication with ServiceNow using the specified app id.
func (m *ServiceNowConnection) GetAuthenticationMethod()(ServiceNowAuthenticationMethodable) {
    return m.authenticationMethod
}
// GetCreatedDateTime gets the createdDateTime property value. Date Time when connection properties were created. The value cannot be modified and is automatically populated when the connection properties were entered.
func (m *ServiceNowConnection) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
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
func (m *ServiceNowConnection) GetIncidentApiUrl()(*string) {
    return m.incidentApiUrl
}
// GetInstanceUrl gets the instanceUrl property value. Indicates the ServiceNow instance URL that Intune will connect to. Saved in the format of https://<instance>.service-now.com
func (m *ServiceNowConnection) GetInstanceUrl()(*string) {
    return m.instanceUrl
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date Time when connection properties were last updated. The value cannot be modified and is automatically populated when the connection properties were updated.
func (m *ServiceNowConnection) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLastQueriedDateTime gets the lastQueriedDateTime property value. Date Time when incidents from ServiceNow were last queried
func (m *ServiceNowConnection) GetLastQueriedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastQueriedDateTime
}
// GetServiceNowConnectionStatus gets the serviceNowConnectionStatus property value. Status of ServiceNow Connection
func (m *ServiceNowConnection) GetServiceNowConnectionStatus()(*ServiceNowConnectionStatus) {
    return m.serviceNowConnectionStatus
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
    m.authenticationMethod = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date Time when connection properties were created. The value cannot be modified and is automatically populated when the connection properties were entered.
func (m *ServiceNowConnection) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetIncidentApiUrl sets the incidentApiUrl property value. Indicates the ServiceNow incident API URL that Intune will use the fetch incidents. Saved in the format of /api/now/table/incident
func (m *ServiceNowConnection) SetIncidentApiUrl(value *string)() {
    m.incidentApiUrl = value
}
// SetInstanceUrl sets the instanceUrl property value. Indicates the ServiceNow instance URL that Intune will connect to. Saved in the format of https://<instance>.service-now.com
func (m *ServiceNowConnection) SetInstanceUrl(value *string)() {
    m.instanceUrl = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date Time when connection properties were last updated. The value cannot be modified and is automatically populated when the connection properties were updated.
func (m *ServiceNowConnection) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLastQueriedDateTime sets the lastQueriedDateTime property value. Date Time when incidents from ServiceNow were last queried
func (m *ServiceNowConnection) SetLastQueriedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastQueriedDateTime = value
}
// SetServiceNowConnectionStatus sets the serviceNowConnectionStatus property value. Status of ServiceNow Connection
func (m *ServiceNowConnection) SetServiceNowConnectionStatus(value *ServiceNowConnectionStatus)() {
    m.serviceNowConnectionStatus = value
}
