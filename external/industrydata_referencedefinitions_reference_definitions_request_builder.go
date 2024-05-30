package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder provides operations to manage the referenceDefinitions property of the microsoft.graph.industryData.industryDataRoot entity.
type IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderGetQueryParameters get a list of the referenceDefinition objects and their properties.
type IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderGetQueryParameters struct {
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
// IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderGetQueryParameters
}
// IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByReferenceDefinitionId provides operations to manage the referenceDefinitions property of the microsoft.graph.industryData.industryDataRoot entity.
// returns a *IndustrydataReferencedefinitionsReferenceDefinitionItemRequestBuilder when successful
func (m *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder) ByReferenceDefinitionId(referenceDefinitionId string)(*IndustrydataReferencedefinitionsReferenceDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if referenceDefinitionId != "" {
        urlTplParams["referenceDefinition%2Did"] = referenceDefinitionId
    }
    return NewIndustrydataReferencedefinitionsReferenceDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewIndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderInternal instantiates a new IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder and sets the default values.
func NewIndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder) {
    m := &IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/referenceDefinitions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewIndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder instantiates a new IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder and sets the default values.
func NewIndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *IndustrydataReferencedefinitionsCountRequestBuilder when successful
func (m *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder) Count()(*IndustrydataReferencedefinitionsCountRequestBuilder) {
    return NewIndustrydataReferencedefinitionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the referenceDefinition objects and their properties.
// returns a ReferenceDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/industrydata-referencedefinition-list?view=graph-rest-beta
func (m *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ReferenceDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateReferenceDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ReferenceDefinitionCollectionResponseable), nil
}
// Post create new navigation property to referenceDefinitions for external
// returns a ReferenceDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder) Post(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ReferenceDefinitionable, requestConfiguration *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderPostRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ReferenceDefinitionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateReferenceDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ReferenceDefinitionable), nil
}
// ToGetRequestInformation get a list of the referenceDefinition objects and their properties.
// returns a *RequestInformation when successful
func (m *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to referenceDefinitions for external
// returns a *RequestInformation when successful
func (m *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.ReferenceDefinitionable, requestConfiguration *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder when successful
func (m *IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder) WithUrl(rawUrl string)(*IndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder) {
    return NewIndustrydataReferencedefinitionsReferenceDefinitionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
