package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemCloudPCsItemTroubleshootRequestBuilder provides operations to call the troubleshoot method.
type UsersItemCloudPCsItemTroubleshootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemCloudPCsItemTroubleshootRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemCloudPCsItemTroubleshootRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemCloudPCsItemTroubleshootRequestBuilderInternal instantiates a new TroubleshootRequestBuilder and sets the default values.
func NewUsersItemCloudPCsItemTroubleshootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemCloudPCsItemTroubleshootRequestBuilder) {
    m := &UsersItemCloudPCsItemTroubleshootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/microsoft.graph.troubleshoot";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemCloudPCsItemTroubleshootRequestBuilder instantiates a new TroubleshootRequestBuilder and sets the default values.
func NewUsersItemCloudPCsItemTroubleshootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemCloudPCsItemTroubleshootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemCloudPCsItemTroubleshootRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation troubleshoot a specific Cloud PC. Use this API to check the health status of the Cloud PC and the session host.
func (m *UsersItemCloudPCsItemTroubleshootRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *UsersItemCloudPCsItemTroubleshootRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post troubleshoot a specific Cloud PC. Use this API to check the health status of the Cloud PC and the session host.
func (m *UsersItemCloudPCsItemTroubleshootRequestBuilder) Post(ctx context.Context, requestConfiguration *UsersItemCloudPCsItemTroubleshootRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
