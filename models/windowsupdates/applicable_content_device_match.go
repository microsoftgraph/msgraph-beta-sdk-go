package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicableContentDeviceMatch 
type ApplicableContentDeviceMatch struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deviceId property
    deviceId *string
    // The OdataType property
    odataType *string
    // The recommendedBy property
    recommendedBy []string
}
// NewApplicableContentDeviceMatch instantiates a new applicableContentDeviceMatch and sets the default values.
func NewApplicableContentDeviceMatch()(*ApplicableContentDeviceMatch) {
    m := &ApplicableContentDeviceMatch{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApplicableContentDeviceMatchFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicableContentDeviceMatchFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicableContentDeviceMatch(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplicableContentDeviceMatch) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceId gets the deviceId property value. The deviceId property
func (m *ApplicableContentDeviceMatch) GetDeviceId()(*string) {
    return m.deviceId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicableContentDeviceMatch) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
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
    res["recommendedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRecommendedBy(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ApplicableContentDeviceMatch) GetOdataType()(*string) {
    return m.odataType
}
// GetRecommendedBy gets the recommendedBy property value. The recommendedBy property
func (m *ApplicableContentDeviceMatch) GetRecommendedBy()([]string) {
    return m.recommendedBy
}
// Serialize serializes information the current object
func (m *ApplicableContentDeviceMatch) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceId", m.GetDeviceId())
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
    if m.GetRecommendedBy() != nil {
        err := writer.WriteCollectionOfStringValues("recommendedBy", m.GetRecommendedBy())
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
func (m *ApplicableContentDeviceMatch) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceId sets the deviceId property value. The deviceId property
func (m *ApplicableContentDeviceMatch) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ApplicableContentDeviceMatch) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecommendedBy sets the recommendedBy property value. The recommendedBy property
func (m *ApplicableContentDeviceMatch) SetRecommendedBy(value []string)() {
    m.recommendedBy = value
}
