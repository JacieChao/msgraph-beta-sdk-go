package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder provides operations to manage the permissionsCreepIndexDistributions property of the microsoft.graph.permissionsAnalytics entity.
type PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetQueryParameters get permissionsCreepIndexDistributions from identityGovernance
type PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetQueryParameters
}
// PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthorizationSystem provides operations to manage the authorizationSystem property of the microsoft.graph.permissionsCreepIndexDistribution entity.
func (m *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) AuthorizationSystem()(*PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilder) {
    return NewPermissionsAnalyticsGcpPermissionsCreepIndexDistributionsItemAuthorizationSystemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal instantiates a new PermissionsCreepIndexDistributionItemRequestBuilder and sets the default values.
func NewPermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    m := &PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/permissionsAnalytics/gcp/permissionsCreepIndexDistributions/{permissionsCreepIndexDistribution%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewPermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder instantiates a new PermissionsCreepIndexDistributionItemRequestBuilder and sets the default values.
func NewPermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property permissionsCreepIndexDistributions for identityGovernance
func (m *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get permissionsCreepIndexDistributions from identityGovernance
func (m *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsCreepIndexDistributionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable), nil
}
// Patch update the navigation property permissionsCreepIndexDistributions in identityGovernance
func (m *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, requestConfiguration *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePermissionsCreepIndexDistributionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable), nil
}
// ToDeleteRequestInformation delete navigation property permissionsCreepIndexDistributions for identityGovernance
func (m *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get permissionsCreepIndexDistributions from identityGovernance
func (m *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property permissionsCreepIndexDistributions in identityGovernance
func (m *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PermissionsCreepIndexDistributionable, requestConfiguration *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) WithUrl(rawUrl string)(*PermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder) {
    return NewPermissionsAnalyticsGcpPermissionsCreepIndexDistributionsPermissionsCreepIndexDistributionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
