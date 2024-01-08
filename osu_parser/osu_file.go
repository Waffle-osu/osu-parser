package osu_parser

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type OsuFile struct {
	Version int32

	General      GeneralSection
	Editor       EditorSection
	Metadata     MetadataSection
	Difficulty   DifficultySection
	Events       EventsSection
	TimingPoints TimingPointSection
	HitObjects   HitObjectsSection

	ParserWarnings []string
}

func ParseFile(filename string) (OsuFile, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return OsuFile{}, err
	}

	return ParseText(string(data))
}

const (
	SectionGeneral      = 0
	SectionEditor       = 1
	SectionMetadata     = 2
	SectionDifficulty   = 3
	SectionEvents       = 4
	SectionTimingPoints = 5
	SectionHitObjects   = 6
)

func ParseText(osuText string) (OsuFile, error) {
	returnOsuFile := OsuFile{}

	lines := strings.Split(osuText, "\n")
	currentSection := SectionGeneral

	version := strings.Replace(lines[0], "osu file format v", "", -1)
	versionParsed, verionParseErr := strconv.ParseInt(version, 10, 64)

	if verionParseErr != nil {
		return OsuFile{}, verionParseErr
	}

	returnOsuFile.Version = int32(versionParsed)

	addWarning := func(line int, key string, err string) {
		returnOsuFile.ParserWarnings = append(returnOsuFile.ParserWarnings, fmt.Sprintf("Line %d: Error Parsing %s: %s", line, key, err))
	}

	parseInt := func(line int, key string, value string, ret *int32) {
		parsed, parseErr := strconv.ParseInt(value, 10, 64)

		if parseErr != nil {
			addWarning(line, key, parseErr.Error())
		}

		*ret = int32(parsed)
	}

	parseDouble := func(line int, key string, value string, ret *float64) {
		parsed, parseErr := strconv.ParseFloat(value, 64)

		if parseErr != nil {
			addWarning(line, key, parseErr.Error())
		}

		*ret = parsed
	}

	for i := 1; i != len(lines); i++ {
		line := strings.Trim(lines[i], "\t\r ")

		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "//") {
			continue
		}

		switch line {
		case "[General]":
			currentSection = SectionGeneral
			continue
		case "[Editor]":
			currentSection = SectionEditor
			continue
		case "[Metadata]":
			currentSection = SectionMetadata
			continue
		case "[Difficulty]":
			currentSection = SectionDifficulty
			continue
		case "[Events]":
			currentSection = SectionEvents
			continue
		case "[TimingPoints]":
			currentSection = SectionTimingPoints
			continue
		case "[HitObjects]":
			currentSection = SectionHitObjects
			continue
		}

		key := ""
		value := ""

		splitKv := strings.Split(line, ":")

		if len(splitKv) > 0 {
			key = splitKv[0]
		}

		if len(splitKv) > 1 {
			value = strings.Trim(splitKv[1], " ")
		}

		switch currentSection {
		case SectionGeneral:
			general := &returnOsuFile.General

			switch key {
			case "AudioFilename":
				general.AudioFilename = value
			case "AudioLeadIn":
				parseInt(i, key, value, &general.AudioLeadIn)
			case "AudioHash":
				general.AudioHash = value
			case "PreviewTime":
				parseInt(i, key, value, &general.PreviewTime)
			case "SampleSet":
				switch value {
				case "Normal":
					general.SampleSet = SampleSetNormal
				case "Soft":
					general.SampleSet = SampleSetSoft
				case "Drum":
					general.SampleSet = SampleSetDrum
				}
			case "StackLeniency":
				parseDouble(i, key, value, &general.StackLeniency)
			case "Mode":
				mode := int32(0)

				parseInt(i, key, value, &mode)

				general.Mode = Playmode(mode)
			case "LetterboxInBreaks":
				general.LetterboxInBreaks = value == "1"
			case "WidescreenStoryboard":
				general.WidescreenStoryboard = value == "1"
			case "AlwaysShowPlayfield":
				general.AlwaysShowPlayfield = value == "1"
			case "EpilepsyWarning":
				general.EpilepsyWarning = value == "1"
			case "SamplesMatchPlaybackRate":
				general.SamplesMatchPlaybackRate = value == "1"
			case "Countdown":
				parseInt(i, key, value, &general.Countdown)
			case "CountdownOffset":
				parseInt(i, key, value, &general.CountdownOffset)
			case "SampleVolume":
				parseInt(i, key, value, &general.SampleVolume)
			case "SkinPreference":
				general.SkinPreference = value
			case "TimelineZoom":
				parseDouble(i, key, value, &general.TimelineZoom)
			}
		case SectionEditor:
			editor := &returnOsuFile.Editor

			switch key {
			case "DistanceSpacing":
				parseDouble(i, key, value, &editor.DistanceSpacing)
			case "BeatDivisor":
				parseInt(i, key, value, &editor.BeatDivisor)
			case "GridSize":
				parseInt(i, key, value, &editor.GridSize)
			case "TimelineZoom":
				parseDouble(i, key, value, &editor.TimelineZoom)
			}
		case SectionMetadata:
			metadata := &returnOsuFile.Metadata

			switch key {
			case "Artist":
				metadata.Artist = value
			case "ArtistUnicode":
				metadata.ArtistUnicode = value
			case "Title":
				metadata.Title = value
			case "TitleUnicode":
				metadata.TitleUnicode = value
			case "Creator":
				metadata.Creator = value
			case "Version":
				metadata.Version = value
			case "Tags":
				metadata.Tags = value
			case "Source":
				metadata.Source = value
			case "BeatmapID":
				parseInt(i, key, value, &metadata.BeatmapID)
			case "BeatmapSetID":
				parseInt(i, key, value, &metadata.BeatmapSetID)
			}
		case SectionDifficulty:
			difficulty := &returnOsuFile.Difficulty

			//We floor the value if the file version is below 13
			//as version 13 introduced decimal difficulty settings
			actualValue := 0.0
			parsed, parseErr := strconv.ParseFloat(value, 10)

			if parseErr != nil {
				addWarning(i, key, parseErr.Error())
			} else {
				if returnOsuFile.Version < 13 {
					actualValue = math.Floor(parsed)
				} else {
					actualValue = parsed
				}
			}

			switch key {
			case "HPDrainRate":
				difficulty.HPDrainRate = actualValue
			case "CircleSize":
				difficulty.CircleSize = actualValue
			case "OverallDifficulty":
				difficulty.OverallDifficulty = actualValue
			case "ApproachRate":
				difficulty.ApproachRate = actualValue
			case "SliderMultiplier":
				parseDouble(i, key, value, &difficulty.SliderMultiplier)
			case "SliderTickRate":
				parseDouble(i, key, value, &difficulty.SliderTickRate)
			}
		case SectionEvents:
			if len(line) == 0 {
				continue
			}

			events := &returnOsuFile.Events

			split := strings.Split(line, ",")

			eventType := int32(0)
			time := int32(0)

			parseInt(i, key, split[0], &eventType)
			parseInt(i, key, split[1], &time)

			switch EventType(eventType) {
			case EventTypeVideo:
				fallthrough
			case EventTypeBackground:
				backgroundImage := strings.Trim(split[2], " ")

				events.Events = append(events.Events, Event{
					EventType:       EventType(eventType),
					EventTime:       int32(time),
					BackgroundImage: backgroundImage,
				})
			case EventTypeBreak:
				breakStart := time
				breakEnd := int32(0)

				parseInt(i, key, split[2], &breakEnd)

				events.Events = append(events.Events, Event{
					EventType:      EventTypeBreak,
					BreakTimeBegin: breakStart,
					BreakTimeEnd:   breakEnd,
				})
			}
		case SectionTimingPoints:
			timing := &returnOsuFile.TimingPoints

			split := strings.Split(line, ",")
			lenSplit := len(split)

			if lenSplit > 2 {
				offset := 0.0
				beatLength := 0.0
				timeSignature := TimeSignatureQuadruple
				sampleSet := SampleSetNormal
				customSampleSet := CustomSampleSetNone
				sampleVolume := int32(100)
				inheritedTimingPoint := false
				special := SpecialNone

				parseDouble(i, key, split[0], &offset)
				parseDouble(i, key, split[1], &beatLength)

				switch split[2] {
				case "0":
					timeSignature = TimeSignatureQuadruple
				case "1":
					timeSignature = TimeSignatureTriplet
				case "5":
					timeSignature = TimeSignature5
				case "6":
					timeSignature = TimeSignature6
				case "7":
					timeSignature = TimeSignature7
				}

				if lenSplit > 4 {
					switch split[4] {
					case "0":
						customSampleSet = CustomSampleSetNone
					case "1":
						customSampleSet = CustomSampleSet1
					case "2":
						customSampleSet = CustomSampleSet2
					}
				}

				if lenSplit > 5 {
					parseInt(i, key, split[5], &sampleVolume)
				}

				if lenSplit > 6 {
					inheritedTimingPoint = split[6] != "1"
				}

				if lenSplit > 7 {
					switch split[7] {
					case "0":
						special = SpecialNone
					case "1":
						special = SpecialKiai
					case "8":
						special = SpecialTaikoOmitBarLine
					}
				}

				timing.TimingPoints = append(timing.TimingPoints, TimingPoint{
					Offset:               offset,
					BeatLength:           beatLength,
					TimeSignature:        timeSignature,
					SampleSet:            sampleSet,
					CustomSampleSet:      customSampleSet,
					Volume:               sampleVolume,
					InheritedTimingPoint: inheritedTimingPoint,
					SpecialFlag:          special,
				})
			} else if lenSplit == 2 {
				offset := 0.0
				beatLength := 0.0

				parseDouble(i, key, split[0], &offset)
				parseDouble(i, key, split[1], &beatLength)

				timing.TimingPoints = append(timing.TimingPoints, TimingPoint{
					Offset:               offset,
					BeatLength:           beatLength,
					TimeSignature:        TimeSignatureQuadruple,
					SampleSet:            returnOsuFile.General.SampleSet,
					CustomSampleSet:      CustomSampleSetNone,
					Volume:               100,
					InheritedTimingPoint: false,
					SpecialFlag:          SpecialNone,
				})
			} else {
				addWarning(i, "[TimingPoints]", "Incorrect formatting of timing point.")
			}
		case SectionHitObjects:
			hitObjects := &returnOsuFile.HitObjects

			split := strings.Split(line, ",")

			posX := 0.0
			posY := 0.0
			time := 0.0
			hitObjectTypeInt := int32(0)
			hitSoundInt := int32(0)
			comboColorOffset := int32(0)
			newCombo := false

			parseDouble(i, "HitObjects", split[0], &posX)
			parseDouble(i, "HitObjects", split[1], &posY)
			parseDouble(i, "HitObjects", split[2], &time)
			parseInt(i, "HitObjects", split[3], &hitObjectTypeInt)
			parseInt(i, "HitObjects", split[4], &hitSoundInt)

			comboColorOffset = (hitObjectTypeInt >> 4) & 7
			newCombo = (hitObjectTypeInt & 4) > 0
			hitSound := HitSoundType(hitSoundInt)
			toSwitchHitObjectType := HitObjectTypeCircle

			if (hitObjectTypeInt & int32(HitObjectTypeCircle)) > 0 {
				toSwitchHitObjectType = HitObjectTypeCircle
			}

			if (hitObjectTypeInt & int32(HitObjectTypeSlider)) > 0 {
				toSwitchHitObjectType = HitObjectTypeSlider
			}

			if (hitObjectTypeInt & int32(HitObjectTypeSpinner)) > 0 {
				toSwitchHitObjectType = HitObjectTypeSpinner
			}

			if (hitObjectTypeInt & int32(HitObjectTypeHold)) > 0 {
				toSwitchHitObjectType = HitObjectTypeHold
			}

			if len(split) > 5 && len(split[5]) > 0 {
				switch toSwitchHitObjectType {
				//Circle Specific things
				case HitObjectTypeCircle:
					//Per object hitsounding data

					hitSoundsSplit := strings.Split(split[5], ":")
					lenHsSplit := len(hitSoundsSplit)

					sampleSetInt := int32(0)
					sampleSetAdditionInt := int32(0)
					customSampleSetInt := int32(0)
					volume := int32(0)
					sample := ""

					parseInt(i, "HitObjects: Per-object hitsounds 0", hitSoundsSplit[0], &sampleSetInt)
					parseInt(i, "HitObjects: Per-object hitsounds 1", hitSoundsSplit[1], &sampleSetAdditionInt)

					if lenHsSplit > 2 {
						parseInt(i, "HitObjects: Per-object hitsounds 2", hitSoundsSplit[2], &customSampleSetInt)

						if lenHsSplit > 3 {
							parseInt(i, "HitObjects: Per-object hitsounds 3", hitSoundsSplit[3], &volume)

							if lenHsSplit > 4 {
								sample = hitSoundsSplit[4]
							}
						}
					}

					hitObjects.HitObjects = append(hitObjects.HitObjects, HitObject{
						Position: Vec2{
							X: posX,
							Y: posY,
						},
						Time: time,
						Type: toSwitchHitObjectType,

						ComboColorOffset:  int(comboColorOffset),
						NewCombo:          newCombo,
						HitSound:          hitSound,
						SampleSet:         SampleSet(sampleSetInt),
						SampleSetAddition: SampleSet(sampleSetAdditionInt),
						CustomSampleSet:   CustomSampleSet(customSampleSetInt),
						Volume:            volume,
						SampleFile:        sample,
					})

				case HitObjectTypeSlider:
					curveType := CurveTypeCatmull
					repeatCount := int32(0)
					length := 0.0

					sliderPointsSplit := strings.Split(split[5], "|")

					//Curve type
					switch sliderPointsSplit[0] {
					case "C":
						curveType = CurveTypeCatmull
					case "B":
						curveType = CurveTypeBezier
					case "L":
						curveType = CurveTypeLinear
					case "P":
						curveType = CurveTypePerfect
					}

					sliderPoints := []Vec2{}

					for j := 1; j != len(sliderPointsSplit); j++ {
						posSplit := strings.Split(sliderPointsSplit[j], ":")

						if len(posSplit) < 2 {
							continue
						}

						pointX := 0.0
						pointY := 0.0

						parseDouble(i, "HitObjects Slider: Slider Points", posSplit[0], &pointX)
						parseDouble(i, "HitObjects Slider: Slider Points", posSplit[1], &pointY)

						sliderPoints = append(sliderPoints, Vec2{
							X: pointX,
							Y: pointY,
						})
					}

					if len(split) > 6 {
						parseInt(i, "HitObjects Slider: Slider repeat count", split[6], &repeatCount)
					}

					if len(split) > 7 {
						parseDouble(i, "HitObjects Slider: Slider length", split[7], &length)
					}

					hitSounds := []HitSoundType{}

					//Slider hitsounds
					if len(split) > 8 && len(split[8]) > 0 {
						additions := strings.Split(split[8], "|")
						lenAdditions := len(additions)

						if len(additions) > 0 {
							additionLength := int(math.Min(float64(lenAdditions), float64(repeatCount+1)))

							for j := 0; j != additionLength; j++ {
								sound := int32(0)

								parseInt(i, "HitObjects Slider: Slider per-thing hitsounds", additions[j], &sound)

								hitSounds = append(hitSounds, HitSoundType(sound))
							}

							for j := additionLength; j < int(repeatCount)+1; j++ {
								hitSounds = append(hitSounds, hitSound)
							}
						}
					}

					sampleSets := []SampleSet{}
					sampleSetAdditions := []SampleSet{}

					if len(split) > 9 && len(split[9]) > 0 {
						sampleSetsSplit := strings.Split(split[9], "|")

						if len(sampleSetsSplit) > 1 {
							for _, element := range sampleSetsSplit {
								splitSampleSets := strings.Split(element, ":")

								sampleSetInt := int32(0)
								sampleSetAdditionInt := int32(0)

								parseInt(i, "HitObjects Slider: Slider SampleSets", splitSampleSets[0], &sampleSetInt)
								parseInt(i, "HitObjects Slider: Slider SampleSets", splitSampleSets[1], &sampleSetAdditionInt)

								sampleSets = append(sampleSets, SampleSet(sampleSetInt))
								sampleSetAdditions = append(sampleSetAdditions, SampleSet(sampleSetAdditionInt))
							}
						}
					}

					sampleSetInt := int32(0)
					sampleSetAdditionInt := int32(0)
					customSampleSetInt := int32(0)
					volume := int32(0)
					sample := ""

					if len(split) > 10 {
						sampleDetailsSplit := strings.Split(split[10], ":")
						lenHsSplit := len(sampleDetailsSplit)

						parseInt(i, "HitObjects Slider: Per-object hitsounds 0", sampleDetailsSplit[0], &sampleSetInt)
						parseInt(i, "HitObjects Slider: Per-object hitsounds 1", sampleDetailsSplit[1], &sampleSetAdditionInt)

						if lenHsSplit > 2 {
							parseInt(i, "HitObjects Slider: Per-object hitsounds 2", sampleDetailsSplit[2], &customSampleSetInt)

							if lenHsSplit > 3 {
								parseInt(i, "HitObjects Slider: Per-object hitsounds 3", sampleDetailsSplit[3], &volume)

								if lenHsSplit > 4 {
									sample = sampleDetailsSplit[4]
								}
							}
						}
					}

					hitObjects.HitObjects = append(hitObjects.HitObjects, HitObject{
						Position: Vec2{
							X: posX,
							Y: posY,
						},
						Time: time,
						Type: toSwitchHitObjectType,

						ComboColorOffset:  int(comboColorOffset),
						NewCombo:          newCombo,
						HitSound:          hitSound,
						SampleSet:         SampleSet(sampleSetInt),
						SampleSetAddition: SampleSet(sampleSetAdditionInt),
						CustomSampleSet:   CustomSampleSet(customSampleSetInt),
						Volume:            volume,
						SampleFile:        sample,

						//Slider Specific
						CurveType:          curveType,
						RepeatCount:        repeatCount,
						SliderLength:       length,
						SliderPoints:       sliderPoints,
						SoundTypes:         hitSounds,
						SampleSets:         sampleSets,
						SampleSetAdditions: sampleSetAdditions,
					})
				case HitObjectTypeSpinner:
					endTime := int32(0)

					parseInt(i, "HitObjects Spinner: Spinner End time", split[5], &endTime)

					sampleSetInt := int32(0)
					sampleSetAdditionInt := int32(0)
					customSampleSetInt := int32(0)
					volume := int32(0)
					sample := ""

					if len(split) > 6 {
						sampleDetailsSplit := strings.Split(split[6], ":")
						lenHsSplit := len(sampleDetailsSplit)

						parseInt(i, "HitObjects Spinner: Per-object hitsounds 0", sampleDetailsSplit[0], &sampleSetInt)
						parseInt(i, "HitObjects Spinner: Per-object hitsounds 1", sampleDetailsSplit[1], &sampleSetAdditionInt)

						if lenHsSplit > 2 {
							parseInt(i, "HitObjects Spinner: Per-object hitsounds 2", sampleDetailsSplit[2], &customSampleSetInt)

							if lenHsSplit > 3 {
								parseInt(i, "HitObjects Spinner: Per-object hitsounds 3", sampleDetailsSplit[3], &volume)

								if lenHsSplit > 4 {
									sample = sampleDetailsSplit[4]
								}
							}
						}
					}

					hitObjects.HitObjects = append(hitObjects.HitObjects, HitObject{
						Position: Vec2{
							X: posX,
							Y: posY,
						},
						Time: time,
						Type: toSwitchHitObjectType,

						ComboColorOffset:  int(comboColorOffset),
						NewCombo:          newCombo,
						HitSound:          hitSound,
						SampleSet:         SampleSet(sampleSetInt),
						SampleSetAddition: SampleSet(sampleSetAdditionInt),
						CustomSampleSet:   CustomSampleSet(customSampleSetInt),
						Volume:            volume,
						SampleFile:        sample,
						EndTime:           endTime,
					})
				case HitObjectTypeHold:
					sampleSetInt := int32(0)
					sampleSetAdditionInt := int32(0)
					customSampleSetInt := int32(0)
					volume := int32(0)
					sample := ""
					endTime := int32(0)

					if len(split) > 5 {
						sampleDetailsSplit := strings.Split(split[5], ":")
						lenHsSplit := len(sampleDetailsSplit)

						parseInt(i, "HitObjects Hold: Hold Endtime", sampleDetailsSplit[1], &endTime)
						parseInt(i, "HitObjects Hold: Per-object hitsounds 0", sampleDetailsSplit[1], &sampleSetInt)
						parseInt(i, "HitObjects Hold: Per-object hitsounds 0", sampleDetailsSplit[1], &sampleSetInt)
						parseInt(i, "HitObjects Hold: Per-object hitsounds 1", sampleDetailsSplit[2], &sampleSetAdditionInt)

						if lenHsSplit > 3 {
							parseInt(i, "HitObjects Hold: Per-object hitsounds 2", sampleDetailsSplit[3], &customSampleSetInt)

							if lenHsSplit > 4 {
								parseInt(i, "HitObjects Hold: Per-object hitsounds 3", sampleDetailsSplit[4], &volume)

								if lenHsSplit > 5 {
									sample = sampleDetailsSplit[5]
								}
							}
						}
					}

					hitObjects.HitObjects = append(hitObjects.HitObjects, HitObject{
						Position: Vec2{
							X: posX,
							Y: posY,
						},
						Time: time,
						Type: toSwitchHitObjectType,

						ComboColorOffset:  int(comboColorOffset),
						NewCombo:          newCombo,
						HitSound:          hitSound,
						SampleSet:         SampleSet(sampleSetInt),
						SampleSetAddition: SampleSet(sampleSetAdditionInt),
						CustomSampleSet:   CustomSampleSet(customSampleSetInt),
						Volume:            volume,
						SampleFile:        sample,
						EndTime:           endTime,
					})
				}
			}
		}
	}

	return returnOsuFile, nil
}
