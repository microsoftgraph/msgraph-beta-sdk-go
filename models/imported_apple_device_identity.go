package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImportedAppleDeviceIdentity the importedAppleDeviceIdentity resource represents the imported device identity of an Apple device .
type ImportedAppleDeviceIdentity struct {
    Entity
}
// NewImportedAppleDeviceIdentity instantiates a new ImportedAppleDeviceIdentity and sets the default values.
func NewImportedAppleDeviceIdentity()(*ImportedAppleDeviceIdentity) {
    m := &ImportedAppleDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateImportedAppleDeviceIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateImportedAppleDeviceIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.importedAppleDeviceIdentityResult":
                        return NewImportedAppleDeviceIdentityResult(), nil
                }
            }
        }
    }
    return NewImportedAppleDeviceIdentity(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Created Date Time of the device
// returns a *Time when successful
func (m *ImportedAppleDeviceIdentity) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. The description of the device
// returns a *string when successful
func (m *ImportedAppleDeviceIdentity) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDiscoverySource gets the discoverySource property value. The discoverySource property
// returns a *DiscoverySource when successful
func (m *ImportedAppleDeviceIdentity) GetDiscoverySource()(*DiscoverySource) {
    val, err := m.GetBackingStore().Get("discoverySource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DiscoverySource)
    }
    return nil
}
// GetEnrollmentState gets the enrollmentState property value. The enrollmentState property
// returns a *EnrollmentState when successful
func (m *ImportedAppleDeviceIdentity) GetEnrollmentState()(*EnrollmentState) {
    val, err := m.GetBackingStore().Get("enrollmentState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EnrollmentState)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ImportedAppleDeviceIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["discoverySource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiscoverySource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscoverySource(val.(*DiscoverySource))
        }
        return nil
    }
    res["enrollmentState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEnrollmentState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentState(val.(*EnrollmentState))
        }
        return nil
    }
    res["isDeleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeleted(val)
        }
        return nil
    }
    res["isSupervised"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSupervised(val)
        }
        return nil
    }
    res["lastContactedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastContactedDateTime(val)
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*Platform))
        }
        return nil
    }
    res["requestedEnrollmentProfileAssignmentDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedEnrollmentProfileAssignmentDateTime(val)
        }
        return nil
    }
    res["requestedEnrollmentProfileId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedEnrollmentProfileId(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    return res
}
// GetIsDeleted gets the isDeleted property value. Indicates if the device is deleted from Apple Business Manager
// returns a *bool when successful
func (m *ImportedAppleDeviceIdentity) GetIsDeleted()(*bool) {
    val, err := m.GetBackingStore().Get("isDeleted")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSupervised gets the isSupervised property value. Indicates if the Apple device is supervised.
// returns a *bool when successful
func (m *ImportedAppleDeviceIdentity) GetIsSupervised()(*bool) {
    val, err := m.GetBackingStore().Get("isSupervised")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastContactedDateTime gets the lastContactedDateTime property value. Last Contacted Date Time of the device
// returns a *Time when successful
func (m *ImportedAppleDeviceIdentity) GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastContactedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPlatform gets the platform property value. The platform property
// returns a *Platform when successful
func (m *ImportedAppleDeviceIdentity) GetPlatform()(*Platform) {
    val, err := m.GetBackingStore().Get("platform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Platform)
    }
    return nil
}
// GetRequestedEnrollmentProfileAssignmentDateTime gets the requestedEnrollmentProfileAssignmentDateTime property value. The time enrollment profile was assigned to the device
// returns a *Time when successful
func (m *ImportedAppleDeviceIdentity) GetRequestedEnrollmentProfileAssignmentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("requestedEnrollmentProfileAssignmentDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRequestedEnrollmentProfileId gets the requestedEnrollmentProfileId property value. Enrollment profile Id admin intends to apply to the device during next enrollment
// returns a *string when successful
func (m *ImportedAppleDeviceIdentity) GetRequestedEnrollmentProfileId()(*string) {
    val, err := m.GetBackingStore().Get("requestedEnrollmentProfileId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSerialNumber gets the serialNumber property value. Device serial number
// returns a *string when successful
func (m *ImportedAppleDeviceIdentity) GetSerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("serialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ImportedAppleDeviceIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetDiscoverySource() != nil {
        cast := (*m.GetDiscoverySource()).String()
        err = writer.WriteStringValue("discoverySource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentState() != nil {
        cast := (*m.GetEnrollmentState()).String()
        err = writer.WriteStringValue("enrollmentState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeleted", m.GetIsDeleted())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSupervised", m.GetIsSupervised())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastContactedDateTime", m.GetLastContactedDateTime())
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
    {
        err = writer.WriteTimeValue("requestedEnrollmentProfileAssignmentDateTime", m.GetRequestedEnrollmentProfileAssignmentDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestedEnrollmentProfileId", m.GetRequestedEnrollmentProfileId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Created Date Time of the device
func (m *ImportedAppleDeviceIdentity) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the device
func (m *ImportedAppleDeviceIdentity) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDiscoverySource sets the discoverySource property value. The discoverySource property
func (m *ImportedAppleDeviceIdentity) SetDiscoverySource(value *DiscoverySource)() {
    err := m.GetBackingStore().Set("discoverySource", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentState sets the enrollmentState property value. The enrollmentState property
func (m *ImportedAppleDeviceIdentity) SetEnrollmentState(value *EnrollmentState)() {
    err := m.GetBackingStore().Set("enrollmentState", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDeleted sets the isDeleted property value. Indicates if the device is deleted from Apple Business Manager
func (m *ImportedAppleDeviceIdentity) SetIsDeleted(value *bool)() {
    err := m.GetBackingStore().Set("isDeleted", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSupervised sets the isSupervised property value. Indicates if the Apple device is supervised.
func (m *ImportedAppleDeviceIdentity) SetIsSupervised(value *bool)() {
    err := m.GetBackingStore().Set("isSupervised", value)
    if err != nil {
        panic(err)
    }
}
// SetLastContactedDateTime sets the lastContactedDateTime property value. Last Contacted Date Time of the device
func (m *ImportedAppleDeviceIdentity) SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastContactedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatform sets the platform property value. The platform property
func (m *ImportedAppleDeviceIdentity) SetPlatform(value *Platform)() {
    err := m.GetBackingStore().Set("platform", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestedEnrollmentProfileAssignmentDateTime sets the requestedEnrollmentProfileAssignmentDateTime property value. The time enrollment profile was assigned to the device
func (m *ImportedAppleDeviceIdentity) SetRequestedEnrollmentProfileAssignmentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("requestedEnrollmentProfileAssignmentDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestedEnrollmentProfileId sets the requestedEnrollmentProfileId property value. Enrollment profile Id admin intends to apply to the device during next enrollment
func (m *ImportedAppleDeviceIdentity) SetRequestedEnrollmentProfileId(value *string)() {
    err := m.GetBackingStore().Set("requestedEnrollmentProfileId", value)
    if err != nil {
        panic(err)
    }
}
// SetSerialNumber sets the serialNumber property value. Device serial number
func (m *ImportedAppleDeviceIdentity) SetSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("serialNumber", value)
    if err != nil {
        panic(err)
    }
}
type ImportedAppleDeviceIdentityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDiscoverySource()(*DiscoverySource)
    GetEnrollmentState()(*EnrollmentState)
    GetIsDeleted()(*bool)
    GetIsSupervised()(*bool)
    GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPlatform()(*Platform)
    GetRequestedEnrollmentProfileAssignmentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestedEnrollmentProfileId()(*string)
    GetSerialNumber()(*string)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDiscoverySource(value *DiscoverySource)()
    SetEnrollmentState(value *EnrollmentState)()
    SetIsDeleted(value *bool)()
    SetIsSupervised(value *bool)()
    SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPlatform(value *Platform)()
    SetRequestedEnrollmentProfileAssignmentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestedEnrollmentProfileId(value *string)()
    SetSerialNumber(value *string)()
}
