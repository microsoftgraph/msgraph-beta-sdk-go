package iosvppapp

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i18fbae264dc0cd6f504171ab8a47ed1533e91ea2b0158e64f3d7a829ffc83d1a "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/app/iosvppapp/revokealllicenses"
    ib544cbf6ce75d3ffb003ffffe0a91e55b85bac7b4ef4f6ead1735b9296125b2a "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/app/iosvppapp/revokeuserlicense"
    id461c5cf6f3351b5c2dec99429bbb1a4130e9ee3e7e9bc94ce9b788e647e6aba "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/userstatuses/item/app/iosvppapp/revokedevicelicense"
)

// IosVppAppRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileApps\{mobileApp-id}\userStatuses\{userAppInstallStatus-id}\app\microsoft.graph.iosVppApp
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
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp_id}/userStatuses/{userAppInstallStatus_id}/app/microsoft.graph.iosVppApp";
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
func (m *IosVppAppRequestBuilder) RevokeAllLicenses()(*i18fbae264dc0cd6f504171ab8a47ed1533e91ea2b0158e64f3d7a829ffc83d1a.RevokeAllLicensesRequestBuilder) {
    return i18fbae264dc0cd6f504171ab8a47ed1533e91ea2b0158e64f3d7a829ffc83d1a.NewRevokeAllLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosVppAppRequestBuilder) RevokeDeviceLicense()(*id461c5cf6f3351b5c2dec99429bbb1a4130e9ee3e7e9bc94ce9b788e647e6aba.RevokeDeviceLicenseRequestBuilder) {
    return id461c5cf6f3351b5c2dec99429bbb1a4130e9ee3e7e9bc94ce9b788e647e6aba.NewRevokeDeviceLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosVppAppRequestBuilder) RevokeUserLicense()(*ib544cbf6ce75d3ffb003ffffe0a91e55b85bac7b4ef4f6ead1735b9296125b2a.RevokeUserLicenseRequestBuilder) {
    return ib544cbf6ce75d3ffb003ffffe0a91e55b85bac7b4ef4f6ead1735b9296125b2a.NewRevokeUserLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
