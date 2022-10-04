package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i300bb30a733f11af1db0965deae41bf705649be04be882517d59a871d06c607d "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/policysets/item/items"
    i8571cd805b9deabc3a16f61b4629f7bf1382d50a1782362d06f4aef6c29369b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/policysets/item/update"
    ic3b46ce500546d2dc4f9ff8392cc8eb4e3e9d6daf420cff3c3fd9477b80afa9a "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/policysets/item/assignments"
    i24f7e4e5877bd5856a1bea9781fe340d00b7d49e45ba3821ab75b408464158c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/policysets/item/assignments/item"
    i99026bae59eb41024babc35aa8e0359f7e3892bf2e64ce282f1766881b0a8dc1 "github.com/microsoftgraph/msgraph-beta-sdk-go/deviceappmanagement/policysets/item/items/item"
)

// PolicySetItemRequestBuilder provides operations to manage the policySets property of the microsoft.graph.deviceAppManagement entity.
type PolicySetItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PolicySetItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PolicySetItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PolicySetItemRequestBuilderGetQueryParameters the PolicySet of Policies and Applications
type PolicySetItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PolicySetItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PolicySetItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PolicySetItemRequestBuilderGetQueryParameters
}
// PolicySetItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PolicySetItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments the assignments property
func (m *PolicySetItemRequestBuilder) Assignments()(*ic3b46ce500546d2dc4f9ff8392cc8eb4e3e9d6daf420cff3c3fd9477b80afa9a.AssignmentsRequestBuilder) {
    return ic3b46ce500546d2dc4f9ff8392cc8eb4e3e9d6daf420cff3c3fd9477b80afa9a.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.policySets.item.assignments.item collection
func (m *PolicySetItemRequestBuilder) AssignmentsById(id string)(*i24f7e4e5877bd5856a1bea9781fe340d00b7d49e45ba3821ab75b408464158c4.PolicySetAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["policySetAssignment%2Did"] = id
    }
    return i24f7e4e5877bd5856a1bea9781fe340d00b7d49e45ba3821ab75b408464158c4.NewPolicySetAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPolicySetItemRequestBuilderInternal instantiates a new PolicySetItemRequestBuilder and sets the default values.
func NewPolicySetItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicySetItemRequestBuilder) {
    m := &PolicySetItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/policySets/{policySet%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPolicySetItemRequestBuilder instantiates a new PolicySetItemRequestBuilder and sets the default values.
func NewPolicySetItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PolicySetItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPolicySetItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property policySets for deviceAppManagement
func (m *PolicySetItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *PolicySetItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the PolicySet of Policies and Applications
func (m *PolicySetItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *PolicySetItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property policySets in deviceAppManagement
func (m *PolicySetItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetable, requestConfiguration *PolicySetItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property policySets for deviceAppManagement
func (m *PolicySetItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PolicySetItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the PolicySet of Policies and Applications
func (m *PolicySetItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PolicySetItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePolicySetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetable), nil
}
// Items the items property
func (m *PolicySetItemRequestBuilder) Items()(*i300bb30a733f11af1db0965deae41bf705649be04be882517d59a871d06c607d.ItemsRequestBuilder) {
    return i300bb30a733f11af1db0965deae41bf705649be04be882517d59a871d06c607d.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceAppManagement.policySets.item.items.item collection
func (m *PolicySetItemRequestBuilder) ItemsById(id string)(*i99026bae59eb41024babc35aa8e0359f7e3892bf2e64ce282f1766881b0a8dc1.PolicySetItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["policySetItem%2Did"] = id
    }
    return i99026bae59eb41024babc35aa8e0359f7e3892bf2e64ce282f1766881b0a8dc1.NewPolicySetItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property policySets in deviceAppManagement
func (m *PolicySetItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetable, requestConfiguration *PolicySetItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePolicySetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PolicySetable), nil
}
// Update the update property
func (m *PolicySetItemRequestBuilder) Update()(*i8571cd805b9deabc3a16f61b4629f7bf1382d50a1782362d06f4aef6c29369b8.UpdateRequestBuilder) {
    return i8571cd805b9deabc3a16f61b4629f7bf1382d50a1782362d06f4aef6c29369b8.NewUpdateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
