package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationOneRosterApiDataProvider 
type EducationOneRosterApiDataProvider struct {
    EducationSynchronizationDataProvider
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
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
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationOneRosterApiDataProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationOneRosterApiDataProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationOneRosterApiDataProvider(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationOneRosterApiDataProvider) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetConnectionSettings gets the connectionSettings property value. The connectionSettings property
func (m *EducationOneRosterApiDataProvider) GetConnectionSettings()(EducationSynchronizationConnectionSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.connectionSettings
    }
}
// GetConnectionUrl gets the connectionUrl property value. The connectionUrl property
func (m *EducationOneRosterApiDataProvider) GetConnectionUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectionUrl
    }
}
// GetCustomizations gets the customizations property value. The customizations property
func (m *EducationOneRosterApiDataProvider) GetCustomizations()(EducationSynchronizationCustomizationsable) {
    if m == nil {
        return nil
    } else {
        return m.customizations
    }
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
                res[i] = *(v.(*string))
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
                res[i] = *(v.(*string))
            }
            m.SetTermIds(res)
        }
        return nil
    }
    return res
}
// GetProviderName gets the providerName property value. The providerName property
func (m *EducationOneRosterApiDataProvider) GetProviderName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.providerName
    }
}
// GetSchoolsIds gets the schoolsIds property value. The schoolsIds property
func (m *EducationOneRosterApiDataProvider) GetSchoolsIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.schoolsIds
    }
}
// GetTermIds gets the termIds property value. The termIds property
func (m *EducationOneRosterApiDataProvider) GetTermIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.termIds
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationOneRosterApiDataProvider) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetConnectionSettings sets the connectionSettings property value. The connectionSettings property
func (m *EducationOneRosterApiDataProvider) SetConnectionSettings(value EducationSynchronizationConnectionSettingsable)() {
    if m != nil {
        m.connectionSettings = value
    }
}
// SetConnectionUrl sets the connectionUrl property value. The connectionUrl property
func (m *EducationOneRosterApiDataProvider) SetConnectionUrl(value *string)() {
    if m != nil {
        m.connectionUrl = value
    }
}
// SetCustomizations sets the customizations property value. The customizations property
func (m *EducationOneRosterApiDataProvider) SetCustomizations(value EducationSynchronizationCustomizationsable)() {
    if m != nil {
        m.customizations = value
    }
}
// SetProviderName sets the providerName property value. The providerName property
func (m *EducationOneRosterApiDataProvider) SetProviderName(value *string)() {
    if m != nil {
        m.providerName = value
    }
}
// SetSchoolsIds sets the schoolsIds property value. The schoolsIds property
func (m *EducationOneRosterApiDataProvider) SetSchoolsIds(value []string)() {
    if m != nil {
        m.schoolsIds = value
    }
}
// SetTermIds sets the termIds property value. The termIds property
func (m *EducationOneRosterApiDataProvider) SetTermIds(value []string)() {
    if m != nil {
        m.termIds = value
    }
}
