package additionalaccesswithaccesspackageid

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AdditionalAccessWithAccessPackageIdRequestBuilder provides operations to call the additionalAccess method.
type AdditionalAccessWithAccessPackageIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AdditionalAccessWithAccessPackageIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AdditionalAccessWithAccessPackageIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAdditionalAccessWithAccessPackageIdRequestBuilderInternal instantiates a new AdditionalAccessWithAccessPackageIdRequestBuilder and sets the default values.
func NewAdditionalAccessWithAccessPackageIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, accessPackageId *string)(*AdditionalAccessWithAccessPackageIdRequestBuilder) {
    m := &AdditionalAccessWithAccessPackageIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackageAssignments/microsoft.graph.additionalAccess(accessPackageId='{accessPackageId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageId != nil {
        urlTplParams[""] = *accessPackageId
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAdditionalAccessWithAccessPackageIdRequestBuilder instantiates a new AdditionalAccessWithAccessPackageIdRequestBuilder and sets the default values.
func NewAdditionalAccessWithAccessPackageIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdditionalAccessWithAccessPackageIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdditionalAccessWithAccessPackageIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function additionalAccess
func (m *AdditionalAccessWithAccessPackageIdRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function additionalAccess
func (m *AdditionalAccessWithAccessPackageIdRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *AdditionalAccessWithAccessPackageIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function additionalAccess
func (m *AdditionalAccessWithAccessPackageIdRequestBuilder) Get()(AdditionalAccessWithAccessPackageIdResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function additionalAccess
func (m *AdditionalAccessWithAccessPackageIdRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *AdditionalAccessWithAccessPackageIdRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(AdditionalAccessWithAccessPackageIdResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateAdditionalAccessWithAccessPackageIdResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(AdditionalAccessWithAccessPackageIdResponseable), nil
}
