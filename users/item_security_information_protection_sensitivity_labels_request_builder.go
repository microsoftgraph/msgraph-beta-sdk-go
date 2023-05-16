package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder provides operations to manage the sensitivityLabels property of the microsoft.graph.security.informationProtection entity.
type ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSecurityInformationProtectionSensitivityLabelsRequestBuilderGetQueryParameters get a list of sensitivityLabel objects associated with a user or organization.
type ItemSecurityInformationProtectionSensitivityLabelsRequestBuilderGetQueryParameters struct {
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
// ItemSecurityInformationProtectionSensitivityLabelsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationProtectionSensitivityLabelsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilderGetQueryParameters
}
// ItemSecurityInformationProtectionSensitivityLabelsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSecurityInformationProtectionSensitivityLabelsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BySensitivityLabelId provides operations to manage the sensitivityLabels property of the microsoft.graph.security.informationProtection entity.
func (m *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder) BySensitivityLabelId(sensitivityLabelId string)(*ItemSecurityInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if sensitivityLabelId != "" {
        urlTplParams["sensitivityLabel%2Did"] = sensitivityLabelId
    }
    return NewItemSecurityInformationProtectionSensitivityLabelsSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSecurityInformationProtectionSensitivityLabelsRequestBuilderInternal instantiates a new SensitivityLabelsRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder) {
    m := &ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/security/informationProtection/sensitivityLabels{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemSecurityInformationProtectionSensitivityLabelsRequestBuilder instantiates a new SensitivityLabelsRequestBuilder and sets the default values.
func NewItemSecurityInformationProtectionSensitivityLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSecurityInformationProtectionSensitivityLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder) Count()(*ItemSecurityInformationProtectionSensitivityLabelsCountRequestBuilder) {
    return NewItemSecurityInformationProtectionSensitivityLabelsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of sensitivityLabel objects associated with a user or organization.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/security-informationprotection-list-sensitivitylabels?view=graph-rest-1.0
func (m *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateSensitivityLabelCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelCollectionResponseable), nil
}
// MicrosoftGraphSecurityEvaluateApplication provides operations to call the evaluateApplication method.
func (m *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder) MicrosoftGraphSecurityEvaluateApplication()(*ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilder) {
    return NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityEvaluateClassificationResults provides operations to call the evaluateClassificationResults method.
func (m *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder) MicrosoftGraphSecurityEvaluateClassificationResults()(*ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilder) {
    return NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateClassificationResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityEvaluateRemoval provides operations to call the evaluateRemoval method.
func (m *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder) MicrosoftGraphSecurityEvaluateRemoval()(*ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilder) {
    return NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityEvaluateRemovalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityExtractContentLabel provides operations to call the extractContentLabel method.
func (m *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder) MicrosoftGraphSecurityExtractContentLabel()(*ItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilder) {
    return NewItemSecurityInformationProtectionSensitivityLabelsMicrosoftGraphSecurityExtractContentLabelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to sensitivityLabels for users
func (m *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder) Post(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilderPostRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateSensitivityLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable), nil
}
// ToGetRequestInformation get a list of sensitivityLabel objects associated with a user or organization.
func (m *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation create new navigation property to sensitivityLabels for users
func (m *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.SensitivityLabelable, requestConfiguration *ItemSecurityInformationProtectionSensitivityLabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
