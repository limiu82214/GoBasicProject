package board_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_domain"
	"github.com/pkg/errors"
)

type setStateUseCase struct {
	loadBoardPort board_application_port_out.ILoadBoardAdapter
}

func NewSetStateUseCase(
	loadBoardPort board_application_port_out.ILoadBoardAdapter,
) board_application_port_in.ISetStateUseCase {
	return &setStateUseCase{
		loadBoardPort: loadBoardPort,
	}
}

func (s *setStateUseCase) SetState(cmd *board_application_port_in.SetStateCmd) error {
	if !cmd.IsValid() {
		panic("檢查是本基")
	}

	board, err := s.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, board_domain.ErrGetEmpty) {
			board = board_domain.NewBoard()
		} else {
			return errors.Wrap(err, "in service setState")
		}
	}

	err = board.SetState(cmd.X, cmd.Y, cmd.S)
	if err != nil {
		return errors.Wrap(err, "in service setState")
	}

	err = s.loadBoardPort.SetBoard(board)
	if err != nil {
		return errors.Wrap(err, "in service setState")
	}

	whoWin := board.WhoWin()
	if whoWin != board_domain.Blank {
		board.ResetBoardState()

		return errors.Errorf("winner is %d", whoWin)
	}

	return nil
}