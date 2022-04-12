package presence

import (
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
// PresenceRequestBuilderDeleteOptions options for Delete
type PresenceRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// PresenceRequestBuilderGetOptions options for Get
type PresenceRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PresenceRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// PresenceRequestBuilderGetQueryParameters get presence from users
type PresenceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PresenceRequestBuilderPatchOptions options for Patch
type PresenceRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Presenceable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ClearPresence the clearPresence property
func (m *PresenceRequestBuilder) ClearPresence()(*i76a36a5967b7a0dbc74be00d4a84f2e878f560504d3a9952fb37d1b227043ead.ClearPresenceRequestBuilder) {
    return i76a36a5967b7a0dbc74be00d4a84f2e878f560504d3a9952fb37d1b227043ead.NewClearPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClearUserPreferredPresence the clearUserPreferredPresence property
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
func (m *PresenceRequestBuilder) CreateDeleteRequestInformation(options *PresenceRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get presence from users
func (m *PresenceRequestBuilder) CreateGetRequestInformation(options *PresenceRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property presence in users
func (m *PresenceRequestBuilder) CreatePatchRequestInformation(options *PresenceRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property presence for users
func (m *PresenceRequestBuilder) Delete(options *PresenceRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get presence from users
func (m *PresenceRequestBuilder) Get(options *PresenceRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Presenceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePresenceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Presenceable), nil
}
// Patch update the navigation property presence in users
func (m *PresenceRequestBuilder) Patch(options *PresenceRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SetPresence the setPresence property
func (m *PresenceRequestBuilder) SetPresence()(*iba58ce57dde26291ce4951142ac6e48a720e189264fcf02a0fa187582dc0208d.SetPresenceRequestBuilder) {
    return iba58ce57dde26291ce4951142ac6e48a720e189264fcf02a0fa187582dc0208d.NewSetPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUserPreferredPresence the setUserPreferredPresence property
func (m *PresenceRequestBuilder) SetUserPreferredPresence()(*i5dd07c47805059d448d8c8f1e7951182fab55575575904704c1de455e4f036e9.SetUserPreferredPresenceRequestBuilder) {
    return i5dd07c47805059d448d8c8f1e7951182fab55575575904704c1de455e4f036e9.NewSetUserPreferredPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
