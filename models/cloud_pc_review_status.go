package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcReviewStatus 
type CloudPcReviewStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The resource ID of the Azure Storage account in which the Cloud PC snapshot is being saved.
    azureStorageAccountId *string
    // The name of the Azure Storage account in which the Cloud PC snapshot is being saved.
    azureStorageAccountName *string
    // The name of the container in an Azure Storage account in which the Cloud PC snapshot is being saved.
    azureStorageContainerName *string
    // True if the Cloud PC is set to in review by the administrator.
    inReview *bool
    // The OdataType property
    odataType *string
    // The specific date and time of the Cloud PC snapshot that was taken and saved automatically, when the Cloud PC is set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
    restorePointDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The specific date and time when the Cloud PC was set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
    reviewStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ID of the Azure subscription in which the Cloud PC snapshot is being saved, in GUID format.
    subscriptionId *string
    // The name of the Azure subscription in which the Cloud PC snapshot is being saved.
    subscriptionName *string
    // The userAccessLevel property
    userAccessLevel *CloudPcUserAccessLevel
}
// NewCloudPcReviewStatus instantiates a new cloudPcReviewStatus and sets the default values.
func NewCloudPcReviewStatus()(*CloudPcReviewStatus) {
    m := &CloudPcReviewStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCloudPcReviewStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcReviewStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcReviewStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcReviewStatus) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAzureStorageAccountId gets the azureStorageAccountId property value. The resource ID of the Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) GetAzureStorageAccountId()(*string) {
    return m.azureStorageAccountId
}
// GetAzureStorageAccountName gets the azureStorageAccountName property value. The name of the Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) GetAzureStorageAccountName()(*string) {
    return m.azureStorageAccountName
}
// GetAzureStorageContainerName gets the azureStorageContainerName property value. The name of the container in an Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) GetAzureStorageContainerName()(*string) {
    return m.azureStorageContainerName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcReviewStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["azureStorageAccountId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureStorageAccountId)
    res["azureStorageAccountName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureStorageAccountName)
    res["azureStorageContainerName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAzureStorageContainerName)
    res["inReview"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetInReview)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["restorePointDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetRestorePointDateTime)
    res["reviewStartDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetReviewStartDateTime)
    res["subscriptionId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubscriptionId)
    res["subscriptionName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSubscriptionName)
    res["userAccessLevel"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCloudPcUserAccessLevel , m.SetUserAccessLevel)
    return res
}
// GetInReview gets the inReview property value. True if the Cloud PC is set to in review by the administrator.
func (m *CloudPcReviewStatus) GetInReview()(*bool) {
    return m.inReview
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CloudPcReviewStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetRestorePointDateTime gets the restorePointDateTime property value. The specific date and time of the Cloud PC snapshot that was taken and saved automatically, when the Cloud PC is set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcReviewStatus) GetRestorePointDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.restorePointDateTime
}
// GetReviewStartDateTime gets the reviewStartDateTime property value. The specific date and time when the Cloud PC was set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcReviewStatus) GetReviewStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.reviewStartDateTime
}
// GetSubscriptionId gets the subscriptionId property value. The ID of the Azure subscription in which the Cloud PC snapshot is being saved, in GUID format.
func (m *CloudPcReviewStatus) GetSubscriptionId()(*string) {
    return m.subscriptionId
}
// GetSubscriptionName gets the subscriptionName property value. The name of the Azure subscription in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) GetSubscriptionName()(*string) {
    return m.subscriptionName
}
// GetUserAccessLevel gets the userAccessLevel property value. The userAccessLevel property
func (m *CloudPcReviewStatus) GetUserAccessLevel()(*CloudPcUserAccessLevel) {
    return m.userAccessLevel
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
func (m *CloudPcReviewStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAzureStorageAccountId sets the azureStorageAccountId property value. The resource ID of the Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) SetAzureStorageAccountId(value *string)() {
    m.azureStorageAccountId = value
}
// SetAzureStorageAccountName sets the azureStorageAccountName property value. The name of the Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) SetAzureStorageAccountName(value *string)() {
    m.azureStorageAccountName = value
}
// SetAzureStorageContainerName sets the azureStorageContainerName property value. The name of the container in an Azure Storage account in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) SetAzureStorageContainerName(value *string)() {
    m.azureStorageContainerName = value
}
// SetInReview sets the inReview property value. True if the Cloud PC is set to in review by the administrator.
func (m *CloudPcReviewStatus) SetInReview(value *bool)() {
    m.inReview = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CloudPcReviewStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRestorePointDateTime sets the restorePointDateTime property value. The specific date and time of the Cloud PC snapshot that was taken and saved automatically, when the Cloud PC is set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcReviewStatus) SetRestorePointDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.restorePointDateTime = value
}
// SetReviewStartDateTime sets the reviewStartDateTime property value. The specific date and time when the Cloud PC was set to in review. The timestamp is shown in ISO 8601 format and Coordinated Universal Time (UTC). For example, midnight UTC on Jan 1, 2014 appears as 2014-01-01T00:00:00Z.
func (m *CloudPcReviewStatus) SetReviewStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewStartDateTime = value
}
// SetSubscriptionId sets the subscriptionId property value. The ID of the Azure subscription in which the Cloud PC snapshot is being saved, in GUID format.
func (m *CloudPcReviewStatus) SetSubscriptionId(value *string)() {
    m.subscriptionId = value
}
// SetSubscriptionName sets the subscriptionName property value. The name of the Azure subscription in which the Cloud PC snapshot is being saved.
func (m *CloudPcReviewStatus) SetSubscriptionName(value *string)() {
    m.subscriptionName = value
}
// SetUserAccessLevel sets the userAccessLevel property value. The userAccessLevel property
func (m *CloudPcReviewStatus) SetUserAccessLevel(value *CloudPcUserAccessLevel)() {
    m.userAccessLevel = value
}
