// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_adapter/in/game_adapter_in_goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_adapter/in/game_adapter_in_player"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_adapter/out/game_adapter_out_leveldb"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_application"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_adapter/in/user_adapter_in_goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_adapter/out/user_adapter_out_game"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_adapter/out/user_adapter_out_leveldb"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/leveldb"
)

// Injectors from wire.go:

func InitUserAdapterInGoPrompt() user_adapter_in_goprompt.IUserGopromptAdapter {
	db := leveldb.GetInst()
	iLoadBoardAdapter := game_adapter_out_leveldb.New(db)
	iGetBoardStateUseCase := game_application.NewGetBoardStateUseCase(iLoadBoardAdapter)
	iSetStateUseCase := game_application.NewSetStateUseCase(iLoadBoardAdapter)
	iResetBoardStateUseCase := game_application.NewResetBoardStateUseCase(iLoadBoardAdapter)
	iWhoWinUseCase := game_application.NewWhoWinUseCase(iLoadBoardAdapter)
	iBoardPlayerAdapter := game_adapter_in_player.New(iGetBoardStateUseCase, iSetStateUseCase, iResetBoardStateUseCase, iWhoWinUseCase)
	iBoardAdapter := user_adapter_out_game.New(iBoardPlayerAdapter)
	user_application_port_inIGetBoardStateUseCase := user_application.NewGetBoardStateUseCase(iBoardAdapter)
	iLoadPlayerAdapter := user_adapter_out_leveldb.New(db)
	iPutChessUseCase := user_application.NewPutChessUseCase(iBoardAdapter, iLoadPlayerAdapter)
	iResetBoardUseCase := user_application.NewResetBoardUseCase(iBoardAdapter)
	user_application_port_inIWhoWinUseCase := user_application.NewWhoWinUseCase(iBoardAdapter)
	iSetPlayerInfoUseCase := user_application.NewSetPlayerInfoUseCase(iLoadPlayerAdapter)
	iUserGopromptAdapter := user_adapter_in_goprompt.New(user_application_port_inIGetBoardStateUseCase, iPutChessUseCase, iResetBoardUseCase, user_application_port_inIWhoWinUseCase, iSetPlayerInfoUseCase)
	return iUserGopromptAdapter
}

func InitBoardAdapterInGoPrompt() game_adapter_in_goprompt.IBoardGopromptAdapter {
	db := leveldb.GetInst()
	iLoadBoardAdapter := game_adapter_out_leveldb.New(db)
	iSetStateUseCase := game_application.NewSetStateUseCase(iLoadBoardAdapter)
	iWhoWinUseCase := game_application.NewWhoWinUseCase(iLoadBoardAdapter)
	iGetBoardStateUseCase := game_application.NewGetBoardStateUseCase(iLoadBoardAdapter)
	iResetBoardStateUseCase := game_application.NewResetBoardStateUseCase(iLoadBoardAdapter)
	iBoardGopromptAdapter := game_adapter_in_goprompt.New(iSetStateUseCase, iWhoWinUseCase, iGetBoardStateUseCase, iResetBoardStateUseCase)
	return iBoardGopromptAdapter
}

// wire.go:

var DBSet = wire.NewSet(leveldb.GetInst)

var BoardApplicationSet = wire.NewSet(game_application.NewGetBoardStateUseCase, game_application.NewSetStateUseCase, game_application.NewResetBoardStateUseCase, game_application.NewWhoWinUseCase)

var UserApplicationSet = wire.NewSet(user_application.NewGetBoardStateUseCase, user_application.NewPutChessUseCase, user_application.NewResetBoardUseCase, user_application.NewWhoWinUseCase, user_application.NewSetPlayerInfoUseCase)
