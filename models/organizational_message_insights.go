package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageInsights contains statistics into how the organizational message was interacted with by clients
type OrganizationalMessageInsights struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The number of times this message was clicked on
    clicks *int32
    // The number of times this message was dismissed by a user. This may not be collected for some surfaces and will be null
    dismisses *int32
    // The number of times this message was shown to all clients
    impressions *int32
    // The OdataType property
    odataType *string
}
// NewOrganizationalMessageInsights instantiates a new organizationalMessageInsights and sets the default values.
func NewOrganizationalMessageInsights()(*OrganizationalMessageInsights) {
    m := &OrganizationalMessageInsights{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.organizationalMessageInsights";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOrganizationalMessageInsightsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalMessageInsightsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalMessageInsights(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationalMessageInsights) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClicks gets the clicks property value. The number of times this message was clicked on
func (m *OrganizationalMessageInsights) GetClicks()(*int32) {
    return m.clicks
}
// GetDismisses gets the dismisses property value. The number of times this message was dismissed by a user. This may not be collected for some surfaces and will be null
func (m *OrganizationalMessageInsights) GetDismisses()(*int32) {
    return m.dismisses
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalMessageInsights) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clicks"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetClicks)
    res["dismisses"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetDismisses)
    res["impressions"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetImpressions)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetImpressions gets the impressions property value. The number of times this message was shown to all clients
func (m *OrganizationalMessageInsights) GetImpressions()(*int32) {
    return m.impressions
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageInsights) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *OrganizationalMessageInsights) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("clicks", m.GetClicks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("dismisses", m.GetDismisses())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("impressions", m.GetImpressions())
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
func (m *OrganizationalMessageInsights) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClicks sets the clicks property value. The number of times this message was clicked on
func (m *OrganizationalMessageInsights) SetClicks(value *int32)() {
    m.clicks = value
}
// SetDismisses sets the dismisses property value. The number of times this message was dismissed by a user. This may not be collected for some surfaces and will be null
func (m *OrganizationalMessageInsights) SetDismisses(value *int32)() {
    m.dismisses = value
}
// SetImpressions sets the impressions property value. The number of times this message was shown to all clients
func (m *OrganizationalMessageInsights) SetImpressions(value *int32)() {
    m.impressions = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageInsights) SetOdataType(value *string)() {
    m.odataType = value
}
