package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/externalconnectors"
    i0076de30666e5a2bbac488b5589060c505c936d20bd513b4af52ba9d47ae03ce "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/groups/item/members"
    i6bcc0207423f237bfcb4579df45dd5e5b9ff58acfaec84876442aecc9aa12072 "github.com/microsoftgraph/msgraph-beta-sdk-go/connections/item/groups/item/members/item"
)

// ExternalGroupItemRequestBuilder provides operations to manage the groups property of the microsoft.graph.externalConnectors.externalConnection entity.
type ExternalGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ExternalGroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExternalGroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ExternalGroupItemRequestBuilderGetQueryParameters get groups from connections
type ExternalGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExternalGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExternalGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExternalGroupItemRequestBuilderGetQueryParameters
}
// ExternalGroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExternalGroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewExternalGroupItemRequestBuilderInternal instantiates a new ExternalGroupItemRequestBuilder and sets the default values.
func NewExternalGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExternalGroupItemRequestBuilder) {
    m := &ExternalGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/connections/{externalConnection%2Did}/groups/{externalGroup%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExternalGroupItemRequestBuilder instantiates a new ExternalGroupItemRequestBuilder and sets the default values.
func NewExternalGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExternalGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExternalGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property groups for connections
func (m *ExternalGroupItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ExternalGroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get groups from connections
func (m *ExternalGroupItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ExternalGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property groups in connections
func (m *ExternalGroupItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalGroupable, requestConfiguration *ExternalGroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property groups for connections
func (m *ExternalGroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ExternalGroupItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get get groups from connections
func (m *ExternalGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExternalGroupItemRequestBuilderGetRequestConfiguration)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalGroupable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.CreateExternalGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalGroupable), nil
}
// Members provides operations to manage the members property of the microsoft.graph.externalConnectors.externalGroup entity.
func (m *ExternalGroupItemRequestBuilder) Members()(*i0076de30666e5a2bbac488b5589060c505c936d20bd513b4af52ba9d47ae03ce.MembersRequestBuilder) {
    return i0076de30666e5a2bbac488b5589060c505c936d20bd513b4af52ba9d47ae03ce.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.externalConnectors.externalGroup entity.
func (m *ExternalGroupItemRequestBuilder) MembersById(id string)(*i6bcc0207423f237bfcb4579df45dd5e5b9ff58acfaec84876442aecc9aa12072.IdentityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["identity%2Did"] = id
    }
    return i6bcc0207423f237bfcb4579df45dd5e5b9ff58acfaec84876442aecc9aa12072.NewIdentityItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property groups in connections
func (m *ExternalGroupItemRequestBuilder) Patch(ctx context.Context, body ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalGroupable, requestConfiguration *ExternalGroupItemRequestBuilderPatchRequestConfiguration)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalGroupable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.CreateExternalGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.ExternalGroupable), nil
}
