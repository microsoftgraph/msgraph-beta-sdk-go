package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CertificateConnectorDetails struct {
    Entity
    // Connector name (set during enrollment).
    connectorName *string;
    // Date/time when this connector was enrolled.
    enrollmentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Date/time when this connector last connected to the service.
    lastCheckinDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Name of the machine hosting this connector service.
    machineName *string;
}
// Instantiates a new certificateConnectorDetails and sets the default values.
func NewCertificateConnectorDetails()(*CertificateConnectorDetails) {
    m := &CertificateConnectorDetails{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the connectorName property value. Connector name (set during enrollment).
func (m *CertificateConnectorDetails) GetConnectorName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectorName
    }
}
// Gets the enrollmentDateTime property value. Date/time when this connector was enrolled.
func (m *CertificateConnectorDetails) GetEnrollmentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentDateTime
    }
}
// Gets the lastCheckinDateTime property value. Date/time when this connector last connected to the service.
func (m *CertificateConnectorDetails) GetLastCheckinDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastCheckinDateTime
    }
}
// Gets the machineName property value. Name of the machine hosting this connector service.
func (m *CertificateConnectorDetails) GetMachineName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.machineName
    }
}
// The deserialization information for the current model
func (m *CertificateConnectorDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connectorName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorName(val)
        }
        return nil
    }
    res["enrollmentDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentDateTime(val)
        }
        return nil
    }
    res["lastCheckinDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCheckinDateTime(val)
        }
        return nil
    }
    res["machineName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *CertificateConnectorDetails) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CertificateConnectorDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the connectorName property value. Connector name (set during enrollment).
// Parameters:
//  - value : Value to set for the connectorName property.
func (m *CertificateConnectorDetails) SetConnectorName(value *string)() {
    m.connectorName = value
}
// Sets the enrollmentDateTime property value. Date/time when this connector was enrolled.
// Parameters:
//  - value : Value to set for the enrollmentDateTime property.
func (m *CertificateConnectorDetails) SetEnrollmentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.enrollmentDateTime = value
}
// Sets the lastCheckinDateTime property value. Date/time when this connector last connected to the service.
// Parameters:
//  - value : Value to set for the lastCheckinDateTime property.
func (m *CertificateConnectorDetails) SetLastCheckinDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCheckinDateTime = value
}
// Sets the machineName property value. Name of the machine hosting this connector service.
// Parameters:
//  - value : Value to set for the machineName property.
func (m *CertificateConnectorDetails) SetMachineName(value *string)() {
    m.machineName = value
}
