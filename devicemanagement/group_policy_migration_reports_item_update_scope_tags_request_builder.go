package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder provides operations to call the updateScopeTags method.
type GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderInternal instantiates a new UpdateScopeTagsRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder) {
    m := &GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport%2Did}/updateScopeTags";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewGroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder instantiates a new UpdateScopeTagsRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateScopeTags
func (m *GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder) Post(ctx context.Context, body GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBodyable, requestConfiguration *GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderPostRequestConfiguration)(GroupPolicyMigrationReportsItemUpdateScopeTagsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateGroupPolicyMigrationReportsItemUpdateScopeTagsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GroupPolicyMigrationReportsItemUpdateScopeTagsResponseable), nil
}
// ToPostRequestInformation invoke action updateScopeTags
func (m *GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body GroupPolicyMigrationReportsItemUpdateScopeTagsPostRequestBodyable, requestConfiguration *GroupPolicyMigrationReportsItemUpdateScopeTagsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
