package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilder provides operations to call the updateScopeTags method.
type GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilderInternal instantiates a new UpdateScopeTagsRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilder) {
    m := &GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport%2Did}/microsoft.graph.updateScopeTags";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewGroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilder instantiates a new UpdateScopeTagsRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateScopeTags
func (m *GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilder) Post(ctx context.Context, body GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsPostRequestBodyable, requestConfiguration *GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilderPostRequestConfiguration)(GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateGroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsResponseable), nil
}
// ToPostRequestInformation invoke action updateScopeTags
func (m *GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsPostRequestBodyable, requestConfiguration *GroupPolicyMigrationReportsItemMicrosoftGraphUpdateScopeTagsUpdateScopeTagsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
