package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkCameraConfiguration 
type TeamworkCameraConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The cameras property
    cameras []TeamworkPeripheralable
    // The configuration for the content camera.
    contentCameraConfiguration TeamworkContentCameraConfigurationable
    // The defaultContentCamera property
    defaultContentCamera TeamworkPeripheralable
    // The OdataType property
    odataType *string
}
// NewTeamworkCameraConfiguration instantiates a new teamworkCameraConfiguration and sets the default values.
func NewTeamworkCameraConfiguration()(*TeamworkCameraConfiguration) {
    m := &TeamworkCameraConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.teamworkCameraConfiguration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTeamworkCameraConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkCameraConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkCameraConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkCameraConfiguration) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCameras gets the cameras property value. The cameras property
func (m *TeamworkCameraConfiguration) GetCameras()([]TeamworkPeripheralable) {
    return m.cameras
}
// GetContentCameraConfiguration gets the contentCameraConfiguration property value. The configuration for the content camera.
func (m *TeamworkCameraConfiguration) GetContentCameraConfiguration()(TeamworkContentCameraConfigurationable) {
    return m.contentCameraConfiguration
}
// GetDefaultContentCamera gets the defaultContentCamera property value. The defaultContentCamera property
func (m *TeamworkCameraConfiguration) GetDefaultContentCamera()(TeamworkPeripheralable) {
    return m.defaultContentCamera
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkCameraConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cameras"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateTeamworkPeripheralFromDiscriminatorValue , m.SetCameras)
    res["contentCameraConfiguration"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkContentCameraConfigurationFromDiscriminatorValue , m.SetContentCameraConfiguration)
    res["defaultContentCamera"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue , m.SetDefaultContentCamera)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkCameraConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TeamworkCameraConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCameras() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCameras())
        err := writer.WriteCollectionOfObjectValues("cameras", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("contentCameraConfiguration", m.GetContentCameraConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("defaultContentCamera", m.GetDefaultContentCamera())
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
func (m *TeamworkCameraConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCameras sets the cameras property value. The cameras property
func (m *TeamworkCameraConfiguration) SetCameras(value []TeamworkPeripheralable)() {
    m.cameras = value
}
// SetContentCameraConfiguration sets the contentCameraConfiguration property value. The configuration for the content camera.
func (m *TeamworkCameraConfiguration) SetContentCameraConfiguration(value TeamworkContentCameraConfigurationable)() {
    m.contentCameraConfiguration = value
}
// SetDefaultContentCamera sets the defaultContentCamera property value. The defaultContentCamera property
func (m *TeamworkCameraConfiguration) SetDefaultContentCamera(value TeamworkPeripheralable)() {
    m.defaultContentCamera = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkCameraConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
