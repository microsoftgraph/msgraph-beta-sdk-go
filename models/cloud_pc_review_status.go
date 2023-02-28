package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// CloudPcReviewStatus 
type CloudPcReviewStatus struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewCloudPcReviewStatus instantiates a new cloudPcReviewStatus and sets the default values.
func NewCloudPcReviewStatus()(*CloudPcReviewStatus) {
    m := &CloudPcReviewStatus{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCloudPcReviewStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcReviewStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcReviewStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcReviewStatus) GetAdditionalData()(map[string]any) {
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
// GetAzureStorageAccountId gets the azureStorageAccountId property value. The resource ID of the Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) GetAzureStorageAccountId()(*string) {
    val, err := m.GetBackingStore().Get("azureStorageAccountId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureStorageAccountName gets the azureStorageAccountName property value. The name of the Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) GetAzureStorageAccountName()(*string) {
    val, err := m.GetBackingStore().Get("azureStorageAccountName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureStorageContainerName gets the azureStorageContainerName property value. The name of the container in an Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) GetAzureStorageContainerName()(*string) {
    val, err := m.GetBackingStore().Get("azureStorageContainerName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *CloudPcReviewStatus) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcReviewStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["azureStorageAccountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureStorageAccountId(val)
        }
        return nil
    }
    res["azureStorageAccountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureStorageAccountName(val)
        }
        return nil
    }
    res["azureStorageContainerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureStorageContainerName(val)
        }
        return nil
    }
    res["inReview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInReview(val)
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
    res["restorePointDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestorePointDateTime(val)
        }
        return nil
    }
    res["reviewStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewStartDateTime(val)
        }
        return nil
    }
    res["subscriptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionId(val)
        }
        return nil
    }
    res["subscriptionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubscriptionName(val)
        }
        return nil
    }
    res["userAccessLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcUserAccessLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccessLevel(val.(*CloudPcUserAccessLevel))
        }
        return nil
    }
    return res
}
// GetInReview gets the inReview property value. True if the Cloud PC is set to in review by the administrator.
func (m *CloudPcReviewStatus) GetInReview()(*bool) {
    val, err := m.GetBackingStore().Get("inReview")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcReviewStatus) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRestorePointDateTime gets the restorePointDateTime property value. The specific date and time of the Cloud PC snapshot that was taken and saved automatically, when the Cloud PC is set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcReviewStatus) GetRestorePointDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("restorePointDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetReviewStartDateTime gets the reviewStartDateTime property value. The specific date and time when the Cloud PC was set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcReviewStatus) GetReviewStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("reviewStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSubscriptionId gets the subscriptionId property value. The ID of the Azure subscription in which the Cloud PC snapshot is being saved, in GUID format.
func (m *CloudPcReviewStatus) GetSubscriptionId()(*string) {
    val, err := m.GetBackingStore().Get("subscriptionId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubscriptionName gets the subscriptionName property value. The name of the Azure subscription in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) GetSubscriptionName()(*string) {
    val, err := m.GetBackingStore().Get("subscriptionName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserAccessLevel gets the userAccessLevel property value. The userAccessLevel property
func (m *CloudPcReviewStatus) GetUserAccessLevel()(*CloudPcUserAccessLevel) {
    val, err := m.GetBackingStore().Get("userAccessLevel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcUserAccessLevel)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudPcReviewStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("azureStorageAccountId", m.GetAzureStorageAccountId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("azureStorageAccountName", m.GetAzureStorageAccountName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("azureStorageContainerName", m.GetAzureStorageContainerName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("inReview", m.GetInReview())
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
        err := writer.WriteTimeValue("restorePointDateTime", m.GetRestorePointDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("reviewStartDateTime", m.GetReviewStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriptionId", m.GetSubscriptionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subscriptionName", m.GetSubscriptionName())
        if err != nil {
            return err
        }
    }
    if m.GetUserAccessLevel() != nil {
        cast := (*m.GetUserAccessLevel()).String()
        err := writer.WriteStringValue("userAccessLevel", &cast)
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
func (m *CloudPcReviewStatus) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureStorageAccountId sets the azureStorageAccountId property value. The resource ID of the Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) SetAzureStorageAccountId(value *string)() {
    err := m.GetBackingStore().Set("azureStorageAccountId", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureStorageAccountName sets the azureStorageAccountName property value. The name of the Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) SetAzureStorageAccountName(value *string)() {
    err := m.GetBackingStore().Set("azureStorageAccountName", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureStorageContainerName sets the azureStorageContainerName property value. The name of the container in an Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) SetAzureStorageContainerName(value *string)() {
    err := m.GetBackingStore().Set("azureStorageContainerName", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *CloudPcReviewStatus) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetInReview sets the inReview property value. True if the Cloud PC is set to in review by the administrator.
func (m *CloudPcReviewStatus) SetInReview(value *bool)() {
    err := m.GetBackingStore().Set("inReview", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcReviewStatus) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRestorePointDateTime sets the restorePointDateTime property value. The specific date and time of the Cloud PC snapshot that was taken and saved automatically, when the Cloud PC is set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcReviewStatus) SetRestorePointDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("restorePointDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewStartDateTime sets the reviewStartDateTime property value. The specific date and time when the Cloud PC was set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcReviewStatus) SetReviewStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("reviewStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSubscriptionId sets the subscriptionId property value. The ID of the Azure subscription in which the Cloud PC snapshot is being saved, in GUID format.
func (m *CloudPcReviewStatus) SetSubscriptionId(value *string)() {
    err := m.GetBackingStore().Set("subscriptionId", value)
    if err != nil {
        panic(err)
    }
}
// SetSubscriptionName sets the subscriptionName property value. The name of the Azure subscription in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) SetSubscriptionName(value *string)() {
    err := m.GetBackingStore().Set("subscriptionName", value)
    if err != nil {
        panic(err)
    }
}
// SetUserAccessLevel sets the userAccessLevel property value. The userAccessLevel property
func (m *CloudPcReviewStatus) SetUserAccessLevel(value *CloudPcUserAccessLevel)() {
    err := m.GetBackingStore().Set("userAccessLevel", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcReviewStatusable 
type CloudPcReviewStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureStorageAccountId()(*string)
    GetAzureStorageAccountName()(*string)
    GetAzureStorageContainerName()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetInReview()(*bool)
    GetOdataType()(*string)
    GetRestorePointDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReviewStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubscriptionId()(*string)
    GetSubscriptionName()(*string)
    GetUserAccessLevel()(*CloudPcUserAccessLevel)
    SetAzureStorageAccountId(value *string)()
    SetAzureStorageAccountName(value *string)()
    SetAzureStorageContainerName(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetInReview(value *bool)()
    SetOdataType(value *string)()
    SetRestorePointDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReviewStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubscriptionId(value *string)()
    SetSubscriptionName(value *string)()
    SetUserAccessLevel(value *CloudPcUserAccessLevel)()
}
