package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Financials 
type Financials struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The companies property
    companies []Companyable
    // The OdataType property
    odataType *string
}
// NewFinancials instantiates a new Financials and sets the default values.
func NewFinancials()(*Financials) {
    m := &Financials{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.financials";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateFinancialsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFinancialsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFinancials(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Financials) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCompanies gets the companies property value. The companies property
func (m *Financials) GetCompanies()([]Companyable) {
    return m.companies
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Financials) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["companies"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateCompanyFromDiscriminatorValue , m.SetCompanies)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Financials) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *Financials) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCompanies() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetCompanies())
        err := writer.WriteCollectionOfObjectValues("companies", cast)
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
func (m *Financials) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCompanies sets the companies property value. The companies property
func (m *Financials) SetCompanies(value []Companyable)() {
    m.companies = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Financials) SetOdataType(value *string)() {
    m.odataType = value
}
