package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type TrainingCampaignReportOverview struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTrainingCampaignReportOverview instantiates a new TrainingCampaignReportOverview and sets the default values.
func NewTrainingCampaignReportOverview()(*TrainingCampaignReportOverview) {
    m := &TrainingCampaignReportOverview{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTrainingCampaignReportOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTrainingCampaignReportOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrainingCampaignReportOverview(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TrainingCampaignReportOverview) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *TrainingCampaignReportOverview) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TrainingCampaignReportOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["trainingModuleCompletion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTrainingEventsContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingModuleCompletion(val.(TrainingEventsContentable))
        }
        return nil
    }
    res["trainingNotificationDeliveryStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTrainingNotificationDeliveryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrainingNotificationDeliveryStatus(val.(TrainingNotificationDeliveryable))
        }
        return nil
    }
    res["userCompletionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserTrainingCompletionSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserCompletionStatus(val.(UserTrainingCompletionSummaryable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *TrainingCampaignReportOverview) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTrainingModuleCompletion gets the trainingModuleCompletion property value. The trainingModuleCompletion property
// returns a TrainingEventsContentable when successful
func (m *TrainingCampaignReportOverview) GetTrainingModuleCompletion()(TrainingEventsContentable) {
    val, err := m.GetBackingStore().Get("trainingModuleCompletion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TrainingEventsContentable)
    }
    return nil
}
// GetTrainingNotificationDeliveryStatus gets the trainingNotificationDeliveryStatus property value. The trainingNotificationDeliveryStatus property
// returns a TrainingNotificationDeliveryable when successful
func (m *TrainingCampaignReportOverview) GetTrainingNotificationDeliveryStatus()(TrainingNotificationDeliveryable) {
    val, err := m.GetBackingStore().Get("trainingNotificationDeliveryStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TrainingNotificationDeliveryable)
    }
    return nil
}
// GetUserCompletionStatus gets the userCompletionStatus property value. The userCompletionStatus property
// returns a UserTrainingCompletionSummaryable when successful
func (m *TrainingCampaignReportOverview) GetUserCompletionStatus()(UserTrainingCompletionSummaryable) {
    val, err := m.GetBackingStore().Get("userCompletionStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(UserTrainingCompletionSummaryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TrainingCampaignReportOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("trainingModuleCompletion", m.GetTrainingModuleCompletion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("trainingNotificationDeliveryStatus", m.GetTrainingNotificationDeliveryStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("userCompletionStatus", m.GetUserCompletionStatus())
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
func (m *TrainingCampaignReportOverview) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TrainingCampaignReportOverview) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TrainingCampaignReportOverview) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTrainingModuleCompletion sets the trainingModuleCompletion property value. The trainingModuleCompletion property
func (m *TrainingCampaignReportOverview) SetTrainingModuleCompletion(value TrainingEventsContentable)() {
    err := m.GetBackingStore().Set("trainingModuleCompletion", value)
    if err != nil {
        panic(err)
    }
}
// SetTrainingNotificationDeliveryStatus sets the trainingNotificationDeliveryStatus property value. The trainingNotificationDeliveryStatus property
func (m *TrainingCampaignReportOverview) SetTrainingNotificationDeliveryStatus(value TrainingNotificationDeliveryable)() {
    err := m.GetBackingStore().Set("trainingNotificationDeliveryStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetUserCompletionStatus sets the userCompletionStatus property value. The userCompletionStatus property
func (m *TrainingCampaignReportOverview) SetUserCompletionStatus(value UserTrainingCompletionSummaryable)() {
    err := m.GetBackingStore().Set("userCompletionStatus", value)
    if err != nil {
        panic(err)
    }
}
type TrainingCampaignReportOverviewable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetTrainingModuleCompletion()(TrainingEventsContentable)
    GetTrainingNotificationDeliveryStatus()(TrainingNotificationDeliveryable)
    GetUserCompletionStatus()(UserTrainingCompletionSummaryable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetTrainingModuleCompletion(value TrainingEventsContentable)()
    SetTrainingNotificationDeliveryStatus(value TrainingNotificationDeliveryable)()
    SetUserCompletionStatus(value UserTrainingCompletionSummaryable)()
}
