package clusters

import (
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/api"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/clusters/ocm"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/clusters/types"
	"github.com/bf2fc6cc711aee1a0c2a/kas-fleet-manager/pkg/config"
	"github.com/pkg/errors"
)

//go:generate moq -out provider_moq.go . Provider
type Provider interface {
	// Create using the information provided to request a new OpenShift/k8s cluster from the provider
	Create(request *types.ClusterRequest) (*types.ClusterSpec, error)
	// Delete delete the cluster from the provider
	Delete(spec *types.ClusterSpec) (bool, error)
	// CheckClusterStatus check the status of the cluster. This will be called periodically during cluster provisioning phase to see if the cluster is ready.
	// It should set the status in the returned `ClusterSpec` to either `provisioning`, `ready` or `failed`.
	// If there is additional data that needs to be preserved and passed between checks, add it to the returned `ClusterSpec` and it will be saved to the database and passed into this function again next time it is called.
	CheckClusterStatus(spec *types.ClusterSpec) (*types.ClusterSpec, error)
	// AddIdentityProvider add an identity provider to the cluster
	AddIdentityProvider(clusterSpec *types.ClusterSpec, identityProvider types.IdentityProviderInfo) (*types.IdentityProviderInfo, error)
	// ApplyResources apply openshift/k8s resources to the cluster
	ApplyResources(clusterSpec *types.ClusterSpec, resources types.ResourceSet) (*types.ResourceSet, error)
	// ScaleUp scale the cluster up with the number of additional nodes specified
	ScaleUp(clusterSpec *types.ClusterSpec, increment int) (*types.ClusterSpec, error)
	// ScaleDown scale the cluster down with the number of nodes specified
	ScaleDown(clusterSpec *types.ClusterSpec, decrement int) (*types.ClusterSpec, error)
	// SetComputeNodes set the number of desired compute nodes for the cluster
	SetComputeNodes(clusterSpec *types.ClusterSpec, numNodes int) (*types.ClusterSpec, error)
	// GetComputeNodes get the number of compute nodes for the cluster
	GetComputeNodes(spec *types.ClusterSpec) (*types.ComputeNodesInfo, error)
	// GetClusterDNS Get the dns of the cluster
	GetClusterDNS(clusterSpec *types.ClusterSpec) (string, error)
	// GetCloudProviders Get the information about supported cloud providers from the cluster provider
	GetCloudProviders() (*types.CloudProviderInfoList, error)
	// GetCloudProviderRegions Get the regions information for the given cloud provider from the cluster provider
	GetCloudProviderRegions(providerInf types.CloudProviderInfo) (*types.CloudProviderRegionInfoList, error)
}

//go:generate moq -out addon_provider_moq.go . AddonProvider
type AddonProvider interface {
	InstallAddon(clusterSpec *types.ClusterSpec, addonID string) (bool, error)
	InstallAddonWithParams(clusterSpec *types.ClusterSpec, addonId string, addonParams []ocm.AddonParameter) (bool, error)
}

// ProviderFactory used to return an instance of Provider implementation
//go:generate moq -out provider_factory_moq.go . ProviderFactory
type ProviderFactory interface {
	GetProvider(providerType api.ClusterProviderType) (Provider, error)
	GetAddonProvider(providerType api.ClusterProviderType) (AddonProvider, error)
}

// DefaultProviderFactory the default implementation for ProviderFactory
type DefaultProviderFactory struct {
	ocmClient ocm.Client
	config    *config.ApplicationConfig

	ocmProvider        *OCMProvider
	standaloneProvider *StandaloneProvider
}

func NewDefaultProviderFactory(ocmClient ocm.Client, appConfig *config.ApplicationConfig) *DefaultProviderFactory {
	return &DefaultProviderFactory{
		ocmClient:          ocmClient,
		config:             appConfig,
		ocmProvider:        nil,
		standaloneProvider: nil,
	}
}

func (d *DefaultProviderFactory) GetProvider(providerType api.ClusterProviderType) (Provider, error) {
	if providerType == "" {
		providerType = api.ClusterProviderOCM
	}
	switch providerType {
	case api.ClusterProviderOCM:
		if d.ocmProvider == nil {
			cb := ocm.NewClusterBuilder(d.config.AWS, d.config.OSDClusterConfig)
			d.ocmProvider = newOCMProvider(d.ocmClient, cb)
		}
		return d.ocmProvider, nil
	case api.ClusterProviderStandalone:
		if d.standaloneProvider == nil {
			d.standaloneProvider = newStandaloneProvider()
		}
		return d.standaloneProvider, nil
	default:
		return nil, errors.Errorf("invalid provider type: %v", providerType)
	}
}

func (d *DefaultProviderFactory) GetAddonProvider(providerType api.ClusterProviderType) (AddonProvider, error) {
	if providerType == "" {
		providerType = api.ClusterProviderOCM
	}
	switch providerType {
	case api.ClusterProviderOCM:
		if d.ocmProvider == nil {
			cb := ocm.NewClusterBuilder(d.config.AWS, d.config.OSDClusterConfig)
			d.ocmProvider = newOCMProvider(d.ocmClient, cb)
		}
		return d.ocmProvider, nil
	default:
		return nil, errors.Errorf("invalid provider type: %v", providerType)
	}
}