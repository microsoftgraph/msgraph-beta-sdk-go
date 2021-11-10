package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CartToClassAssociation struct {
    Entity
    // Identifiers of classrooms to be associated with device carts.
    classroomIds []string;
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Admin provided description of the CartToClassAssociation.
    description *string;
    // Identifiers of device carts to be associated with classes.
    deviceCartIds []string;
    // Admin provided name of the device configuration.
    displayName *string;
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Version of the CartToClassAssociation.
    version *int32;
}
// Instantiates a new cartToClassAssociation and sets the default values.
func NewCartToClassAssociation()(*CartToClassAssociation) {
    m := &CartToClassAssociation{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the classroomIds property value. Identifiers of classrooms to be associated with device carts.
func (m *CartToClassAssociation) GetClassroomIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.classroomIds
    }
}
// Gets the createdDateTime property value. DateTime the object was created.
func (m *CartToClassAssociation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Admin provided description of the CartToClassAssociation.
func (m *CartToClassAssociation) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the deviceCartIds property value. Identifiers of device carts to be associated with classes.
func (m *CartToClassAssociation) GetDeviceCartIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.deviceCartIds
    }
}
// Gets the displayName property value. Admin provided name of the device configuration.
func (m *CartToClassAssociation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *CartToClassAssociation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the version property value. Version of the CartToClassAssociation.
func (m *CartToClassAssociation) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// The deserialization information for the current model
func (m *CartToClassAssociation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classroomIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetClassroomIds(res)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceCartIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDeviceCartIds(res)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *CartToClassAssociation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CartToClassAssociation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("classroomIds", m.GetClassroomIds())
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
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("deviceCartIds", m.GetDeviceCartIds())
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
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the classroomIds property value. Identifiers of classrooms to be associated with device carts.
// Parameters:
//  - value : Value to set for the classroomIds property.
func (m *CartToClassAssociation) SetClassroomIds(value []string)() {
    m.classroomIds = value
}
// Sets the createdDateTime property value. DateTime the object was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *CartToClassAssociation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Admin provided description of the CartToClassAssociation.
// Parameters:
//  - value : Value to set for the description property.
func (m *CartToClassAssociation) SetDescription(value *string)() {
    m.description = value
}
// Sets the deviceCartIds property value. Identifiers of device carts to be associated with classes.
// Parameters:
//  - value : Value to set for the deviceCartIds property.
func (m *CartToClassAssociation) SetDeviceCartIds(value []string)() {
    m.deviceCartIds = value
}
// Sets the displayName property value. Admin provided name of the device configuration.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *CartToClassAssociation) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastModifiedDateTime property value. DateTime the object was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *CartToClassAssociation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the version property value. Version of the CartToClassAssociation.
// Parameters:
//  - value : Value to set for the version property.
func (m *CartToClassAssociation) SetVersion(value *int32)() {
    m.version = value
}
