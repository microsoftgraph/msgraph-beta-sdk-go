package iosvppapp

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i24b5eee907f3023ad6b32b470f7d6d36d9dbb47745fa42f740f7ec20c352f5db "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/iosvppapp/revokedevicelicense"
    i6e18af1428f7b694ac5766c2ccdc52630011ab10de0b02134747274ffb915696 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/iosvppapp/revokealllicenses"
    ieba7d49970f8ba88a7065c6eb5ffd1bee021c90edc2d1e7786356454063de65b "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/mobileapps/item/iosvppapp/revokeuserlicense"
)

// IosVppAppRequestBuilder builds and executes requests for operations under \deviceAppManagement\mobileApps\{mobileApp-id}\microsoft.graph.iosVppApp
type IosVppAppRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewIosVppAppRequestBuilderInternal instantiates a new IosVppAppRequestBuilder and sets the default values.
func NewIosVppAppRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosVppAppRequestBuilder) {
    m := &IosVppAppRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/mobileApps/{mobileApp%2Did}/microsoft.graph.iosVppApp";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIosVppAppRequestBuilder instantiates a new IosVppAppRequestBuilder and sets the default values.
func NewIosVppAppRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosVppAppRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosVppAppRequestBuilderInternal(urlParams, requestAdapter)
}
// RevokeAllLicenses the revokeAllLicenses property
func (m *IosVppAppRequestBuilder) RevokeAllLicenses()(*i6e18af1428f7b694ac5766c2ccdc52630011ab10de0b02134747274ffb915696.RevokeAllLicensesRequestBuilder) {
    return i6e18af1428f7b694ac5766c2ccdc52630011ab10de0b02134747274ffb915696.NewRevokeAllLicensesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeDeviceLicense the revokeDeviceLicense property
func (m *IosVppAppRequestBuilder) RevokeDeviceLicense()(*i24b5eee907f3023ad6b32b470f7d6d36d9dbb47745fa42f740f7ec20c352f5db.RevokeDeviceLicenseRequestBuilder) {
    return i24b5eee907f3023ad6b32b470f7d6d36d9dbb47745fa42f740f7ec20c352f5db.NewRevokeDeviceLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RevokeUserLicense the revokeUserLicense property
func (m *IosVppAppRequestBuilder) RevokeUserLicense()(*ieba7d49970f8ba88a7065c6eb5ffd1bee021c90edc2d1e7786356454063de65b.RevokeUserLicenseRequestBuilder) {
    return ieba7d49970f8ba88a7065c6eb5ffd1bee021c90edc2d1e7786356454063de65b.NewRevokeUserLicenseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
