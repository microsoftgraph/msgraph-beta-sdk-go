package informationprotection

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0277c2a76ea648d8b78b47eb740e4426e8b70cafa779ae7cfc25d4423af78cd7 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/bitlocker"
    i216e4fe6427aea6a66c8bb5a2498b41402122929dcc455372eacd46b1a7f45b8 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/sensitivitypolicysettings"
    i6ce06805c7dd5778af7b3657ba097e40ca55b94215fb7f2d731ec82a57198d25 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/threatassessmentrequests"
    i8810e040e3ab5580410e10af11c08bbf8c2f5d7c936b0901f09b6047e80747f4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/policy"
    ia8c5de9643bc2b5ba04000acd631436eef2e4037d151c2c9be0fa9de45b5db4e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/sensitivitylabels"
    ib9aa15bb15b353f1d239567b5090a9ed690a845fd979fc5fc5236c37f2c040ec "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/datalosspreventionpolicies"
    i400f881fad4edba6edc2afa2747654d483b40cd2db34116a77814ec312135432 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/threatassessmentrequests/item"
    iafc786e37c049dc894b91f10f3a1a6a1d5d6c2a2aadbde9b9478f38371b96dd6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/datalosspreventionpolicies/item"
    ic2a0f6c7e03401afbd27ed2c227884da9fe6f7c80345dc54e72b4c6d67543da9 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/informationprotection/sensitivitylabels/item"
)

// InformationProtectionRequestBuilder provides operations to manage the informationProtection property of the microsoft.graph.user entity.
type InformationProtectionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// InformationProtectionRequestBuilderDeleteOptions options for Delete
type InformationProtectionRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// InformationProtectionRequestBuilderGetOptions options for Get
type InformationProtectionRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *InformationProtectionRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// InformationProtectionRequestBuilderGetQueryParameters get informationProtection from me
type InformationProtectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// InformationProtectionRequestBuilderPatchOptions options for Patch
type InformationProtectionRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Bitlocker the bitlocker property
func (m *InformationProtectionRequestBuilder) Bitlocker()(*i0277c2a76ea648d8b78b47eb740e4426e8b70cafa779ae7cfc25d4423af78cd7.BitlockerRequestBuilder) {
    return i0277c2a76ea648d8b78b47eb740e4426e8b70cafa779ae7cfc25d4423af78cd7.NewBitlockerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewInformationProtectionRequestBuilderInternal instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewInformationProtectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionRequestBuilder) {
    m := &InformationProtectionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/informationProtection{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewInformationProtectionRequestBuilder instantiates a new InformationProtectionRequestBuilder and sets the default values.
func NewInformationProtectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*InformationProtectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewInformationProtectionRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property informationProtection for me
func (m *InformationProtectionRequestBuilder) CreateDeleteRequestInformation(options *InformationProtectionRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get informationProtection from me
func (m *InformationProtectionRequestBuilder) CreateGetRequestInformation(options *InformationProtectionRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property informationProtection in me
func (m *InformationProtectionRequestBuilder) CreatePatchRequestInformation(options *InformationProtectionRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DataLossPreventionPolicies the dataLossPreventionPolicies property
func (m *InformationProtectionRequestBuilder) DataLossPreventionPolicies()(*ib9aa15bb15b353f1d239567b5090a9ed690a845fd979fc5fc5236c37f2c040ec.DataLossPreventionPoliciesRequestBuilder) {
    return ib9aa15bb15b353f1d239567b5090a9ed690a845fd979fc5fc5236c37f2c040ec.NewDataLossPreventionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DataLossPreventionPoliciesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.informationProtection.dataLossPreventionPolicies.item collection
func (m *InformationProtectionRequestBuilder) DataLossPreventionPoliciesById(id string)(*iafc786e37c049dc894b91f10f3a1a6a1d5d6c2a2aadbde9b9478f38371b96dd6.DataLossPreventionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["dataLossPreventionPolicy%2Did"] = id
    }
    return iafc786e37c049dc894b91f10f3a1a6a1d5d6c2a2aadbde9b9478f38371b96dd6.NewDataLossPreventionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property informationProtection for me
func (m *InformationProtectionRequestBuilder) Delete(options *InformationProtectionRequestBuilderDeleteOptions)(error) {
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
// Get get informationProtection from me
func (m *InformationProtectionRequestBuilder) Get(options *InformationProtectionRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateInformationProtectionFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.InformationProtectionable), nil
}
// Patch update the navigation property informationProtection in me
func (m *InformationProtectionRequestBuilder) Patch(options *InformationProtectionRequestBuilderPatchOptions)(error) {
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
// Policy the policy property
func (m *InformationProtectionRequestBuilder) Policy()(*i8810e040e3ab5580410e10af11c08bbf8c2f5d7c936b0901f09b6047e80747f4.PolicyRequestBuilder) {
    return i8810e040e3ab5580410e10af11c08bbf8c2f5d7c936b0901f09b6047e80747f4.NewPolicyRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabels the sensitivityLabels property
func (m *InformationProtectionRequestBuilder) SensitivityLabels()(*ia8c5de9643bc2b5ba04000acd631436eef2e4037d151c2c9be0fa9de45b5db4e.SensitivityLabelsRequestBuilder) {
    return ia8c5de9643bc2b5ba04000acd631436eef2e4037d151c2c9be0fa9de45b5db4e.NewSensitivityLabelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SensitivityLabelsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.informationProtection.sensitivityLabels.item collection
func (m *InformationProtectionRequestBuilder) SensitivityLabelsById(id string)(*ic2a0f6c7e03401afbd27ed2c227884da9fe6f7c80345dc54e72b4c6d67543da9.SensitivityLabelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sensitivityLabel%2Did"] = id
    }
    return ic2a0f6c7e03401afbd27ed2c227884da9fe6f7c80345dc54e72b4c6d67543da9.NewSensitivityLabelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SensitivityPolicySettings the sensitivityPolicySettings property
func (m *InformationProtectionRequestBuilder) SensitivityPolicySettings()(*i216e4fe6427aea6a66c8bb5a2498b41402122929dcc455372eacd46b1a7f45b8.SensitivityPolicySettingsRequestBuilder) {
    return i216e4fe6427aea6a66c8bb5a2498b41402122929dcc455372eacd46b1a7f45b8.NewSensitivityPolicySettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatAssessmentRequests the threatAssessmentRequests property
func (m *InformationProtectionRequestBuilder) ThreatAssessmentRequests()(*i6ce06805c7dd5778af7b3657ba097e40ca55b94215fb7f2d731ec82a57198d25.ThreatAssessmentRequestsRequestBuilder) {
    return i6ce06805c7dd5778af7b3657ba097e40ca55b94215fb7f2d731ec82a57198d25.NewThreatAssessmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThreatAssessmentRequestsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.informationProtection.threatAssessmentRequests.item collection
func (m *InformationProtectionRequestBuilder) ThreatAssessmentRequestsById(id string)(*i400f881fad4edba6edc2afa2747654d483b40cd2db34116a77814ec312135432.ThreatAssessmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["threatAssessmentRequest%2Did"] = id
    }
    return i400f881fad4edba6edc2afa2747654d483b40cd2db34116a77814ec312135432.NewThreatAssessmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
