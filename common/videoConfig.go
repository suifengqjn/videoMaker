package common



type CutFront struct {
	Switch bool
	Value int
}

type CutBack struct {
	Switch bool
	Value int
}

type ClearWater struct {
	Switch bool
	X int
	Y int
	W int
	H int
}

type ExtractSubtitles struct {
	Switch bool
}

type Composite struct {
	Switch bool
	Style int  //合成模式
	Voice string  // 播音员
	Volume int // 音量
	SpeechRate int //语速
	PitchRate int // 语调
	BreakTime int // 停顿
}


//字幕样式
type Subtitles struct {
	Switch bool
	FontSize int
	FontColor string
	MarginV int
	BjColor string
	BjAlpha int  // 0 - 10  0完全透明
	CoverBj bool  //遮盖原字幕
	CoverH int
}



type WaterText struct {
	Switch bool
	Content string
	Path string
	Size int
	Color string
	Alpha float32
	Style int
	Sp1 int
	Sp2 int
}

type RunWaterText struct {
	Switch  bool
	Content string
	Path    string
	Size    int
	Color   string
	IsTop   int
	LeftToRight   int
	Sp      int
}

type WaterImage struct {
	Switch bool
	Path string
	Style int
	Sp1 int
	Sp2 int
}

type AddBgm struct {
	Switch bool
	Dir string
	Keep int
}

type FilmTitle struct {
	Switch bool
	Path   string
}

type FilmEnd struct {
	Switch bool
	Path   string
}
