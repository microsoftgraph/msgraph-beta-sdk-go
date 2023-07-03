package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CertificateConnectorDetails 
type CertificateConnectorDetails struct {
    Entity
}
// NewCertificateConnectorDetails instantiates a new CertificateConnectorDetails and sets the default values.
func NewCertificateConnectorDetails()(*CertificateConnectorDetails) {
    m := &CertificateConnectorDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCertificateConnectorDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateConnectorDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateConnectorDetails(), nil
}
// GetConnectorName gets the connectorName property value. Connector name (set during enrollment).
func (m *CertificateConnectorDetails) GetConnectorName()(*string) {
    val, err := m.GetBackingStore().Get("connectorName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConnectorVersion gets the connectorVersion property value. Version of the connector installed.
func (m *CertificateConnectorDetails) GetConnectorVersion()(*string) {
    val, err := m.GetBackingStore().Get("connectorVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnrollmentDateTime gets the enrollmentDateTime property value. Date/time when this connector was enrolled.
func (m *CertificateConnectorDetails) GetEnrollmentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("enrollmentDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CertificateConnectorDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connectorName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorName(val)
        }
        return nil
    }
    res["connectorVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorVersion(val)
        }
        return nil
    }
    res["enrollmentDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentDateTime(val)
        }
        return nil
    }
    res["lastCheckinDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckinDateTime(val)
        }
        return nil
    }
    res["machineName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMachineName(val)
        }
        return nil
    }
    return res
}
// GetLastCheckinDateTime gets the lastCheckinDateTime property value. Date/time when this connector last connected to the service.
func (m *CertificateConnectorDetails) GetLastCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastCheckinDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMachineName gets the machineName property value. Name of the machine hosting this connector service.
func (m *CertificateConnectorDetails) GetMachineName()(*string) {
    val, err := m.GetBackingStore().Get("machineName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CertificateConnectorDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("connectorName", m.GetConnectorName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("connectorVersion", m.GetConnectorVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("enrollmentDateTime", m.GetEnrollmentDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastCheckinDateTime", m.GetLastCheckinDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("machineName", m.GetMachineName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnectorName sets the connectorName property value. Connector name (set during enrollment).
func (m *CertificateConnectorDetails) SetConnectorName(value *string)() {
    err := m.GetBackingStore().Set("connectorName", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectorVersion sets the connectorVersion property value. Version of the connector installed.
func (m *CertificateConnectorDetails) SetConnectorVersion(value *string)() {
    err := m.GetBackingStore().Set("connectorVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentDateTime sets the enrollmentDateTime property value. Date/time when this connector was enrolled.
func (m *CertificateConnectorDetails) SetEnrollmentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("enrollmentDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastCheckinDateTime sets the lastCheckinDateTime property value. Date/time when this connector last connected to the service.
func (m *CertificateConnectorDetails) SetLastCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastCheckinDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMachineName sets the machineName property value. Name of the machine hosting this connector service.
func (m *CertificateConnectorDetails) SetMachineName(value *string)() {
    err := m.GetBackingStore().Set("machineName", value)
    if err != nil {
        panic(err)
    }
}
// CertificateConnectorDetailsable 
type CertificateConnectorDetailsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectorName()(*string)
    GetConnectorVersion()(*string)
    GetEnrollmentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMachineName()(*string)
    SetConnectorName(value *string)()
    SetConnectorVersion(value *string)()
    SetEnrollmentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMachineName(value *string)()
}
