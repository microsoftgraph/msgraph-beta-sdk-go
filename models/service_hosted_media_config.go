package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServiceHostedMediaConfig struct {
    MediaConfig
}
// NewServiceHostedMediaConfig instantiates a new ServiceHostedMediaConfig and sets the default values.
func NewServiceHostedMediaConfig()(*ServiceHostedMediaConfig) {
    m := &ServiceHostedMediaConfig{
        MediaConfig: *NewMediaConfig(),
    }
    odataTypeValue := "#microsoft.graph.serviceHostedMediaConfig"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateServiceHostedMediaConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceHostedMediaConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceHostedMediaConfig(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServiceHostedMediaConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MediaConfig.GetFieldDeserializers()
    res["liveCaptionOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLiveCaptionOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLiveCaptionOptions(val.(LiveCaptionOptionsable))
        }
        return nil
    }
    res["preFetchMedia"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMediaInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MediaInfoable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(MediaInfoable)
                }
            }
            m.SetPreFetchMedia(res)
        }
        return nil
    }
    return res
}
// GetLiveCaptionOptions gets the liveCaptionOptions property value. The liveCaptionOptions property
// returns a LiveCaptionOptionsable when successful
func (m *ServiceHostedMediaConfig) GetLiveCaptionOptions()(LiveCaptionOptionsable) {
    val, err := m.GetBackingStore().Get("liveCaptionOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LiveCaptionOptionsable)
    }
    return nil
}
// GetPreFetchMedia gets the preFetchMedia property value. The list of media to prefetch.
// returns a []MediaInfoable when successful
func (m *ServiceHostedMediaConfig) GetPreFetchMedia()([]MediaInfoable) {
    val, err := m.GetBackingStore().Get("preFetchMedia")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]MediaInfoable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ServiceHostedMediaConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MediaConfig.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("liveCaptionOptions", m.GetLiveCaptionOptions())
        if err != nil {
            return err
        }
    }
    if m.GetPreFetchMedia() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPreFetchMedia()))
        for i, v := range m.GetPreFetchMedia() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("preFetchMedia", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLiveCaptionOptions sets the liveCaptionOptions property value. The liveCaptionOptions property
func (m *ServiceHostedMediaConfig) SetLiveCaptionOptions(value LiveCaptionOptionsable)() {
    err := m.GetBackingStore().Set("liveCaptionOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetPreFetchMedia sets the preFetchMedia property value. The list of media to prefetch.
func (m *ServiceHostedMediaConfig) SetPreFetchMedia(value []MediaInfoable)() {
    err := m.GetBackingStore().Set("preFetchMedia", value)
    if err != nil {
        panic(err)
    }
}
type ServiceHostedMediaConfigable interface {
    MediaConfigable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLiveCaptionOptions()(LiveCaptionOptionsable)
    GetPreFetchMedia()([]MediaInfoable)
    SetLiveCaptionOptions(value LiveCaptionOptionsable)()
    SetPreFetchMedia(value []MediaInfoable)()
}
