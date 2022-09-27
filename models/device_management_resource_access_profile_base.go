package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementResourceAccessProfileBase base Profile Type for Resource Access
type DeviceManagementResourceAccessProfileBase struct {
    Entity
    // The list of assignments for the device configuration profile.
    assignments []DeviceManagementResourceAccessProfileAssignmentable
    // DateTime profile was created
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Profile description
    description *string
    // Profile display name
    displayName *string
    // DateTime profile was last modified
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Scope Tags
    roleScopeTagIds []string
    // Version of the profile
    version *int32
}
// NewDeviceManagementResourceAccessProfileBase instantiates a new deviceManagementResourceAccessProfileBase and sets the default values.
func NewDeviceManagementResourceAccessProfileBase()(*DeviceManagementResourceAccessProfileBase) {
    m := &DeviceManagementResourceAccessProfileBase{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementResourceAccessProfileBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementResourceAccessProfileBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementResourceAccessProfileBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windows10XCertificateProfile":
                        return NewWindows10XCertificateProfile(), nil
                    case "#microsoft.graph.windows10XSCEPCertificateProfile":
                        return NewWindows10XSCEPCertificateProfile(), nil
                    case "#microsoft.graph.windows10XTrustedRootCertificate":
                        return NewWindows10XTrustedRootCertificate(), nil
                    case "#microsoft.graph.windows10XVpnConfiguration":
                        return NewWindows10XVpnConfiguration(), nil
                    case "#microsoft.graph.windows10XWifiConfiguration":
                        return NewWindows10XWifiConfiguration(), nil
                }
            }
        }
    }
    return NewDeviceManagementResourceAccessProfileBase(), nil
}
// GetAssignments gets the assignments property value. The list of assignments for the device configuration profile.
func (m *DeviceManagementResourceAccessProfileBase) GetAssignments()([]DeviceManagementResourceAccessProfileAssignmentable) {
    return m.assignments
}
// GetCreationDateTime gets the creationDateTime property value. DateTime profile was created
func (m *DeviceManagementResourceAccessProfileBase) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.creationDateTime
}
// GetDescription gets the description property value. Profile description
func (m *DeviceManagementResourceAccessProfileBase) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Profile display name
func (m *DeviceManagementResourceAccessProfileBase) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementResourceAccessProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateDeviceManagementResourceAccessProfileAssignmentFromDiscriminatorValue , m.SetAssignments)
    res["creationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreationDateTime)
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["lastModifiedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastModifiedDateTime)
    res["roleScopeTagIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetRoleScopeTagIds)
    res["version"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetVersion)
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime profile was last modified
func (m *DeviceManagementResourceAccessProfileBase) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. Scope Tags
func (m *DeviceManagementResourceAccessProfileBase) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// GetVersion gets the version property value. Version of the profile
func (m *DeviceManagementResourceAccessProfileBase) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceManagementResourceAccessProfileBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignments())
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
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
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
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
// SetAssignments sets the assignments property value. The list of assignments for the device configuration profile.
func (m *DeviceManagementResourceAccessProfileBase) SetAssignments(value []DeviceManagementResourceAccessProfileAssignmentable)() {
    m.assignments = value
}
// SetCreationDateTime sets the creationDateTime property value. DateTime profile was created
func (m *DeviceManagementResourceAccessProfileBase) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
// SetDescription sets the description property value. Profile description
func (m *DeviceManagementResourceAccessProfileBase) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Profile display name
func (m *DeviceManagementResourceAccessProfileBase) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime profile was last modified
func (m *DeviceManagementResourceAccessProfileBase) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. Scope Tags
func (m *DeviceManagementResourceAccessProfileBase) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// SetVersion sets the version property value. Version of the profile
func (m *DeviceManagementResourceAccessProfileBase) SetVersion(value *int32)() {
    m.version = value
}
