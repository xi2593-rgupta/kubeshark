package cmd

import (
	"errors"
	"fmt"

	"github.com/creasty/defaults"
	"github.com/kubeshark/kubeshark/config"
	"github.com/kubeshark/kubeshark/config/configStructs"
	"github.com/kubeshark/kubeshark/errormessage"
	"github.com/kubeshark/kubeshark/misc"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var tapCmd = &cobra.Command{
	Use:   "tap [POD REGEX]",
	Short: "Capture the network traffic in your Kubernetes cluster",
	RunE: func(cmd *cobra.Command, args []string) error {
		tap()
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			config.Config.Tap.PodRegexStr = args[0]
		} else if len(args) > 1 {
			return errors.New("unexpected number of arguments")
		}

		if err := config.Config.Tap.Validate(); err != nil {
			return errormessage.FormatError(err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(tapCmd)

	defaultTapConfig := configStructs.TapConfig{}
	if err := defaults.Set(&defaultTapConfig); err != nil {
		log.Debug().Err(err).Send()
	}

	tapCmd.Flags().StringP(configStructs.DockerRegistryLabel, "r", defaultTapConfig.Docker.Registry, "The Docker registry that's hosting the images")
	tapCmd.Flags().StringP(configStructs.DockerTagLabel, "t", defaultTapConfig.Docker.Tag, "The tag of the Docker images that are going to be pulled")
	tapCmd.Flags().String(configStructs.DockerImagePullPolicy, defaultTapConfig.Docker.ImagePullPolicy, "ImagePullPolicy for the Docker images")
	tapCmd.Flags().StringSlice(configStructs.DockerImagePullSecrets, defaultTapConfig.Docker.ImagePullSecrets, "ImagePullSecrets for the Docker images")
	tapCmd.Flags().Uint16(configStructs.ProxyFrontPortLabel, defaultTapConfig.Proxy.Front.SrcPort, "Provide a custom port for the front-end proxy/port-forward")
	tapCmd.Flags().Uint16(configStructs.ProxyHubPortLabel, defaultTapConfig.Proxy.Hub.SrcPort, "Provide a custom port for the Hub proxy/port-forward")
	tapCmd.Flags().String(configStructs.ProxyHostLabel, defaultTapConfig.Proxy.Host, "Provide a custom host for the proxy/port-forward")
	tapCmd.Flags().StringSliceP(configStructs.NamespacesLabel, "n", defaultTapConfig.Namespaces, "Namespaces selector")
	tapCmd.Flags().BoolP(configStructs.AllNamespacesLabel, "A", defaultTapConfig.AllNamespaces, "Tap all namespaces")
	tapCmd.Flags().String(configStructs.StorageLimitLabel, defaultTapConfig.StorageLimit, "Override the default storage limit. (per node)")
	tapCmd.Flags().Bool(configStructs.DryRunLabel, defaultTapConfig.DryRun, "Preview of all pods matching the regex, without tapping them")
	tapCmd.Flags().StringP(configStructs.PcapLabel, "p", defaultTapConfig.Pcap, fmt.Sprintf("Capture from a PCAP snapshot of %s (.tar.gz) using your Docker Daemon instead of Kubernetes", misc.Software))
	tapCmd.Flags().Bool(configStructs.ServiceMeshLabel, defaultTapConfig.ServiceMesh, "Capture the encrypted traffic if the cluster is configured with a service mesh and with mTLS")
	tapCmd.Flags().Bool(configStructs.TlsLabel, defaultTapConfig.Tls, "Capture the traffic that's encrypted with OpenSSL or Go crypto/tls libraries")
	tapCmd.Flags().Bool(configStructs.DebugLabel, defaultTapConfig.Debug, "Enable the debug mode")
}
