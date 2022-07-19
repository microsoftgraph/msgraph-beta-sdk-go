package privilegedsignupstatus

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i505eff958ea7946e725489038a9fbe85c297626b944e12d9c6f68697f1a65258 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/issignedup"
    i576c2492be4604b88fea75e47108c8bd8db61a3de84b266733d2d7849bec46eb "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/cansignup"
    ib823f830e3ea0500a214b973cf170985ed37b35d907e2b81e06df8aaa8f66256 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/completesetup"
    ibe660ff6aa6da80e38b282109151294b9409fe46a692bb37484dd8a1705b5a94 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/signup"
    id44dae2f191b15a24754b7171e85d2ba201233b96e37c2e42c8cad8b291e85ff "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedsignupstatus/count"
)

// PrivilegedSignupStatusRequestBuilder provides operations to manage the collection of privilegedSignupStatus entities.
type PrivilegedSignupStatusRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PrivilegedSignupStatusRequestBuilderGetQueryParameters get entities from privilegedSignupStatus
type PrivilegedSignupStatusRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// PrivilegedSignupStatusRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedSignupStatusRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedSignupStatusRequestBuilderGetQueryParameters
}
// PrivilegedSignupStatusRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedSignupStatusRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CanSignUp provides operations to call the canSignUp method.
func (m *PrivilegedSignupStatusRequestBuilder) CanSignUp()(*i576c2492be4604b88fea75e47108c8bd8db61a3de84b266733d2d7849bec46eb.CanSignUpRequestBuilder) {
    return i576c2492be4604b88fea75e47108c8bd8db61a3de84b266733d2d7849bec46eb.NewCanSignUpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteSetup the completeSetup property
func (m *PrivilegedSignupStatusRequestBuilder) CompleteSetup()(*ib823f830e3ea0500a214b973cf170985ed37b35d907e2b81e06df8aaa8f66256.CompleteSetupRequestBuilder) {
    return ib823f830e3ea0500a214b973cf170985ed37b35d907e2b81e06df8aaa8f66256.NewCompleteSetupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPrivilegedSignupStatusRequestBuilderInternal instantiates a new PrivilegedSignupStatusRequestBuilder and sets the default values.
func NewPrivilegedSignupStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedSignupStatusRequestBuilder) {
    m := &PrivilegedSignupStatusRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedSignupStatus{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPrivilegedSignupStatusRequestBuilder instantiates a new PrivilegedSignupStatusRequestBuilder and sets the default values.
func NewPrivilegedSignupStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedSignupStatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedSignupStatusRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *PrivilegedSignupStatusRequestBuilder) Count()(*id44dae2f191b15a24754b7171e85d2ba201233b96e37c2e42c8cad8b291e85ff.CountRequestBuilder) {
    return id44dae2f191b15a24754b7171e85d2ba201233b96e37c2e42c8cad8b291e85ff.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation get entities from privilegedSignupStatus
func (m *PrivilegedSignupStatusRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get entities from privilegedSignupStatus
func (m *PrivilegedSignupStatusRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *PrivilegedSignupStatusRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation add new entity to privilegedSignupStatus
func (m *PrivilegedSignupStatusRequestBuilder) CreatePostRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedSignupStatusable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration add new entity to privilegedSignupStatus
func (m *PrivilegedSignupStatusRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedSignupStatusable, requestConfiguration *PrivilegedSignupStatusRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get entities from privilegedSignupStatus
func (m *PrivilegedSignupStatusRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedSignupStatusCollectionResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get entities from privilegedSignupStatus
func (m *PrivilegedSignupStatusRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *PrivilegedSignupStatusRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedSignupStatusCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedSignupStatusCollectionResponseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedSignupStatusCollectionResponseable), nil
}
// IsSignedUp provides operations to call the isSignedUp method.
func (m *PrivilegedSignupStatusRequestBuilder) IsSignedUp()(*i505eff958ea7946e725489038a9fbe85c297626b944e12d9c6f68697f1a65258.IsSignedUpRequestBuilder) {
    return i505eff958ea7946e725489038a9fbe85c297626b944e12d9c6f68697f1a65258.NewIsSignedUpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post add new entity to privilegedSignupStatus
func (m *PrivilegedSignupStatusRequestBuilder) Post(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedSignupStatusable)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedSignupStatusable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler add new entity to privilegedSignupStatus
func (m *PrivilegedSignupStatusRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedSignupStatusable, requestConfiguration *PrivilegedSignupStatusRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedSignupStatusable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedSignupStatusFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedSignupStatusable), nil
}
// SignUp the signUp property
func (m *PrivilegedSignupStatusRequestBuilder) SignUp()(*ibe660ff6aa6da80e38b282109151294b9409fe46a692bb37484dd8a1705b5a94.SignUpRequestBuilder) {
    return ibe660ff6aa6da80e38b282109151294b9409fe46a692bb37484dd8a1705b5a94.NewSignUpRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
