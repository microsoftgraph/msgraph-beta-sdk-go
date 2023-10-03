package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CertificateConnectorDetailsItemGetHealthMetricsResponse 
// Deprecated: This class is obsolete. Use getHealthMetricsPostResponse instead.
type CertificateConnectorDetailsItemGetHealthMetricsResponse struct {
    CertificateConnectorDetailsItemGetHealthMetricsPostResponse
}
// NewCertificateConnectorDetailsItemGetHealthMetricsResponse instantiates a new CertificateConnectorDetailsItemGetHealthMetricsResponse and sets the default values.
func NewCertificateConnectorDetailsItemGetHealthMetricsResponse()(*CertificateConnectorDetailsItemGetHealthMetricsResponse) {
    m := &CertificateConnectorDetailsItemGetHealthMetricsResponse{
        CertificateConnectorDetailsItemGetHealthMetricsPostResponse: *NewCertificateConnectorDetailsItemGetHealthMetricsPostResponse(),
    }
    return m
}
// CreateCertificateConnectorDetailsItemGetHealthMetricsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCertificateConnectorDetailsItemGetHealthMetricsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCertificateConnectorDetailsItemGetHealthMetricsResponse(), nil
}
// CertificateConnectorDetailsItemGetHealthMetricsResponseable 
// Deprecated: This class is obsolete. Use getHealthMetricsPostResponse instead.
type CertificateConnectorDetailsItemGetHealthMetricsResponseable interface {
    CertificateConnectorDetailsItemGetHealthMetricsPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
