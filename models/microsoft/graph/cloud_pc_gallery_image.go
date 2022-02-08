package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudPcGalleryImage 
type CloudPcGalleryImage struct {
    Entity
    // The official display name of the gallery image. Read-only.
    displayName *string;
    // The date in which this image is no longer within long-term support. The Cloud PC will continue to provide short-term support. Read-only.
    endDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The date when the image is no longer available. Read-only.
    expirationDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The offer name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
    offer *string;
    // The official display offer name of the gallery image. For example, Windows 10 Enterprise + OS Optimizations. Read-only.
    offerDisplayName *string;
    // The publisher name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
    publisher *string;
    // Recommended Cloud PC SKU for this gallery image. Read-only.
    recommendedSku *string;
    // The size of this image in gigabytes. Read-only.
    sizeInGB *int32;
    // The SKU name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
    sku *string;
    // The official display stock keeping unit (SKU) name of this gallery image. For example, 2004. Read-only.
    skuDisplayName *string;
    // The date when the image becomes available. Read-only.
    startDate *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly;
    // The status of the gallery image on the Cloud PC. Possible values are: supported, supportedWithWarning, notSupported, unknownFutureValue. Read-only.
    status *CloudPcGalleryImageStatus;
}
// NewCloudPcGalleryImage instantiates a new cloudPcGalleryImage and sets the default values.
func NewCloudPcGalleryImage()(*CloudPcGalleryImage) {
    m := &CloudPcGalleryImage{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. The official display name of the gallery image. Read-only.
func (m *CloudPcGalleryImage) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEndDate gets the endDate property value. The date in which this image is no longer within long-term support. The Cloud PC will continue to provide short-term support. Read-only.
func (m *CloudPcGalleryImage) GetEndDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.endDate
    }
}
// GetExpirationDate gets the expirationDate property value. The date when the image is no longer available. Read-only.
func (m *CloudPcGalleryImage) GetExpirationDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.expirationDate
    }
}
// GetOffer gets the offer property value. The offer name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) GetOffer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.offer
    }
}
// GetOfferDisplayName gets the offerDisplayName property value. The official display offer name of the gallery image. For example, Windows 10 Enterprise + OS Optimizations. Read-only.
func (m *CloudPcGalleryImage) GetOfferDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.offerDisplayName
    }
}
// GetPublisher gets the publisher property value. The publisher name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// GetRecommendedSku gets the recommendedSku property value. Recommended Cloud PC SKU for this gallery image. Read-only.
func (m *CloudPcGalleryImage) GetRecommendedSku()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recommendedSku
    }
}
// GetSizeInGB gets the sizeInGB property value. The size of this image in gigabytes. Read-only.
func (m *CloudPcGalleryImage) GetSizeInGB()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sizeInGB
    }
}
// GetSku gets the sku property value. The SKU name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) GetSku()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sku
    }
}
// GetSkuDisplayName gets the skuDisplayName property value. The official display stock keeping unit (SKU) name of this gallery image. For example, 2004. Read-only.
func (m *CloudPcGalleryImage) GetSkuDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuDisplayName
    }
}
// GetStartDate gets the startDate property value. The date when the image becomes available. Read-only.
func (m *CloudPcGalleryImage) GetStartDate()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly) {
    if m == nil {
        return nil
    } else {
        return m.startDate
    }
}
// GetStatus gets the status property value. The status of the gallery image on the Cloud PC. Possible values are: supported, supportedWithWarning, notSupported, unknownFutureValue. Read-only.
func (m *CloudPcGalleryImage) GetStatus()(*CloudPcGalleryImageStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudPcGalleryImage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["expirationDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDate(val)
        }
        return nil
    }
    res["offer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOffer(val)
        }
        return nil
    }
    res["offerDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfferDisplayName(val)
        }
        return nil
    }
    res["publisher"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["recommendedSku"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendedSku(val)
        }
        return nil
    }
    res["sizeInGB"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInGB(val)
        }
        return nil
    }
    res["sku"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSku(val)
        }
        return nil
    }
    res["skuDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuDisplayName(val)
        }
        return nil
    }
    res["startDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDate(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *CloudPcGalleryImage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudPcGalleryImage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m != nil {
        m.displayName = value
    }
}
// SetEndDate sets the endDate property value. The date in which this image is no longer within long-term support. The Cloud PC will continue to provide short-term support. Read-only.
func (m *CloudPcGalleryImage) SetEndDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.endDate = value
    }
}
// SetExpirationDate sets the expirationDate property value. The date when the image is no longer available. Read-only.
func (m *CloudPcGalleryImage) SetExpirationDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.expirationDate = value
    }
}
// SetOffer sets the offer property value. The offer name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) SetOffer(value *string)() {
    if m != nil {
        m.offer = value
    }
}
// SetOfferDisplayName sets the offerDisplayName property value. The official display offer name of the gallery image. For example, Windows 10 Enterprise + OS Optimizations. Read-only.
func (m *CloudPcGalleryImage) SetOfferDisplayName(value *string)() {
    if m != nil {
        m.offerDisplayName = value
    }
}
// SetPublisher sets the publisher property value. The publisher name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) SetPublisher(value *string)() {
    if m != nil {
        m.publisher = value
    }
}
// SetRecommendedSku sets the recommendedSku property value. Recommended Cloud PC SKU for this gallery image. Read-only.
func (m *CloudPcGalleryImage) SetRecommendedSku(value *string)() {
    if m != nil {
        m.recommendedSku = value
    }
}
// SetSizeInGB sets the sizeInGB property value. The size of this image in gigabytes. Read-only.
func (m *CloudPcGalleryImage) SetSizeInGB(value *int32)() {
    if m != nil {
        m.sizeInGB = value
    }
}
// SetSku sets the sku property value. The SKU name of the gallery image. This value will be passed to Azure to get the image resource. Read-only.
func (m *CloudPcGalleryImage) SetSku(value *string)() {
    if m != nil {
        m.sku = value
    }
}
// SetSkuDisplayName sets the skuDisplayName property value. The official display stock keeping unit (SKU) name of this gallery image. For example, 2004. Read-only.
func (m *CloudPcGalleryImage) SetSkuDisplayName(value *string)() {
    if m != nil {
        m.skuDisplayName = value
    }
}
// SetStartDate sets the startDate property value. The date when the image becomes available. Read-only.
func (m *CloudPcGalleryImage) SetStartDate(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)() {
    if m != nil {
        m.startDate = value
    }
}
// SetStatus sets the status property value. The status of the gallery image on the Cloud PC. Possible values are: supported, supportedWithWarning, notSupported, unknownFutureValue. Read-only.
func (m *CloudPcGalleryImage) SetStatus(value *CloudPcGalleryImageStatus)() {
    if m != nil {
        m.status = value
    }
}
