package game

type GameState struct {
	Round       RoundState
	Stats       GameStats
	PlayersInfo map[Player]PlayerInfo
	TeamsInfo   map[Team]TeamInfo
	NextAction  NextAction
}

func newGameState(g Game) GameState {
	return GameState{
		Round:       g.currentRound().state(),
		Stats:       g.stats,
		PlayersInfo: g.playersInfo,
		TeamsInfo:   g.teamsInfo,
		NextAction:  g.nextAction(),
	}
}

func (gs GameState) ViewFor(p Player) GameView {
	return GameView{
		You:           p,
		LeftOpponent:  p.leftOpponent(),
		Teammate:      p.teammate(),
		RightOpponent: p.rightOpponent(),

		YourTeam:      p.team(),
		OpponentsTeam: p.opponentTeam(),

		PlayersInfo: gs.PlayersInfo,
		TeamsInfo:   gs.TeamsInfo,

		Round:      gs.Round.ViewFor(p),
		Stats:      gs.Stats,
		NextAction: gs.NextAction,
	}
}

type GameView struct {
	You           Player
	LeftOpponent  Player
	Teammate      Player
	RightOpponent Player

	YourTeam      Team
	OpponentsTeam Team

	PlayersInfo map[Player]PlayerInfo
	TeamsInfo   map[Team]TeamInfo

	Round      RoundView
	Stats      GameStats
	NextAction NextAction
}
