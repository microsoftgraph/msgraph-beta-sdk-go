package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConnectorGroup struct {
    Entity
    // Read-only. Nullable.
    applications []Application;
    // Indicates the type of hybrid agent. This pre-set by the system. Possible values are: applicationProxy. Read-only.
    connectorGroupType *ConnectorGroupType;
    // Indicates if the connectorGroup is the default connectorGroup. Only a single connector group can be the default connectorGroup and this is pre-set by the system. Read-only.
    isDefault *bool;
    // Read-only. Nullable.
    members []Connector;
    // The name associated with the connectorGroup.
    name *string;
    // The region the connectorGroup is assigned to and will optimize traffic for. This region can only be set if no connectors or applications are assigned to the connectorGroup. The possible values are: nam (for North America), eur (for Europe), aus (for Australia), asia (for Asia), ind (for India), and unknownFutureValue.
    region *ConnectorGroupRegion;
}
// Instantiates a new connectorGroup and sets the default values.
func NewConnectorGroup()(*ConnectorGroup) {
    m := &ConnectorGroup{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the applications property value. Read-only. Nullable.
func (m *ConnectorGroup) GetApplications()([]Application) {
    if m == nil {
        return nil
    } else {
        return m.applications
    }
}
// Gets the connectorGroupType property value. Indicates the type of hybrid agent. This pre-set by the system. Possible values are: applicationProxy. Read-only.
func (m *ConnectorGroup) GetConnectorGroupType()(*ConnectorGroupType) {
    if m == nil {
        return nil
    } else {
        return m.connectorGroupType
    }
}
// Gets the isDefault property value. Indicates if the connectorGroup is the default connectorGroup. Only a single connector group can be the default connectorGroup and this is pre-set by the system. Read-only.
func (m *ConnectorGroup) GetIsDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefault
    }
}
// Gets the members property value. Read-only. Nullable.
func (m *ConnectorGroup) GetMembers()([]Connector) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// Gets the name property value. The name associated with the connectorGroup.
func (m *ConnectorGroup) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the region property value. The region the connectorGroup is assigned to and will optimize traffic for. This region can only be set if no connectors or applications are assigned to the connectorGroup. The possible values are: nam (for North America), eur (for Europe), aus (for Australia), asia (for Asia), ind (for India), and unknownFutureValue.
func (m *ConnectorGroup) GetRegion()(*ConnectorGroupRegion) {
    if m == nil {
        return nil
    } else {
        return m.region
    }
}
// The deserialization information for the current model
func (m *ConnectorGroup) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApplication() })
        if err != nil {
            return err
        }
        res := make([]Application, len(val))
        for i, v := range val {
            res[i] = *(v.(*Application))
        }
        m.SetApplications(res)
        return nil
    }
    res["connectorGroupType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectorGroupType)
        if err != nil {
            return err
        }
        cast := val.(ConnectorGroupType)
        m.SetConnectorGroupType(&cast)
        return nil
    }
    res["isDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefault(val)
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConnector() })
        if err != nil {
            return err
        }
        res := make([]Connector, len(val))
        for i, v := range val {
            res[i] = *(v.(*Connector))
        }
        m.SetMembers(res)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["region"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseConnectorGroupRegion)
        if err != nil {
            return err
        }
        cast := val.(ConnectorGroupRegion)
        m.SetRegion(&cast)
        return nil
    }
    return res
}
func (m *ConnectorGroup) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ConnectorGroup) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApplications()))
        for i, v := range m.GetApplications() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("applications", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConnectorGroupType() != nil {
        cast := m.GetConnectorGroupType().String()
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        cast := m.GetRegion().String()
        err = writer.WriteStringValue("region", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the applications property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the applications property.
func (m *ConnectorGroup) SetApplications(value []Application)() {
    m.applications = value
}
// Sets the connectorGroupType property value. Indicates the type of hybrid agent. This pre-set by the system. Possible values are: applicationProxy. Read-only.
// Parameters:
//  - value : Value to set for the connectorGroupType property.
func (m *ConnectorGroup) SetConnectorGroupType(value *ConnectorGroupType)() {
    m.connectorGroupType = value
}
// Sets the isDefault property value. Indicates if the connectorGroup is the default connectorGroup. Only a single connector group can be the default connectorGroup and this is pre-set by the system. Read-only.
// Parameters:
//  - value : Value to set for the isDefault property.
func (m *ConnectorGroup) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// Sets the members property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the members property.
func (m *ConnectorGroup) SetMembers(value []Connector)() {
    m.members = value
}
// Sets the name property value. The name associated with the connectorGroup.
// Parameters:
//  - value : Value to set for the name property.
func (m *ConnectorGroup) SetName(value *string)() {
    m.name = value
}
// Sets the region property value. The region the connectorGroup is assigned to and will optimize traffic for. This region can only be set if no connectors or applications are assigned to the connectorGroup. The possible values are: nam (for North America), eur (for Europe), aus (for Australia), asia (for Asia), ind (for India), and unknownFutureValue.
// Parameters:
//  - value : Value to set for the region property.
func (m *ConnectorGroup) SetRegion(value *ConnectorGroupRegion)() {
    m.region = value
}
