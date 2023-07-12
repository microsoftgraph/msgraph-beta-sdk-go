package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAndAppManagementAssignmentFilter a class containing the properties used for Assignment Filter.
type DeviceAndAppManagementAssignmentFilter struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewDeviceAndAppManagementAssignmentFilter instantiates a new deviceAndAppManagementAssignmentFilter and sets the default values.
func NewDeviceAndAppManagementAssignmentFilter()(*DeviceAndAppManagementAssignmentFilter) {
    m := &DeviceAndAppManagementAssignmentFilter{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceAndAppManagementAssignmentFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAndAppManagementAssignmentFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.payloadCompatibleAssignmentFilter":
                        return NewPayloadCompatibleAssignmentFilter(), nil
                }
            }
        }
    }
    return NewDeviceAndAppManagementAssignmentFilter(), nil
}
// GetAssignmentFilterManagementType gets the assignmentFilterManagementType property value. Supported filter management types whether its devices or apps.
func (m *DeviceAndAppManagementAssignmentFilter) GetAssignmentFilterManagementType()(*AssignmentFilterManagementType) {
    val, err := m.GetBackingStore().Get("assignmentFilterManagementType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AssignmentFilterManagementType)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The creation time of the assignment filter. The value cannot be modified and is automatically populated during new assignment filter process. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'.
func (m *DeviceAndAppManagementAssignmentFilter) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. Optional description of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The name of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAndAppManagementAssignmentFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentFilterManagementType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssignmentFilterManagementType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterManagementType(val.(*AssignmentFilterManagementType))
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["payloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePayloadByFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PayloadByFilterable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(PayloadByFilterable)
                }
            }
            m.SetPayloads(res)
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDevicePlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*DevicePlatformType))
        }
        return nil
    }
    res["roleScopeTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetRoleScopeTags(res)
        }
        return nil
    }
    res["rule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRule(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modified time of the Assignment Filter. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
func (m *DeviceAndAppManagementAssignmentFilter) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPayloads gets the payloads property value. Indicates associated assignments for a specific filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetPayloads()([]PayloadByFilterable) {
    val, err := m.GetBackingStore().Get("payloads")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]PayloadByFilterable)
    }
    return nil
}
// GetPlatform gets the platform property value. Supported platform types.
func (m *DeviceAndAppManagementAssignmentFilter) GetPlatform()(*DevicePlatformType) {
    val, err := m.GetBackingStore().Get("platform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DevicePlatformType)
    }
    return nil
}
// GetRoleScopeTags gets the roleScopeTags property value. Indicates role scope tags assigned for the assignment filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetRoleScopeTags()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTags")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRule gets the rule property value. Rule definition of the assignment filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetRule()(*string) {
    val, err := m.GetBackingStore().Get("rule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DeviceAndAppManagementAssignmentFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignmentFilterManagementType() != nil {
        cast := (*m.GetAssignmentFilterManagementType()).String()
        err = writer.WriteStringValue("assignmentFilterManagementType", &cast)
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPayloads() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPayloads()))
        for i, v := range m.GetPayloads() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("payloads", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTags() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTags", m.GetRoleScopeTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rule", m.GetRule())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignmentFilterManagementType sets the assignmentFilterManagementType property value. Supported filter management types whether its devices or apps.
func (m *DeviceAndAppManagementAssignmentFilter) SetAssignmentFilterManagementType(value *AssignmentFilterManagementType)() {
    err := m.GetBackingStore().Set("assignmentFilterManagementType", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The creation time of the assignment filter. The value cannot be modified and is automatically populated during new assignment filter process. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'.
func (m *DeviceAndAppManagementAssignmentFilter) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Optional description of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modified time of the Assignment Filter. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'
func (m *DeviceAndAppManagementAssignmentFilter) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPayloads sets the payloads property value. Indicates associated assignments for a specific filter.
func (m *DeviceAndAppManagementAssignmentFilter) SetPayloads(value []PayloadByFilterable)() {
    err := m.GetBackingStore().Set("payloads", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatform sets the platform property value. Supported platform types.
func (m *DeviceAndAppManagementAssignmentFilter) SetPlatform(value *DevicePlatformType)() {
    err := m.GetBackingStore().Set("platform", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTags sets the roleScopeTags property value. Indicates role scope tags assigned for the assignment filter.
func (m *DeviceAndAppManagementAssignmentFilter) SetRoleScopeTags(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTags", value)
    if err != nil {
        panic(err)
    }
}
// SetRule sets the rule property value. Rule definition of the assignment filter.
func (m *DeviceAndAppManagementAssignmentFilter) SetRule(value *string)() {
    err := m.GetBackingStore().Set("rule", value)
    if err != nil {
        panic(err)
    }
}
// DeviceAndAppManagementAssignmentFilterable 
type DeviceAndAppManagementAssignmentFilterable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentFilterManagementType()(*AssignmentFilterManagementType)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPayloads()([]PayloadByFilterable)
    GetPlatform()(*DevicePlatformType)
    GetRoleScopeTags()([]string)
    GetRule()(*string)
    SetAssignmentFilterManagementType(value *AssignmentFilterManagementType)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPayloads(value []PayloadByFilterable)()
    SetPlatform(value *DevicePlatformType)()
    SetRoleScopeTags(value []string)()
    SetRule(value *string)()
}
