package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder provides operations to manage the appliesTo property of the microsoft.graph.featureRolloutPolicy entity.
type PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderGetQueryParameters nullable. Specifies a list of directoryObjects that feature is enabled for.
type PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderGetQueryParameters struct {
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
// PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderGetQueryParameters
}
// PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderInternal instantiates a new AppliesToRequestBuilder and sets the default values.
func NewPoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder) {
    m := &PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/policies/featureRolloutPolicies/{featureRolloutPolicy%2Did}/appliesTo{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder instantiates a new AppliesToRequestBuilder and sets the default values.
func NewPoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder) Count()(*PoliciesFeatureRolloutPoliciesItemAppliesToCountRequestBuilder) {
    return NewPoliciesFeatureRolloutPoliciesItemAppliesToCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation nullable. Specifies a list of directoryObjects that feature is enabled for.
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePostRequestInformation create new navigation property to appliesTo for policies
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder) CreatePostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get nullable. Specifies a list of directoryObjects that feature is enabled for.
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder) Get(ctx context.Context, requestConfiguration *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// GetByIds provides operations to call the getByIds method.
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder) GetByIds()(*PoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilder) {
    return NewPoliciesFeatureRolloutPoliciesItemAppliesToGetByIdsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetUserOwnedObjects provides operations to call the getUserOwnedObjects method.
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder) GetUserOwnedObjects()(*PoliciesFeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilder) {
    return NewPoliciesFeatureRolloutPoliciesItemAppliesToGetUserOwnedObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Post create new navigation property to appliesTo for policies
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, requestConfiguration *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectable), nil
}
// Ref provides operations to manage the collection of policyRoot entities.
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder) Ref()(*PoliciesFeatureRolloutPoliciesItemAppliesToRefRequestBuilder) {
    return NewPoliciesFeatureRolloutPoliciesItemAppliesToRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ValidateProperties provides operations to call the validateProperties method.
func (m *PoliciesFeatureRolloutPoliciesItemAppliesToRequestBuilder) ValidateProperties()(*PoliciesFeatureRolloutPoliciesItemAppliesToValidatePropertiesRequestBuilder) {
    return NewPoliciesFeatureRolloutPoliciesItemAppliesToValidatePropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
