package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    ic425ce86de24a034e6f015478562056e313cd3283aeee6815e39b42a068dd2e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/categories/item/recommendedsettings"
    i8cca88b6428aa5e0aef1f88b2eea712467f69b4bcd843cf74b8d35246b3c0873 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/templates/item/categories/item/recommendedsettings/item"
)

// DeviceManagementTemplateSettingCategoryItemRequestBuilder builds and executes requests for operations under \deviceManagement\templates\{deviceManagementTemplate-id}\categories\{deviceManagementTemplateSettingCategory-id}
type DeviceManagementTemplateSettingCategoryItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementTemplateSettingCategoryItemRequestBuilderDeleteOptions options for Delete
type DeviceManagementTemplateSettingCategoryItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementTemplateSettingCategoryItemRequestBuilderGetOptions options for Get
type DeviceManagementTemplateSettingCategoryItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementTemplateSettingCategoryItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementTemplateSettingCategoryItemRequestBuilderGetQueryParameters collection of setting categories within the template
type DeviceManagementTemplateSettingCategoryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementTemplateSettingCategoryItemRequestBuilderPatchOptions options for Patch
type DeviceManagementTemplateSettingCategoryItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementTemplateSettingCategory;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal instantiates a new DeviceManagementTemplateSettingCategoryItemRequestBuilder and sets the default values.
func NewDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementTemplateSettingCategoryItemRequestBuilder) {
    m := &DeviceManagementTemplateSettingCategoryItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate_id}/categories/{deviceManagementTemplateSettingCategory_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementTemplateSettingCategoryItemRequestBuilder instantiates a new DeviceManagementTemplateSettingCategoryItemRequestBuilder and sets the default values.
func NewDeviceManagementTemplateSettingCategoryItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementTemplateSettingCategoryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementTemplateSettingCategoryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation collection of setting categories within the template
func (m *DeviceManagementTemplateSettingCategoryItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementTemplateSettingCategoryItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation collection of setting categories within the template
func (m *DeviceManagementTemplateSettingCategoryItemRequestBuilder) CreateGetRequestInformation(options *DeviceManagementTemplateSettingCategoryItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation collection of setting categories within the template
func (m *DeviceManagementTemplateSettingCategoryItemRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementTemplateSettingCategoryItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete collection of setting categories within the template
func (m *DeviceManagementTemplateSettingCategoryItemRequestBuilder) Delete(options *DeviceManagementTemplateSettingCategoryItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get collection of setting categories within the template
func (m *DeviceManagementTemplateSettingCategoryItemRequestBuilder) Get(options *DeviceManagementTemplateSettingCategoryItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementTemplateSettingCategory, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewDeviceManagementTemplateSettingCategory() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.DeviceManagementTemplateSettingCategory), nil
}
// Patch collection of setting categories within the template
func (m *DeviceManagementTemplateSettingCategoryItemRequestBuilder) Patch(options *DeviceManagementTemplateSettingCategoryItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceManagementTemplateSettingCategoryItemRequestBuilder) RecommendedSettings()(*ic425ce86de24a034e6f015478562056e313cd3283aeee6815e39b42a068dd2e1.RecommendedSettingsRequestBuilder) {
    return ic425ce86de24a034e6f015478562056e313cd3283aeee6815e39b42a068dd2e1.NewRecommendedSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecommendedSettingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.templates.item.categories.item.recommendedSettings.item collection
func (m *DeviceManagementTemplateSettingCategoryItemRequestBuilder) RecommendedSettingsById(id string)(*i8cca88b6428aa5e0aef1f88b2eea712467f69b4bcd843cf74b8d35246b3c0873.DeviceManagementSettingInstanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementSettingInstance_id"] = id
    }
    return i8cca88b6428aa5e0aef1f88b2eea712467f69b4bcd843cf74b8d35246b3c0873.NewDeviceManagementSettingInstanceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}