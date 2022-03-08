package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConnectorGroup provides operations to manage the collection of application entities.
type ConnectorGroup struct {
    Entity
    // Read-only. Nullable.
    applications []Applicationable;
    // Indicates the type of hybrid agent. This pre-set by the system. Possible values are: applicationProxy. Read-only.
    connectorGroupType *ConnectorGroupType;
    // Indicates if the connectorGroup is the default connectorGroup. Only a single connector group can be the default connectorGroup and this is pre-set by the system. Read-only.
    isDefault *bool;
    // Read-only. Nullable.
    members []Connectorable;
    // The name associated with the connectorGroup.
    name *string;
    // The region the connectorGroup is assigned to and will optimize traffic for. This region can only be set if no connectors or applications are assigned to the connectorGroup. The possible values are: nam (for North America), eur (for Europe), aus (for Australia), asia (for Asia), ind (for India), and unknownFutureValue.
    region *ConnectorGroupRegion;
}
// NewConnectorGroup instantiates a new connectorGroup and sets the default values.
func NewConnectorGroup()(*ConnectorGroup) {
    m := &ConnectorGroup{
        Entity: *NewEntity(),
    }
    return m
}
// CreateConnectorGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConnectorGroupFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewConnectorGroup(), nil
}
// GetApplications gets the applications property value. Read-only. Nullable.
func (m *ConnectorGroup) GetApplications()([]Applicationable) {
    if m == nil {
        return nil
    } else {
        return m.applications
    }
}
// GetConnectorGroupType gets the connectorGroupType property value. Indicates the type of hybrid agent. This pre-set by the system. Possible values are: applicationProxy. Read-only.
func (m *ConnectorGroup) GetConnectorGroupType()(*ConnectorGroupType) {
    if m == nil {
        return nil
    } else {
        return m.connectorGroupType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConnectorGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Applicationable, len(val))
            for i, v := range val {
                res[i] = v.(Applicationable)
            }
            m.SetApplications(res)
        }
        return nil
    }
    res["connectorGroupType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectorGroupType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorGroupType(val.(*ConnectorGroupType))
        }
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Connectorable, len(val))
            for i, v := range val {
                res[i] = v.(Connectorable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["region"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectorGroupRegion)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val.(*ConnectorGroupRegion))
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. Indicates if the connectorGroup is the default connectorGroup. Only a single connector group can be the default connectorGroup and this is pre-set by the system. Read-only.
func (m *ConnectorGroup) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// GetMembers gets the members property value. Read-only. Nullable.
func (m *ConnectorGroup) GetMembers()([]Connectorable) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// GetName gets the name property value. The name associated with the connectorGroup.
func (m *ConnectorGroup) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetRegion gets the region property value. The region the connectorGroup is assigned to and will optimize traffic for. This region can only be set if no connectors or applications are assigned to the connectorGroup. The possible values are: nam (for North America), eur (for Europe), aus (for Australia), asia (for Asia), ind (for India), and unknownFutureValue.
func (m *ConnectorGroup) GetRegion()(*ConnectorGroupRegion) {
    if m == nil {
        return nil
    } else {
        return m.region
    }
}
func (m *ConnectorGroup) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConnectorGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApplications() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApplications()))
        for i, v := range m.GetApplications() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("applications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConnectorGroupType() != nil {
        cast := (*m.GetConnectorGroupType()).String()
        err = writer.WriteStringValue("connectorGroupType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetRegion() != nil {
        cast := (*m.GetRegion()).String()
        err = writer.WriteStringValue("region", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplications sets the applications property value. Read-only. Nullable.
func (m *ConnectorGroup) SetApplications(value []Applicationable)() {
    if m != nil {
        m.applications = value
    }
}
// SetConnectorGroupType sets the connectorGroupType property value. Indicates the type of hybrid agent. This pre-set by the system. Possible values are: applicationProxy. Read-only.
func (m *ConnectorGroup) SetConnectorGroupType(value *ConnectorGroupType)() {
    if m != nil {
        m.connectorGroupType = value
    }
}
// SetIsDefault sets the isDefault property value. Indicates if the connectorGroup is the default connectorGroup. Only a single connector group can be the default connectorGroup and this is pre-set by the system. Read-only.
func (m *ConnectorGroup) SetIsDefault(value *bool)() {
    if m != nil {
        m.isDefault = value
    }
}
// SetMembers sets the members property value. Read-only. Nullable.
func (m *ConnectorGroup) SetMembers(value []Connectorable)() {
    if m != nil {
        m.members = value
    }
}
// SetName sets the name property value. The name associated with the connectorGroup.
func (m *ConnectorGroup) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetRegion sets the region property value. The region the connectorGroup is assigned to and will optimize traffic for. This region can only be set if no connectors or applications are assigned to the connectorGroup. The possible values are: nam (for North America), eur (for Europe), aus (for Australia), asia (for Asia), ind (for India), and unknownFutureValue.
func (m *ConnectorGroup) SetRegion(value *ConnectorGroupRegion)() {
    if m != nil {
        m.region = value
    }
}
