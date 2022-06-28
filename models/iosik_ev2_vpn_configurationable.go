package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosikEv2VpnConfigurationable 
type IosikEv2VpnConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    IosVpnConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowDefaultChildSecurityAssociationParameters()(*bool)
    GetAllowDefaultSecurityAssociationParameters()(*bool)
    GetAlwaysOnConfiguration()(AppleVpnAlwaysOnConfigurationable)
    GetChildSecurityAssociationParameters()(IosVpnSecurityAssociationParametersable)
    GetClientAuthenticationType()(*VpnClientAuthenticationType)
    GetDeadPeerDetectionRate()(*VpnDeadPeerDetectionRate)
    GetDisableMobilityAndMultihoming()(*bool)
    GetDisableRedirect()(*bool)
    GetEnableAlwaysOnConfiguration()(*bool)
    GetEnableCertificateRevocationCheck()(*bool)
    GetEnableEAP()(*bool)
    GetEnablePerfectForwardSecrecy()(*bool)
    GetEnableUseInternalSubnetAttributes()(*bool)
    GetLocalIdentifier()(*VpnLocalIdentifier)
    GetMtuSizeInBytes()(*int32)
    GetRemoteIdentifier()(*string)
    GetSecurityAssociationParameters()(IosVpnSecurityAssociationParametersable)
    GetServerCertificateCommonName()(*string)
    GetServerCertificateIssuerCommonName()(*string)
    GetServerCertificateType()(*VpnServerCertificateType)
    GetSharedSecret()(*string)
    GetTlsMaximumVersion()(*string)
    GetTlsMinimumVersion()(*string)
    SetAllowDefaultChildSecurityAssociationParameters(value *bool)()
    SetAllowDefaultSecurityAssociationParameters(value *bool)()
    SetAlwaysOnConfiguration(value AppleVpnAlwaysOnConfigurationable)()
    SetChildSecurityAssociationParameters(value IosVpnSecurityAssociationParametersable)()
    SetClientAuthenticationType(value *VpnClientAuthenticationType)()
    SetDeadPeerDetectionRate(value *VpnDeadPeerDetectionRate)()
    SetDisableMobilityAndMultihoming(value *bool)()
    SetDisableRedirect(value *bool)()
    SetEnableAlwaysOnConfiguration(value *bool)()
    SetEnableCertificateRevocationCheck(value *bool)()
    SetEnableEAP(value *bool)()
    SetEnablePerfectForwardSecrecy(value *bool)()
    SetEnableUseInternalSubnetAttributes(value *bool)()
    SetLocalIdentifier(value *VpnLocalIdentifier)()
    SetMtuSizeInBytes(value *int32)()
    SetRemoteIdentifier(value *string)()
    SetSecurityAssociationParameters(value IosVpnSecurityAssociationParametersable)()
    SetServerCertificateCommonName(value *string)()
    SetServerCertificateIssuerCommonName(value *string)()
    SetServerCertificateType(value *VpnServerCertificateType)()
    SetSharedSecret(value *string)()
    SetTlsMaximumVersion(value *string)()
    SetTlsMinimumVersion(value *string)()
}
