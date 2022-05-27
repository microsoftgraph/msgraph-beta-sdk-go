package getstorageaccountswithsubscriptionid

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetStorageAccountsWithSubscriptionIdRequestBuilder provides operations to call the getStorageAccounts method.
type GetStorageAccountsWithSubscriptionIdRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetStorageAccountsWithSubscriptionIdRequestBuilderInternal instantiates a new GetStorageAccountsWithSubscriptionIdRequestBuilder and sets the default values.
func NewGetStorageAccountsWithSubscriptionIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, subscriptionId *string)(*GetStorageAccountsWithSubscriptionIdRequestBuilder) {
    m := &GetStorageAccountsWithSubscriptionIdRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/snapshots/microsoft.graph.getStorageAccounts(subscriptionId='{subscriptionId}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if subscriptionId != nil {
        urlTplParams["subscriptionId"] = *subscriptionId
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetStorageAccountsWithSubscriptionIdRequestBuilder instantiates a new GetStorageAccountsWithSubscriptionIdRequestBuilder and sets the default values.
func NewGetStorageAccountsWithSubscriptionIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetStorageAccountsWithSubscriptionIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetStorageAccountsWithSubscriptionIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getStorageAccounts
func (m *GetStorageAccountsWithSubscriptionIdRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function getStorageAccounts
func (m *GetStorageAccountsWithSubscriptionIdRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *GetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get invoke function getStorageAccounts
func (m *GetStorageAccountsWithSubscriptionIdRequestBuilder) Get()(GetStorageAccountsWithSubscriptionIdResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function getStorageAccounts
func (m *GetStorageAccountsWithSubscriptionIdRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *GetStorageAccountsWithSubscriptionIdRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetStorageAccountsWithSubscriptionIdResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetStorageAccountsWithSubscriptionIdResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetStorageAccountsWithSubscriptionIdResponseable), nil
}
