package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder provides operations to manage the cloudPCConnectivityIssues property of the microsoft.graph.deviceManagement entity.
type CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderGetQueryParameters the list of CloudPC Connectivity Issue.
type CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderGetQueryParameters
}
// CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderInternal instantiates a new CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder and sets the default values.
func NewCloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder) {
    m := &CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/cloudPCConnectivityIssues/{cloudPCConnectivityIssue%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder instantiates a new CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder and sets the default values.
func NewCloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property cloudPCConnectivityIssues for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the list of CloudPC Connectivity Issue.
// returns a CloudPCConnectivityIssueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCConnectivityIssueable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPCConnectivityIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCConnectivityIssueable), nil
}
// Patch update the navigation property cloudPCConnectivityIssues in deviceManagement
// returns a CloudPCConnectivityIssueable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCConnectivityIssueable, requestConfiguration *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCConnectivityIssueable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPCConnectivityIssueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCConnectivityIssueable), nil
}
// ToDeleteRequestInformation delete navigation property cloudPCConnectivityIssues for deviceManagement
// returns a *RequestInformation when successful
func (m *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of CloudPC Connectivity Issue.
// returns a *RequestInformation when successful
func (m *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property cloudPCConnectivityIssues in deviceManagement
// returns a *RequestInformation when successful
func (m *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCConnectivityIssueable, requestConfiguration *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder when successful
func (m *CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder) WithUrl(rawUrl string)(*CloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder) {
    return NewCloudpcconnectivityissuesCloudPCConnectivityIssueItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
