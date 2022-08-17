package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAndAppManagementAssignmentFilter 
type DeviceAndAppManagementAssignmentFilter struct {
    Entity
    // Creation time of the Assignment Filter.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description of the Assignment Filter.
    description *string
    // DisplayName of the Assignment Filter.
    displayName *string
    // Last modified time of the Assignment Filter.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Supported platform types.
    platform *DevicePlatformType
    // RoleScopeTags of the Assignment Filter.
    roleScopeTags []string
    // Rule definition of the Assignment Filter.
    rule *string
}
// NewDeviceAndAppManagementAssignmentFilter instantiates a new DeviceAndAppManagementAssignmentFilter and sets the default values.
func NewDeviceAndAppManagementAssignmentFilter()(*DeviceAndAppManagementAssignmentFilter) {
    m := &DeviceAndAppManagementAssignmentFilter{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceAndAppManagementAssignmentFilter";
    m.SetOdataType(&odataTypeValue);
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
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.payloadCompatibleAssignmentFilter":
                        return NewPayloadCompatibleAssignmentFilter(), nil
                }
            }
        }
    }
    return NewDeviceAndAppManagementAssignmentFilter(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Creation time of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. DisplayName of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAndAppManagementAssignmentFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
                res[i] = *(v.(*string))
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
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modified time of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPlatform gets the platform property value. Supported platform types.
func (m *DeviceAndAppManagementAssignmentFilter) GetPlatform()(*DevicePlatformType) {
    return m.platform
}
// GetRoleScopeTags gets the roleScopeTags property value. RoleScopeTags of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetRoleScopeTags()([]string) {
    return m.roleScopeTags
}
// GetRule gets the rule property value. Rule definition of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) GetRule()(*string) {
    return m.rule
}
// Serialize serializes information the current object
func (m *DeviceAndAppManagementAssignmentFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
// SetCreatedDateTime sets the createdDateTime property value. Creation time of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. DisplayName of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modified time of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPlatform sets the platform property value. Supported platform types.
func (m *DeviceAndAppManagementAssignmentFilter) SetPlatform(value *DevicePlatformType)() {
    m.platform = value
}
// SetRoleScopeTags sets the roleScopeTags property value. RoleScopeTags of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) SetRoleScopeTags(value []string)() {
    m.roleScopeTags = value
}
// SetRule sets the rule property value. Rule definition of the Assignment Filter.
func (m *DeviceAndAppManagementAssignmentFilter) SetRule(value *string)() {
    m.rule = value
}
