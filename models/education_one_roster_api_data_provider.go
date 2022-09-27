package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationOneRosterApiDataProvider 
type EducationOneRosterApiDataProvider struct {
    EducationSynchronizationDataProvider
    // The connectionSettings property
    connectionSettings EducationSynchronizationConnectionSettingsable
    // The connectionUrl property
    connectionUrl *string
    // The customizations property
    customizations EducationSynchronizationCustomizationsable
    // The providerName property
    providerName *string
    // The schoolsIds property
    schoolsIds []string
    // The termIds property
    termIds []string
}
// NewEducationOneRosterApiDataProvider instantiates a new EducationOneRosterApiDataProvider and sets the default values.
func NewEducationOneRosterApiDataProvider()(*EducationOneRosterApiDataProvider) {
    m := &EducationOneRosterApiDataProvider{
        EducationSynchronizationDataProvider: *NewEducationSynchronizationDataProvider(),
    }
    odataTypeValue := "#microsoft.graph.educationOneRosterApiDataProvider";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEducationOneRosterApiDataProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationOneRosterApiDataProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationOneRosterApiDataProvider(), nil
}
// GetConnectionSettings gets the connectionSettings property value. The connectionSettings property
func (m *EducationOneRosterApiDataProvider) GetConnectionSettings()(EducationSynchronizationConnectionSettingsable) {
    return m.connectionSettings
}
// GetConnectionUrl gets the connectionUrl property value. The connectionUrl property
func (m *EducationOneRosterApiDataProvider) GetConnectionUrl()(*string) {
    return m.connectionUrl
}
// GetCustomizations gets the customizations property value. The customizations property
func (m *EducationOneRosterApiDataProvider) GetCustomizations()(EducationSynchronizationCustomizationsable) {
    return m.customizations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationOneRosterApiDataProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationSynchronizationDataProvider.GetFieldDeserializers()
    res["connectionSettings"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEducationSynchronizationConnectionSettingsFromDiscriminatorValue , m.SetConnectionSettings)
    res["connectionUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetConnectionUrl)
    res["customizations"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEducationSynchronizationCustomizationsFromDiscriminatorValue , m.SetCustomizations)
    res["providerName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetProviderName)
    res["schoolsIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetSchoolsIds)
    res["termIds"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetTermIds)
    return res
}
// GetProviderName gets the providerName property value. The providerName property
func (m *EducationOneRosterApiDataProvider) GetProviderName()(*string) {
    return m.providerName
}
// GetSchoolsIds gets the schoolsIds property value. The schoolsIds property
func (m *EducationOneRosterApiDataProvider) GetSchoolsIds()([]string) {
    return m.schoolsIds
}
// GetTermIds gets the termIds property value. The termIds property
func (m *EducationOneRosterApiDataProvider) GetTermIds()([]string) {
    return m.termIds
}
// Serialize serializes information the current object
func (m *EducationOneRosterApiDataProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationSynchronizationDataProvider.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("connectionSettings", m.GetConnectionSettings())
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
    {
        err = writer.WriteStringValue("providerName", m.GetProviderName())
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
    if m.GetTermIds() != nil {
        err = writer.WriteCollectionOfStringValues("termIds", m.GetTermIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnectionSettings sets the connectionSettings property value. The connectionSettings property
func (m *EducationOneRosterApiDataProvider) SetConnectionSettings(value EducationSynchronizationConnectionSettingsable)() {
    m.connectionSettings = value
}
// SetConnectionUrl sets the connectionUrl property value. The connectionUrl property
func (m *EducationOneRosterApiDataProvider) SetConnectionUrl(value *string)() {
    m.connectionUrl = value
}
// SetCustomizations sets the customizations property value. The customizations property
func (m *EducationOneRosterApiDataProvider) SetCustomizations(value EducationSynchronizationCustomizationsable)() {
    m.customizations = value
}
// SetProviderName sets the providerName property value. The providerName property
func (m *EducationOneRosterApiDataProvider) SetProviderName(value *string)() {
    m.providerName = value
}
// SetSchoolsIds sets the schoolsIds property value. The schoolsIds property
func (m *EducationOneRosterApiDataProvider) SetSchoolsIds(value []string)() {
    m.schoolsIds = value
}
// SetTermIds sets the termIds property value. The termIds property
func (m *EducationOneRosterApiDataProvider) SetTermIds(value []string)() {
    m.termIds = value
}
