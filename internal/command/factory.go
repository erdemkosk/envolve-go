package command

import (
	config "github.com/erdemkosk/envolve-go/internal"
)

func CommandFactory(commandType config.CommandType) ICommand {
	if commandType == config.SYNC {
		return &SyncCommand{}
	}
	if commandType == config.SHOW {
		return &ShowCommand{}
	}

	return nil
}
