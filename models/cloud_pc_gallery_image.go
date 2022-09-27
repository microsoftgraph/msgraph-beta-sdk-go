package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcGalleryImage 
type CloudPcGalleryImage struct {
    Entity
    // The official display name of the gallery image. Read-only.
    displayName *string
    // The date in which this image is no longer within long-term support. The Cloud PC will continue to provide short-term support. Read-only.
    endDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The date when the image is no longer available. Read-only.
    expirationDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The offer name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
    offer *string
    // The official display offer name of the gallery image. For example, Windows 10 Enterprise + OS Optimizations. Read-only.
    offerDisplayName *string
    // The publisher name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
    publisher *string
    // Recommended Cloud PC SKU for this gallery image. Read-only.
    recommendedSku *string
    // The size of this image in gigabytes. Read-only.
    sizeInGB *int32
    // The SKU name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
    sku *string
    // The official display stock keeping unit (SKU) name of this gallery image. For example, 2004. Read-only.
    skuDisplayName *string
    // The date when the image becomes available. Read-only.
    startDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The status of the gallery image on the Cloud PC. Possible values are: supported, supportedWithWarning, notSupported, unknownFutureValue. Read-only.
    status *CloudPcGalleryImageStatus
}
// NewCloudPcGalleryImage instantiates a new CloudPcGalleryImage and sets the default values.
func NewCloudPcGalleryImage()(*CloudPcGalleryImage) {
    m := &CloudPcGalleryImage{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.cloudPcGalleryImage";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateCloudPcGalleryImageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcGalleryImageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcGalleryImage(), nil
}
// GetDisplayName gets the displayName property value. The official display name of the gallery image. Read-only.
func (m *CloudPcGalleryImage) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndDate gets the endDate property value. The date in which this image is no longer within long-term support. The Cloud PC will continue to provide short-term support. Read-only.
func (m *CloudPcGalleryImage) GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endDate
}
// GetExpirationDate gets the expirationDate property value. The date when the image is no longer available. Read-only.
func (m *CloudPcGalleryImage) GetExpirationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.expirationDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcGalleryImage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["endDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetEndDate)
    res["expirationDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetExpirationDate)
    res["offer"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOffer)
    res["offerDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOfferDisplayName)
    res["publisher"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetPublisher)
    res["recommendedSku"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRecommendedSku)
    res["sizeInGB"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetSizeInGB)
    res["sku"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSku)
    res["skuDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSkuDisplayName)
    res["startDate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetDateOnlyValue(m.SetStartDate)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseCloudPcGalleryImageStatus , m.SetStatus)
    return res
}
// GetOffer gets the offer property value. The offer name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) GetOffer()(*string) {
    return m.offer
}
// GetOfferDisplayName gets the offerDisplayName property value. The official display offer name of the gallery image. For example, Windows 10 Enterprise + OS Optimizations. Read-only.
func (m *CloudPcGalleryImage) GetOfferDisplayName()(*string) {
    return m.offerDisplayName
}
// GetPublisher gets the publisher property value. The publisher name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) GetPublisher()(*string) {
    return m.publisher
}
// GetRecommendedSku gets the recommendedSku property value. Recommended Cloud PC SKU for this gallery image. Read-only.
func (m *CloudPcGalleryImage) GetRecommendedSku()(*string) {
    return m.recommendedSku
}
// GetSizeInGB gets the sizeInGB property value. The size of this image in gigabytes. Read-only.
func (m *CloudPcGalleryImage) GetSizeInGB()(*int32) {
    return m.sizeInGB
}
// GetSku gets the sku property value. The SKU name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) GetSku()(*string) {
    return m.sku
}
// GetSkuDisplayName gets the skuDisplayName property value. The official display stock keeping unit (SKU) name of this gallery image. For example, 2004. Read-only.
func (m *CloudPcGalleryImage) GetSkuDisplayName()(*string) {
    return m.skuDisplayName
}
// GetStartDate gets the startDate property value. The date when the image becomes available. Read-only.
func (m *CloudPcGalleryImage) GetStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.startDate
}
// GetStatus gets the status property value. The status of the gallery image on the Cloud PC. Possible values are: supported, supportedWithWarning, notSupported, unknownFutureValue. Read-only.
func (m *CloudPcGalleryImage) GetStatus()(*CloudPcGalleryImageStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *CloudPcGalleryImage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("expirationDate", m.GetExpirationDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("offer", m.GetOffer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("offerDisplayName", m.GetOfferDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recommendedSku", m.GetRecommendedSku())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("sizeInGB", m.GetSizeInGB())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sku", m.GetSku())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skuDisplayName", m.GetSkuDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("startDate", m.GetStartDate())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The official display name of the gallery image. Read-only.
func (m *CloudPcGalleryImage) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndDate sets the endDate property value. The date in which this image is no longer within long-term support. The Cloud PC will continue to provide short-term support. Read-only.
func (m *CloudPcGalleryImage) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endDate = value
}
// SetExpirationDate sets the expirationDate property value. The date when the image is no longer available. Read-only.
func (m *CloudPcGalleryImage) SetExpirationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.expirationDate = value
}
// SetOffer sets the offer property value. The offer name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) SetOffer(value *string)() {
    m.offer = value
}
// SetOfferDisplayName sets the offerDisplayName property value. The official display offer name of the gallery image. For example, Windows 10 Enterprise + OS Optimizations. Read-only.
func (m *CloudPcGalleryImage) SetOfferDisplayName(value *string)() {
    m.offerDisplayName = value
}
// SetPublisher sets the publisher property value. The publisher name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) SetPublisher(value *string)() {
    m.publisher = value
}
// SetRecommendedSku sets the recommendedSku property value. Recommended Cloud PC SKU for this gallery image. Read-only.
func (m *CloudPcGalleryImage) SetRecommendedSku(value *string)() {
    m.recommendedSku = value
}
// SetSizeInGB sets the sizeInGB property value. The size of this image in gigabytes. Read-only.
func (m *CloudPcGalleryImage) SetSizeInGB(value *int32)() {
    m.sizeInGB = value
}
// SetSku sets the sku property value. The SKU name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) SetSku(value *string)() {
    m.sku = value
}
// SetSkuDisplayName sets the skuDisplayName property value. The official display stock keeping unit (SKU) name of this gallery image. For example, 2004. Read-only.
func (m *CloudPcGalleryImage) SetSkuDisplayName(value *string)() {
    m.skuDisplayName = value
}
// SetStartDate sets the startDate property value. The date when the image becomes available. Read-only.
func (m *CloudPcGalleryImage) SetStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.startDate = value
}
// SetStatus sets the status property value. The status of the gallery image on the Cloud PC. Possible values are: supported, supportedWithWarning, notSupported, unknownFutureValue. Read-only.
func (m *CloudPcGalleryImage) SetStatus(value *CloudPcGalleryImageStatus)() {
    m.status = value
}
