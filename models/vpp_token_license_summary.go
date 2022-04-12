package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VppTokenLicenseSummary license summary of a given app in a token.
type VppTokenLicenseSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The Apple Id associated with the given Apple Volume Purchase Program Token.
    appleId *string
    // The number of VPP licenses available.
    availableLicenseCount *int32
    // The organization associated with the Apple Volume Purchase Program Token.
    organizationName *string
    // The number of VPP licenses in use.
    usedLicenseCount *int32
    // Identifier of the VPP token.
    vppTokenId *string
}
// NewVppTokenLicenseSummary instantiates a new vppTokenLicenseSummary and sets the default values.
func NewVppTokenLicenseSummary()(*VppTokenLicenseSummary) {
    m := &VppTokenLicenseSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVppTokenLicenseSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVppTokenLicenseSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVppTokenLicenseSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VppTokenLicenseSummary) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppleId gets the appleId property value. The Apple Id associated with the given Apple Volume Purchase Program Token.
func (m *VppTokenLicenseSummary) GetAppleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appleId
    }
}
// GetAvailableLicenseCount gets the availableLicenseCount property value. The number of VPP licenses available.
func (m *VppTokenLicenseSummary) GetAvailableLicenseCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.availableLicenseCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VppTokenLicenseSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppleId(val)
        }
        return nil
    }
    res["availableLicenseCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailableLicenseCount(val)
        }
        return nil
    }
    res["organizationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationName(val)
        }
        return nil
    }
    res["usedLicenseCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsedLicenseCount(val)
        }
        return nil
    }
    res["vppTokenId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVppTokenId(val)
        }
        return nil
    }
    return res
}
// GetOrganizationName gets the organizationName property value. The organization associated with the Apple Volume Purchase Program Token.
func (m *VppTokenLicenseSummary) GetOrganizationName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.organizationName
    }
}
// GetUsedLicenseCount gets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *VppTokenLicenseSummary) GetUsedLicenseCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.usedLicenseCount
    }
}
// GetVppTokenId gets the vppTokenId property value. Identifier of the VPP token.
func (m *VppTokenLicenseSummary) GetVppTokenId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.vppTokenId
    }
}
// Serialize serializes information the current object
func (m *VppTokenLicenseSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appleId", m.GetAppleId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("availableLicenseCount", m.GetAvailableLicenseCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("organizationName", m.GetOrganizationName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("usedLicenseCount", m.GetUsedLicenseCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("vppTokenId", m.GetVppTokenId())
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
func (m *VppTokenLicenseSummary) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppleId sets the appleId property value. The Apple Id associated with the given Apple Volume Purchase Program Token.
func (m *VppTokenLicenseSummary) SetAppleId(value *string)() {
    if m != nil {
        m.appleId = value
    }
}
// SetAvailableLicenseCount sets the availableLicenseCount property value. The number of VPP licenses available.
func (m *VppTokenLicenseSummary) SetAvailableLicenseCount(value *int32)() {
    if m != nil {
        m.availableLicenseCount = value
    }
}
// SetOrganizationName sets the organizationName property value. The organization associated with the Apple Volume Purchase Program Token.
func (m *VppTokenLicenseSummary) SetOrganizationName(value *string)() {
    if m != nil {
        m.organizationName = value
    }
}
// SetUsedLicenseCount sets the usedLicenseCount property value. The number of VPP licenses in use.
func (m *VppTokenLicenseSummary) SetUsedLicenseCount(value *int32)() {
    if m != nil {
        m.usedLicenseCount = value
    }
}
// SetVppTokenId sets the vppTokenId property value. Identifier of the VPP token.
func (m *VppTokenLicenseSummary) SetVppTokenId(value *string)() {
    if m != nil {
        m.vppTokenId = value
    }
}
