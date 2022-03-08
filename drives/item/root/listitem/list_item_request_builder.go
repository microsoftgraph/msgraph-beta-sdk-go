package listitem

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i22e35d2792e670303b941a9b608688f7dc9b73d8a4a6d66d171406ee49bc893b "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i40019e998df2381900e59377a82e6fe6030872422907ef524f26281257ab70c2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/listitem/fields"
    i97dcf288e4fc9f2d9166a701fd4f1c71b926e4775df689aebacf9ca8c393ffe3 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/listitem/activities"
    i99379bb58eacfac530bfe11a684ff121bb27e436e312d0dc1955011ec7d60495 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/listitem/versions"
    i9fe24f6cb7d91d94fc33b6b689f3c0d49e2df0342dc2697dc4d9a0ffcd0b1f76 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/listitem/driveitem"
    id355a249cf15d1eea62c24adf38de7ecde11df4888be34ebf29320703dce751d "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/listitem/createlink"
    if5a67baf1923fef8c0a0c9fc17dec9eaadb90a1759d861715a03bf7f3a128ad7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/listitem/analytics"
    i75b987bc9c33f49af73a7c2a6ad687fae6eb2bb24476249436175b92279ec554 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/listitem/activities/item"
    ide9e4b527e4c13791a6bab833ef60864c5f12eef9e2ef994b8a29d99b3c2045e "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/root/listitem/versions/item"
)

// ListItemRequestBuilder provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ListItemRequestBuilderDeleteOptions options for Delete
type ListItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemRequestBuilderGetOptions options for Get
type ListItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ListItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemRequestBuilderGetQueryParameters for drives in SharePoint, the associated document library list item. Read-only. Nullable.
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ListItemRequestBuilderPatchOptions options for Patch
type ListItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ListItemRequestBuilder) Activities()(*i97dcf288e4fc9f2d9166a701fd4f1c71b926e4775df689aebacf9ca8c393ffe3.ActivitiesRequestBuilder) {
    return i97dcf288e4fc9f2d9166a701fd4f1c71b926e4775df689aebacf9ca8c393ffe3.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.listItem.activities.item collection
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*i75b987bc9c33f49af73a7c2a6ad687fae6eb2bb24476249436175b92279ec554.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i75b987bc9c33f49af73a7c2a6ad687fae6eb2bb24476249436175b92279ec554.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*if5a67baf1923fef8c0a0c9fc17dec9eaadb90a1759d861715a03bf7f3a128ad7.AnalyticsRequestBuilder) {
    return if5a67baf1923fef8c0a0c9fc17dec9eaadb90a1759d861715a03bf7f3a128ad7.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/root/listItem{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemRequestBuilder instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property listItem for drives
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation(options *ListItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) CreateGetRequestInformation(options *ListItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ListItemRequestBuilder) CreateLink()(*id355a249cf15d1eea62c24adf38de7ecde11df4888be34ebf29320703dce751d.CreateLinkRequestBuilder) {
    return id355a249cf15d1eea62c24adf38de7ecde11df4888be34ebf29320703dce751d.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property listItem in drives
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(options *ListItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property listItem for drives
func (m *ListItemRequestBuilder) Delete(options *ListItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) DriveItem()(*i9fe24f6cb7d91d94fc33b6b689f3c0d49e2df0342dc2697dc4d9a0ffcd0b1f76.DriveItemRequestBuilder) {
    return i9fe24f6cb7d91d94fc33b6b689f3c0d49e2df0342dc2697dc4d9a0ffcd0b1f76.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i40019e998df2381900e59377a82e6fe6030872422907ef524f26281257ab70c2.FieldsRequestBuilder) {
    return i40019e998df2381900e59377a82e6fe6030872422907ef524f26281257ab70c2.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) Get(options *ListItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateListItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, interval *string, endDateTime *string)(*i22e35d2792e670303b941a9b608688f7dc9b73d8a4a6d66d171406ee49bc893b.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i22e35d2792e670303b941a9b608688f7dc9b73d8a4a6d66d171406ee49bc893b.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, interval, endDateTime);
}
// Patch update the navigation property listItem in drives
func (m *ListItemRequestBuilder) Patch(options *ListItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) Versions()(*i99379bb58eacfac530bfe11a684ff121bb27e436e312d0dc1955011ec7d60495.VersionsRequestBuilder) {
    return i99379bb58eacfac530bfe11a684ff121bb27e436e312d0dc1955011ec7d60495.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.root.listItem.versions.item collection
func (m *ListItemRequestBuilder) VersionsById(id string)(*ide9e4b527e4c13791a6bab833ef60864c5f12eef9e2ef994b8a29d99b3c2045e.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return ide9e4b527e4c13791a6bab833ef60864c5f12eef9e2ef994b8a29d99b3c2045e.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
