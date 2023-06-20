package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationPowerSchoolDataProvider 
type EducationPowerSchoolDataProvider struct {
    EducationSynchronizationDataProvider
}
// NewEducationPowerSchoolDataProvider instantiates a new EducationPowerSchoolDataProvider and sets the default values.
func NewEducationPowerSchoolDataProvider()(*EducationPowerSchoolDataProvider) {
    m := &EducationPowerSchoolDataProvider{
        EducationSynchronizationDataProvider: *NewEducationSynchronizationDataProvider(),
    }
    odataTypeValue := "#microsoft.graph.educationPowerSchoolDataProvider"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationPowerSchoolDataProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationPowerSchoolDataProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationPowerSchoolDataProvider(), nil
}
// GetAllowTeachersInMultipleSchools gets the allowTeachersInMultipleSchools property value. Indicates whether the source has multiple identifiers for a single student or teacher.
func (m *EducationPowerSchoolDataProvider) GetAllowTeachersInMultipleSchools()(*bool) {
    val, err := m.GetBackingStore().Get("allowTeachersInMultipleSchools")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetClientId gets the clientId property value. The client ID used to connect to PowerSchool.
func (m *EducationPowerSchoolDataProvider) GetClientId()(*string) {
    val, err := m.GetBackingStore().Get("clientId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetClientSecret gets the clientSecret property value. The client secret to authenticate the connection to the PowerSchool instance.
func (m *EducationPowerSchoolDataProvider) GetClientSecret()(*string) {
    val, err := m.GetBackingStore().Get("clientSecret")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConnectionUrl gets the connectionUrl property value. The connection URL to the PowerSchool instance.
func (m *EducationPowerSchoolDataProvider) GetConnectionUrl()(*string) {
    val, err := m.GetBackingStore().Get("connectionUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomizations gets the customizations property value. Optional customization to be applied to the synchronization profile.
func (m *EducationPowerSchoolDataProvider) GetCustomizations()(EducationSynchronizationCustomizationsable) {
    val, err := m.GetBackingStore().Get("customizations")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationSynchronizationCustomizationsable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationPowerSchoolDataProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationSynchronizationDataProvider.GetFieldDeserializers()
    res["allowTeachersInMultipleSchools"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTeachersInMultipleSchools(val)
        }
        return nil
    }
    res["clientId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["clientSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
        }
        return nil
    }
    res["connectionUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionUrl(val)
        }
        return nil
    }
    res["customizations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationCustomizationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomizations(val.(EducationSynchronizationCustomizationsable))
        }
        return nil
    }
    res["schoolsIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSchoolsIds(res)
        }
        return nil
    }
    res["schoolYear"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchoolYear(val)
        }
        return nil
    }
    return res
}
// GetSchoolsIds gets the schoolsIds property value. The list of schools to sync.
func (m *EducationPowerSchoolDataProvider) GetSchoolsIds()([]string) {
    val, err := m.GetBackingStore().Get("schoolsIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSchoolYear gets the schoolYear property value. The school year to sync.
func (m *EducationPowerSchoolDataProvider) GetSchoolYear()(*string) {
    val, err := m.GetBackingStore().Get("schoolYear")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationPowerSchoolDataProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationSynchronizationDataProvider.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowTeachersInMultipleSchools", m.GetAllowTeachersInMultipleSchools())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientSecret", m.GetClientSecret())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("connectionUrl", m.GetConnectionUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("customizations", m.GetCustomizations())
        if err != nil {
            return err
        }
    }
    if m.GetSchoolsIds() != nil {
        err = writer.WriteCollectionOfStringValues("schoolsIds", m.GetSchoolsIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("schoolYear", m.GetSchoolYear())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowTeachersInMultipleSchools sets the allowTeachersInMultipleSchools property value. Indicates whether the source has multiple identifiers for a single student or teacher.
func (m *EducationPowerSchoolDataProvider) SetAllowTeachersInMultipleSchools(value *bool)() {
    err := m.GetBackingStore().Set("allowTeachersInMultipleSchools", value)
    if err != nil {
        panic(err)
    }
}
// SetClientId sets the clientId property value. The client ID used to connect to PowerSchool.
func (m *EducationPowerSchoolDataProvider) SetClientId(value *string)() {
    err := m.GetBackingStore().Set("clientId", value)
    if err != nil {
        panic(err)
    }
}
// SetClientSecret sets the clientSecret property value. The client secret to authenticate the connection to the PowerSchool instance.
func (m *EducationPowerSchoolDataProvider) SetClientSecret(value *string)() {
    err := m.GetBackingStore().Set("clientSecret", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionUrl sets the connectionUrl property value. The connection URL to the PowerSchool instance.
func (m *EducationPowerSchoolDataProvider) SetConnectionUrl(value *string)() {
    err := m.GetBackingStore().Set("connectionUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomizations sets the customizations property value. Optional customization to be applied to the synchronization profile.
func (m *EducationPowerSchoolDataProvider) SetCustomizations(value EducationSynchronizationCustomizationsable)() {
    err := m.GetBackingStore().Set("customizations", value)
    if err != nil {
        panic(err)
    }
}
// SetSchoolsIds sets the schoolsIds property value. The list of schools to sync.
func (m *EducationPowerSchoolDataProvider) SetSchoolsIds(value []string)() {
    err := m.GetBackingStore().Set("schoolsIds", value)
    if err != nil {
        panic(err)
    }
}
// SetSchoolYear sets the schoolYear property value. The school year to sync.
func (m *EducationPowerSchoolDataProvider) SetSchoolYear(value *string)() {
    err := m.GetBackingStore().Set("schoolYear", value)
    if err != nil {
        panic(err)
    }
}
// EducationPowerSchoolDataProviderable 
type EducationPowerSchoolDataProviderable interface {
    EducationSynchronizationDataProviderable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowTeachersInMultipleSchools()(*bool)
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetConnectionUrl()(*string)
    GetCustomizations()(EducationSynchronizationCustomizationsable)
    GetSchoolsIds()([]string)
    GetSchoolYear()(*string)
    SetAllowTeachersInMultipleSchools(value *bool)()
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetConnectionUrl(value *string)()
    SetCustomizations(value EducationSynchronizationCustomizationsable)()
    SetSchoolsIds(value []string)()
    SetSchoolYear(value *string)()
}
