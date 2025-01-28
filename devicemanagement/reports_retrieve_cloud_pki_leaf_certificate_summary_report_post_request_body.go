package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody instantiates a new ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody and sets the default values.
func NewReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody()(*ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody) {
    m := &ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCertificationAuthorityId gets the certificationAuthorityId property value. The certificationAuthorityId property
// returns a *string when successful
func (m *ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody) GetCertificationAuthorityId()(*string) {
    val, err := m.GetBackingStore().Get("certificationAuthorityId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificationAuthorityId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthorityId(val)
        }
        return nil
    }
    res["select"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSelectEscaped(res)
        }
        return nil
    }
    return res
}
// GetSelectEscaped gets the select property value. The select property
// returns a []string when successful
func (m *ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody) GetSelectEscaped()([]string) {
    val, err := m.GetBackingStore().Get("selectEscaped")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("certificationAuthorityId", m.GetCertificationAuthorityId())
        if err != nil {
            return err
        }
    }
    if m.GetSelectEscaped() != nil {
        err := writer.WriteCollectionOfStringValues("select", m.GetSelectEscaped())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCertificationAuthorityId sets the certificationAuthorityId property value. The certificationAuthorityId property
func (m *ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody) SetCertificationAuthorityId(value *string)() {
    err := m.GetBackingStore().Set("certificationAuthorityId", value)
    if err != nil {
        panic(err)
    }
}
// SetSelectEscaped sets the select property value. The select property
func (m *ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBody) SetSelectEscaped(value []string)() {
    err := m.GetBackingStore().Set("selectEscaped", value)
    if err != nil {
        panic(err)
    }
}
type ReportsRetrieveCloudPkiLeafCertificateSummaryReportPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCertificationAuthorityId()(*string)
    GetSelectEscaped()([]string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCertificationAuthorityId(value *string)()
    SetSelectEscaped(value []string)()
}
