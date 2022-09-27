package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserPrint 
type UserPrint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // The recentPrinterShares property
    recentPrinterShares []PrinterShareable
}
// NewUserPrint instantiates a new userPrint and sets the default values.
func NewUserPrint()(*UserPrint) {
    m := &UserPrint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.userPrint";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserPrintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserPrintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserPrint(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserPrint) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserPrint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["recentPrinterShares"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreatePrinterShareFromDiscriminatorValue , m.SetRecentPrinterShares)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserPrint) GetOdataType()(*string) {
    return m.odataType
}
// GetRecentPrinterShares gets the recentPrinterShares property value. The recentPrinterShares property
func (m *UserPrint) GetRecentPrinterShares()([]PrinterShareable) {
    return m.recentPrinterShares
}
// Serialize serializes information the current object
func (m *UserPrint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRecentPrinterShares() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetRecentPrinterShares())
        err := writer.WriteCollectionOfObjectValues("recentPrinterShares", cast)
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
func (m *UserPrint) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserPrint) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecentPrinterShares sets the recentPrinterShares property value. The recentPrinterShares property
func (m *UserPrint) SetRecentPrinterShares(value []PrinterShareable)() {
    m.recentPrinterShares = value
}
