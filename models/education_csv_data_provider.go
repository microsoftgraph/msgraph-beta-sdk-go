package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationCsvDataProvider 
type EducationCsvDataProvider struct {
    EducationSynchronizationDataProvider
}
// NewEducationCsvDataProvider instantiates a new educationCsvDataProvider and sets the default values.
func NewEducationCsvDataProvider()(*EducationCsvDataProvider) {
    m := &EducationCsvDataProvider{
        EducationSynchronizationDataProvider: *NewEducationSynchronizationDataProvider(),
    }
    odataTypeValue := "#microsoft.graph.educationCsvDataProvider"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationCsvDataProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationCsvDataProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationCsvDataProvider(), nil
}
// GetCustomizations gets the customizations property value. Optional customizations to be applied to the synchronization profile.
func (m *EducationCsvDataProvider) GetCustomizations()(EducationSynchronizationCustomizationsable) {
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
func (m *EducationCsvDataProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationSynchronizationDataProvider.GetFieldDeserializers()
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationCsvDataProvider) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationCsvDataProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationSynchronizationDataProvider.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("customizations", m.GetCustomizations())
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
    return nil
}
// SetCustomizations sets the customizations property value. Optional customizations to be applied to the synchronization profile.
func (m *EducationCsvDataProvider) SetCustomizations(value EducationSynchronizationCustomizationsable)() {
    err := m.GetBackingStore().Set("customizations", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationCsvDataProvider) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// EducationCsvDataProviderable 
type EducationCsvDataProviderable interface {
    EducationSynchronizationDataProviderable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomizations()(EducationSynchronizationCustomizationsable)
    GetOdataType()(*string)
    SetCustomizations(value EducationSynchronizationCustomizationsable)()
    SetOdataType(value *string)()
}
