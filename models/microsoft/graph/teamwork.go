package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Teamwork provides operations to manage the teamwork singleton.
type Teamwork struct {
    Entity
    // The Teams devices provisioned for the tenant.
    devices []TeamworkDeviceable;
    // A workforce integration with shifts.
    workforceIntegrations []WorkforceIntegrationable;
}
// NewTeamwork instantiates a new teamwork and sets the default values.
func NewTeamwork()(*Teamwork) {
    m := &Teamwork{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamwork(), nil
}
// GetDevices gets the devices property value. The Teams devices provisioned for the tenant.
func (m *Teamwork) GetDevices()([]TeamworkDeviceable) {
    if m == nil {
        return nil
    } else {
        return m.devices
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Teamwork) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["devices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamworkDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(TeamworkDeviceable)
            }
            m.SetDevices(res)
        }
        return nil
    }
    res["workforceIntegrations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWorkforceIntegrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WorkforceIntegrationable, len(val))
            for i, v := range val {
                res[i] = v.(WorkforceIntegrationable)
            }
            m.SetWorkforceIntegrations(res)
        }
        return nil
    }
    return res
}
// GetWorkforceIntegrations gets the workforceIntegrations property value. A workforce integration with shifts.
func (m *Teamwork) GetWorkforceIntegrations()([]WorkforceIntegrationable) {
    if m == nil {
        return nil
    } else {
        return m.workforceIntegrations
    }
}
func (m *Teamwork) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Teamwork) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDevices() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDevices()))
        for i, v := range m.GetDevices() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("devices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWorkforceIntegrations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWorkforceIntegrations()))
        for i, v := range m.GetWorkforceIntegrations() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("workforceIntegrations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDevices sets the devices property value. The Teams devices provisioned for the tenant.
func (m *Teamwork) SetDevices(value []TeamworkDeviceable)() {
    if m != nil {
        m.devices = value
    }
}
// SetWorkforceIntegrations sets the workforceIntegrations property value. A workforce integration with shifts.
func (m *Teamwork) SetWorkforceIntegrations(value []WorkforceIntegrationable)() {
    if m != nil {
        m.workforceIntegrations = value
    }
}
