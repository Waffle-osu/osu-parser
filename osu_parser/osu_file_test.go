package osu_parser_test

import (
	"fmt"
	"testing"

	"github.com/Waffle-osu/osu-parser/osu_parser"
)

func TestOsuParser(t *testing.T) {
	osuFile := `osu file format v9

[General]
AudioFilename: 0254B84A50FB69AB02.mp3
AudioLeadIn: 0
PreviewTime: -1
Countdown: 0
SampleSet: Normal
StackLeniency: 0.7
Mode: 0
LetterboxInBreaks: 1

[Editor]
DistanceSpacing: 1
BeatDivisor: 4
GridSize: 8

[Metadata]
Title:サトリムソウ
Artist:COOL&CREATE
Creator:Furball
Version:Insane
Source:東方地霊殿　～ Subterranean Animism.
Tags:touhou satori Subterranean Animism cool create th11

[Difficulty]
HPDrainRate:6
CircleSize:5
OverallDifficulty:6
ApproachRate:8
SliderMultiplier:1.4
SliderTickRate:1

[Events]
//Background and Video events
0,0,"tapeciarnia.pl-243136_touhou_komeiji_satori.jpg"
//Break Periods
//Storyboard Layer 0 (Background)
//Storyboard Layer 1 (Fail)
//Storyboard Layer 2 (Pass)
//Storyboard Layer 3 (Foreground)
//Storyboard Sound Samples
//Background Colour Transformations
3,100,163,162,255

[TimingPoints]
2692.25806451613,324.324324324324,4,1,0,100,1,0
13070,-76.9230769230769,4,1,0,100,0,0

[Colours]
Combo1 : 128,0,255
Combo2 : 255,128,255
Combo3 : 255,0,255
Combo4 : 255,0,128

[HitObjects]
102,50,2692,2,0,B|132:123|224:105,1,140,4|0
278,73,3178,1,8
312,72,3259,1,0
346,73,3340,1,0
397,121,3503,2,0,B|422:176|419:216,1,70,8|0
403,255,3827,1,0
376,319,3989,2,0,B|301:319,1,70,2|0
243,350,4313,2,0,B|172:329,1,70,2|0
104,311,4638,2,2,B|61:267|54:171,1,140,2|0
151,80,5286,6,0,B|163:133|146:235,1,140,4|0
188,278,5773,1,8
222,277,5854,1,0
256,283,5935,1,0
310,239,6097,2,0,B|323:190|352:164,1,70,8|0
397,136,6421,1,0
332,108,6584,2,0,B|294:114|248:109,1,70,2|0
199,79,6908,2,0,B|161:73|115:78,1,70,2|0
80,127,7232,1,2
218,345,7881,5,4
285,222,8205,1,8
346,256,8367,1,0
346,256,8449,1,0
346,256,8530,2,0,B|390:265|445:256,1,70,0|0
486,207,8854,1,8
388,338,9016,1,0
285,222,9178,2,0,B|278:180|282:142,1,70,2|0
239,289,9503,2,0,B|203:266|177:237,1,70,2|0
273,365,9827,2,2,B|224:372|128:353,1,140,2|0
81,295,10313,1,0
112,64,10476,6,0,B|183:32|174:114|252:81,1,140,4|8
308,41,10962,1,0
444,178,11124,2,0,B|372:146|381:228|303:195,1,140,4|8
257,151,11611,1,0
187,174,11773,2,0,B|108:178,1,70,6|0
216,247,12097,2,0,B|152:295,1,70,6|0
293,226,12421,1,12
324,242,12503,1,0
358,248,12584,1,4
392,245,12665,1,0
421,226,12746,2,0,B|436:146,1,70,12|0
413,50,13070,6,2,B|344:25|325:117|230:61,1,181.999996745586,2|8
56,129,13557,1,0
56,129,13638,1,0
56,129,13719,2,0,B|108:109|175:148,1,90.9999983727932,2|2
308,195,14043,2,0,B|256:215|189:176,1,90.9999983727932,8|0
194,24,14367,1,2
230,103,14530,1,0
144,207,14692,1,10
180,286,14855,1,0
309,194,15016,2,0,B|396:164|417:270|488:210,1,181.999996745586,2|8
512,160,15503,1,0
256,120,15665,6,0,B|160:120,1,90.9999983727932
80,88,15989,1,0
117,268,16151,1,0
117,268,16232,1,0
117,268,16313,1,0
186,211,16476,2,0,B|232:195|298:227,1,90.9999983727932
338,289,16800,1,0
309,137,16962,2,0,B|306:85|316:39,1,90.9999983727932
363,212,17286,2,0,B|409:235|444:267,1,90.9999983727932
264,213,17611,2,0,B|74:216,1,181.999996745586
	
	`

	parsedOsuFile, _ := osu_parser.ParseText(osuFile)

	if len(parsedOsuFile.ParserWarnings) != 0 {
		for _, warning := range parsedOsuFile.ParserWarnings {
			fmt.Printf("%s", warning)
		}
	}
}
