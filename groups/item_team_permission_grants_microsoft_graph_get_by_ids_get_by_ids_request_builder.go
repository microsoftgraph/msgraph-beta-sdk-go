package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilder provides operations to call the getByIds method.
type ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilderInternal instantiates a new GetByIdsRequestBuilder and sets the default values.
func NewItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilder) {
    m := &ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/permissionGrants/microsoft.graph.getByIds";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilder instantiates a new GetByIdsRequestBuilder and sets the default values.
func NewItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post return the directory objects specified in a list of IDs. Some common uses for this function are to:
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/directoryobject-getbyids?view=graph-rest-1.0
func (m *ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilder) Post(ctx context.Context, body ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsPostRequestBodyable, requestConfiguration *ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilderPostRequestConfiguration)(ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsResponseable), nil
}
// ToPostRequestInformation return the directory objects specified in a list of IDs. Some common uses for this function are to:
func (m *ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsPostRequestBodyable, requestConfiguration *ItemTeamPermissionGrantsMicrosoftGraphGetByIdsGetByIdsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
