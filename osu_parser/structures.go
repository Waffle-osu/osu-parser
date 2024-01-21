package osu_parser

type Playmode int32

const (
	PlaymodeOsu   Playmode = 0
	PlaymodeTaiko Playmode = 1
	PlaymodeCatch Playmode = 2
	PlaymodeMania Playmode = 3
)

type GeneralSection struct {
	AudioFilename            string
	AudioLeadIn              int32
	AudioHash                string
	PreviewTime              int32
	Countdown                int32
	SampleSet                SampleSet
	StackLeniency            float64
	Mode                     Playmode
	LetterboxInBreaks        bool
	WidescreenStoryboard     bool
	EditorBookmarks          []int32
	EditorDistanceSpacing    float64
	StoryFireInFront         bool
	UseSkinSprites           bool
	SampleVolume             int32
	SkinPreference           string
	AlwaysShowPlayfield      bool
	EpilepsyWarning          bool
	CountdownOffset          int32
	TimelineZoom             float64
	SamplesMatchPlaybackRate bool
}

type EditorSection struct {
	DistanceSpacing float64
	BeatDivisor     int32
	GridSize        int32
	Bookmarks       []int32
	TimelineZoom    float64
}

type MetadataSection struct {
	Title         string
	TitleUnicode  string
	Artist        string
	ArtistUnicode string
	Creator       string
	Version       string
	Source        string
	Tags          string
	BeatmapID     int32
	BeatmapSetID  int32
}

type DifficultySection struct {
	HPDrainRate       float64
	CircleSize        float64
	OverallDifficulty float64
	ApproachRate      float64
	SliderMultiplier  float64
	SliderTickRate    float64
}

type EventType int32

const (
	EventTypeBackground EventType = 0
	EventTypeVideo      EventType = 1
	EventTypeBreak      EventType = 2
	EventTypeColor      EventType = 3
	EventTypeSprite     EventType = 4
	EventTypeSample     EventType = 5
	EventTypeAnimation  EventType = 6
)

type Color struct {
	R int32
	G int32
	B int32
}

type Vec2 struct {
	X float64
	Y float64
}

type Event struct {
	EventType EventType

	EventTime       int32
	BackgroundImage string
	BreakTimeBegin  int32
	BreakTimeEnd    int32
}

type EventsSection struct {
	Events []Event
}

type HitObjectType int32
type HitSoundType int32
type SampleSet int32
type CustomSampleSet int32
type CurveType int32

const (
	HitObjectTypeCircle         HitObjectType = 1
	HitObjectTypeSlider         HitObjectType = 2
	HitObjectTypeNewCombo       HitObjectType = 4
	HitObjectTypeCircleNewCombo HitObjectType = 5
	HitObjectTypeSliderNewCombo HitObjectType = 6
	HitObjectTypeSpinner        HitObjectType = 8
	HitObjectTypeColorHax       HitObjectType = 112
	HitObjectTypeHold           HitObjectType = 128

	HitSoundTypeNone    HitSoundType = 0
	HitSoundTypeDefault HitSoundType = 1
	HitSoundTypeWhistle HitSoundType = 2
	HitSoundTypeFinish  HitSoundType = 4
	HitSoundTypeClap    HitSoundType = 8

	SampleSetNone   SampleSet = 0
	SampleSetNormal SampleSet = 1
	SampleSetSoft   SampleSet = 2
	SampleSetDrum   SampleSet = 3

	CustomSampleSetNone CustomSampleSet = 0
	CustomSampleSet1    CustomSampleSet = 1
	CustomSampleSet2    CustomSampleSet = 2

	CurveTypeCatmull CurveType = 0
	CurveTypeBezier  CurveType = 1
	CurveTypeLinear  CurveType = 2
	CurveTypePerfect CurveType = 3
)

type HitObject struct {
	Type              HitObjectType
	Position          Vec2
	Time              float64
	NewCombo          bool
	HitSound          HitSoundType
	ComboColorOffset  int
	SampleSet         SampleSet
	SampleSetAddition SampleSet
	CustomSampleSet   CustomSampleSet
	Volume            int32
	SampleFile        string

	//Slider Specific
	CurveType          CurveType
	RepeatCount        int32
	SliderLength       float64
	SliderPoints       []Vec2
	SoundTypes         []HitSoundType
	SampleSets         []SampleSet
	SampleSetAdditions []SampleSet

	//Spinner/Hold specific
	EndTime int32
}

type HitObjectsSection struct {
	CountNormal  int64
	CountSlider  int64
	CountSpinner int64
	CountHold    int64

	List []HitObject
}

type TimeSignature int32
type Special int32

const (
	TimeSignatureQuadruple TimeSignature = 0
	TimeSignatureTriplet   TimeSignature = 1
	TimeSignature5         TimeSignature = 5
	TimeSignature6         TimeSignature = 6
	TimeSignature7         TimeSignature = 7

	SpecialNone             Special = 0
	SpecialKiai             Special = 1
	SpecialTaikoOmitBarLine Special = 8
)

type TimingPoint struct {
	Offset               float64
	BeatLength           float64
	TimeSignature        TimeSignature
	SampleSet            SampleSet
	CustomSampleSet      CustomSampleSet
	Volume               int32
	InheritedTimingPoint bool
	SpecialFlag          Special
}

type TimingPointSection struct {
	TimingPoints []TimingPoint
}
