package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilder provides operations to call the addGroup method.
type ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilderInternal instantiates a new MicrosoftGraphAddGroupRequestBuilder and sets the default values.
func NewItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilder) {
    m := &ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/groupLifecyclePolicies/{groupLifecyclePolicy%2Did}/microsoft.graph.addGroup";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilder instantiates a new MicrosoftGraphAddGroupRequestBuilder and sets the default values.
func NewItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action addGroup
func (m *ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilder) Post(ctx context.Context, body ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupAddGroupPostRequestBodyable, requestConfiguration *ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilderPostRequestConfiguration)(ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupAddGroupResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupAddGroupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupAddGroupResponseable), nil
}
// ToPostRequestInformation invoke action addGroup
func (m *ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupAddGroupPostRequestBodyable, requestConfiguration *ItemGroupLifecyclePoliciesItemMicrosoftGraphAddGroupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
