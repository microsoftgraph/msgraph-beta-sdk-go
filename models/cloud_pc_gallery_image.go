package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudPcGalleryImage 
type CloudPcGalleryImage struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewCloudPcGalleryImage instantiates a new cloudPcGalleryImage and sets the default values.
func NewCloudPcGalleryImage()(*CloudPcGalleryImage) {
    m := &CloudPcGalleryImage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudPcGalleryImageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCloudPcGalleryImageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudPcGalleryImage(), nil
}
// GetDisplayName gets the displayName property value. The official display name of the gallery image. Read-only.
func (m *CloudPcGalleryImage) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEndDate gets the endDate property value. The date in which this image is no longer within long-term support. The Cloud PC continues to provide short-term support. Read-only.
func (m *CloudPcGalleryImage) GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("endDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetExpirationDate gets the expirationDate property value. The date when the image is no longer available. Read-only.
func (m *CloudPcGalleryImage) GetExpirationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("expirationDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcGalleryImage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["expirationDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDate(val)
        }
        return nil
    }
    res["offer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffer(val)
        }
        return nil
    }
    res["offerDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfferDisplayName(val)
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["recommendedSku"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedSku(val)
        }
        return nil
    }
    res["sizeInGB"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInGB(val)
        }
        return nil
    }
    res["sku"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSku(val)
        }
        return nil
    }
    res["skuDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuDisplayName(val)
        }
        return nil
    }
    res["startDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDate(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcGalleryImageStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CloudPcGalleryImageStatus))
        }
        return nil
    }
    return res
}
// GetOffer gets the offer property value. The offer name of the gallery image. This value is passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) GetOffer()(*string) {
    val, err := m.GetBackingStore().Get("offer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOfferDisplayName gets the offerDisplayName property value. The official display offer name of the gallery image. For example, Windows 10 Enterprise + OS Optimizations. Read-only.
func (m *CloudPcGalleryImage) GetOfferDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("offerDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPublisher gets the publisher property value. The publisher name of the gallery image. This value is passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) GetPublisher()(*string) {
    val, err := m.GetBackingStore().Get("publisher")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecommendedSku gets the recommendedSku property value. Recommended Cloud PC SKU for this gallery image. Read-only.
func (m *CloudPcGalleryImage) GetRecommendedSku()(*string) {
    val, err := m.GetBackingStore().Get("recommendedSku")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSizeInGB gets the sizeInGB property value. The size of this image in gigabytes. Read-only.
func (m *CloudPcGalleryImage) GetSizeInGB()(*int32) {
    val, err := m.GetBackingStore().Get("sizeInGB")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetSku gets the sku property value. The SKU name of the gallery image. This value is passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) GetSku()(*string) {
    val, err := m.GetBackingStore().Get("sku")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSkuDisplayName gets the skuDisplayName property value. The official display stock keeping unit (SKU) name of this gallery image. For example, 2004. Read-only.
func (m *CloudPcGalleryImage) GetSkuDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("skuDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStartDate gets the startDate property value. The date when the image becomes available. Read-only.
func (m *CloudPcGalleryImage) GetStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("startDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetStatus gets the status property value. The status of the gallery image on the Cloud PC. Possible values are: supported, supportedWithWarning, notSupported, unknownFutureValue. Read-only.
func (m *CloudPcGalleryImage) GetStatus()(*CloudPcGalleryImageStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudPcGalleryImageStatus)
    }
    return nil
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
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEndDate sets the endDate property value. The date in which this image is no longer within long-term support. The Cloud PC continues to provide short-term support. Read-only.
func (m *CloudPcGalleryImage) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("endDate", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDate sets the expirationDate property value. The date when the image is no longer available. Read-only.
func (m *CloudPcGalleryImage) SetExpirationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("expirationDate", value)
    if err != nil {
        panic(err)
    }
}
// SetOffer sets the offer property value. The offer name of the gallery image. This value is passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) SetOffer(value *string)() {
    err := m.GetBackingStore().Set("offer", value)
    if err != nil {
        panic(err)
    }
}
// SetOfferDisplayName sets the offerDisplayName property value. The official display offer name of the gallery image. For example, Windows 10 Enterprise + OS Optimizations. Read-only.
func (m *CloudPcGalleryImage) SetOfferDisplayName(value *string)() {
    err := m.GetBackingStore().Set("offerDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetPublisher sets the publisher property value. The publisher name of the gallery image. This value is passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) SetPublisher(value *string)() {
    err := m.GetBackingStore().Set("publisher", value)
    if err != nil {
        panic(err)
    }
}
// SetRecommendedSku sets the recommendedSku property value. Recommended Cloud PC SKU for this gallery image. Read-only.
func (m *CloudPcGalleryImage) SetRecommendedSku(value *string)() {
    err := m.GetBackingStore().Set("recommendedSku", value)
    if err != nil {
        panic(err)
    }
}
// SetSizeInGB sets the sizeInGB property value. The size of this image in gigabytes. Read-only.
func (m *CloudPcGalleryImage) SetSizeInGB(value *int32)() {
    err := m.GetBackingStore().Set("sizeInGB", value)
    if err != nil {
        panic(err)
    }
}
// SetSku sets the sku property value. The SKU name of the gallery image. This value is passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) SetSku(value *string)() {
    err := m.GetBackingStore().Set("sku", value)
    if err != nil {
        panic(err)
    }
}
// SetSkuDisplayName sets the skuDisplayName property value. The official display stock keeping unit (SKU) name of this gallery image. For example, 2004. Read-only.
func (m *CloudPcGalleryImage) SetSkuDisplayName(value *string)() {
    err := m.GetBackingStore().Set("skuDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDate sets the startDate property value. The date when the image becomes available. Read-only.
func (m *CloudPcGalleryImage) SetStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("startDate", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status of the gallery image on the Cloud PC. Possible values are: supported, supportedWithWarning, notSupported, unknownFutureValue. Read-only.
func (m *CloudPcGalleryImage) SetStatus(value *CloudPcGalleryImageStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// CloudPcGalleryImageable 
type CloudPcGalleryImageable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetExpirationDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetOffer()(*string)
    GetOfferDisplayName()(*string)
    GetPublisher()(*string)
    GetRecommendedSku()(*string)
    GetSizeInGB()(*int32)
    GetSku()(*string)
    GetSkuDisplayName()(*string)
    GetStartDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetStatus()(*CloudPcGalleryImageStatus)
    SetDisplayName(value *string)()
    SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetExpirationDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetOffer(value *string)()
    SetOfferDisplayName(value *string)()
    SetPublisher(value *string)()
    SetRecommendedSku(value *string)()
    SetSizeInGB(value *int32)()
    SetSku(value *string)()
    SetSkuDisplayName(value *string)()
    SetStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetStatus(value *CloudPcGalleryImageStatus)()
}
