package cmd

import (
	"fmt"
	_ "github.com/spf13/cobra"
	"k8s.io/api/node/v1alpha1"
	_ "k8s.io/apiserver/pkg/admission"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/client-go/informers"
	"net"

	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

const defaultEtcdPathPrefix = "/registry/pizza-store.packyzbq.io"

type CustomServerOptions struct {
	RecommandOptions      *genericoptions.RecommendedOptions
	SharedInformerFactory informers.SharedInformerFactory
}

func NewCustomServerOptions() *CustomServerOptions {
	o := &CustomServerOptions{
		RecommandOptions: genericoptions.NewRecommendedOptions(
			defaultEtcdPathPrefix,
			apiserver.Codecs.LegacyCodec(v1alpha1.SchemeGroupVersion),
			genericoptions.NewProcessInfo("pizza-apiserver", "pizza-apiserver"),
		),
	}
	return o
}

//启动 api server
func (o CustomServerOptions) Run(stopCh <-chan struct{}) error {
	config, err := o.Config()
	if err != nil {
		return err
	}

	server, err := config.Complete().New()
	if err != nil {
		return err
	}

	_ = server.GenericAPIServer.AddPostStartHook("start-pizza-apiserver-informers",
		func(context genericapiserver.PostStartHookContext) error {
			config.GenericConfig.SharedInformerFactory.Start(context.StopCh)
			o.SharedInformerFactory.Start(context.StopCh)
			return nil
		},
	)

	return server.GenericAPIServer.PrepareRun().Run(stopCh)
}

// NewCommandStartCustomServer provides a CLI handler for 'start master' command
// with a default CustomServerOptions.
//func NewCommandStartCustomServer(defaults *CustomServerOptions, stopCh <-chan struct{}) *cobra.Command {
//	o := *defaults
//	cmd := &cobra.Command{
//		Short: "Launch a custom API cmd",
//		Long:  "Launch a custom API cmd",
//		RunE: func(c *cobra.Command, args []string) error {
//			if err := o.Complete(); err != nil {
//				return err
//			}
//			if err := o.Validate(); err != nil {
//				return err
//			}
//			if err := o.Run(stopCh); err != nil {
//				return err
//			}
//			return nil
//		},
//	}
//
//	flags := cmd.Flags()
//	o.RecommandOptions.AddFlags(flags)
//
//	return cmd
//}

func (o CustomServerOptions) Validate() error {
	errors := []error{}
	errors = append(errors, o.RecommandOptions.Validate()...)
	return utilerrors.NewAggregate(errors)
}

//func (o *CustomServerOptions) Complete() error {
//	// register admission plugins
//	pizzatoppings.Register(o.RecommandOptions.Admission.Plugins)
//
//	// add admisison plugins to the RecommendedPluginOrder
//	o.RecommandOptions.Admission.RecommendedPluginOrder = append(o.RecommandOptions.Admission.RecommendedPluginOrder, "PizzaToppings")
//
//	return nil
//}

// 将 Options 转换为 api server 启动时需要的 config，
func (o *CustomServerOptions) Config() (*apiserver.Config, error) {
	// TODO have a "real" external address
	if err := o.RecommandOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	// 自定义api相关逻辑
	//o.RecommandOptions.ExtraAdmissionInitializers = func(c *genericapiserver.RecommendedConfig) ([]admission.PluginInitializer, error) {
	//	client, err := clientset.NewForConfig(c.LoopbackClientConfig)
	//	if err != nil {
	//		return nil, err
	//	}
	//	informerFactory := informers.NewSharedInformerFactory(client, c.LoopbackClientConfig.Timeout)
	//	o.SharedInformerFactory = informerFactory
	//	return []admission.PluginInitializer{custominitializer.New(informerFactory)}, nil
	//}

	serverConfig := genericapiserver.NewRecommendedConfig(apiserver.Codecs)
	if err := o.RecommandOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}

	config := &apiserver.Config{
		GenericConfig: serverConfig,
		ExtraConfig:   apiserver.ExtraConfig{},
	}
	return config, nil
}
