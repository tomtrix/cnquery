package builder

// FIXME: DEPRECATED, remove in v9.0 vv
// Yes...the entire file
import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.mondoo.com/cnquery/apps/cnquery/cmd/builder/common"
)

func NewSlimProviderCommand(opts CommandOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:     opts.Use,
		Aliases: opts.Aliases,
		Short:   opts.Short,
		Long:    opts.Long,
		PreRun: func(cmd *cobra.Command, args []string) {
			if opts.PreRun != nil {
				opts.PreRun(cmd, args)
			}
			if opts.CommonPreRun != nil {
				opts.CommonPreRun(cmd, args)
			}
		},
		ValidArgsFunction: opts.ValidArgsFunction,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				log.Error().Msg("provider " + args[0] + " does not exist")
				cmd.Help()
				os.Exit(1)
			}

			if viper.GetString("inventory-file") != "" {
				// when the user provided an inventory file, users do not need to provide a provider
				opts.Run(cmd, args)
				return
			}

			log.Info().Msg("no provider specified, defaulting to local.\n  Use --help for a list of available providers.")
			opts.Run(cmd, args)
		},
	}
	opts.CommonFlags(cmd)
	buildCmd(cmd, opts.CommonFlags, opts.CommonPreRun, opts.Run, opts.Docs)
	return cmd
}

// CommandOpts is a helper command to create a cobra.Command
type CommandOpts struct {
	Use               string
	Aliases           []string
	Short             string
	Long              string
	Run               common.RunFn
	CommonFlags       common.CommonFlagsFn
	CommonPreRun      common.CommonPreRunFn
	Docs              common.CommandsDocs
	PreRun            func(cmd *cobra.Command, args []string)
	ValidArgsFunction func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)
}

func buildCmd(baseCmd *cobra.Command, commonCmdFlags common.CommonFlagsFn, preRun common.CommonPreRunFn, runFn common.RunFn, docs common.CommandsDocs) {
	containerCmd := common.ContainerProviderCmd(commonCmdFlags, preRun, runFn, docs)
	containerImageCmd := common.ContainerImageProviderCmd(commonCmdFlags, preRun, runFn, docs)
	containerCmd.AddCommand(containerImageCmd)
	containerRegistryCmd := common.ContainerRegistryProviderCmd(commonCmdFlags, preRun, runFn, docs)
	containerCmd.AddCommand(containerRegistryCmd)

	dockerCmd := common.DockerProviderCmd(commonCmdFlags, preRun, runFn, docs)
	dockerImageCmd := common.DockerImageProviderCmd(commonCmdFlags, preRun, runFn, docs)
	dockerCmd.AddCommand(dockerImageCmd)
	dockerContainerCmd := common.DockerContainerProviderCmd(commonCmdFlags, preRun, runFn, docs)
	dockerCmd.AddCommand(dockerContainerCmd)

	// aws subcommand
	awsCmd := common.AwsProviderCmd(commonCmdFlags, preRun, runFn, docs)
	awsEc2 := common.AwsEc2ProviderCmd(commonCmdFlags, preRun, runFn, docs)
	awsCmd.AddCommand(awsEc2)

	awsEc2Connect := common.AwsEc2ConnectProviderCmd(commonCmdFlags, preRun, runFn, docs)
	awsEc2.AddCommand(awsEc2Connect)

	awsEc2EbsCmd := common.AwsEc2EbsProviderCmd(commonCmdFlags, preRun, runFn, docs)
	awsEc2EbsVolumeCmd := common.AwsEc2EbsVolumeProviderCmd(commonCmdFlags, preRun, runFn, docs)
	awsEc2EbsCmd.AddCommand(awsEc2EbsVolumeCmd)
	awsEc2EbsSnapshotCmd := common.AwsEc2EbsSnapshotProviderCmd(commonCmdFlags, preRun, runFn, docs)
	awsEc2EbsCmd.AddCommand(awsEc2EbsSnapshotCmd)
	awsEc2.AddCommand(awsEc2EbsCmd)

	awsEc2Ssm := common.AwsEc2SsmProviderCmd(commonCmdFlags, preRun, runFn, docs)
	awsEc2.AddCommand(awsEc2Ssm)

	// gcp subcommand
	gcpCmd := common.ScanGcpCmd(commonCmdFlags, preRun, runFn, docs)
	gcpGcrCmd := common.ScanGcpGcrCmd(commonCmdFlags, preRun, runFn, docs)
	gcpCmd.AddCommand(gcpGcrCmd)
	gcpCmd.AddCommand(common.ScanGcpOrgCmd(commonCmdFlags, preRun, runFn, docs))
	gcpCmd.AddCommand(common.ScanGcpProjectCmd(commonCmdFlags, preRun, runFn, docs))
	gcpCmd.AddCommand(common.ScanGcpFolderCmd(commonCmdFlags, preRun, runFn, docs))

	// vsphere subcommand
	vsphereCmd := common.VsphereProviderCmd(commonCmdFlags, preRun, runFn, docs)
	vsphereVmCmd := common.VsphereVmProviderCmd(commonCmdFlags, preRun, runFn, docs)
	vsphereCmd.AddCommand(vsphereVmCmd)

	// github subcommand
	githubCmd := common.ScanGithubCmd(commonCmdFlags, preRun, runFn, docs)
	githubOrgCmd := common.GithubProviderOrganizationCmd(commonCmdFlags, preRun, runFn, docs)
	githubCmd.AddCommand(githubOrgCmd)
	githubRepositoryCmd := common.GithubProviderRepositoryCmd(commonCmdFlags, preRun, runFn, docs)
	githubCmd.AddCommand(githubRepositoryCmd)
	githubUserCmd := common.GithubProviderUserCmd(commonCmdFlags, preRun, runFn, docs)
	githubCmd.AddCommand(githubUserCmd)

	// terraform subcommand
	terraformCmd := common.TerraformProviderCmd(commonCmdFlags, preRun, runFn, docs)
	terraformPlanCmd := common.TerraformProviderPlanCmd(commonCmdFlags, preRun, runFn, docs)
	terraformCmd.AddCommand(terraformPlanCmd)
	terraformStateCmd := common.TerraformProviderStateCmd(commonCmdFlags, preRun, runFn, docs)
	terraformCmd.AddCommand(terraformStateCmd)

	// subcommands
	baseCmd.AddCommand(common.LocalProviderCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(common.MockProviderCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(common.VagrantCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(terraformCmd)
	baseCmd.AddCommand(common.SshProviderCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(common.WinrmProviderCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(containerCmd)
	baseCmd.AddCommand(dockerCmd)
	baseCmd.AddCommand(common.KubernetesProviderCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(awsCmd)
	baseCmd.AddCommand(common.AzureProviderCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(gcpCmd)
	baseCmd.AddCommand(vsphereCmd)
	baseCmd.AddCommand(githubCmd)
	baseCmd.AddCommand(common.GitlabProviderCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(common.Ms365ProviderCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(common.HostProviderCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(common.AristaProviderCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(common.ScanOktaCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(common.ScanGoogleWorkspaceCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(common.ScanSlackCmd(commonCmdFlags, preRun, runFn, docs))
	baseCmd.AddCommand(common.ScanVcdCmd(commonCmdFlags, preRun, runFn, docs))
}

// ^^