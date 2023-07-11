package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImportedDeviceIdentity the importedDeviceIdentity resource represents a unique hardware identity of a device that has been pre-staged for pre-enrollment configuration.
type ImportedDeviceIdentity struct {
    Entity
}
// NewImportedDeviceIdentity instantiates a new importedDeviceIdentity and sets the default values.
func NewImportedDeviceIdentity()(*ImportedDeviceIdentity) {
    m := &ImportedDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateImportedDeviceIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedDeviceIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.importedDeviceIdentityResult":
                        return NewImportedDeviceIdentityResult(), nil
                }
            }
        }
    }
    return NewImportedDeviceIdentity(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Created Date Time of the device
func (m *ImportedDeviceIdentity) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
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
func (m *ImportedDeviceIdentity) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnrollmentState gets the enrollmentState property value. The enrollmentState property
func (m *ImportedDeviceIdentity) GetEnrollmentState()(*EnrollmentState) {
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
func (m *ImportedDeviceIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["importedDeviceIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportedDeviceIdentifier(val)
        }
        return nil
    }
    res["importedDeviceIdentityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportedDeviceIdentityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportedDeviceIdentityType(val.(*ImportedDeviceIdentityType))
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
    return res
}
// GetImportedDeviceIdentifier gets the importedDeviceIdentifier property value. Imported Device Identifier
func (m *ImportedDeviceIdentity) GetImportedDeviceIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("importedDeviceIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetImportedDeviceIdentityType gets the importedDeviceIdentityType property value. The importedDeviceIdentityType property
func (m *ImportedDeviceIdentity) GetImportedDeviceIdentityType()(*ImportedDeviceIdentityType) {
    val, err := m.GetBackingStore().Get("importedDeviceIdentityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ImportedDeviceIdentityType)
    }
    return nil
}
// GetLastContactedDateTime gets the lastContactedDateTime property value. Last Contacted Date Time of the device
func (m *ImportedDeviceIdentity) GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastContactedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last Modified DateTime of the description
func (m *ImportedDeviceIdentity) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ImportedDeviceIdentity) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPlatform gets the platform property value. The platform property
func (m *ImportedDeviceIdentity) GetPlatform()(*Platform) {
    val, err := m.GetBackingStore().Get("platform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Platform)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ImportedDeviceIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetEnrollmentState() != nil {
        cast := (*m.GetEnrollmentState()).String()
        err = writer.WriteStringValue("enrollmentState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("importedDeviceIdentifier", m.GetImportedDeviceIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetImportedDeviceIdentityType() != nil {
        cast := (*m.GetImportedDeviceIdentityType()).String()
        err = writer.WriteStringValue("importedDeviceIdentityType", &cast)
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
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Created Date Time of the device
func (m *ImportedDeviceIdentity) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the device
func (m *ImportedDeviceIdentity) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentState sets the enrollmentState property value. The enrollmentState property
func (m *ImportedDeviceIdentity) SetEnrollmentState(value *EnrollmentState)() {
    err := m.GetBackingStore().Set("enrollmentState", value)
    if err != nil {
        panic(err)
    }
}
// SetImportedDeviceIdentifier sets the importedDeviceIdentifier property value. Imported Device Identifier
func (m *ImportedDeviceIdentity) SetImportedDeviceIdentifier(value *string)() {
    err := m.GetBackingStore().Set("importedDeviceIdentifier", value)
    if err != nil {
        panic(err)
    }
}
// SetImportedDeviceIdentityType sets the importedDeviceIdentityType property value. The importedDeviceIdentityType property
func (m *ImportedDeviceIdentity) SetImportedDeviceIdentityType(value *ImportedDeviceIdentityType)() {
    err := m.GetBackingStore().Set("importedDeviceIdentityType", value)
    if err != nil {
        panic(err)
    }
}
// SetLastContactedDateTime sets the lastContactedDateTime property value. Last Contacted Date Time of the device
func (m *ImportedDeviceIdentity) SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastContactedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last Modified DateTime of the description
func (m *ImportedDeviceIdentity) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ImportedDeviceIdentity) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPlatform sets the platform property value. The platform property
func (m *ImportedDeviceIdentity) SetPlatform(value *Platform)() {
    err := m.GetBackingStore().Set("platform", value)
    if err != nil {
        panic(err)
    }
}
// ImportedDeviceIdentityable 
type ImportedDeviceIdentityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetEnrollmentState()(*EnrollmentState)
    GetImportedDeviceIdentifier()(*string)
    GetImportedDeviceIdentityType()(*ImportedDeviceIdentityType)
    GetLastContactedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    GetPlatform()(*Platform)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetEnrollmentState(value *EnrollmentState)()
    SetImportedDeviceIdentifier(value *string)()
    SetImportedDeviceIdentityType(value *ImportedDeviceIdentityType)()
    SetLastContactedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
    SetPlatform(value *Platform)()
}
