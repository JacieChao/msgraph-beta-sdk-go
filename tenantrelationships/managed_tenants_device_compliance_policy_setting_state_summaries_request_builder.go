package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/managedtenants"
)

// ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
type ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderGetQueryParameters summary information for device compliance policy setting states across managed tenants.
type ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderGetQueryParameters struct {
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
// ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderGetQueryParameters
}
// ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceCompliancePolicySettingStateSummaryId provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.managedTenants.managedTenant entity.
func (m *ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder) ByDeviceCompliancePolicySettingStateSummaryId(deviceCompliancePolicySettingStateSummaryId string)(*ManagedTenantsDeviceCompliancePolicySettingStateSummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceCompliancePolicySettingStateSummaryId != "" {
        urlTplParams["deviceCompliancePolicySettingStateSummary%2Did"] = deviceCompliancePolicySettingStateSummaryId
    }
    return NewManagedTenantsDeviceCompliancePolicySettingStateSummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal instantiates a new DeviceCompliancePolicySettingStateSummariesRequestBuilder and sets the default values.
func NewManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    m := &ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/managedTenants/deviceCompliancePolicySettingStateSummaries{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder instantiates a new DeviceCompliancePolicySettingStateSummariesRequestBuilder and sets the default values.
func NewManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder) Count()(*ManagedTenantsDeviceCompliancePolicySettingStateSummariesCountRequestBuilder) {
    return NewManagedTenantsDeviceCompliancePolicySettingStateSummariesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get summary information for device compliance policy setting states across managed tenants.
func (m *ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderGetRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceCompliancePolicySettingStateSummaryCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateDeviceCompliancePolicySettingStateSummaryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceCompliancePolicySettingStateSummaryCollectionResponseable), nil
}
// Post create new navigation property to deviceCompliancePolicySettingStateSummaries for tenantRelationships
func (m *ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder) Post(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceCompliancePolicySettingStateSummaryable, requestConfiguration *ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderPostRequestConfiguration)(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceCompliancePolicySettingStateSummaryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.CreateDeviceCompliancePolicySettingStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceCompliancePolicySettingStateSummaryable), nil
}
// ToGetRequestInformation summary information for device compliance policy setting states across managed tenants.
func (m *ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceCompliancePolicySettingStateSummaries for tenantRelationships
func (m *ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilder) ToPostRequestInformation(ctx context.Context, body i72d786f54cc0bb289c971b085dd642b2fc3af6394328682e69783fd7e229b582.DeviceCompliancePolicySettingStateSummaryable, requestConfiguration *ManagedTenantsDeviceCompliancePolicySettingStateSummariesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
