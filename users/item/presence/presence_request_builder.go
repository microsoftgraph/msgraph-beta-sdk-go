package presence

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1d8eba405fb9a387de3bc92cd3706580c107d66d82467147261e3b29e92a01bb "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/presence/clearuserpreferredpresence"
    i5dd07c47805059d448d8c8f1e7951182fab55575575904704c1de455e4f036e9 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/presence/setuserpreferredpresence"
    i76a36a5967b7a0dbc74be00d4a84f2e878f560504d3a9952fb37d1b227043ead "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/presence/clearpresence"
    iba58ce57dde26291ce4951142ac6e48a720e189264fcf02a0fa187582dc0208d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/presence/setpresence"
)

// PresenceRequestBuilder provides operations to manage the presence property of the microsoft.graph.user entity.
type PresenceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PresenceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresenceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PresenceRequestBuilderGetQueryParameters get a user's presence information.
type PresenceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PresenceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresenceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PresenceRequestBuilderGetQueryParameters
}
// PresenceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresenceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClearPresence provides operations to call the clearPresence method.
func (m *PresenceRequestBuilder) ClearPresence()(*i76a36a5967b7a0dbc74be00d4a84f2e878f560504d3a9952fb37d1b227043ead.ClearPresenceRequestBuilder) {
    return i76a36a5967b7a0dbc74be00d4a84f2e878f560504d3a9952fb37d1b227043ead.NewClearPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClearUserPreferredPresence provides operations to call the clearUserPreferredPresence method.
func (m *PresenceRequestBuilder) ClearUserPreferredPresence()(*i1d8eba405fb9a387de3bc92cd3706580c107d66d82467147261e3b29e92a01bb.ClearUserPreferredPresenceRequestBuilder) {
    return i1d8eba405fb9a387de3bc92cd3706580c107d66d82467147261e3b29e92a01bb.NewClearUserPreferredPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPresenceRequestBuilderInternal instantiates a new PresenceRequestBuilder and sets the default values.
func NewPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresenceRequestBuilder) {
    m := &PresenceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/presence{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPresenceRequestBuilder instantiates a new PresenceRequestBuilder and sets the default values.
func NewPresenceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property presence for users
func (m *PresenceRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PresenceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get a user's presence information.
func (m *PresenceRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PresenceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property presence in users
func (m *PresenceRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Presenceable, requestConfiguration *PresenceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property presence for users
func (m *PresenceRequestBuilder) Delete(ctx context.Context, requestConfiguration *PresenceRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get a user's presence information.
func (m *PresenceRequestBuilder) Get(ctx context.Context, requestConfiguration *PresenceRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Presenceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePresenceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Presenceable), nil
}
// Patch update the navigation property presence in users
func (m *PresenceRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Presenceable, requestConfiguration *PresenceRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Presenceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePresenceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Presenceable), nil
}
// SetPresence provides operations to call the setPresence method.
func (m *PresenceRequestBuilder) SetPresence()(*iba58ce57dde26291ce4951142ac6e48a720e189264fcf02a0fa187582dc0208d.SetPresenceRequestBuilder) {
    return iba58ce57dde26291ce4951142ac6e48a720e189264fcf02a0fa187582dc0208d.NewSetPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUserPreferredPresence provides operations to call the setUserPreferredPresence method.
func (m *PresenceRequestBuilder) SetUserPreferredPresence()(*i5dd07c47805059d448d8c8f1e7951182fab55575575904704c1de455e4f036e9.SetUserPreferredPresenceRequestBuilder) {
    return i5dd07c47805059d448d8c8f1e7951182fab55575575904704c1de455e4f036e9.NewSetUserPreferredPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
