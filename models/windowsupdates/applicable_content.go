package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicableContent 
type ApplicableContent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The catalogEntry property
    catalogEntry CatalogEntryable
    // Collection of devices and recommendations for applicable catalog content.
    matchedDevices []ApplicableContentDeviceMatchable
    // The OdataType property
    odataType *string
}
// NewApplicableContent instantiates a new applicableContent and sets the default values.
func NewApplicableContent()(*ApplicableContent) {
    m := &ApplicableContent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateApplicableContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicableContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplicableContent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApplicableContent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCatalogEntry gets the catalogEntry property value. The catalogEntry property
func (m *ApplicableContent) GetCatalogEntry()(CatalogEntryable) {
    return m.catalogEntry
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicableContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["catalogEntry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCatalogEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalogEntry(val.(CatalogEntryable))
        }
        return nil
    }
    res["matchedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApplicableContentDeviceMatchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApplicableContentDeviceMatchable, len(val))
            for i, v := range val {
                res[i] = v.(ApplicableContentDeviceMatchable)
            }
            m.SetMatchedDevices(res)
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
// GetMatchedDevices gets the matchedDevices property value. Collection of devices and recommendations for applicable catalog content.
func (m *ApplicableContent) GetMatchedDevices()([]ApplicableContentDeviceMatchable) {
    return m.matchedDevices
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ApplicableContent) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ApplicableContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("catalogEntry", m.GetCatalogEntry())
        if err != nil {
            return err
        }
    }
    if m.GetMatchedDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMatchedDevices()))
        for i, v := range m.GetMatchedDevices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("matchedDevices", cast)
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
func (m *ApplicableContent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCatalogEntry sets the catalogEntry property value. The catalogEntry property
func (m *ApplicableContent) SetCatalogEntry(value CatalogEntryable)() {
    m.catalogEntry = value
}
// SetMatchedDevices sets the matchedDevices property value. Collection of devices and recommendations for applicable catalog content.
func (m *ApplicableContent) SetMatchedDevices(value []ApplicableContentDeviceMatchable)() {
    m.matchedDevices = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ApplicableContent) SetOdataType(value *string)() {
    m.odataType = value
}
