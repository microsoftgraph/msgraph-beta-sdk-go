package iosvppapp

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4b68c14bf56a3f6c7820c357040415d8849dbf591fc19bcb6d8186cee764e07d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/devicestatuses/item/app/iosvppapp/revokedevicelicense"
    i6005d3fb2d54c99d04a2a9dde84a70404a8c8f6274eacc2eb425348d2e10e214 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/devicestatuses/item/app/iosvppapp/revokealllicenses"
    i6307b530460ab7a670e27732aa36f0cb56c7e09f42777c5ed6dbc19877cd4f73 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/devicestatuses/item/app/iosvppapp/revokeuserlicense"
)

// IosVppAppRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileApps\{mobileApp-id}\userStatuses\{userAppInstallStatus-id}\deviceStatuses\{mobileAppInstallStatus-id}\app\microsoft.graph.iosVppApp
type IosVppAppRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// NewIosVppAppRequestBuilderInternal instantiates a new IosVppAppRequestBuilder and sets the default values.
func NewIosVppAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosVppAppRequestBuilder) {
    m := &IosVppAppRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp_id}/userStatuses/{userAppInstallStatus_id}/deviceStatuses/{mobileAppInstallStatus_id}/app/microsoft.graph.iosVppApp";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIosVppAppRequestBuilder instantiates a new IosVppAppRequestBuilder and sets the default values.
func NewIosVppAppRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*IosVppAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosVppAppRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *IosVppAppRequestBuilder) RevokeAllLicenses()(*i6005d3fb2d54c99d04a2a9dde84a70404a8c8f6274eacc2eb425348d2e10e214.RevokeAllLicensesRequestBuilder) {
    return i6005d3fb2d54c99d04a2a9dde84a70404a8c8f6274eacc2eb425348d2e10e214.NewRevokeAllLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosVppAppRequestBuilder) RevokeDeviceLicense()(*i4b68c14bf56a3f6c7820c357040415d8849dbf591fc19bcb6d8186cee764e07d.RevokeDeviceLicenseRequestBuilder) {
    return i4b68c14bf56a3f6c7820c357040415d8849dbf591fc19bcb6d8186cee764e07d.NewRevokeDeviceLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosVppAppRequestBuilder) RevokeUserLicense()(*i6307b530460ab7a670e27732aa36f0cb56c7e09f42777c5ed6dbc19877cd4f73.RevokeUserLicenseRequestBuilder) {
    return i6307b530460ab7a670e27732aa36f0cb56c7e09f42777c5ed6dbc19877cd4f73.NewRevokeUserLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
