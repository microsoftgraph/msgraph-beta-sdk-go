package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalMessageLogoDimensions contains the required size dimensions of a logo
type OrganizationalMessageLogoDimensions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Maximum height of the logo
    maxHeight *int32
    // Maximum width of the logo
    maxWidth *int32
    // Minimum height of the logo
    minHeight *int32
    // Minimum width of the logo
    minWidth *int32
    // The OdataType property
    odataType *string
}
// NewOrganizationalMessageLogoDimensions instantiates a new organizationalMessageLogoDimensions and sets the default values.
func NewOrganizationalMessageLogoDimensions()(*OrganizationalMessageLogoDimensions) {
    m := &OrganizationalMessageLogoDimensions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.organizationalMessageLogoDimensions";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateOrganizationalMessageLogoDimensionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationalMessageLogoDimensionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationalMessageLogoDimensions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrganizationalMessageLogoDimensions) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalMessageLogoDimensions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maxHeight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxHeight(val)
        }
        return nil
    }
    res["maxWidth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxWidth(val)
        }
        return nil
    }
    res["minHeight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinHeight(val)
        }
        return nil
    }
    res["minWidth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinWidth(val)
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
// GetMaxHeight gets the maxHeight property value. Maximum height of the logo
func (m *OrganizationalMessageLogoDimensions) GetMaxHeight()(*int32) {
    return m.maxHeight
}
// GetMaxWidth gets the maxWidth property value. Maximum width of the logo
func (m *OrganizationalMessageLogoDimensions) GetMaxWidth()(*int32) {
    return m.maxWidth
}
// GetMinHeight gets the minHeight property value. Minimum height of the logo
func (m *OrganizationalMessageLogoDimensions) GetMinHeight()(*int32) {
    return m.minHeight
}
// GetMinWidth gets the minWidth property value. Minimum width of the logo
func (m *OrganizationalMessageLogoDimensions) GetMinWidth()(*int32) {
    return m.minWidth
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageLogoDimensions) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *OrganizationalMessageLogoDimensions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("maxHeight", m.GetMaxHeight())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxWidth", m.GetMaxWidth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minHeight", m.GetMinHeight())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minWidth", m.GetMinWidth())
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
func (m *OrganizationalMessageLogoDimensions) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetMaxHeight sets the maxHeight property value. Maximum height of the logo
func (m *OrganizationalMessageLogoDimensions) SetMaxHeight(value *int32)() {
    m.maxHeight = value
}
// SetMaxWidth sets the maxWidth property value. Maximum width of the logo
func (m *OrganizationalMessageLogoDimensions) SetMaxWidth(value *int32)() {
    m.maxWidth = value
}
// SetMinHeight sets the minHeight property value. Minimum height of the logo
func (m *OrganizationalMessageLogoDimensions) SetMinHeight(value *int32)() {
    m.minHeight = value
}
// SetMinWidth sets the minWidth property value. Minimum width of the logo
func (m *OrganizationalMessageLogoDimensions) SetMinWidth(value *int32)() {
    m.minWidth = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OrganizationalMessageLogoDimensions) SetOdataType(value *string)() {
    m.odataType = value
}
