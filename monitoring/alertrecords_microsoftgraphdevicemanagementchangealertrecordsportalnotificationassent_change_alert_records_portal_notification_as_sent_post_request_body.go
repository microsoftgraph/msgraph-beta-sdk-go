package monitoring

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody instantiates a new AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody and sets the default values.
func NewAlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody()(*AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody) {
    m := &AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAlertRecordIds gets the alertRecordIds property value. The alertRecordIds property
// returns a []string when successful
func (m *AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody) GetAlertRecordIds()([]string) {
    val, err := m.GetBackingStore().Get("alertRecordIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alertRecordIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetAlertRecordIds(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAlertRecordIds() != nil {
        err := writer.WriteCollectionOfStringValues("alertRecordIds", m.GetAlertRecordIds())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAlertRecordIds sets the alertRecordIds property value. The alertRecordIds property
func (m *AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody) SetAlertRecordIds(value []string)() {
    err := m.GetBackingStore().Set("alertRecordIds", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
type AlertrecordsMicrosoftgraphdevicemanagementchangealertrecordsportalnotificationassentChangeAlertRecordsPortalNotificationAsSentPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlertRecordIds()([]string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    SetAlertRecordIds(value []string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
}
