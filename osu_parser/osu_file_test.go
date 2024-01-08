package osu_parser_test

import (
	"fmt"
	"testing"

	"github.com/Waffle-osu/osu-parser/osu_parser"
)

func TestOsuParser(t *testing.T) {
	osuFile := `osu file format v14

[General]
AudioFilename: audio.mp3
AudioLeadIn: 0
PreviewTime: 12408
Countdown: 0
SampleSet: Normal
StackLeniency: 0.7
Mode: 0
LetterboxInBreaks: 0
WidescreenStoryboard: 1

[Editor]
Bookmarks: 31,1522,13448,37298,49209,61134,76041,99892,111818,135669,147594,159519,174426,198277
DistanceSpacing: 0.3
BeatDivisor: 4
GridSize: 32
TimelineZoom: 2.2

[Metadata]
Title:Corpse Voyage ~ Be of good cheer!
TitleUnicode:Corpse Voyage ~ Be of good cheer!
Artist:ZUN
ArtistUnicode:ZUN
Creator:Furball
Version:parser test
Source:東方地霊殿　～ Subterranean Animism.
Tags:
BeatmapID:0
BeatmapSetID:-1

[Difficulty]
HPDrainRate:6
CircleSize:5
OverallDifficulty:7
ApproachRate:5
SliderMultiplier:1.4
SliderTickRate:1

[Events]
//Background and Video events
0,0,"Touhou.600.856359_waifu2x_noise1_scale4x.jpg",0,0
//Break Periods
//Storyboard Layer 0 (Background)
//Storyboard Layer 1 (Fail)
//Storyboard Layer 2 (Pass)
//Storyboard Layer 3 (Foreground)
//Storyboard Layer 4 (Overlay)
//Storyboard Sound Samples

[TimingPoints]
16.888537380475,372.670807453416,4,1,0,100,1,0
13433,-100,4,1,0,100,0,1
37283,-66.6666666666667,4,1,0,100,0,0
111818,-133.333333333333,4,1,0,100,0,1
134178,-50,4,1,0,100,0,0


[HitObjects]
101,124,16,1,2,0:0:0:0:
174,89,110,38,0,P|224:80|262:122,1,105,8|4,0:0|0:0,0:0:0:0:
303,181,482,2,4,B|347:196|395:212|460:174,1,140,2|8,0:0|0:0,0:0:0:0:
256,192,948,12,2,1693,0:0:0:0:
400,106,1787,6,0,L|363:216,2,105
	
	`

	parsedOsuFile, _ := osu_parser.ParseText(osuFile)

	if len(parsedOsuFile.ParserWarnings) != 0 {
		for _, warning := range parsedOsuFile.ParserWarnings {
			fmt.Printf("%s", warning)
		}
	}
}
