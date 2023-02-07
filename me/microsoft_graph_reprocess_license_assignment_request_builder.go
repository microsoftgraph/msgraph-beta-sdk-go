package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftGraphReprocessLicenseAssignmentRequestBuilder provides operations to call the reprocessLicenseAssignment method.
type MicrosoftGraphReprocessLicenseAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphReprocessLicenseAssignmentRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphReprocessLicenseAssignmentRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphReprocessLicenseAssignmentRequestBuilderInternal instantiates a new MicrosoftGraphReprocessLicenseAssignmentRequestBuilder and sets the default values.
func NewMicrosoftGraphReprocessLicenseAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphReprocessLicenseAssignmentRequestBuilder) {
    m := &MicrosoftGraphReprocessLicenseAssignmentRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.reprocessLicenseAssignment";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphReprocessLicenseAssignmentRequestBuilder instantiates a new MicrosoftGraphReprocessLicenseAssignmentRequestBuilder and sets the default values.
func NewMicrosoftGraphReprocessLicenseAssignmentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphReprocessLicenseAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphReprocessLicenseAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reprocess all group-based license assignments for the user. To learn more about group-based licensing, see What is group-based licensing in Azure Active Directory. Also see Identify and resolve license assignment problems for a group in Azure Active Directory for more details.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/user-reprocesslicenseassignment?view=graph-rest-1.0
func (m *MicrosoftGraphReprocessLicenseAssignmentRequestBuilder) Post(ctx context.Context, requestConfiguration *MicrosoftGraphReprocessLicenseAssignmentRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Userable), nil
}
// ToPostRequestInformation reprocess all group-based license assignments for the user. To learn more about group-based licensing, see What is group-based licensing in Azure Active Directory. Also see Identify and resolve license assignment problems for a group in Azure Active Directory for more details.
func (m *MicrosoftGraphReprocessLicenseAssignmentRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphReprocessLicenseAssignmentRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
