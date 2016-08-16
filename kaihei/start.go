package kaihei

import (
	"fmt"

	"github.com/eris-ltd/eris-cli/chains"
	"github.com/eris-ltd/eris-cli/util"
	srv "github.com/eris-ltd/eris-cli/services"
	"github.com/eris-ltd/eris-cli/definitions"
)

func StartUpEris(do *definitions.Do) error {

	fmt.Println("starting up your services...")

	// start services
	listOfServices := util.ErisContainersByType(definitions.TypeService, false)

	if len(listOfServices) == 0 {
		return fmt.Errorf("no existing services to start")
	}

	names := make([]string, len(listOfServices))
	for i, serviceName := range listOfServices {
		names[i] = serviceName.ShortName
	}

	fmt.Println(names)

	doStart := definitions.NowDo()
	doStart.ServicesSlice = names
	if err := srv.StartService(doStart); err != nil {
		return err
	}

	// start chain
	// doChain.Name    - name of the chain (optional)
	if do.ChainName != ""{
		doChain := definitions.NowDo()
		doChain.Name = do.ChainName

		fmt.Println("starting up your chain...")
		if err := chains.StartChain(doChain); err != nil {
			return err
		}
	}
	
	return nil
}

func ShutUpEris(do *definitions.Do) error {

	fmt.Println("shutting down your services...")

	// start services
	listOfServices := util.ErisContainersByType(definitions.TypeService, false)

	if len(listOfServices) == 0 {
		return fmt.Errorf("no existing services to stop")
	}

	names := make([]string, len(listOfServices))
	for i, serviceName := range listOfServices {
		names[i] = serviceName.ShortName
	}

	fmt.Println(names)

	doStop := definitions.NowDo()
	doStop.Operations.Args = names
	doStop.Timeout = 10
	if err := srv.KillService(doStop); err != nil {
		return err
	}

	// shutdown all chains
	listOfChains := util.ErisContainersByType(definitions.TypeChain, false)

	if len(listOfChains) == 0 {
		return fmt.Errorf("no existing chains to stop")
	}

	namez := make([]string, len(listOfChains))
	for i, chainName := range listOfChains {
		namez[i] = chainName.ShortName
	}

	fmt.Println(namez)

	doStopChain := definitions.NowDo()
	doStopChain.Operations.Args = namez
	doStopChain.Timeout = 10
	if err := chains.KillChain(doStopChain); err != nil {
		return err
	}


	return nil
}
