package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i3783ecca2e180950ea8455123aaec3dca5384fc8410f062bca07795be89adfa0 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/presences/item/clearpresence"
    i5ed3d0ca2ce865846a0f2ec773a6d70c98a29e8021c06bb49fee978222054e11 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/presences/item/setuserpreferredpresence"
    i63204af1eb97ddfc7f34adf891a9a23250dbd8ffb28475e6ee571d26aadfd107 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/presences/item/setpresence"
    ib61b07a61e3d9e1ae2e83dcf306c1142b1e2f1e15e289603a751e740648320b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/communications/presences/item/clearuserpreferredpresence"
)

// PresenceItemRequestBuilder provides operations to manage the presences property of the microsoft.graph.cloudCommunications entity.
type PresenceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PresenceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresenceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PresenceItemRequestBuilderGetQueryParameters get presences from communications
type PresenceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PresenceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresenceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PresenceItemRequestBuilderGetQueryParameters
}
// PresenceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresenceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClearPresence the clearPresence property
func (m *PresenceItemRequestBuilder) ClearPresence()(*i3783ecca2e180950ea8455123aaec3dca5384fc8410f062bca07795be89adfa0.ClearPresenceRequestBuilder) {
    return i3783ecca2e180950ea8455123aaec3dca5384fc8410f062bca07795be89adfa0.NewClearPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClearUserPreferredPresence the clearUserPreferredPresence property
func (m *PresenceItemRequestBuilder) ClearUserPreferredPresence()(*ib61b07a61e3d9e1ae2e83dcf306c1142b1e2f1e15e289603a751e740648320b1.ClearUserPreferredPresenceRequestBuilder) {
    return ib61b07a61e3d9e1ae2e83dcf306c1142b1e2f1e15e289603a751e740648320b1.NewClearUserPreferredPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPresenceItemRequestBuilderInternal instantiates a new PresenceItemRequestBuilder and sets the default values.
func NewPresenceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresenceItemRequestBuilder) {
    m := &PresenceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/presences/{presence%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPresenceItemRequestBuilder instantiates a new PresenceItemRequestBuilder and sets the default values.
func NewPresenceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresenceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPresenceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property presences for communications
func (m *PresenceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PresenceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get presences from communications
func (m *PresenceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PresenceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property presences in communications
func (m *PresenceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Presenceable, requestConfiguration *PresenceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property presences for communications
func (m *PresenceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PresenceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get presences from communications
func (m *PresenceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PresenceItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Presenceable, error) {
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
// Patch update the navigation property presences in communications
func (m *PresenceItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Presenceable, requestConfiguration *PresenceItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Presenceable, error) {
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
// SetPresence the setPresence property
func (m *PresenceItemRequestBuilder) SetPresence()(*i63204af1eb97ddfc7f34adf891a9a23250dbd8ffb28475e6ee571d26aadfd107.SetPresenceRequestBuilder) {
    return i63204af1eb97ddfc7f34adf891a9a23250dbd8ffb28475e6ee571d26aadfd107.NewSetPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUserPreferredPresence the setUserPreferredPresence property
func (m *PresenceItemRequestBuilder) SetUserPreferredPresence()(*i5ed3d0ca2ce865846a0f2ec773a6d70c98a29e8021c06bb49fee978222054e11.SetUserPreferredPresenceRequestBuilder) {
    return i5ed3d0ca2ce865846a0f2ec773a6d70c98a29e8021c06bb49fee978222054e11.NewSetUserPreferredPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
