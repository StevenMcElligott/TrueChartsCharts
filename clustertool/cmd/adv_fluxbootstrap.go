package cmd

import (
    "context"

    talhelperCfg "github.com/budimanjojo/talhelper/v3/pkg/config"
    "github.com/rs/zerolog/log"
    "github.com/spf13/cobra"
    "github.com/truecharts/public/clustertool/pkg/fluxhandler"
    "github.com/truecharts/public/clustertool/pkg/helper"
    "github.com/truecharts/public/clustertool/pkg/initfiles"
    "github.com/truecharts/public/clustertool/pkg/sops"
)

var fluxbootstrap = &cobra.Command{
    Use:   "fluxbootstrap",
    Short: "Manually bootstrap fluxcd on existing cluster",
    Run: func(cmd *cobra.Command, args []string) {
        ctx := context.Background()

        if err := sops.DecryptFiles(); err != nil {
            log.Info().Msgf("Error decrypting files: %v\n", err)
        }
        initfiles.LoadTalEnv()
        _, err := talhelperCfg.LoadAndValidateFromFile(helper.TalConfigFile, []string{helper.ClusterEnvFile}, false)
        if err != nil {
            log.Fatal().Err(err).Msg("failed to parse talconfig or talenv file: %s")
        }

        fluxhandler.FluxBootstrap(ctx)
    },
}

func init() {
    adv.AddCommand(fluxbootstrap)
}
