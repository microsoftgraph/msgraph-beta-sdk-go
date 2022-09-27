package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataProcessorServiceForWindowsFeaturesOnboarding a configuration entity for MEM features that utilize Data Processor Service for Windows (DPSW) data.
type DataProcessorServiceForWindowsFeaturesOnboarding struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Indicates whether the tenant has enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. When TRUE, the tenant has enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. When FALSE, the tenant has not enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. Default value is FALSE.
    areDataProcessorServiceForWindowsFeaturesEnabled *bool
    // Indicates whether the tenant has required Windows license. When TRUE, the tenant has the required Windows license. When FALSE, the tenant does not have the required Windows license. Default value is FALSE.
    hasValidWindowsLicense *bool
    // The OdataType property
    odataType *string
}
// NewDataProcessorServiceForWindowsFeaturesOnboarding instantiates a new dataProcessorServiceForWindowsFeaturesOnboarding and sets the default values.
func NewDataProcessorServiceForWindowsFeaturesOnboarding()(*DataProcessorServiceForWindowsFeaturesOnboarding) {
    m := &DataProcessorServiceForWindowsFeaturesOnboarding{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.dataProcessorServiceForWindowsFeaturesOnboarding";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDataProcessorServiceForWindowsFeaturesOnboardingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataProcessorServiceForWindowsFeaturesOnboardingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataProcessorServiceForWindowsFeaturesOnboarding(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAreDataProcessorServiceForWindowsFeaturesEnabled gets the areDataProcessorServiceForWindowsFeaturesEnabled property value. Indicates whether the tenant has enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. When TRUE, the tenant has enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. When FALSE, the tenant has not enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. Default value is FALSE.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) GetAreDataProcessorServiceForWindowsFeaturesEnabled()(*bool) {
    return m.areDataProcessorServiceForWindowsFeaturesEnabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["areDataProcessorServiceForWindowsFeaturesEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetAreDataProcessorServiceForWindowsFeaturesEnabled)
    res["hasValidWindowsLicense"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetHasValidWindowsLicense)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetHasValidWindowsLicense gets the hasValidWindowsLicense property value. Indicates whether the tenant has required Windows license. When TRUE, the tenant has the required Windows license. When FALSE, the tenant does not have the required Windows license. Default value is FALSE.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) GetHasValidWindowsLicense()(*bool) {
    return m.hasValidWindowsLicense
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("areDataProcessorServiceForWindowsFeaturesEnabled", m.GetAreDataProcessorServiceForWindowsFeaturesEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasValidWindowsLicense", m.GetHasValidWindowsLicense())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAreDataProcessorServiceForWindowsFeaturesEnabled sets the areDataProcessorServiceForWindowsFeaturesEnabled property value. Indicates whether the tenant has enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. When TRUE, the tenant has enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. When FALSE, the tenant has not enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. Default value is FALSE.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) SetAreDataProcessorServiceForWindowsFeaturesEnabled(value *bool)() {
    m.areDataProcessorServiceForWindowsFeaturesEnabled = value
}
// SetHasValidWindowsLicense sets the hasValidWindowsLicense property value. Indicates whether the tenant has required Windows license. When TRUE, the tenant has the required Windows license. When FALSE, the tenant does not have the required Windows license. Default value is FALSE.
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) SetHasValidWindowsLicense(value *bool)() {
    m.hasValidWindowsLicense = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DataProcessorServiceForWindowsFeaturesOnboarding) SetOdataType(value *string)() {
    m.odataType = value
}
