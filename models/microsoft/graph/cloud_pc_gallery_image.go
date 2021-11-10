package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPcGalleryImage struct {
    Entity
    // 
    displayName *string;
    // 
    endDate *string;
    // 
    expirationDate *string;
    // 
    offer *string;
    // 
    offerDisplayName *string;
    // 
    publisher *string;
    // 
    recommendedSku *string;
    // 
    sizeInGB *int32;
    // 
    sku *string;
    // 
    skuDisplayName *string;
    // 
    startDate *string;
    // 
    status *CloudPcGalleryImageStatus;
}
// Instantiates a new cloudPcGalleryImage and sets the default values.
func NewCloudPcGalleryImage()(*CloudPcGalleryImage) {
    m := &CloudPcGalleryImage{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. 
func (m *CloudPcGalleryImage) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the endDate property value. 
func (m *CloudPcGalleryImage) GetEndDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endDate
    }
}
// Gets the expirationDate property value. 
func (m *CloudPcGalleryImage) GetExpirationDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.expirationDate
    }
}
// Gets the offer property value. 
func (m *CloudPcGalleryImage) GetOffer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.offer
    }
}
// Gets the offerDisplayName property value. 
func (m *CloudPcGalleryImage) GetOfferDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.offerDisplayName
    }
}
// Gets the publisher property value. 
func (m *CloudPcGalleryImage) GetPublisher()(*string) {
    if m == nil {
        return nil
    } else {
        return m.publisher
    }
}
// Gets the recommendedSku property value. 
func (m *CloudPcGalleryImage) GetRecommendedSku()(*string) {
    if m == nil {
        return nil
    } else {
        return m.recommendedSku
    }
}
// Gets the sizeInGB property value. 
func (m *CloudPcGalleryImage) GetSizeInGB()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.sizeInGB
    }
}
// Gets the sku property value. 
func (m *CloudPcGalleryImage) GetSku()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sku
    }
}
// Gets the skuDisplayName property value. 
func (m *CloudPcGalleryImage) GetSkuDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuDisplayName
    }
}
// Gets the startDate property value. 
func (m *CloudPcGalleryImage) GetStartDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startDate
    }
}
// Gets the status property value. 
func (m *CloudPcGalleryImage) GetStatus()(*CloudPcGalleryImageStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["expirationDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
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
        val, err := n.GetStringValue()
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
            cast := val.(CloudPcGalleryImageStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *CloudPcGalleryImage) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
        err = writer.WriteStringValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("expirationDate", m.GetExpirationDate())
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
        err = writer.WriteStringValue("startDate", m.GetStartDate())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *CloudPcGalleryImage) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the endDate property value. 
// Parameters:
//  - value : Value to set for the endDate property.
func (m *CloudPcGalleryImage) SetEndDate(value *string)() {
    m.endDate = value
}
// Sets the expirationDate property value. 
// Parameters:
//  - value : Value to set for the expirationDate property.
func (m *CloudPcGalleryImage) SetExpirationDate(value *string)() {
    m.expirationDate = value
}
// Sets the offer property value. 
// Parameters:
//  - value : Value to set for the offer property.
func (m *CloudPcGalleryImage) SetOffer(value *string)() {
    m.offer = value
}
// Sets the offerDisplayName property value. 
// Parameters:
//  - value : Value to set for the offerDisplayName property.
func (m *CloudPcGalleryImage) SetOfferDisplayName(value *string)() {
    m.offerDisplayName = value
}
// Sets the publisher property value. 
// Parameters:
//  - value : Value to set for the publisher property.
func (m *CloudPcGalleryImage) SetPublisher(value *string)() {
    m.publisher = value
}
// Sets the recommendedSku property value. 
// Parameters:
//  - value : Value to set for the recommendedSku property.
func (m *CloudPcGalleryImage) SetRecommendedSku(value *string)() {
    m.recommendedSku = value
}
// Sets the sizeInGB property value. 
// Parameters:
//  - value : Value to set for the sizeInGB property.
func (m *CloudPcGalleryImage) SetSizeInGB(value *int32)() {
    m.sizeInGB = value
}
// Sets the sku property value. 
// Parameters:
//  - value : Value to set for the sku property.
func (m *CloudPcGalleryImage) SetSku(value *string)() {
    m.sku = value
}
// Sets the skuDisplayName property value. 
// Parameters:
//  - value : Value to set for the skuDisplayName property.
func (m *CloudPcGalleryImage) SetSkuDisplayName(value *string)() {
    m.skuDisplayName = value
}
// Sets the startDate property value. 
// Parameters:
//  - value : Value to set for the startDate property.
func (m *CloudPcGalleryImage) SetStartDate(value *string)() {
    m.startDate = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *CloudPcGalleryImage) SetStatus(value *CloudPcGalleryImageStatus)() {
    m.status = value
}
