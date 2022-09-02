package importappledeviceidentitylist

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ImportAppleDeviceIdentityListRequestBuilder provides operations to call the importAppleDeviceIdentityList method.
type ImportAppleDeviceIdentityListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ImportAppleDeviceIdentityListRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ImportAppleDeviceIdentityListRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewImportAppleDeviceIdentityListRequestBuilderInternal instantiates a new ImportAppleDeviceIdentityListRequestBuilder and sets the default values.
func NewImportAppleDeviceIdentityListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportAppleDeviceIdentityListRequestBuilder) {
    m := &ImportAppleDeviceIdentityListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/depOnboardingSettings/{depOnboardingSetting%2Did}/importedAppleDeviceIdentities/microsoft.graph.importAppleDeviceIdentityList";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewImportAppleDeviceIdentityListRequestBuilder instantiates a new ImportAppleDeviceIdentityListRequestBuilder and sets the default values.
func NewImportAppleDeviceIdentityListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ImportAppleDeviceIdentityListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewImportAppleDeviceIdentityListRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action importAppleDeviceIdentityList
func (m *ImportAppleDeviceIdentityListRequestBuilder) CreatePostRequestInformation(body ImportAppleDeviceIdentityListPostRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action importAppleDeviceIdentityList
func (m *ImportAppleDeviceIdentityListRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ImportAppleDeviceIdentityListPostRequestBodyable, requestConfiguration *ImportAppleDeviceIdentityListRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action importAppleDeviceIdentityList
func (m *ImportAppleDeviceIdentityListRequestBuilder) Post(ctx context.Context, body ImportAppleDeviceIdentityListPostRequestBodyable, requestConfiguration *ImportAppleDeviceIdentityListRequestBuilderPostRequestConfiguration)(ImportAppleDeviceIdentityListResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateImportAppleDeviceIdentityListResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ImportAppleDeviceIdentityListResponseable), nil
}
