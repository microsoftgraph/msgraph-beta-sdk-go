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
// AdditionalAccessWithAccessPackageIdRequestBuilderGetOptions options for Get
type AdditionalAccessWithAccessPackageIdRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
func (m *AdditionalAccessWithAccessPackageIdRequestBuilder) CreateGetRequestInformation(options *AdditionalAccessWithAccessPackageIdRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get invoke function additionalAccess
func (m *AdditionalAccessWithAccessPackageIdRequestBuilder) Get(options *AdditionalAccessWithAccessPackageIdRequestBuilderGetOptions)(AdditionalAccessWithAccessPackageIdResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateAdditionalAccessWithAccessPackageIdResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(AdditionalAccessWithAccessPackageIdResponseable), nil
}
