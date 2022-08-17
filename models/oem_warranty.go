package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OemWarranty oEM Warranty information for a given device
type OemWarranty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // List of additional warranty offers. This collection can contain a maximum of 100 elements.
    additionalWarranties []WarrantyOfferable
    // List of base warranty offers. This collection can contain a maximum of 100 elements.
    baseWarranties []WarrantyOfferable
    // Device configuration page URL
    deviceConfigurationUrl *string
    // Device warranty page URL
    deviceWarrantyUrl *string
    // The OdataType property
    odataType *string
}
// NewOemWarranty instantiates a new oemWarranty and sets the default values.
func NewOemWarranty()(*OemWarranty) {
    m := &OemWarranty{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.oemWarranty";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOemWarrantyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOemWarrantyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOemWarranty(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OemWarranty) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAdditionalWarranties gets the additionalWarranties property value. List of additional warranty offers. This collection can contain a maximum of 100 elements.
func (m *OemWarranty) GetAdditionalWarranties()([]WarrantyOfferable) {
    return m.additionalWarranties
}
// GetBaseWarranties gets the baseWarranties property value. List of base warranty offers. This collection can contain a maximum of 100 elements.
func (m *OemWarranty) GetBaseWarranties()([]WarrantyOfferable) {
    return m.baseWarranties
}
// GetDeviceConfigurationUrl gets the deviceConfigurationUrl property value. Device configuration page URL
func (m *OemWarranty) GetDeviceConfigurationUrl()(*string) {
    return m.deviceConfigurationUrl
}
// GetDeviceWarrantyUrl gets the deviceWarrantyUrl property value. Device warranty page URL
func (m *OemWarranty) GetDeviceWarrantyUrl()(*string) {
    return m.deviceWarrantyUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OemWarranty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalWarranties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWarrantyOfferFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WarrantyOfferable, len(val))
            for i, v := range val {
                res[i] = v.(WarrantyOfferable)
            }
            m.SetAdditionalWarranties(res)
        }
        return nil
    }
    res["baseWarranties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWarrantyOfferFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WarrantyOfferable, len(val))
            for i, v := range val {
                res[i] = v.(WarrantyOfferable)
            }
            m.SetBaseWarranties(res)
        }
        return nil
    }
    res["deviceConfigurationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceConfigurationUrl(val)
        }
        return nil
    }
    res["deviceWarrantyUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceWarrantyUrl(val)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OemWarranty) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *OemWarranty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAdditionalWarranties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdditionalWarranties()))
        for i, v := range m.GetAdditionalWarranties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("additionalWarranties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBaseWarranties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBaseWarranties()))
        for i, v := range m.GetBaseWarranties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("baseWarranties", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceConfigurationUrl", m.GetDeviceConfigurationUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceWarrantyUrl", m.GetDeviceWarrantyUrl())
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
func (m *OemWarranty) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAdditionalWarranties sets the additionalWarranties property value. List of additional warranty offers. This collection can contain a maximum of 100 elements.
func (m *OemWarranty) SetAdditionalWarranties(value []WarrantyOfferable)() {
    m.additionalWarranties = value
}
// SetBaseWarranties sets the baseWarranties property value. List of base warranty offers. This collection can contain a maximum of 100 elements.
func (m *OemWarranty) SetBaseWarranties(value []WarrantyOfferable)() {
    m.baseWarranties = value
}
// SetDeviceConfigurationUrl sets the deviceConfigurationUrl property value. Device configuration page URL
func (m *OemWarranty) SetDeviceConfigurationUrl(value *string)() {
    m.deviceConfigurationUrl = value
}
// SetDeviceWarrantyUrl sets the deviceWarrantyUrl property value. Device warranty page URL
func (m *OemWarranty) SetDeviceWarrantyUrl(value *string)() {
    m.deviceWarrantyUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OemWarranty) SetOdataType(value *string)() {
    m.odataType = value
}
