package main

const (
	Neutral int = 0b0000
	Chaotic int = 0b0001
	Lawful  int = 0b0010
	Evil    int = 0b0100
	Good    int = 0b1000
)

type Monster5e struct {
	Size             string
	Type             string
	Name             string
	Alignment        int
	AC               int
	HP               int
	HPDiceType       int
	HPDiceNumber     int
	HPDiceAdd        int
	Walk             int
	Swim             int
	Fly              int
	Str              int
	Dex              int
	Con              int
	Int              int
	Wis              int
	Cha              int
	SaveStr          int
	SaveDex          int
	SaveCon          int
	SaveInt          int
	SaveWis          int
	SaveCha          int
	Acrobatics       int
	AnimalHanding    int
	Arcana           int
	Athletics        int
	Deception        int
	History          int
	Insight          int
	Intimidation     int
	Investigation    int
	Medicine         int
	Nature           int
	Perception       int
	Performance      int
	Persuasion       int
	Religion         int
	SlightOfHand     int
	Stealth          int
	Survival         int
	Darkvision       int
	Blindsight       int
	Tremorsense      int
	Languages        string
	Telepathy        int
	CR               int
	CRDenominator    int
	CRXP             int
	Abilities        string
	Actions          string
	LegendaryActions string
}

type Dice5e struct {
	DiceType   int
	DiceNumber int
	DiceAdd    int
}
