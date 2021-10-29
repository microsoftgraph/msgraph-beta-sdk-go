package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DirectoryDefinition struct {
    Entity
    // Read only value indicating what type of discovery the app supports. Possible values are: AttributeDataTypes, AttributeNames, AttributeReadOnly, None, ReferenceAttributes, UnknownFutureValue.
    discoverabilities *DirectoryDefinitionDiscoverabilities;
    // Represents the discovery date and time using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    discoveryDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Name of the directory. Must be unique within the synchronization schema. Not nullable.
    name *string;
    // Collection of objects supported by the directory.
    objects []ObjectDefinition;
    // 
    readOnly *bool;
    // Read only value that indicates version discovered. null if discovery has not yet occurred.
    version *string;
}
// Instantiates a new directoryDefinition and sets the default values.
func NewDirectoryDefinition()(*DirectoryDefinition) {
    m := &DirectoryDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the discoverabilities property value. Read only value indicating what type of discovery the app supports. Possible values are: AttributeDataTypes, AttributeNames, AttributeReadOnly, None, ReferenceAttributes, UnknownFutureValue.
func (m *DirectoryDefinition) GetDiscoverabilities()(*DirectoryDefinitionDiscoverabilities) {
    if m == nil {
        return nil
    } else {
        return m.discoverabilities
    }
}
// Gets the discoveryDateTime property value. Represents the discovery date and time using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *DirectoryDefinition) GetDiscoveryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.discoveryDateTime
    }
}
// Gets the name property value. Name of the directory. Must be unique within the synchronization schema. Not nullable.
func (m *DirectoryDefinition) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the objects property value. Collection of objects supported by the directory.
func (m *DirectoryDefinition) GetObjects()([]ObjectDefinition) {
    if m == nil {
        return nil
    } else {
        return m.objects
    }
}
// Gets the readOnly property value. 
func (m *DirectoryDefinition) GetReadOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.readOnly
    }
}
// Gets the version property value. Read only value that indicates version discovered. null if discovery has not yet occurred.
func (m *DirectoryDefinition) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *DirectoryDefinition) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["discoverabilities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDirectoryDefinitionDiscoverabilities)
        if err != nil {
            return err
        }
        cast := val.(DirectoryDefinitionDiscoverabilities)
        m.SetDiscoverabilities(&cast)
        return nil
    }
    res["discoveryDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDiscoveryDateTime(val)
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
    res["objects"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewObjectDefinition() })
        if err != nil {
            return err
        }
        res := make([]ObjectDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*ObjectDefinition))
        }
        m.SetObjects(res)
        return nil
    }
    res["readOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetReadOnly(val)
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVersion(val)
        return nil
    }
    return res
}
func (m *DirectoryDefinition) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DirectoryDefinition) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDiscoverabilities() != nil {
        cast := m.GetDiscoverabilities().String()
        err = writer.WriteStringValue("discoverabilities", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("discoveryDateTime", m.GetDiscoveryDateTime())
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetObjects()))
        for i, v := range m.GetObjects() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("objects", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("readOnly", m.GetReadOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the discoverabilities property value. Read only value indicating what type of discovery the app supports. Possible values are: AttributeDataTypes, AttributeNames, AttributeReadOnly, None, ReferenceAttributes, UnknownFutureValue.
// Parameters:
//  - value : Value to set for the discoverabilities property.
func (m *DirectoryDefinition) SetDiscoverabilities(value *DirectoryDefinitionDiscoverabilities)() {
    m.discoverabilities = value
}
// Sets the discoveryDateTime property value. Represents the discovery date and time using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the discoveryDateTime property.
func (m *DirectoryDefinition) SetDiscoveryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.discoveryDateTime = value
}
// Sets the name property value. Name of the directory. Must be unique within the synchronization schema. Not nullable.
// Parameters:
//  - value : Value to set for the name property.
func (m *DirectoryDefinition) SetName(value *string)() {
    m.name = value
}
// Sets the objects property value. Collection of objects supported by the directory.
// Parameters:
//  - value : Value to set for the objects property.
func (m *DirectoryDefinition) SetObjects(value []ObjectDefinition)() {
    m.objects = value
}
// Sets the readOnly property value. 
// Parameters:
//  - value : Value to set for the readOnly property.
func (m *DirectoryDefinition) SetReadOnly(value *bool)() {
    m.readOnly = value
}
// Sets the version property value. Read only value that indicates version discovered. null if discovery has not yet occurred.
// Parameters:
//  - value : Value to set for the version property.
func (m *DirectoryDefinition) SetVersion(value *string)() {
    m.version = value
}
