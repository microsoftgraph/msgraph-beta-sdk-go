package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessPackageResourceEnvironment struct {
    Entity
    accessPackageResources []AccessPackageResource;
    connectionInfo *ConnectionInfo;
    createdBy *string;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    description *string;
    displayName *string;
    isDefaultEnvironment *bool;
    modifiedBy *string;
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    originId *string;
    originSystem *string;
}
func NewAccessPackageResourceEnvironment()(*AccessPackageResourceEnvironment) {
    m := &AccessPackageResourceEnvironment{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AccessPackageResourceEnvironment) GetAccessPackageResources()([]AccessPackageResource) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageResources
    }
}
func (m *AccessPackageResourceEnvironment) GetConnectionInfo()(*ConnectionInfo) {
    if m == nil {
        return nil
    } else {
        return m.connectionInfo
    }
}
func (m *AccessPackageResourceEnvironment) GetCreatedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
func (m *AccessPackageResourceEnvironment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *AccessPackageResourceEnvironment) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *AccessPackageResourceEnvironment) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *AccessPackageResourceEnvironment) GetIsDefaultEnvironment()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDefaultEnvironment
    }
}
func (m *AccessPackageResourceEnvironment) GetModifiedBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.modifiedBy
    }
}
func (m *AccessPackageResourceEnvironment) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.modifiedDateTime
    }
}
func (m *AccessPackageResourceEnvironment) GetOriginId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originId
    }
}
func (m *AccessPackageResourceEnvironment) GetOriginSystem()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originSystem
    }
}
func (m *AccessPackageResourceEnvironment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageResource() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageResource, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageResource))
        }
        m.SetAccessPackageResources(res)
        return nil
    }
    res["connectionInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConnectionInfo() })
        if err != nil {
            return err
        }
        m.SetConnectionInfo(val.(*ConnectionInfo))
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCreatedBy(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["isDefaultEnvironment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsDefaultEnvironment(val)
        return nil
    }
    res["modifiedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetModifiedBy(val)
        return nil
    }
    res["modifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetModifiedDateTime(val)
        return nil
    }
    res["originId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOriginId(val)
        return nil
    }
    res["originSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOriginSystem(val)
        return nil
    }
    return res
}
func (m *AccessPackageResourceEnvironment) IsNil()(bool) {
    return m == nil
}
func (m *AccessPackageResourceEnvironment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageResources()))
        for i, v := range m.GetAccessPackageResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageResources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("connectionInfo", m.GetConnectionInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultEnvironment", m.GetIsDefaultEnvironment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("modifiedBy", m.GetModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originId", m.GetOriginId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originSystem", m.GetOriginSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AccessPackageResourceEnvironment) SetAccessPackageResources(value []AccessPackageResource)() {
    m.accessPackageResources = value
}
func (m *AccessPackageResourceEnvironment) SetConnectionInfo(value *ConnectionInfo)() {
    m.connectionInfo = value
}
func (m *AccessPackageResourceEnvironment) SetCreatedBy(value *string)() {
    m.createdBy = value
}
func (m *AccessPackageResourceEnvironment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *AccessPackageResourceEnvironment) SetDescription(value *string)() {
    m.description = value
}
func (m *AccessPackageResourceEnvironment) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *AccessPackageResourceEnvironment) SetIsDefaultEnvironment(value *bool)() {
    m.isDefaultEnvironment = value
}
func (m *AccessPackageResourceEnvironment) SetModifiedBy(value *string)() {
    m.modifiedBy = value
}
func (m *AccessPackageResourceEnvironment) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
func (m *AccessPackageResourceEnvironment) SetOriginId(value *string)() {
    m.originId = value
}
func (m *AccessPackageResourceEnvironment) SetOriginSystem(value *string)() {
    m.originSystem = value
}
