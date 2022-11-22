package getsupportedcloudpcremoteactions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GetSupportedCloudPcRemoteActionsRequestBuilder provides operations to call the getSupportedCloudPcRemoteActions method.
type GetSupportedCloudPcRemoteActionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetSupportedCloudPcRemoteActionsRequestBuilderGetQueryParameters invoke function getSupportedCloudPcRemoteActions
type GetSupportedCloudPcRemoteActionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// GetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GetSupportedCloudPcRemoteActionsRequestBuilderGetQueryParameters
}
// NewGetSupportedCloudPcRemoteActionsRequestBuilderInternal instantiates a new GetSupportedCloudPcRemoteActionsRequestBuilder and sets the default values.
func NewGetSupportedCloudPcRemoteActionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetSupportedCloudPcRemoteActionsRequestBuilder) {
    m := &GetSupportedCloudPcRemoteActionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}/microsoft.graph.getSupportedCloudPcRemoteActions(){?%24top,%24skip,%24search,%24filter,%24count}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetSupportedCloudPcRemoteActionsRequestBuilder instantiates a new GetSupportedCloudPcRemoteActionsRequestBuilder and sets the default values.
func NewGetSupportedCloudPcRemoteActionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetSupportedCloudPcRemoteActionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetSupportedCloudPcRemoteActionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function getSupportedCloudPcRemoteActions
func (m *GetSupportedCloudPcRemoteActionsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function getSupportedCloudPcRemoteActions
func (m *GetSupportedCloudPcRemoteActionsRequestBuilder) Get(ctx context.Context, requestConfiguration *GetSupportedCloudPcRemoteActionsRequestBuilderGetRequestConfiguration)(GetSupportedCloudPcRemoteActionsResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, CreateGetSupportedCloudPcRemoteActionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GetSupportedCloudPcRemoteActionsResponseable), nil
}
