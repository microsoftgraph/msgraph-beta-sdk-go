package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationOneRosterApiDataProvider 
type EducationOneRosterApiDataProvider struct {
    EducationSynchronizationDataProvider
}
// NewEducationOneRosterApiDataProvider instantiates a new educationOneRosterApiDataProvider and sets the default values.
func NewEducationOneRosterApiDataProvider()(*EducationOneRosterApiDataProvider) {
    m := &EducationOneRosterApiDataProvider{
        EducationSynchronizationDataProvider: *NewEducationSynchronizationDataProvider(),
    }
    odataTypeValue := "#microsoft.graph.educationOneRosterApiDataProvider"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationOneRosterApiDataProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationOneRosterApiDataProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationOneRosterApiDataProvider(), nil
}
// GetConnectionSettings gets the connectionSettings property value. The connectionSettings property
func (m *EducationOneRosterApiDataProvider) GetConnectionSettings()(EducationSynchronizationConnectionSettingsable) {
    val, err := m.GetBackingStore().Get("connectionSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationSynchronizationConnectionSettingsable)
    }
    return nil
}
// GetConnectionUrl gets the connectionUrl property value. The connection URL to the OneRoster instance.
func (m *EducationOneRosterApiDataProvider) GetConnectionUrl()(*string) {
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
func (m *EducationOneRosterApiDataProvider) GetCustomizations()(EducationSynchronizationCustomizationsable) {
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
func (m *EducationOneRosterApiDataProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationSynchronizationDataProvider.GetFieldDeserializers()
    res["connectionSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationConnectionSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionSettings(val.(EducationSynchronizationConnectionSettingsable))
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
    res["providerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderName(val)
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
    res["termIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTermIds(res)
        }
        return nil
    }
    return res
}
// GetProviderName gets the providerName property value. The OneRoster Service Provider name as defined by the [OneRoster specification][oneroster].
func (m *EducationOneRosterApiDataProvider) GetProviderName()(*string) {
    val, err := m.GetBackingStore().Get("providerName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSchoolsIds gets the schoolsIds property value. The list of [School/Org][orgs] sourcedId to sync.
func (m *EducationOneRosterApiDataProvider) GetSchoolsIds()([]string) {
    val, err := m.GetBackingStore().Get("schoolsIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetTermIds gets the termIds property value. The list of [academic sessions][terms] to sync.
func (m *EducationOneRosterApiDataProvider) GetTermIds()([]string) {
    val, err := m.GetBackingStore().Get("termIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
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
    err := m.GetBackingStore().Set("connectionSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetConnectionUrl sets the connectionUrl property value. The connection URL to the OneRoster instance.
func (m *EducationOneRosterApiDataProvider) SetConnectionUrl(value *string)() {
    err := m.GetBackingStore().Set("connectionUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomizations sets the customizations property value. Optional customization to be applied to the synchronization profile.
func (m *EducationOneRosterApiDataProvider) SetCustomizations(value EducationSynchronizationCustomizationsable)() {
    err := m.GetBackingStore().Set("customizations", value)
    if err != nil {
        panic(err)
    }
}
// SetProviderName sets the providerName property value. The OneRoster Service Provider name as defined by the [OneRoster specification][oneroster].
func (m *EducationOneRosterApiDataProvider) SetProviderName(value *string)() {
    err := m.GetBackingStore().Set("providerName", value)
    if err != nil {
        panic(err)
    }
}
// SetSchoolsIds sets the schoolsIds property value. The list of [School/Org][orgs] sourcedId to sync.
func (m *EducationOneRosterApiDataProvider) SetSchoolsIds(value []string)() {
    err := m.GetBackingStore().Set("schoolsIds", value)
    if err != nil {
        panic(err)
    }
}
// SetTermIds sets the termIds property value. The list of [academic sessions][terms] to sync.
func (m *EducationOneRosterApiDataProvider) SetTermIds(value []string)() {
    err := m.GetBackingStore().Set("termIds", value)
    if err != nil {
        panic(err)
    }
}
// EducationOneRosterApiDataProviderable 
type EducationOneRosterApiDataProviderable interface {
    EducationSynchronizationDataProviderable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectionSettings()(EducationSynchronizationConnectionSettingsable)
    GetConnectionUrl()(*string)
    GetCustomizations()(EducationSynchronizationCustomizationsable)
    GetProviderName()(*string)
    GetSchoolsIds()([]string)
    GetTermIds()([]string)
    SetConnectionSettings(value EducationSynchronizationConnectionSettingsable)()
    SetConnectionUrl(value *string)()
    SetCustomizations(value EducationSynchronizationCustomizationsable)()
    SetProviderName(value *string)()
    SetSchoolsIds(value []string)()
    SetTermIds(value []string)()
}
