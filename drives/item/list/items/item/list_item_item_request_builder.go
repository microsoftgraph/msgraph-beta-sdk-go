package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i54028a2dda969202a2685ec86a5a8a4e891dafa2c1d74908b1d7d78bf6cb88ac "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/items/item/versions"
    i8e0f7da76797ba3f8bae0d6b0383996b9ab20a49eb4905d6bb6d793dfd61963a "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/items/item/driveitem"
    i969880607fa486281b4aa025225e1041cab6f84a5285cdafea6d91652a7164ae "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/items/item/fields"
    ia5b0b5dbeb80174adc999c070c1fc324825198480c336efaebbec5ba2ccdabf1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/items/item/analytics"
    ib1d9e40882d6f5ef9afe802be79f211fa5b174f9aaaac1bf53a82e537b9fe3e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    ic7171b347963c876fc971bc15a12d3ea614217a23b1dbd21e993da8e3014e3cb "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/items/item/activities"
    idbb93eb42db70975c84850dec586ff94d09778b415741f360894dd8f3c91db4d "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/items/item/createlink"
    id44f41fdff4c40056126341175af022d07e262210aaf29c45e032b601ffa3965 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/items/item/versions/item"
    if862063aea9f6b597354ed27e13c1416d68bb0e7ff4d8cf9c199bed8bd689e2b "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/list/items/item/activities/item"
)

// ListItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.list entity.
type ListItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ListItemItemRequestBuilderDeleteOptions options for Delete
type ListItemItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemItemRequestBuilderGetOptions options for Get
type ListItemItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ListItemItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemItemRequestBuilderGetQueryParameters all items contained in the list.
type ListItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ListItemItemRequestBuilderPatchOptions options for Patch
type ListItemItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ListItemItemRequestBuilder) Activities()(*ic7171b347963c876fc971bc15a12d3ea614217a23b1dbd21e993da8e3014e3cb.ActivitiesRequestBuilder) {
    return ic7171b347963c876fc971bc15a12d3ea614217a23b1dbd21e993da8e3014e3cb.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.list.items.item.activities.item collection
func (m *ListItemItemRequestBuilder) ActivitiesById(id string)(*if862063aea9f6b597354ed27e13c1416d68bb0e7ff4d8cf9c199bed8bd689e2b.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return if862063aea9f6b597354ed27e13c1416d68bb0e7ff4d8cf9c199bed8bd689e2b.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemItemRequestBuilder) Analytics()(*ia5b0b5dbeb80174adc999c070c1fc324825198480c336efaebbec5ba2ccdabf1.AnalyticsRequestBuilder) {
    return ia5b0b5dbeb80174adc999c070c1fc324825198480c336efaebbec5ba2ccdabf1.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemItemRequestBuilderInternal instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewListItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemItemRequestBuilder) {
    m := &ListItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/list/items/{listItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemItemRequestBuilder instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewListItemItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property items for drives
func (m *ListItemItemRequestBuilder) CreateDeleteRequestInformation(options *ListItemItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation all items contained in the list.
func (m *ListItemItemRequestBuilder) CreateGetRequestInformation(options *ListItemItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ListItemItemRequestBuilder) CreateLink()(*idbb93eb42db70975c84850dec586ff94d09778b415741f360894dd8f3c91db4d.CreateLinkRequestBuilder) {
    return idbb93eb42db70975c84850dec586ff94d09778b415741f360894dd8f3c91db4d.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property items in drives
func (m *ListItemItemRequestBuilder) CreatePatchRequestInformation(options *ListItemItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property items for drives
func (m *ListItemItemRequestBuilder) Delete(options *ListItemItemRequestBuilderDeleteOptions)(error) {
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
func (m *ListItemItemRequestBuilder) DriveItem()(*i8e0f7da76797ba3f8bae0d6b0383996b9ab20a49eb4905d6bb6d793dfd61963a.DriveItemRequestBuilder) {
    return i8e0f7da76797ba3f8bae0d6b0383996b9ab20a49eb4905d6bb6d793dfd61963a.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemItemRequestBuilder) Fields()(*i969880607fa486281b4aa025225e1041cab6f84a5285cdafea6d91652a7164ae.FieldsRequestBuilder) {
    return i969880607fa486281b4aa025225e1041cab6f84a5285cdafea6d91652a7164ae.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the list.
func (m *ListItemItemRequestBuilder) Get(options *ListItemItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable, error) {
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
func (m *ListItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, interval *string, endDateTime *string)(*ib1d9e40882d6f5ef9afe802be79f211fa5b174f9aaaac1bf53a82e537b9fe3e4.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return ib1d9e40882d6f5ef9afe802be79f211fa5b174f9aaaac1bf53a82e537b9fe3e4.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, interval, endDateTime);
}
// Patch update the navigation property items in drives
func (m *ListItemItemRequestBuilder) Patch(options *ListItemItemRequestBuilderPatchOptions)(error) {
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
func (m *ListItemItemRequestBuilder) Versions()(*i54028a2dda969202a2685ec86a5a8a4e891dafa2c1d74908b1d7d78bf6cb88ac.VersionsRequestBuilder) {
    return i54028a2dda969202a2685ec86a5a8a4e891dafa2c1d74908b1d7d78bf6cb88ac.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.list.items.item.versions.item collection
func (m *ListItemItemRequestBuilder) VersionsById(id string)(*id44f41fdff4c40056126341175af022d07e262210aaf29c45e032b601ffa3965.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return id44f41fdff4c40056126341175af022d07e262210aaf29c45e032b601ffa3965.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
