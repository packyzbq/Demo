package apiserver

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/version"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

var (
	Scheme = runtime.NewScheme()
	Codecs = serializer.NewCodecFactory(Scheme)
)

type ExtraConfig struct {
	// Place you custom config here.
}

type Config struct {
	GenericConfig *genericapiserver.RecommendedConfig
	ExtraConfig   ExtraConfig
}

// CustomServer contains state for a Kubernetes cluster master/api cmd.
type CustomServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

type completedConfig struct {
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}

type CompletedConfig struct {
	// Embed a private pointer that cannot be instantiated outside of this package.
	*completedConfig
}

// Complete fills in any fields not set that are required to have valid data. It's mutating the receiver.
// 以默认选项配置必要项
func (cfg *Config) Complete() CompletedConfig {
	c := completedConfig{
		cfg.GenericConfig.Complete(),
		&cfg.ExtraConfig,
	}

	c.GenericConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}

	return CompletedConfig{&c}
}

// New returns a new instance of CustomServer from the given config.
func (c completedConfig) New() (*CustomServer, error) {
	genericServer, err := c.GenericConfig.New("pizza-apiserver", genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	s := &CustomServer{
		GenericAPIServer: genericServer,
	}

	//apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(restaurant.GroupName, Scheme, metav1.ParameterCodec, Codecs)
	//
	//v1alpha1storage := map[string]rest.Storage{}
	//v1alpha1storage["pizzas"] = customregistry.RESTInPeace(pizzastorage.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter))
	//v1alpha1storage["toppings"] = customregistry.RESTInPeace(toppingstorage.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter))
	//apiGroupInfo.VersionedResourcesStorageMap["v1alpha1"] = v1alpha1storage
	//
	//v1beta1storage := map[string]rest.Storage{}
	//v1beta1storage["pizzas"] = customregistry.RESTInPeace(pizzastorage.NewREST(Scheme, c.GenericConfig.RESTOptionsGetter))
	//apiGroupInfo.VersionedResourcesStorageMap["v1beta1"] = v1beta1storage
	//
	//if err := s.GenericAPIServer.InstallAPIGroup(&apiGroupInfo); err != nil {
	//	return nil, err
	//}

	return s, nil
}
