package iosvppapp

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i01615e369fddd26fbf9f2a9762b1d4ac304ddda1d4948d54db1d56bf795e2221 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses/item/app/iosvppapp/revokedevicelicense"
    i246c46f95b555db6d8964aa1e90be55f7be5eac10e1764d3c3f5441f722d171f "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses/item/app/iosvppapp/revokealllicenses"
    iea6b669b897a4c035a59b5f017a44d3620d7cabef826776dd92c368838079fae "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/devicestatuses/item/app/iosvppapp/revokeuserlicense"
)

// IosVppAppRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileApps\{mobileApp-id}\deviceStatuses\{mobileAppInstallStatus-id}\app\microsoft.graph.iosVppApp
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
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp_id}/deviceStatuses/{mobileAppInstallStatus_id}/app/microsoft.graph.iosVppApp";
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
func (m *IosVppAppRequestBuilder) RevokeAllLicenses()(*i246c46f95b555db6d8964aa1e90be55f7be5eac10e1764d3c3f5441f722d171f.RevokeAllLicensesRequestBuilder) {
    return i246c46f95b555db6d8964aa1e90be55f7be5eac10e1764d3c3f5441f722d171f.NewRevokeAllLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosVppAppRequestBuilder) RevokeDeviceLicense()(*i01615e369fddd26fbf9f2a9762b1d4ac304ddda1d4948d54db1d56bf795e2221.RevokeDeviceLicenseRequestBuilder) {
    return i01615e369fddd26fbf9f2a9762b1d4ac304ddda1d4948d54db1d56bf795e2221.NewRevokeDeviceLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *IosVppAppRequestBuilder) RevokeUserLicense()(*iea6b669b897a4c035a59b5f017a44d3620d7cabef826776dd92c368838079fae.RevokeUserLicenseRequestBuilder) {
    return iea6b669b897a4c035a59b5f017a44d3620d7cabef826776dd92c368838079fae.NewRevokeUserLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
