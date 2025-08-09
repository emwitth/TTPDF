package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PdfMapItemType int

const (
	UnknownPdfItemType   PdfMapItemType = -1
	Monster5ePdfItemType PdfMapItemType = 0
)

type PdfMapItemSection struct {
	itemText string
	itemType PdfMapItemType
}

// Strips text of newlines and converts all to lowercase.
func (a *App) normalizeTextForParsing(text string) string {
	// Standardize the text to one line and lowercase
	nonl := strings.ReplaceAll(text, "\n", "")
	safeText := strings.ToLower(nonl)
	return safeText
}

// Looks at an item and tries to determine the pdf map items it contains.
// Returns an array of pdfMapItemSections.
// Type 'unknown' means we have no idea what is going on.
func (a *App) deducePdfMapItemTypes(text string) []PdfMapItemSection {
	safeText := a.normalizeTextForParsing(text)
	// For now, we assume one type per text, but that may change.
	// Likewise, we are starting with just parsing 5e Monsters.
	// This will obviously get more involved as we progress.
	// Honestly, this is prime realestate for an LLM,
	// but I don't feel like dealing with one of those for my little
	// hobby side project...
	startOfMonsterR := regexp.MustCompile(`([a-z]+ ){0,3}(tiny|small|large|medium|huge|gargantuan).+(chaotic|lawful|neutral) (evil|good|neutral)`)
	hpR := regexp.MustCompile(`hit points [0-9]+`)
	strR := regexp.MustCompile(`str [0-9]+`)
	mapItemSections := []PdfMapItemSection{}

	if startOfMonsterR.MatchString(safeText) && hpR.MatchString(safeText) && strR.MatchString(safeText) {
		charCoordsArray := startOfMonsterR.FindAllStringIndex(safeText, -1)

		for i := 0; i < len(charCoordsArray); i++ {
			section := PdfMapItemSection{}
			startIndex := charCoordsArray[i][0]
			endIndex := 0
			if i >= len(charCoordsArray)-1 {
				endIndex = len(safeText)
			} else {
				endIndex = charCoordsArray[i+1][0]
			}
			section.itemText = safeText[startIndex:endIndex]
			section.itemType = Monster5ePdfItemType
			mapItemSections = append(mapItemSections, section)
		}
	}

	return mapItemSections
}

// Takes a string and parses out 5e monster attributes.
// Returns these in a Monster5e struct.
func (a *App) parse5eMonsterText(text string) Monster5e {
	if text == "" {
		return Monster5e{}
	}
	safeText := a.normalizeTextForParsing(text)
	monster := Monster5e{}
	var err error

	// Create Regex objects
	diceR := regexp.MustCompile(`[0-9]+d[0-9]+ (\+|\-) [0-9]+`)
	diceNoModR := regexp.MustCompile(`[0-9]+d[0-9]+`)
	numberR := regexp.MustCompile(`[0-9]+`)
	numberCommaR := regexp.MustCompile(`[0-9,]+`)
	nameR := regexp.MustCompile(`([a-z]+ ){1,3}(tiny|small|large|medium|huge|gargantuan)`)
	sizeR := regexp.MustCompile(`tiny|small|large|medium|huge|gargantuan`)
	typeR := regexp.MustCompile(`(tiny|small|large|medium|huge|gargantuan) [a-z]+`)
	alignmentR := regexp.MustCompile(`(chaotic|lawful|neutral) (evil|good|neutral)`)
	chaoticR := regexp.MustCompile(`chaotic`)
	lawfulR := regexp.MustCompile(`lawful`)
	goodR := regexp.MustCompile(`good`)
	evilR := regexp.MustCompile(`evil`)
	hpR := regexp.MustCompile(`hit points [0-9]+ \([0-9]+d[0-9]+ (\+|\-) [0-9]+\)`)
	hpNoModR := regexp.MustCompile(`hit points [0-9]+ \([0-9]+d[0-9]+\)`)
	speedR := regexp.MustCompile(`speed [0-9]+`)
	swimR := regexp.MustCompile(`swim [0-9]+`)
	flyR := regexp.MustCompile(`fly [0-9]+`)
	strR := regexp.MustCompile(`str [0-9]+`)
	dexR := regexp.MustCompile(`dex [0-9]+`)
	conR := regexp.MustCompile(`con [0-9]+`)
	intR := regexp.MustCompile(`int [0-9]+`)
	wisR := regexp.MustCompile(`wis [0-9]+`)
	chaR := regexp.MustCompile(`cha [0-9]+`)
	savesR := regexp.MustCompile(`saving throws`)
	sstrR := regexp.MustCompile(`str \+[0-9]+`)
	sdexR := regexp.MustCompile(`dex \+[0-9]+`)
	sconR := regexp.MustCompile(`con \+[0-9]+`)
	sintR := regexp.MustCompile(`int \+[0-9]+`)
	swisR := regexp.MustCompile(`wis \+[0-9]+`)
	schaR := regexp.MustCompile(`cha \+[0-9]+`)
	skillsR := regexp.MustCompile(`skills`)
	acrobaticsR := regexp.MustCompile(`acrobatics \+[0-9]+`)
	animalHandingR := regexp.MustCompile(`animal handling \+[0-9]+`)
	arcanaR := regexp.MustCompile(`arcana \+[0-9]+`)
	athleticsR := regexp.MustCompile(`athletics \+[0-9]+`)
	deceptionR := regexp.MustCompile(`deception \+[0-9]+`)
	historyR := regexp.MustCompile(`history \+[0-9]+`)
	insightR := regexp.MustCompile(`insight \+[0-9]+`)
	intimidationR := regexp.MustCompile(`intimidation \+[0-9]+`)
	investigationR := regexp.MustCompile(`investigation \+[0-9]+`)
	medicineR := regexp.MustCompile(`medicine \+[0-9]+`)
	natureR := regexp.MustCompile(`nature \+[0-9]+`)
	perceptionR := regexp.MustCompile(`perception \+[0-9]+`)
	performanceR := regexp.MustCompile(`performance \+[0-9]+`)
	persuasionR := regexp.MustCompile(`persuasion \+[0-9]+`)
	religionR := regexp.MustCompile(`religion \+[0-9]+`)
	slightOfHandR := regexp.MustCompile(`slight of hand \+[0-9]+`)
	stealthR := regexp.MustCompile(`stealth \+[0-9]+`)
	survivalR := regexp.MustCompile(`survival \+[0-9]+`)
	darkvisionR := regexp.MustCompile(`darkvision [0-9]+`)
	blindsightR := regexp.MustCompile(`blindsight [0-9]+`)
	tremorsenseR := regexp.MustCompile(`tremorsense [0-9]+`)
	languagesTelepathyR := regexp.MustCompile(`languages [a-z0-9', ]+ telepathy [0-9]+`)
	languagesChallengeR := regexp.MustCompile(`languages [a-z0-9', ]+ challenge`)
	challengeFractionR := regexp.MustCompile(`challenge 1\/[0-9]+ \([0-9]+ xp\)`)
	challengeR := regexp.MustCompile(`challenge [0-9]+ \([0-9,]+ xp\)`)

	telepathyR := regexp.MustCompile(`telepathy`)

	// Parse out values from string
	// NAME
	nameArray := strings.Split(nameR.FindString(safeText), " ")
	if len(nameArray) > 0 {
		monster.Name = nameArray[0]
		for i := 1; i < len(nameArray)-1; i++ {
			monster.Name = fmt.Sprintf("%s %s", monster.Name, nameArray[i])
		}
	}
	// SIZE
	monster.Size = sizeR.FindString(safeText)
	// TYPE
	typeLine := typeR.FindString(safeText)
	typeArray := strings.Split(typeLine, " ")
	if len(typeArray) == 2 {
		monster.Type = typeArray[1]
	}
	// ALIGNMENT
	if alignmentR.MatchString(safeText) {
		alignmentString := alignmentR.FindString(safeText)
		if chaoticR.MatchString(alignmentString) {
			monster.Alignment = monster.Alignment + Chaotic
		}
		if lawfulR.MatchString(alignmentString) {
			monster.Alignment = monster.Alignment + Lawful
		}
		if goodR.MatchString(alignmentString) {
			monster.Alignment = monster.Alignment + Good
		}
		if evilR.MatchString(alignmentString) {
			monster.Alignment = monster.Alignment + Evil
		}
	}
	// HIT POINTS
	hpLine := "0"
	hpDiceLine := "0d0"
	if hpR.MatchString(safeText) {
		hpLine = hpR.FindString(safeText)
		hpDiceLine = diceR.FindString(safeText)
	} else if hpNoModR.MatchString(safeText) {
		hpLine = hpNoModR.FindString(safeText)
		hpDiceLine = diceNoModR.FindString(safeText)
	}
	monster.HP, err = strconv.Atoi(numberR.FindString(hpLine))
	if err != nil {
		panic(err)
	}
	hpDice, err := a.parseDiceText(hpDiceLine)
	if err != nil {
		panic(err)
	}
	monster.HPDiceNumber = hpDice.DiceNumber
	monster.HPDiceType = hpDice.DiceType
	monster.HPDiceAdd = hpDice.DiceAdd
	// SPEED
	if speedR.MatchString(safeText) {
		walkString := speedR.FindString(safeText)
		walk := numberR.FindString(walkString)
		monster.Walk, err = strconv.Atoi(walk)
		if err != nil {
			panic(err)
		}
	}
	if swimR.MatchString(safeText) {
		swimString := swimR.FindString(safeText)
		swim := numberR.FindString(swimString)
		monster.Swim, err = strconv.Atoi(swim)
		if err != nil {
			panic(err)
		}
	}
	if flyR.MatchString(safeText) {
		flyString := flyR.FindString(safeText)
		fly := numberR.FindString(flyString)
		monster.Fly, err = strconv.Atoi(fly)
		if err != nil {
			panic(err)
		}
	}
	// ABILITY SCORES
	if strR.MatchString(safeText) {
		str := numberR.FindString(strR.FindString(safeText))
		monster.Str, err = strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
	}
	if dexR.MatchString(safeText) {
		dex := numberR.FindString(dexR.FindString(safeText))
		monster.Dex, err = strconv.Atoi(dex)
		if err != nil {
			panic(err)
		}
	}
	if conR.MatchString(safeText) {
		con := numberR.FindString(conR.FindString(safeText))
		monster.Con, err = strconv.Atoi(con)
		if err != nil {
			panic(err)
		}
	}
	if intR.MatchString(safeText) {
		intn := numberR.FindString(intR.FindString(safeText))
		monster.Int, err = strconv.Atoi(intn)
		if err != nil {
			panic(err)
		}
	}
	if wisR.MatchString(safeText) {
		wis := numberR.FindString(wisR.FindString(safeText))
		monster.Wis, err = strconv.Atoi(wis)
		if err != nil {
			panic(err)
		}
	}
	if chaR.MatchString(safeText) {
		cha := numberR.FindString(chaR.FindString(safeText))
		monster.Cha, err = strconv.Atoi(cha)
		if err != nil {
			panic(err)
		}
	}
	// Saving Throws
	if savesR.MatchString(safeText) {
		if sstrR.MatchString(safeText) {
			sstrString := numberR.FindString(sstrR.FindString(safeText))
			sstr, err := strconv.Atoi(sstrString)
			if err != nil {
				panic(err)
			}
			monster.SaveStr = sstr - a.determineModifierFromScore(monster.Str)
		}
		if sdexR.MatchString(safeText) {
			sdexString := numberR.FindString(sdexR.FindString(safeText))
			sdex, err := strconv.Atoi(sdexString)
			if err != nil {
				panic(err)
			}
			monster.SaveDex = sdex - a.determineModifierFromScore(monster.Dex)
		}
		if sconR.MatchString(safeText) {
			sconString := numberR.FindString(sconR.FindString(safeText))
			scon, err := strconv.Atoi(sconString)
			if err != nil {
				panic(err)
			}
			monster.SaveCon = scon - a.determineModifierFromScore(monster.Con)
		}
		if sintR.MatchString(safeText) {
			sintString := numberR.FindString(sintR.FindString(safeText))
			sint, err := strconv.Atoi(sintString)
			if err != nil {
				panic(err)
			}
			monster.SaveInt = sint - a.determineModifierFromScore(monster.Int)
		}
		if swisR.MatchString(safeText) {
			swisString := numberR.FindString(swisR.FindString(safeText))
			swis, err := strconv.Atoi(swisString)
			if err != nil {
				panic(err)
			}
			monster.SaveWis = swis - a.determineModifierFromScore(monster.Wis)
		}
		if schaR.MatchString(safeText) {
			schaString := numberR.FindString(schaR.FindString(safeText))
			scha, err := strconv.Atoi(schaString)
			if err != nil {
				panic(err)
			}
			monster.SaveCha = scha - a.determineModifierFromScore(monster.Cha)
		}
	}
	// SKILLS
	if skillsR.MatchString(safeText) {
		if acrobaticsR.MatchString(safeText) {
			acrobaticsString := numberR.FindString(acrobaticsR.FindString(safeText))
			acrobatics, err := strconv.Atoi(acrobaticsString)
			if err != nil {
				panic(err)
			}
			monster.Acrobatics = acrobatics - a.determineModifierFromScore(monster.Dex)
		}
		if animalHandingR.MatchString(safeText) {
			animalHandlingString := numberR.FindString(animalHandingR.FindString(safeText))
			animalHandling, err := strconv.Atoi(animalHandlingString)
			if err != nil {
				panic(err)
			}
			monster.AnimalHanding = animalHandling - a.determineModifierFromScore(monster.Wis)
		}
		if arcanaR.MatchString(safeText) {
			arcanaString := numberR.FindString(arcanaR.FindString(safeText))
			arcana, err := strconv.Atoi(arcanaString)
			if err != nil {
				panic(err)
			}
			monster.Arcana = arcana - a.determineModifierFromScore(monster.Int)
		}
		if athleticsR.MatchString(safeText) {
			athleticsString := numberR.FindString(athleticsR.FindString(safeText))
			athletics, err := strconv.Atoi(athleticsString)
			if err != nil {
				panic(err)
			}
			monster.Athletics = athletics - a.determineModifierFromScore(monster.Str)
		}
		if deceptionR.MatchString(safeText) {
			deceptionString := numberR.FindString(deceptionR.FindString(safeText))
			deception, err := strconv.Atoi(deceptionString)
			if err != nil {
				panic(err)
			}
			monster.Deception = deception - a.determineModifierFromScore(monster.Cha)
		}
		if historyR.MatchString(safeText) {
			historyString := numberR.FindString(historyR.FindString(safeText))
			history, err := strconv.Atoi(historyString)
			if err != nil {
				panic(err)
			}
			monster.History = history - a.determineModifierFromScore(monster.Int)
		}
		if insightR.MatchString(safeText) {
			insightString := numberR.FindString(insightR.FindString(safeText))
			insight, err := strconv.Atoi(insightString)
			if err != nil {
				panic(err)
			}
			monster.Insight = insight - a.determineModifierFromScore(monster.Wis)
		}
		if intimidationR.MatchString(safeText) {
			intimidationString := numberR.FindString(intimidationR.FindString(safeText))
			intimidation, err := strconv.Atoi(intimidationString)
			if err != nil {
				panic(err)
			}
			monster.Intimidation = intimidation - a.determineModifierFromScore(monster.Cha)
		}
		if investigationR.MatchString(safeText) {
			investigationString := numberR.FindString(investigationR.FindString(safeText))
			investigation, err := strconv.Atoi(investigationString)
			if err != nil {
				panic(err)
			}
			monster.Investigation = investigation - a.determineModifierFromScore(monster.Int)
		}
		if medicineR.MatchString(safeText) {
			medicineString := numberR.FindString(medicineR.FindString(safeText))
			medicine, err := strconv.Atoi(medicineString)
			if err != nil {
				panic(err)
			}
			monster.Medicine = medicine - a.determineModifierFromScore(monster.Wis)
		}
		if natureR.MatchString(safeText) {
			natureString := numberR.FindString(natureR.FindString(safeText))
			nature, err := strconv.Atoi(natureString)
			if err != nil {
				panic(err)
			}
			monster.Nature = nature - a.determineModifierFromScore(monster.Int)
		}
		if perceptionR.MatchString(safeText) {
			perceptionString := numberR.FindString(perceptionR.FindString(safeText))
			perception, err := strconv.Atoi(perceptionString)
			if err != nil {
				panic(err)
			}
			monster.Perception = perception - a.determineModifierFromScore(monster.Wis)
		}
		if performanceR.MatchString(safeText) {
			performanceString := numberR.FindString(performanceR.FindString(safeText))
			performance, err := strconv.Atoi(performanceString)
			if err != nil {
				panic(err)
			}
			monster.Performance = performance - a.determineModifierFromScore(monster.Cha)
		}
		if persuasionR.MatchString(safeText) {
			persuasionString := numberR.FindString(persuasionR.FindString(safeText))
			persuasion, err := strconv.Atoi(persuasionString)
			if err != nil {
				panic(err)
			}
			monster.Persuasion = persuasion - a.determineModifierFromScore(monster.Cha)
		}
		if religionR.MatchString(safeText) {
			religionString := numberR.FindString(religionR.FindString(safeText))
			religion, err := strconv.Atoi(religionString)
			if err != nil {
				panic(err)
			}
			monster.Religion = religion - a.determineModifierFromScore(monster.Int)
		}
		if slightOfHandR.MatchString(safeText) {
			slightOfHandString := numberR.FindString(slightOfHandR.FindString(safeText))
			slightOfHand, err := strconv.Atoi(slightOfHandString)
			if err != nil {
				panic(err)
			}
			monster.SlightOfHand = slightOfHand - a.determineModifierFromScore(monster.Dex)
		}
		if stealthR.MatchString(safeText) {
			stealthString := numberR.FindString(stealthR.FindString(safeText))
			stealth, err := strconv.Atoi(stealthString)
			if err != nil {
				panic(err)
			}
			monster.Stealth = stealth - a.determineModifierFromScore(monster.Dex)
		}
		if survivalR.MatchString(safeText) {
			survivalString := numberR.FindString(survivalR.FindString(safeText))
			survival, err := strconv.Atoi(survivalString)
			if err != nil {
				panic(err)
			}
			monster.Survival = survival - a.determineModifierFromScore(monster.Wis)
		}
	}
	// SENSES
	if darkvisionR.MatchString(safeText) {
		darkvisionString := numberR.FindString(darkvisionR.FindString(safeText))
		monster.Darkvision, err = strconv.Atoi(darkvisionString)
		if err != nil {
			panic(err)
		}
	}
	if blindsightR.MatchString(safeText) {
		blindsightString := numberR.FindString(blindsightR.FindString(safeText))
		monster.Blindsight, err = strconv.Atoi(blindsightString)
		if err != nil {
			panic(err)
		}
	}
	if tremorsenseR.MatchString(safeText) {
		tremorsenseString := numberR.FindString(tremorsenseR.FindString(safeText))
		monster.Tremorsense, err = strconv.Atoi(tremorsenseString)
		if err != nil {
			panic(err)
		}
	}
	// LANGUAGES
	languagesString := ""
	if languagesTelepathyR.MatchString(safeText) {
		languagesString = languagesTelepathyR.FindString(safeText)
	}
	if languagesChallengeR.MatchString(safeText) {
		languagesString = languagesChallengeR.FindString(safeText)
	}
	languagesArray := strings.Split(languagesString, ",")
	// remove "languages " from beginning and " challenge" from end
	if len(languagesArray) != 1 {
		monster.Languages = languagesArray[0][10:len(languagesArray[0])]
	} else if languagesArray[0] != "" {
		monster.Languages = languagesArray[0][10 : len(languagesArray[0])-10]
	}
	for i := 1; i < len(languagesArray); i++ {
		if telepathyR.MatchString(languagesArray[i]) {
			monster.Telepathy, err = strconv.Atoi(numberR.FindString(languagesArray[i]))
			if err != nil {
				panic(err)
			}
		} else if i == len(languagesArray) {
			// telepathy isn't a match, so this ends in "challenge"
			monster.Languages = fmt.Sprintf("%s, %s", monster.Languages, strings.TrimSpace(languagesArray[i])[0:len(languagesArray[0])-10])
		} else {
			monster.Languages = fmt.Sprintf("%s, %s", monster.Languages, strings.TrimSpace(languagesArray[i]))
		}
	}
	// CHALLENGE
	if challengeFractionR.MatchString(safeText) {
		challengeNumberArray := numberR.FindAllString(challengeFractionR.FindString(safeText), -1)
		if len(challengeNumberArray) == 3 {
			numerator, err := strconv.Atoi(challengeNumberArray[0])
			if err != nil {
				panic(err)
			}
			denominator, err := strconv.Atoi(challengeNumberArray[1])
			if err != nil {
				panic(err)
			}
			xp, err := strconv.Atoi(challengeNumberArray[2])
			if err != nil {
				panic(err)
			}
			monster.CR = numerator
			monster.CRDenominator = denominator
			monster.CRXP = xp
		}
	}
	if challengeR.MatchString(safeText) {
		challengeNumberArray := numberCommaR.FindAllString(challengeR.FindString(safeText), -1)
		if len(challengeNumberArray) == 2 {
			cr, err := strconv.Atoi(challengeNumberArray[0])
			if err != nil {
				panic(err)
			}
			xp, err := strconv.Atoi(strings.ReplaceAll(challengeNumberArray[1], ",", ""))
			if err != nil {
				panic(err)
			}
			monster.CR = cr
			monster.CRDenominator = 1
			monster.CRXP = xp
		}
	}

	fmt.Println(monster)
	return monster
}

func (a *App) parseDiceText(text string) (Dice5e, error) {
	dice := Dice5e{}
	var err error
	diceR := regexp.MustCompile(`[0-9]+d[0-9]+`)
	numberR := regexp.MustCompile(`[0-9]+`)
	addR := regexp.MustCompile(`\+ [0-9]+`)
	subtractR := regexp.MustCompile(`\- [0-9]+`)

	if !diceR.MatchString(text) {
		return dice, errors.New("no dice match")
	}
	// if it's a match, we know we have at least two numbers
	diceNumbers := numberR.FindAllString(text, -1)
	dice.DiceNumber, err = strconv.Atoi(diceNumbers[0])
	if err != nil {
		return dice, err
	}
	dice.DiceType, err = strconv.Atoi(diceNumbers[1])
	if err != nil {
		return dice, err
	}

	if addR.MatchString(text) {
		addModifier := addR.FindString(text)
		addModNumberString := numberR.FindString(addModifier)
		dice.DiceAdd, err = strconv.Atoi(addModNumberString)
		if err != nil {
			return dice, err
		}
	} else if subtractR.MatchString(text) {
		subtractModifier := subtractR.FindString(text)
		subtractModNumberString := numberR.FindString(subtractModifier)
		subtractModNumber, err := strconv.Atoi(subtractModNumberString)
		if err != nil {
			return dice, err
		}
		dice.DiceAdd = 0 - subtractModNumber
	} else {
		dice.DiceAdd = 0
	}

	return dice, nil
}

func (a *App) determineModifierFromScore(abilityScore int) int {
	if abilityScore > 10 {
		return (abilityScore - 10) / 2
	}
	switch abilityScore {
	case 9:
		return -1
	case 8:
		return -1
	case 7:
		return -2
	case 6:
		return -2
	case 5:
		return -3
	case 4:
		return -3
	case 3:
		return -4
	case 2:
		return -4
	case 1:
		return -5
	default:
		return -5
	}
}
