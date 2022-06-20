package additionalaccesswithaccesspackageidwithincompatibleaccesspackageid

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder provides operations to call the additionalAccess method.
type AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal instantiates a new AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder and sets the default values.
func NewAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, accessPackageId *string, incompatibleAccessPackageId *string)(*AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    m := &AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/microsoft.graph.additionalAccess(accessPackageId='{accessPackageId}',incompatibleAccessPackageId='{incompatibleAccessPackageId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageId != nil {
        urlTplParams["accessPackageId"] = *accessPackageId
    }
    if incompatibleAccessPackageId != nil {
        urlTplParams["incompatibleAccessPackageId"] = *incompatibleAccessPackageId
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder instantiates a new AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder and sets the default values.
func NewAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// CreateGetRequestInformation invoke function additionalAccess
func (m *AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function additionalAccess
func (m *AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function additionalAccess
func (m *AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) Get()(AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function additionalAccess
func (m *AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateAdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(AdditionalAccessWithAccessPackageIdWithIncompatibleAccessPackageIdResponseable), nil
}
