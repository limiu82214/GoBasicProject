package player_application_port_in

type IPutChessUseCase interface {
	PutChess(cmd *PutChessCmd) error
}