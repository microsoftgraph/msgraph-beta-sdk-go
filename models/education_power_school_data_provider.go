package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationPowerSchoolDataProvider 
type EducationPowerSchoolDataProvider struct {
    EducationSynchronizationDataProvider
    // Indicates whether the source has multiple identifiers for a single student or teacher.
    allowTeachersInMultipleSchools *bool
    // The client ID used to connect to PowerSchool.
    clientId *string
    // The client secret to authenticate the connection to the PowerSchool instance.
    clientSecret *string
    // The connection URL to the PowerSchool instance.
    connectionUrl *string
    // Optional customization to be applied to the synchronization profile.
    customizations EducationSynchronizationCustomizationsable
    // The list of schools to sync.
    schoolsIds []string
    // The school year to sync.
    schoolYear *string
}
// NewEducationPowerSchoolDataProvider instantiates a new EducationPowerSchoolDataProvider and sets the default values.
func NewEducationPowerSchoolDataProvider()(*EducationPowerSchoolDataProvider) {
    m := &EducationPowerSchoolDataProvider{
        EducationSynchronizationDataProvider: *NewEducationSynchronizationDataProvider(),
    }
    odataTypeValue := "#microsoft.graph.educationPowerSchoolDataProvider";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEducationPowerSchoolDataProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationPowerSchoolDataProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationPowerSchoolDataProvider(), nil
}
// GetAllowTeachersInMultipleSchools gets the allowTeachersInMultipleSchools property value. Indicates whether the source has multiple identifiers for a single student or teacher.
func (m *EducationPowerSchoolDataProvider) GetAllowTeachersInMultipleSchools()(*bool) {
    return m.allowTeachersInMultipleSchools
}
// GetClientId gets the clientId property value. The client ID used to connect to PowerSchool.
func (m *EducationPowerSchoolDataProvider) GetClientId()(*string) {
    return m.clientId
}
// GetClientSecret gets the clientSecret property value. The client secret to authenticate the connection to the PowerSchool instance.
func (m *EducationPowerSchoolDataProvider) GetClientSecret()(*string) {
    return m.clientSecret
}
// GetConnectionUrl gets the connectionUrl property value. The connection URL to the PowerSchool instance.
func (m *EducationPowerSchoolDataProvider) GetConnectionUrl()(*string) {
    return m.connectionUrl
}
// GetCustomizations gets the customizations property value. Optional customization to be applied to the synchronization profile.
func (m *EducationPowerSchoolDataProvider) GetCustomizations()(EducationSynchronizationCustomizationsable) {
    return m.customizations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationPowerSchoolDataProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationSynchronizationDataProvider.GetFieldDeserializers()
    res["allowTeachersInMultipleSchools"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAllowTeachersInMultipleSchools)
    res["clientId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetClientId)
    res["clientSecret"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetClientSecret)
    res["connectionUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetConnectionUrl)
    res["customizations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEducationSynchronizationCustomizationsFromDiscriminatorValue , m.SetCustomizations)
    res["schoolsIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSchoolsIds)
    res["schoolYear"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSchoolYear)
    return res
}
// GetSchoolsIds gets the schoolsIds property value. The list of schools to sync.
func (m *EducationPowerSchoolDataProvider) GetSchoolsIds()([]string) {
    return m.schoolsIds
}
// GetSchoolYear gets the schoolYear property value. The school year to sync.
func (m *EducationPowerSchoolDataProvider) GetSchoolYear()(*string) {
    return m.schoolYear
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
    m.allowTeachersInMultipleSchools = value
}
// SetClientId sets the clientId property value. The client ID used to connect to PowerSchool.
func (m *EducationPowerSchoolDataProvider) SetClientId(value *string)() {
    m.clientId = value
}
// SetClientSecret sets the clientSecret property value. The client secret to authenticate the connection to the PowerSchool instance.
func (m *EducationPowerSchoolDataProvider) SetClientSecret(value *string)() {
    m.clientSecret = value
}
// SetConnectionUrl sets the connectionUrl property value. The connection URL to the PowerSchool instance.
func (m *EducationPowerSchoolDataProvider) SetConnectionUrl(value *string)() {
    m.connectionUrl = value
}
// SetCustomizations sets the customizations property value. Optional customization to be applied to the synchronization profile.
func (m *EducationPowerSchoolDataProvider) SetCustomizations(value EducationSynchronizationCustomizationsable)() {
    m.customizations = value
}
// SetSchoolsIds sets the schoolsIds property value. The list of schools to sync.
func (m *EducationPowerSchoolDataProvider) SetSchoolsIds(value []string)() {
    m.schoolsIds = value
}
// SetSchoolYear sets the schoolYear property value. The school year to sync.
func (m *EducationPowerSchoolDataProvider) SetSchoolYear(value *string)() {
    m.schoolYear = value
}
