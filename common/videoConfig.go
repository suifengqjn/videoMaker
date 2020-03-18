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
